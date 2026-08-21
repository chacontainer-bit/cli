<#
.SYNOPSIS
    Instala la bóveda de JARVIS y la capa de voz en Windows.

.DESCRIPTION
    Copia el andamiaje (CLAUDE.md, permisos y las cuatro skills) a la carpeta de
    la bóveda, y prepara un entorno virtual de Python con el bucle de voz.
    No sobrescribe ficheros que ya existan salvo que pases -Sobrescribir.

.EXAMPLE
    powershell -ExecutionPolicy Bypass -File .\instalar.ps1

.EXAMPLE
    powershell -ExecutionPolicy Bypass -File .\instalar.ps1 -Boveda D:\boveda -SinVoz
#>

[CmdletBinding()]
param(
    # Dónde vive la bóveda.
    [string]$Boveda = (Join-Path $HOME 'boveda'),

    # Dónde vive el bucle de voz y su entorno virtual.
    [string]$Destino = (Join-Path $HOME 'jarvis'),

    # Salta la instalación de Python/voz (fases 1 y 2 solamente).
    [switch]$SinVoz,

    # Reemplaza ficheros existentes de la bóveda. Cuidado: pisa tu CLAUDE.md.
    [switch]$Sobrescribir
)

$ErrorActionPreference = 'Stop'
$origen = $PSScriptRoot

function Paso($texto)  { Write-Host "`n==> $texto" -ForegroundColor Cyan }
function Bien($texto)  { Write-Host "    OK   $texto" -ForegroundColor Green }
function Aviso($texto) { Write-Host "    !    $texto" -ForegroundColor Yellow }

# --- 1. Comprobaciones previas ----------------------------------------------

Paso 'Comprobando requisitos'

if (-not (Get-Command claude -ErrorAction SilentlyContinue)) {
    Aviso 'No encuentro el comando "claude".'
    Aviso 'Instálalo con:  npm install -g @anthropic-ai/claude-code'
    Aviso 'Continúo con la bóveda, pero no podrás probarla hasta instalarlo.'
} else {
    Bien "claude encontrado: $((Get-Command claude).Source)"
}

if (-not (Get-Command git -ErrorAction SilentlyContinue)) {
    Aviso 'No encuentro git. Claude Code lo necesita en Windows para su shell.'
    Aviso 'Instálalo con:  winget install Git.Git'
}

# --- 2. Copiar la bóveda ------------------------------------------------------

Paso "Instalando la bóveda en $Boveda"

$plantilla = Join-Path $origen 'boveda'
if (-not (Test-Path $plantilla)) {
    throw "No encuentro la plantilla en $plantilla. ¿Ejecutas el script desde jarvis\scaffold?"
}

New-Item -ItemType Directory -Force -Path $Boveda | Out-Null

Get-ChildItem -Path $plantilla -Recurse -File | ForEach-Object {
    $relativa = $_.FullName.Substring($plantilla.Length).TrimStart('\', '/')

    # dot-claude\ en el repo se convierte en .claude\ en la bóveda
    $relativa = $relativa -replace '^dot-claude', '.claude'

    if ($_.Name -eq '.gitkeep') {
        New-Item -ItemType Directory -Force -Path (Join-Path $Boveda (Split-Path $relativa)) | Out-Null
        return
    }

    $final  = Join-Path $Boveda $relativa
    $carpeta = Split-Path $final
    New-Item -ItemType Directory -Force -Path $carpeta | Out-Null

    if ((Test-Path $final) -and -not $Sobrescribir) {
        Aviso "ya existe, lo dejo como está: $relativa"
    } else {
        Copy-Item $_.FullName -Destination $final -Force
        Bien $relativa
    }
}

foreach ($sub in 'raw', 'wiki', 'outputs') {
    New-Item -ItemType Directory -Force -Path (Join-Path $Boveda $sub) | Out-Null
}

# --- 3. Capa de voz -----------------------------------------------------------

if ($SinVoz) {
    Paso 'Salto la capa de voz (-SinVoz)'
} else {
    Paso "Instalando la capa de voz en $Destino"

    $python = Get-Command py -ErrorAction SilentlyContinue
    if (-not $python) {
        $python = Get-Command python -ErrorAction SilentlyContinue
    }
    if (-not $python) {
        Aviso 'No encuentro Python. Instálalo con:  winget install Python.Python.3.12'
        Aviso 'Después vuelve a ejecutar este script.'
    } else {
        New-Item -ItemType Directory -Force -Path $Destino | Out-Null
        Copy-Item (Join-Path $origen 'jarvis.py')        -Destination $Destino -Force
        Copy-Item (Join-Path $origen 'requirements.txt') -Destination $Destino -Force
        Bien 'jarvis.py y requirements.txt copiados'

        $venv = Join-Path $Destino '.venv'
        if (-not (Test-Path $venv)) {
            & $python.Source -m venv $venv
            Bien 'entorno virtual creado'
        }

        $pip = Join-Path $venv 'Scripts\pip.exe'
        Write-Host '    instalando dependencias (tarda un par de minutos)…'
        & $pip install --quiet --upgrade pip
        & $pip install --quiet -r (Join-Path $Destino 'requirements.txt')
        Bien 'dependencias instaladas'
    }
}

# --- 4. Atajo en el perfil de PowerShell -------------------------------------

Paso 'Creando el atajo "jarvis"'

$exe = Join-Path $Destino '.venv\Scripts\python.exe'
$funcion = @"

# --- JARVIS ---
function jarvis {
    `$env:JARVIS_VAULT = '$Boveda'
    & '$exe' '$(Join-Path $Destino 'jarvis.py')' @args
}
function boveda { Set-Location '$Boveda' }
# --- fin JARVIS ---
"@

if (-not (Test-Path $PROFILE)) {
    New-Item -ItemType File -Force -Path $PROFILE | Out-Null
}

if ((Get-Content $PROFILE -Raw -ErrorAction SilentlyContinue) -match '# --- JARVIS ---') {
    Aviso 'el atajo ya estaba en tu perfil, no lo duplico'
} else {
    Add-Content -Path $PROFILE -Value $funcion -Encoding UTF8
    Bien "añadido a $PROFILE"
}

# --- 5. Resumen ---------------------------------------------------------------

Paso 'Listo'
Write-Host @"

    Bóveda:      $Boveda
    Bucle voz:   $Destino

    Siguientes pasos:

    1. Abre $Boveda\wiki\contexto.md y rellénalo. Es lo que más rinde.
    2. Abre Obsidian -> Open folder as vault -> $Boveda
    3. Prueba el cerebro sin voz:
           boveda
           claude -p "apunta que hay que renovar el seguro en octubre"
           claude -p "que hago hoy"
    4. Abre una terminal NUEVA (para cargar el atajo) y prueba la voz:
           jarvis --voces     # comprueba que hay una voz en español
           jarvis --texto     # prueba el bucle escribiendo
           jarvis             # prueba el bucle hablando

"@ -ForegroundColor Gray

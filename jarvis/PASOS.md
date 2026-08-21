# JARVIS — Pasos para el ordenador (Windows nativo)

Guía de ejecución en **Windows sin WSL**: PowerShell, winget, voz SAPI y
Programador de tareas.

Cada fase termina con una **prueba**. No pases a la siguiente sin superarla.

> **Atajo:** las fases 1 y 2 ya están construidas en
> [`scaffold/`](./scaffold/). Si quieres saltarte el copiar y pegar, ve directo
> a [«Instalación rápida»](#instalación-rápida) y vuelve luego a leer las fases
> para entender qué se ha instalado.

En toda la guía la bóveda vive en `C:\Users\<tú>\boveda`, que en PowerShell se
escribe `$HOME\boveda`.

---

## Instalación rápida

Con los requisitos de la fase 0 ya puestos:

```powershell
git clone https://github.com/chacontainer-bit/cli.git $HOME\cli
cd $HOME\cli\jarvis\scaffold
powershell -ExecutionPolicy Bypass -File .\instalar.ps1
```

Eso deja la bóveda en `$HOME\boveda` con `CLAUDE.md`, permisos y las cuatro
skills, monta el entorno de Python en `$HOME\jarvis` y añade los atajos `jarvis`
y `boveda` a tu perfil de PowerShell. No pisa nada que ya exista.

Opciones: `-Boveda D:\ruta`, `-SinVoz` (solo fases 1 y 2), `-Sobrescribir`.

Luego sigue desde la [prueba de la fase 2](#prueba-de-la-fase-2).

---

## Fase 0 — Requisitos (30 min)

### 0.1 Instalar

Abre **PowerShell como administrador** y pega:

```powershell
winget install --id Git.Git -e
winget install --id OpenJS.NodeJS.LTS -e
winget install --id Python.Python.3.12 -e
winget install --id Obsidian.Obsidian -e
```

Git no es opcional: Claude Code lo usa en Windows para su herramienta de shell.

**Cierra y vuelve a abrir PowerShell** (ahora ya sin administrador) para que
tome el PATH nuevo.

### 0.2 Instalar y autenticar Claude Code

```powershell
npm install -g @anthropic-ai/claude-code
claude
```

Se abre el navegador para el login. Cuando termine, `/exit`.

### 0.3 Permitir scripts

Solo hace falta si vas a usar `instalar.ps1`:

```powershell
Set-ExecutionPolicy -Scope CurrentUser RemoteSigned
```

### 0.4 Voz en español de Windows

Comprueba que tienes una voz `es-ES` o `es-MX` instalada:

```powershell
Add-Type -AssemblyName System.Speech
(New-Object System.Speech.Synthesis.SpeechSynthesizer).GetInstalledVoices() |
  ForEach-Object { $_.VoiceInfo.Name + '  [' + $_.VoiceInfo.Culture.Name + ']' }
```

Si no aparece ninguna con `[es-...]`:
**Configuración → Hora e idioma → Idioma y región → Español → ⋯ → Opciones de
idioma → Voz → Descargar.** Reinicia la terminal después.

### 0.5 Permiso de micrófono

**Configuración → Privacidad y seguridad → Micrófono →** activa *"Permitir que
las aplicaciones de escritorio accedan al micrófono"*. Si está apagado, Python
grabará silencio sin dar ningún error, y perderás media tarde buscando el fallo.

**Prueba de la fase 0:**
```powershell
claude --version
claude -p "responde solo: ok"
```

---

## Fase 1 — Cablea el cerebro y la memoria (30 min)

Qué instala el andamiaje, y por qué.

### 1.1 La bóveda

```powershell
New-Item -ItemType Directory -Force $HOME\boveda\raw, $HOME\boveda\wiki, `
    $HOME\boveda\outputs, $HOME\boveda\.claude\skills | Out-Null
cd $HOME\boveda
git init      # opcional, pero tener historial de tu memoria vale mucho
```

Abre Obsidian → *Open folder as vault* → `C:\Users\<tú>\boveda`.

### 1.2 `boveda\CLAUDE.md`

Se carga en cada sesión: es lo que convierte a Claude Code en JARVIS. Fija el
tono (frases cortas, para leer en voz alta), el mapa de carpetas y las reglas
duras (buscar antes de responder, no inventar enlaces, no rellenar huecos).

→ Fichero listo: [`scaffold/boveda/CLAUDE.md`](./scaffold/boveda/CLAUDE.md)

### 1.3 `boveda\.claude\settings.json`

Sin esto, el bucle de voz se queda colgado pidiendo confirmación. Con esto,
JARVIS puede leer toda la bóveda y escribir **solo** en `raw/` y `outputs/`.
`wiki/` queda denegado por escrito.

→ Fichero listo: [`scaffold/boveda/dot-claude/settings.json`](./scaffold/boveda/dot-claude/settings.json)

> Nunca uses `--permission-mode bypassPermissions` en el bucle de voz. El
> dictado se equivoca, y con permisos totales ejecutaría lo que *crea* haber
> oído.

### 1.4 `boveda\wiki\contexto.md`

Quién eres, proyectos activos, objetivos del trimestre, cómo quieres que te
hable, nombres propios que el dictado suele fallar. **Es la nota más rentable
del sistema entero**: quince minutos aquí ahorran una pregunta tonta en cada
conversación.

→ Plantilla: [`scaffold/boveda/wiki/contexto.md`](./scaffold/boveda/wiki/contexto.md)

**Prueba de la fase 1:**
```powershell
cd $HOME\boveda
claude -p "que sabes de mi? responde en dos frases"
```
Debe responder desde `wiki/contexto.md`, no en genérico. Si responde en
genérico, es que no has rellenado la nota.

---

## Fase 2 — Las skills (1–2 h)

Cada skill es una carpeta con un `SKILL.md`. El `description` del frontmatter es
lo que hace que Claude la elija solo: **escríbelo con las palabras que dirías en
voz alta**, no con las que usarías en documentación.

Las cinco del andamiaje, todas sin integraciones externas:

| Skill | Se dispara con | Qué hace |
|---|---|---|
| [`plan-hoy`](./scaffold/boveda/dot-claude/skills/plan-hoy/SKILL.md) | "qué hago hoy", "prioridades", "empecemos el día" | Busca lo abierto en `raw/`, propone 3 prioridades, las escribe |
| [`cierre-dia`](./scaffold/boveda/dot-claude/skills/cierre-dia/SKILL.md) | "cierra el día", "qué he hecho hoy" | Te pregunta prioridad a prioridad y escribe el cierre en `outputs/` |
| [`cierre-auto`](./scaffold/boveda/dot-claude/skills/cierre-auto/SKILL.md) | "cierre automático" (tarea programada) | El mismo cierre pero sin preguntar nada, solo con lo observable |
| [`nota-rapida`](./scaffold/boveda/dot-claude/skills/nota-rapida/SKILL.md) | "apunta", "anota", "recuérdame" | Captura al fichero del día sin interrumpirte |
| [`buscar-boveda`](./scaffold/boveda/dot-claude/skills/buscar-boveda/SKILL.md) | "qué dije sobre", "dónde apunté" | Grep, lee 3 ficheros como mucho, responde con fecha |

Instalación manual (si no usaste `instalar.ps1`): copia cada carpeta de
`scaffold/boveda/dot-claude/skills/` a `$HOME\boveda\.claude\skills\`.

```powershell
Copy-Item -Recurse -Force `
  $HOME\cli\jarvis\scaffold\boveda\dot-claude\skills\* `
  $HOME\boveda\.claude\skills\
```

### Prueba de la fase 2

Todavía **sin voz**. Así separas los fallos de skill de los fallos de micro:

```powershell
cd $HOME\boveda
claude -p "apunta que hay que renovar el seguro del almacen en octubre"
claude -p "que hago hoy"
claude -p "que dije sobre el seguro"
```

Las tres deben disparar la skill correcta y dejar ficheros en `raw\`.
Ábrelos en Obsidian y compruébalo.

**Si una skill no se dispara**, el problema está en su `description`: añádele
más maneras de pedir lo mismo. Es el ajuste que más veces vas a repetir.

---

## Fase 3 — La voz (1–2 h)

### 3.1 Entorno

```powershell
mkdir $HOME\jarvis -Force
cd $HOME\jarvis
py -m venv .venv
.\.venv\Scripts\Activate.ps1
pip install faster-whisper sounddevice numpy
```

La primera ejecución descarga el modelo de Whisper (unos 500 MB para `small`) y
tarda. Las siguientes arrancan en segundos.

### 3.2 El bucle

→ Fichero listo: [`scaffold/jarvis.py`](./scaffold/jarvis.py)

```powershell
Copy-Item $HOME\cli\jarvis\scaffold\jarvis.py $HOME\jarvis\
```

Qué hace, en orden: graba mientras hablas → transcribe en local con
faster-whisper → lanza `claude -p` con `cwd` en la bóveda → lee la respuesta en
voz alta con la voz SAPI de Windows. Sin `--continue`: cada comando es una
sesión limpia, y la memoria vive en la bóveda.

Tres modos:

```powershell
python jarvis.py --voces    # lista las voces instaladas y sale
python jarvis.py --texto    # escribes en vez de hablar (para depurar)
python jarvis.py            # el bucle completo
python jarvis.py --modelo base   # más rápido, algo menos preciso
```

### 3.3 Atajo permanente

```powershell
notepad $PROFILE
```

Añade al final (ajusta las rutas si cambiaste algo):

```powershell
function jarvis {
    $env:JARVIS_VAULT = "$HOME\boveda"
    & "$HOME\jarvis\.venv\Scripts\python.exe" "$HOME\jarvis\jarvis.py" @args
}
function boveda { Set-Location "$HOME\boveda" }
```

Abre una terminal **nueva** y ya tienes `jarvis` y `boveda` en cualquier sitio.

**Prueba de la fase 3:** ejecuta `jarvis`, di *"apunta que mañana hay que llamar
al transportista"*, y comprueba que responde en voz alta y que la nota aparece
en `boveda\raw\`.

### 3.4 Si algo falla

| Síntoma | Causa casi siempre | Arreglo |
|---|---|---|
| Graba silencio, no te oye | Permiso de micro de apps de escritorio | Fase 0.5 |
| `PortAudioError` al abrir el micro | Dispositivo por defecto equivocado | `python -c "import sounddevice; print(sounddevice.query_devices())"` y fija `sd.default.device = <n>` en `jarvis.py` |
| Habla en inglés o con acento raro | No hay voz `es-*` instalada | Fase 0.4, o `jarvis --voces` para ver cuál coge |
| Acentos rotos en la consola | Página de códigos | `chcp 65001` antes de lanzarlo |
| "No encuentro el comando claude" | PATH sin refrescar | Cierra y reabre PowerShell |
| Tarda más de 25 s por respuesta | Una skill está leyendo demasiado | Acota la skill: Grep primero, tres ficheros máximo |

---

## Fase 4 — El HUD (una tarde)

Dos caminos. **Recomiendo el A**, porque el stack ya lo tienes montado en este
mismo repo.

### Camino A — ruta `/hud` en `chacontainer/web`

Ya tienes React + Vite + Tailwind en `chacontainer/web` y un servidor Go en
`chacontainer/cmd/server`. El HUD es una ruta más, no un proyecto aparte.

1. Endpoint `GET /api/vault` en el servidor Go que lea la carpeta de la bóveda y
   devuelva: las 10 notas más recientes (nombre, fecha, primeras 200 letras), el
   plan de hoy si existe, y conteos (notas totales, de esta semana, tamaño).
2. Ruta `/hud` con cuatro paneles: vitales, panel de comandos, agenda de hoy,
   últimas notas.
3. Que escuche **solo en `127.0.0.1`**. Es tu máquina; no publiques tu memoria.

Prompt para arrancarlo, desde la raíz del repo:

```
Añade un HUD en chacontainer/web: ruta /hud, tema terminal oscuro, cuatro
paneles — vitales del sistema, panel de comandos, agenda de hoy y últimas notas.
Los datos vienen de un endpoint nuevo GET /api/vault en cmd/server que lee la
carpeta de la bóveda (ruta configurable por variable de entorno JARVIS_VAULT).
Sin dependencias nuevas de frontend: usa el Tailwind que ya está configurado.
El servidor escucha solo en 127.0.0.1.
```

### Camino B — un solo fichero

`hud.html` servido con `python -m http.server 8787 --directory $HOME\boveda`.
Cero dependencias, cero integración, lo tienes hoy mismo.

**Prueba de la fase 4:** abres el HUD y ves el plan de hoy que generaste en la
fase 2.

---

## Fase 5 — Rutinas y crecimiento (continuo)

### 5.1 Tareas programadas

En Windows, con `schtasks`. Ojo: la tarea necesita ejecutarse **dentro** de la
bóveda, así que va envuelta en un `cmd /c`:

```powershell
# Resumen matutino, laborables a las 7:00
schtasks /create /tn "JARVIS resumen matutino" /sc weekly /d MON,TUE,WED,THU,FRI /st 07:00 `
  /tr "cmd /c cd /d %USERPROFILE%\boveda && claude -p \"resumen matutino\" >> outputs\cron.log 2>&1"

# Cierre automático, laborables a las 19:00
schtasks /create /tn "JARVIS cierre" /sc weekly /d MON,TUE,WED,THU,FRI /st 19:00 `
  /tr "cmd /c cd /d %USERPROFILE%\boveda && claude -p \"cierre automatico\" >> outputs\cron.log 2>&1"
```

> Fíjate en que la tarea de la tarde llama a **`cierre-auto`**, no a
> `cierre-dia`. La segunda es conversacional: te pregunta prioridad por
> prioridad, y sin nadie delante se quedaría hablando sola. `cierre-auto` no
> pregunta nada, registra solo lo observable en las notas del día y distingue
> «sin señal» de «no hecho», que no es lo mismo.
>
> Las dos escriben en `outputs\AAAA-MM-DD-cierre.md` y no se pisan: el
> automático marca `origen: automatico` y se aparta si ya hay uno manual; el
> manual lo reemplaza siempre, porque tú estabas delante confirmando. Puedes
> dejar la tarea programada puesta y cerrar el día a mano igualmente cuando te
> apetezca.

Para revisar o borrar: `schtasks /query /tn "JARVIS*"`, `schtasks /delete /tn "JARVIS cierre"`.

### 5.2 Siguientes skills, por rentabilidad

1. `resumen-bandeja` — necesita el conector de Gmail.
2. `agenda-hoy` — conector de Google Calendar.
3. `metricas` — solo si publicas contenido y tienes las APIs a mano.
4. **Skills de CHACONTAINER**: presupuestos, seguimiento de envíos, informes de
   inventario. Aquí es donde el sistema empieza a pagarse solo, porque los
   scripts ya existen en `chacontainer/scripts/`.

### 5.3 Higiene semanal

Una skill `revisar-boveda` que una vez por semana proponga qué notas de `raw/`
merecen subir a `wiki/`. Tú apruebas; ella no escribe en `wiki/` por su cuenta —
los permisos no se lo permiten, y esa es la idea.

---

## Resumen de un vistazo

```
Fase 0   30 min   requisitos + voz es-ES   -> claude -p responde
Fase 1   30 min   bóveda + CLAUDE.md       -> te conoce
Fase 2    1-2 h   4 skills                 -> se disparan solas, por teclado
Fase 3    1-2 h   voz local                -> hablas y te responde
Fase 4  1 tarde   HUD                      -> una pantalla
Fase 5 continuo   tareas + integraciones   -> funciona sin ti
```

Si solo tienes una hora esta semana: fases 0, 1 y 2. Son el 80 % del valor. La
voz y el HUD son la capa vistosa, y sin skills buenas no sirven de nada.

---

## Apéndice — si acabas migrando a WSL2 o macOS

Solo cambian tres cosas: los instaladores (`apt`/`brew` en vez de `winget`), el
TTS (Piper o el comando `say` de macOS en vez de SAPI) y las rutinas (`cron` en
vez de `schtasks`). La bóveda, el `CLAUDE.md`, los permisos y las skills son
idénticos: son ficheros de texto.

Aviso concreto sobre WSL2: el micrófono se accede desde Windows, no desde WSL.
Si migras, deja el bucle de voz en Windows y apunta `JARVIS_VAULT` a la bóveda
por `\\wsl$\...`, o al revés. Es la parte que más rompe.

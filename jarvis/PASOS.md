# JARVIS — Pasos para el ordenador

Guía de ejecución. Cada fase termina con una **prueba**: no pases a la siguiente
sin superarla.

Asunciones: usas macOS, Linux o Windows con WSL2. Si estás en Windows nativo sin
WSL, dímelo y adapto los comandos (PowerShell + Task Scheduler en vez de bash +
cron).

En toda la guía, `~/boveda` es la carpeta de la bóveda. Cámbiala si prefieres
otra ruta.

---

## Fase 0 — Requisitos (30 min)

### 0.1 Instalar lo básico

| Pieza | macOS | Linux / WSL2 | Windows nativo |
|---|---|---|---|
| Node 18+ | `brew install node` | `sudo apt install nodejs npm` | instalador de nodejs.org |
| Python 3.11+ | `brew install python@3.11` | `sudo apt install python3 python3-pip python3-venv` | instalador de python.org |
| ffmpeg | `brew install ffmpeg` | `sudo apt install ffmpeg` | `winget install ffmpeg` |
| Obsidian | `brew install --cask obsidian` | AppImage de obsidian.md | `winget install Obsidian.Obsidian` |
| Claude Code | `npm i -g @anthropic-ai/claude-code` | igual | igual |

### 0.2 Autenticar Claude Code

```bash
claude
# sigue el login en el navegador, luego /exit
```

**Prueba de fase 0:**
```bash
claude --version && claude -p "responde solo: ok"
```
Debe imprimir la versión y `ok`.

---

## Fase 1 — Cablea el cerebro y la memoria (30 min)

### 1.1 Crear la bóveda

```bash
mkdir -p ~/boveda/{raw,wiki,outputs,.claude/skills}
cd ~/boveda
git init          # opcional pero recomendado: historial de tu memoria
```

Abre Obsidian → *Open folder as vault* → `~/boveda`.

### 1.2 `~/boveda/CLAUDE.md` — las reglas de la casa

Este fichero se carga en cada sesión. Es lo que convierte a Claude Code en JARVIS.

```markdown
# JARVIS — instrucciones operativas

Eres mi asistente personal. Respondes en español, en frases cortas,
pensadas para ser **leídas en voz alta**: sin markdown, sin listas con viñetas,
sin bloques de código, salvo que se pida explícitamente por escrito.

## La bóveda

- `raw/`     — capturas en bruto. Puedes escribir aquí libremente.
- `outputs/` — informes y entregables que generas. Puedes escribir aquí.
- `wiki/`    — conocimiento curado. **Solo lectura.** Si crees que algo debe ir
               al wiki, propónlo; no lo escribas.

## Reglas

1. Antes de responder algo que dependa de mi contexto, busca en la bóveda con
   Grep. No leas la bóveda entera nunca.
2. Toda nota nueva lleva frontmatter: `fecha`, `tipo`, `tags`.
3. Enlaza con wikilinks `[[nota]]` a notas que ya existan. No inventes enlaces.
4. Si no encuentras algo en la bóveda, dilo. No rellenes huecos.
5. Respuesta por defecto: máximo 4 frases. Si hace falta más, escribe el detalle
   en `outputs/` y resume en voz alta dónde lo dejaste.

## Nombres de fichero

`raw/AAAA-MM-DD-tema.md`, `outputs/AAAA-MM-DD-tipo.md`.
```

### 1.3 `~/boveda/.claude/settings.json` — permisos acotados

Sin esto, el bucle de voz se bloqueará pidiendo confirmación. Con esto, queda
limitado a la bóveda.

```json
{
  "permissions": {
    "allow": [
      "Read",
      "Grep",
      "Glob",
      "Write(./raw/**)",
      "Write(./outputs/**)",
      "Edit(./raw/**)",
      "Edit(./outputs/**)",
      "Bash(date:*)",
      "Bash(ls:*)"
    ],
    "deny": [
      "Write(./wiki/**)",
      "Edit(./wiki/**)",
      "Bash(rm:*)",
      "Bash(curl:*)"
    ]
  }
}
```

> No uses `--permission-mode bypassPermissions` en el bucle de voz. El STT se
> equivoca y ejecutaría lo que crea haber oído.

### 1.4 Una nota semilla

```bash
cat > ~/boveda/wiki/contexto.md <<'EOF'
---
fecha: 2026-08-12
tipo: contexto
tags: [personal]
---
# Contexto

Quién soy, en qué trabajo, qué proyectos tengo activos y qué me importa esta
semana. JARVIS lee esto para no preguntar lo obvio.
EOF
```
Edítala con tu contexto real. Es la nota más rentable de todo el sistema.

**Prueba de fase 1:**
```bash
cd ~/boveda && claude -p "¿qué sabes de mí? Responde en dos frases."
```
Debe responder usando `wiki/contexto.md`, no genéricamente.

---

## Fase 2 — Las skills (1–2 h)

Cada skill es una carpeta con un `SKILL.md`. El frontmatter `description` es lo
que hace que Claude la elija sola: **escríbelo como se dispararía en voz**.

### 2.1 `plan-hoy`

```bash
mkdir -p ~/boveda/.claude/skills/plan-hoy
cat > ~/boveda/.claude/skills/plan-hoy/SKILL.md <<'EOF'
---
name: plan-hoy
description: Fija las tres prioridades del día y las escribe en la bóveda. Úsala cuando pida "plan de hoy", "qué hago hoy", "prioridades", "organiza mi día" o "empecemos el día".
---

# Plan de hoy

1. Lee `raw/` de los últimos 3 días y `outputs/` de ayer con Grep, buscando
   tareas pendientes o compromisos sin cerrar.
2. Lee `wiki/contexto.md`.
3. Propón exactamente 3 prioridades, ordenadas por impacto. Nada más.
4. Escribe `raw/AAAA-MM-DD-plan.md` con frontmatter `tipo: plan` y las 3
   prioridades como checklist.
5. Léelas en voz alta en una sola frase por prioridad.

Si ya existe el plan de hoy, no lo dupliques: léelo y pregunta si hay cambios.
EOF
```

### 2.2 `cierre-dia`

```bash
mkdir -p ~/boveda/.claude/skills/cierre-dia
cat > ~/boveda/.claude/skills/cierre-dia/SKILL.md <<'EOF'
---
name: cierre-dia
description: Cierra el día, registra qué se hizo y deja encolado mañana. Úsala cuando pida "cierra el día", "resumen del día", "qué he hecho hoy" o "terminemos".
---

# Cierre del día

1. Lee `raw/AAAA-MM-DD-plan.md` de hoy.
2. Pregúntame en voz alta, una por una, qué pasó con cada prioridad. Espera
   respuesta antes de pasar a la siguiente.
3. Escribe `outputs/AAAA-MM-DD-cierre.md`: qué se cerró, qué quedó abierto, una
   línea de reflexión, y lo que arrastra a mañana.
4. Enlaza al plan del día con `[[AAAA-MM-DD-plan]]`.
5. Termina con una sola frase: qué es lo primero de mañana.
EOF
```

### 2.3 `nota-rapida`

```bash
mkdir -p ~/boveda/.claude/skills/nota-rapida
cat > ~/boveda/.claude/skills/nota-rapida/SKILL.md <<'EOF'
---
name: nota-rapida
description: Captura una idea, recordatorio o dato suelto en la bóveda sin interrumpir. Úsala cuando diga "apunta", "anota", "guarda esto", "recuérdame" o suelte una idea sin pedir nada más.
---

# Nota rápida

1. Escribe lo dicho en `raw/AAAA-MM-DD-notas.md` (añade al fichero del día, no
   crees uno nuevo por nota).
2. Añádele un `tipo` inferido: idea, tarea, contacto, dato.
3. Si menciona algo que ya existe en `wiki/`, enlázalo con wikilink.
4. Confirma en **una sola frase corta**. No repitas la nota entera.
EOF
```

### 2.4 `buscar-boveda`

```bash
mkdir -p ~/boveda/.claude/skills/buscar-boveda
cat > ~/boveda/.claude/skills/buscar-boveda/SKILL.md <<'EOF'
---
name: buscar-boveda
description: Responde preguntas sobre lo que ya está guardado en la bóveda. Úsala cuando pregunte "qué dije sobre", "cuándo hablamos de", "búscame", "qué sé de" o cualquier pregunta sobre mi propio historial.
---

# Buscar en la bóveda

1. Grep por los términos clave. Prueba sinónimos si no hay resultados.
2. Lee como mucho los 3 ficheros más relevantes. Nunca más.
3. Responde con la conclusión primero y la fecha de la nota después.
4. Si no hay nada, dilo claramente: "no hay nada en la bóveda sobre eso".
   No respondas de conocimiento general.
EOF
```

**Prueba de fase 2** (por teclado, todavía sin voz):
```bash
cd ~/boveda
claude -p "apunta que hay que renovar el seguro del almacén en octubre"
claude -p "qué hago hoy"
claude -p "qué dije sobre el seguro"
```
Las tres deben disparar la skill correcta y dejar ficheros en `raw/`.
Si una skill no se dispara: el problema está en su `description`, amplíala con
más formas de pedirlo.

---

## Fase 3 — La voz (1–2 h)

### 3.1 Entorno Python

```bash
mkdir -p ~/jarvis && cd ~/jarvis
python3 -m venv .venv && source .venv/bin/activate
pip install faster-whisper sounddevice numpy
```

### 3.2 TTS

- **macOS:** nada que instalar, usa `say -v Monica "hola"`.
- **Linux / WSL / Windows:**
  ```bash
  pip install piper-tts
  mkdir -p ~/jarvis/voces && cd ~/jarvis/voces
  # descarga es_ES-davefx-medium.onnx y .onnx.json desde
  # https://huggingface.co/rhasspy/piper-voices/tree/main/es/es_ES/davefx/medium
  ```

### 3.3 El bucle

```bash
cat > ~/jarvis/jarvis.py <<'PY'
#!/usr/bin/env python3
"""Bucle mínimo de voz: grabar -> STT local -> Claude Code -> TTS local."""
import json, os, platform, queue, subprocess, sys, wave
import sounddevice as sd
from faster_whisper import WhisperModel

VAULT = os.path.expanduser("~/boveda")
SR, WAV_IN, WAV_OUT = 16000, "/tmp/jarvis_in.wav", "/tmp/jarvis_out.wav"
ES_MAC = platform.system() == "Darwin"

print("cargando modelo de voz…")
stt = WhisperModel("small", device="cpu", compute_type="int8")


def grabar(path):
    q = queue.Queue()
    with sd.InputStream(samplerate=SR, channels=1, dtype="int16",
                        callback=lambda d, *_: q.put(bytes(d))):
        input("  grabando… ENTER para parar ")
    with wave.open(path, "wb") as w:
        w.setnchannels(1); w.setsampwidth(2); w.setframerate(SR)
        while not q.empty():
            w.writeframes(q.get())


def transcribir(path):
    segs, _ = stt.transcribe(path, language="es")
    return " ".join(s.text for s in segs).strip()


def preguntar(texto):
    r = subprocess.run(
        ["claude", "-p", texto, "--output-format", "json"],
        cwd=VAULT, capture_output=True, text=True)
    if r.returncode != 0:
        return f"Error del motor: {r.stderr.strip()[:200]}"
    return json.loads(r.stdout).get("result", "").strip()


def hablar(texto):
    if ES_MAC:
        subprocess.run(["say", "-v", "Monica", texto])
        return
    subprocess.run(["piper", "--model",
                    os.path.expanduser("~/jarvis/voces/es_ES-davefx-medium.onnx"),
                    "--output_file", WAV_OUT], input=texto, text=True)
    subprocess.run(["ffplay", "-nodisp", "-autoexit", "-loglevel", "quiet", WAV_OUT])


def main():
    print("JARVIS listo. Ctrl+C para salir.")
    while True:
        input("\nENTER para hablar ")
        grabar(WAV_IN)
        texto = transcribir(WAV_IN)
        if not texto:
            print("  (no te he oído)"); continue
        print(f"  tú: {texto}")
        respuesta = preguntar(texto)
        print(f"  jarvis: {respuesta}")
        hablar(respuesta)


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\nhasta luego.")
PY
chmod +x ~/jarvis/jarvis.py
```

### 3.4 Alias

```bash
echo "alias jarvis='~/jarvis/.venv/bin/python ~/jarvis/jarvis.py'" >> ~/.zshrc
source ~/.zshrc     # usa ~/.bashrc si tu shell es bash
```

**Prueba de fase 3:** ejecuta `jarvis`, di *"apunta que mañana hay que llamar al
transportista"*, y comprueba que responde en voz alta y que aparece la nota en
`~/boveda/raw/`.

Si el micrófono no se detecta: `python -c "import sounddevice; print(sounddevice.query_devices())"`
y fija el dispositivo con `sd.default.device = <índice>`.

---

## Fase 4 — El HUD (una tarde)

Tienes dos caminos. **Recomiendo el A**, porque ya tienes el stack montado.

### Camino A — ruta nueva en `chacontainer/web` (React + Vite + Tailwind)

1. Añade en el servidor Go (`chacontainer/cmd/server`) un endpoint
   `GET /api/vault` que devuelva, leyendo `~/boveda`:
   - las 10 notas más recientes (nombre, fecha, primeras 200 letras),
   - el plan de hoy si existe,
   - conteos: notas totales, notas de esta semana, tamaño de la bóveda.
2. Crea la ruta `/hud` en la app web con cuatro paneles: vitales, panel de
   comandos, agenda de hoy, últimas notas.
3. Que escuche solo en `127.0.0.1`. Es tu máquina, no publiques la bóveda.

Prompt para arrancarlo, ejecutado desde la raíz del repo:

```
Añade un HUD en chacontainer/web: ruta /hud, tema terminal oscuro, cuatro
paneles — vitales del sistema, panel de comandos, agenda de hoy y últimas notas.
Los datos vienen de un endpoint nuevo GET /api/vault en cmd/server que lee la
carpeta ~/boveda. Sin dependencias nuevas de frontend; usa el Tailwind que ya
está configurado. El servidor escucha solo en 127.0.0.1.
```

### Camino B — HUD independiente de un solo fichero

`~/jarvis/hud.html` + `python3 -m http.server 8787 --directory ~/boveda`.
Cero dependencias, cero integración. Úsalo si quieres verlo funcionando hoy.

**Prueba de fase 4:** abres el HUD y ves el plan de hoy que generaste en la fase 2.

---

## Fase 5 — Rutinas y crecimiento (continuo)

### 5.1 Rutinas programadas

`crontab -e` (macOS/Linux/WSL):

```cron
0 7 * * 1-5 cd ~/boveda && claude -p "resumen matutino" >> ~/boveda/outputs/cron.log 2>&1
0 19 * * 1-5 cd ~/boveda && claude -p "cierra el día" >> ~/boveda/outputs/cron.log 2>&1
```

En Windows nativo, lo mismo con el Programador de tareas.

> Ojo: `cierre-dia` es conversacional (te pregunta). Para cron, escribe una
> variante `cierre-auto` que no espere respuestas y solo registre lo observable.

### 5.2 Siguientes skills, por orden de rentabilidad

1. `resumen-bandeja` — necesita el conector de Gmail.
2. `agenda-hoy` — conector de Google Calendar.
3. `metricas` — solo si publicas contenido y tienes las APIs a mano.
4. Skills de CHACONTAINER: presupuestos, seguimiento de envíos, informes de
   inventario. Aquí es donde el sistema empieza a pagar por sí mismo, porque
   ya tienes los scripts en `chacontainer/scripts/`.

### 5.3 Higiene semanal

Una skill `revisar-boveda` que una vez por semana proponga qué notas de `raw/`
merecen promoverse a `wiki/`. Tú apruebas; ella no escribe en `wiki/` sola.

---

## Resumen de un vistazo

```
Fase 0  30 min   requisitos            → claude -p responde
Fase 1  30 min   bóveda + CLAUDE.md    → te conoce
Fase 2   1-2 h   4 skills              → se disparan solas por teclado
Fase 3   1-2 h   voz local             → hablas y responde
Fase 4   1 tarde HUD                   → una pantalla
Fase 5  continuo cron + integraciones  → funciona sin ti
```

Si solo tienes una hora esta semana, haz las fases 0, 1 y 2. Es el 80 % del
valor. La voz y el HUD son la parte vistosa, y sin skills buenas no sirven de
nada.

#!/usr/bin/env python3
"""JARVIS — bucle de voz para Windows.

grabar (micro) -> STT local (faster-whisper) -> Claude Code headless -> TTS local (SAPI)

Nada de audio sale de la máquina. El texto transcrito sí viaja a la API de
Anthropic, que es quien responde.

Uso:  python jarvis.py
      python jarvis.py --texto        (sin micro: escribes en vez de hablar)
      python jarvis.py --modelo base  (más rápido, menos preciso)
"""

from __future__ import annotations

import argparse
import json
import os
import queue
import shutil
import subprocess
import sys
import wave
from pathlib import Path

# --- Configuración -----------------------------------------------------------

VAULT = Path(os.environ.get("JARVIS_VAULT", Path.home() / "boveda"))
SAMPLE_RATE = 16_000
TMP = Path(os.environ.get("TEMP", ".")) / "jarvis"
WAV_IN = TMP / "entrada.wav"

# Voz de Windows. Vacío = la primera voz en español que encuentre.
# Para ver las instaladas:  python jarvis.py --voces
VOZ = os.environ.get("JARVIS_VOZ", "")

CLAUDE = shutil.which("claude")

# --- Salida de consola en UTF-8 (si no, los acentos se rompen en Windows) -----

for flujo in (sys.stdout, sys.stderr):
    try:
        flujo.reconfigure(encoding="utf-8")
    except (AttributeError, ValueError):
        pass


# --- Voz: escuchar -----------------------------------------------------------


def grabar(destino: Path) -> None:
    """Graba desde el micrófono hasta que pulses ENTER."""
    import sounddevice as sd

    trozos: queue.Queue[bytes] = queue.Queue()

    def recoger(datos, *_):
        trozos.put(bytes(datos))

    with sd.InputStream(samplerate=SAMPLE_RATE, channels=1,
                        dtype="int16", callback=recoger):
        input("   grabando…  ENTER para parar ")

    destino.parent.mkdir(parents=True, exist_ok=True)
    with wave.open(str(destino), "wb") as w:
        w.setnchannels(1)
        w.setsampwidth(2)
        w.setframerate(SAMPLE_RATE)
        while not trozos.empty():
            w.writeframes(trozos.get())


def cargar_stt(tamano: str):
    from faster_whisper import WhisperModel

    print(f"   cargando modelo de voz «{tamano}» (la primera vez se descarga)…")
    return WhisperModel(tamano, device="cpu", compute_type="int8")


def transcribir(stt, origen: Path) -> str:
    segmentos, _ = stt.transcribe(str(origen), language="es", vad_filter=True)
    return " ".join(s.text for s in segmentos).strip()


# --- Cerebro: Claude Code headless -------------------------------------------


def preguntar(texto: str) -> str:
    """Lanza Claude Code en la bóveda y devuelve la respuesta en texto plano."""
    if CLAUDE is None:
        return ("No encuentro el comando claude. Instálalo con "
                "npm install -g @anthropic-ai/claude-code y reabre la terminal.")

    try:
        proceso = subprocess.run(
            [CLAUDE, "-p", texto, "--output-format", "json"],
            cwd=str(VAULT),
            capture_output=True,
            text=True,
            encoding="utf-8",
            errors="replace",
            timeout=180,
        )
    except subprocess.TimeoutExpired:
        return "El motor ha tardado demasiado. Prueba a pedirme algo más concreto."

    if proceso.returncode != 0:
        detalle = (proceso.stderr or proceso.stdout or "").strip()
        return f"Error del motor: {detalle[:300]}"

    try:
        return json.loads(proceso.stdout).get("result", "").strip()
    except json.JSONDecodeError:
        return proceso.stdout.strip()


# --- Voz: hablar (SAPI, ya viene con Windows) --------------------------------

_PS_LISTAR = (
    "Add-Type -AssemblyName System.Speech;"
    "(New-Object System.Speech.Synthesis.SpeechSynthesizer).GetInstalledVoices()"
    " | ForEach-Object { $_.VoiceInfo.Name + '  [' + $_.VoiceInfo.Culture.Name + ']' }"
)

_PS_HABLAR = """
Add-Type -AssemblyName System.Speech
$s = New-Object System.Speech.Synthesis.SpeechSynthesizer
$pedida = $env:JARVIS_VOZ
if ($pedida) {
    $s.SelectVoice($pedida)
} else {
    $es = $s.GetInstalledVoices() |
          Where-Object { $_.VoiceInfo.Culture.Name -like 'es*' } |
          Select-Object -First 1
    if ($es) { $s.SelectVoice($es.VoiceInfo.Name) }
}
$s.Rate = 1
$s.Speak([Console]::In.ReadToEnd())
"""


def _powershell(script: str, entrada: str | None = None) -> subprocess.CompletedProcess:
    return subprocess.run(
        ["powershell", "-NoProfile", "-NonInteractive", "-Command", script],
        input=entrada,
        capture_output=True,
        text=True,
        encoding="utf-8",
        errors="replace",
    )


def listar_voces() -> None:
    salida = _powershell(_PS_LISTAR)
    print(salida.stdout or salida.stderr)
    print("Si no aparece ninguna [es-..], instala el paquete de voz:")
    print("  Configuración → Hora e idioma → Idioma y región → Español → Voz")


def hablar(texto: str) -> None:
    if not texto:
        return
    entorno_previo = os.environ.get("JARVIS_VOZ")
    if VOZ:
        os.environ["JARVIS_VOZ"] = VOZ
    try:
        resultado = _powershell(_PS_HABLAR, entrada=texto)
        if resultado.returncode != 0:
            print(f"   (TTS falló: {resultado.stderr.strip()[:200]})")
    finally:
        if entorno_previo is None:
            os.environ.pop("JARVIS_VOZ", None)


# --- Bucle -------------------------------------------------------------------


def main() -> int:
    opciones = argparse.ArgumentParser(description="JARVIS — asistente por voz")
    opciones.add_argument("--texto", action="store_true",
                          help="escribir en vez de hablar (útil para depurar)")
    opciones.add_argument("--modelo", default="small",
                          help="tamaño de Whisper: tiny, base, small, medium")
    opciones.add_argument("--voces", action="store_true",
                          help="listar las voces de Windows disponibles y salir")
    args = opciones.parse_args()

    if args.voces:
        listar_voces()
        return 0

    if not VAULT.is_dir():
        print(f"No existe la bóveda en {VAULT}.")
        print("Créala o define la variable JARVIS_VAULT.")
        return 1

    stt = None if args.texto else cargar_stt(args.modelo)

    print(f"\n   JARVIS listo.  Bóveda: {VAULT}")
    print("   Ctrl+C para salir.\n")

    while True:
        if args.texto:
            dicho = input("tú> ").strip()
        else:
            input("   ENTER para hablar ")
            grabar(WAV_IN)
            dicho = transcribir(stt, WAV_IN)
            print(f"   tú:     {dicho or '(no te he oído)'}")

        if not dicho:
            continue

        respuesta = preguntar(dicho)
        print(f"   jarvis: {respuesta}\n")
        hablar(respuesta)


if __name__ == "__main__":
    try:
        sys.exit(main())
    except KeyboardInterrupt:
        print("\n   hasta luego.")

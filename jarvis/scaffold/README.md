# Andamiaje de JARVIS

Las fases 1 y 2 de [`../PASOS.md`](../PASOS.md), ya construidas. Instálalo en un
comando y empieza por la prueba de la fase 2.

## Instalar (Windows)

```powershell
powershell -ExecutionPolicy Bypass -File .\instalar.ps1
```

| Opción | Para qué |
|---|---|
| `-Boveda D:\ruta` | Poner la bóveda en otro sitio (por defecto `$HOME\boveda`) |
| `-Destino D:\ruta` | Dónde va el bucle de voz (por defecto `$HOME\jarvis`) |
| `-SinVoz` | Solo la bóveda y las skills; salta Python y el micrófono |
| `-Sobrescribir` | Pisar ficheros existentes. **Cuidado:** reemplaza tu `CLAUDE.md` |

El script no toca nada que ya exista salvo que se lo pidas. Puedes volver a
ejecutarlo sin miedo para actualizar lo que falte.

## Qué contiene

```
boveda/
  CLAUDE.md              instrucciones operativas: tono, mapa, reglas duras
  wiki/contexto.md       plantilla que TIENES que rellenar (lo que más rinde)
  dot-claude/
    settings.json        permisos: lee todo, escribe solo en raw/ y outputs/
    skills/
      plan-hoy/          3 prioridades del día, escritas en la bóveda
      cierre-dia/        repaso conversacional y cierre en outputs/
      cierre-auto/       el mismo cierre sin nadie delante, para schtasks
      nota-rapida/       captura sin interrumpir
      buscar-boveda/     preguntas sobre tu propio historial
jarvis.py                bucle de voz: micro -> Whisper -> claude -p -> SAPI
requirements.txt         faster-whisper, sounddevice, numpy
instalar.ps1             copia todo y monta el entorno
```

`dot-claude/` se llama así en el repo a propósito: el instalador lo renombra a
`.claude/` al copiarlo. Así estas skills no se cargan al abrir Claude Code en
este repositorio, solo en tu bóveda.

## Después de instalar

1. Rellena `wiki\contexto.md`. Sin eso, JARVIS no te conoce.
2. Abre la bóveda en Obsidian (*Open folder as vault*).
3. Prueba sin voz: `boveda`, luego `claude -p "que hago hoy"`.
4. Terminal nueva y prueba la voz: `jarvis --voces`, `jarvis --texto`, `jarvis`.

## Cómo se ajusta

- **Una skill no se dispara** → amplía su `description` con más formas de pedir
  lo mismo, en las palabras que usarías hablando.
- **Responde demasiado largo** → regla 5 de `CLAUDE.md`.
- **Tarda demasiado** → la skill está leyendo de más: Grep primero, tres
  ficheros como máximo.
- **Quieres una skill nueva** → copia la carpeta de `nota-rapida`, que es la más
  simple, y cambia `name`, `description` y los pasos.

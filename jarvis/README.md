# JARVIS

Asistente personal por voz: hablas, Claude Code ejecuta, Obsidian recuerda.
Montado para **Windows nativo** (PowerShell, voz SAPI, Programador de tareas).

## Empieza aquí

```powershell
cd jarvis\scaffold
powershell -ExecutionPolicy Bypass -File .\instalar.ps1
```

Requisitos previos en la [fase 0](./PASOS.md#fase-0--requisitos-30-min).

## Los documentos

| | |
|---|---|
| [`ANALISIS.md`](./ANALISIS.md) | Qué del sistema es real y qué es marketing, la arquitectura, el presupuesto de latencia, las decisiones previas y los riesgos. **Léelo primero.** |
| [`PASOS.md`](./PASOS.md) | Las cinco fases con los comandos concretos de PowerShell y una prueba al final de cada una. |
| [`scaffold/`](./scaffold/) | Las fases 1 y 2 ya construidas: bóveda, permisos, cuatro skills y el bucle de voz. |

## El resumen corto

El valor está en las **skills** y en la **bóveda**, no en la voz ni en el HUD.
Las fases 0–2 son unas dos horas y dan el 80 % del resultado; la voz y la
pantalla son la capa vistosa, y sin skills buenas no sirven de nada.

Dos avisos que el carrusel original no da: el audio no sale de tu máquina pero
**el texto transcrito sí** viaja a la API, y **el cerebro no es local** — Claude
Code corre contra modelos de Anthropic; lo local es la voz.

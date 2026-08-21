---
name: cierre-auto
description: Cierre del día desatendido, para la tarea programada de la tarde. Úsala cuando pida "cierre automatico", "cierre desatendido" o "registra el dia sin preguntarme". No la uses si estoy delante, para eso está cierre-dia.
---

# Cierre automático

Esta skill se ejecuta **sin nadie delante**, desde el Programador de tareas. La
regla que manda sobre todas las demás:

> **No hagas ni una sola pregunta.** Nadie va a contestarla. Si te falta un dato,
> escríbelo como desconocido y sigue.

## Señales que puedes usar

Solo estas, y solo con Grep y Glob. No leas la bóveda entera.

1. `raw/AAAA-MM-DD-plan.md` de hoy — las prioridades y sus casillas.
2. `raw/AAAA-MM-DD-notas.md` de hoy — lo capturado durante la jornada.
3. Cualquier otro fichero de `raw/` o `outputs/` con la fecha de hoy.

Si no existe ninguno de los tres, ve directo al apartado «Día sin señal».

## Qué escribir

Escribe `outputs/AAAA-MM-DD-cierre.md`:

```
---
fecha: AAAA-MM-DD
tipo: cierre
origen: automatico
tags: [cierre, diario]
---
# Cierre del AAAA-MM-DD

## Cerrado
Prioridades con la casilla marcada, y trabajo que conste en las notas del día.

## Sin señal
Prioridades que siguen sin marcar. Redáctalas como "sin señal", nunca como
"no hecho": no tienes forma de saberlo.

## Capturado hoy
Una línea por nota de `raw/` de hoy, resumida en media frase.

## Arrastra a mañana
Todo lo de "Sin señal" más lo que las notas dejen abierto explícitamente.
```

Enlaza al plan con `[[AAAA-MM-DD-plan]]` si el fichero existe.

## Límites duros

- **No marques casillas** en el plan de hoy. Eso solo lo hace `cierre-dia`,
  con una persona confirmando. Aquí solo observas.
- **No interpretes intenciones.** Si una nota dice "hablar con el
  transportista", eso es una nota, no una tarea completada.
- **No inventes una reflexión.** El apartado de reflexión es de `cierre-dia`,
  porque requiere que alguien piense. Aquí no va.
- **No escribas en `wiki/`.** Los permisos ya lo impiden; que no haga falta.

## Si ya existe el cierre de hoy

- Con `origen: manual` → **no toques nada**. La persona ya cerró el día a mano y
  su versión vale más que la tuya. Termina diciendo solo: "ya había cierre
  manual, no hago nada".
- Con `origen: automatico` → reescríbelo con los datos actuales.

## Día sin señal

Si no hay ni plan ni notas de hoy, escribe igualmente el fichero con el cuerpo
vacío y una única línea: "Sin actividad registrada en la bóveda." Es información:
así el hueco queda visible en el historial en vez de desaparecer.

## Salida por consola

Como esto acaba en un fichero de log, termina con **una sola línea**: qué has
escrito y cuántas prioridades quedan arrastrando. Sin saludos ni resúmenes.

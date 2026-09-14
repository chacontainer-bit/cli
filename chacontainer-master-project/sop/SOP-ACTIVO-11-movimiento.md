# SOP-ACTIVO-11: Movimiento

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre el registro de cada desplazamiento físico del activo retornable (contenedor, IBC,
tarima o tambor) una vez que ya tiene custodio confirmado: la salida física, el tránsito entre ubicaciones
o hubs intermedios, y la entrega en destino. Corresponde a las transiciones `Asignado → En tránsito`,
`En tránsito → En uso` y el self-loop `En tránsito → En tránsito` (transbordo) de la
[máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones) (filas 12, 13 y 14), y actualiza
[Capa 2 · Ubicación](../README.md#etapa-3--modelo-de-packaging-systems) del
[ciclo de vida del activo](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado) en las etapas
"Salida" y "Movimiento". Se construyó a partir del cruce entre el
[catálogo Systems](../catalogo-systems.md#7-trazabilidad) y la máquina de estados — **no a partir de una
entrevista real con el responsable de logística**. No debe declararse estándar hasta validar en operación
los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Movimiento es el proceso que mantiene actualizada la ubicación física de un activo retornable desde el
momento en que sale de instalaciones de CHACONTAINER (o del punto donde fue asignado) hasta que llega a
manos del custodio final. Cubre tres eventos: la salida física que dispara `Asignado → En tránsito`, los
pasos intermedios por hubs o puntos de transbordo (self-loop `En tránsito → En tránsito`), y la entrega en
destino que dispara `En tránsito → En uso`. Sin este proceso, [Capa 2 · Ubicación](../README.md#etapa-3--modelo-de-packaging-systems)
deja de reflejar la realidad apenas el activo sale de planta, y la [Trazabilidad](../catalogo-systems.md#7-trazabilidad)
se rompe justo en el tramo de mayor riesgo de pérdida.

## 2. Objetivo

Garantizar que todo desplazamiento físico del activo entre la asignación y la entrega en destino queda
registrado en tiempo cercano al real en Capa 2 · Ubicación, incluyendo cualquier parada o transbordo
intermedio, de forma que en ningún momento del tránsito se pierda visibilidad de dónde está el activo.

## 3. Alcance

Aplica al desplazamiento físico del activo desde que tiene custodio confirmado (`Asignado`) hasta que llega
a manos de ese custodio (`En uso`), incluyendo pasos por hubs, centros de consolidación o transbordo entre
transportistas. **No incluye:**

- La asignación misma (definición del custodio y del criterio de reglas aplicables) — ver
  [SOP-ACTIVO-10](./SOP-ACTIVO-10-asignacion.md) (Asignación).
- Un cambio de custodio sin que CHACONTAINER ejecute o coordine el movimiento físico — ver
  [SOP-ACTIVO-12](./SOP-ACTIVO-12-transferencia-de-custodia.md) (Transferencia de custodia).
- El desplazamiento de retorno (`En uso → En tránsito` y lo que sigue) — ver
  [SOP-ACTIVO-13](./SOP-ACTIVO-13-retorno.md) (Retorno), que reutiliza el mismo mecanismo de registro de
  movimiento pero con dirección e intención distintas (regresar, no entregar).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Logística `[VALIDAR nombre exacto del puesto/área]` | Confirma la salida física del activo, coordina y registra cada transbordo intermedio. |
| Sistema (automático, vía escaneo QR/RFID) | Dispara `En tránsito → En uso` al detectar el evento de entrega en destino, según la [tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones) (fila 13). |
| Custodio receptor | Confirma la recepción del activo cuando no hay escaneo automático disponible en destino `[VALIDAR mecanismo de confirmación cuando el custodio no tiene lector]`. |
| Responsable de hub/transbordo `[VALIDAR si existe como rol o si lo asume Logística]` | Registra el paso del activo por un punto intermedio antes de reenviarlo a destino. |

## 5. Diagrama de flujo

```
[Activo en estado "Asignado" — custodio confirmado, pendiente de salida física]
                ↓
[Logística confirma salida física del activo]
                ↓
[Transición Asignado → En tránsito · Capa 2 se actualiza: ubicación = "en ruta"]
                ↓
        ¿Pasa por un hub o punto de transbordo intermedio?
        ↓ sí                                          ↓ no
[Registrar llegada al hub/transbordo                [Continuar en tránsito
 → self-loop En tránsito → En tránsito                directo a destino]
 · Capa 2 se actualiza: ubicación = hub]
        ↓
[Reanudar tránsito hacia destino final]
        ↓
[Entrega en destino: escaneo QR/RFID o confirmación del custodio]
        ↓
[Transición En tránsito → En uso · Capa 2 se actualiza: ubicación = sitio del custodio]
        ↓
[Activo queda bajo trazabilidad pasiva mientras está "En uso" —
 ver SOP-ACTIVO-13 para el inicio del retorno]
```

## 6. SOP paso a paso

1. Verificar que el activo está en estado `Asignado` con custodio confirmado antes de iniciar cualquier
   movimiento (precondición heredada de [SOP-ACTIVO-10](./SOP-ACTIVO-10-asignacion.md)).
2. Logística confirma la salida física del activo de la ubicación de origen (planta, almacén o punto de
   entrega acordado).
3. Registrar el evento de salida: dispara la transición `Asignado → En tránsito`
   ([tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones), fila 12) y actualiza
   Capa 2 · Ubicación con el destino planeado y la fecha/hora de salida.
4. Si la ruta incluye uno o más hubs o puntos de transbordo, registrar cada llegada intermedia como un
   evento de trazabilidad independiente: dispara el self-loop `En tránsito → En tránsito`
   (fila 14) sin cambiar el estado, pero sí la ubicación registrada en Capa 2
   `[VALIDAR si el transbordo requiere escaneo obligatorio o basta el registro manual de Logística]`.
5. Reanudar el tránsito hacia el destino final tras cada transbordo, dejando la ubicación anterior como
   histórico y no como ubicación vigente.
6. En destino, confirmar la entrega mediante escaneo QR/RFID (evento automático) o, si no hay lector
   disponible, mediante confirmación explícita del custodio receptor `[VALIDAR canal de confirmación
   manual — firma física, confirmación digital, llamada registrada]`.
7. Registrar el evento de entrega: dispara la transición `En tránsito → En uso`
   (fila 13) y actualiza Capa 2 · Ubicación con la ubicación final del custodio.
8. Dejar evidencia de cada tramo (salida, transbordos, entrega) en
   [Trazabilidad](../catalogo-systems.md#7-trazabilidad), con marca de tiempo, ubicación y actor de cada
   evento — sin excepción, según el invariante 5 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar).
9. Actualizar el indicador de tiempo de tránsito por ruta y reportar cualquier tramo sin evento registrado
   por más de `[VALIDAR umbral de horas/días sin evento antes de generar alerta]`.

## 7. Checklist

- [ ] Activo verificado en estado `Asignado` con custodio confirmado antes de la salida.
- [ ] Salida física registrada con fecha, hora y destino planeado.
- [ ] Cada transbordo/hub intermedio registrado como evento propio (no omitido "porque es el mismo viaje").
- [ ] Capa 2 · Ubicación refleja siempre la última ubicación conocida, no la ubicación de origen.
- [ ] Entrega en destino confirmada por escaneo o por confirmación explícita del custodio.
- [ ] Transición `En tránsito → En uso` registrada el mismo día de la entrega real.
- [ ] Ningún tramo del trayecto queda sin evento de trazabilidad asociado.

## 8. KPI

- Tiempo de tránsito por ruta (salida → entrega confirmada).
- % de movimientos con confirmación de entrega registrada dentro del plazo esperado
  `[VALIDAR plazo esperado por tipo de ruta]`.
- % de activos con al menos un tramo de tránsito sin evento de trazabilidad (gap de datos).
- Tiempo promedio detenido en hub/punto de transbordo.
- Contribuye directamente al KPI de **Tiempo de ciclo** de la [ETAPA 6](../README.md#etapa-6--indicadores-del-sistema).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Activo en tránsito sin ningún evento intermedio registrado (se pierde visibilidad durante días) | Exigir registro de cada transbordo como evento propio (paso 4); alertar cuando un tramo excede el umbral de horas sin evento `[VALIDAR umbral]`. |
| Confirmación de entrega no registrada — el activo queda "En tránsito" indefinidamente aunque ya esté en uso | Escaneo automático como método preferente; confirmación manual del custodio como respaldo documentado, no como excepción silenciosa. |
| Transbordo tratado como "el mismo movimiento" y no como evento propio — se pierde el historial de manos por las que pasó el activo | Modelar explícitamente el self-loop `En tránsito → En tránsito` como transición registrada, no como continuación implícita. |
| Dependencia total de la confirmación manual del custodio para transicionar a `En uso` en rutas sin lector | Priorizar despliegue de identificación (QR como mínimo) en todo destino recurrente; documentar por escrito el canal alterno de confirmación. |
| Ubicación en Capa 2 desactualizada porque Logística registra el evento días después del movimiento real | Definir un plazo máximo de registro posterior al evento físico `[VALIDAR plazo]` y medirlo como parte del KPI de tiempo de tránsito. |

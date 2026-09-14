# SOP-ACTIVO-10: Asignación

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre la etapa **Asignación** del [ciclo de vida del activo](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado),
ejecutada por los servicios [Renta](../catalogo-solutions.md#2-renta) y [Venta](../catalogo-solutions.md#1-venta)
de Solutions y sostenida por [Gestión de custodios](../catalogo-systems.md#8-gestión-de-custodios) y
[Reglas operativas](../catalogo-systems.md#11-reglas-operativas) de Systems, es decir, la transición
`Disponible → Asignado` de la [máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones)
(transición #11). Se construyó a partir del cruce entre esos catálogos y el ciclo de vida — **no a partir de
una entrevista real con el responsable comercial u operativo**. No debe declararse estándar hasta validar en
operación los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Asignación es el proceso que convierte un activo en estado `Disponible` en un activo `Asignado`: elige qué
unidad específica del saldo disponible entregar, valida que cumple las reglas mínimas para salir (tipo
correcto, condición permitida) y registra al custodio inicial que se hace responsable de él — el primer dato
real de [Capa 3 · Custodia](../README.md#etapa-3--modelo-de-packaging-systems). Sin este paso no hay cadena
de responsabilidad: es la frontera entre "el activo existe en inventario" y "alguien específico responde por
él" (ver [invariante 1](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar) de
estados-del-activo.md — ningún activo pasa de `Disponible` a `En uso` sin pasar antes por `Asignado`).

## 2. Objetivo

Garantizar que todo activo que sale de inventario `Disponible` se elige y se valida contra reglas objetivas
(no criterio informal), y que la transferencia de custodia queda registrada antes de que el activo salga
físicamente de instalaciones de CHACONTAINER.

## 3. Alcance

Aplica a todo activo en estado `Disponible` que se asigna a un cliente, proyecto o ruta a partir de un
servicio [Renta](../catalogo-solutions.md#2-renta) o [Venta](../catalogo-solutions.md#1-venta) de Solutions.
Cubre: selección del activo específico dentro del saldo disponible, validación de reglas previas a la
asignación, y registro del custodio inicial.

No incluye la salida física del activo de instalaciones (transición `Asignado → En tránsito`, transición #12
de la máquina de estados) ni la logística de entrega, que corresponden a
[SOP-ACTIVO-11 · Movimiento](./SOP-ACTIVO-11-movimiento.md). No incluye la negociación comercial (precio,
condiciones de renta/venta), que se resuelve antes de llegar a este SOP. No incluye la actualización continua
del saldo de inventario tras la asignación, cubierta por [SOP-ACTIVO-09 · Inventario](./SOP-ACTIVO-09-inventario.md).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Comercial / Operaciones `[VALIDAR quién dispara la transición exactamente — según la tabla de transiciones puede ser cualquiera de los dos]` | Origina la solicitud de asignación (venta o renta cerrada) y dispara la transición `Disponible → Asignado` (transición #11). |
| Responsable de inventario `[VALIDAR]` | Selecciona el activo específico dentro del saldo disponible según el criterio de asignación vigente (ver §6, paso 2) y confirma que sigue en estado `Disponible` al momento de asignarlo. |
| Responsable de reglas operativas `[VALIDAR]` | Mantiene actualizado el criterio de "condición permitida" y "tipo correcto" por cliente/uso en [Reglas operativas](../catalogo-systems.md#11-reglas-operativas), contra el que se valida cada asignación. |
| Gestión de custodios (rol operativo, no solo capacidad de sistema) `[VALIDAR]` | Registra al custodio inicial (cliente, operador, transportista, según corresponda) y confirma que la transferencia de custodia quedó formalizada antes de liberar el activo a salida. |

## 5. Diagrama de flujo

```
[Solicitud de asignación: Renta o Venta cerrada, con tipo de
 activo, cantidad y destino requeridos]
                ↓
[Consultar saldo "Disponible" (SOP-ACTIVO-09) por tipo y ubicación]
                ↓
        ¿Hay saldo suficiente del tipo requerido?
        ↓ no                                  ↓ sí
[Reportar faltante — no se        [Seleccionar activo(s) específico(s)
 puede asignar; escalar a          según criterio de asignación vigente
 disponibilidad/gobernanza]        (FIFO / por condición) [VALIDAR]]
                                                ↓
                                   [Validar tipo correcto y condición
                                    permitida contra Reglas operativas]
                                                ↓
                                        ¿Cumple ambas reglas?
                                        ↓ no                ↓ sí
                              [Rechazar la unidad,   [Confirmar transición
                               regresar a selección    Disponible → Asignado]
                               o escalar excepción]              ↓
                                                        [Registrar custodio inicial
                                                         (Capa 3) y fecha de asignación]
                                                                  ↓
                                                        [Enviar a Movimiento / Salida
                                                         (SOP-ACTIVO-11)]
```

## 6. SOP paso a paso

1. Recibir la solicitud de asignación desde el servicio [Renta](../catalogo-solutions.md#2-renta) o
   [Venta](../catalogo-solutions.md#1-venta) de Solutions, con tipo de activo, cantidad y destino/custodio
   requeridos.
2. Consultar el saldo `Disponible` vigente (ver [SOP-ACTIVO-09 · Inventario](./SOP-ACTIVO-09-inventario.md))
   por tipo y ubicación, y confirmar que hay unidades suficientes para cubrir la solicitud.
3. Si no hay saldo suficiente, detener la asignación, reportar el faltante y escalar según corresponda
   (buscar en otra ubicación, lista de espera, o señal de déficit hacia
   [Analítica](../catalogo-systems.md#14-analítica) — ver [ETAPA 4 del README](../README.md#etapa-4--modelo-de-gobernanza),
   "¿qué planta tiene déficit?").
4. Si hay saldo suficiente, seleccionar el/los activo(s) específico(s) a asignar dentro del saldo disponible
   según el criterio vigente `[VALIDAR criterio de asignación: FIFO (más antiguo primero), por condición
   (mejor condición primero o condición más ajustada al requerimiento), o combinación — no hay definición
   operativa confirmada]`.
5. Validar que el activo seleccionado cumple **tipo correcto**: coincide exactamente con lo solicitado
   (familia, medida, configuración — incluye variantes de modificación/dunnage si el cliente las requiere,
   ver [catálogo Solutions #6](../catalogo-solutions.md#6-modificación)).
6. Validar que el activo seleccionado cumple **condición permitida** para el uso/cliente destino, según la
   regla configurada en [Reglas operativas](../catalogo-systems.md#11-reglas-operativas) (Capa 5 · Gobierno)
   para ese tipo de activo o cliente `[VALIDAR: existe hoy una tabla de condición permitida por cliente, o
   se decide caso por caso]`.
7. Si el activo no cumple alguna de las dos reglas, rechazarlo para esta asignación (vuelve al saldo
   disponible para otro destino compatible) y seleccionar otra unidad, o escalar como excepción si ninguna
   unidad del saldo cumple.
8. Si el activo cumple ambas reglas, confirmar la transición `Disponible → Asignado` (transición #11 de la
   [máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones)).
9. Registrar el custodio inicial — cliente, operador, transportista, planta o área, según corresponda al
   tipo de operación — en [Gestión de custodios](../catalogo-systems.md#8-gestión-de-custodios), actualizando
   Capa 3 con fecha de asignación `[VALIDAR si se requiere acta de entrega firmada en este punto o solo en
   la salida física, ver SOP-12 planificado]`.
10. Dejar evidencia de la asignación: activo, custodio, fecha, criterio de selección aplicado y resultado de
    la validación de reglas.
11. Enviar el activo asignado al flujo de [Movimiento / Salida](./SOP-ACTIVO-11-movimiento.md), siguiente
    etapa del [ciclo de vida](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado).
12. Actualizar el saldo de inventario disponible (dispara el recálculo cubierto en
    [SOP-ACTIVO-09](./SOP-ACTIVO-09-inventario.md)) para reflejar que el activo salió de `Disponible`.

## 7. Checklist

- [ ] Saldo `Disponible` consultado y confirmado suficiente antes de seleccionar un activo específico.
- [ ] Activo seleccionado según el criterio de asignación vigente (no al azar ni por conveniencia del
      operador) `[VALIDAR criterio]`.
- [ ] Tipo correcto validado contra lo solicitado por el cliente/proyecto.
- [ ] Condición permitida validada contra Reglas operativas para ese cliente/tipo de activo.
- [ ] Custodio inicial registrado en el sistema con fecha, antes de que el activo avance a salida física.
- [ ] Evidencia de la asignación (activo, custodio, fecha, criterio, validación) queda registrada.
- [ ] Transición `Disponible → Asignado` reflejada en el sistema el mismo día de la asignación.
- [ ] Saldo de inventario disponible actualizado tras la asignación.

## 8. KPI

- **Tiempo de asignación**: desde la solicitud hasta la confirmación de `Asignado`.
- **% de solicitudes de asignación rechazadas por falta de saldo** — insumo directo para el déficit por
  planta/ubicación (ver [README ETAPA 4](../README.md#etapa-4--modelo-de-gobernanza)).
- **% de activos rechazados en validación de tipo/condición** durante la selección — mide qué tan limpio
  está el saldo `Disponible` (activos que están en el saldo pero no deberían estar disponibles para ciertos
  usos).
- **Utilización**: activos en uso / activos disponibles × 100 (ver
  [README ETAPA 6](../README.md#etapa-6--indicadores-del-sistema)) — depende directamente de que la
  asignación se ejecute sin fricción.
- **% de asignaciones con custodio registrado antes de la salida física** (objetivo: 100%).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Activo asignado sin custodio registrado — se pierde la cadena de responsabilidad (Capa 3) desde el inicio | Paso 9 obligatorio antes de enviar a Movimiento; el invariante 1 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar) prohíbe saltar directo a `En uso`. |
| Criterio de selección de activo (FIFO vs. condición) no definido — cada operador decide distinto, generando inconsistencia y quejas de clientes que reciben activos en peor condición que otros | `[VALIDAR criterio único]`; mientras no se defina, documentar caso por caso qué criterio se aplicó (paso 10) para poder auditar el patrón después. |
| Activo con condición no permitida asignado a un cliente que exige un estándar más alto (ej. industria alimentaria/farmacéutica) | Validación obligatoria contra Reglas operativas (paso 6) antes de confirmar; sin regla configurada para ese cliente, escalar en vez de asignar por defecto `[VALIDAR comportamiento cuando falta la regla — ver invariante 6 de estados-del-activo.md]`. |
| Activo bloqueado o con incidencia abierta se asigna por error porque el saldo `Disponible` no refleja su verdadero estado | Confirmar el estado real del activo (no solo su presencia en el saldo agregado) inmediatamente antes de confirmar la asignación individual — depende de que [SOP-ACTIVO-09](./SOP-ACTIVO-09-inventario.md) mantenga el saldo sincronizado. |
| Solicitud de asignación no se puede cubrir y no queda registrada como déficit — se pierde la señal de "qué planta tiene déficit" | Paso 3 obligatorio: todo faltante se reporta y escala, no se descarta silenciosamente. |

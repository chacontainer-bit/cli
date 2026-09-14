# SOP-ACTIVO-09: Inventario

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre la etapa **Inventario disponible** del [ciclo de vida del activo](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado),
sostenida por la capacidad [Inventario digital](../catalogo-systems.md#6-inventario-digital) de Systems, y
la transición `Liberado → Disponible` de la [máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones)
(transición #10), incluyendo su ciclo de reincorporación (transición #10 repetida tras `SOP-ACTIVO-08` y tras
mantenimiento). Se construyó a partir del cruce entre el [catálogo Systems](../catalogo-systems.md#6-inventario-digital),
el [catálogo Solutions](../catalogo-solutions.md#11-inventarios-físicos) y el ciclo de vida — **no a partir de
una entrevista real con el responsable de inventario**. No debe declararse estándar hasta validar en
operación los puntos marcados `[VALIDAR]`.

**Distinción clave de este SOP:** este documento cubre el **mantenimiento continuo** del saldo de inventario
disponible (un proceso de sistema que corre todo el tiempo, activo por activo, evento por evento) y la
**conciliación** de ese saldo contra un conteo físico cuando hay discrepancia. Es un proceso distinto al
servicio [Inventarios físicos](../catalogo-solutions.md#11-inventarios-físicos) del catálogo Solutions, que es
una **auditoría puntual** que CHACONTAINER vende o ejecuta sobre el parque de un cliente (o sobre su propio
parque) en una ventana de tiempo definida. Este SOP-ACTIVO-09 es lo que existe *antes y después* de esa
auditoría: el saldo digital que la auditoría contrasta, y el proceso que corrige ese saldo cuando la
auditoría (o cualquier otra señal) revela que está mal.

---

## 1. Resumen ejecutivo

Inventario es el proceso que mantiene actualizado, en todo momento, el saldo de activos en estado `Disponible`
por ubicación y tipo — la respuesta permanente a "¿cuánto tengo disponible ahora?", no solo el día de un
conteo. Recibe cada evento de cambio de estado o ubicación generado en cualquier otra etapa del ciclo
(liberación, asignación, retorno, baja) y recalcula el saldo. Cuando un
[inventario físico](../catalogo-solutions.md#11-inventarios-físicos) (propio o de cliente) revela una
discrepancia entre lo contado y lo que dice el sistema, este SOP también cubre cómo se investiga la causa y
se ajusta el saldo digital, dejando evidencia del ajuste.

## 2. Objetivo

Garantizar que el saldo de inventario `Disponible` que muestra el sistema refleja, con la mayor fidelidad
posible, la cantidad de activos físicamente utilizables en cada ubicación — la
[disponibilidad real](../README.md#etapa-6--indicadores-del-sistema) del [README](../README.md), no solo el
dato contable — y que toda discrepancia detectada se investiga, se corrige y queda registrada, en vez de
sobrescribirse en silencio.

## 3. Alcance

Aplica a:

1. **Mantenimiento continuo del saldo `Disponible`** — cada activo que entra a `Disponible` (tras
   [Liberación](./SOP-ACTIVO-08-liberacion.md)) o sale de `Disponible` (tras [Asignación](./SOP-ACTIVO-10-asignacion.md))
   o vuelve a `Disponible` (tras reincorporación desde mantenimiento).
2. **Conciliación contra inventario físico** — cuando se ejecuta el servicio
   [Inventarios físicos](../catalogo-solutions.md#11-inventarios-físicos) (auditoría puntual, propia o de
   cliente) y su resultado se contrasta contra el saldo digital.

No incluye la ejecución operativa del conteo físico en sitio (eso lo define el servicio
[Inventarios físicos](../catalogo-solutions.md#11-inventarios-físicos) de Solutions como proceso propio, con
su propio alcance de cliente/ubicación/ventana de tiempo). No incluye la decisión de qué activo específico
asignar de ese saldo disponible, que es [SOP-ACTIVO-10 · Asignación](./SOP-ACTIVO-10-asignacion.md).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Sistema (automático) | Recalcula el saldo de inventario digital ante cada evento de estado/ubicación registrado en [trazabilidad](../catalogo-systems.md#7-trazabilidad); ejecuta la transición `Liberado → Disponible` (transición #10) cuando no hay intervención manual configurada. |
| Responsable de inventario `[VALIDAR nombre exacto del puesto]` | Supervisa el saldo, confirma manualmente el ingreso a `Disponible` cuando la regla no está automatizada, y encabeza la conciliación cuando hay discrepancia. |
| Operador del servicio Inventarios físicos (Solutions) `[VALIDAR si es rol interno o de un tercero contratado]` | Ejecuta el conteo físico según [catálogo Solutions #11](../catalogo-solutions.md#11-inventarios-físicos) y entrega el reporte de conteo vs. contable. |
| Gobernanza `[VALIDAR]` | Revisa discrepancias que superan el umbral definido `[VALIDAR umbral]`, decide si hay incidencia formal (SOP-15) y ajusta reglas operativas si la causa es sistemática. |

## 5. Diagrama de flujo

```
[Evento de ciclo de vida ocurre: Liberación, Asignación, Retorno,
 Movimiento, Baja, Reincorporación tras mantenimiento]
                ↓
[Sistema registra el evento en trazabilidad y actualiza el estado del activo]
                ↓
[Inventario digital recalcula el saldo por ubicación, tipo y estado]
                ↓
[Saldo "Disponible" queda actualizado — consultable en cualquier corte de fecha]
                ↓
        ¿Se ejecuta un inventario físico (servicio Solutions #11)?
        ↓ no (operación normal, sin acción adicional)     ↓ sí (auditoría puntual)
   [El ciclo continúa]                          [Contrastar conteo físico
                                                   vs. saldo del sistema]
                                                              ↓
                                                    ¿Coincide el conteo?
                                                    ↓ sí            ↓ no
                                          [Cerrar conciliación   [Investigar causa: evento no
                                           sin ajuste,             registrado, activo sin
                                           dejar evidencia]        identificar, error de
                                                                    captura, pérdida/robo]
                                                                          ↓
                                                                [Ajustar saldo digital con
                                                                 evidencia del ajuste y
                                                                 su causa]
                                                                          ↓
                                                                [¿Discrepancia supera el
                                                                 umbral? → registrar
                                                                 incidencia (SOP-15) y
                                                                 escalar a Gobernanza]
```

## 6. SOP paso a paso

### A. Mantenimiento continuo del saldo

1. Recibir el evento generado por cualquier otra etapa del ciclo (liberación, asignación, retorno,
   movimiento, baja, reincorporación) a través de [trazabilidad](../catalogo-systems.md#7-trazabilidad).
2. Confirmar que el evento actualizó correctamente el estado del activo en la
   [máquina de estados](../estados-del-activo.md#1-catálogo-de-estados) antes de recalcular el saldo — un
   saldo no puede ser más confiable que el estado del que depende.
3. Recalcular el saldo de inventario por ubicación, tipo y estado (Módulo 9 de
   [CHACONTAINER OS](../catalogo-systems.md#16-chacontainer-os)).
4. Cuando el evento es `Liberado → Disponible` (transición #10), confirmar el ingreso a inventario
   disponible — automático si hay [reglas operativas](../catalogo-systems.md#11-reglas-operativas)
   configuradas para ese tipo de activo, manual si no `[VALIDAR: qué tipos de activo tienen esta regla
   automatizada hoy]`.
5. Dejar el saldo consultable en cualquier corte de fecha, sin depender de que alguien lo recalcule a mano.
6. Definir y respetar la **frecuencia de revisión del saldo** por parte del responsable de inventario, más
   allá del recálculo automático por evento `[VALIDAR frecuencia: diaria, por turno, continua]`.

### B. Conciliación contra inventario físico

7. Cuando se ejecuta el servicio [Inventarios físicos](../catalogo-solutions.md#11-inventarios-físicos)
   (propio o de un cliente), recibir su reporte de conteo físico vs. contable.
8. Contrastar el conteo físico contra el saldo `Disponible` del sistema para la misma ubicación/tipo/ventana
   de tiempo.
9. Si el conteo coincide con el saldo, cerrar la conciliación sin ajuste y dejar evidencia de que se validó
   (fecha, alcance, resultado "sin discrepancia").
10. Si hay discrepancia, investigar la causa antes de ajustar: evento no registrado en trazabilidad, activo
    sin identificación (no escaneado en el conteo), error de captura en el sistema, o pérdida/robo real.
11. Ajustar el saldo digital al valor confirmado por el conteo físico, dejando evidencia del ajuste: qué se
    cambió, por qué y quién lo autorizó `[VALIDAR nivel de autorización requerido para un ajuste manual de
    saldo]`.
12. Si la discrepancia supera el umbral definido `[VALIDAR umbral]`, registrar una incidencia formal
    (SOP-15, [Gestión de incidencias](../catalogo-systems.md#10-gestión-de-incidencias)) y escalar a
    Gobernanza, especialmente si el patrón se repite en la misma ubicación o cliente
    ([ETAPA 4](../README.md#etapa-4--modelo-de-gobernanza)).
13. Definir la **frecuencia de conciliación** (con qué periodicidad se ejecuta un inventario físico sobre el
    parque propio o de cada cliente, más allá de las auditorías puntuales solicitadas) `[VALIDAR frecuencia]`.

## 7. Checklist

- [ ] Todo evento de cambio de estado/ubicación generó una actualización de trazabilidad antes de recalcular
      el saldo.
- [ ] Saldo `Disponible` por ubicación y tipo consultable en tiempo real (no requiere recálculo manual para
      estar al día).
- [ ] Transición `Liberado → Disponible` ejecutada (automática o manual) sin activos "atascados" en
      `Liberado`.
- [ ] Al ejecutar un inventario físico, el conteo se contrastó explícitamente contra el saldo del sistema
      (no se archivó como dato aislado).
- [ ] Toda discrepancia detectada tiene una causa investigada antes de ajustar el saldo — no se ajusta "a
      ciegas" para que cuadre.
- [ ] Todo ajuste de saldo queda con evidencia (qué, por qué, quién lo autorizó).
- [ ] Discrepancias sobre el umbral generaron incidencia formal y fueron escaladas a Gobernanza.

## 8. KPI

- **Disponibilidad real**: inventario físicamente utilizable / inventario contable × 100 (ver
  [README ETAPA 6](../README.md#etapa-6--indicadores-del-sistema)).
- **% de discrepancia físico-contable** por conciliación ejecutada.
- **Tiempo de resolución de discrepancias**: desde detectada hasta ajustada y, si aplica, cerrada como
  incidencia.
- **Frecuencia real de conciliación** vs. la frecuencia objetivo `[VALIDAR frecuencia objetivo]`.
- **% de activos sin identificación** detectados durante conciliación (ver
  [catálogo Solutions #11](../catalogo-solutions.md#11-inventarios-físicos)) — mide el riesgo estructural
  detrás de futuras discrepancias.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Saldo digital desincronizado porque un evento del ciclo (retorno, baja) no se registró en trazabilidad | Todo SOP del ciclo de vida está obligado a registrar su evento antes de considerarse cerrado (ver invariante 5 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)); auditar periódicamente eventos huérfanos. |
| Ajuste de saldo hecho "a ciegas" para que el sistema cuadre con el conteo, sin investigar la causa | Paso 10-11 obligan a documentar la causa antes de ajustar; ningún ajuste se acepta sin evidencia. |
| Frecuencia de conciliación no definida — el saldo digital se desvía del real por meses sin que nadie lo note | `[VALIDAR frecuencia de conciliación]`; mientras no se defina, tratar como riesgo abierto y no asumir que el saldo digital es confiable sin verificación reciente. |
| Activo sin identificación (QR/RFID) no se cuenta en el conteo físico ni se puede vincular al saldo digital | Priorizar [Identificación](../catalogo-systems.md#3-identificación) del parque antes de confiar en la disponibilidad real reportada; registrar el hallazgo como insumo para vender identificación (ver [catálogo Solutions, lectura cruzada](../catalogo-solutions.md#lectura-cruzada-de-dónde-entra-cada-servicio-a-systems)). |
| Confundir este SOP con el servicio comercial de Inventarios físicos y tratar la conciliación como un evento aislado en vez de un proceso recurrente de mantenimiento del saldo | Este documento — dejar explícita la distinción (ver nota al inicio) en la capacitación de todo responsable de inventario. |

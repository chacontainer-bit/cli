# SOP-ACTIVO-03: Inspección

**Estado: Borrador v0.1 — pendiente de validación.**

Este SOP documenta el estado `En inspección` de la [máquina de
estados](../estados-del-activo.md#1-catálogo-de-estados) del Proyecto Maestro
CHACONTAINER. A diferencia de la mayoría de los SOP-ACTIVO-XX (que cubren una
transición limpia de un estado a otro), `En inspección` es un estado con
**múltiples puntos de entrada y salida** — el activo pasa por él más de una
vez durante su ciclo de vida, con un propósito distinto cada vez:

| Momento | Entra desde | Transición | Propósito |
|---|---|---|---|
| **Inspección inicial** | `Identificado` | #3 | Primera evaluación de condición de un activo recién identificado, antes de que haya circulado. |
| **Inspección de validación (post-lavado)** | `Sucio` | #7 | Confirmar que el lavado dejó el activo apto, o detectar daño que el lavado expuso. |
| **Inspección de validación (post-reparación)** | `En reparación` | #8 | Confirmar que la reparación/reacondicionamiento fue efectivo antes de liberar el activo. |
| **Reinspección** | `En tránsito` | #15 | Evaluar la condición de un activo que retorna de campo, antes de decidir su siguiente paso. |

Este documento cubre el proceso de evaluación en sí (el checklist de
aceptación/rechazo y la lógica de dictamen), aplicable a los cuatro momentos
de la tabla — el paso 1 del [§6](#6-sop-paso-a-paso) identifica cuál de los
cuatro aplica en cada ejecución. Se construyó a partir del cruce entre el
[catálogo Solutions](../catalogo-solutions.md#9-inspección), el [ciclo de
vida del activo](../ciclo-de-vida-del-activo.md) y la [máquina de
estados](../estados-del-activo.md) — **no a partir de una entrevista real
con el responsable operativo**. No debe declararse estándar hasta validar en
operación los puntos marcados `[VALIDAR]`.

El checklist del [§7](#7-checklist) toma como referencia de tono y nivel de
detalle el checklist de inspección post-lavado de
[SOP-02](../../chacontainer/docs/sop/SOP-02-lavado-de-activos-retornables.md)
del SaaS `chacontainer/` — **no es una copia**: SOP-02 cubre un checklist
acotado a limpieza tras el lavado, mientras que este documento cubre la
evaluación de condición completa (estructural, funcional, normativa) en los
cuatro momentos de la tabla anterior, y usa exclusivamente la terminología de
estados de [estados-del-activo.md](../estados-del-activo.md).

---

## 1. Resumen ejecutivo

Inspección es el proceso que evalúa **un activo individual** contra un
criterio de aceptación documentado y determina su siguiente paso en el
ciclo: apto para liberar, requiere lavado, o requiere reparación. Ocurre
cada vez que el activo entra al estado `En inspección`, en cuatro momentos
distintos del ciclo (inspección inicial, dos variantes de validación
post-intervención, y reinspección tras retorno — ver tabla introductoria).
Es, según el [catálogo Solutions](../catalogo-solutions.md#9-inspección), el
punto donde se genera el dato de **Capa 4 · Estado** con mayor
confiabilidad de todo el ciclo: alimenta directamente el modelo de
gobernanza (¿cuántos activos están dañados?, ¿cuál es su condición real?).
Inspección evalúa activos uno por uno; cuando lo que llega es un lote
heterogéneo sin separar, ese lote pasa primero por
[Clasificación](./SOP-ACTIVO-04-clasificacion-de-condicion.md) antes de que
cada activo (o subgrupo) llegue a este SOP.

## 2. Objetivo

Garantizar que todo activo que entra al estado `En inspección` — sin
importar desde cuál de los cuatro momentos — es evaluado contra un checklist
de aceptación documentado y consistente por tipo de activo, y que el
dictamen (apto / requiere lavado / requiere reparación) queda registrado y
trazable antes de disparar la transición de estado correspondiente, sin
depender del criterio informal de un solo inspector.

## 3. Alcance

Aplica a todo activo retornable (contenedor, IBC, tarima, tambor) que se
encuentra en estado `En inspección`, en cualquiera de los cuatro momentos
descritos en la tabla introductoria. Incluye la evaluación individual del
activo, el registro del dictamen y la evidencia asociada, y el disparo de la
transición de estado que corresponda (`Sucio`, `En reparación` o
`Liberado`).

**No incluye:**

- La **Clasificación** de un lote heterogéneo por tipo/condición/propietario
  — eso es un proceso distinto que opera sobre el lote completo, no sobre un
  activo individual (ver [SOP-ACTIVO-04](./SOP-ACTIVO-04-clasificacion-de-condicion.md#3-alcance),
  que deja explícita la distinción). Cuando un retorno llega como lote
  heterogéneo, Clasificación se ejecuta primero para separar el lote en
  subgrupos manejables, y **cada activo dentro de esos subgrupos** pasa
  después por este SOP para su dictamen individual definitivo.
- La ejecución física del lavado ([SOP-ACTIVO-05](./SOP-ACTIVO-05-lavado.md))
  o de la reparación ([SOP-ACTIVO-06](./SOP-ACTIVO-06-reparacion.md)) — este
  SOP solo decide si el activo necesita esos procesos, no los ejecuta.
- La confirmación final de liberación e ingreso a inventario disponible
  (SOP-ACTIVO-08, pendiente de redactar) — este SOP dispara la transición a
  `Liberado`, pero el paso de `Liberado` a `Disponible` es un checkpoint
  separado (invariante 2 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)).
- La apertura y seguimiento de una incidencia cuando la inspección detecta
  una disputa de custodia o condición — eso corresponde a Gestión de
  incidencias (SOP-15, pendiente); este SOP solo dispara el registro inicial
  (ver paso 8 del [§6](#6-sop-paso-a-paso)).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Inspector `[VALIDAR nombre exacto del puesto]` | Evalúa el activo contra el checklist de aceptación, registra el dictamen y dispara la transición de estado correspondiente (fila 3, 4, 5, 6, 7, 8 de la [tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones)). |
| Recepción `[VALIDAR si es el mismo rol que Inspector o uno distinto]` | Recibe el activo que llega de un retorno (`En tránsito`) y lo ingresa al flujo de reinspección (fila 15 de la tabla de transiciones). |
| Responsable de planta / Gobernanza `[VALIDAR división exacta de autoridad]` | Define y actualiza el criterio de aceptación por tipo de activo, resuelve excepciones y disputas escaladas, y autoriza casos límite no cubiertos por el checklist. |
| Operador de lavado / Técnico de reparación | Entregan el activo a este SOP tras completar su propia intervención (transiciones #7 y #8); no ejecutan la inspección, pero su registro previo (evidencia, insumos usados) es insumo para el dictamen de validación. |
| Sistema (CHACONTAINER OS) | Registra automáticamente cada transición de estado en trazabilidad con marca de tiempo, actor y estado anterior/nuevo, conforme al invariante 5 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar). |

## 5. Diagrama de flujo

```
[Activo entra a "En inspección" desde uno de cuatro momentos]
        │
        ├─ Identificado ─────────────► (#3)  Inspección inicial
        ├─ Sucio (lavado completado) ─► (#7)  Validación post-lavado
        ├─ En reparación (reparación
        │   completada) ──────────────► (#8)  Validación post-reparación
        └─ En tránsito (retorno
            a planta) ─────────────────► (#15) Reinspección
                        ↓
        [Identificar el motivo de la inspección (cuál de los 4 momentos)]
                        ↓
        [Si viene de un retorno masivo sin separar, confirmar que ya
         pasó por Clasificación (SOP-ACTIVO-04) antes de continuar]
                        ↓
        [Identificar tipo de activo; si es reinspección, contrastar
         contra la condición registrada en la salida]
                        ↓
        [Evaluar contra el checklist de aceptación por tipo (§7)]
                        ↓
                ¿Cumple criterio de aceptación?
                │
                ├─ Sí, sin intervención pendiente ──► Liberado (#6)
                │                                      → SOP-ACTIVO-08 (pendiente)
                │
                ├─ No, requiere limpieza ────────────► Sucio (#4)
                │                                      → SOP-ACTIVO-05 (Lavado)
                │
                └─ No, hay daño estructural/
                    funcional ──────────────────────► En reparación (#5)
                                                        → SOP-ACTIVO-06 (Reparación)
```

## 6. SOP paso a paso

1. Identificar cuál de los cuatro momentos de inspección aplica (ver tabla
   introductoria): inicial (`Identificado`), validación post-lavado
   (`Sucio`), validación post-reparación (`En reparación`), o reinspección
   (`En tránsito`).
2. Si el activo llega como parte de un lote heterogéneo desde un retorno
   masivo, confirmar que el lote ya pasó por
   [Clasificación](./SOP-ACTIVO-04-clasificacion-de-condicion.md); si no,
   derivarlo a ese proceso antes de continuar con la inspección individual.
3. Registrar el inicio de la inspección en el sistema, vinculando el evento
   a la trazabilidad del activo (ver
   [catálogo Systems, Trazabilidad](../catalogo-systems.md#7-trazabilidad)).
4. Identificar el tipo de activo (contenedor, IBC, tarima, tambor) y, si es
   reinspección o validación post-intervención, consultar el historial de
   condición registrado previamente (salida, lavado o reparación) para
   contrastar contra el estado actual.
5. Evaluar el activo contra el checklist de aceptación por tipo de activo
   (ver [§7 Checklist](#7-checklist)): integridad estructural, componentes
   funcionales, limpieza aparente, etiquetado y, si aplica, cumplimiento
   normativo `[VALIDAR criterio exacto de aceptación/rechazo por tipo de
   activo — pendiente de definición operativa]`.
6. Registrar el dictamen de la evaluación: **apto** (sin intervención
   pendiente), **requiere lavado**, o **requiere reparación**.
7. Dejar evidencia fotográfica y el checklist registrado en el sistema,
   especialmente cuando el dictamen es distinto a "apto".
8. Si la inspección detecta una disputa de condición o custodia (p. ej. el
   cliente reclama que el daño no es atribuible al uso), registrar el
   hallazgo como incidencia (SOP-15, pendiente) y escalar al Responsable de
   planta / Gobernanza antes de cerrar el dictamen — conforme al invariante
   3 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar).
9. Disparar la transición de estado correspondiente al dictamen:
   - Apto → `Liberado` (transición #6).
   - Requiere lavado → `Sucio` (transición #4).
   - Requiere reparación → `En reparación` (transición #5).
10. Confirmar que la transición quedó registrada en trazabilidad con marca
    de tiempo, actor y estado anterior/nuevo (invariante 5), y actualizar
    el indicador de tiempo de inspección y, si aplica, el de tasa de daño
    por cliente/ruta/proveedor (ver [§8 KPI](#8-kpi)).
11. Recordar que un dictamen "apto" nunca mueve el activo directamente a
    `Disponible` — pasa obligatoriamente por `Liberado` primero (invariante
    2 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)).

## 7. Checklist

- [ ] Motivo de la inspección identificado (inicial / validación post-lavado
      / validación post-reparación / reinspección).
- [ ] Si el activo viene de un lote heterogéneo, confirmado que ya pasó por
      Clasificación (SOP-ACTIVO-04) antes de esta evaluación individual.
- [ ] Inicio de la inspección registrado en trazabilidad.
- [ ] Tipo de activo identificado y, si es reinspección o validación,
      contrastado contra el historial de condición previo.
- [ ] Identificación (QR/RFID) legible y vinculada al ID correcto del
      registro maestro.
- [ ] Integridad estructural verificada: sin fisuras, deformaciones ni
      perforaciones visibles que comprometan el uso `[VALIDAR tolerancia
      máxima aceptable por tipo de activo]`.
- [ ] Componentes funcionales (válvulas, tapas, bisagras, ruedas, empaques,
      según tipo) presentes y operativos `[VALIDAR lista exacta de
      componentes críticos por tipo de activo]`.
- [ ] Limpieza aparente evaluada: si no cumple, se deriva a `Sucio`, no se
      fuerza un dictamen "apto" con reservas.
- [ ] Etiquetado/rotulado vigente, legible y sin residuo de etiquetado
      anterior no removido.
- [ ] Cumplimiento normativo aplicable verificado, cuando corresponde
      (materiales peligrosos, alimentos, etc.) `[VALIDAR normativa
      aplicable por tipo de producto/cliente]`.
- [ ] Evidencia fotográfica capturada, especialmente en dictámenes distintos
      a "apto".
- [ ] Dictamen (apto / requiere lavado / requiere reparación) registrado en
      el sistema el mismo día de la evaluación.
- [ ] Transición de estado correspondiente disparada y reflejada en
      trazabilidad (invariante 5).
- [ ] Si se detectó disputa de condición/custodia, incidencia registrada y
      escalada antes de cerrar el dictamen.

## 8. KPI

- % de activos por categoría de condición resultante (apto / requiere
  lavado / requiere reparación) — ver [catálogo Solutions,
  Inspección](../catalogo-solutions.md#9-inspección).
- Tasa de daño por cliente/ruta/proveedor, cruzada con [Analítica (Systems
  #14)](../catalogo-systems.md#14-analítica) para responder "dónde se
  producen las pérdidas/daños" ([ETAPA 4 del
  README](../README.md#etapa-4--modelo-de-gobernanza)).
- Tiempo de inspección por activo, y tiempo promedio en el estado `En
  inspección` (detecta cuello de botella cuando un activo queda "atrapado"
  sin dictamen).
- % de reinspecciones con dictamen distinto al de la inspección/validación
  anterior — mide consistencia del criterio entre inspectores y entre
  momentos del ciclo.
- % de activos que, tras `Liberado`, generan una incidencia posterior por
  daño no detectado — mide la calidad real del checklist de aceptación.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Dependencia de un solo inspector que concentra el criterio de aceptación (mismo riesgo documentado en SOP-02 del SaaS para lavado) | Documentar por escrito el checklist de aceptación/rechazo por tipo de activo (§7) y capacitar a más de un inspector. |
| Criterio inconsistente entre los cuatro momentos de inspección (p. ej. la reinspección es más laxa que la inicial) | Un único checklist de aceptación por tipo de activo, aplicado sin variación según el momento — el motivo de la inspección (paso 1) se registra, pero no cambia el criterio de evaluación. |
| Activo liberado sin pasar por el checklist completo, por presión de tiempo o de despacho | El dictamen "apto" solo se registra tras completar el checklist del §7; ninguna transición a `Liberado` (#6) se dispara sin el registro completo. |
| Activo "apto" que en realidad tenía daño no detectado en la evaluación visual | Checklist estructurado por componente (no solo "a simple vista"), evidencia fotográfica obligatoria, y el KPI de incidencias post-`Liberado` (§8) como control retrospectivo. |
| Falta de criterio exacto de aceptación por tipo de activo (hoy sin definir, marcado `[VALIDAR]`) | No inventar el criterio: operar con el criterio documentado que exista, escalar al Responsable de planta cualquier caso sin criterio claro, y cerrar el `[VALIDAR]` antes de declarar este SOP como estándar. |
| Disputa de condición/custodia resuelta informalmente sin dejar registro | Toda disputa detectada en inspección se registra como incidencia (paso 8) y se escala, conforme al invariante 3 de estados-del-activo.md — no se resuelve "de palabra" entre inspector y custodio. |
| Transiciones automáticas de reglas operativas mal configuradas derivan en un dictamen sin base | Las transiciones que dependan de reglas automáticas (p. ej. umbrales de aceptación configurados en Systems) requieren regla explícita por tipo de activo/cliente; sin regla configurada, no se ejecuta con un valor por defecto arbitrario (invariante 6 de estados-del-activo.md). |

---

**Documentos relacionados:** [catalogo-solutions.md](../catalogo-solutions.md) ·
[catalogo-systems.md](../catalogo-systems.md) ·
[ciclo-de-vida-del-activo.md](../ciclo-de-vida-del-activo.md) ·
[estados-del-activo.md](../estados-del-activo.md) ·
[SOP-ACTIVO-02 · Identificación QR/RFID](./SOP-ACTIVO-02-identificacion-qr-rfid.md) ·
[SOP-ACTIVO-04 · Clasificación de condición](./SOP-ACTIVO-04-clasificacion-de-condicion.md) ·
[SOP-ACTIVO-05 · Lavado](./SOP-ACTIVO-05-lavado.md) ·
[SOP-ACTIVO-06 · Reparación / reacondicionamiento](./SOP-ACTIVO-06-reparacion.md)

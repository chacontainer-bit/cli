# SOP-ACTIVO-06: Reparación / reacondicionamiento

**Estado: Borrador v0.1 — pendiente de validación.**

Este SOP documenta la etapa **Reparación / reacondicionamiento** del [ciclo
de vida del activo](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado)
del Proyecto Maestro CHACONTAINER: la transición del estado `En reparación`
hacia `En inspección` (si el activo queda apto) o hacia `Baja` (si no es
reparable), según la [máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones)
(transiciones #8 y #9). Corresponde a los servicios **Reparación** (#4) y
**Reacondicionamiento** (#5) del [catálogo Solutions](../catalogo-solutions.md),
y se apoya en **Gestión de incidencias** (#10) y **Trazabilidad** (#7) del
[catálogo Systems](../catalogo-systems.md) para registrar diagnóstico y
resultado.

Es un documento propio de este proyecto. Donde aplica, se referencia
[SOP-02](../../chacontainer/docs/sop/SOP-02-lavado-de-activos-retornables.md)
del SaaS `chacontainer/` solo como referencia de tono/formato, no como fuente
de contenido — ese SOP cubre lavado, no reparación, y usa una terminología de
estados distinta a la de [estados-del-activo.md](../estados-del-activo.md).

---

## 1. Resumen ejecutivo

Un activo llega a estado `En reparación` cuando una inspección (SOP-ACTIVO-03,
pendiente) detecta daño que le impide reingresar a inventario disponible tal
como está (transición #5: `En inspección` → `En reparación`). Este SOP cubre
lo que ocurre mientras el activo está en ese estado: diagnóstico del daño,
decisión de si conviene repararlo puntualmente o reacondicionarlo
integralmente, ejecución de la intervención, y el resultado final —
transición a `En inspección` para validación final (transición #8) si quedó
apto, o a `Baja` (transición #9) si se determina que no es reparable.

## 2. Objetivo

Garantizar que todo activo en estado `En reparación` recibe un diagnóstico
documentado, una decisión de reparar-vs-dar de baja basada en un criterio
explícito (no solo en la experiencia del técnico de turno), y que el
resultado de la intervención queda registrado y trazable antes de que el
activo salga de ese estado.

## 3. Alcance

Aplica a todo activo retornable en estado `En reparación`, sin importar la
vía de entrada:

- Vía inspección inicial o reinspección: `En inspección` → `En reparación`
  (transición #5).
- Vía incidencia reportada en campo (daño detectado durante uso, retorno o
  logística inversa) que deriva a reparación conforme a SOP-15 (Incidencias).

Este SOP cubre **dos modalidades de intervención** que el [catálogo
Solutions](../catalogo-solutions.md) distingue explícitamente y que este
documento trata bajo el mismo estado `En reparación` pero con alcance
distinto:

| Modalidad | Definición (según catálogo Solutions) | Cuándo aplica |
|---|---|---|
| **Reparación puntual** ([Solutions #4](../catalogo-solutions.md#4-reparación)) | Daño estructural o funcional localizado que impide el reingreso del activo tal como está. | El activo tiene una falla identificable y acotada (p. ej. una válvula, un panel, una soldadura). |
| **Reacondicionamiento integral** ([Solutions #5](../catalogo-solutions.md#5-reacondicionamiento)) | El activo es funcional pero está degradado por uso/tiempo de forma generalizada (no un daño puntual); se interviene integralmente para devolverlo a condición "como nuevo" o al estándar acordado. | El desgaste es general (múltiples puntos, estético y/o estructural menor) y no se resuelve con una intervención localizada. |

La decisión de qué modalidad aplica se toma en el diagnóstico (paso 1 del
§6), no antes. Un activo puede iniciar el ciclo como candidato a reparación
puntual y, tras el diagnóstico, resultar en que conviene reacondicionamiento
integral (o viceversa, si el reacondicionamiento revela que basta una
reparación puntual) `[VALIDAR: si existe un umbral objetivo de número de
puntos dañados o % de superficie afectada que obligue a pasar de reparación
puntual a reacondicionamiento integral, o si la decisión queda a criterio
del técnico/responsable de planta]`.

**No incluye:**

- El lavado del activo antes o después de la reparación (cubierto por
  [SOP-ACTIVO-05](./SOP-ACTIVO-05-lavado.md)), salvo la limpieza que el
  reacondicionamiento integral incorpore como parte de su propia
  intervención `[VALIDAR: si el reacondicionamiento ejecuta su propio paso
  de limpieza o delega a SOP-ACTIVO-05]`.
- El reetiquetado o modificación del activo (SOP-ACTIVO-07, fuera de
  alcance de este documento), aunque el reacondicionamiento integral pueda
  derivar en una venta de esos servicios (ver [Solutions #5](../catalogo-solutions.md#5-reacondicionamiento),
  campo "servicio que puede venderse después").
- La ejecución de la baja una vez decidida (SOP-18/19/20, fuera de alcance
  — este SOP solo dispara la transición a `Baja`, no gestiona scrap ni
  disposición final).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Técnico de reparación `[VALIDAR]` | Diagnostica el daño, ejecuta la reparación o el reacondicionamiento, y registra el resultado antes/después. |
| Responsable de planta / Gobernanza `[VALIDAR división exacta de autoridad entre ambos roles]` | Decide reparar-vs-dar de baja cuando el costo o la viabilidad técnica no son evidentes, y autoriza la transición a `Baja` (transición #9 exige explícitamente este rol según la tabla de transiciones). |
| Comercial `[VALIDAR]` | Cotiza la intervención cuando el costo aplica al cliente (activo propiedad del cliente, o reparación fuera de garantía). |
| Sistema (CHACONTAINER OS) | Registra la transición de estado y, si el diagnóstico proviene de una incidencia, vincula el registro de [Gestión de incidencias](../catalogo-systems.md#10-gestión-de-incidencias). |

## 5. Diagrama de flujo

```
[Activo en estado "En reparación"]
              ↓
[Diagnosticar daño: puntual vs. degradación generalizada]
              ↓
       ¿Reparación puntual o reacondicionamiento integral?
       ↓ puntual                          ↓ integral
[Reparar el punto de falla          [Intervenir integralmente:
 identificado]                       múltiples puntos + limpieza
       ↓                             profunda + estándar acordado]
       └───────────────┬──────────────────┘
                        ↓
        [Registrar evidencia antes/después + insumos usados]
                        ↓
              ¿Resultado: activo reparable?
              ↓ sí                        ↓ no
   [Transición: En reparación      [Transición: En reparación
    → En inspección]                → Baja]
        (validación final,              (SOP-ACTIVO-06 → SOP-18,
         transición #8)                  transición #9, requiere
                                          autorización de Responsable
                                          de planta / Gobernanza)
```

## 6. SOP paso a paso

1. Diagnosticar el daño del activo en estado `En reparación`: registrar
   tipo de falla, ubicación, y si el patrón corresponde a daño puntual o a
   degradación generalizada (define reparación vs. reacondicionamiento
   integral según §3).
2. Si el diagnóstico proviene de una incidencia (SOP-15), vincular el
   registro de reparación al ticket de incidencia correspondiente en
   [Gestión de incidencias](../catalogo-systems.md#10-gestión-de-incidencias).
3. Cotizar la intervención cuando el costo aplique al cliente
   `[VALIDAR criterio de cuándo el costo es del cliente vs. de CHACONTAINER
   — depende de si el activo es propio, rentado o vendido, y de condiciones
   de garantía]`.
4. Aplicar el criterio de "no reparable" antes de ejecutar cualquier
   intervención, para evitar invertir recursos en un activo que de todos
   modos se dará de baja `[VALIDAR: criterio objetivo de no-reparabilidad
   — hoy no está definido; candidatos típicos: costo de reparación mayor a
   un % del valor de reemplazo, daño estructural que compromete seguridad,
   o pieza sin refacción disponible]`.
5. Si el activo es reparable, ejecutar la intervención según la modalidad
   determinada en el paso 1:
   - **Reparación puntual:** intervenir el punto de falla identificado.
   - **Reacondicionamiento integral:** intervenir de forma integral
     (múltiples puntos, limpieza profunda, repintado/reetiquetado si
     aplica) hasta el estándar acordado `[VALIDAR estándar "como nuevo" —
     definirlo por tipo de activo]`.
6. Registrar evidencia fotográfica antes/después, refacciones o insumos
   usados, y tiempo invertido en la intervención.
7. Si el activo no es reparable (o el costo de reparación no se justifica),
   escalar la decisión al Responsable de planta / Gobernanza para
   autorizar la transición a `Baja` — este paso requiere explícitamente
   ese rol según la transición #9 de la tabla de transiciones.
8. Registrar el resultado de la intervención (reparado / no reparable) en
   el sistema, con marca de tiempo, actor y evidencia.
9. Disparar la transición de estado correspondiente:
   - Si el activo quedó apto: `En reparación` → `En inspección`
     (transición #8), para validación final antes de `Liberado`.
   - Si no es reparable: `En reparación` → `Baja` (transición #9), que
     inicia el flujo hacia SOP-18 (Baja) fuera del alcance de este
     documento.
10. Confirmar que la transición quedó registrada en trazabilidad conforme
    al invariante 5 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar) —
    ninguna transición de estado se considera completa sin este registro.

## 7. Checklist

- [ ] Diagnóstico documentado: tipo de daño, ubicación, y clasificación
      puntual vs. generalizada.
- [ ] Modalidad de intervención determinada (reparación puntual o
      reacondicionamiento integral) y justificada.
- [ ] Vínculo con ticket de incidencia registrado, si el origen fue una
      incidencia (SOP-15).
- [ ] Cotización generada, si el costo aplica al cliente.
- [ ] Criterio de "no reparable" aplicado antes de ejecutar la
      intervención (evita reparar lo que de todos modos se dará de baja).
- [ ] Evidencia antes/después e insumos/refacciones usados registrados.
- [ ] Si el resultado es "no reparable", autorización del Responsable de
      planta / Gobernanza documentada antes de transicionar a `Baja`.
- [ ] Transición de estado (`En inspección` o `Baja`) reflejada en
      trazabilidad el mismo día del cierre de la intervención.

## 8. KPI

- % de activos que ingresan a `En reparación` y salen como reparados
  (`En inspección`) vs. dados de `Baja`.
- Costo por reparación puntual vs. costo por reacondicionamiento integral
  (insumo directo para la decisión reparar-vs-reemplazar de la [ETAPA 4](../README.md#etapa-4--modelo-de-gobernanza)).
- Tiempo de ciclo en estado `En reparación`, por modalidad (puntual vs.
  integral).
- % de activos reacondicionados que, tras la intervención, requieren una
  segunda intervención en un plazo corto `[VALIDAR ventana de tiempo]` —
  indicador de calidad del reacondicionamiento.
- Tasa de recuperación vs. baja por tipo/familia de activo, cruzada con
  [Analítica (Systems #14)](../catalogo-systems.md#14-analítica) para
  responder "qué activos conviene reparar" ([ETAPA 4](../README.md#etapa-4--modelo-de-gobernanza)).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Reparar un activo cuyo costo de intervención supera su valor de reemplazo | Aplicar el criterio de no-reparabilidad (paso 4) antes de ejecutar la intervención, con umbral económico explícito `[VALIDAR]`. |
| Confundir reparación puntual con reacondicionamiento integral (subestimar el alcance) | Diagnóstico documentado como paso obligatorio previo a cualquier intervención (paso 1), no una decisión implícita del técnico en el momento. |
| Decisión de baja tomada sin autoridad definida | La transición #9 (`En reparación` → `Baja`) requiere explícitamente Responsable de planta / Gobernanza, nunca al criterio único del técnico. |
| Activo "reparado" con daño residual no detectado | La transición nunca va directo a `Disponible` — pasa obligatoriamente por `En inspección` para validación final (invariante 2 de estados-del-activo.md). |
| Dependencia de un solo técnico que concentra el criterio de diagnóstico | Documentar por escrito el criterio de clasificación puntual/integral y de no-reparabilidad; capacitar a más de un técnico. |
| Reparación sin vínculo a la incidencia que la originó | Registrar el vínculo con el ticket de [Gestión de incidencias](../catalogo-systems.md#10-gestión-de-incidencias) desde el inicio del diagnóstico (paso 2), para no perder el patrón de fallas recurrentes. |

---

**Documentos relacionados:** [catalogo-solutions.md](../catalogo-solutions.md) ·
[catalogo-systems.md](../catalogo-systems.md) ·
[ciclo-de-vida-del-activo.md](../ciclo-de-vida-del-activo.md) ·
[estados-del-activo.md](../estados-del-activo.md) ·
[SOP-ACTIVO-05 · Lavado](./SOP-ACTIVO-05-lavado.md)

# SOP-ACTIVO-08: Liberación

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre la etapa **Liberación** del [ciclo de vida del activo](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado):
el checkpoint de validación final que precede al ingreso a inventario disponible, formalizado en la
[máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones) como la transición #6
(`En inspección` → `Liberado`). Se construyó a partir del cruce entre el
[catálogo Systems](../catalogo-systems.md#11-reglas-operativas), el ciclo de vida y la máquina de estados —
**no a partir de una entrevista real con el responsable operativo**. No debe declararse estándar hasta
validar en operación los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Liberación es el checkpoint que confirma, antes de autorizar el ingreso de un activo a inventario
disponible, que este cumple las reglas de aceptación vigentes ([Reglas operativas — catálogo Systems #11](../catalogo-systems.md#11-reglas-operativas))
para su tipo y, si aplica, su cliente. No es una intervención física sobre el activo: es una decisión de
control que se apoya en la evidencia acumulada en las etapas previas (inspección, limpieza,
reparación/reacondicionamiento y, si aplicó, [reetiquetado — SOP-ACTIVO-07](./SOP-ACTIVO-07-reetiquetado.md)).
Si el activo cumple, transiciona a estado `Liberado`; si no cumple, regresa a la etapa que corresponda
(inspección, limpieza o reparación) en vez de avanzar por excepción.

## 2. Objetivo

Garantizar que ningún activo pasa al estado `Disponible` sin haber sido validado explícitamente contra las
reglas de aceptación vigentes, evitando que activos con hallazgos pendientes (condición física, limpieza,
etiquetado, identificación) se asignen a un cliente.

## 3. Alcance

Aplica a todo activo que llega al estado `En inspección` con resultado preliminar apto — ya sea de forma
directa (transición #6 de la [tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones): "apto
sin intervención") o después de pasar por `Sucio` (lavado), `En reparación` (reparación/reacondicionamiento)
y/o el proceso de [reetiquetado](./SOP-ACTIVO-07-reetiquetado.md). Cubre la decisión de aprobar (transición
a `Liberado`) o rechazar (retorno a la etapa correspondiente).

No incluye el ingreso físico a la zona de inventario disponible ni la actualización del saldo de inventario
digital, que es la transición siguiente (`Liberado` → `Disponible`, tabla de transiciones fila 10,
SOP-ACTIVO-09 `[VALIDAR]`). No incluye la ejecución de la reparación, el lavado ni el reetiquetado mismos
(SOP separados) — este SOP solo verifica su resultado.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Inspector de calidad | Autoridad primaria de liberación: verifica el activo contra las reglas de aceptación y decide aprobar o rechazar (tabla de transiciones, fila 6) `[VALIDAR si es el mismo rol que ejecutó la inspección inicial, o un rol distinto/superior tipo "cuatro ojos"]`. |
| Responsable de reglas operativas `[VALIDAR]` | Configura y mantiene actualizadas las reglas de aceptación por tipo de activo y, si aplica, por cliente ([catálogo Systems #11](../catalogo-systems.md#11-reglas-operativas)). |
| Responsable de planta / Gobernanza | Resuelve excepciones y casos límite; interviene cuando el hallazgo no encaja claramente en una regla configurada. |
| Operador de inventario `[VALIDAR]` | Recibe el activo `Liberado` y ejecuta o confirma la transición siguiente hacia `Disponible` (SOP-ACTIVO-09). |

## 5. Diagrama de flujo

```
[Activo llega a "En inspección" con resultado preliminar apto]
   (directo, o tras Sucio→Lavado / En reparación→Reparación /
    Reetiquetado SOP-ACTIVO-07)
                ↓
[Consultar reglas de aceptación vigentes (Systems #11)
 para el tipo de activo / cliente]
                ↓
[Verificar el activo contra cada regla: condición física,
 limpieza, etiquetado/normativa, identificación, evidencia documental]
                ↓
        ¿Cumple todas las reglas de aceptación?
        ↓ sí                                    ↓ no
[Inspector de calidad aprueba          [Registrar hallazgo específico
 → estado "Liberado"]                   → devolver a la etapa que
        ↓                                corresponda: En inspección /
[Registrar evidencia de liberación      Sucio / En reparación /
 en trazabilidad]                       Reetiquetado]
        ↓
[Enviar a SOP-ACTIVO-09 · Inventario disponible
 (transición "Liberado" → "Disponible")]
```

## 6. SOP paso a paso

1. Recibir el activo en estado `En inspección` con resultado preliminar apto, identificando de qué etapa
   proviene: sin intervención previa, o desde `Sucio` (SOP-ACTIVO-05 `[VALIDAR]`), `En reparación`
   (SOP-ACTIVO-06 `[VALIDAR]`) o [reetiquetado](./SOP-ACTIVO-07-reetiquetado.md).
2. Consultar las reglas de aceptación vigentes configuradas en
   [Reglas operativas](../catalogo-systems.md#11-reglas-operativas) para ese tipo de activo y, si aplica,
   ese cliente específico `[VALIDAR si las reglas de aceptación varían por cliente o son un único estándar
   global]`.
3. Verificar el activo contra cada regla configurada, como mínimo: condición física final (sin daño
   pendiente), limpieza (si pasó por `Sucio`), etiquetado y normativa aplicable (si pasó por
   [SOP-ACTIVO-07](./SOP-ACTIVO-07-reetiquetado.md)), identificación QR/RFID presente y vinculada al
   registro maestro, y evidencia documental completa de las etapas anteriores `[VALIDAR checklist
   exhaustivo de reglas de aceptación por tipo de activo]`.
4. Si el activo cumple todas las reglas, el Inspector de calidad (o el rol con autoridad definida, ver
   [§4 Roles](#4-roles)) aprueba la liberación: el activo transiciona a estado `Liberado` (tabla de
   transiciones, fila 6).
5. Si el activo **no** cumple una o más reglas, registrar el hallazgo específico (qué regla no se cumplió y
   por qué) y devolver el activo a la etapa que corresponda: `En inspección` (si requiere nueva evaluación
   general), `Sucio` (si el hallazgo es de limpieza), `En reparación` (si el hallazgo es de condición
   física) o al proceso de [reetiquetado](./SOP-ACTIVO-07-reetiquetado.md) (si el hallazgo es de etiquetado
   o normativa) — el activo **no** avanza a `Liberado` bajo ninguna excepción informal.
6. Registrar en [trazabilidad](../catalogo-systems.md#7-trazabilidad) el evento de liberación (o de
   rechazo) con fecha, actor y reglas verificadas, conforme al invariante 5 de
   [estados-del-activo.md §4](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar).
7. Enviar el activo liberado al flujo de Inventario disponible (SOP-ACTIVO-09 `[VALIDAR]`), donde ocurre la
   transición `Liberado` → `Disponible` (tabla de transiciones, fila 10) — automática (sistema) o manual
   (responsable de inventario), según lo que defina ese SOP.
8. Actualizar los indicadores de liberación (ver [§8 KPI](#8-kpi)).

## 7. Checklist

- [ ] Activo llega a este checkpoint en estado `En inspección` con resultado preliminar apto.
- [ ] Reglas de aceptación vigentes consultadas para el tipo de activo (y cliente, si aplica) antes de
      decidir.
- [ ] Activo verificado contra cada regla configurada (condición, limpieza, etiquetado, identificación,
      evidencia documental).
- [ ] Decisión de liberación tomada por el rol con autoridad definida — no por cualquier operador
      disponible.
- [ ] Si no cumple, hallazgo específico registrado y activo devuelto a la etapa correspondiente (no
      descartado sin registro).
- [ ] Evento de liberación (o rechazo) registrado en trazabilidad el mismo día.
- [ ] Activo liberado enviado al flujo de Inventario disponible sin quedar en `Liberado` indefinidamente.

## 8. KPI

- % de activos que pasan el checkpoint de liberación a la primera (sin devolución).
- Tiempo entre "En inspección: apto preliminar" y confirmación de `Liberado`.
- % de activos devueltos por hallazgo, desglosado por tipo de regla incumplida (condición / limpieza /
  etiquetado / identificación).
- Activos liberados por periodo — insumo directo del KPI **Disponibilidad** de la
  [ETAPA 6 del README](../README.md#etapa-6--indicadores-del-sistema).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Activo liberado sin verificar todas las reglas de aceptación ("liberación de confianza") | Checklist obligatorio contra reglas operativas configuradas (paso 3), sin excepción informal. |
| Reglas de aceptación no configuradas para un tipo de activo/cliente | Conforme al invariante 6 de [estados-del-activo.md §4](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar), la transición automática correspondiente debe quedar deshabilitada, no usar un umbral por defecto arbitrario `[VALIDAR comportamiento cuando falta la regla]`. |
| Autoridad de liberación concentrada en una sola persona (mismo riesgo que en lavado, ver [SOP-02 del SaaS](../../chacontainer/docs/sop/SOP-02-lavado-de-activos-retornables.md)) | `[VALIDAR]` definir si la liberación requiere un rol distinto al que ejecutó la inspección inicial. |
| Activo con hallazgo devuelto sin registro claro de la causa | Registro obligatorio del hallazgo específico (paso 5) antes de derivar a la etapa correspondiente. |
| Activo asignado a un cliente sin haber pasado por `Liberado` | El invariante 2 de [estados-del-activo.md §4](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar) ("ningún activo entra a `Disponible` sin pasar por `Liberado`") debe aplicarse como restricción dura en el sistema (Módulo 1 de [CHACONTAINER OS](../catalogo-systems.md#16-chacontainer-os)), no solo como proceso documentado. |

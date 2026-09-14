# KPI del sistema — especificación

**Fase:** I · Fundación — paso 6 (última pieza pendiente) de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

> **Nota de secuencia:** este documento debió cerrar la FASE I junto con los
> pasos 1-5. Se redactó después de la FASE II porque, en la práctica, cada
> KPI de la [ETAPA 6](./README.md#etapa-6--indicadores-del-sistema) necesita
> una fuente de datos concreta para dejar de ser una fórmula abstracta — y
> esa fuente solo quedó definida en [modelo-de-datos.md](./modelo-de-datos.md)
> (paso 8). El resultado es el mismo que si se hubiera hecho en orden: los
> 10 KPI de la ETAPA 6, ahora con fórmula, fuente de datos real (o brecha
> pendiente), cadencia y dueño.

Los 10 KPI ya están listados en la [ETAPA 6](./README.md#etapa-6--indicadores-del-sistema)
del proyecto maestro. Este documento los formaliza contra:
[modelo-de-datos.md](./modelo-de-datos.md) (¿de qué tabla sale el dato?),
[modelo-de-gobernanza.md](./modelo-de-gobernanza.md) (¿quién lo revisa y
con qué cadencia?), y el esquema real de `chacontainer/` (¿ya se puede
calcular hoy, o falta una tabla del roadmap de [arquitectura-os.md](./arquitectura-os.md)?).

---

## 1. Los 10 KPI, especificados

| KPI | Fórmula (ETAPA 6) | Fuente de datos | ¿Calculable hoy? | Cadencia / dueño (ver [modelo-de-gobernanza.md §2](./modelo-de-gobernanza.md#2-las-16-preguntas-con-dueño-y-cadencia)) |
|---|---|---|---|---|
| **Disponibilidad** | Activos disponibles / activos requeridos × 100 | `assets.status = 'available'` (numerador); "requeridos" no existe — depende de `operational_rules` (nivel mínimo) | 🟡 Parcial — numerador sí, denominador no | Continua / Nivel 1 |
| **Utilización** | Activos en uso / activos disponibles × 100 | `assets.status IN ('in_use')` vs. total activo | ✅ Sí, con el esquema real | Continua / Nivel 1 |
| **Tiempo de ciclo** | Tiempo desde salida hasta retorno | `asset_events` (evento de salida → evento de retorno) | ✅ Sí, con `asset_events.occurred_at` | Continua / Nivel 1 (consulta), Nivel 3 (tendencia) |
| **Cumplimiento de retorno** | Retornos en tiempo / retornos esperados × 100 | Requiere comparar `asset_events` (retorno real) contra `operational_rules` (punto de retorno esperado) | 🔴 No — depende de `operational_rules` (Fase C del roadmap de [arquitectura-os.md](./arquitectura-os.md#3-roadmap-de-implementación-propuesto)) | Semanal / Nivel 2-3 |
| **Pérdida** | Activos no recuperados / activos administrados × 100 | Requiere `asset_state_detail` con estado `Perdido` (no existe hoy — `assets.status = 'lost'` es la aproximación agregada) | 🟡 Aproximado — con el estado agregado `lost`, sin el detalle de `Retenido`→`Perdido` | Mensual / Nivel 3 |
| **Daño** | Activos dañados / activos retornados × 100 | Requiere `incidents` (tipo `daño`) o `asset_state_detail` con `En reparación` — ninguna existe hoy | 🔴 No — depende de Fase A y D del roadmap | Mensual / Nivel 3 |
| **Tiempo detenido** | Horas o días sin movimiento | `asset_events.occurred_at` más reciente vs. ahora — calculable con lo que ya existe | ✅ Sí | Diaria (para alertas) / Nivel 1-2 |
| **Costo por ciclo** | Costo total del sistema / ciclos completados | `maintenance_records.cost` cubre parte del costo; falta costo de logística inversa, recuperación, incidencias | 🟡 Parcial | Mensual / Nivel 3 |
| **Rotación** | Número de ciclos por activo | Se deriva de contar transiciones `Disponible`→`Asignado`→...→`Disponible` en `asset_events` | ✅ Sí, con el esquema real (requiere una consulta, no una tabla nueva) | Mensual / Nivel 3 |
| **Disponibilidad real** | Inventario físicamente utilizable, no solo contable | Requiere que `asset_state_detail` distinga `Disponible` de `Sucio`/`En reparación`/`Bloqueado` (que hoy caen todos en `assets.status = 'maintenance'` o `'lost'`, ver [modelo-de-datos.md §2](./modelo-de-datos.md#2-brecha-principal-assetsstatus-6-valores-vs-estados-del-activomd-16-estados)) | 🔴 No — es, de hecho, la razón de ser de `asset_state_detail` (Fase A del roadmap) | Continua / Nivel 1 |

**Resumen: 4 de 10 calculables hoy sin cambios de esquema, 3 parciales, 3
bloqueados por tablas que aún no existen** (`operational_rules`,
`asset_state_detail`, `incidents` — exactamente las tres piezas de las
Fases A, C y D del roadmap de [arquitectura-os.md](./arquitectura-os.md#3-roadmap-de-implementación-propuesto)).
Esto no es una coincidencia: los KPI que gobernanza necesita son,
literalmente, la razón por la que el roadmap técnico prioriza esas tablas.

## 2. Qué hacer mientras tanto (antes de que exista el esquema completo)

El piloto de la [ETAPA 7](./README.md#etapa-7--piloto-packaging-systems)
no tiene que esperar a que las 10 tablas existan. Los 4 KPI ya calculables
(Utilización, Tiempo de ciclo, Tiempo detenido, Rotación) son suficientes
para arrancar la **medición inicial** que pide la ETAPA 7; los 6 restantes
se incorporan según avance el roadmap de `arquitectura-os.md`, no bloquean
el inicio del piloto.

## 3. Umbrales de alerta (conectan con reglas-operativas.md)

Ningún KPI de esta lista tiene, por sí solo, un umbral de "bien/mal" —
eso lo define cada regla en [reglas-operativas.md](./reglas-operativas.md).
Ejemplo: **Disponibilidad real** baja no es una alerta hasta que cruza el
`nivel_minimo_unidades` configurado para ese tipo de activo/planta. Los KPI
de este documento son lo que se **mide**; las reglas operativas son lo que
convierte esa medición en una **acción**.

---

## Cierre de FASE I

Con este documento, los 6 pasos de la FASE I · Fundación quedan completos:
[catálogo Solutions](./catalogo-solutions.md) (1),
[catálogo Systems](./catalogo-systems.md) (2),
[ciclo de vida del activo](./ciclo-de-vida-del-activo.md) (3),
[estados del activo](./estados-del-activo.md) (4),
[20 SOP-ACTIVO](./sop/README.md) (5), y este documento de KPI (6). La
FASE II · Sistema ([gobernanza](./modelo-de-gobernanza.md),
[datos](./modelo-de-datos.md), [reglas](./reglas-operativas.md),
[arquitectura](./arquitectura-os.md)) ya estaba resuelta antes de cerrar
este paso — quedan ambas fases completas.

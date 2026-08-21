# Modelo de datos

**Fase:** II · Sistema — paso 8 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

Este documento traduce las [5 capas](./README.md#etapa-3--modelo-de-packaging-systems),
los [16 estados](./estados-del-activo.md) y las capacidades de
[Systems](./catalogo-systems.md) a un esquema de datos, y lo **coteja contra
el esquema real** que ya existe en `chacontainer/migrations/001_initial_schema.sql`
y `002_dashboard_views.sql`. No es un modelo de datos desde cero: es el
modelo objetivo de este proyecto maestro superpuesto al SaaS que ya opera,
con las brechas explícitas.

> **Aclaración de alcance:** este documento describe y compara — no modifica
> el código ni las migraciones de `chacontainer/`. Cualquier cambio de
> esquema real es decisión de quien mantiene ese proyecto.

---

## 1. Lo que ya existe (esquema real, resumido)

| Tabla | Cubre | Capa(s) que alimenta |
|---|---|---|
| `tenants`, `users` | Multi-tenant y usuarios/roles | — (infraestructura) |
| `plants`, `plant_zones` | Ubicaciones físicas propias | Capa 2 · Ubicación (parcial — no cubre ubicaciones de terceros) |
| `clients`, `client_contacts`, `client_interactions` | CRM de clientes | Capa 3 · Custodia (parcial — un cliente, no un historial de custodios) |
| `assets` | Registro maestro del activo, con `status` (6 valores), `qr_code`, `client_id` | Capa 1 · Activo + Capa 4 · Estado (parcial) |
| `asset_events` | Historial de cambios de estado y ubicación por activo | Capa 4 (historial) + Capa 2 (historial) — es la base real de Trazabilidad |
| `qr_scans` | Log de cada escaneo QR | Identificación / Trazabilidad (evento crudo) |
| `shipments`, `shipment_lines`, `shipment_events` | Envíos y su seguimiento | Capa 2 (movimiento) — cubre "Movimiento" más que "asset_events" para el caso de envío consolidado |
| `maintenance_records` | Registro de mantenimiento (tipo, fechas, técnico, costo) | Evidencia de Lavado/Reparación (Solutions) |
| `airtable_sync_log` | Integración externa | — (infraestructura) |

Este esquema real cubre razonablemente bien Capa 1, Capa 2 y una versión
simplificada de Capa 4. **No cubre**, o cubre de forma muy parcial: Capa 3
(custodia con historial), Capa 5 (gobierno/reglas), ni las capacidades de
Alertas, Incidencias y Analítica de [catalogo-systems.md](./catalogo-systems.md).

## 2. Brecha principal: `assets.status` (6 valores) vs. estados-del-activo.md (16 estados)

```sql
status VARCHAR(20) NOT NULL DEFAULT 'available'
  CHECK (status IN ('available','in_use','in_transit','maintenance','retired','lost'))
```

Esto es una versión **colapsada** de los 16 estados del [catálogo de
estados](./estados-del-activo.md#1-catálogo-de-estados). Mapeo propuesto
(no destructivo — el valor real se puede seguir usando como "estado
agregado" mientras el detalle vive en una tabla nueva, ver §3):

| Estado real (`assets.status`) | Estados objetivo que agrupa |
|---|---|
| `available` | `Disponible` |
| `in_use` | `Asignado`, `En uso` |
| `in_transit` | `En tránsito` |
| `maintenance` | `Registrado`, `Identificado`, `En inspección`, `Sucio`, `En reparación`, `Liberado` |
| `retired` | `Baja`, `Scrap`, `Dispuesto` |
| `lost` | `Retenido`, `Bloqueado`, `Perdido` |

`maintenance` y `lost` son los que más comprimen información: el valor
`maintenance` no distingue si el activo está sucio, dañado, en inspección o
ya liberado pero sin mover a inventario — justo la granularidad que
[SOP-ACTIVO-03 a 08](./sop/README.md) necesitan para operar. Recomendación:
no reemplazar `assets.status` (rompería lo que ya consume ese campo, como
`mv_plant_inventory`), sino **agregar** un campo de estado detallado (§3) y
mantener `assets.status` como proyección agregada de ese detalle.

## 3. Esquema objetivo (nuevas tablas propuestas)

### `asset_state_detail` (extiende Capa 4)

Guarda el estado fino de los 16 de `estados-del-activo.md`, con
`assets.status` derivado de él (por trigger o por la aplicación).

| Columna | Tipo | Nota |
|---|---|---|
| `asset_id` | UUID (FK `assets.id`) | PK — un registro vigente por activo |
| `estado_detalle` | VARCHAR(30) | uno de los 16 estados; `CHECK` con la lista cerrada |
| `desde` | TIMESTAMPTZ | inicio del estado actual |
| `sop_origen` | VARCHAR(20) | qué SOP-ACTIVO disparó la transición (ver [estados-del-activo.md §2](./estados-del-activo.md#2-tabla-de-transiciones)) |

El historial de transiciones ya tiene hogar natural: **reutilizar
`asset_events`**, agregando `estado_detalle_from` / `estado_detalle_to` junto
a los `from_status`/`to_status` agregados que ya existen, en vez de crear una
tabla de historial paralela.

### `asset_custody` (Capa 3 · Custodia — no existe hoy)

Hoy `assets.client_id` es un solo campo, sin historial ni tipos de custodio
más allá de "cliente". No alcanza para [SOP-ACTIVO-12](./sop/SOP-ACTIVO-12-transferencia-de-custodia.md)
(transferencias sin pasar por CHACONTAINER) ni para responder "¿qué
proveedor retiene activos?" ([ETAPA 4](./README.md#etapa-4--modelo-de-gobernanza)).

| Columna | Tipo | Nota |
|---|---|---|
| `id` | UUID | PK |
| `asset_id` | UUID (FK) | — |
| `custodian_type` | VARCHAR(20) | `cliente`, `proveedor`, `operador`, `transportista`, `planta_propia`, `usuario` |
| `custodian_ref_id` | UUID | FK polimórfica a `clients.id` u otra tabla según `custodian_type` `[VALIDAR: modelo polimórfico vs. tabla `custodians` unificada — decisión de diseño pendiente]` |
| `desde` | TIMESTAMPTZ | — |
| `hasta` | TIMESTAMPTZ (nullable) | NULL = custodia vigente |
| `origen_evento` | VARCHAR(20) | `asignacion` (SOP-ACTIVO-10) o `transferencia` (SOP-ACTIVO-12) |

### `operational_rules` (Capa 5 · Gobierno — no existe hoy)

Ver especificación completa en [reglas-operativas.md](./reglas-operativas.md)
(paso 9). Esqueleto de columnas: `scope` (tenant/cliente/tipo de
activo/planta), `parametro` (p. ej. `tiempo_maximo_fuera_dias`), `valor`,
`accion` (alerta / transición automática), `prioridad`.

### `incidents` (Gestión de incidencias — no existe hoy)

Respalda [SOP-ACTIVO-15](./sop/SOP-ACTIVO-15-incidencias.md),
[16](./sop/SOP-ACTIVO-16-activo-perdido.md) y
[17](./sop/SOP-ACTIVO-17-activo-bloqueado.md).

| Columna | Tipo | Nota |
|---|---|---|
| `id` | UUID | PK |
| `asset_id` | UUID (FK, nullable) | nullable porque una incidencia puede ser de un lote, no de un activo (ver SOP-ACTIVO-04) |
| `tipo` | VARCHAR(20) | `daño`, `retraso`, `faltante`, `perdida`, `disputa` |
| `estado` | VARCHAR(20) | `abierta`, `en_proceso`, `resuelta` |
| `sop_relacionado` | VARCHAR(20) | — |
| `abierta_por` / `resuelta_por` | UUID (FK `users.id`) | — |
| `abierta_at` / `resuelta_at` | TIMESTAMPTZ | — |

### `alerts` (Alertas — no existe hoy)

| Columna | Tipo | Nota |
|---|---|---|
| `id` | UUID | PK |
| `asset_id` | UUID (FK) | — |
| `rule_id` | UUID (FK `operational_rules.id`) | qué regla la disparó |
| `estado` | VARCHAR(20) | `activa`, `atendida`, `descartada` |
| `notificado_a` | UUID (FK `users.id`) | — |
| `creada_at` / `atendida_at` | TIMESTAMPTZ | — |

## 4. Identificación: QR ya cubierto, RFID no

`assets.qr_code` y `qr_scans` cubren bien [catalogo-systems.md #4 (QR)](./catalogo-systems.md#4-qr).
No hay ningún campo ni tabla para [RFID (#5)](./catalogo-systems.md#5-rfid).
Propuesta mínima: agregar `assets.rfid_tag VARCHAR(100)` y una tabla
`rfid_reads` paralela a `qr_scans` (mismas columnas relevantes: `asset_id`,
`plant_id`, `zone_id`, `read_at`), o generalizar `qr_scans` a
`identifier_reads` con un campo `method` (`qr`/`rfid`) — **decisión de
diseño para quien mantiene el SaaS, no de este proyecto** `[VALIDAR con el
equipo técnico]`.

## 5. Resumen: qué falta construir

| Capa / capacidad | Cobertura actual | Falta |
|---|---|---|
| Capa 1 · Activo | Alta (`assets`) | Soporte de variantes/configuración (Solutions #6 Modificación) |
| Capa 2 · Ubicación | Alta (`plants`, `plant_zones`, `asset_events`) | Ubicaciones de terceros (cliente/proveedor) más allá de `client_id` |
| Capa 3 · Custodia | Baja (solo `client_id`) | `asset_custody` con historial y tipos de custodio |
| Capa 4 · Estado | Media (`assets.status`, 6 valores) | `asset_state_detail`, 16 estados, ver §2 |
| Capa 5 · Gobierno | Nula | `operational_rules` (paso 9) |
| Trazabilidad | Alta (`asset_events`, `qr_scans`) | Extender con estado detallado y RFID |
| Gestión de incidencias | Nula | `incidents` |
| Alertas | Nula | `alerts` |
| Analítica | Parcial (vistas materializadas de inventario/envíos) | Vistas sobre `incidents`, `alerts`, `asset_state_detail` para los KPI de la [ETAPA 6](./README.md#etapa-6--indicadores-del-sistema) |

---

## Próximo paso sugerido

Paso 9 de la FASE II: **reglas** — especificar `operational_rules` en
detalle (tipos de parámetro, alcance, prioridad, qué acción dispara cada
una), que es el insumo que le falta a `estados-del-activo.md` para que las
transiciones automáticas (invariante 6) dejen de ser un supuesto y sean una
tabla real.

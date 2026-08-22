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

**Decidido: tabla `custodians` unificada, no FK polimórfica.** Una FK que
apunte a `clients.id` o a otra tabla según `custodian_type` no se puede
declarar como `REFERENCES` real en Postgres — obliga a triggers o a
renunciar a la integridad referencial, justo el tipo de brecha silenciosa
que este proyecto lleva insistiendo en cerrar (ver el mismo argumento
aplicado a estados en [estados-del-activo.md, invariante 5](./estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)).
Una tabla `custodians` unificada —con `client_id UUID REFERENCES clients(id)`
nullable para el caso `cliente`, y campos propios (`nombre`, `tipo`,
`contacto`) para `proveedor`/`operador`/`transportista`/`planta_propia`/`usuario`—
mantiene la integridad referencial real y permite un solo join limpio desde
`asset_custody`, en vez de resolver el tipo en código de aplicación cada
vez que se consulta.

| Columna | Tipo | Nota |
|---|---|---|
| `id` | UUID | PK |
| `asset_id` | UUID (FK) | — |
| `custodian_id` | UUID (FK `custodians.id`) | reemplaza `custodian_type` + `custodian_ref_id`; el tipo vive en `custodians.tipo` |
| `desde` | TIMESTAMPTZ | — |
| `hasta` | TIMESTAMPTZ (nullable) | NULL = custodia vigente |
| `origen_evento` | VARCHAR(20) | `asignacion` (SOP-ACTIVO-10) o `transferencia` (SOP-ACTIVO-12) |

```sql
CREATE TABLE custodians (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tenant_id  UUID        NOT NULL REFERENCES tenants(id) ON DELETE CASCADE,
    tipo       VARCHAR(20) NOT NULL
                 CHECK (tipo IN ('cliente','proveedor','operador','transportista','planta_propia','usuario')),
    client_id  UUID        REFERENCES clients(id),   -- solo si tipo = 'cliente'
    nombre     VARCHAR(200) NOT NULL,                 -- denormalizado: útil aun cuando client_id es NULL
    contacto   JSONB       NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

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
Decisión y plan de migración diferido: [escala-rfid.md §4](./escala-rfid.md#4-impacto-en-el-modelo-de-datos)
(generalizar a `identifier_reads` con campo `method`, ejecutado solo cuando
se dispare la Fase F de RFID — no antes).

## Actor de sistema para eventos automáticos

`asset_events.user_id` y `qr_scans.scanned_by` son `NOT NULL REFERENCES users(id)`
en el esquema real. Las transiciones automáticas del MVP (Fase C del
[roadmap](./arquitectura-os.md#3-roadmap-de-implementación-propuesto):
reglas operativas + alertas, ver [reglas-operativas.md](./reglas-operativas.md))
y las lecturas de portal RFID futuras necesitan escribir estos eventos sin
que haya una persona detrás.

**Decidido: usuario sistema reservado por tenant, no relajar el `NOT NULL`.**
Cada tenant tiene, desde su alta, una fila en `users` de tipo sistema:

```sql
-- Se agrega 'system' al CHECK de users.role, y se inserta una fila
-- reservada por tenant en el mismo flujo que crea el tenant:
INSERT INTO users (tenant_id, email, name, role, active)
VALUES ($tenant_id, 'system@' || $tenant_slug || '.internal', 'Sistema (automático)', 'system', TRUE);
```

Toda transición o lectura automática usa el `id` de esa fila como
`user_id`/`scanned_by`. Ventajas sobre permitir `NULL`: la integridad
referencial se mantiene sin excepción, ningún reporte ni vista downstream
necesita un `COALESCE` o un caso especial para "sin usuario", y el propio
registro de auditoría distingue automático de manual con el mismo campo que
ya usa para todo lo demás — sin agregar una columna booleana aparte.

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

# Arquitectura de CHACONTAINER OS

**Fase:** II · Sistema — paso 10 (último) de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

Cierra la FASE II reconciliando los 12 módulos del MVP definidos en la
[ETAPA 5](./README.md#etapa-5--chacontainer-os) contra la arquitectura
técnica real ya documentada en
[`chacontainer/docs/architecture.md`](../chacontainer/docs/architecture.md),
y contra las tablas nuevas propuestas en
[modelo-de-datos.md](./modelo-de-datos.md) y
[reglas-operativas.md](./reglas-operativas.md). Es la especificación de
"qué construir" que conecta este proyecto maestro con el SaaS que ya existe.

> Como en los documentos anteriores de esta fase: esto describe y propone,
> no modifica el código real de `chacontainer/`.

---

## 1. Stack real (resumen de architecture.md)

- **API**: Go / `net/http`, en [`chacontainer/cmd/server`](../chacontainer/cmd/server) + [`chacontainer/internal/`](../chacontainer/internal/)
- **Base de datos**: PostgreSQL 15 multi-tenant + vistas materializadas
- **Cache/rate limiting**: Redis
- **Frontend**: React 19 + Vite + Tailwind, en [`chacontainer/web/`](../chacontainer/web/)
- **Automatización batch**: scripts Python en [`chacontainer/scripts/`](../chacontainer/scripts/) (sync Airtable, alertas, reportes)
- **Integraciones**: Airtable, Make.com, ERP externo (SAP/Odoo)
- **Auth**: JWT HS256; identidad QR con HMAC-SHA256

Dominios ya existentes en `internal/domain/`: `tenant`, `asset`, `qr`,
`shipment`, `plant`, `client`. Es decir: la base de Capa 1 (Activo), Capa 2
(Ubicación parcial) y CRM ya tiene código propio. **No existen** dominios
`custody`, `rules`, `incidents`, ni `alerts` — coincide exactamente con el
gap identificado en [modelo-de-datos.md §5](./modelo-de-datos.md#5-resumen-qué-falta-construir).

## 2. Los 12 módulos del MVP, mapeados

| # | Módulo (ETAPA 5) | Dominio/tabla real | Estado | Qué falta |
|---|---|---|---|---|
| 1 | Activos (registro maestro) | `internal/domain/asset` + `assets` | ✅ Existe | Ninguna brecha estructural |
| 2 | Identificación (QR/RFID) | `internal/domain/qr` + `qr_code`, `qr_scans` | 🟡 Parcial | QR completo; RFID no existe (ver [modelo-de-datos.md §4](./modelo-de-datos.md#4-identificación-qr-ya-cubierto-rfid-no)) |
| 3 | Ubicaciones | `internal/domain/plant` + `plants`, `plant_zones` | 🟡 Parcial | Cubre plantas propias; no modela ubicaciones de terceros más allá de `client_id` |
| 4 | Movimientos | `internal/domain/shipment` + `shipments`, `shipment_events`, `asset_events` | ✅ Existe | Ninguna brecha estructural |
| 5 | Custodia | — | 🔴 No existe | Nuevo dominio `custody` + tabla `asset_custody` ([modelo-de-datos.md](./modelo-de-datos.md#asset_custody-capa-3--custodia--no-existe-hoy)) |
| 6 | Condición | `assets.status` (colapsado) | 🟡 Parcial | `asset_state_detail` con los 16 estados ([modelo-de-datos.md §2-3](./modelo-de-datos.md#2-brecha-principal-assetsstatus-6-valores-vs-estados-del-activomd-16-estados)) |
| 7 | Mantenimiento | `maintenance_records` | 🟡 Parcial | Cubre evidencia de lavado/reparación; no está vinculado a `asset_state_detail` ni a los SOP-ACTIVO-05/06 explícitamente |
| 8 | Incidencias | — | 🔴 No existe | Nuevo dominio `incidents` + tabla `incidents` |
| 9 | Inventario (disponible vs. requerido) | `mv_plant_inventory` | 🟡 Parcial | La vista existe para conteo por estado agregado; falta "requerido" (nivel mínimo, ver [reglas-operativas.md](./reglas-operativas.md)) |
| 10 | Alertas | — | 🔴 No existe | Nuevo dominio `alerts` + tabla `alerts`, depende de `operational_rules` |
| 11 | Dashboard (indicadores) | `mv_plant_inventory`, `mv_shipment_kpis`, `mv_daily_scans` | 🟡 Parcial | Cubre inventario/envíos/escaneos; faltan los KPI de la [ETAPA 6](./README.md#etapa-6--indicadores-del-sistema) que dependen de `incidents`/`alerts`/`asset_custody` (pérdida, daño, cumplimiento de retorno) |
| 12 | Evidencia | `maintenance_records.notes`, `asset_events.notes`, `geo` (JSONB) | 🟡 Parcial | No hay almacenamiento de archivos/fotos explícito — solo campos de texto/JSON `[VALIDAR: ¿dónde se guardan las fotos hoy — Airtable, S3, ninguna?]` |

**Resumen**: 2 de 12 módulos completos, 8 parciales, 2 inexistentes
(Custodia, Alertas — Incidencias también inexistente aunque cuenta aparte
en la tabla por claridad). El SaaS real ya resuelve el "tracking" (dónde
está, qué es) pero no la "gobernanza" (quién responde, qué regla se
incumplió, qué se hizo al respecto) — exactamente la distinción que hace la
[ETAPA 4](./README.md#etapa-4--modelo-de-gobernanza) entre tracking y
gobernanza.

## 3. Roadmap de implementación propuesto

Orden por dependencia, no por facilidad — cada fase habilita la siguiente:

**Fase A — Estado detallado** (habilita SOP-ACTIVO-03 a 08 en el sistema real)
1. Tabla `asset_state_detail` + extender `asset_events` con estado detallado.
2. Endpoint/lógica que derive `assets.status` (agregado) desde `asset_state_detail` (detallado), sin romper `mv_plant_inventory`.

**Fase B — Custodia** (habilita SOP-ACTIVO-10, 12, 13)
3. Dominio `custody` + tabla `asset_custody`.
4. Migrar `assets.client_id` a un primer registro de `asset_custody` (custodia inicial), sin eliminar el campo mientras haya código que lo consuma.

**Fase C — Reglas y alertas** (habilita el Nivel 1 de [modelo-de-gobernanza.md](./modelo-de-gobernanza.md#1-los-tres-niveles-de-gobernanza))
5. Tabla `operational_rules` con al menos las reglas `global` de los 4 tipos (§1 de [reglas-operativas.md](./reglas-operativas.md)).
6. Tabla `alerts` + job periódico (candidato natural: un script Python nuevo en `chacontainer/scripts/`, siguiendo el patrón ya usado para sync/alertas) que evalúe reglas contra `asset_state_detail` y `asset_custody`.

**Fase D — Incidencias** (habilita SOP-ACTIVO-15, 16, 17)
7. Tabla `incidents`, vinculable a `asset_id` (nullable, ver [modelo-de-datos.md](./modelo-de-datos.md#incidents-gestión-de-incidencias--no-existe-hoy)) o a un lote (SOP-ACTIVO-04).

**Fase E — Indicadores completos** (cierra la [ETAPA 6](./README.md#etapa-6--indicadores-del-sistema))
8. Vistas materializadas nuevas sobre `incidents`, `alerts`, `asset_custody`, `asset_state_detail` para los KPI que hoy no se pueden calcular (pérdida, daño, cumplimiento de retorno, tiempo detenido, costo por ciclo).

**Fase F — RFID** (opcional, por caso de uso — ver [catalogo-systems.md #5](./catalogo-systems.md#5-rfid))
9. Solo si un piloto concreto ([ETAPA 7](./README.md#etapa-7--piloto-packaging-systems)) lo requiere; no es prerequisito de las fases A-E.

## 4. Qué no cambia

Este roadmap es aditivo: ninguna fase requiere romper el esquema o los
endpoints que ya existen (`assets`, `shipments`, `qr_scans`, las vistas
materializadas). El diseño de `asset_state_detail` y `asset_custody` como
tablas nuevas con `assets.status`/`assets.client_id` como proyecciones
agregadas es deliberado: permite construir Systems encima de Solutions sin
una migración disruptiva del SaaS que ya opera con clientes reales.

---

## Cierre de FASE II

Con este documento se completan los 4 pasos de la FASE II · Sistema:
[modelo de gobernanza](./modelo-de-gobernanza.md) (paso 7),
[modelo de datos](./modelo-de-datos.md) (paso 8),
[reglas operativas](./reglas-operativas.md) (paso 9) y esta arquitectura
(paso 10). Sigue la **FASE III · Producto** ([ETAPA 14](./README.md#etapa-14--orden-de-ejecución)):
MVP, QR, Dashboard, Alertas, Piloto — que ahora tiene, en el roadmap del
§3, una traducción directa a trabajo de ingeniería sobre `chacontainer/`.

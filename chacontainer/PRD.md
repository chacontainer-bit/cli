# PRD — CHACONTAINER Núcleo Operativo

Estado: v1 — auto-generado a partir de `docs/architecture.md` y del código
existente, y completado por Claude en la rama `claude/build-from-single-prd-d8ngy0`.

## Contexto

CHACONTAINER ya tenía dos mitades inconexas cuando arrancó este PRD:

- **Backend (Go)**: dominios, handlers HTTP y middleware JWT reales, pero con
  persistencia 100% en memoria (`internal/api/stubs.go` — nada se guardaba).
- **Frontend (React/Vite)**: pantallas visualmente terminadas pero conectadas
  a `src/data/mockData.js` — sin login real, sin llamadas a la API, con un
  vocabulario de producto distinto al del backend (contenedores de "lavado"
  en Chile vs. activos logísticos multi-planta en México).

Este PRD fija el backend (`docs/architecture.md`) como fuente de verdad del
modelo de datos, y reconstruye el frontend del núcleo operativo sobre la API
real. Marketplace, Proyectos, Calculadora, Asistente y AssetFactory son
herramientas satélite ya funcionales sobre datos simulados: quedan
**fuera de alcance** de este PRD.

## objetivo

Que un usuario pueda registrar su empresa (tenant), iniciar sesión, y
gestionar de punta a punta el ciclo de vida de sus activos retornables
(contenedores, pallets, IBCs, etc.) en un sistema multi-tenant real:
alta de plantas, alta de activos con QR, escaneo de eventos, envíos
(shipments) con línea de tiempo de estado, y una cartera de clientes (CRM)
básica — todo respaldado por persistencia real, no por datos de ejemplo.

## pantallas

1. **Login / Registro** (`/login`)
   - Iniciar sesión con email + contraseña.
   - Crear cuenta: nombre de empresa + nombre de usuario + email + contraseña
     → crea un tenant nuevo y un usuario `admin` para ese tenant.
2. **Dashboard** (`/dashboard`, protegida)
   - KPIs reales: activos totales, en tránsito, envíos de hoy, envíos
     atrasados, escaneos últimas 24h.
   - Distribución de activos por estado (pie) y tendencia de activos/envíos
     (barras), leídos de `/api/v1/dashboard/*`.
3. **Activos / Trazabilidad** (`/traceability`, protegida)
   - Listar/buscar/filtrar activos por planta, tipo, estado, texto.
   - Alta de activo (genera código QR firmado).
   - Simular escaneo de QR (check-in/check-out/transferencia/inspección) y
     ver la línea de tiempo de eventos por activo.
4. **Backoffice — Envíos y Clientes** (`/backoffice`, protegida)
   - Envíos: listar, crear, transicionar de estado (`draft → confirmed →
     in_transit → delivered`, o `cancelled`/`exception`), ver línea de tiempo.
   - Clientes: listar, crear, alta rápida de contacto.
5. **Plantas** (`/plants`, protegida)
   - Listar/crear plantas y sus zonas (bodega, andén, patio, producción,
     cuarentena).

Rutas protegidas redirigen a `/login` si no hay sesión válida.

## reglas

- **Aislamiento por tenant**: toda fila lleva `tenant_id`; toda consulta se
  filtra por el tenant del JWT. Un usuario nunca ve datos de otro tenant.
- **Autenticación**: JWT HS256, 24h de vigencia, emitido sólo por
  `POST /api/v1/auth/login` tras verificar `bcrypt`. Todas las rutas
  `/api/v1/*` exigen `Authorization: Bearer <token>`.
- **Registro**: el primer usuario de un tenant nuevo siempre es `admin`.
  Contraseña mínima 8 caracteres. Email único por tenant.
- **Roles**: `admin`, `manager`, `operator`, `viewer`, `api` (ya definidos en
  `domain/tenant`). `viewer` no debería escribir — queda documentado como
  regla de negocio; la aplicación estricta por endpoint es backlog (ver
  `docs/architecture.md` `RequireRole`, hoy no conectado a ninguna ruta).
- **Activos**: `code` y `qr_code` únicos por tenant. Todo alta registra un
  evento `check_in` en el historial del activo (auditoría inmutable).
- **Escaneo QR**: requiere que el `qr_code` exista para el tenant; cada
  escaneo agrega un evento y actualiza `last_scan_at`/planta/zona del activo.
- **Envíos**: nacen en estado `draft`; cada transición de estado queda
  registrada como evento con usuario y fecha (no se sobreescribe historial).
- **Persistencia real**: SQLite embebido (motor puro Go, sin cgo) para poder
  correr y probar el stack sin infraestructura externa; el esquema es
  compatible en forma con `migrations/001_initial_schema.sql` (Postgres
  sigue siendo el objetivo de producción, documentado en
  `docs/architecture.md`).

## pruebas

- **Backend (Go, `go test ./...` dentro de `chacontainer/`)**:
  - Registro y login: credenciales válidas emiten JWT; credenciales
    inválidas y contraseña corta se rechazan.
  - CRUD de activos, plantas, clientes y envíos contra SQLite en memoria.
  - Aislamiento por tenant: un tenant no puede leer/listar recursos de otro.
  - Transición de estado de envío queda reflejada en `GET .../timeline`.
  - Escaneo de QR actualiza el activo y agrega evento.
- **Frontend (Vitest + Testing Library)**:
  - Flujo de login: credenciales incorrectas muestran error; correctas
    navegan a `/dashboard`.
  - Dashboard renderiza KPIs a partir de una respuesta de API simulada
    (`fetch` mockeado).
  - Ruta protegida redirige a `/login` sin sesión.
- **Verificación manual**: levantar backend + frontend localmente
  (ver `chacontainer/README.md`), registrar un tenant nuevo desde la UI,
  crear una planta, un activo, escanearlo, crear un envío y transicionarlo,
  confirmando que los datos persisten tras reiniciar el backend.

## Fuera de alcance (backlog explícito)

- Persistencia Postgres real en producción, Redis, workers cron.
- Webhooks salientes reales hacia Make.com/ERP/Airtable (los clientes HTTP
  ya existen en `internal/integrations/*` pero no están conectados al flujo
  de escritura; hoy sólo se registra el evento entrante).
- Aplicación de `RequireRole` por endpoint.
- Marketplace, Proyectos, Calculadora, Asistente, AssetFactory (ya
  funcionales sobre datos simulados, no tocados por este PRD).

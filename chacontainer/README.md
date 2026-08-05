# CHACONTAINER

Plataforma multi-tenant para gestión de activos retornables (contenedores,
pallets, IBCs, racks…) en la industria automotriz: escaneo QR, envíos con
línea de tiempo de estado, y una cartera de clientes básica.

Ver `PRD.md` para objetivo/pantallas/reglas/pruebas del alcance actual y
`docs/architecture.md` para el diseño técnico completo (incluyendo el
objetivo de producción con Postgres/Redis).

## Arquitectura de este repo

- `cmd/server`, `internal/` — API REST en Go (módulo propio,
  `github.com/chacontainer/backend`, independiente del `go.mod` raíz del
  CLI de `gh`).
- `web/` — frontend React + Vite + Tailwind.
- `migrations/` — esquema de referencia para Postgres (producción).
  El backend usa hoy SQLite embebido (`internal/store/sqlite`) para poder
  correr sin infraestructura externa; ver PRD.md, sección "Fuera de alcance".
- `scripts/` — automatizaciones Python (QR, sync Airtable, alertas, reportes).

## Levantar el backend

```bash
cd chacontainer
go run ./cmd/server
```

Variables de entorno (todas opcionales en desarrollo):

| Variable | Default | Notas |
|---|---|---|
| `PORT` | `8080` | |
| `DB_PATH` | `./chacontainer.db` | archivo SQLite, se crea/migra solo |
| `JWT_SECRET` | secreto de desarrollo inseguro | **requerido** si `ENVIRONMENT != development` |
| `ENVIRONMENT` | `development` | |
| `MAKE_WEBHOOK_SECRET`, `AIRTABLE_*`, `ERP_*` | vacío | integraciones opcionales, ver architecture.md |

En el primer arranque, si la base está vacía, se siembra un tenant demo con
datos de ejemplo (plantas, activos, cliente, envío) y se imprime en el log:

```
seeded demo tenant — login with demo@chacontainer.mx / Demo12345!
```

## Levantar el frontend

```bash
cd chacontainer/web
npm install
VITE_API_URL=http://localhost:8080 npm run dev
```

Abre `http://localhost:5173`. Inicia sesión con las credenciales demo de
arriba, o crea una cuenta nueva desde la pantalla de login (pestaña "Crear
cuenta") — cada registro crea un tenant nuevo y aislado.

## Pruebas

```bash
# Backend (Go) — httptest + SQLite en memoria
cd chacontainer && go test ./...

# Frontend (Vitest + Testing Library)
cd chacontainer/web && npm test
```

## Alcance actual vs. backlog

Ver la sección "Fuera de alcance" de `PRD.md`. En resumen: este PRD cubre el
núcleo operativo (auth, dashboard, activos/QR, envíos, clientes, plantas)
con persistencia SQLite real y probada. Postgres en producción, Redis,
workers cron, y el envío saliente de webhooks a Make.com/ERP/Airtable siguen
pendientes. Marketplace, Proyectos, Calculadora, Asistente y Fábrica de
Assets son herramientas ya existentes sobre datos simulados, no tocadas por
este PRD.

# CLAUDE.md

Este repo es un fork de GitHub CLI (`gh`) que además aloja **chacontainer/**,
mi SaaS real de gestión de activos industriales. Son dos proyectos distintos
en el mismo repo, con historiales, stacks y convenciones separadas:

- Todo fuera de `chacontainer/` → es el fork de `gh`. Sigue **AGENTS.md**
  (Go, arquitectura Options+Factory, `make lint`, tests con `testify`/`moq`).
  No mezcles cambios de `chacontainer/` en el mismo commit que cambios del
  CLI, y viceversa.
- `chacontainer/` → mi producto. Ver detalle abajo.

## chacontainer/ — estructura y stack

```
chacontainer/
├── cmd/server/main.go        # entrypoint del binario Go (net/http, sin framework)
├── internal/
│   ├── api/                  # router.go + handlers/ + middleware/
│   ├── config/                # carga de env vars
│   ├── domain/                # asset, client, plant, qr, shipment, tenant
│   └── integrations/           # airtable/, erp/, make_com/
├── web/                       # React 19 + Vite + Tailwind (SPA)
│   └── src/{components,pages,hooks,context,lib,data,assets}/
├── scripts/                   # jobs batch en Python (cron)
├── migrations/                # SQL Postgres, numeradas (001_, 002_...)
└── docs/
    ├── architecture.md        # referencia técnica completa (rutas API, schema, deploy)
    ├── sop/                   # SOP-01..SOP-13, ciclo comercial completo
    ├── lista-precios-costos.md
    └── fundamentos-agentes-ia.md / agente-ingeniero-procesos-ia.md
```

Antes de asumir cómo funciona algo en chacontainer/, lee
`chacontainer/docs/architecture.md` — tiene el diagrama de sistema, las
rutas REST completas, el modelo multi-tenant, el formato del QR, las
integraciones (Airtable/Make.com/ERP) y las env vars de producción.

### Estado real del backend (importante)

`chacontainer/internal/api/router.go` conecta los handlers a **stores stub
en memoria** (`stubAssetStore`, `stubShipmentStore`, etc. en `stubs.go`), no
a Postgres todavía, aunque `chacontainer/migrations/` ya define el schema
real (multi-tenant por `tenant_id`, incluyendo vistas materializadas). No
asumas que los datos persisten entre requests ni que las migraciones están
aplicadas — si una tarea depende de eso, dilo explícitamente en vez de
fingir que ya existe.

### Stack por área
| Área | Stack | Comando de calidad |
|---|---|---|
| API | Go, `net/http` puro (sin framework), JWT HS256 | `go build ./chacontainer/...`, `go vet ./chacontainer/...` |
| Web | React 19 + Vite + Tailwind, `react-router-dom`, `recharts` | `cd chacontainer/web && npm run lint` (no hay `npm test` — no asumas cobertura) |
| Jobs batch | Python (`scripts/*.py`): sync Airtable, alertas de envíos, reportes, generación de QR | sin lint/test formal definido — revisa manualmente |
| DB | PostgreSQL, multi-tenant por `tenant_id`, migraciones SQL numeradas | — |

### Convenciones de commits en chacontainer/
El historial usa Conventional Commits con scope, mensaje en español:
`docs(chacontainer): ...`, `feat(chacontainer): ...`, `fix(lint): ...`,
`chore(deps): ...`. Sigue este patrón para cambios en `chacontainer/`. El
resto del repo (fork de `gh`) usa mensajes en inglés siguiendo el estilo
del proyecto upstream.

### SOPs (`chacontainer/docs/sop/`)
Documentan el ciclo comercial completo (13 pasos). Formato fijo obligatorio
para cualquier SOP nuevo: **Resumen ejecutivo, Objetivo, Alcance, Roles,
Diagrama de flujo, SOP paso a paso, Checklist, KPI, Riesgos y controles.**
Sigue este formato SIEMPRE al documentar un proceso nuevo — no lo
abrevies ni reordenes.

## Antes de terminar cualquier cambio en chacontainer/
- Go: `go build ./chacontainer/...` y `go vet ./chacontainer/...`
- Web: `cd chacontainer/web && npm run lint`
- No des una tarea por completa si el build/lint no corrió.

## Antes de terminar cualquier cambio fuera de chacontainer/ (fork de `gh`)
- `go test ./...` y `make lint` (ver AGENTS.md para el detalle de arquitectura,
  patrones de comandos, mocking HTTP y estilo de código).

## Otros directorios relevantes
- `skills/gh/` y `skills/gh-skill/` — skill del propio CLI `gh` (no confundir
  con `.claude/skills/`).
- `.claude/skills/` — biblioteca grande de skills de marketing (ads, SEO,
  copywriting, pricing, etc.) instalada en el repo; no tiene relación con
  `chacontainer/` como producto ni con el fork de `gh` — son herramientas de
  referencia para tareas de marketing, no código a mantener.

## Límites de este entorno (Claude Code Remote)
- NO tengo acceso a tu computadora/disco local, solo a este repo clonado.
- Los conectores MCP que requieren login (Slack, Notion, Google Drive, etc.)
  no se pueden autorizar aquí — necesitan una sesión interactiva
  (`claude mcp` o `/mcp` en la CLI local). Si me pides "conecta X" o "lee un
  archivo de mi compu" en esta sesión remota, dime explícitamente que no se
  puede en vez de improvisar otra entrega.

## Tareas repetitivas (N documentos/páginas con la misma plantilla)
- Usa varios sub-agentes en paralelo (Agent tool), no una sesión secuencial
  de N commits.

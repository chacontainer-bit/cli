# CLAUDE.md

Este repo es un fork de GitHub CLI (`gh`) que además aloja **chacontainer/**,
mi SaaS real de gestión de activos industriales. Son dos proyectos distintos
en el mismo repo:

- Todo fuera de `chacontainer/` → sigue AGENTS.md (gh CLI, Go, `make lint`).
- `chacontainer/` → mi producto. Stack:
  - `chacontainer/cmd/server` + `chacontainer/internal/` — API Go (net/http)
  - `chacontainer/web/` — frontend React 19 + Vite + Tailwind (`npm run lint`, no hay `npm test` todavía — no asumas cobertura)
  - `chacontainer/scripts/*.py` — jobs batch (Airtable sync, alertas, reportes)
  - `chacontainer/migrations/*.sql` — Postgres, multi-tenant por `tenant_id`
  - `chacontainer/docs/sop/` — SOPs del ciclo comercial (SOP-01 a SOP-13).
    Al documentar un proceso usa SIEMPRE la skill `/sop-chacontainer`
    (plantilla completa: Marco de decisión + 12 secciones).

## Antes de terminar cualquier cambio en chacontainer/
- Corre la skill `/verify` (build/vet de Go + lint web). No des una tarea por
  completa si el build/lint no corrió.

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

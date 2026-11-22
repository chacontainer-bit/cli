---
name: verify
description: Corre los checks obligatorios antes de dar por terminado o commitear un cambio. Usar antes de cada commit, o cuando el usuario pida verificar, compilar o lintear.
---

## Archivos cambiados

!`git status --short`

## Instrucciones

Decide qué checks aplican según las rutas cambiadas arriba (y los commits de la rama si no hay cambios sin commitear):

- Cambios en `chacontainer/` (Go: `cmd/`, `internal/`):
  `go build ./chacontainer/... && go vet ./chacontainer/...`
- Cambios en `chacontainer/web/`:
  `cd chacontainer/web && npm run lint`
- Cambios en `chacontainer/scripts/*.py`:
  `python3 -m py_compile <archivo>` por cada script tocado
- Cambios fuera de `chacontainer/` (gh CLI): sigue AGENTS.md (`make lint`, `go test ./...` del paquete tocado)

Reglas:
- Corre cada check que aplique; no los omitas aunque el cambio parezca trivial.
- Reporta cada check como PASA / FALLA con la salida relevante.
- Separa los errores que tu cambio introdujo de los que ya existían (compara con `git stash` o con `origin` si hace falta). No ocultes ni "arregles" errores ajenos sin avisar.
- No des la tarea por completa si un check que aplica no corrió o falló por tu cambio.

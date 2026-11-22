---
name: verify
description: Verifica cambios antes de commit o de dar una tarea por terminada. Build y vet de Go para chacontainer/, lint del frontend chacontainer/web, y golangci-lint para el resto del repo (gh CLI).
---

# Verificación previa a commit

Revisa qué rutas cambiaron (`git status --short` y `git diff --name-only HEAD`) y corre solo lo que aplica. Ejecuta desde la raíz del repo.

## Cambios en chacontainer/ (Go: cmd/, internal/, tools/)
```bash
go build ./chacontainer/... && go vet ./chacontainer/...
```

## Cambios en chacontainer/web/
```bash
cd chacontainer/web && npm run lint
```
No hay `npm test`. No afirmes que hay cobertura de pruebas.

## Cambios en chacontainer/scripts/*.py
```bash
python3 -m py_compile chacontainer/scripts/<archivo>.py
```

## Cambios en chacontainer/migrations/*.sql
No hay validador automático. Revisa a mano que toda tabla nueva lleve `tenant_id` y que la migración sea aditiva o tenga reversa.

## Cambios fuera de chacontainer/ (gh CLI)
Sigue AGENTS.md: `make lint` (golangci-lint) y `go test` del paquete que tocaste.

## Cómo reportar
- Reporta comando por comando: pasó / falló, con la salida del error.
- Separa los errores que ya existían de los nuevos: corre `git stash` y repite el mismo comando para sacar la línea base, luego `git stash pop`. Solo bloquea el commit un error nuevo.
- Línea base conocida (2026-11-22): `npm run lint` ya marca 4 errores en AppContext.jsx, Dashboard.jsx (Legend sin usar) y Marketplace.jsx (ChevronDown sin usar, función impura en render).
- Si un comando no se pudo correr, dilo. No des la tarea por completa.
- Cambios que solo tocan docs (`*.md`): no corras nada.

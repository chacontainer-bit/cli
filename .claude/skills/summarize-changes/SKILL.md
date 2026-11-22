---
description: Resume los cambios sin commitear y señala riesgos. Usar cuando el usuario pregunte qué cambió, pida un mensaje de commit o quiera revisar su diff.
---

## Cambios actuales

!`git diff HEAD`

## Instrucciones

Resume los cambios en dos o tres viñetas y luego lista los riesgos que veas: manejo de errores faltante, valores hardcodeados, consultas sin filtro por `tenant_id`, migraciones irreversibles o pruebas que haya que actualizar. Si el diff está vacío, di que no hay cambios sin commitear.

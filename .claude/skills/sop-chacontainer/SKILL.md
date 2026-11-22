---
name: sop-chacontainer
description: Plantilla y reglas para documentar un SOP o proceso de CHACONTAINER en chacontainer/docs/sop/. Usar al crear, revisar o actualizar un SOP del ciclo comercial u operativo (prospección, diagnóstico, cotización, lavado, implementación, postventa).
argument-hint: "[nombre del proceso]"
paths: chacontainer/docs/sop/**
---

# SOP de CHACONTAINER

Proceso a documentar: $ARGUMENTS

## Antes de escribir
1. Lee `chacontainer/docs/agente-ingeniero-procesos-ia.md` (metodología) y un SOP existente como referencia de tono y profundidad (por ejemplo `SOP-13-primer-contacto.md`).
2. Ubica el proceso en el ciclo comercial: indica qué SOP lo alimenta y a cuál entrega. Revisa `ls chacontainer/docs/sop/` para no duplicar uno que ya existe.
3. Nombre del archivo: `SOP-NN-kebab-case.md`, con NN = siguiente número libre.

## Reglas de contenido
- Si el SOP no salió de entrevistar a quien hoy ejecuta el proceso, márcalo como **Borrador v0.1 — pendiente de validación** y dilo en el encabezado.
- Marca con `[VALIDAR]` todo responsable, plazo, herramienta, cifra o umbral que no esté confirmado. No inventes precios, capacidades, certificaciones ni clientes.
- KPI: escribe la fórmula y quién da el dato, no solo el nombre del indicador.
- En Riesgos incluye siempre los de CHACONTAINER: retención de activos en custodia, penalizaciones, sobrepromesa, capital inmovilizado y cobro tardío, cuando apliquen.
- En automatización, di explícitamente qué NO se automatiza todavía y por qué.

## Estructura obligatoria (en este orden, con estos títulos)

```markdown
# SOP-NN: <Nombre>

**Estado: Borrador v0.1 — pendiente de validación.**
<Párrafo: paso del ciclo, SOP de entrada y de salida, origen de la información, qué falta para declararlo estándar.>

---

## Marco de decisión aplicado
1. **¿Cuál es el objetivo del proceso?**
2. **¿Qué problema busca resolver?**
3. **¿Quién es el cliente?** (interno y externo)
4. **¿Cuáles son las entradas?**
5. **¿Cuáles son las salidas esperadas?**
6. **¿Quién es el responsable?**
7. **¿Qué indicadores dicen que el proceso funciona?**
8. **¿Qué riesgos existen?**
9. **¿Qué actividades no agregan valor?**
10. **¿Qué partes pueden estandarizarse?**
11. **¿Qué partes pueden automatizarse?**
12. **¿Qué depende únicamente del fundador?**

---

## 1. Resumen ejecutivo
## 2. Objetivo
## 3. Alcance            (desde / hasta / qué NO incluye)
## 4. Roles              (tabla Rol | Responsabilidad)
## 5. Diagrama de flujo  (bloque de texto con flechas y decisiones ¿sí/no?)
## 6. SOP paso a paso    (pasos numerados: quién, qué, entrada, salida)
## 7. Checklist          (- [ ] verificables)
## 8. KPI                (tabla KPI | Fórmula | Meta [VALIDAR] | Fuente)
## 9. Riesgos y controles (tabla Riesgo | Control)
## 10. Mejoras
## 11. Recomendaciones de automatización
## 12. Versión, fecha y autor (versión, fecha, autor, próxima acción)
```

## Al terminar
- Confirma que las 12 secciones numeradas y el Marco de decisión están presentes: `grep -c '^## ' <archivo>` debe dar 13.
- Si creas varios SOPs con esta plantilla, usa sub-agentes en paralelo (uno por SOP), según CLAUDE.md.

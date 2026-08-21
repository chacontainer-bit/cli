---
name: plan-hoy
description: Fija las tres prioridades del día y las escribe en la bóveda. Úsala cuando pida "plan de hoy", "qué hago hoy", "cuáles son mis prioridades", "organiza mi día", "empecemos el día" o "por dónde empiezo".
---

# Plan de hoy

1. Comprueba primero si ya existe `raw/AAAA-MM-DD-plan.md` de hoy. Si existe,
   **no lo dupliques**: léemelo en voz alta y pregúntame si hay cambios.
2. Si no existe, busca con Grep en `raw/` de los últimos 3 días y en `outputs/`
   de ayer: tareas sin cerrar, compromisos y cosas que arrastro.
3. Lee `wiki/contexto.md` para saber qué me importa esta semana.
4. Propón **exactamente 3 prioridades**, ordenadas por impacto. Ni una más.
5. Escribe `raw/AAAA-MM-DD-plan.md`:

   ```
   ---
   fecha: AAAA-MM-DD
   tipo: plan
   tags: [plan, diario]
   ---
   # Plan del AAAA-MM-DD

   - [ ] Prioridad 1
   - [ ] Prioridad 2
   - [ ] Prioridad 3

   ## De dónde salen
   Una línea por prioridad diciendo de qué nota viene.
   ```

6. Léemelas en voz alta: una frase por prioridad, sin preámbulo.

Si no hay nada en la bóveda de los últimos días, dilo y pregúntame las tres
prioridades en vez de inventarlas.

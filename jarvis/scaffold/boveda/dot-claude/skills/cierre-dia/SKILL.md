---
name: cierre-dia
description: Cierra el día, registra qué se hizo y deja encolado mañana. Úsala cuando pida "cierra el día", "resumen del día", "qué he hecho hoy", "terminemos por hoy" o "cierre".
---

# Cierre del día

1. Lee `raw/AAAA-MM-DD-plan.md` de hoy. Si no existe, lee las notas de `raw/`
   de hoy y trabaja con eso.
2. Pregúntame **una por una** qué pasó con cada prioridad. Espera mi respuesta
   antes de pasar a la siguiente. Una frase por pregunta.
3. Escribe `outputs/AAAA-MM-DD-cierre.md`:

   ```
   ---
   fecha: AAAA-MM-DD
   tipo: cierre
   tags: [cierre, diario]
   ---
   # Cierre del AAAA-MM-DD

   ## Cerrado
   ## Abierto
   ## Reflexión
   Una línea.
   ## Mañana
   Lo que arrastra.
   ```

4. Enlaza al plan del día con `[[AAAA-MM-DD-plan]]`.
5. Marca en el plan de hoy las casillas de lo que se cerró.
6. Termina con **una sola frase**: qué es lo primero de mañana.

No me hagas más preguntas de las necesarias. Si algo está claro por las notas
del día, no me lo preguntes: dalo por hecho y dímelo.

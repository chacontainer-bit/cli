---
name: cierre-dia
description: Cierra el día conmigo delante, registra qué se hizo y deja encolado mañana. Úsala cuando pida "cierra el día", "resumen del día", "qué he hecho hoy", "terminemos por hoy" o "cierre". Para el cierre desatendido de la tarea programada usa cierre-auto.
---

# Cierre del día

1. Lee `raw/AAAA-MM-DD-plan.md` de hoy. Si no existe, lee las notas de `raw/`
   de hoy y trabaja con eso.
2. Si ya existe `outputs/AAAA-MM-DD-cierre.md` con `origen: automatico`, léelo:
   te ahorra preguntas. Lo vas a reemplazar por tu versión, que vale más porque
   yo estoy delante confirmando.
3. Pregúntame **una por una** qué pasó con cada prioridad. Espera mi respuesta
   antes de pasar a la siguiente. Una frase por pregunta.
4. Escribe `outputs/AAAA-MM-DD-cierre.md`:

   ```
   ---
   fecha: AAAA-MM-DD
   tipo: cierre
   origen: manual
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

5. Enlaza al plan del día con `[[AAAA-MM-DD-plan]]`.
6. Marca en el plan de hoy las casillas de lo que se cerró. Esto solo lo haces
   tú, con mi confirmación: el cierre automático nunca marca casillas.
7. Termina con **una sola frase**: qué es lo primero de mañana.

No me hagas más preguntas de las necesarias. Si algo está claro por las notas
del día, no me lo preguntes: dalo por hecho y dímelo.

---
name: nota-rapida
description: Captura una idea, recordatorio, dato o contacto en la bóveda sin interrumpir. Úsala cuando diga "apunta", "anota", "guarda esto", "recuérdame", "toma nota" o cuando simplemente suelte una idea sin pedirme nada a cambio.
---

# Nota rápida

1. Añade lo dicho al final de `raw/AAAA-MM-DD-notas.md`. **Un fichero por día**,
   no uno por nota. Si no existe, créalo con frontmatter
   `fecha`, `tipo: notas`, `tags: [captura]`.
2. Cada entrada va como:

   ```
   ## HH:MM — <título de 3 o 4 palabras>
   <lo que dije, limpio de muletillas pero sin reinterpretar>
   tipo:: idea | tarea | contacto | dato
   ```

3. Si menciono algo que ya existe en `wiki/`, enlázalo con wikilink. Comprueba
   antes con Glob que la nota existe.
4. Si es una tarea con fecha, ponla explícita en la entrada.
5. Confirma en **una sola frase corta**. No me repitas la nota entera: ya sé lo
   que he dicho.

No me hagas preguntas de aclaración salvo que la nota sea incomprensible. El
objetivo de esta skill es no interrumpirme.

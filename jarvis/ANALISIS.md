# JARVIS — Análisis previo a construir

Análisis del sistema del carrusel (voz → Claude Code → Obsidian → voz + HUD)
antes de escribir una sola línea. El objetivo es separar lo que es real de lo
que es marketing, y decidir qué construir y en qué orden.

---

## 1. Qué promete el carrusel y qué es cierto

| Afirmación del carrusel | Veredicto | Detalle |
|---|---|---|
| "Tú hablas, JARVIS ejecuta el trabajo" | ✅ Real | Es un bucle: STT → CLI → TTS. Nada exótico. |
| "Claude Code es el motor" | ✅ Real | Claude Code tiene modo headless (`claude -p`) pensado justo para esto. |
| "Obsidian es la memoria" | ✅ Real | Una bóveda es una carpeta de Markdown. Obsidian solo la visualiza. |
| "Sin base de datos" | ✅ Real | Ficheros + `grep`. Para volumen personal sobra. |
| "Construido con Fable 5 — funciona en cualquier modelo local" | ⚠️ Engañoso | Claude Code corre contra modelos de Anthropic. Lo **local** es la voz (STT/TTS), no el cerebro. Se puede apuntar `ANTHROPIC_BASE_URL` a un proxy con modelo local, pero pierdes skills, tool-use fiable y agentes: deja de ser el mismo producto. |
| "Tu audio nunca sale de la máquina" | ✅ Real *con matiz* | El audio no sale (Whisper y Piper corren en local). **El texto transcrito sí sale**: viaja a la API de Anthropic. Privacidad de audio ≠ privacidad de contenido. |
| "~15 minutos para cablear la voz" | ❌ Optimista | Con drivers de audio, modelos que descargar y ajustes de micro: 1–2 horas la primera vez. |
| "El HUD: una tarde" | ✅ Realista | Si lees ficheros locales y no inventas backend. |
| Métricas del HUD (135K, 202K consultas…) | 🎬 Atrezzo | Son números de la demo del creador, no un objetivo. |

**Conclusión:** el sistema es real y construible, pero el valor no está en la
voz ni en el HUD. Está en las **skills** y en la **bóveda**. Voz y HUD son la
capa bonita; si las skills son malas, tendrás un juguete caro que habla.

---

## 2. Arquitectura real

```
┌─────────────┐   audio    ┌──────────────┐  texto   ┌─────────────────┐
│ Micrófono   │ ─────────► │ STT local    │ ───────► │  jarvis.py      │
│ (push-to-   │            │ faster-      │          │  (el bucle)     │
│  talk)      │            │ whisper      │          └────────┬────────┘
└─────────────┘            └──────────────┘                   │
                                                              │ claude -p
                                                              ▼
┌─────────────┐   audio    ┌──────────────┐  texto   ┌─────────────────┐
│ Altavoz     │ ◄───────── │ TTS local    │ ◄─────── │  Claude Code    │
└─────────────┘            │ Piper / say  │          │  (headless)     │
                           └──────────────┘          └────────┬────────┘
                                                              │
                                        lee/escribe           │
                           ┌──────────────────────────────────┴──────┐
                           ▼                                          ▼
                 ┌──────────────────┐                    ┌───────────────────┐
                 │  .claude/skills/ │                    │   bóveda/         │
                 │  plan-hoy        │                    │   raw/  wiki/     │
                 │  cierre-dia      │                    │   outputs/        │
                 │  buscar-boveda   │                    │   (Markdown)      │
                 │  …               │                    └─────────┬─────────┘
                 └──────────────────┘                              │ lee
                                                                    ▼
                                                          ┌───────────────────┐
                                                          │  HUD (web local)  │
                                                          └───────────────────┘
```

Piezas y decisión técnica de cada una:

### 2.1 El motor — Claude Code headless
- Comando base: `claude -p "texto" --output-format json`, ejecutado con `cwd`
  en la carpeta de la bóveda.
- El enrutado a la skill correcta **no lo programas tú**: lo hace Claude leyendo
  las descripciones de las skills. Tu trabajo es escribir descripciones que
  disparen bien.
- Continuidad: `--continue` reanuda la última sesión. Decisión importante — ver §4.

### 2.2 La memoria — bóveda Obsidian
- Tres carpetas, como en el carrusel: `raw/` (capturado en bruto), `wiki/`
  (conocimiento depurado), `outputs/` (lo que JARVIS entrega).
- El "grafo Kárpaty" del carrusel es simplemente **wikilinks** `[[nota]]` +
  frontmatter YAML. No hay magia; el grafo es un efecto visual de Obsidian.
- Obsidian es **opcional para el sistema** y obligatorio para ti: el sistema
  funciona con ficheros; Obsidian es cómo tú los miras.

### 2.3 Los oídos y la boca — voz local
- **STT:** `faster-whisper` (Python, `pip install`) es el camino más corto.
  `whisper.cpp` es más rápido en CPU pero hay que compilarlo. Modelo `small`
  con español: buena relación precisión/velocidad.
- **TTS:** en macOS el comando `say` ya viene instalado y es instantáneo — úsalo
  en la v1. En Windows/Linux, **Piper** (voces `es_ES-davefx-medium` o
  `es_MX-claude-high`).
- **Push-to-talk:** la v1 usa ENTER para empezar y parar. Los atajos globales de
  teclado y la detección de voz (VAD) son la fuente número uno de tiempo perdido
  al principio. Se añaden después, cuando el resto funciona.

### 2.4 La cara — HUD
- **No construyas un backend nuevo.** Ya tienes `chacontainer/web` (React + Vite
  + Tailwind) y un servidor Go en `chacontainer/cmd/server`. El HUD es una ruta
  más de esa app leyendo la bóveda, no un proyecto aparte.
- Alternativa mínima si quieres separarlo: un único `index.html` + `python -m
  http.server` sobre la carpeta de la bóveda. Cero dependencias.

---

## 3. El punto crítico: el presupuesto de latencia

Esto decide si lo usas a diario o lo abandonas en dos semanas.

| Etapa | Tiempo típico | Se puede recortar |
|---|---|---|
| Grabar (tú hablando) | 3–8 s | No |
| STT local (`small`, CPU) | 1–3 s | Sí: modelo `base`, o GPU |
| **Claude Code** | **4–25 s** | **Sí, y es donde está todo el margen** |
| TTS local | 0,5–2 s | Sí: empezar a hablar por frases |
| **Total** | **9–38 s** | |

El cuello de botella es Claude, y casi siempre por una razón: **la skill hace
que lea demasiado**. Una skill que abre la bóveda entera tarda 25 s y cuesta
diez veces más que una que hace un `grep` y lee dos ficheros.

Regla operativa: **cada skill declara qué ficheros puede tocar y por qué**. Si
una respuesta tarda más de 15 s, la skill está mal escrita, no el modelo.

---

## 4. Decisiones que hay que tomar antes de empezar

**1. ¿Sesión persistente o una sesión por comando?**
- Una por comando (sin `--continue`): barato, predecible, sin contexto previo.
- Persistente (`--continue`): recuerda la conversación, pero el contexto crece y
  cada comando se encarece hasta que compacta.
- **Recomendación:** una sesión por comando. La memoria vive en la bóveda, no en
  el contexto. Es justo el punto del carrusel: *"si no está en la bóveda, no
  pasó"*.

**2. ¿Qué permisos le das al bucle de voz?**
- En headless, si la herramienta no está permitida, la ejecución falla o se
  salta el paso. La tentación es `--permission-mode bypassPermissions`.
- **No lo hagas a nivel global.** Un bucle de voz con permisos totales ejecuta
  lo que le parezca haber oído — y el STT se equivoca. Define una allowlist en
  `<bóveda>/.claude/settings.json` limitada a lectura/escritura dentro de la
  bóveda y a un puñado de comandos concretos.

**3. ¿Escritura automática en la bóveda?**
- Sí, pero **solo en `raw/` y `outputs/`**. `wiki/` es tu conocimiento curado; que
  lo edite un agente sin revisión es como dejar que se autoedite tu memoria.

**4. ¿Qué skills primero?**
El carrusel propone métricas/bandeja/tendencias, que dependen de APIs externas
(YouTube, Gmail) y son las que más tardan en funcionar. Para tu caso, empieza
por las que **no necesitan ninguna integración**:
`plan-hoy`, `cierre-dia`, `nota-rapida`, `buscar-boveda`. Con esas cuatro ya
tienes un sistema usable el mismo día. Las integraciones vienen en la fase 5.

---

## 5. Riesgos reales y mitigación

| Riesgo | Probabilidad | Mitigación |
|---|---|---|
| El STT entiende mal y se ejecuta algo no deseado | Alta | Allowlist estricta; leer en voz alta la acción antes de ejecutarla si es destructiva |
| Coste por token se dispara | Media | Skills pequeñas, `grep` antes de `read`, sin `--continue`, revisar consumo la primera semana |
| Se convierte en un juguete que no usas | **Alta** | Empezar por 3 skills que resuelvan trabajo que ya haces hoy a mano |
| Problemas de micro/drivers bloquean el proyecto | Media | Fase 3 va después de tener el sistema funcionando por teclado |
| La bóveda se llena de basura autogenerada | Media | Todo lo automático a `raw/`; promoción a `wiki/` manual |
| Dependencia de conexión | Cierta | Sin internet no hay JARVIS. La bóveda sí sigue siendo tuya y legible. |

---

## 6. Orden de construcción recomendado

El carrusel lo ordena cerebro → memoria → voz → cara. Es correcto, con un
cambio: **prueba todo por teclado antes de meter voz.**

| Fase | Qué | Tiempo real | Criterio de "funciona" |
|---|---|---|---|
| 0 | Requisitos instalados | 30 min | `claude --version` responde |
| 1 | Bóveda + `CLAUDE.md` + permisos | 30 min | `claude -p "resume la bóveda"` responde bien |
| 2 | 4 skills escritas a mano | 1–2 h | Cada skill se dispara sola al pedirla en lenguaje natural |
| 3 | Capa de voz | 1–2 h | Hablas y responde en voz alta en <20 s |
| 4 | HUD | 1 tarde | Una pantalla con vitales, agenda y últimas notas |
| 5 | Rutinas programadas + integraciones | continuo | El resumen de las 7:00 aparece solo |

**Total hasta algo usable a diario: un fin de semana.** El carrusel vende que
son minutos; son minutos por paso solo si nada falla.

---

## 7. Lo que NO hay que construir todavía

- Wake word ("Oye JARVIS"). Consume CPU en bucle y da falsos positivos. Push-to-talk
  primero.
- Base de datos o índice vectorial. `grep` sobre Markdown aguanta miles de notas.
- Autenticación en el HUD. Es local, escucha solo en `127.0.0.1`.
- Modelo local para el cerebro. Es el cambio de mayor coste y menor beneficio;
  hazlo, si acaso, cuando el sistema ya te sea imprescindible.

---

Siguiente documento: [`PASOS.md`](./PASOS.md) — los comandos concretos para el ordenador.

# MVP de CHACONTAINER OS — alcance

**Fase:** III · Producto — paso 11 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

La [ETAPA 5](./README.md#etapa-5--chacontainer-os) lista 12 módulos como
"MVP". Doce módulos no son un MVP: son el producto completo. Este documento
decide **qué subconjunto se construye primero**, con un criterio explícito:
lo mínimo para que el piloto de la [ETAPA 7](./README.md#etapa-7--piloto-packaging-systems)
pueda medir un antes y un después que el cliente sienta.

Todo lo que sigue se apoya en el roadmap técnico de
[arquitectura-os.md §3](./arquitectura-os.md#3-roadmap-de-implementación-propuesto)
(fases A-F) y en el estado real del código de `chacontainer/`.

---

## 1. El criterio de corte

Un módulo entra al MVP si cumple **las dos** condiciones:

1. **El piloto no puede medirse sin él.** La ETAPA 7 exige comparar
   inventario, pérdidas, tiempos, disponibilidad, daños, retrasos y costos
   antes vs. después.
2. **No puede sustituirse por trabajo manual durante 60-90 días sobre
   200-500 activos.** A esa escala, mucho se puede llevar en una hoja de
   cálculo sin perder el piloto; lo que no se puede es reconstruir
   retroactivamente el historial de estado y custodia.

El segundo criterio es el que más recorta. Ejemplo: registrar incidencias
en una hoja durante el piloto es viable y no invalida la medición;
reconstruir "cuánto tiempo estuvo cada activo realmente disponible" al final
del piloto, sin haberlo capturado en el momento, es imposible.

## 2. Qué entra y qué no

| # | Módulo (ETAPA 5) | ¿MVP? | Razón |
|---|---|---|---|
| 1 | Activos | ✅ Ya existe | Base de todo; `assets` + `internal/domain/asset` operan hoy |
| 2 | Identificación (QR) | ✅ Ya existe | QR completo; RFID fuera (ver [qr.md](./qr.md)) |
| 3 | Ubicaciones | ✅ Ya existe | `plants`, `plant_zones` bastan para un piloto de una sola planta |
| 4 | Movimientos | ✅ Ya existe | `asset_events`, `shipments` |
| 5 | **Custodia** | ✅ **Construir** (Fase B) | El piloto existe para descubrir quién retiene activos — sin historial de custodia no hay respuesta |
| 6 | **Condición** | ✅ **Construir** (Fase A) | `asset_state_detail`: sin los 16 estados no hay "disponibilidad real", el KPI que justifica el proyecto |
| 7 | Mantenimiento | 🟡 Parcial, sin cambios | `maintenance_records` ya captura lavado/reparación; basta para el piloto |
| 8 | Incidencias | ❌ Fuera (Fase D, post-piloto) | Registrable manualmente a escala de piloto sin invalidar la medición |
| 9 | **Inventario** | ✅ **Construir** (parcial) | El conteo por estado ya existe; falta el "requerido" (nivel mínimo) que viene con las reglas |
| 10 | **Alertas** | ✅ **Construir** (Fase C) | Es la única forma de detectar "activo detenido" en el momento, no al final (ver [alertas.md](./alertas.md)) |
| 11 | **Dashboard** | ✅ **Construir** (parcial) | Es el entregable visible del piloto (ver [dashboard.md](./dashboard.md)) |
| 12 | Evidencia | 🟡 Parcial, sin cambios | Campos de texto/JSON actuales bastan; almacenamiento de fotos queda fuera `[VALIDAR dónde se guardan hoy]` |

**MVP = 4 módulos nuevos o extendidos (5, 6, 9, 10, 11) sobre 7 que ya
funcionan.** Equivale a las fases **A + B + C** del roadmap, más una porción
de **E** (los KPI que el dashboard necesita). Quedan fuera **D**
(incidencias) y **F** (RFID).

## 3. Qué habilita este recorte

Con A + B + C, el sistema pasa a responder 12 de las 16 preguntas de
gobernanza de la [ETAPA 4](./README.md#etapa-4--modelo-de-gobernanza).
Las 4 que siguen sin respuesta —¿cuántos están dañados?, ¿cuánto cuesta cada
ciclo?, y las dos de analítica de patrón— dependen de incidencias
(Fase D) y de tener historia acumulada, que por definición un piloto de 90
días todavía no tiene.

De los 10 KPI de [kpi.md](./kpi.md), el MVP hace calculables **8**: los 4 que
ya lo eran (utilización, tiempo de ciclo, tiempo detenido, rotación) más
disponibilidad, disponibilidad real, cumplimiento de retorno y pérdida.
Quedan fuera daño y costo por ciclo.

## 4. Criterios de aceptación del MVP

El MVP está listo cuando, para un activo cualquiera del piloto, el sistema
puede responder sin intervención manual:

1. En cuál de los [16 estados](./estados-del-activo.md#1-catálogo-de-estados) está, y desde cuándo.
2. Quién es su custodio actual y quiénes lo fueron antes, con fechas.
3. Si incumple alguna [regla operativa](./reglas-operativas.md) vigente, y desde hace cuánto.
4. Cuántos activos como él están realmente disponibles hoy (no contablemente).
5. Todo lo anterior consultable desde el dashboard sin abrir la base de datos.

Si alguno de los cinco requiere que alguien revise una hoja de cálculo, el
MVP no está terminado.

## 5. Lo que el MVP deliberadamente no resuelve

- **Multi-planta**: el piloto es de una planta ([ETAPA 7](./README.md#etapa-7--piloto-packaging-systems)). El esquema ya es multi-tenant y multi-planta, así que no hay deuda estructural, pero no se optimiza para ello.
- **Autoservicio del cliente**: el cliente del piloto recibe un reporte ([modelo-de-gobernanza.md §3](./modelo-de-gobernanza.md#3-ritual-de-gobernanza-propuesto)), no un login. Eso es Nivel 6 de la [escalera comercial](./README.md#etapa-9--escalera-comercial), no MVP.
- **Integraciones ERP/MES**: FASE V. Airtable y Make.com ya existen y siguen operando como están.

---

## Próximo paso

Pasos 12-14 detallan los tres módulos que el MVP construye de cara al
usuario: [QR](./qr.md), [Dashboard](./dashboard.md) y
[Alertas](./alertas.md). El paso 15 es el [plan de piloto](./piloto.md) que
los pone a prueba.

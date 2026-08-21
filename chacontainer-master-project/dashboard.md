# Dashboard — indicadores y vistas

**Fase:** III · Producto — paso 13 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

En `chacontainer/web/` ya existen las páginas `Dashboard.jsx` y
`Traceability.jsx`, y en la base tres vistas materializadas
(`mv_plant_inventory`, `mv_shipment_kpis`, `mv_daily_scans`). Este documento
no parte de cero: define **qué debe mostrar el dashboard para sostener el
modelo de gobernanza**, y qué de eso todavía no puede calcularse.

---

## 1. Principio de organización: por cadencia, no por módulo

El error natural sería una pantalla por módulo (activos, custodia, alertas…).
Eso produce un dashboard que nadie usa, porque nadie tiene el trabajo de
"revisar el módulo de custodia".

La organización correcta sale de
[modelo-de-gobernanza.md §1](./modelo-de-gobernanza.md#1-los-tres-niveles-de-gobernanza):
tres personas distintas miran tres cosas distintas con tres cadencias
distintas. Por lo tanto, **tres vistas**:

| Vista | Para quién | Cadencia | Pregunta que contesta |
|---|---|---|---|
| **Operación** | Responsable de planta (Nivel 2) | Varias veces al día | ¿Qué requiere mi atención ahora? |
| **Estado del parque** | Operación y comercial (Nivel 1, consulta) | Bajo demanda | ¿Qué tengo, dónde, en qué condición? |
| **Gobernanza** | Gobernanza (Nivel 3) y cliente | Mensual | ¿Está mejorando el sistema y dónde está el dinero? |

## 2. Vista Operación — la bandeja de excepciones

No es un tablero de métricas: es una **lista de trabajo**. Todo lo que
aparece aquí exige una acción o una decisión hoy.

| Bloque | Contenido | Fuente |
|---|---|---|
| Alertas activas | Activos que incumplen una regla, ordenados por antigüedad del incumplimiento | `alerts` ([alertas.md](./alertas.md)) |
| Activos detenidos | En `Retenido`, con días transcurridos y custodio responsable | `asset_state_detail` + `asset_custody` |
| Pendientes de liberación | En `Liberado` sin pasar a `Disponible` — inventario que existe pero nadie puede asignar | `asset_state_detail` |
| Bloqueados | En `Bloqueado`, con la disputa asociada | `asset_state_detail` ([SOP-17](./sop/SOP-ACTIVO-17-activo-bloqueado.md)) |
| Escaneos rechazados | Intentos de transición inválida ([qr.md §3](./qr.md#3-validación-de-transición-en-el-escaneo)) | `qr_scans` |

El bloque de escaneos rechazados es el más fácil de omitir y uno de los más
útiles: un operador que insiste en despachar activos sin lavar es un
problema de proceso que ningún KPI mensual va a revelar a tiempo.

## 3. Vista Estado del parque

Extiende lo que hoy hace `mv_plant_inventory` (conteo por planta/tipo/estado
agregado) con la granularidad de los 16 estados.

- **Conteo por estado detallado**, no por los 6 valores de `assets.status`. La diferencia práctica: hoy `maintenance` mezcla `Sucio`, `En reparación`, `En inspección` y `Liberado`; el responsable de planta necesita distinguirlos porque cada uno tiene un cuello de botella distinto.
- **Disponibilidad real vs. contable**, lado a lado. Es el número que justifica todo el proyecto ([kpi.md](./kpi.md#1-los-10-kpi-especificados)) y debe leerse de un vistazo: *"tienes 480 activos; 310 están realmente disponibles"*.
- **Distribución por custodio**, con antigüedad. Responde "¿quién los tiene?" y "¿desde cuándo?" en una sola tabla.
- **Trazabilidad por activo**: la página `Traceability.jsx` ya existe; se extiende para mostrar la línea de tiempo de estados y custodios, no solo de ubicaciones.

## 4. Vista Gobernanza

Los KPI de la [ETAPA 6](./README.md#etapa-6--indicadores-del-sistema), con
dos reglas de presentación que importan más que la selección de gráficas:

1. **Siempre contra la línea base del piloto**, nunca en absoluto. Un 82% de
   disponibilidad no dice nada; "82% contra 61% al inicio" es el argumento
   comercial completo ([ETAPA 7](./README.md#etapa-7--piloto-packaging-systems)).
2. **Segmentable por custodio, planta y familia de activo**, porque las
   decisiones de Nivel 3 son casi siempre "¿dónde intervenir primero?" y esa
   respuesta es un segmento, no un promedio.

KPI disponibles en el MVP (8 de 10, ver [mvp.md §3](./mvp.md#3-qué-habilita-este-recorte)):
disponibilidad, disponibilidad real, utilización, tiempo de ciclo,
cumplimiento de retorno, pérdida, tiempo detenido, rotación. Daño y costo
por ciclo quedan pendientes de la Fase D del roadmap y **deben mostrarse
como "no disponible aún"**, no omitirse en silencio: un KPI ausente sin
explicación se lee como un KPI en cero.

## 5. Implementación

- **Vistas materializadas nuevas** sobre `asset_state_detail`, `asset_custody` y `alerts`, siguiendo el patrón ya establecido (refresh periódico, índice único por tenant) — corresponde a la Fase E de [arquitectura-os.md](./arquitectura-os.md#3-roadmap-de-implementación-propuesto).
- **La vista Operación no debe leer de vistas materializadas.** Un refresh cada 5 minutos es correcto para KPI mensuales e inútil para una bandeja de excepciones: si un activo entra en `Retenido`, el responsable de planta tiene que verlo ya. Esa vista consulta las tablas directamente.
- El reporte mensual al cliente ([modelo-de-gobernanza.md §3](./modelo-de-gobernanza.md#3-ritual-de-gobernanza-propuesto)) es una exportación de la vista Gobernanza filtrada a ese cliente, no un desarrollo aparte.

---

## Próximo paso

[Alertas](./alertas.md) (paso 14): el mecanismo que llena la bandeja de
excepciones de la vista Operación.

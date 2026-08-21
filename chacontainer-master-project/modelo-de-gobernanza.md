# Modelo de gobernanza

**Fase:** II · Sistema — paso 7 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

La [ETAPA 4](./README.md#etapa-4--modelo-de-gobernanza) del proyecto maestro
lista 16 preguntas que el sistema debe poder responder, y
[catalogo-systems.md #15](./catalogo-systems.md#15-gobernanza) ya cruza cada
pregunta contra la capacidad que la responde. Lo que falta — y es el objeto
de este documento — es **quién** responde con esa información, **con qué
cadencia**, **qué umbral** convierte un dato en una acción, y **qué pasa**
cuando la acción no es automática.

Gobernanza es la única capacidad de Systems sin dato propio
([catalogo-systems.md §Lectura cruzada](./catalogo-systems.md#lectura-cruzada-qué-capacidad-de-systems-resuelve-cada-pregunta-de-gobernanza)):
consume lo que producen las otras 15. Este documento formaliza esa función
de consumo.

---

## 1. Los tres niveles de gobernanza

No toda excepción necesita al mismo responsable. Modelamos tres niveles,
correspondientes a la [ETAPA 9 · Escalera comercial](./README.md#etapa-9--escalera-comercial)
(Niveles 4, 5 y 6):

| Nivel | Quién actúa | Qué resuelve | Cadencia |
|---|---|---|---|
| **1 · Automático** | El sistema (Reglas operativas + Alertas, ver [paso 9](./reglas-operativas.md)) | Excepciones con acción predefinida sin ambigüedad (activo excede tiempo máximo fuera → pasa a `Retenido` y notifica) | Tiempo real |
| **2 · Operativo** | Responsable de planta / operador designado `[VALIDAR rol exacto]` | Excepciones que requieren una decisión de campo, pero dentro de un criterio ya definido (¿reparar o dar de baja este activo según el criterio de costo ya acordado?) | Diaria / por turno |
| **3 · Estratégico** | Gobernanza (fundador/responsable de cuenta) `[VALIDAR quién exactamente por cliente]` | Decisiones que ajustan el criterio mismo: cambiar un umbral, renegociar condiciones con un custodio recurrentemente moroso, decidir si un patrón de pérdida amerita rediseñar la ruta o el contrato | Semanal / mensual (ver §3) |

El nivel 1 depende de que existan reglas configuradas (ver
[reglas-operativas.md](./reglas-operativas.md)); sin regla, toda excepción
sube por defecto al nivel 2 — nunca debe quedar sin dueño.

## 2. Las 16 preguntas, con dueño y cadencia

Extiende la tabla de [catalogo-systems.md](./catalogo-systems.md#lectura-cruzada-qué-capacidad-de-systems-resuelve-cada-pregunta-de-gobernanza)
agregando responsable y cadencia de revisión.

| Pregunta (ETAPA 4) | Capacidad que responde | Nivel | Cadencia de revisión |
|---|---|---|---|
| ¿Cuántos activos existen? | Registro + Inventario digital | 1 (consulta) | Continua (dashboard) |
| ¿Cuántos están realmente disponibles? | Inventario digital | 1 (consulta) | Continua |
| ¿Dónde están? | Trazabilidad | 1 (consulta) | Continua |
| ¿Quién los tiene? | Gestión de custodios | 1 (consulta) | Continua |
| ¿Cuánto tiempo llevan ahí? | Trazabilidad + Reglas operativas | 1 → 2 si excede umbral | Continua / por excepción |
| ¿Cuál es su condición? | Registro (Capa 4) | 1 (consulta) | Continua |
| ¿Cuándo deberían regresar? | Reglas operativas | 1 (consulta) | Continua |
| ¿Cuántos están detenidos? | Alertas | 2 | Diaria |
| ¿Cuántos están dañados? | Gestión de incidencias | 2 | Diaria |
| ¿Cuántos faltan? | Inventario digital vs. registro | 2 → 3 si es recurrente | Semanal |
| ¿Cuánto cuesta cada ciclo? | Indicadores | 3 | Mensual |
| ¿Dónde se producen las pérdidas? | Analítica | 3 | Mensual |
| ¿Qué proveedor/custodio retiene activos? | Gestión de custodios + Analítica | 3 | Mensual |
| ¿Qué planta tiene exceso/déficit? | Analítica | 3 | Mensual |
| ¿Qué activos conviene reparar? | Analítica | 2 (caso a caso) / 3 (política) | Diaria (caso) / mensual (política) |
| ¿Qué activos deben darse de baja? | Analítica + Gobernanza | 3 | Mensual (ligado a SOP-ACTIVO-18) |

Patrón general: las preguntas de **estado actual** (dónde, quién, cuánto)
son consulta continua de nivel 1; las preguntas de **excepción puntual**
(detenidos, dañados) son revisión operativa diaria de nivel 2; las
preguntas de **patrón y costo** (pérdida, exceso/déficit, política de
reparación) son revisión estratégica mensual de nivel 3.

## 3. Ritual de gobernanza propuesto

| Ritual | Frecuencia | Participantes `[VALIDAR]` | Insumo | Salida |
|---|---|---|---|---|
| Revisión de excepciones abiertas | Diaria | Responsable de planta | Alertas activas, incidencias abiertas (SOP-ACTIVO-15/16/17) | Excepciones resueltas o escaladas a nivel 3 |
| Comité de gobernanza | Mensual `[VALIDAR periodicidad real deseada]` | Gobernanza + responsables de planta | Indicadores (ETAPA 6), Analítica, bitácora de excepciones del mes | Ajustes a reglas operativas, decisiones de baja masiva, escalamiento comercial a clientes/custodios problemáticos |
| Reporte de gobernanza al cliente | Mensual (si el contrato es Nivel 5/6) | Gobernanza + contacto del cliente | Mismo insumo que el comité, filtrado a los activos de ese cliente | Reporte entregable (evidencia del [Módulo 12 de CHACONTAINER OS](./catalogo-systems.md#16-chacontainer-os)) — este reporte es en sí mismo la prueba de valor que sostiene el Nivel 5/6 de la escalera comercial |

## 4. Qué activa una escalada de nivel 2 a nivel 3

No toda excepción de nivel 2 debe escalar. Escala cuando:

1. **Es recurrente** — el mismo tipo de excepción ocurre 3+ veces `[VALIDAR umbral]` con el mismo cliente/custodio/planta en el período de revisión.
2. **Excede un umbral económico** — el costo acumulado de la excepción (reparaciones repetidas, activos perdidos) supera un monto definido `[VALIDAR monto]`.
3. **Requiere cambiar una regla, no solo resolver un caso** — si resolver la excepción puntual no evita que se repita, el problema es de política (nivel 3), no de ejecución (nivel 2).
4. **Involucra una relación comercial** — cualquier excepción que amerite conversación con el cliente/custodio sobre condiciones, no solo con el operador interno.

## 5. Relación con el resto del proyecto

- Este documento asume que existen [Reglas operativas](./reglas-operativas.md) configurables — es el paso 9, siguiente en esta misma fase.
- Los umbrales que aparecen aquí como `[VALIDAR]` son, en su mayoría, los mismos parámetros de Capa 5 · Gobierno de la [ETAPA 3](./README.md#etapa-3--modelo-de-packaging-systems) (tiempo máximo fuera, punto de retorno, nivel mínimo de inventario) — no se inventan de nuevo, se validan una sola vez y se referencian desde ambos documentos.
- El nivel 3 (Gobernanza) es lo que se vende como Nivel 5/6 de la [escalera comercial](./README.md#etapa-9--escalera-comercial) — este documento es, en el fondo, la especificación operativa de qué implica cobrar un fee de gobernanza.

---

## Próximo paso sugerido

Paso 8 de la FASE II: **modelo de datos** — traducir las 5 capas, los 16
estados y las reglas/alertas/incidencias de este documento a un esquema de
base de datos concreto, cotejado contra el esquema real que ya existe en
`chacontainer/migrations/`.

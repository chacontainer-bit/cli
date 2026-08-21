# Escala · Integraciones ERP / MES

**Fase:** V · Escala — paso 24 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

El SaaS ya tiene infraestructura de integración: un paquete
`internal/integrations/erp`, conectores de Airtable y Make.com, y una tabla
`airtable_sync_log` con dirección `push`/`pull`. Este documento no diseña
plomería nueva. Resuelve la pregunta que decide si una integración funciona
o se vuelve una fuente permanente de discusiones: **quién manda sobre cada
dato**.

---

## 1. El problema real: dos sistemas, una verdad

El ERP del cliente ya tiene registros de sus activos: cuántos son, cuánto
valen, cuándo se compraron, a qué centro de costo pertenecen. CHACONTAINER
OS también los tiene. En el momento en que ambos existen, cada campo
duplicado es una pregunta pendiente: si difieren, ¿cuál tiene razón?

La respuesta no puede ser "los sincronizamos". Sincronizar sin decidir el
dueño de cada campo es el modo de falla clásico: dos sistemas escribiendo el
mismo dato, un conflicto que se resuelve por orden de llegada, y nadie
capaz de explicar por qué el número cambió.

## 2. La frontera propuesta

El criterio es sencillo: **manda quien observa el hecho.**

| Dominio | Sistema de registro | Por qué |
|---|---|---|
| Existencia del activo, identificador, tipo, modelo | **ERP** (o CHACONTAINER OS si el activo es propio) | El alta nace de una compra, y la compra vive en el ERP |
| Valor contable, depreciación, centro de costo, propiedad | **ERP** | Es información financiera; el OS no tiene por qué opinar |
| Estado físico, ubicación, custodia, condición | **CHACONTAINER OS** | Nace de un escaneo o una inspección en piso — el ERP nunca observa esto |
| Historial de movimientos y transiciones | **CHACONTAINER OS** | Es la línea de tiempo que produce la trazabilidad |
| Incidencias, alertas, reglas | **CHACONTAINER OS** | No tienen equivalente en el ERP |
| Órdenes de compra y facturación | **ERP** | — |

De ahí sale la regla operativa: **ningún campo se escribe desde los dos
lados.** Cada campo tiene un dueño y el otro sistema lo lee. Si el ERP es
dueño del valor contable, el OS lo muestra pero no lo edita; si el OS es
dueño del estado físico, el ERP lo consume pero no lo sobrescribe.

Esta frontera además refuerza el argumento comercial de la
[Oferta 4](./oferta-systems.md#oferta-4--gobernanza-del-sistema-nivel-5):
el ERP del cliente sabe cuántos activos *compró*; solo CHACONTAINER OS sabe
cuántos puede *usar hoy*. Son preguntas distintas y por eso conviven dos
sistemas.

## 3. ERP y MES no son el mismo problema

Se mencionan juntos y se integran distinto.

| | ERP | MES |
|---|---|---|
| Qué gobierna | Compras, inventario contable, finanzas | Ejecución de producción en piso |
| Cadencia | Batch — diaria o por evento de negocio | Tiempo real o casi |
| Qué necesita de CHACONTAINER | Disponibilidad real y valor del parque | Si hay empaque disponible **ahora** para la corrida que arranca |
| Qué aporta a CHACONTAINER | Alta de activos, datos de propiedad y costo | La demanda real de empaque por línea y turno |
| Patrón de integración | Sincronización programada | Consulta bajo demanda o evento |

El valor de MES es el que se subestima: es la única fuente que dice
**cuántos activos se necesitan realmente**, que es el denominador del KPI de
[disponibilidad](./kpi.md#1-los-10-kpi-especificados) — hoy marcado como no
calculable justamente porque "activos requeridos" no existe en ningún lado.
Una integración MES convierte ese denominador de supuesto en dato.

## 4. Qué lo dispara

- **ERP**: cuando el cliente exige que el alta de activos no se capture dos veces, o cuando auditoría le pide conciliar físico contra contable de forma recurrente. Suele aparecer solo en cuentas de [Oferta 4](./oferta-systems.md#oferta-4--gobernanza-del-sistema-nivel-5).
- **MES**: cuando el problema declarado del cliente es **paro de línea por falta de empaque**, no pérdida de activos. Son dolores distintos y el segundo no se resuelve con trazabilidad sola.

Ninguna integración debe ofrecerse antes del piloto. Integrar es caro,
específico de cada cliente, y solo tiene sentido cuando ya se demostró que
el dato de CHACONTAINER vale algo.

## 5. Qué no hay que romper hoy

- **Mantener `airtable_sync_log` como patrón, no como excepción.** Toda integración necesita bitácora con dirección, estado y error. Ya existe el patrón; conviene que ERP y MES lo reutilicen en vez de inventar su propio log.
- **No dejar que una integración escriba estados.** Un ERP que pueda marcar un activo como `Disponible` rompe las [invariantes de la máquina de estados](./estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar) — el estado se gana pasando por `Liberado`, no por declaración externa. Las integraciones leen estado; no lo escriben.
- **No prometer tiempo real sobre una base pensada para batch.** Las vistas materializadas se refrescan periódicamente ([dashboard.md §5](./dashboard.md#5-implementación)); un MES que consulte disponibilidad necesita leer tablas directas, con el costo que eso implica.

---

Continúa en [Escala geográfica](./escala-geografica.md) (pasos 25-27).

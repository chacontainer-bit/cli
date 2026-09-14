# Piloto Packaging Systems — plan de ejecución

**Fase:** III · Producto — paso 15 (último) de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

La [ETAPA 7](./README.md#etapa-7--piloto-packaging-systems) define el piloto
en términos generales: una planta, un cliente, 200-500 activos, 60-90 días,
medir antes y después. Este documento lo convierte en un plan ejecutable y
resuelve el problema que la ETAPA 7 deja abierto: **cómo se mide el "antes"
si el sistema que mide es justamente lo que no existe todavía**.

---

## 1. El problema de la línea base

La ETAPA 7 pide comparar inventario, pérdidas, tiempos, disponibilidad,
daños, retrasos y costos antes vs. después. Pero antes de intervenir no hay
`asset_state_detail`, ni custodia registrada, ni alertas. El cliente no sabe
esos números — esa ignorancia es la razón por la que se contrata el piloto.

La salida no es estimar: es **usar un servicio de Solutions como instrumento
de medición**. Un inventario físico
([catalogo-solutions.md #11](./catalogo-solutions.md#11-inventarios-físicos))
produce, en una sola intervención, el conteo real, los faltantes y el
porcentaje de activos sin identificar. Eso es la línea base, y además es
facturable.

Esto no es un rodeo: es exactamente el mecanismo de la
[ETAPA 8](./README.md#etapa-8--modelo-comercial-solutions--systems). El
inventario físico es el punto de entrada comercial *y* el instrumento de
medición del piloto. La misma actividad hace las dos cosas.

**Lo que la línea base sí puede capturar** (día 0, con inventario físico):
conteo real vs. contable, faltantes, % sin identificar, condición por
categoría (vía [clasificación](./sop/SOP-ACTIVO-04-clasificacion-de-condicion.md)),
y —por entrevista, no por sistema— el tiempo de ciclo percibido y las
pérdidas anuales estimadas por el cliente `[VALIDAR: dejar constancia
escrita de que estas dos son declaradas, no medidas]`.

**Lo que no puede capturar y hay que aceptar**: tiempo detenido real,
cumplimiento de retorno y costo por ciclo. Son series temporales; no existen
retroactivamente. Para estos tres, la comparación honesta no es
antes-vs-después sino **primer mes del piloto vs. último mes**.

## 2. Criterios de selección

La ETAPA 7 lista qué seleccionar. Los criterios para elegir bien:

| Dimensión | Criterio | Por qué |
|---|---|---|
| Cliente | Que ya sea cliente de Solutions, con dolor declarado de faltantes | El piloto no es el momento de abrir cuenta nueva; se apoya en una relación existente ([SOP-10 comercial](../chacontainer/docs/sop/SOP-10-postventa-y-expansion.md)) |
| Planta | Una sola, propia o del cliente, con volumen suficiente de rotación | Multi-planta introduce variables que el piloto no puede aislar |
| Familia de activo | Una sola, homogénea | Si se mezclan IBC y tarimas, no se sabrá a cuál atribuir la mejora |
| Volumen | 200-500 activos | Menos no da señal estadística; más encarece la línea base sin agregar aprendizaje |
| Ruta | Una, con retorno recurrente | El KPI de cumplimiento de retorno necesita ciclos completos dentro de la ventana |
| Problema | Uno concreto y declarado por el cliente | El éxito se juzga contra el problema que el cliente dijo tener, no contra el que nosotros veamos |

**El criterio de la ruta es el que más restringe la duración**: el piloto
debe durar al menos 2-3 ciclos completos de esa ruta. Si el ciclo es de 30
días, 90 días es el mínimo, no el máximo. Si el ciclo es de 10 días, 60 días
sobran. La duración se deriva del ciclo, no del calendario.

## 3. Cronograma

| Semana | Actividad | Entregable |
|---|---|---|
| −2 a 0 | Diagnóstico ([Systems #1](./catalogo-systems.md#1-diagnóstico-de-activos)) y acuerdo de alcance | Alcance firmado, problema declarado por escrito |
| 0 | Inventario físico + clasificación | **Línea base** (§1) |
| 0-1 | Identificación QR del parque del piloto | 100% del parque identificado ([SOP-02](./sop/SOP-ACTIVO-02-identificacion-qr-rfid.md)) |
| 1 | Configuración de reglas globales y del cliente | `operational_rules` cargadas ([reglas-operativas.md](./reglas-operativas.md)) |
| 1-2 | Capacitación de operadores en el catálogo de escaneo | Operadores usando las acciones de [qr.md §2](./qr.md#2-catálogo-de-acciones-de-escaneo) |
| 2-12 | Operación con el sistema; revisión diaria de excepciones | Bitácora de excepciones; alertas atendidas |
| 4, 8, 12 | Comité de gobernanza mensual | Ajustes a reglas; reporte al cliente |
| 12-13 | Medición final y comparación | Reporte de cierre del piloto |

Las semanas 0-2 no son preparación: son ya el servicio. El cliente recibe
valor (inventario real, parque identificado) antes de que el software
demuestre nada. Si el piloto se cancelara en la semana 3, el cliente ya
tendría algo que no tenía.

## 4. Criterios de éxito

El piloto es exitoso si al cierre se cumplen **los tres**:

1. **Cuantitativo**: mejora medible en el indicador ligado al problema
   declarado por el cliente en la semana −2. No en todos los KPI — en el que
   el cliente dijo que le dolía. `[VALIDAR: fijar el umbral de mejora
   con el cliente ANTES de empezar, no al final]`
2. **Operativo**: los operadores usan el sistema sin supervisión especial en
   las últimas 4 semanas. Si a la semana 12 alguien de CHACONTAINER todavía
   tiene que recordarles escanear, el proceso no se adoptó.
3. **De datos**: el sistema puede responder los 5 criterios de aceptación del
   [MVP §4](./mvp.md#4-criterios-de-aceptación-del-mvp) para cualquier
   activo del piloto, sin intervención manual.

El criterio 2 es el que más pilotos reprueba y el que menos se mide.

## 5. Conversión a contrato

El piloto termina en una de tres salidas, y conviene decidir de antemano
cuál corresponde a cada resultado:

| Resultado | Salida | Nivel de la [escalera](./README.md#etapa-9--escalera-comercial) |
|---|---|---|
| Los tres criterios cumplidos | Contrato de gobernanza recurrente sobre el parque del piloto, con plan de extensión a otras plantas/familias | Nivel 5 |
| Cuantitativo y de datos, pero no operativo | Extensión del piloto 60 días con foco en adopción, no en más funcionalidad | Sigue en piloto |
| No cuantitativo | Cierre con entrega de la línea base y los hallazgos; se conserva la relación de Solutions | Nivel 1-2 |

La tercera fila importa: un piloto que no mejora el indicador **no es un
fracaso comercial** si dejó al cliente con su parque identificado y su
inventario real. Ese es el piso que el diseño de §3 garantiza.

## 6. Riesgos

| Riesgo | Control |
|---|---|
| El cliente cambia el alcance a mitad (más activos, otra planta) | Alcance firmado en semana −2; cualquier cambio reinicia la línea base o se documenta como piloto distinto |
| Los operadores no escanean y el dato queda incompleto | Métrica de adopción semanal (escaneos esperados vs. reales) desde la semana 2, no al final |
| La mejora se atribuye al piloto pero viene de otra causa (temporada, cambio de proveedor) | Registrar en la bitácora todo cambio operativo externo al piloto; mencionarlo en el reporte de cierre |
| El parque del piloto se mezcla con activos fuera de alcance | La identificación de la semana 0-1 delimita físicamente el alcance: sin QR del piloto, el activo no es del piloto |
| Se descubren tantos faltantes que la conversación se vuelve conflictiva | Acordar en semana −2 que la línea base es diagnóstico, no auditoría de responsabilidades `[VALIDAR encuadre con el cliente]` |

---

## Cierre de FASE III

Con este documento se completan los 5 pasos de la FASE III · Producto:
[MVP](./mvp.md) (11), [QR](./qr.md) (12), [Dashboard](./dashboard.md) (13),
[Alertas](./alertas.md) (14) y este plan de piloto (15).

Las tres primeras fases del proyecto —Fundación, Sistema y Producto— quedan
cerradas. Sigue la **FASE IV · Comercial** ([ETAPA 14](./README.md#etapa-14--orden-de-ejecución)):
oferta Systems, pricing, diagnóstico comercial, piloto pagado y conversión
Solutions → Systems.

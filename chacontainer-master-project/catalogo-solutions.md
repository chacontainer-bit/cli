# Catálogo maestro · CHACONTAINER Packaging Solutions

**Fase:** I · Fundación — paso 1 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

Este catálogo cubre los 15 servicios de Packaging Solutions definidos en el
[Proyecto Maestro](./README.md#etapa-0--fundamento-del-modelo). Cada servicio
se documenta con los 7 campos que exige la ETAPA 0, para que Solutions deje de
ser solo ejecución y empiece a producir el dato que alimenta Packaging Systems.

Convención: activo retornable = contenedor, IBC, tarima/pallet o tambor
(terminología alineada con [SOP-02](../chacontainer/docs/sop/SOP-02-lavado-de-activos-retornables.md)
y [lista-precios-costos.md](../chacontainer/docs/lista-precios-costos.md) del SaaS existente).
Los campos marcados `[VALIDAR]` dependen de una decisión operativa o comercial
que aún no está definida — no inventar el dato, confirmarlo con el responsable
antes de operar sobre él.

---

## Índice

1. [Venta](#1-venta)
2. [Renta](#2-renta)
3. [Lavado](#3-lavado)
4. [Reparación](#4-reparación)
5. [Reacondicionamiento](#5-reacondicionamiento)
6. [Modificación](#6-modificación)
7. [Dunnage](#7-dunnage)
8. [Reetiquetado](#8-reetiquetado)
9. [Inspección](#9-inspección)
10. [Clasificación](#10-clasificación)
11. [Inventarios físicos](#11-inventarios-físicos)
12. [Recuperación](#12-recuperación)
13. [Compra de scrap](#13-compra-de-scrap)
14. [Disposición final](#14-disposición-final)
15. [Recuperación de valor](#15-recuperación-de-valor)

---

## 1. Venta

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El cliente necesita empaque retornable (IBC, tambor, tarima, contenedor) nuevo o certificado y no quiere gestionar fabricación, importación o inventario propio. |
| **Qué recibe CHACONTAINER** | Especificación técnica del activo (tipo, capacidad, material, cantidad), destino de uso y volumen estimado de compra. |
| **Actividad que ejecuta** | Cotiza, produce o abastece el activo, factura y entrega en la ubicación pactada. |
| **Evidencia que genera** | Orden de venta, ficha técnica del activo entregado, guía de entrega firmada, número de serie/SKU asignado. |
| **Indicador que afecta** | Ticket promedio, margen por SKU, tiempo de entrega, activos nuevos incorporados al parque total administrable. |
| **Información que produce** | Alta del activo en el registro maestro con ID único desde el día cero — primer dato de [Capa 1 · Activo](./README.md#etapa-3--modelo-de-packaging-systems). |
| **Servicio que puede venderse después** | Identificación (QR/RFID) del activo recién vendido, renta de activos complementarios, contrato de lavado/mantenimiento periódico. |

## 2. Renta

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El cliente necesita empaque retornable por un periodo definido sin inmovilizar capital en compra ni asumir el mantenimiento del activo. |
| **Qué recibe CHACONTAINER** | Solicitud con tipo de activo, cantidad, periodo de renta y ubicación de entrega/retorno. |
| **Actividad que ejecuta** | Asigna activos disponibles del parque propio, entrega, factura por periodo, monitorea vencimiento y gestiona el retorno. |
| **Evidencia que genera** | Contrato de renta, acta de entrega y de retorno, registro de condición del activo en ambos extremos. |
| **Indicador que afecta** | Tasa de utilización del parque en renta, tiempo de ciclo, cumplimiento de retorno, ingresos recurrentes por activo. |
| **Información que produce** | Historial de custodia (Capa 3) y de ubicación (Capa 2) del activo durante todo el periodo; primer dato real de "tiempo fuera" del activo. |
| **Servicio que puede venderse después** | Trazabilidad/alertas de vencimiento, lavado al retorno, renovación o conversión a contrato de gestión (Nivel 3+ de la [escalera comercial](./README.md#etapa-9--escalera-comercial)). |

## 3. Lavado

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El activo retorna sucio y no puede reingresar a inventario disponible ni entregarse a un nuevo cliente sin riesgo de contaminación cruzada o reclamo. |
| **Qué recibe CHACONTAINER** | Activo sucio proveniente de retorno de campo, con o sin identificación previa. |
| **Actividad que ejecuta** | Ver [SOP-02](../chacontainer/docs/sop/SOP-02-lavado-de-activos-retornables.md): clasifica, lava según tipo/producto previo, inspecciona y libera o rechaza. |
| **Evidencia que genera** | Registro de ingreso, checklist de inspección post-lavado, foto y/o escaneo QR del resultado (apto/rechazado). |
| **Indicador que afecta** | Tiempo de ciclo de lavado, % de rechazo, % de reclamos por limpieza, rotación de inventario limpio disponible. |
| **Información que produce** | Condición del activo (Capa 4 · Estado) en dos momentos (sucio → limpio), y evidencia de que el activo *existe y sigue en circulación* aunque no tenga dueño identificado aún. |
| **Servicio que puede venderse después** | Identificación (el punto de contacto más frecuente para detectar activos sin QR/RFID — ver [ETAPA 8](./README.md#etapa-8--modelo-comercial-solutions--systems)), inspección estructural, reparación de lo que el lavado deja visible. |

## 4. Reparación

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El activo tiene daño estructural o funcional que le impide reingresar a inventario disponible tal como está. |
| **Qué recibe CHACONTAINER** | Activo dañado, generalmente derivado de un rechazo en lavado o inspección, con diagnóstico preliminar del daño. |
| **Actividad que ejecuta** | Diagnostica el daño, cotiza (si aplica costo al cliente), repara o determina que no es reparable, y libera o deriva a baja/scrap. |
| **Evidencia que genera** | Diagnóstico de falla, orden de reparación, registro de refacciones/insumos usados, evidencia fotográfica antes/después. |
| **Indicador que afecta** | % de activos dañados sobre retornados, costo por reparación, tiempo de reparación, tasa de recuperación vs. baja. |
| **Información que produce** | Tipo de falla más frecuente por SKU/familia de activo — insumo directo para decidir qué activos "conviene reparar" ([ETAPA 4](./README.md#etapa-4--modelo-de-gobernanza)). |
| **Servicio que puede venderse después** | Reacondicionamiento si el daño es mayor a reparación puntual, modificación si el cliente aprovecha la intervención para actualizar el activo, contrato de mantenimiento preventivo. |

## 5. Reacondicionamiento

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El activo es funcional pero está degradado por uso/tiempo (estético, estructural menor) y el cliente quiere extender su vida útil en vez de reemplazarlo. |
| **Qué recibe CHACONTAINER** | Activo usado con desgaste generalizado, no un daño puntual (lo distingue de reparación). |
| **Actividad que ejecuta** | Interviene integralmente el activo (limpieza profunda, reparación de múltiples puntos, repintado/reetiquetado si aplica) para devolverlo a condición "como nuevo" o al estándar acordado. |
| **Evidencia que genera** | Diagnóstico integral pre-intervención, orden de reacondicionamiento, registro de condición final vs. estándar. |
| **Indicador que afecta** | Costo de reacondicionamiento vs. costo de reemplazo (decisión reparar-vs-dar de baja), vida útil extendida por ciclo, rotación. |
| **Información que produce** | Curva de degradación del activo por número de ciclos — dato clave para fijar "punto de retorno" y condición permitida ([Capa 5 · Gobierno](./README.md#etapa-3--modelo-de-packaging-systems)). |
| **Servicio que puede venderse después** | Reetiquetado, modificación, contrato de gestión del ciclo de vida completo (el reacondicionamiento es un punto de entrada natural a Systems porque exige diagnóstico consultivo — ver [SOP-03](../chacontainer/docs/sop/SOP-03-diagnostico-consultivo.md)). |

## 6. Modificación

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El activo estándar no encaja con un requerimiento específico del cliente (dimensiones, accesorios, compatibilidad con línea de producción). |
| **Qué recibe CHACONTAINER** | Activo base (propio o del cliente) más especificación técnica de la modificación requerida. |
| **Actividad que ejecuta** | Diseña y ejecuta la modificación (corte, adaptación, montaje de accesorios), valida contra la especificación. |
| **Evidencia que genera** | Especificación técnica aprobada, orden de modificación, evidencia de la modificación ejecutada, validación de conformidad. |
| **Indicador que afecta** | Tiempo de diseño+ejecución, % de modificaciones aceptadas a la primera, margen por proyecto (suele cotizarse a medida). |
| **Información que produce** | Variantes de SKU por cliente/línea — obliga a que el registro maestro de activos (Módulo 1 de CHACONTAINER OS) soporte configuraciones no estándar. |
| **Servicio que puede venderse después** | Dunnage (modificaciones suelen derivar en necesidad de protección interna a medida), venta de más unidades de la variante ya validada, identificación específica para la nueva configuración. |

## 7. Dunnage

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El producto del cliente se daña en tránsito o almacenamiento por falta de protección/fijación adecuada dentro del empaque retornable. |
| **Qué recibe CHACONTAINER** | Especificación del producto a proteger (dimensiones, fragilidad, orientación) y del activo que lo contendrá. |
| **Actividad que ejecuta** | Diseña e instala el sistema de protección interna (espuma, separadores, fijaciones) dentro del activo retornable. |
| **Evidencia que genera** | Diseño de dunnage aprobado, prototipo validado, orden de fabricación/instalación, evidencia de prueba (si aplica prueba de transporte). |
| **Indicador que afecta** | % de reclamos por daño en tránsito del producto del cliente (no del activo), tiempo de diseño, costo por unidad de dunnage. |
| **Información que produce** | Relación producto-del-cliente ↔ activo ↔ configuración de protección — dato de valor para vender "solución completa" en vez de solo el contenedor. |
| **Servicio que puede venderse después** | Venta o renta del activo configurado con el dunnage ya integrado, reposición periódica del dunnage (es consumible en muchos casos), modificación cuando el producto del cliente cambia. |

## 8. Reetiquetado

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El activo cambia de propietario, contenido o destino y su etiquetación (marca, código de producto, información regulatoria) queda obsoleta o incorrecta. |
| **Qué recibe CHACONTAINER** | Activo con etiquetación previa y la nueva especificación de etiquetado (marca, cliente, contenido, normativa aplicable). |
| **Actividad que ejecuta** | Retira etiquetación previa, aplica la nueva conforme a especificación y normativa, valida legibilidad y adherencia. |
| **Evidencia que genera** | Orden de reetiquetado, evidencia fotográfica del resultado, checklist de cumplimiento normativo si aplica (materiales peligrosos, alimentos, etc.). |
| **Indicador que afecta** | Tiempo de reetiquetado, % de rechazos por incumplimiento normativo, activos habilitados para nuevo cliente/uso. |
| **Información que produce** | Historial de propietarios/contenidos sucesivos del activo — relevante para trazabilidad y para decidir si un activo puede reasignarse entre clientes. |
| **Servicio que puede venderse después** | Identificación QR/RFID (el reetiquetado físico es el momento natural para instalar o reemplazar el identificador digital), liberación e ingreso a inventario disponible. |

## 9. Inspección

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Nadie sabe con certeza la condición real de un activo o de un lote antes de decidir si reingresa, se repara, se da de baja o se factura como dañado a un tercero. |
| **Qué recibe CHACONTAINER** | Activo o lote de activos a evaluar, criterio de aceptación aplicable (por tipo de activo/cliente). |
| **Actividad que ejecuta** | Evalúa el activo contra el checklist de condición, clasifica el resultado y determina el siguiente paso del ciclo. |
| **Evidencia que genera** | Checklist de inspección firmado/registrado, fotografías, dictamen de condición (apto / requiere reparación / baja). |
| **Indicador que afecta** | % de activos por categoría de condición, tasa de daño por cliente/ruta/proveedor, tiempo de inspección por activo. |
| **Información que produce** | Es el punto donde se genera el dato de **Capa 4 · Estado** con mayor confiabilidad — alimenta directamente el modelo de gobernanza (¿cuántos están dañados?, ¿cuál es su condición?). |
| **Servicio que puede venderse después** | El servicio que corresponda según el dictamen (lavado, reparación, reacondicionamiento, scrap), y — si la inspección se hace de forma recurrente para un cliente — un contrato de inspección periódica (Nivel 2 de la escalera comercial). |

## 10. Clasificación

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Un lote de activos heterogéneo (mezclado en tipo, condición o propietario) no puede procesarse ni facturarse como unidad; hay que separarlo antes de decidir qué hacer con cada parte. |
| **Qué recibe CHACONTAINER** | Lote de activos sin clasificar, típicamente proveniente de una recuperación masiva, cierre de operación de un cliente o inventario físico. |
| **Actividad que ejecuta** | Separa el lote por tipo, condición, propietario y destino (reutilizable, reparable, scrap), y documenta el conteo resultante por categoría. |
| **Evidencia que genera** | Reporte de clasificación con conteo por categoría, evidencia fotográfica del lote, acta de conformidad con el cliente si aplica. |
| **Indicador que afecta** | Tiempo de clasificación por volumen de lote, % del lote recuperable vs. scrap, precisión de la clasificación (validada contra inspección posterior). |
| **Información que produce** | Primer conteo real de un lote que antes solo existía como estimado — insumo directo para inventarios físicos y para decisiones de recuperación de valor. |
| **Servicio que puede venderse después** | Inventarios físicos (si la clasificación revela que el cliente no tiene control de su parque), inspección detallada por categoría, compra de scrap para la fracción no recuperable. |

## 11. Inventarios físicos

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El cliente (o CHACONTAINER mismo) no sabe cuántos activos tiene realmente, dónde están ni en qué condición, más allá de lo que dice el registro contable. |
| **Qué recibe CHACONTAINER** | Alcance del inventario (ubicación, familia de activos, ventana de tiempo) y acceso al sitio donde se hará el conteo. |
| **Actividad que ejecuta** | Cuenta físicamente los activos, los identifica (si ya tienen QR/RFID) o registra su ausencia de identificación, y contrasta contra el inventario contable/esperado. |
| **Evidencia que genera** | Reporte de inventario físico vs. contable, listado de faltantes, listado de activos sin identificación detectados, evidencia fotográfica del conteo. |
| **Indicador que afecta** | % de discrepancia físico-contable, número de faltantes detectados, % de activos sin identificar. |
| **Información que produce** | La primera fotografía real y auditable del parque de activos de un cliente — es el disparador más directo hacia Packaging Systems descrito en la [ETAPA 8](./README.md#etapa-8--modelo-comercial-solutions--systems): "se descubren faltantes → se necesita trazabilidad". |
| **Servicio que puede venderse después** | Identificación QR/RFID de todo el parque, implementación de trazabilidad continua (Systems), recuperación de los activos localizados fuera de sitio. |

## 12. Recuperación

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Hay activos retornables fuera de circulación (en sitio del cliente, de un proveedor, detenidos en tránsito) que no están generando valor ni rotando. |
| **Qué recibe CHACONTAINER** | Ubicación estimada o confirmada de los activos a recuperar, y la autorización del custodio/propietario para retirarlos. |
| **Actividad que ejecuta** | Localiza, coordina logística inversa, retira los activos y los reingresa al flujo (a inspección/lavado) o a evaluación de scrap. |
| **Evidencia que genera** | Orden de recuperación, acta de retiro firmada por el custodio, registro de condición al momento de la recuperación. |
| **Indicador que afecta** | % de activos recuperados sobre activos "perdidos"/detenidos identificados, costo de recuperación por activo, tiempo detenido antes de recuperar. |
| **Información que produce** | Confirma o corrige la Capa 2 (ubicación) y Capa 3 (custodia) del activo, y genera el dato de "dónde se producen las pérdidas" y "qué proveedor retiene activos" ([ETAPA 4](./README.md#etapa-4--modelo-de-gobernanza)). |
| **Servicio que puede venderse después** | Logística inversa recurrente (Systems), lavado/inspección del activo recuperado, contrato de gobernanza si el patrón de pérdida es sistemático con un cliente/proveedor específico. |

## 13. Compra de scrap

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El cliente tiene activos retornables que ya no puede o no quiere seguir usando (dañados sin reparación viable, obsoletos) y necesita liberarse de ellos generando algún retorno económico. |
| **Qué recibe CHACONTAINER** | Lote de activos declarados como scrap por el cliente, con o sin evaluación previa de CHACONTAINER. |
| **Actividad que ejecuta** | Evalúa el lote, cotiza su valor de scrap, lo adquiere y lo retira. |
| **Evidencia que genera** | Cotización de compra, acta de compraventa/retiro, registro de peso o unidades adquiridas. |
| **Indicador que afecta** | Volumen de scrap comprado, margen entre precio de compra y valor de recuperación posterior, tasa de conversión de lotes evaluados a comprados. |
| **Información que produce** | Confirma la baja definitiva del activo en el registro maestro (cierre del ciclo de vida) y da visibilidad de cuánto valor "sale" del parque administrado. |
| **Servicio que puede venderse después** | Recuperación de valor sobre el mismo lote (separar piezas/materiales aprovechables antes de disposición final), y — indirectamente — venta de activos nuevos que reemplacen lo dado de baja. |

## 14. Disposición final

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Un activo (o lo que queda de él tras recuperación de valor) no tiene ningún uso posible y debe salir del sistema de forma responsable y trazable, no solo "desaparecer". |
| **Qué recibe CHACONTAINER** | Activo o remanente de material sin valor de reuso ni de scrap recuperable. |
| **Actividad que ejecuta** | Gestiona la disposición conforme a normativa ambiental aplicable (reciclaje, disposición controlada), con el proveedor autorizado que corresponda `[VALIDAR proveedor(es) de disposición homologados]`. |
| **Evidencia que genera** | Certificado o comprobante de disposición final, registro del proveedor de disposición utilizado, baja definitiva en el sistema. |
| **Indicador que afecta** | % de activos con disposición certificada (vs. sin evidencia), costo de disposición por unidad/tipo. |
| **Información que produce** | Cierre auditable del ciclo de vida completo del activo — dato relevante para reportes de sostenibilidad/ESG frente a clientes industriales. |
| **Servicio que puede venderse después** | Ninguno sobre ese activo específico (es el fin del ciclo); a nivel cuenta, el certificado de disposición sirve para vender el siguiente ciclo de venta/renta como reemplazo responsable. |

## 15. Recuperación de valor

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Un activo dado de baja o comprado como scrap todavía contiene componentes, materiales o piezas con valor de reventa o reuso que se perderían si se dispusiera directamente. |
| **Qué recibe CHACONTAINER** | Activos o lotes ya clasificados como no reutilizables en su forma original (scrap comprado, bajas, remanentes de reacondicionamiento). |
| **Actividad que ejecuta** | Desarma, separa y evalúa componentes/materiales aprovechables (piezas, metal, plástico, accesorios reutilizables) para reventa o reuso interno. |
| **Evidencia que genera** | Reporte de componentes recuperados, registro de destino de cada componente (reventa, reuso interno, disposición del remanente). |
| **Indicador que afecta** | % de valor recuperado sobre el valor de scrap adquirido, ingresos por venta de componentes, volumen que efectivamente evita disposición final. |
| **Información que produce** | Tasa real de recuperación de valor por tipo de activo — insumo para decidir, en la ETAPA 4, "qué activos conviene reparar" vs. "qué activos deben darse de baja" con datos económicos reales en vez de criterio informal. |
| **Servicio que puede venderse después** | Venta de componentes/piezas recuperadas como repuesto para reparación de otros activos del mismo parque, disposición final del remanente sin valor. |

---

## Lectura cruzada: de dónde entra cada servicio a Systems

| Servicio Solutions | Dato que expone primero | Puerta de entrada a Systems (ver [ETAPA 8](./README.md#etapa-8--modelo-comercial-solutions--systems)) |
|---|---|---|
| Lavado | Activos sin identificación | Identificación QR/RFID |
| Inventarios físicos | Faltantes reales | Trazabilidad |
| Recuperación | Activos detenidos/fuera de sitio | Logística inversa + alertas |
| Inspección | Condición real del parque | Gobernanza / reglas de aceptación |
| Reparación | Costo real de mantener el activo vivo | Decisión reparar-vs-baja basada en datos |
| Clasificación | Composición real de un lote | Inventarios físicos + recuperación de valor |

Este cruce es el argumento comercial central de la [ETAPA 11](./README.md#etapa-11--propuesta-comercial): cualquiera de estos 15 servicios, ejecutado con disciplina de datos, puede convertirse en el punto de entrada hacia una relación de gobernanza recurrente.

---

## Próximo paso sugerido

Con este catálogo cerrado, el paso 2 de la FASE I es el catálogo definitivo de Packaging Systems (mismo formato de 7 campos, aplicado a: diagnóstico, identificación, inventario digital, trazabilidad, custodios, logística inversa, incidencias, reglas, alertas, indicadores, analítica, gobernanza, CHACONTAINER OS).

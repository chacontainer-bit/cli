# Catálogo maestro · CHACONTAINER Packaging Systems

**Fase:** I · Fundación — paso 2 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

Este catálogo cubre las 16 capacidades de Packaging Systems definidas en el
[Proyecto Maestro](./README.md#etapa-0--fundamento-del-modelo), documentadas
con los mismos 7 campos usados en el [catálogo de Solutions](./catalogo-solutions.md).
La diferencia de fondo entre ambos catálogos: en Solutions el "cliente" casi
siempre paga por una intervención física puntual; en Systems, lo que se vende
es la garantía de que esa información sigue viva después de la intervención
— por eso casi todas las entradas de este catálogo terminan alimentando la
misma capa: [Capa 5 · Gobierno](./README.md#etapa-3--modelo-de-packaging-systems).

Los campos marcados `[VALIDAR]` dependen de una decisión técnica o comercial
pendiente — no inventar el dato, confirmarlo antes de operar sobre él.

---

## Índice

1. [Diagnóstico de activos](#1-diagnóstico-de-activos)
2. [Registro](#2-registro)
3. [Identificación](#3-identificación)
4. [QR](#4-qr)
5. [RFID](#5-rfid)
6. [Inventario digital](#6-inventario-digital)
7. [Trazabilidad](#7-trazabilidad)
8. [Gestión de custodios](#8-gestión-de-custodios)
9. [Logística inversa](#9-logística-inversa)
10. [Gestión de incidencias](#10-gestión-de-incidencias)
11. [Reglas operativas](#11-reglas-operativas)
12. [Alertas](#12-alertas)
13. [Indicadores](#13-indicadores)
14. [Analítica](#14-analítica)
15. [Gobernanza](#15-gobernanza)
16. [CHACONTAINER OS](#16-chacontainer-os)

---

## 1. Diagnóstico de activos

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El cliente no sabe si su forma actual de administrar el empaque retornable (o la ausencia total de administración) le está costando dinero, y no tiene con qué argumentar una intervención de Systems. |
| **Qué recibe CHACONTAINER** | Acceso al proceso actual del cliente: cómo compra, usa, retorna y da seguimiento (o no) a su empaque retornable. |
| **Actividad que ejecuta** | Levanta el estado actual (ver [SOP-03](../chacontainer/docs/sop/SOP-03-diagnostico-consultivo.md)): parque estimado, puntos ciegos, pérdidas percibidas, costo de no saber. |
| **Evidencia que genera** | Informe de diagnóstico con hallazgos y estimación de impacto económico, línea base de los KPI que se van a comparar después. |
| **Indicador que afecta** | Ninguno propio — genera la **línea base** contra la que se medirán todos los KPI de la [ETAPA 6](./README.md#etapa-6--indicadores-del-sistema) una vez implementado el piloto. |
| **Información que produce** | Primer inventario *estimado* del parque del cliente (a falta de uno físico), y el mapa de las 5 capas ([Activo, Ubicación, Custodia, Estado, Gobierno](./README.md#etapa-3--modelo-de-packaging-systems)) aplicado a su operación real. |
| **Servicio que puede venderse después** | Inventarios físicos (Solutions) para reemplazar el estimado por un conteo real, y el piloto de Packaging Systems descrito en la [ETAPA 7](./README.md#etapa-7--piloto-packaging-systems). |

## 2. Registro

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Un activo que existe físicamente pero no existe en ningún sistema no puede gestionarse, solo puede buscarse. |
| **Qué recibe CHACONTAINER** | Datos del activo (tipo, modelo, medida, propietario, número de parte, condición, valor — Capa 1) capturados en cualquier punto de contacto: venta, lavado, inventario físico, recuperación. |
| **Actividad que ejecuta** | Da de alta el activo en el registro maestro (Módulo 1 de [CHACONTAINER OS](#16-chacontainer-os)) con un ID único, evitando duplicados. |
| **Evidencia que genera** | Ficha del activo en el registro maestro, con fecha y origen del alta (venta, hallazgo en inventario físico, recuperación, etc.). |
| **Indicador que afecta** | Cobertura de registro: % del parque real que efectivamente existe en el sistema (base de "disponibilidad real", no solo contable). |
| **Información que produce** | El registro maestro mismo — la tabla de la que dependen todas las demás capacidades de este catálogo. |
| **Servicio que puede venderse después** | Identificación física (QR/RFID) para que el ID digital tenga una contraparte física escaneable, inventario digital. |

## 3. Identificación

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Aun con el activo registrado en el sistema, nadie en campo puede confirmar en segundos *cuál* activo físico corresponde a *qué* registro. |
| **Qué recibe CHACONTAINER** | El activo físico ya dado de alta en el registro, sin marcaje digital todavía. |
| **Actividad que ejecuta** | Decide y aplica el método de identificación (QR o RFID, ver [#4](#4-qr) y [#5](#5-rfid)) según el caso de uso, y lo vincula al ID del registro maestro. |
| **Evidencia que genera** | Registro del identificador asignado (código QR o tag RFID) vinculado al ID del activo, evidencia fotográfica de la instalación. |
| **Indicador que afecta** | % del parque identificado sobre el parque registrado — brecha crítica que suele descubrirse en inventarios físicos o lavado. |
| **Información que produce** | El puente entre el mundo físico y el registro digital — sin esto, ninguna de las capacidades siguientes (trazabilidad, alertas, indicadores) tiene datos reales que procesar. |
| **Servicio que puede venderse después** | Inventario digital, trazabilidad continua — es, junto con inventarios físicos, la puerta de entrada más frecuente descrita en la [ETAPA 8](./README.md#etapa-8--modelo-comercial-solutions--systems). |

## 4. QR

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Se necesita un método de identificación de bajo costo y despliegue inmediato, con lectura desde cualquier smartphone sin hardware adicional. |
| **Qué recibe CHACONTAINER** | Activo a identificar y el payload que debe codificarse (ID del activo, y opcionalmente tenant/tipo — ver formato en [architecture.md](../chacontainer/docs/architecture.md#qr-code-system)). |
| **Actividad que ejecuta** | Genera el código QR, lo imprime/aplica sobre el activo (etiqueta resistente a lavado e intemperie `[VALIDAR material y proveedor]`) y verifica que escanea correctamente. |
| **Evidencia que genera** | Código QR físico instalado y verificado, registro del evento de escaneo inicial en el sistema. |
| **Indicador que afecta** | Costo de identificación por activo (más bajo que RFID), tiempo de despliegue por lote. |
| **Información que produce** | Cada escaneo posterior genera un evento de trazabilidad (quién, cuándo, dónde) — ver flujo de escaneo en [architecture.md](../chacontainer/docs/architecture.md#scan-flow). |
| **Servicio que puede venderse después** | Trazabilidad basada en escaneo manual (más barata que RFID, requiere que alguien escanee en cada punto de control), inventario digital. |

## 5. RFID

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El caso de uso exige lectura sin línea de vista ni intervención manual (portales de entrada/salida, lectura masiva de un lote, activos que rotan a alta frecuencia donde escanear uno a uno no es viable). |
| **Qué recibe CHACONTAINER** | Activo a identificar, especificación del punto de lectura (portal fijo, lector móvil) y volumen/frecuencia de paso esperado. |
| **Actividad que ejecuta** | Instala el tag RFID en el activo y, si aplica, despliega la infraestructura de lectura (portal, antenas, lectores) en los puntos de control acordados `[VALIDAR alcance: solo tag vs. tag + infraestructura de lectura]`. |
| **Evidencia que genera** | Tag RFID instalado y verificado, registro de la infraestructura de lectura desplegada (si aplica), prueba de lectura en el punto de control. |
| **Indicador que afecta** | Costo de identificación por activo (mayor que QR, justificado por volumen/frecuencia), cobertura de puntos de control con lectura automática. |
| **Información que produce** | Eventos de movimiento automáticos sin depender de que una persona escanee — habilita trazabilidad casi en tiempo real en los puntos con lector instalado. |
| **Servicio que puede venderse después** | Trazabilidad automática de alto volumen, contrato de gobernanza (Nivel 5) donde el costo de la infraestructura se justifica por el fee recurrente de administración. |

## 6. Inventario digital

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El inventario físico es una fotografía puntual que empieza a desactualizarse apenas termina; el cliente necesita saber "cuánto tengo disponible" en cualquier momento, no solo el día del conteo. |
| **Qué recibe CHACONTAINER** | El registro maestro de activos identificados, más los eventos de movimiento/estado que los van actualizando. |
| **Actividad que ejecuta** | Mantiene el saldo de inventario por ubicación, tipo y estado actualizado en el sistema (Módulo 9 de [CHACONTAINER OS](#16-chacontainer-os)) a partir de cada evento registrado. |
| **Evidencia que genera** | Reporte de inventario en cualquier corte de fecha, historial de saldos por ubicación. |
| **Indicador que afecta** | Disponibilidad y disponibilidad real ([ETAPA 6](./README.md#etapa-6--indicadores-del-sistema)): inventario físicamente utilizable, no solo contable. |
| **Información que produce** | La respuesta permanente (no puntual) a "¿cuántos activos disponibles tengo y dónde?" — el dato que reemplaza al inventario físico como evento aislado. |
| **Servicio que puede venderse después** | Alertas de nivel mínimo de inventario, analítica de rotación por ubicación/cliente. |

## 7. Trazabilidad

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Se sabe que un activo existe pero no su historia: dónde ha estado, quién lo ha tenido, cuánto tiempo lleva en cada punto — información esencial para resolver disputas, detectar cuellos de botella y cumplir requisitos de clientes regulados. |
| **Qué recibe CHACONTAINER** | Eventos de escaneo/lectura (QR o RFID) y de cambio de estado o custodia generados por cualquier otro proceso (lavado, renta, movimiento, retorno). |
| **Actividad que ejecuta** | Registra cada evento con marca de tiempo, ubicación y actor, y construye la línea de tiempo completa del activo (Módulo 4 de CHACONTAINER OS). |
| **Evidencia que genera** | Historial completo por activo, consultable en cualquier momento; exportable como evidencia ante un cliente o auditoría. |
| **Indicador que afecta** | Tiempo de ciclo, tiempo detenido, cumplimiento de retorno — todos requieren la línea de tiempo que produce trazabilidad. |
| **Información que produce** | La serie temporal de movimientos y estados por activo — la materia prima de la que se derivan indicadores y analítica. |
| **Servicio que puede venderse después** | Alertas basadas en tiempo (activo detenido más de X días), analítica de patrones de movimiento, logística inversa dirigida por los propios datos de trazabilidad. |

## 8. Gestión de custodios

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Cuando un activo se pierde o se daña, nadie puede decir con certeza quién lo tenía a cargo en ese momento — sin custodio identificado no hay a quién exigir responsabilidad ni con quién coordinar el retorno. |
| **Qué recibe CHACONTAINER** | Eventos de asignación y transferencia de custodia (Capa 3): cliente, proveedor, operador, transportista, planta, área o usuario que recibe el activo. |
| **Actividad que ejecuta** | Registra cada transferencia de custodia y mantiene el custodio actual de cada activo consultable en todo momento. |
| **Evidencia que genera** | Historial de custodios por activo, con fecha de cada transferencia; acta de transferencia cuando el proceso lo requiere (ver [SOP-12](./README.md#etapa-2--sop-del-ciclo-completo-del-activo) planificado). |
| **Indicador que afecta** | Cumplimiento de retorno por custodio, tasa de pérdida por custodio — insumo directo para responder "qué proveedor retiene activos" ([ETAPA 4](./README.md#etapa-4--modelo-de-gobernanza)). |
| **Información que produce** | Responsabilidad trazable activo por activo — condición necesaria para cualquier conversación de gobernanza con un cliente o proveedor. |
| **Servicio que puede venderse después** | Alertas de custodia (tiempo máximo fuera excedido por custodio específico), logística inversa dirigida a los custodios con mayor retención. |

## 9. Logística inversa

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Los activos retornables no regresan solos; sin un proceso activo de recolección, se acumulan indefinidamente en sitio del cliente o del proveedor y dejan de rotar. |
| **Qué recibe CHACONTAINER** | Lista de activos a recolectar (por vencimiento de renta, por alerta de tiempo detenido, o por cierre de operación de un cliente) y su ubicación según trazabilidad. |
| **Actividad que ejecuta** | Planea y coordina la ruta de recolección, retira los activos y los reingresa al flujo físico (inspección/lavado) actualizando su estado y ubicación en el sistema. |
| **Evidencia que genera** | Ruta de recolección ejecutada, acta de retiro por punto, actualización de estado/ubicación de cada activo recolectado. |
| **Indicador que afecta** | Tiempo de ciclo, % de activos recolectados sobre pendientes de recolección, costo de recolección por ruta/activo. |
| **Información que produce** | Cierra el loop entre "dónde debería estar el activo" (según reglas) y "dónde está" (según trazabilidad) — convierte una alerta en una acción física ejecutada. |
| **Servicio que puede venderse después** | Recuperación (Solutions) cuando la recolección detecta activos dañados o extraviados, contrato de logística inversa recurrente (Nivel 4 de la escalera comercial). |

## 10. Gestión de incidencias

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Daños, retrasos, faltantes y pérdidas ocurren de forma dispersa (por WhatsApp, llamada, correo) y sin un registro único no se puede saber si un problema es aislado o un patrón. |
| **Qué recibe CHACONTAINER** | Reporte de una incidencia (daño, retraso, faltante, pérdida) desde cualquier punto: cliente, operador de campo, proceso interno de inspección. |
| **Actividad que ejecuta** | Registra la incidencia con tipo, activo(s) involucrado(s), fecha y responsable; da seguimiento hasta su resolución (Módulo 8 de CHACONTAINER OS). |
| **Evidencia que genera** | Ticket de incidencia con estado (abierta/en proceso/resuelta), evidencia asociada (fotos, comunicaciones), fecha de resolución. |
| **Indicador que afecta** | Pérdida, daño, número de incidencias por tipo/cliente/ruta, tiempo promedio de resolución. |
| **Información que produce** | El registro estructurado de todo lo que "sale mal" en el ciclo — sin esto, gobernanza (#15) no tiene con qué priorizar. |
| **Servicio que puede venderse después** | Reglas operativas ajustadas para prevenir la incidencia recurrente, alertas tempranas basadas en los mismos patrones detectados. |

## 11. Reglas operativas

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Sin reglas explícitas, "qué debería pasar" con cada activo depende del criterio informal de quien esté a cargo ese día — no es consistente ni escalable. |
| **Qué recibe CHACONTAINER** | Parámetros de negocio acordados con el cliente o definidos internamente: tiempo máximo fuera, punto de retorno, nivel mínimo de inventario, condición permitida (Capa 5). |
| **Actividad que ejecuta** | Configura estas reglas en el sistema por tipo de activo, cliente o ruta, de forma que el sistema pueda evaluarlas automáticamente contra el estado real. |
| **Evidencia que genera** | Catálogo de reglas vigentes por cliente/tipo de activo, historial de cambios a las reglas. |
| **Indicador que afecta** | Ninguno directamente — las reglas son el criterio contra el que se miden cumplimiento de retorno, tiempo detenido y disponibilidad. |
| **Información que produce** | El estándar contra el cual el sistema puede clasificar automáticamente cada activo como "en regla" o "en excepción". |
| **Servicio que puede venderse después** | Alertas (la ejecución automática de estas reglas), gobernanza (el ajuste continuo de las reglas según lo que muestran los indicadores). |

## 12. Alertas

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Nadie revisa un dashboard todo el día; si una excepción (activo detenido, faltante, fuera de tiempo) no se comunica activamente, se descubre demasiado tarde. |
| **Qué recibe CHACONTAINER** | El cruce continuo entre el estado real de cada activo (trazabilidad) y las reglas operativas (#11) vigentes para él. |
| **Actividad que ejecuta** | Detecta automáticamente cuándo un activo incumple una regla y notifica al responsable definido (Módulo 10 de CHACONTAINER OS). |
| **Evidencia que genera** | Registro de la alerta emitida, a quién se notificó y cuándo, y si fue atendida. |
| **Indicador que afecta** | Tiempo detenido (se reduce si la alerta se atiende rápido), cumplimiento de retorno. |
| **Información que produce** | El primer momento en que una excepción se vuelve visible para alguien que puede actuar — convierte datos pasivos en acción. |
| **Servicio que puede venderse después** | Logística inversa (la alerta dispara la recolección), escalamiento a gobernanza si la misma alerta se repite de forma sistemática con un cliente o custodio. |

## 13. Indicadores

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | El cliente pregunta "¿está funcionando esto?" y sin métricas consolidadas la respuesta es anecdótica, no demostrable. |
| **Qué recibe CHACONTAINER** | Los eventos crudos generados por registro, identificación, trazabilidad, custodia e incidencias. |
| **Actividad que ejecuta** | Calcula y presenta los KPI definidos en la [ETAPA 6](./README.md#etapa-6--indicadores-del-sistema) (disponibilidad, utilización, tiempo de ciclo, cumplimiento de retorno, pérdida, daño, tiempo detenido, costo por ciclo, rotación, disponibilidad real). |
| **Evidencia que genera** | Dashboard de indicadores (Módulo 11 de CHACONTAINER OS), reportes periódicos exportables. |
| **Indicador que afecta** | Es el propio producto: la capacidad de calcular todos los KPI del sistema. |
| **Información que produce** | La respuesta cuantitativa a las preguntas de gobernanza de la [ETAPA 4](./README.md#etapa-4--modelo-de-gobernanza). |
| **Servicio que puede venderse después** | Analítica (comparaciones, tendencias y proyecciones sobre estos mismos indicadores), reporte ejecutivo periódico como parte de un contrato de gobernanza. |

## 14. Analítica

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Saber el valor actual de un KPI no dice si va mejorando, empeorando, ni dónde intervenir primero entre varias plantas, rutas o clientes. |
| **Qué recibe CHACONTAINER** | La serie histórica de indicadores y eventos ya calculados (#13), segmentable por planta, cliente, ruta, familia de activo. |
| **Actividad que ejecuta** | Analiza tendencias, compara segmentos, identifica outliers (la planta con más pérdida, el proveedor que más retiene) y prioriza dónde intervenir. |
| **Evidencia que genera** | Reporte analítico o comparativo, recomendación priorizada de dónde actuar primero. |
| **Indicador que afecta** | No genera un KPI nuevo — mejora la calidad de la decisión sobre los KPI existentes. |
| **Información que produce** | Priorización basada en datos de "qué activos conviene reparar", "qué planta tiene exceso/déficit" ([ETAPA 4](./README.md#etapa-4--modelo-de-gobernanza)). |
| **Servicio que puede venderse después** | Gobernanza (la analítica es el insumo con el que se ajustan las reglas operativas de forma continua). |

## 15. Gobernanza

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Tener datos, reglas, alertas e indicadores no sirve de nada si nadie los usa para decidir y ajustar el sistema de forma continua — sin gobernanza, Systems se queda en "tracking pasivo". |
| **Qué recibe CHACONTAINER** | Todo lo anterior (#1–#14) consolidado, más la autoridad acordada con el cliente para tomar o recomendar decisiones sobre su parque de activos. |
| **Actividad que ejecuta** | Revisa indicadores y analítica de forma periódica, ajusta reglas operativas, escala incidencias sistemáticas, y responde por el desempeño del sistema completo frente al cliente. |
| **Evidencia que genera** | Reporte de gobernanza periódico (mensual `[VALIDAR periodicidad]`), bitácora de ajustes a reglas y decisiones tomadas. |
| **Indicador que afecta** | Todos — es la función que se responsabiliza de que los KPI de la ETAPA 6 se muevan en la dirección correcta. |
| **Información que produce** | El historial de decisiones tomadas sobre el sistema — evidencia de que alguien está administrando el ciclo, no solo midiéndolo. |
| **Servicio que puede venderse después** | Nivel 6 de la [escalera comercial](./README.md#etapa-9--escalera-comercial): el cliente adopta CHACONTAINER OS como infraestructura propia del sistema. |

## 16. CHACONTAINER OS

| Campo | Detalle |
|---|---|
| **Problema que resuelve** | Todas las capacidades anteriores existen como proceso, pero sin una plataforma que las conecte, cada una vive en una hoja de cálculo o en la memoria de alguien — no escala más allá de un puñado de activos. |
| **Qué recibe CHACONTAINER** | Los procesos ya definidos en las etapas 0-4 de este proyecto maestro — el software digitaliza lo que ya fue diseñado, no al revés (ver [ETAPA 5](./README.md#etapa-5--chacontainer-os)). |
| **Actividad que ejecuta** | Opera y mantiene la plataforma (los 12 módulos del MVP: activos, identificación, ubicaciones, movimientos, custodia, condición, mantenimiento, incidencias, inventario, alertas, dashboard, evidencia). Implementación técnica base ya existe en [`chacontainer/cmd/server`](../chacontainer/cmd/server), [`chacontainer/internal/`](../chacontainer/internal/) y [`chacontainer/web/`](../chacontainer/web/) del SaaS actual. |
| **Evidencia que genera** | El sistema mismo: cada registro, evento, alerta y reporte que producen las 15 capacidades anteriores queda persistido aquí. |
| **Indicador que afecta** | Disponibilidad de la plataforma, cobertura de módulos activos por cliente. |
| **Información que produce** | La fuente única de verdad del ciclo de vida de cada activo — el objetivo final descrito en el [Norte del proyecto](./README.md#norte-del-proyecto). |
| **Servicio que puede venderse después** | Licenciamiento (Nivel 6, modelo de ingresos por software de la [ETAPA 10](./README.md#etapa-10--modelo-de-ingresos)): por activo/mes, por usuario, por planta o por módulo. |

---

## Lectura cruzada: qué capacidad de Systems resuelve cada pregunta de gobernanza

| Pregunta de la [ETAPA 4](./README.md#etapa-4--modelo-de-gobernanza) | Capacidad que la responde |
|---|---|
| ¿Cuántos activos existen? | Registro + Inventario digital |
| ¿Cuántos están realmente disponibles? | Inventario digital (disponibilidad real) |
| ¿Dónde están? | Trazabilidad |
| ¿Quién los tiene? | Gestión de custodios |
| ¿Cuánto tiempo llevan ahí? | Trazabilidad + Reglas operativas |
| ¿Cuál es su condición? | Registro (Capa 4 · Estado, alimentada por inspección de Solutions) |
| ¿Cuándo deberían regresar? | Reglas operativas |
| ¿Cuántos están detenidos? | Alertas |
| ¿Cuántos están dañados? | Gestión de incidencias |
| ¿Cuántos faltan? | Inventario digital vs. registro esperado |
| ¿Cuánto cuesta cada ciclo? | Indicadores |
| ¿Dónde se producen las pérdidas? | Analítica |
| ¿Qué proveedor retiene activos? | Gestión de custodios + Analítica |
| ¿Qué planta tiene exceso/déficit? | Analítica |
| ¿Qué activos conviene reparar? | Analítica (cruzada con reparación de Solutions) |
| ¿Qué activos deben darse de baja? | Analítica + Gobernanza |

Este cruce confirma que ninguna de las 16 capacidades es opcional de forma aislada: **Gobernanza (#15)** es la única que no tiene dato propio — depende por completo de que las otras 15 estén funcionando.

---

## Próximo paso sugerido

Con los dos catálogos cerrados (Solutions + Systems), el paso 3 de la FASE I es el **ciclo de vida del activo**: formalizar el diagrama de la [ETAPA 2](./README.md#etapa-2--sop-del-ciclo-completo-del-activo) como el flujo maestro único que conecta cada servicio de Solutions y cada capacidad de Systems documentados en estos dos catálogos, dejando explícito en qué punto del ciclo interviene cada uno.

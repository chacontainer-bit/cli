# Plan Estratégico de Negocio — CHACONTAINER
## Enfoque sector automotriz e industrial (OEM · Tier 1 · Tier 2 · manufactura avanzada · nearshoring)

**Documento base para dirección general.** Construido ejecutando la
secuencia recomendada del [banco de prompts fuente](./prompts-fuente.md):
18 → 1 → 2 → 3 → 4 → 12 → 6 → 7 → 8 → 20.

> **Independiente de [`chacontainer/`](../chacontainer/) y de
> [`chacontainer-master-project/`](../chacontainer-master-project/).**
> Este documento añade la capa que los anteriores no cubrían: la
> especificidad del sector automotriz (OEM, Tier 1, Tier 2, nearshoring) y
> el formato de plan estratégico para dirección/inversión. Donde el
> contenido coincide con lo ya construido en `chacontainer-master-project/`
> (Packaging Solutions, Packaging Systems, modelo de gobernanza, pricing,
> KPIs), este documento lo referencia en vez de reescribirlo desde cero,
> para no generar dos fuentes de verdad divergentes.
>
> **Regla seguida en todo el documento:** no se inventan cifras de mercado,
> ventas, costos ni valuación. Donde hace falta un número real, el texto
> dice explícitamente **`Dato requerido`** y explica qué dato es y por qué
> importa. Donde se necesita una fórmula, se da la fórmula, no un valor de
> ejemplo disfrazado de dato real.

---

## 1. Portada

| Campo | Valor |
|---|---|
| Documento | Plan Estratégico de Negocio — CHACONTAINER |
| Enfoque | Sector automotriz e industrial (OEM, Tier 1, Tier 2, manufactura avanzada, nearshoring) |
| Alcance | Empaques retornables, contenedores colapsables, racks metálicos, servicios asociados y sistema de gobernanza de activos retornables |
| Uso previsto | Documento base para dirección general; convertible a presentación o propuesta para inversionistas/socios |
| Estado | Borrador v0.1 — construido sobre el marco del proyecto maestro CHACONTAINER; pendiente de validación con datos reales de operación |
| Fecha | `Dato requerido` — fecha de emisión formal |
| Responsable | `Dato requerido` — quién firma el documento hacia dirección/inversionistas |

---

## 2. Resumen ejecutivo

CHACONTAINER opera hoy como proveedor transaccional de empaque retornable
industrial — compra, venta, renta, reparación, modificación, limpieza,
compra de scrap y logística inversa — con foco natural en clientes del
sector automotriz e industrial. La oportunidad estratégica es dejar de
competir solo por transacción física y convertirse en el **sistema de
control, operación y rentabilidad de activos retornables** para OEM,
Tier 1, Tier 2 y manufactura avanzada — un movimiento que el resto de este
documento llama **Packaging Solutions → Packaging Systems** (detalle en
§9-10, y ya desarrollado con profundidad operativa en
[`chacontainer-master-project/`](../chacontainer-master-project/)).

El argumento comercial central no es "vendemos contenedores mejores": es
que **la mayoría de los clientes de este segmento no saben cuántos activos
tienen realmente disponibles hoy, ni cuánto les cuesta no saberlo** — y ese
costo, una vez cuantificado, suele ser mayor a lo que el cliente percibe
(ver §5-6). CHACONTAINER ya tiene la puerta de entrada operativa (venta,
renta, reparación, limpieza) para descubrir ese costo en cada intervención;
lo que falta es convertir esa puerta de entrada en un sistema comercial y
de datos sistemático, no anecdótico.

Este documento no sustituye el trabajo de detalle ya hecho en
`chacontainer-master-project/` (catálogos de servicio, SOP, modelo de
datos, modelo de gobernanza, pricing) — lo eleva a nivel de dirección y lo
aterriza específicamente al sector automotriz, donde aplican dinámicas
propias: ciclos de producción JIT/JIS, penalización severa por paro de
línea, contratos de suministro multianuales, y el fenómeno de nearshoring
que está expandiendo la base de plantas Tier 1/Tier 2 en México.

**Lo que este documento no puede resolver por sí solo:** toda cifra de
mercado, ingreso actual, costo real, tamaño de equipo o cartera de clientes
existente requiere datos reales de CHACONTAINER que no están disponibles en
este entorno de trabajo. Cada uno de esos puntos queda marcado
`Dato requerido` con la pregunta exacta que hay que responder.

---

## 3. Contexto del negocio

CHACONTAINER participa en el mercado de empaque retornable industrial con
un catálogo de servicios ya documentado en detalle
([catálogo Solutions](../chacontainer-master-project/catalogo-solutions.md)):
venta, renta, lavado, reparación, reacondicionamiento, modificación,
dunnage, reetiquetado, inspección, clasificación, inventarios físicos,
recuperación, compra de scrap, disposición final y recuperación de valor.

El contexto que este documento añade es el **sectorial**: el cliente
automotriz e industrial (OEM, Tier 1, Tier 2) tiene características que no
comparte con un cliente industrial genérico:

- **Sensibilidad extrema al paro de línea.** Un faltante de empaque puede
  detener una línea de producción con costo por hora muy superior al valor
  del empaque mismo.
- **Contratos de suministro multianuales**, que atan al proveedor de
  empaque a la vida del programa vehicular — la relación no es
  transaccional, es de largo plazo por diseño.
- **Multiplicidad de plantas y proveedores** en la cadena OEM → Tier 1 →
  Tier 2, donde el activo retornable cruza custodios constantemente — el
  escenario exacto que el [modelo de custodia](../chacontainer-master-project/modelo-de-datos.md#asset_custody-capa-3--custodia--no-existe-hoy)
  del proyecto maestro fue diseñado para resolver.
- **Nearshoring**: expansión activa de plantas Tier 1/Tier 2 en México
  atrayendo capacidad manufacturera desde Asia y otras regiones — cada
  planta nueva es, potencialmente, tanto una necesidad de empaque desde
  cero como un cliente sin ningún sistema de trazabilidad todavía.

`Dato requerido`: volumen actual de facturación de CHACONTAINER, mezcla de
ingresos por línea de servicio, número y tipo de clientes activos
(OEM/Tier 1/Tier 2/otros), y plantas atendidas hoy. Sin esto, el
diagnóstico del §4 es cualitativo, no cuantitativo.

---

## 4. Diagnóstico estratégico

### 4.1 Situación actual (cualitativa, sujeta a validación)

CHACONTAINER opera como proveedor de servicios físicos de empaque
retornable con un catálogo amplio (15 servicios documentados) pero, hasta
antes del trabajo de `chacontainer-master-project/`, sin un sistema formal
que capture y explote los datos que esos servicios generan. Es una empresa
con capacidad operativa real (lavado, reparación, logística inversa) y con
una oportunidad de sistema de datos todavía no construida.

### 4.2 FODA aplicado

| | Interno | Externo |
|---|---|---|
| **Positivo** | **Fortalezas:** catálogo de servicio ya amplio y documentado a nivel SOP; capacidad operativa física instalada (lavado, reparación, logística); ya existe un SaaS propio (`chacontainer/`) con base de datos multi-tenant, QR y trazabilidad básica funcionando | **Oportunidades:** nearshoring expandiendo la base de clientes Tier 1/Tier 2 en México; ningún competidor tradicional de empaque retornable ofrece gobernanza de datos como servicio (ver §7); clientes automotrices ya acostumbrados a exigir trazabilidad por otras razones (calidad, IATF) — el terreno cultural para pedir trazabilidad de empaque está más maduro que en otros sectores |
| **Negativo** | **Debilidades:** modelo de datos del SaaS actual soporta solo 6 estados de activo, no la granularidad operativa que exige un sistema de gobernanza real (ver [modelo-de-datos.md](../chacontainer-master-project/modelo-de-datos.md#2-brecha-principal-assetsstatus-6-valores-vs-estados-del-activomd-16-estados)); no existe hoy modelo de custodia con historial; pricing y costos reales sin validar (`lista-precios-costos.md` está en `[VALIDAR]` casi en su totalidad) | **Amenazas:** un competidor de empaque tradicional podría copiar el discurso de "gobernanza" sin la ejecución detrás; volatilidad de la industria automotriz (paros de producción, cambios de programa) afecta directamente la demanda de empaque; dependencia de que el cliente automotriz adopte el escaneo disciplinadamente (riesgo de adopción, ver [piloto.md §6](../chacontainer-master-project/piloto.md#6-riesgos)) |

### 4.3 Capacidades operativas actuales

`Dato requerido`: capacidad instalada real (plantas propias, personal de
lavado/reparación, flota de logística inversa, volumen de activos
administrados hoy). El proyecto maestro asume una operación con capacidad
de piloto (200-500 activos, ver [piloto.md](../chacontainer-master-project/piloto.md#2-criterios-de-selección))
pero no conoce la capacidad real instalada de CHACONTAINER.

### 4.4 Nivel de madurez empresarial (escala 1-5)

| Dimensión | Nivel estimado | Justificación | Evidencia necesaria para confirmar |
|---|---|---|---|
| Estrategia | 3 — estructurado | El proyecto maestro ya formaliza tesis, ciclo de vida, modelo de datos y modelo comercial | Validación del fundador sobre si esta estrategia refleja decisiones ya tomadas u objetivos aspiracionales |
| Ventas | `Dato requerido` | No hay visibilidad de pipeline, CRM o proceso comercial real fuera de los SOP comerciales documentados en `chacontainer/docs/sop/` | Volumen de oportunidades activas, tasa de conversión histórica |
| Operaciones | `Dato requerido` | Existen SOP detallados (20 del ciclo del activo) pero no está confirmado que se ejecuten así en campo hoy | Confirmación operativa punto por punto de los SOP |
| Tecnología | 2-3 — básico a estructurado | SaaS propio ya existe con QR, multi-tenant, trazabilidad básica; falta custodia, reglas, alertas, incidencias (ver [arquitectura-os.md](../chacontainer-master-project/arquitectura-os.md#2-los-12-módulos-del-mvp-mapeados)) | — (ya evaluado directamente contra el código) |
| Control de activos | 2 — básico | 6 estados agregados, sin historial de custodia | — (ya evaluado) |
| Documentación | 4 — escalable | Nivel alto por el trabajo ya hecho en el proyecto maestro; el riesgo es que la documentación vaya adelante de la operación real | Contraste documento vs. campo |
| Capacidad para recibir inversión | `Dato requerido` | Depende de gobierno corporativo, estados financieros y trazabilidad de ingresos — ninguno disponible aquí | Estados financieros auditables o al menos internos consistentes |

### 4.5 Brechas entre lo que CHACONTAINER es hoy y lo que necesita ser

La brecha central, ya identificada en detalle en el proyecto maestro, es
**pasar de vender contenedores a vender certeza sobre el parque de
contenedores**. Las tres brechas concretas que la sostienen:

1. **Brecha de datos**: el modelo de estado actual (6 valores) no distingue
   condiciones operativas críticas (sucio, en inspección, liberado) — ver
   [modelo-de-datos.md §2](../chacontainer-master-project/modelo-de-datos.md#2-brecha-principal-assetsstatus-6-valores-vs-estados-del-activomd-16-estados).
2. **Brecha de custodia**: no hay forma de responder "¿quién tiene este
   activo, desde cuándo?" con historial — crítico en cadenas OEM → Tier 1 →
   Tier 2 donde el activo cambia de manos varias veces.
3. **Brecha comercial**: no existe todavía un mecanismo sistemático que
   convierta un servicio de Solutions en una oportunidad de Systems — ver
   [conversion-solutions-systems.md](../chacontainer-master-project/conversion-solutions-systems.md#1-por-qué-la-secuencia-de-la-etapa-8-no-se-activa-sola).

### 4.6 Recomendaciones prioritarias del diagnóstico

1. Levantar los `Dato requerido` de este documento antes de presentarlo externamente — un documento de dirección con vacíos de cifra pierde autoridad si no están explícitamente enmarcados como pendientes de validar (ya lo están aquí, pero conviene resolverlos antes de un pitch externo).
2. Confirmar con operación real si los 20 SOP del ciclo del activo ([índice](../chacontainer-master-project/sop/README.md)) reflejan la práctica actual o son un diseño objetivo.
3. Priorizar la Fase A del [roadmap técnico](../chacontainer-master-project/arquitectura-os.md#3-roadmap-de-implementación-propuesto) (estado detallado del activo) porque es la dependencia de casi todo lo demás en este documento.

---

## 5. Análisis del mercado

### 5.1 Problemas principales del mercado (segmento automotriz/industrial)

| Problema | Manifestación típica | Por qué el cliente no lo resuelve solo |
|---|---|---|
| Pérdida de activos retornables | Contenedores/racks que no regresan de proveedores o clientes | Sin identificación individual ni trazabilidad, "perdido" y "en tránsito lento" son indistinguibles |
| Mala ubicación / falta de visibilidad | Nadie sabe cuántos activos hay en una planta de proveedor hasta que se necesitan | El inventario vive en hojas de cálculo desconectadas del movimiento físico real |
| Sobre-inventario de seguridad | Se compra de más "por si acaso" | Es la respuesta racional a la incertidumbre — y el costo queda escondido como "activo", no como gasto de la falta de visibilidad (ver [diagnóstico comercial §3](../chacontainer-master-project/diagnostico-comercial.md#3-cuantificar-el-costo-de-no-saber)) |
| Mala condición al momento de uso | Activos dañados o sucios llegan a línea de producción | Sin criterio de inspección estandarizado ni responsable claro de mantenimiento preventivo |
| Falta de trazabilidad entre plantas/proveedores | Nadie puede reconstruir dónde estuvo un activo | La cadena OEM→Tier1→Tier2 multiplica los puntos de custodia sin que ninguno lleve el registro completo |

### 5.2 Dolor económico del cliente

Tres componentes, en orden de visibilidad para el cliente (de más a menos
visible) — desarrollados con más detalle en
[diagnostico-comercial.md §3](../chacontainer-master-project/diagnostico-comercial.md#3-cuantificar-el-costo-de-no-saber):
reposición por pérdida, compras/rentas de emergencia, y el mayor de todos
pero el más invisible — **sobre-inventario de seguridad** (capital
inmovilizado comprado "por si acaso" que nadie contabiliza como costo de no
saber).

`Dato requerido` para cuantificar en un caso real: costo unitario de
reposición por tipo de activo, tasa de pérdida anual declarada por el
cliente, % de sobre-inventario estimado. Ninguno se inventa aquí — se
levantan en el [diagnóstico comercial](../chacontainer-master-project/diagnostico-comercial.md)
con cada prospecto.

### 5.3 Dolor operativo del cliente

Tiempo del equipo de logística/compras dedicado a buscar, contar y
conciliar activos en vez de a su función real; paros de línea o retrasos de
embarque por falta de empaque disponible en el momento exacto que se
necesita (crítico en manufactura JIT/JIS).

### 5.4 Riesgo logístico

En una cadena OEM → Tier 1 → Tier 2, un activo puede estar retenido en
cualquiera de los eslabones sin que el eslabón anterior lo sepa. El riesgo
no es solo perder el activo — es que **el paro de producción llega antes de
que alguien note el problema**, porque no hay alerta temprana sin un umbral
de tiempo configurado ([reglas-operativas.md](../chacontainer-master-project/reglas-operativas.md)
y [alertas.md](../chacontainer-master-project/alertas.md) ya especifican
ese mecanismo).

### 5.5 Cómo entra CHACONTAINER como solución

Por la puerta que la [ETAPA 8 del proyecto maestro](../chacontainer-master-project/README.md#etapa-8--modelo-comercial-solutions--systems)
ya describe: un servicio puntual de Solutions (lavado, reparación,
inventario físico) que revela el vacío de datos, seguido de una oferta de
Systems escalonada — detallada en §14 de este documento.

### 5.6 Clientes con mayor urgencia de compra

| Perfil | Por qué tiene urgencia |
|---|---|
| Tier 1/Tier 2 en expansión por nearshoring | Planta nueva = sistema de empaque desde cero, sin inercia de "así lo hemos hecho siempre" que frene la adopción |
| OEM/Tier 1 con historial reciente de paro de línea por faltante de empaque | El costo ya se materializó y es fácil de señalar en la conversación comercial |
| Empresas con múltiples plantas/proveedores sin visibilidad centralizada | El problema de custodia dispersa es exactamente el que resuelve el modelo de gobernanza |

### 5.7 Qué cliente debería priorizar CHACONTAINER

Tier 1/Tier 2 con 1-3 plantas, en expansión activa, con relación comercial
previa de Solutions (venta o renta) — perfil que además coincide con los
criterios de selección de piloto ya definidos en
[piloto.md §2](../chacontainer-master-project/piloto.md#2-criterios-de-selección).
No se prioriza el OEM como primer cliente: el ciclo de decisión es más
largo y la complejidad organizacional mayor; el OEM es objetivo de Nivel 2
una vez que existan casos de éxito con Tier 1/Tier 2.

### 5.8 Servicios con mayor potencial comercial

Lavado, reparación e inventarios físicos — son los tres servicios que, por
diseño, generan la señal de conversión hacia Systems con mayor frecuencia
(ver [conversion-solutions-systems.md §2](../chacontainer-master-project/conversion-solutions-systems.md#2-señales-de-conversión-por-servicio-de-solutions)).

### 5.9 Barreras de entrada

- Contratos de suministro existentes con proveedores de empaque incumbentes, atados a la vida del programa vehicular.
- Resistencia a integrar un sistema externo a los procesos de calidad ya certificados (IATF 16949 y similares) — cualquier propuesta de trazabilidad debe presentarse como complemento, no como sustituto de esos sistemas.
- Escepticismo genuino: "ya nos han vendido software de trazabilidad antes y no funcionó" — la respuesta es el encuadre de piloto pagado autofinanciable ([piloto-pagado.md](../chacontainer-master-project/piloto-pagado.md)), no la promesa.

### 5.10 Recomendación estratégica de entrada al mercado

Entrar por Tier 1/Tier 2 en expansión de nearshoring, con Solutions como
puerta física y el [diagnóstico comercial](../chacontainer-master-project/diagnostico-comercial.md)
como instrumento de conversación — nunca vendiendo "gobernanza" como primer
mensaje (ver [propuesta comercial, ETAPA 11](../chacontainer-master-project/README.md#etapa-11--propuesta-comercial)).

`Dato requerido` para validar cifras de mercado: no se dispone de datos
públicos verificados de tamaño de mercado de empaque retornable automotriz
en México en este entorno de trabajo. Cualquier cifra de TAM/SAM/SOM debe
investigarse con fuentes como INA (Industria Nacional de Autopartes),
AMIA, o reportes sectoriales — no se debe presentar un número sin esa
fuente.

---

## 6. Problema principal del cliente

Sintetizando §5: el cliente automotriz/industrial objetivo **no sabe cuántos
activos retornables tiene realmente disponibles hoy, ni cuánto le cuesta no
saberlo** — y ese costo de no saber suele manifestarse como sobre-compra
(el componente más grande y más invisible), no solo como pérdida directa.
Es el mismo problema que el [diagnóstico comercial](../chacontainer-master-project/diagnostico-comercial.md)
del proyecto maestro ya usa como apertura de conversación, aquí confirmado
como el problema central también desde el ángulo específicamente automotriz.

---

## 7. Oportunidad estratégica

Ningún competidor tradicional de empaque retornable en el sector automotriz
mexicano combina, hoy, ejecución física (Solutions) con gobernanza de datos
como servicio recurrente (Systems) bajo un mismo proveedor. La oportunidad
es capturar ese espacio antes de que:

(a) un proveedor de software de trazabilidad genérico (sin capacidad
física) entre al sector sin poder ejecutar lavado/reparación/logística
inversa, o

(b) un competidor de empaque tradicional imite el discurso de
"trazabilidad" sin construir la infraestructura de datos detrás.

`Dato requerido`: mapa de competidores directos e indirectos en el mercado
mexicano de empaque retornable automotriz — no se dispone de esa
investigación en este entorno.

---

## 8. Modelo de negocio

| Línea de negocio | Cliente ideal | Problema que resuelve | Tipo de ingreso | Riesgo | KPI principal |
|---|---|---|---|---|---|
| Venta de activos | OEM/Tier 1/Tier 2 en alta o expansión | Necesita empaque nuevo/certificado sin gestionar fabricación | Transaccional | Ciclo largo de decisión en cuentas grandes | Margen por SKU |
| Renta | Cliente con necesidad temporal o de proyecto | No quiere inmovilizar capital en compra | Transaccional / por uso | Retorno tardío del activo rentado | Tasa de utilización del parque en renta |
| Lavado / reparación / reacondicionamiento | Cualquier cliente con parque en circulación | Activo sucio/dañado no puede reingresar | Transaccional recurrente | Dependencia de volumen de retorno de campo | Tiempo de ciclo de intervención |
| Compra de scrap / recuperación de valor | Cliente dando de baja parque obsoleto | Necesita liberarse de activos sin valor con retorno económico | Transaccional | Volumen impredecible | Margen scrap vs. recuperación |
| Diagnóstico de parque (Oferta 1) | Cliente con visibilidad baja, dolor declarado | No sabe cuánto tiene ni dónde | Transaccional (proyecto) | Puede no convertir a Systems | % de diagnósticos que avanzan a piloto |
| Identificación y registro (Oferta 2) | Cliente que ya midió su parque | Tiene el conteo pero no puede darle seguimiento | Setup + recurrente | Setup regalado erosiona margen si no se cobra | % del parque identificado |
| Control y trazabilidad (Oferta 3) | Cliente con pérdidas recurrentes | Sabe lo que tiene, se le sigue perdiendo | Recurrente + variable | Logística inversa sin tope consume margen | Cumplimiento de retorno |
| Gobernanza (Oferta 4) | Cliente Nivel 5, relación madura | Quiere un responsable, no solo un reporte | Recurrente (fee de gobernanza) | Incentivo perverso si no hay piso de contrato (ya corregido, ver [pricing.md §2](../chacontainer-master-project/pricing.md#2-el-incentivo-perverso-que-hay-que-resolver)) | Disponibilidad real |

Detalle completo de qué recibe, qué ejecuta, qué evidencia genera cada
línea: [catalogo-solutions.md](../chacontainer-master-project/catalogo-solutions.md)
y [catalogo-systems.md](../chacontainer-master-project/catalogo-systems.md).
Costos, recursos y alianzas necesarias se detallan en §16-17 de este
documento.

**Escalar el negocio** significa, en este modelo, repetir la secuencia
Solutions → señal → Diagnóstico → Piloto → Gobernanza en más cuentas
Tier 1/Tier 2, apoyado en [hubs y partners](../chacontainer-master-project/escala-hubs-y-partners.md)
una vez que el volumen por ruta lo justifique — no en vender más contenedores
por sí solos.

---

## 9. Packaging Solutions

Definición, alcance y catálogo completo de los 15 servicios ya
desarrollados en
[catalogo-solutions.md](../chacontainer-master-project/catalogo-solutions.md).
En el lenguaje de dirección: **Packaging Solutions es la intervención física
directa sobre el activo** — venta, renta, lavado, reparación, modificación,
scrap. Es lo que CHACONTAINER ya sabe ejecutar.

- **Cliente ideal**: cualquier empresa con parque de empaque retornable en operación, sin importar su nivel de madurez de datos.
- **Forma de cobro**: transaccional, por servicio o por proyecto.
- **Cómo explicarlo a un director de planta**: "resolvemos el problema físico que tienes hoy en tu línea — empaque sucio, dañado o faltante — sin pedirte que cambies tu forma de operar."
- **Cómo explicarlo a compras**: "es una compra o servicio puntual, con el mismo proceso de cotización y orden de compra que ya manejas con cualquier proveedor."

## 10. Packaging Systems

Definición, alcance y las 16 capacidades ya desarrolladas en
[catalogo-systems.md](../chacontainer-master-project/catalogo-systems.md).
En el lenguaje de dirección: **Packaging Systems es la gobernanza del
sistema donde el activo opera** — identificación, trazabilidad, custodia,
reglas, alertas, analítica y el fee de administración recurrente que
sostiene todo eso.

- **Cliente ideal**: cliente que ya pasó por un [diagnóstico comercial](../chacontainer-master-project/diagnostico-comercial.md) y confirmó un problema de visibilidad cuantificable.
- **Forma de cobro**: mixta (setup) y recurrente (fee por activo/mes + fee de gobernanza) — modelo completo en [pricing.md](../chacontainer-master-project/pricing.md).
- **Cómo explicarlo a un director de planta**: "vas a saber, en cualquier momento, cuántos activos tienes disponibles de verdad — no lo que dice el sistema contable."
- **Cómo explicarlo a logística**: "cada movimiento queda registrado automáticamente; ya no dependes de que alguien lo anote a mano."
- **Cómo explicarlo a dirección general**: "convertimos una pérdida operativa recurrente e invisible en un costo medido, con un responsable nombrado de nuestro lado."

### 10.1 Diferencia comercial, operativa y financiera

| | Packaging Solutions | Packaging Systems |
|---|---|---|
| Comercial | Venta/cotización puntual | Contrato de servicio recurrente, escalera de 4 ofertas (ver §14) |
| Operativa | Ejecución física sobre el activo | Captura y gobierno de datos sobre el activo y su ciclo |
| Financiera | Ingreso transaccional, margen por servicio | Ingreso recurrente, margen por administración; ver corrección de incentivo perverso en [pricing.md §2](../chacontainer-master-project/pricing.md#2-el-incentivo-perverso-que-hay-que-resolver) |

### 10.2 Cómo vender primero Solutions y escalar a Systems

Mecanismo completo, con señales por servicio y umbrales de conversión:
[conversion-solutions-systems.md](../chacontainer-master-project/conversion-solutions-systems.md).

---

## 11. Propuesta de valor

### 11.1 Propuesta de valor principal

> CHACONTAINER convierte el costo invisible de no saber dónde está tu
> empaque retornable en un sistema medido, gobernado y con un responsable
> nombrado — sin dejar de resolver, hoy mismo, el problema físico que tienes
> en línea.

### 11.2 Por audiencia

| Audiencia | Propuesta de valor específica |
|---|---|
| OEM | Visibilidad consolidada del parque de empaque a través de toda su cadena de proveedores, sin depender de que cada Tier reporte manualmente |
| Tier 1 | Reducción del costo de reposición y del riesgo de paro de línea propio y hacia su OEM, con un sistema que no exige cambiar su ERP |
| Tier 2 | Acceso a gestión de activos de nivel Tier 1 sin la inversión de construirla internamente |
| Empresas con pérdida de activos | El diagnóstico cuantifica la pérdida real (no la percibida) antes de pedir ningún compromiso |
| Empresas que necesitan reducir CAPEX | Renta y gestión de activos existentes en vez de comprar más parque de seguridad |
| Empresas que necesitan velocidad operativa | Escaneo y reglas automáticas reemplazan la búsqueda manual y la conciliación en hoja de cálculo |

### 11.3 Mensajes por formato

- **Mensaje corto comercial**: "¿Cuántos de tus contenedores están realmente disponibles hoy — y cuánto te cuesta no saberlo?" (tomado directamente de la [ETAPA 11 del proyecto maestro](../chacontainer-master-project/README.md#etapa-11--propuesta-comercial)).
- **Mensaje ejecutivo para dirección**: "Gestionamos el ciclo de vida completo de tu empaque retornable — desde el suministro hasta su recuperación — con datos, no con supuestos."
- **Mensaje técnico para operaciones/logística**: "Cada movimiento de tus activos queda registrado automáticamente vía QR, con alertas cuando algo se sale de la regla que tú definiste."
- **Frase de posicionamiento de marca**: "CHACONTAINER: Operating System for Returnable Packaging" (consistente con el cierre del [proyecto maestro](../chacontainer-master-project/README.md#norte-del-proyecto)).

### 11.4 Diferenciadores frente a un proveedor tradicional

Un proveedor tradicional vende o renta el activo y termina ahí. CHACONTAINER
mantiene la relación después de la entrega — el activo queda dentro de un
sistema que sigue generando datos y valor durante todo su ciclo de vida, no
solo en el momento de la transacción.

---

## 12. Segmentación de clientes

| Segmento | Perfil de cliente ideal | Dolor principal | Servicio recomendado de entrada | Urgencia de compra | Complejidad de cierre | Decisor principal |
|---|---|---|---|---|---|---|
| Tier 1/Tier 2 en expansión (nearshoring) | Planta nueva o en ramp-up en México | Sistema de empaque desde cero, sin trazabilidad | Venta + Identificación (Oferta 2) | Alta | Media | Gerencia de planta / logística |
| Tier 1/Tier 2 con pérdidas recurrentes | Parque en operación, sin visibilidad | Reposición constante, sobre-inventario | Diagnóstico de parque (Oferta 1) | Alta | Media-alta (requiere cuantificar el dolor) | Compras / logística, con validación de dirección |
| OEM con múltiples proveedores | Cadena de suministro compleja, historial de paros por faltante | Falta de visibilidad consolidada entre Tiers | Diagnóstico + piloto en una planta/ruta acotada | Media (ciclo largo) | Alta | Comité de compras/calidad, no una sola persona |
| Empresas con necesidad de renta temporal | Proyecto o pico de producción | Evitar inmovilizar capital | Renta | Media | Baja | Compras |
| Empresas con inventario grande sin control | Parque histórico, crecido sin sistema | No saben cuánto tienen realmente | Inventario físico + clasificación | Media-alta | Media | Logística / finanzas (por el ángulo de capital inmovilizado) |

`Dato requerido` para tickets estimados: no se inventan cifras de ticket
promedio. La fórmula aplicable es la de §16 (fórmula de precios) aplicada al
volumen real de cada prospecto (número de activos, tipo, ubicaciones).

---

## 13. Estrategia comercial y tubería de ventas

### 13.1 Prospección y canales

Apoyada en la base instalada de clientes de Solutions (la conversión
descrita en §10.2) más prospección activa en plantas Tier 1/Tier 2 de
nueva instalación (nearshoring) — estas últimas identificables por
anuncios de inversión industrial, asociaciones sectoriales (INA, clústeres
automotrices estatales) y ferias del sector.

### 13.2 Pipeline comercial por etapas

| Etapa | Objetivo | Acción | Responsable | Evidencia | KPI | Criterio para avanzar |
|---|---|---|---|---|---|---|
| Prospección | Identificar cuentas con perfil de §12 | Investigación + contacto inicial | Comercial | Registro en CRM | # de cuentas contactadas | Contacto calificado responde |
| Diagnóstico comercial | Cuantificar el dolor declarado | Ejecutar las [12 preguntas del diagnóstico](../chacontainer-master-project/diagnostico-comercial.md#2-las-preguntas) | Comercial | Estimado de costo de no saber, firmado por el cliente | # de diagnósticos completados | Cliente acepta medir (pasar a Oferta 1) |
| Oferta 1 · Diagnóstico de parque | Convertir el estimado en línea base real | Inventario físico + identificación parcial | Operación + comercial | Reporte de inventario físico | % de discrepancia encontrada | Cliente firma piloto pagado |
| Piloto pagado | Demostrar valor medido | Ejecutar según [piloto.md](../chacontainer-master-project/piloto.md#3-cronograma) | Operación + gobernanza | Reporte de cierre del piloto | Cumplimiento de los 3 criterios de éxito | Se cumplen criterios cuantitativo + operativo + de datos |
| Cierre a contrato de gobernanza | Convertir piloto en recurrencia | Negociación de contrato Nivel 5 | Comercial + fundador | Contrato firmado | Tasa de conversión piloto→contrato | Firma |
| Escalamiento | Expandir a más plantas/familias del mismo cliente | Propuesta de extensión | Gobernanza | Nuevo alcance firmado | Ingreso recurrente por cuenta | Cliente aprueba extensión |

### 13.3 Argumento de entrada y objeciones

Argumento de apertura: la pregunta de la [ETAPA 11](../chacontainer-master-project/README.md#etapa-11--propuesta-comercial)
(§11.3 de este documento). Objeciones comunes esperables en el sector y
cómo se resuelven:

| Objeción | Respuesta |
|---|---|
| "Ya tenemos proveedor de empaque" | No se pide reemplazarlo; el diagnóstico es gratuito y no exclusivo |
| "Ya probamos trazabilidad y no funcionó" | El piloto es pagado, acotado y autofinanciable — el fracaso previo no compromete presupuesto nuevo sin evidencia primero |
| "No tenemos presupuesto para gobernanza" | El punto de entrada nunca es "gobernanza" — es un servicio de Solutions ya presupuestado (lavado, reparación) |
| "Nuestro sistema de calidad ya cubre esto" | El sistema de calidad certifica el proceso; no da visibilidad de inventario físico en tiempo real — son complementarios, no competidores |

### 13.4 KPIs comerciales

Volumen de diagnósticos comerciales realizados, tasa de conversión
diagnóstico → Oferta 1, tasa de conversión piloto → contrato de gobernanza,
valor de pipeline por etapa. `Dato requerido`: metas numéricas para cada
uno — no se fijan aquí sin datos históricos de conversión real.

---

## 14. Oferta comercial

Las cuatro ofertas ya desarrolladas con alcance cerrado en
[oferta-systems.md](../chacontainer-master-project/oferta-systems.md):
**Diagnóstico de parque** (Nivel 3, puerta de entrada), **Identificación y
registro** (Nivel 3), **Control y trazabilidad** (Nivel 4), **Gobernanza
del sistema** (Nivel 5). Cada una con problema que resuelve, entregable,
qué incluye y qué no, y compromiso — detalle completo en el documento
enlazado, incluida la regla de que **no se venden salteadas** (§3 de ese
documento).

Para el sector automotriz específicamente, dos matices adicionales:

- La **Oferta 1 (Diagnóstico)** debe explícitamente no interferir con procesos de calidad ya certificados — se presenta como complementaria, nunca como auditoría de calidad.
- La **Oferta 4 (Gobernanza)** en cuentas OEM probablemente requiere adaptarse a la cadena multi-Tier — el alcance de "un cliente" puede en la práctica significar "un OEM y los Tier 1/Tier 2 que le suministran directamente", lo cual no está resuelto en el documento base y queda como brecha a definir cuando exista el primer caso real `[VALIDAR: alcance de contrato de gobernanza en cadenas multi-Tier]`.

---

## 15. Modelo de ingresos

Igual estructura que [ETAPA 10 del proyecto maestro](../chacontainer-master-project/README.md#etapa-10--modelo-de-ingresos):
transaccionales (venta, scrap, reparación, lavado, modificación), por uso
(renta, logística, recuperación), recurrentes (administración, inventario,
tracking, mantenimiento), y fee de gobernanza. La mezcla objetivo entre
transaccional y recurrente es, deliberadamente, **no fija**: cada cuenta
avanza en la [escalera comercial](../chacontainer-master-project/README.md#etapa-9--escalera-comercial)
a su propio ritmo.

`Dato requerido`: mezcla actual real de ingresos de CHACONTAINER por línea,
para poder proyectar cómo se movería con la adopción de Systems.

---

## 16. Fórmula de precios

CHACONTAINER ya tiene una regla explícita y vigente
([`lista-precios-costos.md`](../chacontainer/docs/lista-precios-costos.md)):
no inventar precios ni condiciones que no estén validadas. Este documento
la respeta y entrega **fórmula, no cifra**:

```
Precio del servicio = Costo directo
                     + Costo indirecto asignado
                     + Riesgo operativo (ver nota)
                     + Margen deseado
                     + Impuestos aplicables
```

| Componente | Qué incluye | Dato requerido para calcularlo |
|---|---|---|
| Costo directo | Mano de obra de la intervención, insumos/químicos, material de identificación (QR), transporte directo | Costo real de mano de obra por hora, costo de insumos por tipo de activo |
| Costo indirecto asignado | Prorrateo de planta, supervisión, administración | Costos fijos mensuales de operación / volumen de servicios ejecutados |
| Riesgo operativo | Provisión por incidencias esperadas (daño no detectado, retraso) — se calcula como % histórico de incidencias × costo promedio de resolución, no como cifra fija | Tasa histórica de incidencias por tipo de servicio (no existe hasta que haya operación medida — ver [modelo-de-datos.md](../chacontainer-master-project/modelo-de-datos.md)) |
| Margen deseado | Definido por dirección, no calculado | Decisión de dirección — no es un dato a "descubrir" |
| Impuestos aplicables | Según jurisdicción y tipo de operación | Régimen fiscal aplicable — `Dato requerido` de contabilidad |

Para el **fee mensual de gobernanza** (Oferta 4), la fórmula base ya está
resuelta en dirección en
[pricing.md §2](../chacontainer-master-project/pricing.md#2-el-incentivo-perverso-que-hay-que-resolver):
piso fijo sobre el parque administrado al firmar el contrato, recalculado
al renovar — evita el incentivo perverso de cobrar menos cuando el sistema
funciona mejor.

**Variables críticas a levantar en campo antes de cotizar cualquier
cuenta real**: costo real de identificación por activo en sitio, costo
mensual de operar la plataforma por activo administrado, horas mensuales
que consume el ritual de gobernanza — las tres ya señaladas como pendientes
en [pricing.md §5](../chacontainer-master-project/pricing.md#5-lo-que-falta-decidir-antes-de-cotizar).

No se presenta aquí ninguna tabla de precios por SKU o por hora: sería
inventar un dato que el documento fuente prohíbe explícitamente inventar.

---

## 17. Arquitectura operativa

Flujo general, de prospección a renovación:

```
Prospección → Diagnóstico comercial → Oferta 1 (línea base) →
Piloto pagado → Cierre a Gobernanza → Operación continua
(escaneo + reglas + alertas + comité mensual) → Renovación / escalamiento
```

Cada bloque de este flujo ya tiene su documento de detalle en el proyecto
maestro:

| Bloque | Documento de referencia |
|---|---|
| Diagnóstico | [diagnostico-comercial.md](../chacontainer-master-project/diagnostico-comercial.md) |
| Piloto | [piloto.md](../chacontainer-master-project/piloto.md) + [piloto-pagado.md](../chacontainer-master-project/piloto-pagado.md) |
| Operación (SOP del ciclo del activo) | [sop/README.md](../chacontainer-master-project/sop/README.md) (20 SOP) |
| Identificación / escaneo | [qr.md](../chacontainer-master-project/qr.md) |
| Reglas y alertas | [reglas-operativas.md](../chacontainer-master-project/reglas-operativas.md) + [alertas.md](../chacontainer-master-project/alertas.md) |
| Gobernanza (comité, reportes) | [modelo-de-gobernanza.md](../chacontainer-master-project/modelo-de-gobernanza.md) |
| Dashboard / indicadores | [dashboard.md](../chacontainer-master-project/dashboard.md) |

Puntos de control críticos: validación de transición en cada escaneo
(evita que un activo avance de estado sin cumplir el paso anterior, ver
[qr.md §3](../chacontainer-master-project/qr.md#3-validación-de-transición-en-el-escaneo)),
y el comité mensual de gobernanza como punto de control estratégico
(§13.2 de este documento y [modelo-de-gobernanza.md §3](../chacontainer-master-project/modelo-de-gobernanza.md#3-ritual-de-gobernanza-propuesto)).

---

## 18. Estructura organizacional mínima

Roles que exige operar este modelo, tomados de los SOP y del modelo de
gobernanza ya definidos (nombres de puesto exactos pendientes de
confirmar — ver [pendientes-de-validacion.md, Nivel 2 §2.1](../chacontainer-master-project/pendientes-de-validacion.md#21-roles-exactos-la-categoría-más-frecuente)):

| Función | Responsabilidad central |
|---|---|
| Comercial | Prospección, diagnóstico comercial, cierre de ofertas y pilotos |
| Operación de planta (lavado/reparación/inspección) | Ejecución de los SOP físicos del ciclo del activo |
| Responsable de inventario/identificación | Alta, identificación y conciliación del registro maestro |
| Logística inversa | Recolección de activos retenidos/detenidos |
| Gobernanza (Nivel 3) | El fundador, delegable por cuenta según tamaño (decisión ya tomada, ver [pendientes-de-validacion.md](../chacontainer-master-project/pendientes-de-validacion.md)) |
| Técnico / soporte de plataforma | Mantenimiento del sistema, evaluador de reglas y alertas |

`Dato requerido`: organigrama real actual de CHACONTAINER, para mapear
quién cubre cada función hoy y dónde hay vacíos.

---

## 19. KPIs

Sistema completo ya especificado en
[kpi.md](../chacontainer-master-project/kpi.md): los 10 KPI de la
[ETAPA 6 del proyecto maestro](../chacontainer-master-project/README.md#etapa-6--indicadores-del-sistema)
(disponibilidad, utilización, tiempo de ciclo, cumplimiento de retorno,
pérdida, daño, tiempo detenido, costo por ciclo, rotación, disponibilidad
real), cada uno con fórmula, fuente de datos real contra el esquema de
`chacontainer/`, y si es calculable hoy o depende de tablas del roadmap
técnico (4 de 10 lo son hoy sin cambios).

KPIs comerciales específicos de este documento: ver §13.4. Metas
numéricas: `Dato requerido` en todos los casos — se fijan con datos
históricos reales, no por defecto.

---

## 20. Riesgos

| Riesgo | Categoría | Causa probable | Impacto | Probabilidad | Señal temprana | Acción preventiva | Prioridad |
|---|---|---|---|---|---|---|---|
| Baja adopción de escaneo por operadores del cliente | Operativo/comercial | Falta de presión interna sin patrocinador claro | Piloto no genera datos confiables | Media-alta | Escaneos esperados vs. reales cayendo desde semana 2 | Exigir responsable operativo asignado antes de firmar piloto ([piloto-pagado.md §6](../chacontainer-master-project/piloto-pagado.md#6-a-quién-ofrecerlo)) | Alta |
| Incentivo perverso de pricing erosiona margen al funcionar bien | Financiero | Fee ligado 1:1 al conteo de activos sin piso | Ingreso recurrente baja cuando el sistema mejora el resultado del cliente | Media | Caída de fee en cuentas con mejora medida | Ya corregido con piso fijo por vigencia de contrato ([pricing.md §2](../chacontainer-master-project/pricing.md#2-el-incentivo-perverso-que-hay-que-resolver)) | Media (mitigado) |
| Volatilidad de demanda automotriz | Externo/comercial | Paros de producción, cambios de programa vehicular | Caída de demanda de Solutions y de activos administrables | Media | Anuncios de recorte de producción en clientes clave | Diversificar cartera entre OEM/Tier 1/Tier 2 y entre programas vehiculares | Media |
| Régimen aduanal restringe el ciclo de vida en operación México-EE.UU. | Legal/logístico | Importación temporal con obligación de retorno | Decisiones de baja/scrap legalmente restringidas | Baja hoy (FASE V, no inmediata) | — | Consultar asesor aduanal antes de operar cruces ([escala-geografica.md §3.2](../chacontainer-master-project/escala-geografica.md#32-el-régimen-aduanal-restringe-el-ciclo-de-vida)) | Baja (aún no aplica) |
| Dependencia de cuentas grandes concentradas | Comercial/financiero | Pocos clientes Nivel 5 representando ingreso recurrente mayoritario | Pérdida de una cuenta grande afecta desproporcionadamente el ingreso recurrente | `Dato requerido` (depende de concentración real de cartera) | Concentración de ingreso por cliente > umbral | Diversificar activamente el pipeline de Nivel 3-4 | Alta si se confirma concentración |
| Falta de capital para escalar identificación/tecnología | Financiero | Setup de identificación no cobrado o subvaluado | Cada alta de cliente nuevo es pérdida de caja | Media | Margen negativo en cuentas de arranque | Cobrar setup por adelantado, no regalarlo ([pricing.md §4](../chacontainer-master-project/pricing.md#4-estructura-por-oferta)) | Alta |

Matriz de prioridad resumida: **Alta** (adopción de escaneo, dependencia de
capital en setup, concentración de cartera si se confirma) — **Media**
(volatilidad automotriz, incentivo de pricing ya mitigado) — **Baja**
(riesgo aduanal, no aplica hasta FASE V).

---

## 21. Plan de implementación

Fases ya detalladas en el roadmap técnico y operativo del proyecto maestro,
consolidadas aquí para dirección:

| Fase | Objetivo | Entregable | Documento de detalle |
|---|---|---|---|
| Preparación interna | Cerrar decisiones de Nivel 1 pendientes | Lista de decisiones resueltas | [pendientes-de-validacion.md](../chacontainer-master-project/pendientes-de-validacion.md) |
| Documentación y procesos | Confirmar SOP contra operación real | 20 SOP validados en campo | [sop/README.md](../chacontainer-master-project/sop/README.md) |
| Validación comercial | Ejecutar diagnósticos comerciales reales | Primeros 3-5 diagnósticos con clientes Tier 1/Tier 2 | [diagnostico-comercial.md](../chacontainer-master-project/diagnostico-comercial.md) |
| Piloto con cliente | Ejecutar el primer piloto pagado | Reporte de cierre de piloto | [piloto.md](../chacontainer-master-project/piloto.md) |
| Implementación tecnológica | Construir Fase A-C del roadmap técnico (estado detallado, custodia, reglas/alertas) | MVP operando ([mvp.md](../chacontainer-master-project/mvp.md)) | [arquitectura-os.md §3](../chacontainer-master-project/arquitectura-os.md#3-roadmap-de-implementación-propuesto) |
| Estandarización operativa | Ritual de gobernanza corriendo mensualmente | Primer comité de gobernanza ejecutado | [modelo-de-gobernanza.md](../chacontainer-master-project/modelo-de-gobernanza.md) |
| Escalamiento comercial | Repetir el embudo en más cuentas | Pipeline con varias cuentas en cada etapa | §13 de este documento |
| Expansión | Hubs, partners, multi-planta | Ver condiciones de activación, no plan fijo | [escala-hubs-y-partners.md](../chacontainer-master-project/escala-hubs-y-partners.md), [escala-geografica.md](../chacontainer-master-project/escala-geografica.md) |

`Dato requerido` por fase: responsable nombrado, duración estimada real y
recursos asignados — dependen de la estructura organizacional real (§18),
no disponible en este entorno.

---

## 22. Plan a 12 meses

| Trimestre | Objetivo principal | Iniciativas clave | Resultado esperado |
|---|---|---|---|
| T1 | Orden interno y estructura comercial | Resolver decisiones de Nivel 1 pendientes; confirmar SOP contra campo; construir Fase A del roadmap técnico (estado detallado) | Base de datos y documentación lista para operar el piloto |
| T2 | Prospección y primeros pilotos | Ejecutar 3-5 diagnósticos comerciales con Tier 1/Tier 2; cerrar 1-2 pilotos pagados | Primer(os) piloto(s) en ejecución con línea base medida |
| T3 | Escalamiento operativo y tecnología | Cerrar Fase B-C del roadmap técnico (custodia, reglas, alertas); primer comité de gobernanza mensual corriendo | Primer contrato de gobernanza firmado |
| T4 | Consolidación y expansión | Evaluar condiciones de activación de hub/partner (§21 de [escala-hubs-y-partners.md](../chacontainer-master-project/escala-hubs-y-partners.md)); preparar reporte de resultados para dirección/inversión | Caso de negocio con datos reales, no solo proyectado |

`Dato requerido`: responsables, recursos y metas numéricas por trimestre —
dependen de decisiones de dirección y de capacidad instalada real no
disponibles en este entorno.

---

## 23. Recomendaciones ejecutivas

1. **No vender "gobernanza" como primer mensaje.** El punto de entrada siempre es un servicio de Solutions ya presupuestado por el cliente — la conversión a Systems se construye desde ahí (§10.2, §13.3).
2. **Priorizar Tier 1/Tier 2 en expansión de nearshoring como primer segmento**, no el OEM — ciclo de decisión más corto, menor complejidad de cierre (§12).
3. **Cerrar los `Dato requerido` de este documento antes de usarlo externamente** — particularmente cifras de facturación actual, capacidad instalada, y organigrama real.
4. **Construir la Fase A del roadmap técnico (estado detallado del activo) antes que cualquier otra pieza de tecnología** — es la dependencia de la que cuelgan disponibilidad real, custodia y la mayoría de los KPI (§16, §19).
5. **Cobrar el setup de identificación desde la primera cuenta piloto** — regalarlo para "quitar fricción" convierte cada alta de cliente en pérdida de caja (§16, §20).
6. **No comprometer penalización contractual sobre KPI todavía** — el compromiso blando con piso de responsabilidad ya decidido (Oferta 4) es el nivel de riesgo correcto hasta tener casos reales medidos.

---

## 24. Próximos pasos

1. Validar con el fundador los `Dato requerido` marcados en §3, §4.3-4.4, §5.10, §7, §12, §15, §18, §21-22 — priorizando facturación actual, capacidad instalada y organigrama.
2. Seleccionar la primera cuenta piloto real siguiendo los criterios de §12 y [piloto.md §2](../chacontainer-master-project/piloto.md#2-criterios-de-selección).
3. Ejecutar el primer diagnóstico comercial real usando [diagnostico-comercial.md](../chacontainer-master-project/diagnostico-comercial.md) y contrastar los supuestos de este documento (§5-6) contra la respuesta real del cliente.
4. Iniciar la Fase A del roadmap técnico ([arquitectura-os.md §3](../chacontainer-master-project/arquitectura-os.md#3-roadmap-de-implementación-propuesto)) en paralelo a la prospección comercial, no después.

---

*Documento construido sobre el marco ya desarrollado en
[`chacontainer-master-project/`](../chacontainer-master-project/) (27 pasos
de la estrategia maestra, 20 SOP del ciclo del activo, modelo de datos,
gobernanza y pricing), con la capa de especificidad automotriz/OEM/Tier 1/
Tier 2/nearshoring que ese proyecto no cubría. Ver el banco de prompts que
originó este documento en [`prompts-fuente.md`](./prompts-fuente.md).*

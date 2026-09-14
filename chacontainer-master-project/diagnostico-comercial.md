# Diagnóstico comercial — el instrumento de apertura

**Fase:** IV · Comercial — paso 18 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

La [ETAPA 11](./README.md#etapa-11--propuesta-comercial) dice cómo abrir la
conversación: *"¿Cuántos activos tienen realmente disponibles hoy y cuánto
les cuesta no saber dónde están los demás?"*. Es la pregunta correcta, pero
tal cual no se puede ejecutar: si el cliente supiera la respuesta no
necesitaría a CHACONTAINER. Este documento la convierte en un instrumento
estructurado que produce un número que el cliente reconozca como propio.

---

## 1. Tres diagnósticos distintos que conviene no confundir

El proyecto usa la palabra "diagnóstico" en tres lugares. Son cosas
diferentes y ocurren en momentos distintos:

| Nombre | Dónde vive | Qué diagnostica | Cuándo |
|---|---|---|---|
| **Diagnóstico consultivo** | [SOP-03 comercial](../chacontainer/docs/sop/SOP-03-diagnostico-consultivo.md) del SaaS | La necesidad del cliente respecto a un servicio de Solutions | Paso 5 del ciclo comercial existente |
| **Diagnóstico comercial de Systems** | Este documento | Si el cliente tiene un problema de sistema y cuánto le cuesta | Antes de proponer cualquier [oferta de Systems](./oferta-systems.md) |
| **Diagnóstico de activos** | [Systems #1](./catalogo-systems.md#1-diagnóstico-de-activos) + [Oferta 1](./oferta-systems.md#oferta-1--diagnóstico-de-parque-nivel-3-puerta-de-entrada) | El estado real del parque, con datos medidos | Ya contratado y facturado |

El de este documento es **gratuito y de conversación**; el de la Oferta 1 es
**pagado y de campo**. El primero vende al segundo. Confundirlos lleva a
regalar trabajo de inventario o a cobrar por una plática.

## 2. Las preguntas

Doce preguntas, agrupadas por lo que revelan. No es un cuestionario para
leer de corrido: es el material del que se elige según hacia dónde vaya la
conversación.

**Sobre el tamaño del problema**
1. ¿Cuántos activos retornables tienen en total? ¿De dónde sale ese número?
2. Si mañana necesitaran 100 unidades para un embarque, ¿cuántas podrían usar hoy sin comprar ni rentar?
3. ¿Cuántos compraron el año pasado? ¿Cuántos de esos fueron reposición y no crecimiento?

*La pregunta 1 es la más reveladora por su segunda mitad. Si el número sale
de un ERP que nadie concilia contra físico, ya se encontró el problema.*

**Sobre la visibilidad**
4. ¿Dónde están hoy los que no están en su planta?
5. ¿Quién es responsable de un activo cuando está en sitio del cliente?
6. ¿Cómo se enteran de que un activo no volvió? ¿En cuánto tiempo?

*La 6 suele responderse con "cuando hacemos inventario" o "cuando nos hace
falta". Ambas respuestas cuantifican el tiempo detenido invisible.*

**Sobre el costo**
7. ¿Cuántos activos dan por perdidos al año?
8. ¿Han comprado de emergencia o rentado por no encontrar los propios?
9. ¿Han parado o retrasado un embarque por falta de empaque disponible?
10. ¿Cuánto tiempo del equipo se va en buscar, contar o conciliar activos?

**Sobre la disposición a actuar**
11. ¿Alguien tiene hoy la responsabilidad explícita de que los activos regresen?
12. Si tuvieran el dato exacto mañana, ¿qué harían distinto?

*La 12 califica la oportunidad. Si el cliente no sabe qué haría con el dato,
no está listo para Systems por más que el problema exista.*

## 3. Cuantificar el costo de no saber

La conversación tiene que terminar en un número. Cinco componentes; el
cliente normalmente solo ve el primero.

| Componente | Cómo estimarlo | Visibilidad para el cliente |
|---|---|---|
| **Reposición por pérdida** | Activos perdidos al año × costo unitario de reposición | Alta — es el único que suele tener presente |
| **Sobre-inventario de seguridad** | Activos comprados de más para cubrir la incertidumbre × costo de capital anual | **Nula, y suele ser el mayor** |
| **Compras y rentas de emergencia** | Eventos al año × sobreprecio pagado vs. compra planeada | Media |
| **Paros y retrasos por faltante** | Eventos al año × costo del evento (flete urgente, penalización, línea detenida) | Media, pero muy sensible |
| **Horas de conciliación** | Horas/mes buscando y contando × costo por hora × 12 | Baja — se percibe como "parte del trabajo" |

**El sobre-inventario es el hallazgo que cambia la conversación.** Una
empresa que no confía en su disponibilidad real compra un colchón para no
parar nunca. Ese colchón es capital inmovilizado que nadie contabiliza como
costo de la falta de visibilidad — se contabiliza como "activos". Si el
cliente tiene 30% más parque del que su operación necesita, ese 30% es el
precio anual de no saber, y suele ser mayor que todas las pérdidas juntas.

**Regla de honestidad:** todos estos números son declarados por el cliente,
no medidos. Deben presentarse como *estimación construida con sus propios
supuestos*, nunca como diagnóstico. La conversión de estimado a medido es
exactamente lo que se vende en la [Oferta 1](./oferta-systems.md#oferta-1--diagnóstico-de-parque-nivel-3-puerta-de-entrada)
`[VALIDAR: dejar constancia escrita del supuesto de cada cifra]`.

## 4. El cierre

El diagnóstico termina en una sola propuesta: **medir**. No en vender
gobernanza, ni trazabilidad, ni software.

> "Estos números son los suyos, con sus supuestos. Ninguno está medido.
> Propongo empezar por medirlos: un inventario físico de una familia de
> activos en una planta. Si los números resultan mucho menores de lo que
> estimamos, no hay proyecto y usted lo sabrá con certeza. Si resultan
> mayores, ya tenemos la línea base."

Ese encuadre hace tres cosas: baja el riesgo percibido a una intervención
acotada, respeta la posibilidad de que no haya problema, y produce el
entregable que el [piloto](./piloto.md#1-el-problema-de-la-línea-base)
necesita como línea base.

## 5. Cuándo no seguir

Conviene retirarse, o quedarse solo en Solutions, si:

- La respuesta a la pregunta 12 es vaga. Sin una acción prevista, el dato es un reporte que nadie lee.
- Nadie es responsable de los retornos (pregunta 11) y nadie está dispuesto a serlo. Systems no crea esa responsabilidad; la hace visible y exigible, pero alguien del lado del cliente tiene que aceptarla.
- El costo estimado de no saber es menor que el costo de la Oferta 1. Ocurre con parques chicos o de bajo valor unitario, y decirlo abiertamente construye la credibilidad que sostiene la relación de Solutions.

---

## Próximo paso

[Piloto pagado](./piloto-pagado.md) (paso 19): cómo se estructura
comercialmente el paso que sigue a este diagnóstico.

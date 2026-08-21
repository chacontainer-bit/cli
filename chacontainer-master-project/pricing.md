# Pricing de Packaging Systems — modelo, no tarifas

**Fase:** IV · Comercial — paso 17 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

> **Este documento no contiene precios, y es deliberado.**
> [`chacontainer/docs/lista-precios-costos.md`](../chacontainer/docs/lista-precios-costos.md)
> es la fuente única de verdad para cotizar, tiene una regla explícita —*"No
> inventar precios, tiempos ni condiciones que no estén aquí"*— y hoy todos
> sus campos están en `[VALIDAR]`. Inventar aquí una tabla de tarifas
> produciría números con apariencia de autoridad que nadie validó y que
> terminarían en una cotización real. Lo que sigue es el **modelo**: qué se
> cobra, sobre qué unidad y con qué lógica. Las cifras se llenan en la lista
> de precios, no aquí.

---

## 1. La decisión que ordena todo: la métrica de valor

Systems puede cobrarse sobre cinco unidades distintas. La elección no es
cosmética: define qué comportamiento premia el modelo.

| Métrica | Cómo funciona | A favor | En contra |
|---|---|---|---|
| **Por activo / mes** | Fee por cada activo bajo administración | Escala con el tamaño del parque; fácil de explicar; estándar en la industria | Incentivo perverso, ver §2 |
| **Por planta / mes** | Fee fijo por instalación cubierta | Predecible para ambos; no penaliza parques grandes | Un cliente con 200 activos paga igual que uno con 5,000 en la misma planta |
| **Por usuario** | Fee por operador con acceso | Familiar en SaaS | Pésimo encaje: el valor no está en cuánta gente mira, y penaliza que más operadores escaneen — justo lo que el sistema necesita |
| **Fee de gobernanza** | Monto mensual por administrar el sistema | Refleja lo que realmente se entrega en la [Oferta 4](./oferta-systems.md#oferta-4--gobernanza-del-sistema-nivel-5) | Difícil de anclar sin referencia; se negocia caso por caso |
| **Participación en ahorro** | % del ahorro medido contra línea base | Alineación perfecta; el cliente no arriesga | Exige acuerdo sobre cómo se mide el ahorro — fuente garantizada de disputa; y flujo impredecible para CHACONTAINER |

**Recomendación:** por activo/mes como base para las Ofertas 2 y 3, y fee de
gobernanza para la Oferta 4. La participación en ahorro es atractiva
conceptualmente pero requiere una línea base indiscutible y una relación
madura; conviene reservarla, si acaso, como componente adicional en cuentas
grandes y nunca como el modelo principal `[VALIDAR con el fundador]`.

## 2. El incentivo perverso que hay que resolver

Si se cobra por activo bajo administración, **CHACONTAINER gana más cuando
el cliente tiene más activos**. Pero el valor que Systems entrega es
precisamente que el cliente **necesite menos activos**: mejor rotación,
menos pérdida y menor tiempo de ciclo significan que el mismo flujo se
atiende con un parque más chico.

Es decir: si Systems funciona muy bien, el ingreso recurrente baja. El
modelo se paga por fracasar parcialmente.

Tres formas de corregirlo, en orden de preferencia:

1. **Cobrar por activo administrado con piso contractual.** El fee se calcula
   sobre el parque al inicio del contrato y no baja durante la vigencia,
   aunque el parque se reduzca. El cliente captura el ahorro (menos compra
   de reposición) y CHACONTAINER no se penaliza por lograrlo. Al renovar se
   recalcula.
2. **Cobrar por ciclo gestionado en vez de por activo en existencia.** Alinea
   el ingreso con la actividad real y es neutral al tamaño del parque. Más
   complejo de facturar y de pronosticar.
3. **Fee de gobernanza fijo**, desacoplado del conteo. Simple y sin
   perversión, pero pierde la escalabilidad automática.

La opción 1 es la más simple de vender y la que menos fricción genera en la
negociación `[VALIDAR estructura de piso y vigencia]`.

## 3. La pregunta incómoda: ¿Systems canibaliza a Solutions?

Parcialmente sí, y conviene tenerlo pensado antes de que lo pregunte alguien
internamente.

Si el cliente pierde menos activos, compra menos reposición — y la venta de
reposición es ingreso de Solutions. Pero la relación no es de suma cero:

- **Lo que se pierde** es la venta de reposición por pérdida, que es
  ingreso no recurrente, impredecible y que el cliente resiente.
- **Lo que se gana** es ingreso recurrente contratado, más el volumen de
  Solutions que Systems *genera*: cada alerta de activo detenido dispara
  logística inversa; cada retorno dispara lavado e inspección; cada
  hallazgo de daño dispara reparación ([oferta-systems.md §4](./oferta-systems.md#4-qué-mantiene-vivo-a-solutions)).

El efecto neto esperado es sustituir ingreso transaccional volátil por
ingreso recurrente más ingreso transaccional *inducido y predecible*. Vale
la pena medirlo explícitamente durante el primer año `[VALIDAR: definir el
indicador de mezcla Solutions/Systems por cuenta]`.

## 4. Estructura por oferta

| Oferta | Estructura de cobro | Naturaleza |
|---|---|---|
| [1 · Diagnóstico de parque](./oferta-systems.md#oferta-1--diagnóstico-de-parque-nivel-3-puerta-de-entrada) | Precio cerrado por proyecto, en función del volumen de activos y ubicaciones a inventariar | Transaccional |
| [2 · Identificación y registro](./oferta-systems.md#oferta-2--identificación-y-registro-nivel-3) | Precio por activo identificado (una vez) + fee de plataforma por activo/mes | Mixta: setup + recurrente |
| [3 · Control y trazabilidad](./oferta-systems.md#oferta-3--control-y-trazabilidad-nivel-4) | Fee por activo/mes (mayor que Oferta 2) + logística inversa facturada por evento o incluida hasta un tope | Recurrente + variable |
| [4 · Gobernanza](./oferta-systems.md#oferta-4--gobernanza-del-sistema-nivel-5) | Fee mensual de gobernanza + fee por activo/mes | Recurrente |

Dos decisiones de diseño que importan más que el nivel de las tarifas:

- **El setup de identificación se cobra aparte y por adelantado.** Es trabajo real en sitio con costo real; regalarlo para "ganar el contrato" convierte cada alta de cliente en una pérdida de caja que solo se recupera si el contrato dura.
- **La logística inversa necesita un tope explícito** si se incluye en el fee. Sin tope, un cliente con custodios muy morosos consume margen ilimitado — y es justo el cliente al que más le sirve el servicio.

## 5. Lo que falta decidir antes de cotizar

Ninguna cotización de Systems debe salir sin que esto esté resuelto y
cargado en [`lista-precios-costos.md`](../chacontainer/docs/lista-precios-costos.md):

1. Costo real de identificar un activo en sitio (mano de obra + etiqueta + traslado).
2. Costo mensual de operar la plataforma por activo (infraestructura + soporte).
3. Horas mensuales que consume el ritual de gobernanza de la Oferta 4.
4. Vigencia mínima de contrato que hace rentable el setup.
5. Tope de logística inversa incluido, y tarifa del excedente.
6. Descuento máximo autorizado sin aprobación (el SOP-07 comercial ya prevé este rango, hoy en `[VALIDAR]`).

Los puntos 1, 2 y 3 son costos, no precios: hasta no tenerlos, cualquier
tarifa es una apuesta sobre el margen.

---

## Próximo paso

[Diagnóstico comercial](./diagnostico-comercial.md) (paso 18): el instrumento
que abre la conversación donde estas ofertas se presentan.

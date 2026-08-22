# Piloto pagado — estructura comercial

**Fase:** IV · Comercial — paso 19 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

[piloto.md](./piloto.md) (paso 15) define **cómo se ejecuta** el piloto:
selección, cronograma, medición, criterios de éxito. Este documento define
**cómo se vende y se contrata**: por qué se cobra, qué incluye el precio,
qué se promete y qué pasa al final.

---

## 1. Por qué el piloto se cobra

La tentación de regalar el piloto para "quitar fricción" es fuerte y es un
error. Tres razones concretas:

1. **Un piloto gratis no tiene dueño del lado del cliente.** Lo que no se
   paga no se defiende internamente: cuando el operador del cliente esté
   ocupado, el escaneo será lo primero que se caiga, y el criterio de
   adopción —el que más pilotos reprueba
   ([piloto.md §4](./piloto.md#4-criterios-de-éxito))— fallará por falta de
   presión interna, no por el producto.
2. **El trabajo del piloto es trabajo real.** Inventariar 300 activos e
   identificarlos es mano de obra en sitio, etiquetas y traslados. Regalarlo
   convierte cada piloto en una pérdida de caja que solo se recupera si el
   contrato posterior se firma y dura.
3. **El precio es información.** Un cliente que no está dispuesto a pagar por
   medir su parque tampoco pagará por administrarlo. Es más barato
   descubrirlo en la propuesta que en la semana 12.

## 2. Estructura autofinanciable

El piloto se cotiza como la suma de trabajo que tiene valor **aunque el
piloto fracase**, más el componente que está a prueba:

| Componente | Qué es | ¿Valor si el piloto fracasa? |
|---|---|---|
| Inventario físico y clasificación | [Oferta 1](./oferta-systems.md#oferta-1--diagnóstico-de-parque-nivel-3-puerta-de-entrada) | **Sí** — el cliente se queda con su conteo real y sus faltantes identificados |
| Identificación QR del parque | [Oferta 2](./oferta-systems.md#oferta-2--identificación-y-registro-nivel-3), setup | **Sí** — el parque queda identificado y registrado, sirva o no el resto |
| Plataforma y gobernanza durante 90 días | [Oferta 3](./oferta-systems.md#oferta-3--control-y-trazabilidad-nivel-4), temporal | **No** — esto es lo que está a prueba |

Los dos primeros componentes cubren el costo directo del piloto; el tercero
se cotiza al fee recurrente que tendría el contrato, prorrateado a la
duración. Así CHACONTAINER no financia la prueba y el cliente no compra una
promesa: compra dos entregables tangibles más el acceso temporal a lo que se
quiere demostrar.

Cifras y tarifas: en
[`lista-precios-costos.md`](../chacontainer/docs/lista-precios-costos.md),
no aquí ([pricing.md](./pricing.md)).

## 3. Reversión de riesgo: qué ofrecer y qué no

**No ofrecer "devolución si no funciona".** Suena potente y es una trampa:
"funcionar" no está definido, y al final se estaría devolviendo dinero por
un inventario y una identificación que sí se entregaron y que el cliente
conserva. La discusión terminaría en si el fracaso fue del sistema o de la
adopción del cliente — una pelea que no conviene tener con alguien con quien
se quiere firmar un contrato recurrente.

**Sí ofrecer, en orden de preferencia:**

1. **Umbral de mejora acordado por escrito antes de empezar**, sobre el
   indicador ligado al problema que el cliente declaró. Si no se alcanza,
   CHACONTAINER extiende el piloto sin costo adicional en lugar de devolver.
   **Decidido: 30 días de extensión.** Suficiente para un ajuste rápido si
   la causa es adopción (operadores no escaneando) más que diseño del
   sistema; si a los 30 días adicionales tampoco se alcanza el umbral, ahí
   sí corresponde el cierre con entrega descrito en
   [piloto.md §5](./piloto.md#5-conversión-a-contrato), no una segunda
   extensión. Convierte el riesgo en tiempo acotado, no en dinero ni en un
   compromiso indefinido.
2. **Acreditar parte del piloto al contrato.** **Decidido: un porcentaje
   fijo del total pagado en el piloto** (no solo del componente de
   plataforma) se acredita al primer período del contrato, si se firma
   dentro de 30-60 días del cierre. Es la opción más agresiva de las
   consideradas frente a acreditar solo el componente de plataforma/
   gobernanza — reduce más el margen del piloto si el cliente convierte,
   a cambio de un incentivo más claro para decidir rápido. Falta fijar el
   porcentaje exacto `[VALIDAR: % de acreditación]`; debe calcularse contra
   el margen real del piloto (§2) para no volver la conversión rápida más
   cara que no convertir.
3. **Sin permanencia después del piloto.** El contrato que sigue puede
   cancelarse con aviso razonable. Es barato de ofrecer —si el sistema
   funciona, nadie lo cancela— y elimina el miedo al amarre.

## 4. Lo que se promete y lo que no

| Se promete | No se promete |
|---|---|
| Entregar el conteo real del parque en alcance | Un porcentaje específico de recuperación de faltantes |
| Identificar el 100% del parque en alcance | Que todos los activos identificados aparezcan |
| Operar alertas y recolección durante la vigencia | Eliminar las pérdidas |
| Un reporte de cierre con la comparación honesta | Que la comparación sea favorable |
| Un responsable nombrado del lado de CHACONTAINER | Disponibilidad 24/7 |

La columna derecha importa tanto como la izquierda, y conviene decirla en
voz alta durante la propuesta. Un cliente al que se le prometió recuperar
faltantes y recupera pocos considerará fracasado un piloto que fue
técnicamente exitoso.

## 5. El contrato define de antemano las tres salidas

La decisión difícil no es firmar el piloto: es qué pasa después. Dejarlo
abierto garantiza una negociación incómoda en la semana 13, justo cuando el
equipo del cliente está cansado del proyecto.

Las tres salidas de [piloto.md §5](./piloto.md#5-conversión-a-contrato)
—contrato de gobernanza, extensión por adopción, o cierre con entrega— deben
quedar escritas en el contrato del piloto, cada una con su condición
objetiva de activación y su precio ya definido. El resultado se lee contra
lo acordado; no se negocia desde cero.

## 6. A quién ofrecerlo

El piloto no es para cualquier cliente que pase el
[diagnóstico comercial](./diagnostico-comercial.md). Además del problema
cuantificado, hacen falta tres condiciones:

1. **Un patrocinador con presupuesto y dolor propio.** No un interesado: alguien a quien el problema le afecte el indicador.
2. **Un responsable operativo asignado del lado del cliente.** Sin contraparte que empuje el escaneo internamente, el criterio de adopción falla ([piloto.md §4](./piloto.md#4-criterios-de-éxito)).
3. **Una ruta con al menos 2-3 ciclos completos dentro de la ventana.** Restricción operativa, no comercial, pero elimina cuentas donde el piloto no podría concluir nada ([piloto.md §2](./piloto.md#2-criterios-de-selección)).

Un cliente que cumple 1 y 3 pero no 2 no está descartado: la conversación es
sobre conseguir esa persona, y esa conversación es más útil tenerla antes de
firmar que en la semana 4.

---

## Próximo paso

[Conversión Solutions → Systems](./conversion-solutions-systems.md)
(paso 20): cómo se detecta y se activa sistemáticamente esta oportunidad
desde la operación que ya existe.

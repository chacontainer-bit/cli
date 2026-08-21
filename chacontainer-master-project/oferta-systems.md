# Oferta Packaging Systems — qué se vende exactamente

**Fase:** IV · Comercial — paso 16 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

Solutions se vende solo: el cliente ve el activo lavado, reparado o
entregado. Systems no tiene esa ventaja — nadie ve la gobernanza. Este
documento convierte las 16 capacidades del
[catálogo Systems](./catalogo-systems.md) en un número corto de ofertas
nombradas, con alcance cerrado, para que el equipo comercial pueda decir
"esto es lo que te vendo" en vez de describir una filosofía.

---

## 1. El problema de vender lo invisible

Tres consecuencias prácticas de que Systems sea intangible, que la oferta
tiene que resolver por diseño:

1. **El cliente no puede evaluar lo que no puede ver.** Por eso toda oferta
   de Systems arranca con un entregable físico o documental (un inventario,
   un parque identificado, un reporte) — no con acceso a software.
2. **El valor aparece con el tiempo, el costo aparece de inmediato.** Por eso
   existe el [piloto pagado](./piloto-pagado.md): comprime la demostración
   de valor a 60-90 días en vez de pedir un acto de fe anual.
3. **Nadie tiene presupuesto para "gobernanza".** Sí lo tienen para reducir
   compra de contenedores, para dejar de pagar paros de línea, o para
   cerrar un hallazgo de auditoría. La oferta se nombra por el problema que
   resuelve, no por la capacidad que entrega.

## 2. Las cuatro ofertas

Cada oferta agrupa capacidades del [catálogo Systems](./catalogo-systems.md)
y corresponde a un nivel de la [escalera comercial](./README.md#etapa-9--escalera-comercial).
No son planes de software: son alcances de servicio.

### Oferta 1 · Diagnóstico de parque *(Nivel 3, puerta de entrada)*

| | |
|---|---|
| **Problema que compra el cliente** | "No sé cuántos activos tengo ni dónde están" |
| **Capacidades** | [Diagnóstico](./catalogo-systems.md#1-diagnóstico-de-activos) + inventario físico y clasificación de Solutions |
| **Entregable** | Informe con conteo real vs. contable, faltantes, % sin identificar, condición por categoría, y estimación económica del costo de no saber |
| **Duración** | Semanas, no meses |
| **Qué NO incluye** | Software, seguimiento continuo, identificación del parque |
| **Para qué sirve comercialmente** | Es la línea base del piloto ([piloto.md §1](./piloto.md#1-el-problema-de-la-línea-base)) y el instrumento que hace visible el problema |

### Oferta 2 · Identificación y registro *(Nivel 3)*

| | |
|---|---|
| **Problema que compra el cliente** | "Tengo el conteo pero no puedo darle seguimiento a nada" |
| **Capacidades** | [Registro](./catalogo-systems.md#2-registro), [Identificación](./catalogo-systems.md#3-identificación), [QR](./catalogo-systems.md#4-qr), [Inventario digital](./catalogo-systems.md#6-inventario-digital) |
| **Entregable** | Parque 100% identificado con QR, registro maestro cargado, inventario consultable |
| **Qué NO incluye** | Reglas, alertas, gobernanza. El cliente ve su parque; nadie lo administra por él |
| **Nota** | Es la oferta que más naturalmente sale de un servicio de lavado o inventario físico ([ETAPA 8](./README.md#etapa-8--modelo-comercial-solutions--systems)) |

### Oferta 3 · Control y trazabilidad *(Nivel 4)*

| | |
|---|---|
| **Problema que compra el cliente** | "Sé lo que tengo pero se me siguen quedando activos afuera" |
| **Capacidades** | Oferta 2 + [Trazabilidad](./catalogo-systems.md#7-trazabilidad), [Custodios](./catalogo-systems.md#8-gestión-de-custodios), [Reglas](./catalogo-systems.md#11-reglas-operativas), [Alertas](./catalogo-systems.md#12-alertas), [Logística inversa](./catalogo-systems.md#9-logística-inversa) |
| **Entregable** | Lo anterior + alertas operando y recolección ejecutada contra ellas |
| **Qué NO incluye** | La decisión de qué hacer con los patrones que revelan los datos — eso es Oferta 4 |
| **Nota** | Es el primer nivel con **recurrencia real**: hay trabajo mensual (atender alertas, recolectar) |

### Oferta 4 · Gobernanza del sistema *(Nivel 5)*

| | |
|---|---|
| **Problema que compra el cliente** | "Quiero que alguien responda por que esto funcione, no solo que me lo reporte" |
| **Capacidades** | Oferta 3 + [Incidencias](./catalogo-systems.md#10-gestión-de-incidencias), [Indicadores](./catalogo-systems.md#13-indicadores), [Analítica](./catalogo-systems.md#14-analítica), [Gobernanza](./catalogo-systems.md#15-gobernanza) |
| **Entregable** | Comité mensual, reporte de gobernanza, ajuste continuo de reglas, y un responsable nombrado del lado de CHACONTAINER ([modelo-de-gobernanza.md §3](./modelo-de-gobernanza.md#3-ritual-de-gobernanza-propuesto)) |
| **Qué la distingue de Oferta 3** | En la 3 vendemos que el cliente se entere; en la 4 vendemos que nosotros actuemos. Es la diferencia entre un tablero y un responsable |
| **Compromiso** | Es la única oferta donde CHACONTAINER puede comprometerse a un indicador, no solo a una actividad `[VALIDAR: si se aceptan compromisos sobre KPI y con qué límite de responsabilidad]` |

El **Nivel 6 (Plataforma)** de la escalera —el cliente usando CHACONTAINER OS
como infraestructura propia, con sus propios usuarios— no se ofrece todavía:
requiere autoservicio y multi-planta, explícitamente fuera del
[MVP](./mvp.md#5-lo-que-el-mvp-deliberadamente-no-resuelve).

## 3. Regla de secuencia

Las ofertas son acumulativas y **no deben venderse salteadas**. Vender
Oferta 3 a un cliente que no pasó por la 1 significa poner reglas y alertas
sobre un parque cuyo inventario real nadie confirmó: las alertas dispararán
sobre datos falsos y el cliente concluirá, con razón, que el sistema no
funciona.

La única excepción razonable es un cliente que ya tenga su parque
identificado y registrado por su cuenta `[VALIDAR: si se ha dado el caso]`,
donde la Oferta 1 se reduce a validar lo que ya existe.

## 4. Qué mantiene vivo a Solutions

Ninguna de estas ofertas reemplaza a Solutions: todas la consumen. La
Oferta 1 requiere inventario físico y clasificación; la 2 requiere
identificación en sitio; la 3 requiere logística inversa y recuperación; la
4 requiere lavado, reparación y disposición a demanda.

Esto es deliberado y conviene decirlo en la conversación comercial: contratar
Systems **no reduce** el gasto del cliente en Solutions — lo hace predecible
y lo redirige a lo que sí hace falta, en vez de a reponer activos perdidos.
La promesa no es "gastarás menos con nosotros", es "gastarás lo mismo o
menos en total, y sabrás en qué".

---

## Próximo paso

[Pricing](./pricing.md) (paso 17): cómo se cobra cada una de estas cuatro
ofertas.

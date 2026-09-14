# Conversión Solutions → Systems — el mecanismo

**Fase:** IV · Comercial — paso 20 (último) de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

La [ETAPA 8](./README.md#etapa-8--modelo-comercial-solutions--systems) narra
la conversión como una secuencia natural: el cliente pide lavado, se detectan
activos sin identificación, se propone identificación, aparecen faltantes, se
necesita trazabilidad. La lógica es correcta y aun así la conversión no
ocurre sola. Este documento explica por qué, y qué hay que instalar para que
ocurra.

---

## 1. Por qué la secuencia de la ETAPA 8 no se activa sola

La narrativa asume un observador que no existe. Quien lava el activo y nota
que no tiene QR es un operador de planta: no es comercial, no tiene a quién
avisarle, no gana nada por avisar, y tiene una cuota de lavado que cumplir
ese turno. La oportunidad se detecta y se pierde en el mismo minuto.

Tres condiciones para que la conversión sea sistemática en vez de anecdótica:

1. **La detección tiene que ser un subproducto de hacer bien el SOP**, no una tarea adicional de estar atento.
2. **La señal tiene que viajar sola** hasta alguien que pueda actuar, sin depender de que el operador tome la iniciativa.
3. **Alguien tiene que ser dueño** de revisar esas señales con una cadencia fija.

Las tres son instalables. Ninguna requiere que el operador se vuelva
vendedor.

## 2. Señales de conversión, por servicio de Solutions

Cada servicio produce, al ejecutarse correctamente, un dato que revela una
brecha de sistema. La columna clave es la última: dónde queda capturado ese
dato sin trabajo extra.

| Servicio | Señal | Qué revela | Dónde se captura |
|---|---|---|---|
| [Lavado](./catalogo-solutions.md#3-lavado) | % de activos del lote sin QR legible | Parque sin identificar | [SOP-ACTIVO-05](./sop/SOP-ACTIVO-05-lavado.md) — el registro de ingreso ya requiere identificar el activo |
| [Inspección](./catalogo-solutions.md#9-inspección) | Tasa de rechazo muy superior a la esperada | Falta de criterio de condición y de mantenimiento preventivo | [SOP-ACTIVO-03](./sop/SOP-ACTIVO-03-inspeccion.md) |
| [Clasificación](./catalogo-solutions.md#10-clasificación) | Lote con alta dispersión de propietarios o condición | Mezcla de parques, sin control de custodia | [SOP-ACTIVO-04](./sop/SOP-ACTIVO-04-clasificacion-de-condicion.md) |
| [Inventarios físicos](./catalogo-solutions.md#11-inventarios-físicos) | Discrepancia físico-contable | Inventario no confiable | [SOP-ACTIVO-09](./sop/SOP-ACTIVO-09-inventario.md) |
| [Recuperación](./catalogo-solutions.md#12-recuperación) | Activos detenidos mucho tiempo antes de recuperarse | Sin alertas ni logística inversa | [SOP-ACTIVO-14](./sop/SOP-ACTIVO-14-logistica-inversa.md) |
| [Reparación](./catalogo-solutions.md#4-reparación) | Reincidencia del mismo daño en la misma familia | Sin analítica de causa | [SOP-ACTIVO-06](./sop/SOP-ACTIVO-06-reparacion.md) |
| [Renta](./catalogo-solutions.md#2-renta) | Retornos fuera de plazo de forma recurrente | Sin reglas de retorno ni gestión de custodios | [SOP-ACTIVO-13](./sop/SOP-ACTIVO-13-retorno.md) |

Esta tabla es la [lectura cruzada del catálogo Solutions](./catalogo-solutions.md#lectura-cruzada-de-dónde-entra-cada-servicio-a-systems),
ahora con el campo que faltaba: **dónde vive el dato**.

## 3. El umbral, no la anécdota

Un activo sin QR no es una oportunidad comercial: es martes. Lo que
califica es la proporción y la persistencia.

Cada señal necesita un umbral definido —*"más del X% del lote sin
identificación, en dos servicios consecutivos del mismo cliente"*— para
distinguir un caso aislado de un patrón `[VALIDAR umbral por señal con el
responsable comercial]`. Sin umbral, o se reporta todo y nadie lo lee, o no
se reporta nada.

El principio es el mismo que gobierna las
[alertas operativas](./alertas.md#1-el-riesgo-real-el-ruido-no-la-cobertura):
una señal que se dispara siempre deja de ser una señal.

## 4. Quién actúa

| Rol | Responsabilidad | Cadencia |
|---|---|---|
| Operador de planta | Ninguna adicional. Ejecuta su SOP; el dato queda capturado como parte del registro normal | — |
| Responsable de planta | Revisa las señales que superaron umbral en sus servicios del período y las marca como oportunidad | Semanal |
| Responsable comercial | Recibe las oportunidades marcadas, decide cuáles ameritan [diagnóstico comercial](./diagnostico-comercial.md) y lo agenda | Semanal |
| Gobernanza | Revisa la tasa de conversión del embudo y ajusta umbrales | Mensual, en el [comité](./modelo-de-gobernanza.md#3-ritual-de-gobernanza-propuesto) |

El eslabón que suele faltar es el segundo. Sin una revisión semanal
explícita —cinco minutos mirando qué señales se dispararon— la información
queda en la base de datos y no llega a nadie.

## 5. El embudo

```
Servicio de Solutions ejecutado
        ↓  (el SOP captura el dato)
Señal supera umbral
        ↓  (revisión semanal del responsable de planta)
Oportunidad marcada
        ↓  (decisión comercial)
Diagnóstico comercial          ← gratuito, conversación
        ↓  (cierre: "propongo medir")
Oferta 1 · Diagnóstico de parque  ← primer ingreso de Systems
        ↓  (línea base construida)
Piloto pagado                  ← 60-90 días
        ↓  (tres salidas ya contratadas)
Contrato de gobernanza          ← recurrencia
```

Vale la pena medir la conversión de cada escalón por separado. Una caída
entre "señal supera umbral" y "oportunidad marcada" es un problema de
proceso interno; una caída entre "diagnóstico comercial" y "Oferta 1" es un
problema de propuesta o de calificación del cliente. Son fallas distintas
con remedios distintos, y el embudo agregado las esconde `[VALIDAR: definir
la meta de conversión por escalón una vez haya volumen]`.

## 6. Lo que este mecanismo no debe convertirse en

**En presión de venta sobre la operación.** Si el operador percibe que su
registro de "activo sin QR" se traduce en presión comercial hacia un cliente
con el que él trata todos los días, empezará a no registrarlo. La captura
del dato tiene que seguir siendo neutral: se registra porque el SOP lo pide
para operar bien, no para generar leads.

Esa neutralidad es lo que hace sostenible el mecanismo. En el momento en que
el dato operativo se contamina con intención comercial, deja de ser
confiable — y con él se cae también la base del sistema, porque son los
mismos registros que alimentan
[trazabilidad, estados y KPI](./modelo-de-datos.md).

---

## Cierre de FASE IV

Con este documento se completan los 5 pasos de la FASE IV · Comercial:
[oferta Systems](./oferta-systems.md) (16), [pricing](./pricing.md) (17),
[diagnóstico comercial](./diagnostico-comercial.md) (18),
[piloto pagado](./piloto-pagado.md) (19) y esta conversión (20).

Sigue la **FASE V · Escala** ([ETAPA 14](./README.md#etapa-14--orden-de-ejecución)):
hubs, partners, RFID, integraciones ERP/MES, multi-planta, operación
nacional y operación México–Estados Unidos.

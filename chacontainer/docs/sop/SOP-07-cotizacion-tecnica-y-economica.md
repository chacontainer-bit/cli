# SOP-07: Cotización técnica y económica

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Prompt maestro del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) al paso 9 del proceso comercial ideal de CHACONTAINER: convertir el diseño de la solución (paso 8) en una cotización formal con tiempos, exclusiones y condiciones, antes de la negociación y cierre (paso 10). Se construyó a partir de prácticas estándar de cotización B2B industrial, **no a partir de una entrevista real con quien hoy elabora las cotizaciones en CHACONTAINER** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6 de la metodología). Los puntos marcados como `[VALIDAR]` requieren confirmación directa del responsable comercial/financiero.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Traducir el diseño de la solución en una propuesta formal de precio y condiciones —técnica y económica— que el cliente pueda evaluar y aprobar, incluyendo tiempos, exclusiones y condiciones claras.
2. **¿Qué problema busca resolver?** Evitar cotizaciones ambiguas o incompletas que generan negociaciones interminables o malentendidos post-venta sobre qué incluye y qué no; evitar inconsistencia de precios entre clientes similares; evitar que la cotización dependa de que el fundador calcule cada caso desde cero.
3. **¿Quién es el cliente?** Interno: quien negocia y cierra (paso 10) y quien implementa (paso 11), que necesitan que lo cotizado sea exactamente lo que se entrega. Externo: el prospecto que recibirá la propuesta formal.
4. **¿Cuáles son las entradas?** El documento de diseño de la solución (paso 8), la lista de precios y costos base de CHACONTAINER `[VALIDAR lista exacta]`, la política de condiciones comerciales (plazos de pago, garantías, vigencia de la oferta) `[VALIDAR]`, y una plantilla de cotización.
5. **¿Cuáles son las salidas esperadas?** Documento de cotización técnica y económica: especificación técnica de lo cotizado, precio, tiempos de entrega/implementación, exclusiones explícitas, condiciones comerciales y vigencia de la oferta.
6. **¿Quién es el responsable?** `[VALIDAR]` — ejecutivo comercial, posiblemente con revisión de finanzas o del fundador para aprobar precios o condiciones fuera del rango estándar.
7. **¿Qué indicadores dicen que el proceso funciona?** Tiempo entre diseño de solución completado y cotización entregada al cliente, % de cotizaciones que se cierran sin renegociación por ambigüedad de alcance, tasa de conversión de cotización a cierre.
8. **¿Qué riesgos existen?** Disputas post-venta por exclusiones no claras; inconsistencia de precios entre clientes similares; cuello de botella por depender del fundador para aprobar cada cotización; tiempos de entrega prometidos que luego no se cumplen.
9. **¿Qué actividades no agregan valor?** Recalcular precios desde cero para cada cliente sin una base de costos estandarizada; cotizaciones verbales o informales sin documento formal; reescribir la cotización varias veces por errores de formato en vez de contenido.
10. **¿Qué partes pueden estandarizarse?** La plantilla de cotización (técnica, económica, tiempos, exclusiones, condiciones, vigencia), la lista de precios y costos base por tipo de activo y servicio, y la política de condiciones comerciales.
11. **¿Qué partes pueden automatizarse?** El cálculo del precio a partir de la especificación de la solución y la lista de precios base, la generación del documento de cotización a partir de una plantilla, y el control de vigencia de la oferta.
12. **¿Qué depende únicamente del fundador?** `[VALIDAR]` — posiblemente la aprobación de descuentos o condiciones fuera del estándar, y el criterio de precio en casos no cubiertos por la lista base.

---

## 1. Resumen ejecutivo

La cotización técnica y económica convierte el diseño de la solución en una propuesta formal que el cliente puede evaluar y aprobar: qué se entrega, a qué precio, en qué tiempos, qué queda explícitamente excluido y bajo qué condiciones comerciales. Es el documento que sostiene la negociación (paso 10) y que, si es ambiguo, genera disputas o renegociaciones que podrían evitarse.

## 2. Objetivo

Entregar al cliente una cotización clara y completa —técnica y económica— con tiempos, exclusiones y condiciones explícitas, basada en el diseño de la solución ya aprobado internamente, sin depender de un cálculo improvisado por caso.

## 3. Alcance

Aplica desde que el diseño de la solución (paso 8) está completo hasta que la cotización queda enviada formalmente al cliente. No incluye la negociación de condiciones ni el cierre del negocio (paso 10).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Ejecutivo comercial | Calcula el precio con la lista base, define tiempos y exclusiones, y consolida el documento. |
| Responsable financiero/fundador `[VALIDAR]` | Aprueba precios o condiciones que se salen del rango estándar. |
| Fundador (rol transitorio) | Aporta hoy el criterio de precio en casos no cubiertos por la lista base; el objetivo del proceso es reducir esta dependencia con una lista de precios documentada. |

## 5. Diagrama de flujo

```
[Diseño de la solución completo (paso 8)]
                ↓
[Calcular precio con la lista de costos/precios base]
                ↓
[Definir tiempos de entrega/implementación]
                ↓
[Definir exclusiones explícitas]
                ↓
[Definir condiciones comerciales (pago, garantía, vigencia)]
                ↓
     ¿Sale del rango estándar?
     ↓ sí                          ↓ no
[Aprobación de finanzas/fundador]   │
     ↓                             │
     └─────────────┬───────────────┘
                    ↓
[Consolidar documento de cotización técnica y económica]
                    ↓
[Enviar cotización al cliente]
                    ↓
[Entrega a Negociación y cierre (paso 10)]
```

## 6. SOP paso a paso

1. Confirmar que el diseño de la solución (paso 8) está completo y documentado.
2. Calcular el precio de la solución usando la lista de precios y costos base de CHACONTAINER `[VALIDAR lista exacta]`.
3. Definir los tiempos de entrega o implementación, verificados contra la disponibilidad real de inventario o capacidad, no una estimación optimista.
4. Definir explícitamente las exclusiones (qué no incluye la propuesta): instalación, transporte, mantenimiento, u otras, según aplique.
5. Definir las condiciones comerciales: forma y plazo de pago, garantías, vigencia de la oferta.
6. Si el precio o las condiciones requieren salirse del estándar (descuentos, plazos especiales), obtener la aprobación correspondiente antes de enviar `[VALIDAR quién aprueba]`.
7. Consolidar toda la información en el documento estándar de cotización técnica y económica.
8. Enviar la cotización al cliente y registrar la fecha de envío y la vigencia de la oferta.
9. Entregar el registro a Negociación y cierre (paso 10).

## 7. Checklist

- [ ] Diseño de la solución (paso 8) completo antes de cotizar.
- [ ] Precio calculado con la lista de precios/costos base, no "a ojo".
- [ ] Tiempos de entrega/implementación verificados contra disponibilidad real.
- [ ] Exclusiones explícitas documentadas (no implícitas).
- [ ] Condiciones comerciales (pago, garantía, vigencia) documentadas.
- [ ] Aprobación obtenida si la cotización sale del estándar (descuentos, condiciones especiales).
- [ ] Cotización enviada como documento formal, no de forma verbal.
- [ ] Fecha de envío y vigencia de la oferta registradas.

## 8. KPI

- Tiempo entre diseño de solución completado y cotización entregada al cliente.
- % de cotizaciones que se cierran sin renegociación por ambigüedad de alcance.
- Tasa de conversión de cotización a cierre.
- % de cotizaciones con exclusiones y condiciones documentadas explícitamente (auditoría interna).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Disputas post-venta por exclusiones no claras | Exclusiones explícitas obligatorias en la plantilla de cotización. |
| Inconsistencia de precios entre clientes similares | Lista de precios/costos base estandarizada, con aprobación explícita para desviaciones. |
| Cuello de botella por depender del fundador para aprobar cada cotización | Definir un rango estándar que el comercial puede cotizar sin aprobación, y escalar solo las excepciones. |
| Tiempos de entrega no realistas prometidos al cliente | Verificar disponibilidad real de inventario/capacidad antes de comprometer una fecha. |

## 10. Mejoras

- Crear y mantener una lista de precios/costos base por tipo de activo y servicio, para no calcular cada cotización desde cero.
- Definir un rango de aprobación estándar dentro del cual el comercial puede cotizar solo, y escalar únicamente las excepciones.
- Estandarizar la plantilla de cotización con secciones fijas de exclusiones y condiciones, para eliminar la ambigüedad.

## 11. Recomendaciones de automatización

- Calcular automáticamente el precio a partir de la especificación de la solución (paso 8) y la lista de precios base.
- Generar el documento de cotización a partir de una plantilla, reutilizando los datos ya capturados en el diseño de la solución.
- Alertas automáticas de vencimiento de la vigencia de la oferta.
- **No automatizar la aprobación de descuentos o condiciones fuera de estándar**: requiere criterio humano hasta documentar y validar los rangos de aprobación (Paso 6 de la metodología).

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar a quien hoy elabora las cotizaciones (probablemente el fundador) para validar la lista de precios real, los rangos de aprobación y las condiciones comerciales estándar, y probar el proceso en la próxima cotización antes de declararlo estándar.

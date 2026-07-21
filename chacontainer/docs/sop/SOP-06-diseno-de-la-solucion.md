# SOP-06: Diseño de la solución

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Prompt maestro del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) al paso 8 del proceso comercial ideal de CHACONTAINER: seleccionar el servicio adecuado (venta, renta, reparación, trazabilidad o sistemas) a partir del diagnóstico consultivo (paso 5) y el levantamiento técnico (paso 7), antes de cotizar (paso 9). Se construyó a partir de prácticas estándar de diseño de soluciones B2B industrial, **no a partir de una entrevista real con quien hoy diseña estas soluciones en CHACONTAINER** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6 de la metodología). Los puntos marcados como `[VALIDAR]` requieren confirmación directa del responsable comercial/técnico.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Traducir el diagnóstico consultivo y el levantamiento técnico en una solución concreta y justificada —qué servicio (venta, renta, reparación, trazabilidad o sistemas), qué configuración y en qué volumen— antes de cotizar.
2. **¿Qué problema busca resolver?** Evitar ofrecer el mismo producto o servicio "por defecto" a todos los clientes sin conectarlo con lo diagnosticado y medido; evitar recotizaciones por mal diseño; evitar que la decisión de qué vender dependa solo del criterio informal del fundador.
3. **¿Quién es el cliente?** Interno: quien cotiza (paso 9) y quien implementa (paso 11). Externo: el prospecto cuyo diagnóstico y levantamiento ya están completos.
4. **¿Cuáles son las entradas?** El documento de diagnóstico consultivo (paso 5), el documento de levantamiento técnico (paso 7), el catálogo de servicios de CHACONTAINER (venta, renta, reparación, trazabilidad, sistemas) `[VALIDAR catálogo exacto]`, y los criterios de decisión para elegir el servicio adecuado `[VALIDAR]`.
5. **¿Cuáles son las salidas esperadas?** Documento de diseño de solución: servicio(s) seleccionado(s), justificación basada en el diagnóstico y el levantamiento, y especificación (tipo de activo, cantidad, configuración logística).
6. **¿Quién es el responsable?** `[VALIDAR]` — probablemente el mismo ejecutivo comercial/técnico, o un rol de "ingeniería de soluciones" `[VALIDAR si existe hoy]`.
7. **¿Qué indicadores dicen que el proceso funciona?** % de soluciones diseñadas que se cotizan sin retrabajo, % de propuestas aceptadas por el cliente en la negociación, tiempo entre levantamiento técnico completado y diseño de solución entregado.
8. **¿Qué riesgos existen?** Elegir el servicio equivocado (por ejemplo, ofrecer venta cuando el cliente necesitaba renta) por no conectar bien diagnóstico y levantamiento; diseñar una solución no viable logísticamente con las restricciones detectadas; depender del criterio no documentado del fundador para decidir qué ofrecer.
9. **¿Qué actividades no agregan valor?** Diseñar sin revisar completos el diagnóstico y el levantamiento; repetir el mismo tipo de solución para todos los clientes sin analizar si aplica; diseñar y cotizar al mismo tiempo sin dejar registro de por qué se eligió esa solución.
10. **¿Qué partes pueden estandarizarse?** Un árbol de decisión o criterios claros para elegir entre venta, renta, reparación, trazabilidad y sistemas según el perfil del cliente, y una plantilla del documento de diseño de solución.
11. **¿Qué partes pueden automatizarse?** Un cuestionario de apoyo a la decisión que, a partir de los datos ya capturados en diagnóstico y levantamiento, sugiera el/los servicio(s) más adecuados; la generación semi-automática del documento de diseño a partir de esos mismos datos.
12. **¿Qué depende únicamente del fundador?** `[VALIDAR]` — posiblemente el criterio final para casos ambiguos o soluciones combinadas/no estándar, y el conocimiento de qué ha funcionado en clientes similares.

---

## 1. Resumen ejecutivo

El diseño de la solución es el paso donde CHACONTAINER convierte el diagnóstico consultivo y el levantamiento técnico en una propuesta concreta: qué servicio ofrecer (venta, renta, reparación, trazabilidad o sistemas), en qué configuración y volumen, y por qué esa es la solución correcta para ese cliente en particular. Es la bisagra entre "entender el problema" y "cotizarlo"; diseñar mal aquí obliga a recotizar o a ajustar la solución ya en implementación.

## 2. Objetivo

Seleccionar y especificar el servicio o combinación de servicios que mejor resuelve el problema diagnosticado del cliente, dentro de las restricciones técnicas encontradas, dejando registrada la justificación de la decisión antes de cotizar.

## 3. Alcance

Aplica desde que el diagnóstico consultivo (paso 5) y el levantamiento técnico (paso 7) están completos, hasta que queda documentada la solución seleccionada y especificada. No incluye la elaboración de precios y condiciones (cotización, paso 9) ni la negociación (paso 10).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Responsable de diseño de solución `[VALIDAR]` | Revisa diagnóstico y levantamiento, aplica los criterios de decisión y documenta la solución. |
| Ejecutivo comercial | Valida que la solución diseñada sea comunicable y vendible al cliente antes de pasar a cotización. |
| Fundador (rol transitorio) | Aporta hoy el criterio final en casos ambiguos o soluciones no estándar; el objetivo del proceso es documentar ese criterio en un árbol de decisión. |

## 5. Diagrama de flujo

```
[Diagnóstico consultivo + Levantamiento técnico completos]
                ↓
[Revisar ambos documentos en conjunto]
                ↓
[Aplicar criterios de decisión:
 venta / renta / reparación / trazabilidad / sistemas]
                ↓
[Seleccionar el/los servicio(s) adecuado(s)]
                ↓
[Especificar la solución: tipo de activo,
 cantidad, configuración logística]
                ↓
[Validar viabilidad contra las restricciones del levantamiento]
                ↓
[Documentar el diseño de la solución con su justificación]
                ↓
[Entrega a Cotización técnica y económica (paso 9)]
```

## 6. SOP paso a paso

1. Confirmar que el diagnóstico consultivo (paso 5) y el levantamiento técnico (paso 7) están completos y disponibles.
2. Revisar ambos documentos en conjunto: necesidades, objetivos y costos del diagnóstico, junto con medidas, cantidades y restricciones del levantamiento.
3. Aplicar los criterios de decisión para elegir entre los servicios disponibles: venta, renta, reparación, trazabilidad o sistemas `[VALIDAR criterios exactos]`.
4. Seleccionar el o los servicios que mejor resuelven el problema diagnosticado dentro de las restricciones técnicas encontradas.
5. Especificar la solución: tipo de activo, cantidad o volumen, y configuración logística (frecuencia de reposición, ciclo de retorno si aplica).
6. Validar internamente que la solución es viable frente a las restricciones detectadas en el levantamiento (peso, altura, acceso, logística).
7. Documentar el diseño de la solución junto con la justificación basada en el diagnóstico y el levantamiento: por qué esta solución y no otra.
8. Entregar el documento de diseño de solución a Cotización técnica y económica (paso 9).

## 7. Checklist

- [ ] Diagnóstico consultivo y levantamiento técnico revisados juntos antes de diseñar.
- [ ] Criterios de decisión aplicados explícitamente (no elegido "por defecto" o "por costumbre").
- [ ] Servicio(s) seleccionado(s) justificado(s) con datos del diagnóstico y del levantamiento.
- [ ] Especificación de la solución completa (tipo de activo, cantidad, configuración logística).
- [ ] Viabilidad de la solución validada contra las restricciones técnicas encontradas.
- [ ] Documento de diseño de solución registrado antes de pasar a cotización.

## 8. KPI

- % de soluciones diseñadas que se cotizan sin retrabajo.
- % de propuestas aceptadas por el cliente en la negociación (indicador indirecto de que la solución encajó).
- Tiempo entre levantamiento técnico completado y diseño de solución entregado.
- % de soluciones que requieren ajuste después de iniciada la implementación (señal de mal diseño).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Elegir el servicio equivocado (por ejemplo, venta en vez de renta) | Criterios de decisión documentados y aplicados explícitamente, no "por costumbre". |
| Solución no viable logísticamente con las restricciones detectadas | Validación obligatoria contra el levantamiento técnico antes de cerrar el diseño. |
| Dependencia del criterio no documentado del fundador | Documentar el árbol de decisión y los casos típicos ya resueltos. |
| Diseñar y cotizar sin dejar registro de por qué se eligió la solución | Documento de diseño de solución obligatorio, previo y separado de la cotización. |

## 10. Mejoras

- Documentar un árbol de decisión o criterios claros para elegir entre venta, renta, reparación, trazabilidad y sistemas, según el perfil y las restricciones del cliente.
- Separar explícitamente el diseño de la solución (qué se va a ofrecer y por qué) de la cotización (paso 9), aunque ambas las haga la misma persona.
- Registrar los casos ya resueltos (soluciones diseñadas exitosamente) como referencia para casos similares futuros.

## 11. Recomendaciones de automatización

- Cuestionario de apoyo a la decisión que, a partir de los datos ya capturados en diagnóstico y levantamiento, sugiera el/los servicio(s) más adecuados (como apoyo, no como reemplazo del criterio humano).
- Generación semi-automática del documento de diseño de solución a partir de los datos ya capturados en los pasos previos, evitando redigitar información.
- **No automatizar la decisión final del servicio a ofrecer en casos ambiguos o combinados**: requiere criterio humano hasta documentar y validar el árbol de decisión (Paso 6 de la metodología).

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar a quien hoy diseña las soluciones (probablemente el fundador) para documentar el criterio real de decisión entre servicios, y probarlo en el próximo caso antes de declararlo estándar.

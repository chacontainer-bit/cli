# SOP-03: Diagnóstico consultivo

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Prompt maestro del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) al paso 5 del proceso comercial ideal de CHACONTAINER: el diagnóstico consultivo que sigue al primer contacto (paso 4) y precede a la calificación como proveedor (paso 6) y al levantamiento técnico (paso 7). Se construyó a partir de prácticas estándar de venta consultiva B2B industrial, **no a partir de una entrevista real con quien hoy hace estos diagnósticos en CHACONTAINER** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6). Los puntos marcados como `[VALIDAR]` requieren confirmación directa del responsable comercial/técnico.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Entender a profundidad el sistema logístico y de empaque actual del prospecto —flujos, costos visibles e invisibles, riesgos y objetivos— para sustentar el levantamiento técnico y el diseño de la solución, sin vender todavía.
2. **¿Qué problema busca resolver?** Evitar diseñar o cotizar soluciones genéricas sin entender el contexto real del cliente; evitar ofrecer producto en lugar de resolver el problema de fondo; reducir el riesgo de propuestas que el cliente rechaza por no ajustarse a su operación real.
3. **¿Quién es el cliente?** Interno: el equipo que hará el levantamiento técnico (paso 7) y diseñará la solución (paso 8). Externo: el prospecto que ya tuvo primer contacto (paso 4) y acepta profundizar en una reunión de diagnóstico.
4. **¿Cuáles son las entradas?** Registro del prospecto calificado (SOP-01 + calificación inicial, paso 3), agenda de la reunión de diagnóstico confirmada en el primer contacto, guía de preguntas de diagnóstico `[VALIDAR guía exacta]`, acceso a la operación del cliente (visita presencial o información compartida remotamente).
5. **¿Cuáles son las salidas esperadas?** Documento de diagnóstico con el sistema logístico actual, el flujo de empaque, los costos visibles (compra, reposición) e invisibles (mermas, pérdidas de activos, tiempos muertos, daños), los riesgos identificados y los objetivos declarados por el cliente.
6. **¿Quién es el responsable?** `[VALIDAR]` — ejecutivo comercial y/o técnico que realiza la visita o reunión.
7. **¿Qué indicadores dicen que el proceso funciona?** % de diagnósticos que avanzan a levantamiento técnico o cotización, tiempo entre primer contacto y diagnóstico completado, % de diagnósticos que identifican al menos un costo invisible, tasa de cierre de negocios con diagnóstico consultivo completo frente a los que no lo tuvieron.
8. **¿Qué riesgos existen?** Vender antes de diagnosticar (se pierde la naturaleza consultiva y se genera desconfianza); diagnóstico superficial que no detecta costos invisibles reales; depender del criterio informal del fundador para interpretar lo que se observa; información sensible del cliente mal registrada o perdida.
9. **¿Qué actividades no agregan valor?** Reuniones de diagnóstico sin guía estructurada (se pierde información); reconstruir el diagnóstico de memoria días después en vez de registrarlo en el momento; repetir preguntas ya respondidas en el primer contacto.
10. **¿Qué partes pueden estandarizarse?** La guía de preguntas de diagnóstico, la plantilla del documento de diagnóstico, y el checklist de temas obligatorios (logística, empaque, costos, riesgos, objetivos).
11. **¿Qué partes pueden automatizarse?** El envío de la plantilla de diagnóstico al comercial antes de la reunión, la generación del documento de diagnóstico a partir de respuestas estructuradas, y el recordatorio de seguimiento si el diagnóstico no se registra a tiempo.
12. **¿Qué depende únicamente del fundador?** `[VALIDAR]` — hoy probablemente el criterio para identificar costos invisibles y riesgos, basado en su experiencia acumulada. Meta del proceso: transferirlo a una guía documentada que cualquier ejecutivo pueda aplicar igual.

---

## 1. Resumen ejecutivo

El diagnóstico consultivo es la reunión (o visita) en la que CHACONTAINER entiende, sin vender todavía, cómo opera realmente el prospecto: su sistema logístico, su flujo de empaque, sus costos visibles e invisibles, sus riesgos y sus objetivos. El resultado es un documento de diagnóstico que alimenta el levantamiento técnico (paso 7) y el diseño de la solución (paso 8), evitando propuestas genéricas que no encajan con la operación real del cliente.

## 2. Objetivo

Levantar un diagnóstico completo y consistente del sistema logístico y de empaque del prospecto, identificando costos visibles e invisibles, riesgos y objetivos, sin ofrecer producto ni precio, para sustentar las etapas posteriores del ciclo comercial.

## 3. Alcance

Aplica a la reunión o visita de diagnóstico posterior al primer contacto (paso 4), con prospectos que ya mostraron interés en profundizar. No incluye la calificación formal como proveedor (paso 6), el levantamiento técnico detallado con medidas y cantidades (paso 7), ni la cotización (paso 9).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Ejecutivo comercial/técnico `[VALIDAR]` | Prepara la guía de diagnóstico, realiza la reunión/visita y registra la información. |
| Responsable comercial `[VALIDAR]` | Revisa la calidad y completitud del diagnóstico antes de pasarlo a levantamiento técnico. |
| Fundador (rol transitorio) | Aporta hoy el criterio para interpretar costos invisibles y riesgos; el objetivo del proceso es documentar ese criterio para no depender de él. |

## 5. Diagrama de flujo

```
[Prospecto con reunión de diagnóstico agendada (paso 4)]
                ↓
[Preparar guía de preguntas de diagnóstico]
                ↓
[Realizar reunión/visita de diagnóstico
 (sin ofrecer producto ni precio)]
                ↓
[Registrar: sistema logístico, flujo de empaque,
 costos visibles e invisibles, riesgos, objetivos]
                ↓
[Consolidar documento de diagnóstico estándar]
                ↓
[Compartir con el equipo interno]
                ↓
[Entrega a Calificación como proveedor (paso 6)
 y/o Levantamiento técnico (paso 7)]
```

## 6. SOP paso a paso

1. Confirmar la agenda de la reunión de diagnóstico acordada en el primer contacto (paso 4).
2. Preparar la guía de preguntas de diagnóstico antes de la reunión, cubriendo: sistema logístico, flujo de empaque, costos visibles, costos invisibles, riesgos y objetivos del cliente `[VALIDAR guía exacta]`.
3. Realizar la reunión o visita, indagando abiertamente sin ofrecer producto ni precio todavía.
4. Registrar en el momento (o inmediatamente después) las respuestas usando la guía estructurada, evitando reconstruir la información de memoria días después.
5. Documentar los costos invisibles detectados (mermas, pérdidas de activos, tiempos muertos, daños) además de los costos visibles (compra, reposición).
6. Identificar riesgos operativos y los objetivos declarados por el cliente (por ejemplo, reducir costos, reducir accidentes, mejorar trazabilidad).
7. Consolidar toda la información en el documento de diagnóstico estándar.
8. Compartir el diagnóstico con el equipo que hará el levantamiento técnico (paso 7) y diseñará la solución (paso 8).
9. Entregar el registro a Calificación como proveedor (paso 6) si el cliente lo exige, o directamente a Levantamiento técnico (paso 7).

## 7. Checklist

- [ ] Guía de preguntas de diagnóstico usada (reunión no improvisada).
- [ ] Sistema logístico actual del cliente documentado.
- [ ] Flujo de empaque actual documentado.
- [ ] Costos visibles identificados.
- [ ] Costos invisibles identificados (mermas, pérdidas, tiempos muertos, daños).
- [ ] Riesgos operativos identificados.
- [ ] Objetivos del cliente registrados explícitamente.
- [ ] Documento de diagnóstico registrado el mismo día de la reunión.
- [ ] No se ofreció producto ni precio durante la reunión de diagnóstico.

## 8. KPI

- % de diagnósticos que avanzan a levantamiento técnico o cotización.
- Tiempo entre primer contacto y diagnóstico completado.
- % de diagnósticos que identifican al menos un costo invisible (no solo visibles).
- Tasa de cierre de negocios con diagnóstico consultivo completo frente a los que no lo tuvieron.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Vender antes de diagnosticar (se pierde la naturaleza consultiva) | Prohibir explícitamente hablar de precio o producto en esta reunión; eso se separa a la cotización (paso 9). |
| Diagnóstico superficial que no detecta costos invisibles | Guía de preguntas obligatoria que cubra explícitamente los costos invisibles. |
| Depender del criterio informal del fundador para interpretar hallazgos | Documentar la guía y los criterios típicos de costos/riesgos por industria. |
| Información del cliente mal registrada o perdida | Registrar el diagnóstico el mismo día en el documento estándar, no de memoria después. |

## 10. Mejoras

- Crear una guía de preguntas de diagnóstico documentada, por industria si aplica, para no depender de la intuición de quien hace la visita.
- Separar explícitamente el diagnóstico (entender) de la cotización (vender), incluso si ambas etapas ocurren con el mismo interlocutor del cliente.
- Estandarizar la plantilla del documento de diagnóstico para que cualquier ejecutivo la use de la misma forma.

## 11. Recomendaciones de automatización

- Enviar automáticamente la guía de diagnóstico y un recordatorio de preparación al comercial antes de cada reunión agendada.
- Generar el documento de diagnóstico a partir de un formulario estructurado en vez de notas libres, para reducir la carga de redacción y la pérdida de información.
- Recordatorio automático de seguimiento si el diagnóstico no se registra dentro de 24-48 horas de la reunión.
- **No automatizar la interpretación de costos invisibles y riesgos todavía**: depende de criterio humano hasta documentar y validar la guía (Paso 6 de la metodología).

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar a quien hoy realiza los diagnósticos consultivos para validar la guía de preguntas real, y probarla en una reunión real antes de declararla estándar.

# SOP-12: Calificación inicial

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Prompt maestro del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) al paso 3 del proceso comercial ideal de CHACONTAINER: filtrar los prospectos generados en la prospección inversa (SOP-01), usando el perfil de cliente ideal (SOP-11), antes de invertir tiempo en el primer contacto (paso 4) y el diagnóstico consultivo (paso 5). Con este SOP quedan documentados como borrador los 12 pasos del proceso comercial ideal. Se construyó a partir de prácticas estándar de calificación B2B industrial, **no a partir de una entrevista real con quien hoy califica prospectos en CHACONTAINER** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6 de la metodología). Los puntos marcados como `[VALIDAR]` requieren confirmación directa del responsable comercial.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Filtrar los prospectos generados en la prospección (SOP-01) confirmando tres criterios clave —uso de empaque retornable, volumen y responsable identificable— antes de invertir tiempo en agendar una reunión de diagnóstico.
2. **¿Qué problema busca resolver?** Evitar agendar reuniones de diagnóstico con prospectos que no califican (sin volumen suficiente, sin uso real de empaque retornable, sin un interlocutor claro), desperdiciando tiempo comercial; evitar decidir "a ojo" si un prospecto vale la pena.
3. **¿Quién es el cliente?** Interno: quien ejecuta el primer contacto (paso 4) y el diagnóstico consultivo (paso 5), que necesita recibir solo prospectos ya filtrados. Externo: el prospecto registrado en la prospección (SOP-01).
4. **¿Cuáles son las entradas?** La lista de prospectos registrados en la prospección inversa (SOP-01), el perfil de cliente ideal (ICP, SOP-11), un guion o checklist de preguntas de calificación inicial `[VALIDAR]`, y el canal de contacto ya establecido con el prospecto.
5. **¿Cuáles son las salidas esperadas?** El prospecto clasificado como "calificado" (con volumen, uso de empaque retornable y responsable identificable confirmados) o "descartado" (con motivo), listo para pasar a primer contacto (paso 4) si calificó.
6. **¿Quién es el responsable?** `[VALIDAR]` — el mismo ejecutivo/SDR que hizo la prospección, o un rol de calificación distinto.
7. **¿Qué indicadores dicen que el proceso funciona?** % de prospectos de la prospección que pasan la calificación inicial, % de reuniones de diagnóstico con prospectos calificados que terminan en negocio, tiempo entre prospección y calificación completada.
8. **¿Qué riesgos existen?** Calificar de forma superficial o sin preguntas estructuradas, dejando pasar prospectos que no califican realmente; descartar prospectos válidos por aplicar el criterio de forma inconsistente; depender del criterio informal del fundador para decidir quién "vale la pena".
9. **¿Qué actividades no agregan valor?** Agendar reuniones de diagnóstico sin calificación previa; calificar de memoria sin registrar las respuestas; repetir preguntas ya respondidas en la prospección.
10. **¿Qué partes pueden estandarizarse?** El guion de preguntas de calificación inicial (uso de empaque retornable, volumen, responsable identificable), y el criterio de "calificado" frente a "descartado".
11. **¿Qué partes pueden automatizarse?** Un formulario o cuestionario corto que el prospecto responda antes de agendar (según el canal), y el registro automático del resultado de calificación en el CRM.
12. **¿Qué depende únicamente del fundador?** `[VALIDAR]` — posiblemente el criterio para casos límite (volumen ambiguo, industria no cubierta explícitamente en el ICP).

---

## 1. Resumen ejecutivo

La calificación inicial es el filtro que separa los prospectos que vale la pena avanzar hacia una reunión de diagnóstico de los que no, aplicando tres preguntas clave: si usan empaque retornable, cuál es su volumen, y si existe un responsable identificable. Sin este filtro, el equipo comercial invierte tiempo de diagnóstico consultivo (paso 5) en prospectos que nunca debieron pasar de la prospección (SOP-01).

## 2. Objetivo

Confirmar, con un guion estructurado de preguntas, que un prospecto usa o podría usar empaque retornable, tiene un volumen relevante y cuenta con un responsable identificable, antes de agendar el primer contacto (paso 4).

## 3. Alcance

Aplica a todo prospecto registrado en la prospección inversa (SOP-01), desde el primer contacto de calificación hasta que queda clasificado como calificado o descartado. No incluye la reunión de diagnóstico (paso 5) ni el primer contacto formal de agenda (paso 4).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Ejecutivo comercial/SDR `[VALIDAR]` | Contacta al prospecto, aplica el guion de calificación y registra el resultado. |
| Responsable comercial `[VALIDAR]` | Define el guion de calificación y el umbral de volumen, y resuelve casos límite. |
| Fundador (rol transitorio) | Aporta hoy el criterio para casos límite; el objetivo del proceso es documentar ese criterio para no depender de él en cada caso. |

## 5. Diagrama de flujo

```
[Prospecto registrado (SOP-01)]
                ↓
[Contactar al prospecto con el guion de calificación]
                ↓
[Preguntar: uso de empaque retornable / volumen / responsable identificable]
                ↓
        ¿Cumple los 3 criterios?
        ↓ sí                              ↓ no
[Marcar como calificado]           [Marcar como descartado, con motivo]
                ↓
[Entrega a Primer contacto (paso 4)]
```

## 6. SOP paso a paso

1. Tomar el prospecto registrado en la prospección inversa (SOP-01).
2. Contactar al prospecto usando el guion de preguntas de calificación inicial `[VALIDAR guion exacto]`.
3. Confirmar si el prospecto usa o podría usar empaque retornable en su operación.
4. Confirmar el volumen aproximado (unidades, frecuencia) para descartar casos sin escala suficiente `[VALIDAR umbral mínimo]`.
5. Confirmar que existe un responsable identificable (nombre y cargo) con quien continuar el proceso.
6. Si el prospecto cumple los tres criterios, marcarlo como calificado y registrar los datos obtenidos.
7. Si no cumple, marcarlo como descartado y registrar el motivo (sin volumen, sin uso real, sin responsable identificable).
8. Entregar los prospectos calificados a Primer contacto (paso 4).

## 7. Checklist

- [ ] Guion de preguntas de calificación usado, no conversación improvisada.
- [ ] Uso de empaque retornable confirmado explícitamente.
- [ ] Volumen aproximado confirmado, no asumido.
- [ ] Responsable identificable (nombre y cargo) confirmado.
- [ ] Resultado de la calificación (calificado/descartado) registrado en el CRM el mismo día.
- [ ] Motivo de descarte registrado cuando aplica.

## 8. KPI

- % de prospectos de la prospección que pasan la calificación inicial.
- % de reuniones de diagnóstico con prospectos calificados que terminan en negocio.
- Tiempo entre prospección y calificación completada.
- % de prospectos descartados con motivo registrado (frente a simplemente abandonados sin registro).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Calificación superficial que deja pasar prospectos que no califican | Guion de preguntas obligatorio con los 3 criterios explícitos. |
| Descartar prospectos válidos por criterio inconsistente | Documentar el umbral de volumen y el criterio de "responsable identificable". |
| Dependencia del criterio informal del fundador en casos límite | Documentar reglas para casos límite y escalar solo las excepciones reales. |

## 10. Mejoras

- Documentar un guion corto y estándar de calificación inicial con las 3 preguntas clave, en vez de dejarlo a la conversación libre.
- Definir un umbral mínimo de volumen para calificar, en vez de un juicio subjetivo caso por caso.
- Registrar sistemáticamente los descartes con motivo, para retroalimentar el perfil de cliente ideal (SOP-11) si se detectan patrones.

## 11. Recomendaciones de automatización

- Formulario o cuestionario corto que el prospecto pueda responder antes de agendar (según el canal), reduciendo la carga manual.
- Registro automático del resultado de calificación (calificado/descartado + motivo) en el CRM.
- **No automatizar la decisión de calificar en casos límite**: requiere criterio humano hasta documentar reglas claras (Paso 6 de la metodología).

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar a quien hoy califica los prospectos para validar el guion real de preguntas y el umbral de volumen, y probarlo con la próxima tanda de prospectos de SOP-01 antes de declararlo estándar.

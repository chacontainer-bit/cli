# SOP-10: Postventa y expansión

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Prompt maestro del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) al paso 12 del proceso comercial ideal de CHACONTAINER: medir resultados, documentar el valor entregado y buscar contratos recurrentes después de la implementación (paso 11). Se construyó a partir de prácticas estándar de gestión de cuentas y retención B2B industrial, **no a partir de una entrevista real con quien hoy da seguimiento postventa en CHACONTAINER** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6 de la metodología). Los puntos marcados como `[VALIDAR]` requieren confirmación directa del responsable comercial/de cuentas.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Dar seguimiento al cliente después de la implementación para medir resultados, asegurar su satisfacción, resolver incidencias, y detectar oportunidades de expansión (mayor volumen, nuevos servicios, renovación o contratos recurrentes).
2. **¿Qué problema busca resolver?** Evitar que el cliente quede sin seguimiento después de la venta (riesgo de que se vaya con la competencia sin previo aviso); evitar perder oportunidades de expansión por falta de seguimiento sistemático; evitar que la relación postventa dependa del contacto informal y esporádico del fundador.
3. **¿Quién es el cliente?** Interno: comercial, que necesita detectar expansión; dirección, que mide retención y recurrencia como indicador de negocio. Externo: el cliente ya implementado.
4. **¿Cuáles son las entradas?** El registro de la implementación completada (paso 11, acta de entrega), el historial de uso/consumo del cliente (si aplica el servicio de trazabilidad/QR), un cronograma de seguimiento postventa `[VALIDAR]`, y un canal de atención a incidencias.
5. **¿Cuáles son las salidas esperadas?** Reporte periódico de resultados y satisfacción del cliente, incidencias resueltas y registradas, oportunidades de expansión identificadas, y contratos recurrentes o renovaciones gestionados.
6. **¿Quién es el responsable?** `[VALIDAR]` — un rol de postventa/cuentas, o el mismo comercial que cerró el negocio.
7. **¿Qué indicadores dicen que el proceso funciona?** Tasa de retención/renovación de clientes, % de clientes con seguimiento postventa registrado dentro del plazo definido, número de oportunidades de expansión detectadas y convertidas, tiempo de resolución de incidencias reportadas.
8. **¿Qué riesgos existen?** Churn silencioso (cliente insatisfecho que no reporta y simplemente no renueva); incidencias no resueltas a tiempo que erosionan la relación; depender de que el fundador recuerde hacer seguimiento a cada cliente; no capturar el resultado real (ahorro, reducción de mermas) que valida el valor entregado.
9. **¿Qué actividades no agregan valor?** Seguimiento reactivo, solo cuando el cliente se queja, en vez de proactivo; no medir ni documentar resultados, dejando el valor entregado como una anécdota; buscar expansión sin datos de uso o satisfacción que la sustenten.
10. **¿Qué partes pueden estandarizarse?** Un cronograma de seguimiento postventa (frecuencia de contacto), una plantilla de reporte de resultados/satisfacción, y un checklist para identificar oportunidades de expansión.
11. **¿Qué partes pueden automatizarse?** Recordatorios de seguimiento periódico, generación de reportes de uso/consumo a partir del sistema de trazabilidad, y alertas cuando un cliente reduce su consumo (riesgo de churn) o lo aumenta (señal de expansión).
12. **¿Qué depende únicamente del fundador?** `[VALIDAR]` — posiblemente la relación personal con las cuentas más grandes o estratégicas, y el criterio para negociar renovaciones o expansión en esos casos.

---

## 1. Resumen ejecutivo

La postventa y expansión es el paso que cierra y reinicia el ciclo comercial: da seguimiento al cliente después de la implementación para medir si la solución realmente entregó valor, resuelve incidencias antes de que erosionen la relación, y detecta oportunidades de vender más o renovar. Sin este paso sistemático, CHACONTAINER cierra negocios pero no los convierte en relaciones recurrentes, y depende del contacto informal del fundador para retener cuentas.

## 2. Objetivo

Dar seguimiento proactivo y sistemático a cada cliente implementado, midiendo y documentando resultados reales, resolviendo incidencias a tiempo, y detectando oportunidades de expansión o renovación que alimenten de nuevo el ciclo comercial.

## 3. Alcance

Aplica desde que la implementación (paso 11) está completada, de forma continua durante toda la relación con el cliente. Cuando se detecta una oportunidad de expansión, el alcance de este SOP termina al reiniciar el ciclo comercial desde calificación inicial (paso 3) o diagnóstico consultivo (paso 5), según corresponda.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Responsable de postventa/cuentas `[VALIDAR]` | Ejecuta el cronograma de seguimiento, mide resultados y detecta oportunidades de expansión. |
| Ejecutivo comercial | Recibe las oportunidades de expansión detectadas y las convierte en un nuevo ciclo comercial. |
| Fundador (rol transitorio) | Aporta hoy la relación personal con cuentas estratégicas; el objetivo del proceso es que el seguimiento no dependa exclusivamente de él. |

## 5. Diagrama de flujo

```
[Implementación completada (paso 11)]
                ↓
[Programar cronograma de seguimiento postventa]
                ↓
[Dar seguimiento periódico: medir uso, resultados, satisfacción]
                ↓
        ¿Hay incidencias reportadas?
        ↓ sí                          ↓ no
[Resolver incidencia y registrar]      │
        ↓                             │
        └─────────────┬───────────────┘
                       ↓
[Documentar resultados (ahorro, reducción de mermas, etc.)]
                       ↓
[Evaluar oportunidad de expansión/renovación]
                       ↓
        ¿Hay oportunidad?
        ↓ sí                                  ↓ no
[Iniciar nuevo ciclo desde Calificación         [Continuar seguimiento
 inicial (paso 3) o Diagnóstico                  periódico]
 consultivo (paso 5)]
```

## 6. SOP paso a paso

1. Confirmar que la implementación (paso 11) está completada y el acta de entrega firmada.
2. Programar el cronograma de seguimiento postventa con el cliente (frecuencia de contacto) `[VALIDAR frecuencia estándar]`.
3. Dar seguimiento periódico y proactivo, no solo cuando el cliente reporta un problema.
4. Medir y registrar resultados concretos: uso/consumo real, reducción de los costos visibles e invisibles detectados en el diagnóstico (paso 5), y satisfacción del cliente.
5. Si el cliente reporta una incidencia, resolverla dentro de un plazo definido y registrar la resolución.
6. Documentar los resultados obtenidos en un reporte periódico, útil tanto para retener al cliente como para usarlo como evidencia comercial con otros prospectos.
7. Identificar oportunidades de expansión (mayor volumen, nuevos servicios, renovación de contrato) a partir de los datos de uso y de la relación con el cliente.
8. Si se detecta una oportunidad, iniciarla como un nuevo ciclo desde calificación inicial (paso 3) o diagnóstico consultivo (paso 5), según corresponda.
9. Si no hay oportunidad inmediata, continuar el seguimiento periódico según el cronograma.

## 7. Checklist

- [ ] Cronograma de seguimiento postventa definido y activo para cada cliente implementado.
- [ ] Seguimiento realizado de forma proactiva, no solo reactiva a quejas.
- [ ] Resultados concretos medidos y registrados, no solo percepción subjetiva.
- [ ] Incidencias reportadas resueltas dentro del plazo definido y registradas.
- [ ] Reporte periódico de resultados documentado.
- [ ] Oportunidades de expansión evaluadas explícitamente, no dejadas al azar.
- [ ] Si hay oportunidad, iniciado el ciclo correspondiente del proceso comercial.

## 8. KPI

- Tasa de retención/renovación de clientes.
- % de clientes con seguimiento postventa registrado dentro del plazo definido.
- Número de oportunidades de expansión detectadas y convertidas en nuevo negocio.
- Tiempo de resolución de incidencias reportadas.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Churn silencioso (cliente insatisfecho que no reporta y simplemente no renueva) | Seguimiento proactivo periódico, no solo reactivo a quejas. |
| Incidencias no resueltas a tiempo erosionan la relación | Plazo definido de resolución y registro obligatorio de incidencias. |
| Depender de que el fundador recuerde hacer seguimiento | Cronograma de seguimiento documentado y asignado a un responsable, no a la memoria de una persona. |
| No capturar el resultado real entregado (ahorro, reducción de mermas) | Medición y registro obligatorio de resultados concretos, no solo percepción. |

## 10. Mejoras

- Definir un cronograma estándar de seguimiento postventa (frecuencia de contacto) para todos los clientes implementados, en vez de un seguimiento informal.
- Crear una plantilla de reporte de resultados que documente el valor entregado (ahorro, reducción de mermas, mejora de trazabilidad), útil también como evidencia comercial.
- Documentar un checklist para identificar sistemáticamente oportunidades de expansión, en vez de dejarlo a la intuición.

## 11. Recomendaciones de automatización

- Recordatorios automáticos de seguimiento periódico según el cronograma definido por cliente.
- Generación automática de reportes de uso/consumo a partir del sistema de trazabilidad, cuando el cliente tiene ese servicio.
- Alertas cuando un cliente reduce su consumo (señal de riesgo de pérdida) o lo aumenta (señal de oportunidad de expansión).
- **No automatizar la relación con cuentas estratégicas ni la negociación de renovaciones/expansión**: requiere criterio humano hasta documentar y validar el proceso (Paso 6 de la metodología).

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar a quien hoy da seguimiento postventa (probablemente el fundador, o nadie de forma sistemática) para validar el cronograma real y el proceso de detección de expansión, y probarlo con el próximo cliente implementado antes de declararlo estándar.

# SOP-08: Negociación y cierre

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Prompt maestro del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) al paso 10 del proceso comercial ideal de CHACONTAINER: llevar la cotización técnica y económica (paso 9) a un acuerdo cerrado, antes de la implementación, entrega y capacitación (paso 11). Se construyó a partir de prácticas estándar de negociación B2B industrial, **no a partir de una entrevista real con quien hoy negocia y cierra en CHACONTAINER** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6 de la metodología). Los puntos marcados como `[VALIDAR]` requieren confirmación directa del responsable comercial.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Llevar la cotización entregada (paso 9) a un acuerdo firmado o una orden de compra confirmada, manejando objeciones y ajustes dentro de rangos aprobados, sin regalar margen ni perder el negocio por rigidez.
2. **¿Qué problema busca resolver?** Evitar que la negociación se alargue indefinidamente sin un proceso claro; evitar ceder descuentos o condiciones fuera de rango sin aprobación; evitar perder negocios por no dar seguimiento oportuno; evitar que el cierre dependa de que el fundador negocie personalmente cada caso.
3. **¿Quién es el cliente?** Interno: quien implementa (paso 11) y necesita saber exactamente qué se acordó; finanzas, que necesita el contrato con las condiciones reales. Externo: el prospecto que evalúa la cotización.
4. **¿Cuáles son las entradas?** La cotización técnica y económica enviada (paso 9), los rangos de aprobación para descuentos y condiciones `[VALIDAR — parcialmente definidos en SOP-07]`, las objeciones del cliente, y una plantilla de contrato u orden de compra.
5. **¿Cuáles son las salidas esperadas?** Contrato u orden de compra firmada o confirmada, con las condiciones finales acordadas (precio, tiempos, exclusiones, forma de pago), registrada y lista para pasar a implementación (paso 11).
6. **¿Quién es el responsable?** `[VALIDAR]` — ejecutivo comercial, con escalamiento a finanzas/fundador para condiciones fuera de rango.
7. **¿Qué indicadores dicen que el proceso funciona?** Tasa de cierre (cotizaciones que terminan en contrato firmado), tiempo entre cotización enviada y cierre, % de cierres dentro del rango de condiciones aprobado, número de negocios perdidos por objeciones no resueltas a tiempo.
8. **¿Qué riesgos existen?** Ceder descuentos o condiciones fuera de rango sin aprobación, erosionando el margen; perder el negocio por demora en responder objeciones; depender de que el fundador negocie personalmente, generando cuello de botella; cerrar verbalmente sin dejar el acuerdo por escrito.
9. **¿Qué actividades no agregan valor?** Negociar sin un límite claro de hasta dónde se puede ceder; dar seguimiento sin registrar en qué quedó cada objeción; repetir información ya cubierta en la cotización en cada llamada de seguimiento.
10. **¿Qué partes pueden estandarizarse?** Los rangos de negociación (descuento máximo, condiciones negociables y no negociables), un checklist de objeciones frecuentes con respuestas estándar, y la plantilla de contrato/orden de compra.
11. **¿Qué partes pueden automatizarse?** Recordatorios de seguimiento a cotizaciones sin respuesta, la generación del contrato/orden de compra a partir de la cotización aprobada, y el registro del estatus de la negociación.
12. **¿Qué depende únicamente del fundador?** `[VALIDAR]` — posiblemente la aprobación final de condiciones fuera de rango, y la negociación directa en cuentas grandes o estratégicas.

---

## 1. Resumen ejecutivo

La negociación y cierre es el paso donde la cotización deja de ser una propuesta y se convierte en un acuerdo firmado: se manejan las objeciones del cliente, se ajustan condiciones dentro de rangos previamente aprobados, y se formaliza el negocio por escrito antes de pasar a implementación. Sin rangos claros de negociación, este paso depende del criterio individual de quien negocia y puede erosionar el margen o alargarse sin control.

## 2. Objetivo

Cerrar el negocio manejando objeciones y ajustes dentro de rangos de negociación previamente aprobados, formalizando el acuerdo por escrito, sin depender de que cada caso se negocie desde cero o se escale innecesariamente.

## 3. Alcance

Aplica desde que la cotización (paso 9) fue enviada al cliente hasta que el negocio se cierra (contrato u orden de compra firmada) o se pierde de forma definitiva. No incluye la implementación, entrega y capacitación (paso 11).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Ejecutivo comercial | Da seguimiento a la cotización, maneja objeciones dentro del rango aprobado y formaliza el cierre. |
| Responsable financiero/fundador `[VALIDAR]` | Aprueba condiciones que se salen del rango estándar de negociación. |
| Fundador (rol transitorio) | Aporta hoy la negociación directa en cuentas grandes o estratégicas; el objetivo del proceso es que los rangos documentados permitan que el comercial resuelva la mayoría de los casos sin escalar. |

## 5. Diagrama de flujo

```
[Cotización enviada (paso 9)]
                ↓
[Dar seguimiento activo a la cotización]
                ↓
[Cliente presenta objeciones/ajustes] ── no hay objeciones ──→ [Cliente acepta]
                ↓ sí
[Evaluar el ajuste contra los rangos de negociación aprobados]
                ↓
        ¿Dentro del rango?
        ↓ sí                              ↓ no
[Confirmar el ajuste con el cliente]   [Escalar a finanzas/fundador
                ↓                        para aprobación]
                └───────────┬───────────────┘
                            ↓
            [Cliente acepta las condiciones finales]
                            ↓
            [Formalizar: contrato u orden de compra firmada]
                            ↓
            [Registrar cierre y condiciones finales]
                            ↓
    [Entrega a Implementación, entrega y capacitación (paso 11)]
```

## 6. SOP paso a paso

1. Confirmar que la cotización (paso 9) fue enviada y registrar la fecha de seguimiento inicial.
2. Dar seguimiento activo a la cotización dentro del plazo definido, sin esperar a que el cliente escriba primero `[VALIDAR plazo estándar de seguimiento]`.
3. Registrar cada objeción o solicitud de ajuste del cliente (precio, tiempos, condiciones) de forma explícita.
4. Evaluar la objeción contra los rangos de negociación aprobados (definidos junto con la cotización, SOP-07).
5. Si el ajuste solicitado está dentro del rango aprobado, confirmarlo directamente con el cliente.
6. Si el ajuste requiere salirse del rango, escalar la aprobación a finanzas/fundador antes de comprometerse con el cliente `[VALIDAR quién aprueba]`.
7. Una vez que el cliente acepta las condiciones finales, formalizar el acuerdo por escrito (contrato u orden de compra firmada).
8. Registrar el cierre con las condiciones finales acordadas, o registrar el negocio como perdido con el motivo, si no se llega a acuerdo.
9. Entregar el contrato/orden de compra a Implementación, entrega y capacitación (paso 11).

## 7. Checklist

- [ ] Seguimiento a la cotización realizado dentro del plazo definido.
- [ ] Cada objeción del cliente registrada explícitamente, no de memoria.
- [ ] Ajustes evaluados contra los rangos de negociación aprobados antes de confirmarlos.
- [ ] Aprobación de finanzas/fundador obtenida si el ajuste sale del rango.
- [ ] Acuerdo final formalizado por escrito (contrato/orden de compra), no solo verbal.
- [ ] Condiciones finales (precio, tiempos, exclusiones, pago) registradas exactamente como se cerraron.
- [ ] Si el negocio se pierde, el motivo queda registrado para análisis posterior.

## 8. KPI

- Tasa de cierre (cotizaciones que terminan en contrato firmado).
- Tiempo entre cotización enviada y cierre.
- % de cierres dentro del rango de condiciones aprobado (sin excepciones no autorizadas).
- Número de negocios perdidos por objeciones no resueltas a tiempo.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Ceder descuentos o condiciones fuera de rango sin aprobación | Rangos de negociación documentados y escalamiento obligatorio fuera de ellos. |
| Perder el negocio por demora en responder objeciones | Plazo estándar de seguimiento definido y monitoreado. |
| Depender de que el fundador negocie personalmente cada caso | Documentar rangos y objeciones frecuentes para que el comercial resuelva la mayoría sin escalar. |
| Cerrar verbalmente sin dejar el acuerdo por escrito | Formalización obligatoria (contrato/orden de compra) antes de pasar a implementación. |

## 10. Mejoras

- Documentar los rangos de negociación (descuento máximo, condiciones negociables) junto con la cotización, para no improvisar en cada caso.
- Crear un checklist de objeciones frecuentes con respuestas estándar, para agilizar el manejo sin depender de la experiencia individual.
- Registrar sistemáticamente los negocios perdidos con su motivo, para detectar patrones (precio, tiempos, competencia) y ajustar el proceso.

## 11. Recomendaciones de automatización

- Recordatorios automáticos de seguimiento a cotizaciones sin respuesta dentro del plazo definido.
- Generación del contrato/orden de compra a partir de la cotización aprobada, sin redigitar la información.
- Registro automático del estatus de la negociación (en curso, ganado, perdido, motivo) en el CRM.
- **No automatizar la negociación misma** (manejo de objeciones, ajuste de condiciones): requiere criterio humano hasta documentar y validar los rangos de negociación (Paso 6 de la metodología).

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar a quien hoy negocia y cierra (probablemente el fundador) para validar los rangos reales de negociación y las objeciones más frecuentes, y probar el proceso en el próximo cierre antes de declararlo estándar.

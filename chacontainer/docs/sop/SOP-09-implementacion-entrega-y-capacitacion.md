# SOP-09: Implementación, entrega y capacitación

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Prompt maestro del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) al paso 11 del proceso comercial ideal de CHACONTAINER: poner en operación la solución acordada en el cierre (paso 10), antes de pasar a postventa y expansión (paso 12). Se construyó a partir de prácticas estándar de implementación B2B industrial, **no a partir de una entrevista real con quien hoy coordina entregas y capacitaciones en CHACONTAINER** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6 de la metodología). Los puntos marcados como `[VALIDAR]` requieren confirmación directa del responsable de logística/operaciones.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Poner en operación la solución acordada —entregando los activos, instalando lo que corresponda y capacitando al personal del cliente— para que empiece a usarla correctamente desde el primer día.
2. **¿Qué problema busca resolver?** Evitar que un negocio ya cerrado se caiga en la ejecución por mala coordinación de entrega; evitar que el cliente use mal el sistema o los activos por falta de capacitación, generando quejas o pérdidas después; evitar que la implementación dependa del criterio improvisado de quien la ejecuta cada vez.
3. **¿Quién es el cliente?** Interno: postventa (paso 12), que necesita que la implementación quede documentada para dar seguimiento; logística, que coordina la entrega. Externo: el cliente que recibe la solución y su personal, que debe operarla.
4. **¿Cuáles son las entradas?** El contrato u orden de compra firmada (paso 10), la especificación de la solución (paso 8) y el levantamiento técnico (paso 7), el inventario/activos disponibles para entrega, material de capacitación `[VALIDAR si existe]`, y un cronograma de implementación.
5. **¿Cuáles son las salidas esperadas?** Activos entregados e instalados según lo acordado, personal del cliente capacitado en el uso correcto, un acta de entrega/aceptación firmada o confirmada por el cliente, y el caso listo para pasar a postventa (paso 12).
6. **¿Quién es el responsable?** `[VALIDAR]` — logística/operaciones para la entrega, y comercial o un rol técnico para la capacitación.
7. **¿Qué indicadores dicen que el proceso funciona?** % de implementaciones entregadas en la fecha acordada, % de actas de entrega firmadas sin observaciones, % de reclamos posteriores atribuibles a falta de capacitación, tiempo entre el cierre (paso 10) y la implementación completada.
8. **¿Qué riesgos existen?** Retraso en la entrega por falta de coordinación con inventario/logística; activos entregados que no coinciden con lo cotizado o diseñado; capacitación superficial o inexistente que genera mal uso del sistema; falta de registro formal de que el cliente aceptó la entrega.
9. **¿Qué actividades no agregan valor?** Entregar sin verificar contra la especificación de la solución; capacitar de forma improvisada sin material de apoyo; no dejar constancia de la entrega ni de la capacitación.
10. **¿Qué partes pueden estandarizarse?** El checklist de entrega (verificación contra lo cotizado/diseñado), el material de capacitación por tipo de servicio, y el acta de entrega/aceptación.
11. **¿Qué partes pueden automatizarse?** La programación del cronograma de entrega según disponibilidad de inventario, la generación del acta de entrega a partir del contrato y del diseño de solución, y los recordatorios de seguimiento post-implementación.
12. **¿Qué depende únicamente del fundador?** `[VALIDAR]` — posiblemente la capacitación en casos complejos o estratégicos, y la resolución de imprevistos durante la implementación.

---

## 1. Resumen ejecutivo

La implementación, entrega y capacitación es el paso donde el negocio cerrado se vuelve realidad operativa: los activos o el sistema acordado se entregan e instalan en el sitio del cliente, y su personal queda capacitado para usarlos correctamente. Un negocio bien cerrado (paso 10) puede fallar aquí si la entrega no coincide con lo cotizado o si la capacitación es superficial, generando reclamos que en realidad son fallas de ejecución, no del producto.

## 2. Objetivo

Entregar e instalar la solución acordada exactamente como fue diseñada y cotizada, capacitar al personal del cliente en su uso correcto, y dejar constancia formal de la entrega antes de pasar a postventa.

## 3. Alcance

Aplica desde que el contrato u orden de compra (paso 10) está firmado hasta que la implementación queda formalmente aceptada por el cliente. No incluye el seguimiento posterior ni la búsqueda de contratos recurrentes (postventa y expansión, paso 12).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Logística/operaciones `[VALIDAR]` | Coordina el cronograma, verifica inventario y ejecuta la entrega/instalación. |
| Ejecutivo comercial/técnico | Capacita al personal del cliente y levanta el acta de entrega/aceptación. |
| Fundador (rol transitorio) | Aporta hoy la capacitación en casos complejos o estratégicos, y resuelve imprevistos; el objetivo del proceso es documentar el material de capacitación para no depender de él. |

## 5. Diagrama de flujo

```
[Contrato/orden de compra firmada (paso 10)]
                ↓
[Programar cronograma de entrega/implementación]
                ↓
[Verificar disponibilidad de inventario/activos]
                ↓
[Entregar/instalar según especificación de la solución]
                ↓
[Capacitar al personal del cliente]
                ↓
[Levantar acta de entrega/aceptación]
                ↓
        ¿Cliente acepta sin observaciones?
        ↓ sí                          ↓ no
[Registrar implementación         [Resolver observaciones
 completada]                        y volver a validar]
                ↓
[Entrega a Postventa y expansión (paso 12)]
```

## 6. SOP paso a paso

1. Confirmar que el contrato u orden de compra (paso 10) está firmado y disponible.
2. Programar el cronograma de entrega/implementación con el cliente, verificando la disponibilidad real de inventario o activos.
3. Preparar los activos o el sistema a entregar según la especificación de la solución (paso 8) y el levantamiento técnico (paso 7).
4. Entregar e instalar (si aplica) los activos en el sitio del cliente, verificando contra la especificación con un checklist.
5. Capacitar al personal del cliente en el uso correcto, el manejo y, si aplica, el ciclo de retorno/rotación de los activos `[VALIDAR material de capacitación]`.
6. Levantar un acta de entrega/aceptación con el cliente, dejando constancia de lo entregado y de la capacitación realizada.
7. Si el cliente presenta observaciones, resolverlas antes de dar por cerrada la implementación.
8. Registrar la implementación como completada y entregar el caso a Postventa y expansión (paso 12).

## 7. Checklist

- [ ] Contrato/orden de compra firmado disponible antes de programar la entrega.
- [ ] Cronograma de entrega acordado con el cliente y verificado contra disponibilidad de inventario.
- [ ] Activos/sistema entregados verificados contra la especificación de la solución (paso 8).
- [ ] Capacitación realizada al personal del cliente, con material de apoyo, no improvisada.
- [ ] Acta de entrega/aceptación firmada o confirmada por el cliente.
- [ ] Observaciones del cliente, si las hubo, resueltas antes de cerrar la implementación.
- [ ] Caso registrado como completado y entregado a postventa.

## 8. KPI

- % de implementaciones entregadas en la fecha acordada.
- % de actas de entrega firmadas sin observaciones.
- % de reclamos posteriores atribuibles a falta o mala capacitación.
- Tiempo entre el cierre (paso 10) y la implementación completada.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Retraso en la entrega por falta de coordinación logística | Verificar disponibilidad real de inventario antes de comprometer fecha con el cliente. |
| Activos entregados no coinciden con lo cotizado/diseñado | Checklist de entrega que verifica explícitamente contra la especificación de la solución. |
| Capacitación superficial que genera mal uso del sistema | Material de capacitación estándar obligatorio, no improvisado. |
| Falta de constancia formal de la entrega/capacitación | Acta de entrega/aceptación obligatoria, firmada o confirmada por el cliente. |

## 10. Mejoras

- Crear un checklist de entrega que verifique explícitamente contra la especificación de la solución y el levantamiento técnico, antes de dar por entregado.
- Desarrollar material de capacitación estándar por tipo de servicio (venta, renta, trazabilidad, sistemas), para no depender de la improvisación de quien capacita.
- Formalizar un acta de entrega/aceptación única para todos los clientes, en vez de dejar constancia informal o nula.

## 11. Recomendaciones de automatización

- Programar automáticamente el cronograma de entrega según la disponibilidad real de inventario en el sistema.
- Generar el acta de entrega a partir de los datos del contrato y del diseño de la solución, sin redactar desde cero cada vez.
- Recordatorios automáticos para dar seguimiento post-implementación y activar el proceso de postventa (paso 12).
- **No automatizar la capacitación presencial ni la resolución de observaciones del cliente**: requieren interacción humana hasta validar el material y el proceso (Paso 6 de la metodología).

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar a quien hoy coordina las entregas/implementaciones y a quien capacita a los clientes, para validar el checklist de entrega y el material de capacitación reales, y probar el proceso en la próxima implementación antes de declararlo estándar.

# SOP-05: Levantamiento técnico

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Prompt maestro del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) al paso 7 del proceso comercial ideal de CHACONTAINER: medidas, cantidades, fotos, logística y restricciones del sitio del cliente, que siguen al diagnóstico consultivo (paso 5) y a la calificación como proveedor (paso 6), y alimentan el diseño de la solución (paso 8) y la cotización técnica y económica (paso 9). Se construyó a partir de prácticas estándar de levantamiento técnico en campo para soluciones de empaque industrial, **no a partir de una entrevista real con quien hoy realiza estas visitas en CHACONTAINER** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6 de la metodología). Los puntos marcados como `[VALIDAR]` requieren confirmación directa del responsable técnico.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Levantar toda la información técnica del sitio del cliente —medidas, cantidades, fotos, logística y restricciones— necesaria para diseñar una solución de empaque retornable que encaje físicamente y operativamente en su operación real.
2. **¿Qué problema busca resolver?** Evitar diseñar o cotizar una solución basada en supuestos o en lo que el cliente dijo de palabra, que luego no cabe en el sitio, no es compatible con su logística, o ignora una restricción crítica (peso, altura, acceso).
3. **¿Quién es el cliente?** Interno: quien diseñará la solución (paso 8) y quien cotizará (paso 9). Externo: el cliente ya diagnosticado (paso 5) y, si aplica, calificado como corresponde (paso 6).
4. **¿Cuáles son las entradas?** El documento de diagnóstico consultivo (paso 5), acceso físico al sitio del cliente (planta, bodega, andén), instrumentos de medición, cámara para fotos, y un checklist de levantamiento técnico `[VALIDAR checklist exacto]`.
5. **¿Cuáles son las salidas esperadas?** Documento de levantamiento técnico con medidas exactas, cantidades y volúmenes reales, fotos organizadas del sitio, condiciones logísticas (rutas, frecuencia, tipo de vehículo) y restricciones (peso, altura, horarios de acceso, normas de seguridad).
6. **¿Quién es el responsable?** `[VALIDAR]` — técnico o ejecutivo que realiza la visita de campo.
7. **¿Qué indicadores dicen que el proceso funciona?** % de levantamientos completos sin datos faltantes al momento de diseñar la solución, % de propuestas corregidas o rechazadas por errores de medidas o restricciones no detectadas, tiempo entre la calificación/diagnóstico y el levantamiento completado.
8. **¿Qué riesgos existen?** Medidas o cantidades estimadas "a ojo" en vez de medidas reales; fotos insuficientes que obligan a regresar al sitio; restricciones críticas no detectadas (límite de peso del andén, altura de acceso) que invalidan la solución después de cotizada; dependencia de una sola persona para saber qué medir y qué preguntar.
9. **¿Qué actividades no agregan valor?** Visitar el sitio sin checklist y tener que regresar por datos faltantes; tomar fotos desorganizadas sin etiquetar qué muestran; medir sin registrar unidades o referencias claras.
10. **¿Qué partes pueden estandarizarse?** El checklist de levantamiento técnico (qué medir, qué fotografiar, qué preguntar de logística y restricciones), la plantilla del documento de levantamiento, y la nomenclatura de las fotos.
11. **¿Qué partes pueden automatizarse?** Un formulario digital de levantamiento (captura estructurada desde celular/tablet en sitio), el almacenamiento automático de fotos asociadas al cliente/sitio, y la generación automática del documento a partir del formulario.
12. **¿Qué depende únicamente del fundador?** `[VALIDAR]` — posiblemente el criterio para detectar restricciones críticas no obvias, basado en su experiencia, y/o ser quien hoy realiza las visitas técnicas más complejas.

---

## 1. Resumen ejecutivo

El levantamiento técnico es la visita de campo en la que CHACONTAINER convierte el diagnóstico consultivo en datos concretos y verificables: medidas exactas, cantidades reales, fotos del sitio, condiciones logísticas y restricciones críticas. Sin este levantamiento bien hecho, el diseño de la solución (paso 8) y la cotización (paso 9) se basan en supuestos que suelen fallar en la implementación.

## 2. Objetivo

Recopilar de forma completa y verificable las medidas, cantidades, fotos, condiciones logísticas y restricciones del sitio del cliente, dejando un documento único que sustente el diseño de la solución y la cotización.

## 3. Alcance

Aplica a la visita técnica al sitio del cliente que sigue al diagnóstico consultivo (paso 5) y, si corresponde, a la calificación como proveedor (paso 6). No incluye el diseño de la solución (paso 8) ni la elaboración de la cotización (paso 9), aunque este documento es su insumo principal.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Técnico/ejecutivo de campo `[VALIDAR]` | Realiza la visita, mide, cuenta, fotografía y registra logística y restricciones. |
| Responsable de diseño de solución | Recibe el documento de levantamiento y valida que esté completo antes de diseñar (paso 8). |
| Fundador (rol transitorio) | Aporta hoy el criterio para detectar restricciones no obvias en visitas complejas; el objetivo del proceso es documentar ese criterio en el checklist. |

## 5. Diagrama de flujo

```
[Diagnóstico consultivo (y calificación como proveedor) completos]
                ↓
[Agendar visita de levantamiento técnico]
                ↓
[Preparar checklist e instrumentos de medición]
                ↓
[Realizar visita: medir, contar, fotografiar]
                ↓
[Registrar condiciones logísticas y restricciones]
                ↓
[Consolidar documento de levantamiento técnico]
                ↓
[Entrega a Diseño de la solución (paso 8)]
```

## 6. SOP paso a paso

1. Confirmar que el diagnóstico consultivo (paso 5) y, si aplica, la calificación como proveedor (paso 6) están completos antes de agendar el levantamiento.
2. Agendar la visita de levantamiento técnico con el cliente, confirmando acceso al sitio (planta, bodega, andén).
3. Preparar el checklist de levantamiento y los instrumentos de medición antes de la visita `[VALIDAR checklist exacto]`.
4. Durante la visita, medir con instrumento las dimensiones relevantes (espacios, andenes, pasillos, altura de acceso).
5. Contar o estimar las cantidades y volúmenes reales (unidades por periodo, rotación esperada).
6. Fotografiar el sitio de forma organizada y etiquetada, indicando qué muestra cada foto y en qué punto del sitio.
7. Registrar las condiciones logísticas (rutas, frecuencia de entrega/recolección, tipo de vehículo de acceso).
8. Identificar y registrar restricciones críticas (límites de peso, altura, horarios de acceso, normas de seguridad del sitio).
9. Consolidar toda la información en el documento estándar de levantamiento técnico.
10. Entregar el documento a Diseño de la solución (paso 8).

## 7. Checklist

- [ ] Diagnóstico consultivo (y calificación como proveedor, si aplica) completos antes de la visita.
- [ ] Checklist de levantamiento e instrumentos de medición preparados antes de ir al sitio.
- [ ] Medidas del sitio tomadas con instrumento, no estimadas a ojo.
- [ ] Cantidades y volúmenes reales registrados (no solo lo que el cliente dijo de palabra).
- [ ] Fotos tomadas de forma organizada y etiquetada.
- [ ] Condiciones logísticas registradas (rutas, frecuencia, tipo de vehículo).
- [ ] Restricciones críticas identificadas y registradas (peso, altura, horarios, seguridad).
- [ ] Documento de levantamiento consolidado el mismo día de la visita.

## 8. KPI

- % de levantamientos completos sin datos faltantes al momento de diseñar la solución.
- % de propuestas corregidas o rechazadas por errores de medidas o restricciones no detectadas en el levantamiento.
- Tiempo entre la calificación/diagnóstico y el levantamiento técnico completado.
- Número de visitas repetidas al mismo sitio por datos faltantes en la primera visita.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Medidas o cantidades estimadas "a ojo" | Checklist obligatorio con instrumento de medición, no estimación visual. |
| Fotos insuficientes que obligan a regresar al sitio | Checklist de fotos mínimas requeridas por tipo de sitio. |
| Restricciones críticas no detectadas (peso, altura, acceso) | Sección obligatoria de restricciones en el checklist, revisada antes de cerrar la visita. |
| Dependencia de una sola persona para saber qué medir | Documentar el checklist para que cualquier técnico pueda ejecutarlo igual. |

## 10. Mejoras

- Crear un checklist estándar de levantamiento técnico (qué medir, qué fotografiar, qué preguntar de logística y restricciones) en vez de dejarlo a la experiencia de quien visita.
- Definir una nomenclatura y forma de organizar las fotos por cliente/sitio, para que sean útiles después sin necesidad de regresar.
- Documentar las restricciones críticas típicas por tipo de instalación (bodega, planta, centro de distribución) para no descubrirlas hasta la implementación.

## 11. Recomendaciones de automatización

- Formulario digital de levantamiento (celular/tablet) que capture medidas, cantidades, fotos y restricciones de forma estructurada en el sitio.
- Almacenamiento automático de las fotos asociadas al cliente/sitio en el sistema de trazabilidad.
- Generación automática del documento de levantamiento técnico a partir del formulario digital, sin redacción manual posterior.
- **No automatizar la toma de medidas ni la identificación de restricciones críticas**: requieren presencia física y criterio humano hasta validar el checklist (Paso 6 de la metodología).

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar a quien hoy realiza las visitas técnicas para validar el checklist real de medidas y restricciones, y probarlo en la próxima visita antes de declararlo estándar.

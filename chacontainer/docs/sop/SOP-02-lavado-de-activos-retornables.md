# SOP-02: Lavado de activos retornables (contenedores, IBC, tarimas y tambores)

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Prompt maestro del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) a un proceso operativo real de CHACONTAINER: el lavado de activos retornables antes de reingresarlos a inventario disponible o entregarlos a un nuevo cliente. Se construyó a partir de prácticas estándar de lavado de empaque industrial retornable (IBC, tarimas plásticas, contenedores, tambores), **no a partir de una entrevista real con el responsable de planta/lavado** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6). Los puntos marcados como `[VALIDAR]` requieren confirmación directa del responsable de planta.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Dejar los activos retornables (contenedores, IBC, tarimas, tambores) limpios, inspeccionados y aptos para reingresar al inventario disponible o ser entregados a un nuevo cliente, cumpliendo el estándar de limpieza requerido según el producto que transportaron.
2. **¿Qué problema busca resolver?** Evitar contaminación cruzada entre clientes o productos incompatibles, evitar reclamos por entregar activos sucios o dañados, y evitar que el inventario "sucio" se acumule sin rotar por cuellos de botella en el lavado.
3. **¿Quién es el cliente?** Interno: logística y despacho, que necesitan inventario limpio disponible para rotar; comercial, que promete tiempos de entrega. Externo: el cliente que recibe el activo limpio y en buen estado.
4. **¿Cuáles son las entradas?** Activos sucios que regresan de campo (contenedores, IBC, tarimas, tambores), agua, detergentes/químicos de lavado `[VALIDAR cuáles y en qué concentración]`, estación o línea de lavado, EPP para el personal, orden de recepción o registro QR del activo.
5. **¿Cuáles son las salidas esperadas?** Activo lavado e inspeccionado, con evidencia (foto y/o registro QR) de "apto para reingreso" o "rechazado a reparación/baja", y su estado actualizado en el sistema de trazabilidad.
6. **¿Quién es el responsable?** `[VALIDAR]` — operador de planta o encargado de lavado (rol a confirmar).
7. **¿Qué indicadores dicen que el proceso funciona?** Tiempo de ciclo de lavado por activo o por lote, % de activos rechazados tras inspección, % de reclamos de clientes por limpieza deficiente, rotación de inventario limpio disponible.
8. **¿Qué riesgos existen?** Contaminación cruzada entre productos incompatibles; uso incorrecto de químicos (riesgo de seguridad y salud); activos que salen "lavados" con daño no detectado; dependencia de un solo operador que concentra el criterio de aceptación.
9. **¿Qué actividades no agregan valor?** Relavar activos por falta de un criterio claro de aceptación; esperas por falta de espacio en la estación de lavado; mover el activo entre áreas innecesariamente antes de lavarlo.
10. **¿Qué partes pueden estandarizarse?** Criterio de aceptación/rechazo por tipo de activo, método y tiempo de lavado, checklist de inspección post-lavado, registro de trazabilidad.
11. **¿Qué partes pueden automatizarse?** Registro de estado del activo vía escaneo QR al iniciar y terminar el lavado, alertas cuando un lote excede el tiempo estándar en la estación, generación automática de reportes de KPI de lavado.
12. **¿Qué depende únicamente del fundador?** `[VALIDAR]` — posiblemente el criterio final de "qué tan limpio es limpio" para clientes exigentes, y la decisión de rechazar o dar de baja un activo dañado.

---

## 1. Resumen ejecutivo

El lavado es el proceso que convierte un activo retornable sucio (contenedor, IBC, tarima o tambor) en un activo disponible para reingresar al inventario o entregarse a un nuevo cliente. Recibe activos desde el retorno de campo, los lava según el tipo de producto que transportaron, los inspecciona contra un criterio de aceptación, y actualiza su estado en el sistema de trazabilidad antes de liberarlos a despacho.

## 2. Objetivo

Garantizar que todo activo retornable que reingresa a inventario o se entrega a un cliente cumple el estándar de limpieza requerido, de forma consistente y sin depender del criterio informal de una sola persona.

## 3. Alcance

Aplica a todo activo retornable (contenedor, IBC, tarima, tambor) que regresa de campo y requiere lavado antes de reingresar a inventario disponible o ser entregado a un nuevo cliente. No incluye la reparación estructural del activo (proceso aparte) ni el transporte de retorno desde el cliente.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Operador de lavado `[VALIDAR]` | Ejecuta el lavado según el método definido por tipo de activo y registra el resultado. |
| Inspector de calidad `[VALIDAR si es el mismo operador o un rol distinto]` | Verifica el activo contra el checklist de aceptación y decide apto/rechazado. |
| Responsable de planta `[VALIDAR]` | Define y actualiza el criterio de aceptación, los tiempos estándar y resuelve excepciones. |

## 5. Diagrama de flujo

```
[Activo sucio ingresa de retorno de campo]
                ↓
[Registrar ingreso (QR / orden de recepción)]
                ↓
[Clasificar por tipo de activo y producto previo]
                ↓
[Lavar según método definido por tipo]
                ↓
[Inspeccionar contra checklist de aceptación]
                ↓
        ¿Cumple criterio? 
        ↓ sí              ↓ no
[Registrar "apto"    [Registrar "rechazado"
 en trazabilidad]     → enviar a reparación o baja]
                ↓
[Liberar a inventario disponible / despacho]
```

## 6. SOP paso a paso

1. Recibir el activo sucio proveniente del retorno de campo y registrar su ingreso (escaneo QR u orden de recepción).
2. Clasificar el activo por tipo (contenedor, IBC, tarima, tambor) y, si aplica, por el producto que transportó previamente `[VALIDAR si existe riesgo de incompatibilidad entre productos]`.
3. Lavar el activo con el método y los insumos definidos para su tipo `[VALIDAR método y químicos exactos por tipo de activo]`.
4. Inspeccionar visualmente el activo lavado contra el checklist de aceptación (limpieza, ausencia de residuos, integridad estructural básica).
5. Si el activo cumple el criterio, registrar el estado "apto" en el sistema de trazabilidad y moverlo a inventario disponible.
6. Si el activo no cumple el criterio, registrar el estado "rechazado" y enviarlo al proceso de reparación o a baja, según corresponda.
7. Dejar evidencia (foto y/o registro QR) del resultado de la inspección, especialmente en rechazos, para trazabilidad y control de calidad.
8. Actualizar el indicador de tiempo de ciclo del lote lavado.

## 7. Checklist

- [ ] Ingreso del activo registrado (QR/orden de recepción) antes de iniciar el lavado.
- [ ] Tipo de activo y producto previo identificados.
- [ ] Método de lavado correcto aplicado según el tipo de activo.
- [ ] Activo inspeccionado contra el checklist de aceptación (no solo "a simple vista").
- [ ] Evidencia (foto/registro) generada, especialmente si el activo fue rechazado.
- [ ] Estado del activo (apto/rechazado) actualizado en el sistema de trazabilidad el mismo día.
- [ ] Activo aprobado movido físicamente a la zona de inventario disponible.

## 8. KPI

- Tiempo de ciclo de lavado por activo o por lote.
- % de activos rechazados tras inspección post-lavado.
- % de reclamos de clientes atribuibles a limpieza deficiente.
- Rotación de inventario limpio disponible (activos liberados por día/semana).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Contaminación cruzada entre productos incompatibles | Clasificar el activo por producto previo antes de lavar y usar método/insumos según esa clasificación. |
| Uso incorrecto de químicos (riesgo de seguridad y salud) | Definir por escrito el químico y la concentración por tipo de activo, con EPP obligatorio. |
| Activo "lavado" con daño no detectado | Checklist de inspección post-lavado obligatorio, separado del lavado mismo. |
| Dependencia de un solo operador con el criterio de aceptación | Documentar el criterio de aceptación/rechazo por escrito y capacitar a más de una persona. |

## 10. Mejoras

- Documentar por escrito el criterio de aceptación/rechazo por tipo de activo, en vez de dejarlo al juicio del operador de turno.
- Definir el método y los insumos de lavado por tipo de activo y producto previo, para evitar contaminación cruzada.
- Separar explícitamente el paso de "lavado" del paso de "inspección", para que no dependan de la misma persona ni del mismo momento.

## 11. Recomendaciones de automatización

- Registrar el ingreso y la salida del activo de la estación de lavado vía escaneo QR, en vez de registro manual.
- Generar alertas automáticas cuando un lote excede el tiempo estándar en la estación de lavado (cuello de botella).
- Generar automáticamente el reporte semanal de KPI de lavado (tiempo de ciclo, % rechazo, reclamos) a partir de los registros de trazabilidad.
- **No automatizar hasta validar el proceso manual en operación real** (Paso 6 de la metodología): automatizar un criterio de aceptación mal definido solo escala el riesgo de calidad.

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar al responsable de planta/lavado para validar/ajustar los puntos marcados `[VALIDAR]` (método de lavado por tipo, químicos, roles, criterio de aceptación), y probar el proceso en campo antes de declararlo estándar.

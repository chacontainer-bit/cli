# SOP-ACTIVO-14: Logística inversa

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre la recolección activa de activos retornables cuando el retorno espontáneo no ocurre:
la transición `Retenido → En tránsito` de la [máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones)
(transición #18). Se construyó a partir del cruce entre la capacidad
[Logística inversa (#9)](../catalogo-systems.md#9-logística-inversa) del catálogo Systems, el servicio
[Recuperación (#12)](../catalogo-solutions.md#12-recuperación) del catálogo Solutions, y la rama de excepción
"Activo perdido/bloqueado/incidencia" descrita en
[ciclo-de-vida-del-activo.md §4](../ciclo-de-vida-del-activo.md#4-ramas-de-excepción-fuera-del-flujo-feliz) —
**no a partir de una entrevista real con el responsable operativo**. No debe declararse estándar hasta
validar en operación los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Logística inversa es el proceso que convierte una alerta de "activo detenido" en una acción física de
recolección. Un activo en estado `Retenido` (tiempo máximo fuera excedido, según [transición #17](../estados-del-activo.md#2-tabla-de-transiciones))
no regresa solo: sin una recolección planeada y ejecutada, se acumula indefinidamente en sitio del cliente
o del proveedor y deja de rotar. Este SOP recibe la lista de activos retenidos (por alerta de tiempo, por
vencimiento de renta, o por cierre de operación de un cliente), planea la ruta de recolección, ejecuta el
retiro físico y reingresa el activo al flujo — moviéndolo a `En tránsito` para su reinspección en planta. Es
la ejecución operativa de la capacidad [Logística inversa](../catalogo-systems.md#9-logística-inversa) de
Systems, apoyada cuando hace falta por el servicio [Recuperación](../catalogo-solutions.md#12-recuperación)
de Solutions.

## 2. Objetivo

Garantizar que todo activo que entra en estado `Retenido` sea recolectado de forma planeada y documentada,
minimizando el tiempo detenido y evitando que la recolección dependa de gestiones informales o de que el
cliente/custodio devuelva el activo por iniciativa propia.

## 3. Alcance

Aplica a todo activo retornable en estado `Retenido` cuya recolección no depende de que el custodio lo
regrese espontáneamente, sino de que CHACONTAINER organice y ejecute la recolección. Cubre:

1. **Recolección por alerta de tiempo** — el activo excede el tiempo máximo fuera definido en
   [Reglas operativas](../catalogo-systems.md#11-reglas-operativas) y dispara una
   [alerta](../catalogo-systems.md#12-alertas) automática.
2. **Recolección por vencimiento de renta** — el periodo de renta terminó y el activo no fue devuelto.
3. **Recolección masiva por cierre de operación de un cliente** — el cliente deja de operar o cambia de
   proveedor y hay que retirar todo el parque asignado.

No incluye la búsqueda de un activo sin ubicación confirmable (eso es
[SOP-ACTIVO-16 · Activo perdido](./SOP-ACTIVO-16-activo-perdido.md)), ni la resolución de una disputa de
custodia o condición que bloquea la recolección (eso es
[SOP-ACTIVO-17 · Activo bloqueado](./SOP-ACTIVO-17-activo-bloqueado.md)). Tampoco incluye la reinspección o
el mantenimiento que recibe el activo una vez que llega a planta (SOP-ACTIVO-03/04, SOP-ACTIVO-05/06 del
ciclo normal) — este SOP termina cuando el activo queda `En tránsito` hacia planta.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Logística `[VALIDAR nombre exacto del puesto]` | Planea la ruta de recolección, coordina el transporte y ejecuta el retiro físico. Es quien dispara la transición `Retenido → En tránsito` (fila 18 de la [tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones)). |
| Responsable de gestión de custodios `[VALIDAR]` | Provee el listado de custodios con activos retenidos y su historial de retención (ver [Gestión de custodios](../catalogo-systems.md#8-gestión-de-custodios)). |
| Comercial `[VALIDAR]` | Notifica al cliente/custodio la recolección programada cuando la relación comercial lo requiere; escala si el custodio se niega a entregar el activo (posible derivación a [SOP-ACTIVO-17](./SOP-ACTIVO-17-activo-bloqueado.md)). |
| Operador de campo (Recuperación) | Ejecuta el retiro físico cuando el activo requiere localización o negociación adicional en sitio — ver [Recuperación](../catalogo-solutions.md#12-recuperación) de Solutions. |

## 5. Diagrama de flujo

```
[Activo en estado "Retenido" — alerta de tiempo, vencimiento de renta,
 o cierre de operación del cliente]
                ↓
[Consultar trazabilidad y gestión de custodios: última ubicación
 y custodio confirmado del activo]
                ↓
[Agrupar activos retenidos por zona/ruta/custodio]
                ↓
[Planear ruta de recolección y coordinar transporte]
                ↓
[Notificar al custodio la recolección programada [VALIDAR: siempre / solo si aplica]]
                ↓
        ¿Custodio entrega el activo sin objeción?
        ↓ sí                                    ↓ no
[Ejecutar retiro físico                  [Registrar incidencia (SOP-ACTIVO-15)
 y firmar acta de retiro]                  → evaluar si es disputa de custodia/condición
                ↓                            → derivar a SOP-ACTIVO-17 si aplica]
[Registrar condición del activo
 al momento del retiro]
                ↓
[Actualizar estado → "En tránsito" y ubicación/custodia en el sistema]
                ↓
[Enviar a reinspección en planta (SOP-ACTIVO-03/04)]
```

## 6. SOP paso a paso

1. Identificar el origen de la retención: alerta automática de tiempo excedido
   ([Alertas](../catalogo-systems.md#12-alertas)), vencimiento de renta, o cierre de operación de un cliente
   (ver [§3 Alcance](#3-alcance)).
2. Consultar [Trazabilidad](../catalogo-systems.md#7-trazabilidad) y
   [Gestión de custodios](../catalogo-systems.md#8-gestión-de-custodios) para confirmar la última ubicación
   conocida y el custodio actual del activo antes de planear la recolección.
3. Agrupar los activos retenidos por zona geográfica, ruta o custodio, para consolidar recolecciones en vez
   de despachar transporte por activo individual `[VALIDAR criterio de agrupación y frecuencia de rutas]`.
4. Planear la ruta de recolección y coordinar el transporte (propio o tercero
   `[VALIDAR modelo de transporte]`).
5. Notificar al custodio la recolección programada, cuando la relación comercial con el cliente lo requiera
   `[VALIDAR si la notificación previa es obligatoria en todos los casos o solo quien tiene contrato vigente]`.
6. Ejecutar el retiro físico del activo en el punto acordado.
7. Si el custodio se niega a entregar el activo o hay desacuerdo sobre su condición, no forzar el retiro:
   registrar la incidencia según [SOP-ACTIVO-15](./SOP-ACTIVO-15-incidencias.md) y evaluar si corresponde
   abrir el proceso de [SOP-ACTIVO-17 · Activo bloqueado](./SOP-ACTIVO-17-activo-bloqueado.md) (transición
   `Retenido → Bloqueado`, fila 19 de la tabla de transiciones).
8. Si el retiro se ejecuta sin objeción, firmar el acta de retiro y registrar la condición del activo al
   momento de la recolección (evidencia fotográfica mínima).
9. Actualizar el estado del activo a `En tránsito` en el sistema (transición #18) y actualizar su ubicación
   (Capa 2) y custodia (Capa 3, vuelve a custodia de CHACONTAINER).
10. Enviar el activo a reinspección en planta, siguiente etapa del
    [ciclo de vida](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado) (SOP-ACTIVO-03/04).
11. Actualizar el indicador de % de activos recolectados sobre pendientes de recolección (ver
    [§8 KPI](#8-kpi)).

## 7. Checklist

- [ ] Origen de la retención identificado (alerta de tiempo / vencimiento de renta / cierre de operación).
- [ ] Última ubicación y custodio confirmados vía trazabilidad y gestión de custodios antes de planear la ruta.
- [ ] Activos retenidos agrupados por zona/ruta/custodio antes de despachar transporte.
- [ ] Ruta de recolección planeada y transporte coordinado.
- [ ] Custodio notificado cuando aplica.
- [ ] Retiro físico ejecutado y acta de retiro firmada, o incidencia registrada si el custodio no entrega el activo.
- [ ] Condición del activo registrada (con evidencia) al momento del retiro.
- [ ] Estado del activo actualizado a `En tránsito` y ubicación/custodia actualizadas el mismo día del retiro.
- [ ] Activo enviado al flujo de reinspección en planta.

## 8. KPI

- **% de activos recolectados sobre pendientes de recolección** en la ventana de tiempo definida (ver
  [Logística inversa](../catalogo-systems.md#9-logística-inversa) del catálogo Systems).
- Tiempo detenido antes de recolección: días entre la entrada a `Retenido` y la confirmación de `En tránsito`.
- Costo de recolección por ruta / por activo.
- % de recolecciones que derivan en incidencia (custodio se niega, disputa de condición) sobre el total de
  recolecciones planeadas.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Activo queda `Retenido` indefinidamente porque nadie planea su recolección de forma proactiva | Ruta de recolección planeada por zona/frecuencia regular, no solo reactiva a la alerta individual `[VALIDAR periodicidad de rutas]`. |
| Recolección ejecutada sin confirmar ubicación/custodio actualizado, generando viajes en vacío | Consulta obligatoria a trazabilidad y gestión de custodios antes de despachar transporte (paso 2). |
| Custodio se niega a entregar el activo y el operador de campo fuerza el retiro sin documentarlo | Prohibición explícita de forzar el retiro; toda negativa se registra como incidencia y, si corresponde, escala a SOP-ACTIVO-17 (paso 7). |
| Activo recolectado sin registrar su condición, perdiendo evidencia de daño ocurrido durante la retención | Registro de condición con evidencia fotográfica obligatorio en el acta de retiro (paso 8). |
| Recolecciones concentradas en pocos custodios sistemáticamente sin que Gobernanza lo detecte | Indicador de % recolectado por custodio alimenta [Analítica](../catalogo-systems.md#14-analítica) y [Gobernanza](../catalogo-systems.md#15-gobernanza) para identificar patrones de retención. |

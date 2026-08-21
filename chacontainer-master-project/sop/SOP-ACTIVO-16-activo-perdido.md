# SOP-ACTIVO-16: Activo perdido

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre la declaración de un activo como perdido — las transiciones `Retenido → Perdido`
(transición #20) y `Bloqueado → Perdido` (transición #22) de la
[máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones) — y sus dos salidas posibles:
`Perdido → En tránsito` si se recupera (transición #23) o `Perdido → Baja` si se agota el plazo de búsqueda
(transición #24). Se construyó a partir de la fila "Activo perdido" en
[ciclo-de-vida-del-activo.md §4](../ciclo-de-vida-del-activo.md#4-ramas-de-excepción-fuera-del-flujo-feliz),
el servicio [Recuperación (#12)](../catalogo-solutions.md#12-recuperación) del catálogo Solutions y la
capacidad [Gestión de incidencias (#10)](../catalogo-systems.md#10-gestión-de-incidencias) del catálogo
Systems — **no a partir de una entrevista real con el responsable operativo**. No debe declararse estándar
hasta validar en operación los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Un activo entra en riesgo de pérdida cuando lleva `Retenido` sin recolección posible, o cuando una disputa
en estado `Bloqueado` se resuelve sin que el activo pueda ubicarse. Este SOP define cuándo se agotan los
pasos de búsqueda previos y se declara formalmente el activo como `Perdido`, cómo se documenta esa
declaración, y qué pasa después: si el activo aparece y se recupera, vuelve al flujo normal (`En tránsito`);
si se agota el plazo o el costo de búsqueda deja de justificarse, se da de baja (`Baja`) y sale
definitivamente del parque circulante. Es el punto donde
[Recuperación](../catalogo-solutions.md#12-recuperación) de Solutions y
[Gestión de incidencias](../catalogo-systems.md#10-gestión-de-incidencias) de Systems trabajan juntas para
cerrar, con evidencia, uno de los escenarios que más directamente alimenta el indicador de **pérdida**
([ETAPA 6 del README](../README.md#etapa-6--indicadores-del-sistema)).

## 2. Objetivo

Garantizar que ningún activo se declare `Perdido` sin haber agotado un proceso de búsqueda documentado, y
que todo activo `Perdido` tenga una resolución explícita en un plazo definido: recuperación o baja formal —
nunca queda indefinidamente en ese estado.

## 3. Alcance

Aplica a todo activo que:

1. Está en estado `Retenido` y, tras el proceso de recolección de
   [SOP-ACTIVO-14](./SOP-ACTIVO-14-logistica-inversa.md), no puede localizarse ni recolectarse
   (transición #20, `Retenido → Perdido`).
2. Está en estado `Bloqueado` (ver [SOP-ACTIVO-17](./SOP-ACTIVO-17-activo-bloqueado.md)) y la disputa se
   resuelve sin que exista posibilidad de recuperación (transición #22, `Bloqueado → Perdido`).
3. Se reporta como perdido de forma explícita por un cliente, custodio u operador, sin haber pasado
   necesariamente por `Retenido` primero `[VALIDAR: ¿un reporte directo de pérdida siempre pasa antes por
   Retenido, o puede declararse Perdido directamente en casos evidentes — robo, siniestro?]`.

Cubre desde la sospecha de pérdida (agotamiento de pasos de búsqueda) hasta la resolución final del estado
`Perdido`: recuperación (`Perdido → En tránsito`, transición #23) o baja (`Perdido → Baja`, transición #24).
No incluye la ejecución operativa de la recolección en sí cuando el activo sí tiene ubicación confirmable
(eso es [SOP-ACTIVO-14](./SOP-ACTIVO-14-logistica-inversa.md)), ni el proceso de baja/scrap/disposición una
vez que el activo llega a `Baja` (SOP-ACTIVO-18/19/20 del ciclo normal).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Logística / Recuperación `[VALIDAR nombre exacto del puesto]` | Ejecuta los pasos de búsqueda previos a la declaración de pérdida (ver [Recuperación](../catalogo-solutions.md#12-recuperación)). |
| Responsable de gestión de incidencias | Registra y da seguimiento al ticket de incidencia obligatorio asociado al estado `Perdido` (ver [SOP-ACTIVO-15](./SOP-ACTIVO-15-incidencias.md)). |
| Gobernanza `[VALIDAR quién tiene autoridad exacta de resolución]` | Autoriza la declaración formal de `Perdido` tras agotar la búsqueda, decide si se continúa buscando o se pasa a `Baja`, y es quien dispara la transición `Bloqueado → Perdido` (fila 22) cuando la pérdida se origina en una disputa resuelta sin recuperación. |
| Comercial `[VALIDAR]` | Comunica al cliente/custodio la declaración de pérdida cuando la relación contractual lo requiere, y gestiona cualquier cobro o compensación asociada `[VALIDAR si existe penalización contractual por activo perdido]`. |

## 5. Diagrama de flujo

```
[Activo "Retenido" sin recolección posible, o "Bloqueado"
 resuelto sin posibilidad de recuperación]
                ↓
[Ejecutar pasos de búsqueda agotables:
 - reconfirmar última ubicación en trazabilidad
 - contactar al custodio registrado
 - verificar contra inventario físico más reciente
 - intento de recolección en sitio]
                ↓
        ¿Se localiza el activo?
        ↓ sí                              ↓ no (búsqueda agotada)
[Continuar en SOP-ACTIVO-14              [Declarar "Perdido"
 · Logística inversa]                     (transición #20 o #22)]
                                                    ↓
                                          [Ticket de incidencia obligatorio
                                           (SOP-ACTIVO-15) — invariante 3]
                                                    ↓
                                          [Definir plazo de búsqueda extendida
                                           `[VALIDAR duración del plazo]`]
                                                    ↓
                                          ¿Activo localizado dentro del plazo?
                                          ↓ sí                    ↓ no (plazo agotado)
                                   [Recuperar activo         [Gobernanza autoriza
                                    → "En tránsito"            baja → "Baja"]
                                    (transición #23)]           (transición #24)
                                          ↓                          ↓
                                   [Reinspección en planta]   [Continúa en SOP-ACTIVO-18]
```

## 6. SOP paso a paso

1. Identificar el disparador: un activo `Retenido` sin recolección posible tras
   [SOP-ACTIVO-14](./SOP-ACTIVO-14-logistica-inversa.md), un activo `Bloqueado` cuya disputa se resuelve sin
   posibilidad de recuperación (ver [SOP-ACTIVO-17](./SOP-ACTIVO-17-activo-bloqueado.md)), o un reporte
   explícito de pérdida.
2. Agotar los pasos de búsqueda previos antes de declarar la pérdida, como mínimo
   `[VALIDAR lista exhaustiva de pasos de búsqueda obligatorios]`:
   - Reconfirmar la última ubicación conocida en [Trazabilidad](../catalogo-systems.md#7-trazabilidad).
   - Contactar al custodio registrado en [Gestión de custodios](../catalogo-systems.md#8-gestión-de-custodios).
   - Verificar contra el [inventario físico](../catalogo-solutions.md#11-inventarios-físicos) más reciente
     disponible para esa ubicación/cliente.
   - Intentar un retiro/recolección directo en sitio si hay indicio de ubicación.
3. Si alguno de estos pasos localiza el activo, detener este SOP y continuar el proceso de recolección en
   [SOP-ACTIVO-14](./SOP-ACTIVO-14-logistica-inversa.md).
4. Si la búsqueda se agota sin localizar el activo, declarar formalmente el estado `Perdido`:
   - Desde `Retenido`, transición #20 (`Retenido → Perdido`) — disparada por sistema automático o
     Gobernanza tras exceder el umbral adicional de retención `[VALIDAR: ¿qué umbral, en días, separa
     "Retenido" de "Perdido"?]`.
   - Desde `Bloqueado`, transición #22 (`Bloqueado → Perdido`) — disparada por Gobernanza cuando la disputa
     se resuelve sin posibilidad de recuperación.
5. Registrar el ticket de incidencia obligatorio asociado (ver [SOP-ACTIVO-15](./SOP-ACTIVO-15-incidencias.md)
   e [invariante 3](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)): ningún
   activo puede quedar en `Perdido` sin su ticket correspondiente.
6. Definir el plazo de búsqueda extendida a partir de la declaración de `Perdido`, durante el cual sigue
   siendo válido localizar y recuperar el activo `[VALIDAR duración exacta del plazo — probablemente distinta
   por tipo/valor de activo]`.
7. Si el activo se localiza y se recupera dentro del plazo, ejecutar la recuperación
   ([Recuperación](../catalogo-solutions.md#12-recuperación) de Solutions), actualizar el estado a
   `En tránsito` (transición #23) y enviarlo a reinspección en planta como cualquier retorno.
8. Si el plazo se agota sin recuperación, Gobernanza autoriza formalmente la baja del activo:
   `[VALIDAR quién tiene autoridad exacta — responsable de planta, dirección comercial, u otro rol de
   Gobernanza]`.
9. Actualizar el estado a `Baja` (transición #24) y continuar el proceso en SOP-ACTIVO-18 (Baja) del ciclo
   normal, dejando cerrado el ticket de incidencia con la resolución "no recuperado — dado de baja".
10. Actualizar el indicador de pérdida (ver [§8 KPI](#8-kpi)).

## 7. Checklist

- [ ] Pasos de búsqueda mínimos ejecutados y documentados antes de declarar `Perdido` (no se declara pérdida por omisión o abandono del seguimiento).
- [ ] Declaración de `Perdido` respaldada por la transición correcta (#20 desde `Retenido`, #22 desde `Bloqueado`).
- [ ] Ticket de incidencia registrado y vinculado al activo `Perdido` (invariante 3 verificado).
- [ ] Plazo de búsqueda extendida definido y comunicado al responsable de seguimiento.
- [ ] Si se recupera: acta de recuperación, condición registrada, estado actualizado a `En tránsito`.
- [ ] Si se agota el plazo: autorización explícita de Gobernanza registrada antes de mover el activo a `Baja`.
- [ ] Ticket de incidencia cerrado con la resolución final (recuperado / dado de baja) y fecha de cierre.
- [ ] Cliente/custodio notificado cuando la relación contractual lo requiere.

## 8. KPI

- **Tasa de pérdida**: activos no recuperados / activos administrados × 100 ([ETAPA 6 del README](../README.md#etapa-6--indicadores-del-sistema)).
- % de activos declarados `Perdido` que se recuperan dentro del plazo de búsqueda extendida.
- Tiempo promedio entre la declaración de `Perdido` y su resolución final (recuperado o dado de baja).
- Costo de búsqueda/recuperación por activo perdido, comparado contra el valor del activo (insumo para
  decidir si conviene seguir buscando — ver [Analítica](../catalogo-systems.md#14-analítica)).
- Concentración de pérdidas por cliente, custodio o ruta (responde "¿dónde se producen las pérdidas?" —
  [ETAPA 4 del README](../README.md#etapa-4--modelo-de-gobernanza)).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Activo declarado `Perdido` sin agotar pasos de búsqueda reales (por conveniencia o falta de tiempo) | Lista mínima de pasos de búsqueda documentada y obligatoria antes de la declaración (paso 2) `[VALIDAR lista exhaustiva]`. |
| Activo queda en `Perdido` indefinidamente, sin plazo ni resolución final | Plazo de búsqueda extendida obligatorio al declarar la pérdida (paso 6), con resolución forzada al vencer: recuperación o baja. |
| Declaración de `Perdido` sin ticket de incidencia asociado, rompiendo el invariante 3 de la máquina de estados | Registro de ticket obligatorio como parte del mismo paso de declaración (paso 5), verificado en checklist de cierre. |
| Decisión de dar de baja tomada sin autoridad clara, generando disputas internas sobre quién autorizó la pérdida del activo | Autorización explícita de Gobernanza requerida antes de ejecutar la transición `Perdido → Baja` (paso 8) `[VALIDAR rol exacto con autoridad]`. |
| Costo de búsqueda/recuperación mayor al valor del activo, pero se sigue buscando por inercia | Comparación explícita costo de recuperación vs. valor del activo como parte de la decisión de Gobernanza (indicador §8), no solo del vencimiento del plazo. |

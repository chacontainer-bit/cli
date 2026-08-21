# SOP-ACTIVO-17: Activo bloqueado

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre las disputas de custodia o condición sobre un activo — la transición
`Retenido → Bloqueado` (transición #19) de la
[máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones) — y su resolución por Gobernanza,
que puede devolver el activo al flujo normal (`Bloqueado → En tránsito` o `Bloqueado → Retenido`,
transición #21) o confirmar que no hay recuperación posible (`Bloqueado → Perdido`, transición #22). Se
construyó a partir de la fila "Activo bloqueado" en
[ciclo-de-vida-del-activo.md §4](../ciclo-de-vida-del-activo.md#4-ramas-de-excepción-fuera-del-flujo-feliz),
la capacidad [Gestión de incidencias (#10)](../catalogo-systems.md#10-gestión-de-incidencias) y la capacidad
[Gobernanza (#15)](../catalogo-systems.md#15-gobernanza) del catálogo Systems — **no a partir de una
entrevista real con el responsable operativo**. No debe declararse estándar hasta validar en operación los
puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Un activo se bloquea cuando, estando `Retenido`, surge una disputa formal que impide simplemente
recolectarlo y seguir el flujo normal: el custodio niega tenerlo, cuestiona quién es responsable de un daño,
o discute los términos de su devolución. `Bloqueado` es un estado de excepción propio — no existía en la
lista de la ETAPA 3 del proyecto maestro — creado específicamente porque una disputa de custodia o condición
no es lo mismo que un simple retraso: requiere que alguien con autoridad de decisión (Gobernanza) resuelva
el desacuerdo antes de que el activo pueda volver a moverse. Este SOP define cómo se abre el bloqueo, qué
evidencia se recopila mientras está bloqueado, y cómo Gobernanza lo resuelve: liberándolo de vuelta al
flujo, o confirmando que no hay recuperación posible.

## 2. Objetivo

Garantizar que ninguna disputa de custodia o condición detenga un activo de forma indefinida ni informal:
todo bloqueo queda documentado con evidencia, tiene un responsable de resolución con autoridad clara, y se
resuelve dentro de un plazo definido hacia uno de los tres destinos posibles de la transición #21/#22.

## 3. Alcance

Aplica a todo activo en estado `Retenido` sobre el cual surge una disputa formal de:

1. **Custodia** — el custodio registrado niega tener el activo, o hay desacuerdo sobre quién debería
   tenerlo (por ejemplo, transferencia informal no registrada entre dos custodios).
2. **Condición** — hay desacuerdo sobre el estado del activo al momento del intento de recolección (el
   custodio alega que ya estaba dañado antes de recibirlo, o CHACONTAINER detecta un daño que el custodio
   no reconoce).

Cubre desde el reporte de la disputa (típicamente originado durante un intento de recolección de
[SOP-ACTIVO-14](./SOP-ACTIVO-14-logistica-inversa.md)) hasta la resolución de Gobernanza. No incluye la
ejecución de la recolección en sí cuando no hay disputa (eso es
[SOP-ACTIVO-14](./SOP-ACTIVO-14-logistica-inversa.md)), ni la declaración de pérdida cuando el bloqueo se
resuelve sin posibilidad de recuperación (eso continúa en
[SOP-ACTIVO-16 · Activo perdido](./SOP-ACTIVO-16-activo-perdido.md) una vez ejecutada la transición #22).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Comercial / Custodio `[VALIDAR nombre exacto del puesto que reporta la disputa]` | Reporta la disputa cuando surge durante un intento de recolección o cualquier otro punto de contacto con el custodio; recopila la versión del custodio. |
| Responsable de gestión de incidencias | Registra el ticket de incidencia obligatorio asociado al bloqueo (ver [SOP-ACTIVO-15](./SOP-ACTIVO-15-incidencias.md)). |
| **Gobernanza** `[VALIDAR quién tiene autoridad exacta de resolución — ¿dirección comercial, responsable legal, comité interno?]` | Tiene la autoridad exclusiva para resolver el bloqueo: dispara las transiciones `Bloqueado → En tránsito`, `Bloqueado → Retenido` o `Bloqueado → Perdido` (filas 21 y 22 de la [tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones)). Ningún otro rol puede cerrar un bloqueo. |
| Legal/Comercial `[VALIDAR si aplica]` | Apoya cuando la disputa tiene implicación contractual o de facturación (p. ej. cobro por activo dañado o no devuelto). |

## 5. Diagrama de flujo

```
[Activo "Retenido" — surge disputa de custodia o condición,
 típicamente durante intento de recolección (SOP-ACTIVO-14)]
                ↓
[Registrar la disputa: qué se disputa, quién la origina, evidencia inicial]
                ↓
[Transición "Retenido → Bloqueado" (#19)]
                ↓
[Ticket de incidencia obligatorio (SOP-ACTIVO-15) — invariante 3]
                ↓
[Recopilar evidencia de ambas partes:
 - versión del custodio
 - trazabilidad/historial de custodia del activo
 - evidencia fotográfica de condición si aplica]
                ↓
[Escalar a Gobernanza con expediente completo]
                ↓
        Gobernanza resuelve la disputa
        ↓                    ↓                      ↓
[A favor de           [Sin resolución          [Sin posibilidad
 continuar el ciclo:   clara todavía:           de recuperación:
 "En tránsito" o        vuelve a "Retenido"      "Perdido"]
 "Retenido"]            para seguimiento]        (transición #22)
 (transición #21)       (transición #21)               ↓
        ↓                    ↓                  [Continúa en
[Continúa flujo      [Continúa en                SOP-ACTIVO-16]
 normal /              SOP-ACTIVO-14]
 SOP-ACTIVO-14]
```

## 6. SOP paso a paso

1. Detectar la disputa: normalmente surge durante un intento de recolección
   ([SOP-ACTIVO-14](./SOP-ACTIVO-14-logistica-inversa.md)), pero puede originarse en cualquier punto de
   contacto con el custodio (reclamo del custodio, hallazgo de inconsistencia en trazabilidad).
2. Registrar la naturaleza de la disputa: si es de custodia (quién tiene o debería tener el activo) o de
   condición (en qué estado estaba y quién es responsable), quién la origina y la evidencia inicial
   disponible.
3. Ejecutar la transición `Retenido → Bloqueado` (transición #19) — quien la dispara es el rol Comercial/
   Custodio que reporta la disputa, según la [tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones).
4. Registrar el ticket de incidencia obligatorio asociado al bloqueo (ver
   [SOP-ACTIVO-15](./SOP-ACTIVO-15-incidencias.md) e
   [invariante 3](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)): ningún
   activo puede quedar en `Bloqueado` sin su ticket correspondiente.
5. Recopilar el expediente de la disputa: la versión del custodio, el historial de
   [Trazabilidad](../catalogo-systems.md#7-trazabilidad) y
   [Gestión de custodios](../catalogo-systems.md#8-gestión-de-custodios) del activo, y evidencia fotográfica
   de condición cuando la disputa es sobre el estado del activo.
6. Escalar el expediente completo a Gobernanza. Ningún rol distinto a Gobernanza puede resolver el bloqueo
   `[VALIDAR: ¿existe un umbral de valor del activo bajo el cual un rol operativo puede resolver disputas
   menores sin escalar a Gobernanza?]`.
7. Gobernanza evalúa el expediente y resuelve hacia uno de tres destinos (transición #21 o #22):
   - **A favor de continuar el ciclo**: el activo puede recolectarse — pasa a `En tránsito` y continúa en
     [SOP-ACTIVO-14](./SOP-ACTIVO-14-logistica-inversa.md), o vuelve a `Retenido` si la disputa se resuelve
     pero la recolección física aún está pendiente.
   - **Sin posibilidad de recuperación**: la disputa confirma que el activo no puede recuperarse (por
     ejemplo, el custodio confirma pérdida o destrucción) — pasa a `Perdido` (transición #22) y continúa en
     [SOP-ACTIVO-16 · Activo perdido](./SOP-ACTIVO-16-activo-perdido.md).
8. Documentar la decisión de Gobernanza en el ticket de incidencia: qué se resolvió, con base en qué
   evidencia, y quién autorizó la resolución.
9. Actualizar el estado del activo según la decisión y cerrar (o mantener abierto, si vuelve a `Retenido`
   para seguimiento) el ticket de incidencia asociado.
10. Actualizar el indicador de tiempo promedio de resolución de bloqueos (ver [§8 KPI](#8-kpi)).

## 7. Checklist

- [ ] Disputa registrada con su naturaleza (custodia / condición), origen y evidencia inicial.
- [ ] Transición `Retenido → Bloqueado` ejecutada y estado actualizado en el sistema el mismo día del reporte.
- [ ] Ticket de incidencia registrado y vinculado al activo `Bloqueado` (invariante 3 verificado).
- [ ] Expediente completo (versión del custodio, trazabilidad, custodia, evidencia fotográfica) recopilado antes de escalar a Gobernanza.
- [ ] Resolución tomada exclusivamente por Gobernanza, documentada con base y autoridad explícitas.
- [ ] Estado del activo actualizado según la resolución (`En tránsito` / `Retenido` / `Perdido`).
- [ ] Si la resolución es `Perdido`, el caso continúa formalmente en SOP-ACTIVO-16.
- [ ] Ticket de incidencia cerrado con la resolución final, o mantenido abierto explícitamente si el activo vuelve a `Retenido` para seguimiento.

## 8. KPI

- Tiempo promedio de resolución de un bloqueo, desde `Retenido → Bloqueado` hasta la resolución de Gobernanza.
- % de bloqueos resueltos a favor de continuar el ciclo vs. resueltos como `Perdido`.
- Número de bloqueos por custodio/cliente — insumo directo para "qué proveedor retiene activos"
  ([ETAPA 4 del README](../README.md#etapa-4--modelo-de-gobernanza)).
- % de bloqueos cuya causa raíz (según Gobernanza) es atribuible a reglas operativas ambiguas o no
  comunicadas al custodio (insumo para ajustar [Reglas operativas](../catalogo-systems.md#11-reglas-operativas)).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Disputa resuelta informalmente por el operador de campo o comercial, sin pasar por Gobernanza | Autoridad de resolución restringida explícitamente a Gobernanza (paso 6-7); ningún otro rol puede ejecutar la transición #21/#22. |
| Activo queda `Bloqueado` indefinidamente porque nadie escala el expediente a Gobernanza | Escalamiento obligatorio inmediato tras recopilar el expediente inicial (paso 6); indicador de tiempo de resolución visible para Gobernanza (§8). |
| Resolución de Gobernanza sin evidencia documentada, generando disputas sobre la disputa misma | Expediente completo (versión del custodio + trazabilidad + evidencia fotográfica) obligatorio antes de escalar (paso 5); decisión documentada con base explícita (paso 8). |
| Bloqueo declarado sin ticket de incidencia asociado, rompiendo el invariante 3 de la máquina de estados | Registro de ticket obligatorio en el mismo paso que la transición a `Bloqueado` (paso 4), verificado en checklist de cierre. |
| Patrón de bloqueos recurrentes con el mismo custodio/cliente no se detecta a tiempo | Indicador de bloqueos por custodio (§8) alimenta [Analítica](../catalogo-systems.md#14-analítica) y revisión periódica de Gobernanza. |

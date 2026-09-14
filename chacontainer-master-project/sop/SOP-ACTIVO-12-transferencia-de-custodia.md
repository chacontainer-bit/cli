# SOP-ACTIVO-12: Transferencia de custodia

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre el cambio de responsable de un activo retornable (Capa 3 · Custodia) que ocurre
**sin que el activo pase físicamente por CHACONTAINER** — por ejemplo, el cliente subcontrata a un tercero
para operar el activo, o lo transfiere de una de sus plantas a otra. No corresponde a ninguna transición de
[Capa 4 · Estado](../estados-del-activo.md#1-catálogo-de-estados) por sí sola: el activo permanece en
`En uso` o `En tránsito` durante todo el proceso, y solo cambia el custodio registrado en la capacidad
[Gestión de custodios](../catalogo-systems.md#8-gestión-de-custodios). Se construyó a partir del cruce entre
el [catálogo Systems](../catalogo-systems.md#8-gestión-de-custodios) y la máquina de estados — **no a partir
de una entrevista real con el responsable comercial o de cuentas**. No debe declararse estándar hasta
validar en operación los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Transferencia de custodia es el proceso que actualiza quién responde por un activo cuando ese cambio ocurre
fuera del control operativo directo de CHACONTAINER: el cliente decide subcontratar a un tercero, o mover el
activo entre sus propias plantas o áreas, sin devolverlo a CHACONTAINER ni pasar por
`Disponible`. A diferencia de [Movimiento](./SOP-ACTIVO-11-movimiento.md), que registra un desplazamiento
físico que CHACONTAINER ejecuta o coordina, este SOP registra un cambio de responsabilidad que CHACONTAINER
normalmente **solo se entera por notificación** — el activo puede o no desplazarse físicamente, pero eso no
es lo que dispara el proceso. Sin este SOP, la Capa 3 queda desactualizada apenas un cliente reorganiza
internamente su operación, y CHACONTAINER pierde la cadena de responsabilidad exigida por
[Gestión de custodios](../catalogo-systems.md#8-gestión-de-custodios) justo en el escenario donde más se
necesita: cuando algo se pierde o se daña bajo un custodio que ni siquiera está en el registro.

## 2. Objetivo

Garantizar que todo cambio de responsable de un activo que ocurre durante `En uso` o `En tránsito` — sin
retorno a `Disponible` y sin intervención física de CHACONTAINER — queda registrado en Capa 3 · Custodia con
el custodio anterior conservado en el historial, de forma que en cualquier momento se pueda identificar
quién responde por el activo.

## 3. Alcance

Aplica a todo cambio de custodio (Capa 3) que ocurre mientras el activo está en estado `En uso` o
`En tránsito`, iniciado por el cliente o custodio actual, sin que el activo regrese a `Disponible` ni pase
por instalaciones de CHACONTAINER. Casos típicos:

1. **Subcontratación a un tercero** — el cliente entrega el activo a un operador/transportista externo para
   que lo use u opere en su nombre.
2. **Transferencia entre plantas o áreas del mismo cliente** — el activo se mueve de una planta a otra del
   mismo custodio principal, cambiando el custodio de detalle (planta/área/usuario, según la granularidad de
   Capa 3 definida en la [ETAPA 3 del README](../README.md#etapa-3--modelo-de-packaging-systems)).

**No incluye:**

- La primera entrega de un activo desde `Disponible`. Esa es
  [SOP-ACTIVO-10 (Asignación)](./SOP-ACTIVO-10-asignacion.md): dispara `Disponible → Asignado → En tránsito`,
  la ejecuta y confirma CHACONTAINER, y crea el primer registro de custodia del ciclo. Transferencia de
  custodia, en cambio, **no** parte de `Disponible` — solo modifica un custodio que ya existía.
- El desplazamiento físico que CHACONTAINER coordina — ver
  [SOP-ACTIVO-11 (Movimiento)](./SOP-ACTIVO-11-movimiento.md).
- El retorno del activo a CHACONTAINER — ver [SOP-ACTIVO-13 (Retorno)](./SOP-ACTIVO-13-retorno.md).

### Diferencia con SOP-ACTIVO-10 (Asignación)

| | SOP-ACTIVO-10 · Asignación | SOP-ACTIVO-12 · Transferencia de custodia |
|---|---|---|
| Punto de partida | `Disponible` (activo en parque de CHACONTAINER) | `En uso` o `En tránsito` (activo ya fuera, con custodio previo) |
| Quién ejecuta el movimiento físico | CHACONTAINER (Solutions: renta/venta) | El cliente/custodio, fuera del control operativo de CHACONTAINER |
| Transición de Capa 4 que dispara | `Disponible → Asignado` | Ninguna — el estado permanece `En uso`/`En tránsito` |
| Qué actualiza | Capa 3 (primer custodio) y, después, Capa 2 vía Movimiento | Solo Capa 3 (nuevo custodio); Capa 2 se actualiza como referencia, no como trazabilidad propia de CHACONTAINER |
| Origen del evento | Solicitud comercial (renta/venta) | Notificación del cliente/custodio actual |

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Custodio saliente (cliente o custodio actual) | Origina y notifica el cambio de responsable a CHACONTAINER. |
| Custodio entrante (tercero o planta/área receptora) | Recibe la responsabilidad del activo `[VALIDAR si debe confirmar recepción directamente a CHACONTAINER o basta la notificación del custodio saliente]`. |
| Responsable de gestión de custodios `[VALIDAR nombre exacto del puesto — ¿responsable de cuenta comercial u operación?]` | Valida la notificación, registra el nuevo custodio en el sistema y conserva el historial. |
| Gobernanza | Revisa si el nuevo custodio implica reglas operativas distintas (tiempo máximo fuera, condición permitida) y ajusta alertas. |

## 5. Diagrama de flujo

```
[Activo en estado "En uso" o "En tránsito" — custodio actual = Cliente/Custodio A]
                ↓
[Custodio A notifica cambio de responsable:
 subcontratación a tercero / traslado a otra planta o área propia]
                ↓
[Responsable de gestión de custodios valida la notificación
 (identidad del activo, autorización de quien notifica)]
                ↓
        ¿Notificación válida y activo identificable?
        ↓ sí                                       ↓ no
[Registrar nuevo custodio en Capa 3              [Rechazar / solicitar aclaración —
 (Gestión de custodios); custodio anterior         el activo conserva el custodio
 se conserva en el historial, no se sobrescribe]   registrado hasta resolver]
        ↓
[Estado (Capa 4) NO cambia — permanece "En uso" / "En tránsito"]
        ↓
[Verificar si el nuevo custodio requiere reglas operativas distintas
 → ajustar Reglas operativas / Alertas si aplica]
        ↓
[Cerrar: ficha del activo refleja el nuevo custodio con fecha de transferencia]
```

## 6. SOP paso a paso

1. Detectar el evento — a diferencia de Movimiento y Asignación, este proceso normalmente **lo inicia el
   cliente**, no CHACONTAINER, mediante notificación de subcontratación o de traslado interno.
2. Confirmar la identidad del activo (ID, QR/RFID) y verificar que su estado actual es `En uso` o
   `En tránsito` bajo el custodio saliente que notifica el cambio.
3. Validar la autorización de quien notifica el cambio
   `[VALIDAR mecanismo — ¿basta la notificación del custodio saliente, o se requiere confirmación explícita
   del custodio entrante antes de registrar la transferencia?]`.
4. Si la notificación no es válida o el activo no es identificable, rechazarla y mantener el custodio
   registrado previamente hasta que se resuelva la discrepancia con el responsable de gestión de custodios.
5. Si la notificación es válida, registrar el nuevo custodio en Capa 3 mediante
   [Gestión de custodios](../catalogo-systems.md#8-gestión-de-custodios), **anexando** el registro — el
   custodio anterior queda en el historial, no se elimina.
6. Confirmar explícitamente que el estado del activo (Capa 4) no se altera por esta transacción: sigue
   `En uso` o `En tránsito` según corresponda; solo cambió el responsable.
7. Revisar si el custodio entrante tiene [Reglas operativas](../catalogo-systems.md#11-reglas-operativas)
   distintas a las del custodio saliente (tiempo máximo fuera, condición permitida) y ajustar las
   [Alertas](../catalogo-systems.md#12-alertas) correspondientes si aplica.
8. Si la transferencia coincide con un desplazamiento físico ejecutado por el propio cliente (no por
   CHACONTAINER), registrar la ubicación reportada en Capa 2 como referencia informativa, dejando explícito
   que no proviene de un evento de trazabilidad propio de CHACONTAINER
   `[VALIDAR si esta ubicación referencial se distingue visualmente de una ubicación confirmada por
   escaneo/Movimiento]`.
9. Dejar evidencia de la transferencia (notificación recibida, fecha, custodio saliente y entrante)
   `[VALIDAR si se requiere un documento firmado tipo "acta de transferencia" o basta el registro digital
   de la notificación]`.
10. Actualizar el indicador de transferencias de custodia registradas por cliente/periodo.

## 7. Checklist

- [ ] Identidad del activo confirmada y estado verificado como `En uso` o `En tránsito`.
- [ ] Notificación de cambio de custodio recibida y validada (no inferida ni asumida).
- [ ] Nuevo custodio registrado en Capa 3 sin sobrescribir el historial del custodio anterior.
- [ ] Estado (Capa 4) confirmado sin cambios — no se transiciona a `Disponible` ni a ningún otro estado por
      esta transacción.
- [ ] Reglas operativas/alertas revisadas contra el perfil del nuevo custodio.
- [ ] Evidencia de la notificación (fecha, custodio saliente y entrante) registrada en la ficha del activo.
- [ ] Transferencia registrada en un plazo razonable desde la notificación del cliente
      `[VALIDAR plazo máximo]`.

## 8. KPI

- Número de transferencias de custodia registradas por cliente/periodo.
- % de transferencias notificadas por el cliente vs. detectadas tardíamente (por ejemplo, en un inventario
  físico o incidencia) — mide el riesgo de Capa 3 desactualizada.
- Tiempo entre la fecha real de la transferencia (según el cliente) y su registro en el sistema.
- % de activos con custodio desactualizado detectado en auditoría o inventario físico — alimenta la
  pregunta "¿quién los tiene?" de la [ETAPA 4 del README](../README.md#etapa-4--modelo-de-gobernanza).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Cliente subcontrata o transfiere el activo sin notificar a CHACONTAINER — Capa 3 queda desactualizada y se pierde la cadena de responsabilidad | Establecer contractualmente la obligación de notificar transferencias de custodia; detectar discrepancias en auditorías de custodia periódicas `[VALIDAR periodicidad]`. |
| Tercero receptor no tiene relación contractual directa con CHACONTAINER — no está claro a quién exigir el retorno | Registrar explícitamente si el custodio entrante es una parte con acuerdo directo o un subcustodio bajo responsabilidad del custodio original `[VALIDAR modelo de responsabilidad solidaria o subsidiaria]`. |
| Confusión entre este SOP y Asignación — se intenta usar transferencia de custodia como atajo para una primera entrega, saltándose `Disponible → Asignado` | Reforzar la distinción de la [§3](#3-alcance): este SOP nunca parte de `Disponible`; si el activo está `Disponible`, el proceso correcto es siempre Asignación. |
| Custodio entrante no cumple las reglas operativas del tipo de activo (p. ej. condición de almacenamiento) | Verificación obligatoria de reglas operativas del nuevo custodio antes de cerrar el registro (paso 7). |
| Notificación de transferencia aceptada sin validar autorización real de quien la envía | Definir por escrito el mecanismo de autorización mínimo `[VALIDAR]` antes de registrar cualquier cambio de custodio. |

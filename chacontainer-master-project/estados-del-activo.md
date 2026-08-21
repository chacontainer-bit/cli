# Estados del activo — máquina de estados (Capa 4)

**Fase:** I · Fundación — paso 4 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

Este documento toma la columna "Estado resultante" de la tabla del
[ciclo de vida del activo](./ciclo-de-vida-del-activo.md#2-tabla-de-trazabilidad-cruzada-por-etapa)
y la convierte en una máquina de estados explícita: estados válidos,
transiciones permitidas, qué las dispara y quién (o qué) tiene autoridad
para disparar cada una. Es la especificación funcional del campo `estado`
del registro maestro de activos — el dato central del Módulo 1 de
[CHACONTAINER OS](./catalogo-systems.md#16-chacontainer-os).

> **Nota de alcance:** la [ETAPA 3](./README.md#etapa-3--modelo-de-packaging-systems)
> del proyecto maestro lista 10 estados de ejemplo para la Capa 4
> (Disponible, Asignado, En tránsito, En uso, Retenido, Sucio, Dañado,
> Reparación, Perdido, Scrap). Este documento los conserva todos pero los
> **extiende a 16** porque el ciclo de vida y los SOP-15 a SOP-20 exigen
> distinguir estados que la lista original agrupaba de forma implícita
> (p. ej. "Dañado" se separa en `En reparación` como estado activo y
> `Baja` como decisión terminal; se agrega `Bloqueado` para cubrir SOP-17,
> que no tenía estado propio). Cualquier estado nuevo respecto a la ETAPA 3
> queda marcado como tal.

---

## 1. Catálogo de estados

### Operativos (ciclo normal del activo)

| Estado | Significado | ¿Nuevo vs. ETAPA 3? |
|---|---|---|
| `Registrado` | Alta completa en Capa 1, sin identificación física todavía. | Nuevo — implícito en Alta, sin nombre propio en la ETAPA 3. |
| `Identificado` | Tiene QR/RFID vinculado, pendiente de primera inspección. | Nuevo — implícito en Identificación. |
| `En inspección` | Bajo evaluación de condición (inspección inicial o reinspección). | Nuevo — cubre tanto Inspección como Reinspección del ciclo. |
| `Sucio` | Requiere lavado antes de continuar. | De la ETAPA 3. |
| `En reparación` | Daño confirmado, en intervención activa (reparación o reacondicionamiento). | De la ETAPA 3 (`Reparación`), renombrado para distinguirlo de `Dañado` como diagnóstico vs. `En reparación` como proceso en curso. |
| `Liberado` | Pasó inspección/reparación/reetiquetado; pendiente de entrar a inventario disponible. | Nuevo — checkpoint entre "apto" y "disponible para asignar". |
| `Disponible` | Apto y listo para asignar. | De la ETAPA 3. |
| `Asignado` | Con custodio confirmado, pendiente de salida física. | De la ETAPA 3. |
| `En tránsito` | En movimiento físico: salida, entre ubicaciones intermedias, o en retorno. | De la ETAPA 3. |
| `En uso` | En poder del custodio, fuera de instalaciones de CHACONTAINER. | De la ETAPA 3. |
| `Retenido` | Plazo de retorno vencido según reglas operativas, sin incidencia formal todavía. | De la ETAPA 3. |

### De excepción

| Estado | Significado | ¿Nuevo vs. ETAPA 3? |
|---|---|---|
| `Bloqueado` | Disputa de custodia o de condición abierta (SOP-17); no puede avanzar hasta resolverse. | Nuevo — SOP-17 no tenía estado propio en la ETAPA 3. |
| `Perdido` | Sin ubicación confirmable tras exceder el umbral de `Retenido`, o reporte explícito. | De la ETAPA 3. |

### Terminales

| Estado | Significado | ¿Nuevo vs. ETAPA 3? |
|---|---|---|
| `Baja` | Decisión tomada de sacar el activo del parque circulante; destino final aún no ejecutado. | Nuevo — la ETAPA 3 no distinguía "decisión de baja" de "scrap ejecutado". |
| `Scrap` | Comprado/valorado como material, en proceso de recuperación de valor. | De la ETAPA 3. |
| `Dispuesto` | Disposición final ejecutada y certificada. Fin absoluto del ciclo de vida. | Nuevo — cierre que la ETAPA 3 no representaba como estado. |

**Total: 16 estados** (11 operativos + 2 de excepción + 3 terminales).

---

## 2. Tabla de transiciones

| # | Estado origen | Estado destino | Evento que dispara | Quién/qué lo dispara | SOP / capacidad relacionada |
|---|---|---|---|---|---|
| 1 | — (nuevo activo) | `Registrado` | Alta ejecutada | Operador de venta/alta | SOP-01 |
| 2 | `Registrado` | `Identificado` | QR/RFID vinculado al ID | Operador de identificación | SOP-02 |
| 3 | `Identificado` | `En inspección` | Inspección iniciada | Inspector | SOP-03 |
| 4 | `En inspección` | `Sucio` | Resultado: requiere lavado | Inspector | SOP-03 |
| 5 | `En inspección` | `En reparación` | Resultado: daño detectado | Inspector | SOP-03 → SOP-06 |
| 6 | `En inspección` | `Liberado` | Resultado: apto sin intervención | Inspector | SOP-08 |
| 7 | `Sucio` | `En inspección` | Lavado completado, pasa a inspección final | Operador de lavado | SOP-05 |
| 8 | `En reparación` | `En inspección` | Reparación/reacondicionamiento completado, requiere validación | Técnico de reparación | SOP-06 |
| 9 | `En reparación` | `Baja` | Diagnóstico: no reparable | Responsable de planta / Gobernanza | SOP-06 → SOP-18 |
| 10 | `Liberado` | `Disponible` | Ingreso a inventario confirmado | Sistema (automático) o responsable de inventario | SOP-09 |
| 11 | `Disponible` | `Asignado` | Asignación a renta/venta/proyecto | Comercial / Operaciones | SOP-10 |
| 12 | `Asignado` | `En tránsito` | Salida física confirmada | Logística | SOP-11 |
| 13 | `En tránsito` | `En uso` | Entrega en destino confirmada (escaneo o confirmación de custodio) | Sistema (automático, vía escaneo) | SOP-11 |
| 14 | `En tránsito` | `En tránsito` | Movimiento entre ubicaciones intermedias (hub, transbordo) | Logística | SOP-11 |
| 15 | `En tránsito` | `En inspección` | Llegada a planta tras retorno, entra a reinspección | Recepción | SOP-03, SOP-04 |
| 16 | `En uso` | `En tránsito` | Inicio de retorno | Custodio / Logística | SOP-13 |
| 17 | `En uso` | `Retenido` | Tiempo máximo fuera excedido, sin movimiento reportado | Sistema (automático, vía Reglas operativas + Alertas) | SOP-15, Reglas operativas |
| 18 | `Retenido` | `En tránsito` | Recolección ejecutada | Logística inversa | SOP-14 |
| 19 | `Retenido` | `Bloqueado` | Disputa reportada (custodia o condición) | Comercial / Custodio | SOP-17 |
| 20 | `Retenido` | `Perdido` | Excede umbral adicional sin recuperación ni disputa | Sistema (automático) / Gobernanza | SOP-16 |
| 21 | `Bloqueado` | `En tránsito` / `Retenido` | Disputa resuelta a favor de continuar el ciclo | Gobernanza | SOP-17 |
| 22 | `Bloqueado` | `Perdido` | Disputa resuelta sin recuperación posible | Gobernanza | SOP-17 → SOP-16 |
| 23 | `Perdido` | `En tránsito` | Activo localizado y recuperado | Recuperación (Solutions) | SOP-13 |
| 24 | `Perdido` | `Baja` | Agotado plazo/costo de recuperación | Gobernanza | SOP-18 |
| 25 | `Baja` | `Scrap` | Se determina valor de scrap | Compra de scrap (Solutions) | SOP-19 |
| 26 | `Baja` | `Dispuesto` | Sin valor de scrap, disposición directa | Disposición final (Solutions) | SOP-20 |
| 27 | `Scrap` | `Dispuesto` | Remanente sin valor tras recuperación de valor | Disposición final (Solutions) | SOP-20 |

## 3. Diagrama simplificado

```
                    ┌─────────────┐
                    │ Registrado  │
                    └──────┬──────┘
                           ↓
                    ┌─────────────┐
                    │ Identificado│
                    └──────┬──────┘
                           ↓
              ┌────────────────────────┐
     ┌────────│     En inspección      │◄───────────┐
     ↓        └───────────┬────────────┘            │
  ┌──────┐                ↓                          │
  │ Sucio│──────►┌───────────────┐                   │
  └──────┘       │  En reparación │──────► Baja       │
                  └───────────────┘  (no reparable)   │
                          ↓                            │
                    ┌───────────┐                      │
                    │ Liberado  │                      │
                    └─────┬─────┘                      │
                          ↓                             │
                    ┌───────────┐                       │
                    │Disponible │                       │
                    └─────┬─────┘                       │
                          ↓                              │
                    ┌───────────┐                        │
                    │ Asignado  │                        │
                    └─────┬─────┘                        │
                          ↓                                │
                    ┌────────────┐   (transbordo, self-loop)
                    │ En tránsito│◄──────────────┐          │
                    └─────┬──────┘               │          │
                          ↓                       │          │
                    ┌───────────┐                 │          │
                    │  En uso   │                 │          │
                    └─────┬─────┘                 │          │
              inicia retorno │  excede tiempo      │          │
                          ↓        ↓                │          │
                    (En tránsito) ┌───────────┐    │          │
                                  │ Retenido  │────┘ (recolectado)
                                  └─────┬─────┘  ────────────────┘ (llega a planta)
                        disputa │       │ excede umbral
                                ↓       ↓
                        ┌───────────┐ ┌─────────┐
                        │ Bloqueado │ │ Perdido │
                        └─────┬─────┘ └────┬────┘
                     resuelto │    localizado│  agotado plazo
                              ↓             ↓         ↓
                   (En tránsito/Retenido) (En tránsito) Baja
                                                          ↓
                                              ┌───────────┴──────────┐
                                              ↓                      ↓
                                         ┌────────┐            ┌───────────┐
                                         │ Scrap  │───────────►│ Dispuesto │
                                         └────────┘            └───────────┘
```

## 4. Invariantes (reglas que el sistema debe garantizar)

1. **Ningún activo salta de `Disponible` a `En uso` directamente.** Debe pasar por `Asignado` y `En tránsito` — de lo contrario no hay custodio registrado (Capa 3) y se pierde la cadena de responsabilidad.
2. **Ningún activo entra a `Disponible` sin pasar por `Liberado`.** Evita que un activo recién lavado o reparado, pero sin validación final, se asigne por error.
3. **`Bloqueado` y `Perdido` siempre generan una incidencia asociada (SOP-15).** No son solo un valor de campo — deben quedar acompañados de un ticket de incidencia consultable.
4. **Los tres estados terminales (`Baja`, `Scrap`, `Dispuesto`) no tienen transición de regreso al ciclo operativo.** Una vez en `Baja`, el único camino posible es hacia `Scrap` o `Dispuesto`.
5. **Toda transición queda registrada en trazabilidad** ([Systems #7](./catalogo-systems.md#7-trazabilidad)) con marca de tiempo, actor (persona o "sistema"), estado anterior y estado nuevo — sin excepción, incluidas las transiciones automáticas.
6. **Las transiciones marcadas "Sistema (automático)" dependen de que existan Reglas operativas** ([Systems #11](./catalogo-systems.md#11-reglas-operativas)) configuradas para ese tipo de activo/cliente; sin regla configurada, esa transición debe quedar deshabilitada, no ejecutarse con un valor por defecto arbitrario `[VALIDAR: comportamiento cuando falta la regla — ¿no transiciona, o usa un umbral global?]`.

---

## Próximo paso sugerido

Con el ciclo y sus estados formalizados, el paso 5 de la FASE I es redactar
los **SOP del ciclo del activo** (SOP-01 a SOP-20) siguiendo el formato fijo
de CLAUDE.md (Resumen ejecutivo, Objetivo, Alcance, Roles, Diagrama de flujo,
SOP paso a paso, Checklist, KPI, Riesgos y controles) — usando este documento
y el [ciclo de vida](./ciclo-de-vida-del-activo.md) como insumo directo para
cada SOP individual. Por tratarse de un espacio de trabajo separado de
`chacontainer/`, estos SOP deben numerarse con un prefijo distinto (p. ej.
`SOP-ACTIVO-01`) para no colisionar con los SOP-01 a SOP-13 del ciclo
comercial que ya existen en `chacontainer/docs/sop/`.

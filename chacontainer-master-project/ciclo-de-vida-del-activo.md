# Ciclo de vida del activo — flujo maestro

**Fase:** I · Fundación — paso 3 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

Este documento formaliza el diagrama de la [ETAPA 2](./README.md#etapa-2--sop-del-ciclo-completo-del-activo)
como el flujo único que atraviesa todo activo retornable, anotando en cada
punto **qué servicio de [Solutions](./catalogo-solutions.md) interviene**,
**qué capacidad de [Systems](./catalogo-systems.md) lo sostiene**, y **qué
capa de datos se actualiza** ([Activo / Ubicación / Custodia / Estado /
Gobierno](./README.md#etapa-3--modelo-de-packaging-systems)). Es el punto de
unión entre los dos catálogos anteriores: nada de lo que documentan tiene
sentido operativo fuera de este flujo.

Los SOP referenciados (`SOP-01` a `SOP-20`) son los que la [ETAPA 2](./README.md#etapa-2--sop-del-ciclo-completo-del-activo)
define como pendientes de redactar — este documento es el mapa que los
ordena, no el reemplazo de cada SOP individual.

---

## 1. El flujo principal, anotado

```
Alta ──────────────────────────── SOP-01 · Venta (Solutions) → Registro (Systems) → Capa 1
  ↓
Identificación ──────────────────  SOP-02 · Identificación / QR / RFID (Systems) → Capa 1 + Capa 4 (estado: identificado)
  ↓
Inspección ──────────────────────  SOP-03 · Inspección (Solutions) → Capa 4 (estado: condición evaluada)
  ↓
Limpieza ─────────────────────────  SOP-05 · Lavado (Solutions) → Capa 4 (estado: sucio → limpio)
  ↓
Reparación / reacondicionamiento ── SOP-06 · Reparación / Reacondicionamiento (Solutions) → Capa 4 (estado: dañado → apto)
  ↓
Reetiquetado / reconfiguración ──── SOP-07 · Reetiquetado / Modificación / Dunnage (Solutions) → Capa 1 (variante/config)
  ↓
Liberación ───────────────────────  SOP-08 · Reglas operativas (Systems, criterio de aceptación) → Capa 4 (estado: liberado)
  ↓
Inventario disponible ────────────  SOP-09 · Inventario digital (Systems) → Capa 4 (estado: disponible)
  ↓
Asignación ────────────────────────  SOP-10 · Renta / Venta (Solutions) + Gestión de custodios (Systems) → Capa 3
  ↓
Salida ───────────────────────────  SOP-11 · Movimiento (Systems) → Capa 2
  ↓
Uso ──────────────────────────────  (fuera de custodia directa de CHACONTAINER) → Trazabilidad (Systems) sigue activa vía QR/RFID
  ↓
Movimiento ───────────────────────  SOP-11 · Movimiento (Systems) → Capa 2 (actualización continua)
  ↓
Trazabilidad ─────────────────────  (capacidad continua, no un paso puntual — ver §3)
  ↓
Retorno ──────────────────────────  SOP-13 · Recuperación / Logística inversa (Solutions + Systems) → Capa 2 + Capa 3
  ↓
Reinspección ─────────────────────  SOP-03/04 · Inspección / Clasificación (Solutions) → Capa 4
  ↓
Mantenimiento ────────────────────  SOP-05/06 · Lavado / Reparación (Solutions) → Capa 4
  ↓
        ¿Apto para reingresar?
        ↓ sí                                    ↓ no
Reincorporación                          Baja / Scrap / Recuperación de valor / Disposición final
(vuelve a "Inventario disponible")       SOP-18/19/20 · Compra de scrap, Recuperación de valor,
                                          Disposición final (Solutions) → cierre en Capa 1
```

## 2. Tabla de trazabilidad cruzada por etapa

| Etapa del ciclo | SOP asociado | Servicio(s) Solutions | Capacidad(es) Systems | Capa de datos | Estado resultante (Capa 4) |
|---|---|---|---|---|---|
| Alta | SOP-01 | Venta | Registro | Capa 1 | — (activo nace en el sistema) |
| Identificación | SOP-02 | — | Identificación, QR, RFID | Capa 1 | Identificado |
| Inspección | SOP-03 | Inspección | Trazabilidad (registra el evento) | Capa 4 | Condición evaluada |
| Clasificación (si aplica, lotes) | SOP-04 | Clasificación | Inventario digital | Capa 1 + Capa 4 | `En inspección` (categorizado)¹ |
| Limpieza | SOP-05 | Lavado | Trazabilidad | Capa 4 | Sucio → limpio |
| Reparación / reacondicionamiento | SOP-06 | Reparación, Reacondicionamiento | Gestión de incidencias (si viene de una) | Capa 4 | Dañado → apto |
| Reetiquetado / reconfiguración | SOP-07 | Reetiquetado, Modificación, Dunnage | Registro (variante) | Capa 1 | Reconfigurado |
| Liberación | SOP-08 | Inspección (validación final) | Reglas operativas | Capa 4 | Liberado |
| Inventario disponible | SOP-09 | Inventarios físicos (auditoría) | Inventario digital | Capa 4 | Disponible |
| Asignación | SOP-10 | Renta, Venta | Gestión de custodios, Reglas operativas | Capa 3 | Asignado |
| Salida | SOP-11 | — | Trazabilidad, Movimiento | Capa 2 | En tránsito |
| Uso | — | — | Trazabilidad (pasiva, vía escaneo esporádico) | Capa 2 + Capa 3 | En uso |
| Movimiento | SOP-11 | — | Trazabilidad, Movimiento | Capa 2 | En tránsito / en uso |
| Retorno | SOP-13, SOP-14 | Recuperación | Logística inversa, Gestión de custodios | Capa 2 + Capa 3 | Retenido → en retorno |
| Reinspección | SOP-03, SOP-04 | Inspección, Clasificación | Gestión de incidencias (si hay hallazgo) | Capa 4 | Condición reevaluada |
| Mantenimiento | SOP-05, SOP-06 | Lavado, Reparación | Trazabilidad | Capa 4 | En mantenimiento |
| Reincorporación | SOP-08, SOP-09 | — | Inventario digital | Capa 4 | Disponible (vuelve al ciclo) |
| Baja | SOP-18 | Compra de scrap | Registro (cierre) | Capa 1 | Baja |
| Scrap | SOP-19 | Compra de scrap, Recuperación de valor | Registro (cierre) | Capa 1 | Scrap |
| Disposición final | SOP-20 | Disposición final | Registro (cierre) | Capa 1 | Dispuesto (fin del ciclo) |

¹ La columna "Estado resultante" de esta tabla es descriptiva, no la lista
canónica — esa vive en [estados-del-activo.md](./estados-del-activo.md), que
formaliza 16 estados a partir de esta tabla y de la ETAPA 3. En particular,
**Clasificación no genera un estado propio**: categoriza un lote y cada
activo individual queda en `En inspección` con su categoría asignada,
para luego tomar la transición que le corresponda (`Sucio`, `En reparación`,
`Liberado` o, si el lote lo determina no recuperable, directo a `Baja`) — ver
[estados-del-activo.md §1](./estados-del-activo.md#1-catálogo-de-estados).

## 3. Lo que no es un paso puntual: trazabilidad y gobernanza

Dos elementos del diagrama de la ETAPA 2 no son etapas discretas como las demás — son capacidades que corren **en paralelo a todo el ciclo**, no en una casilla específica:

- **Trazabilidad** registra un evento en cada punto de contacto (escaneo QR/RFID, cambio de estado, transferencia de custodia) desde Alta hasta Disposición final. No "ocurre" entre Movimiento y Retorno como sugiere la posición del diagrama original — ocurre en los dieciséis puntos de esta tabla.
- **Gobernanza** ([Systems #15](./catalogo-systems.md#15-gobernanza)) no interviene en ninguna casilla: consume la tabla completa para decidir, en la bifurcación final, si un activo se reincorpora o se da de baja, y para ajustar las reglas operativas ([Systems #11](./catalogo-systems.md#11-reglas-operativas)) que rigen el resto del ciclo.

## 4. Ramas de excepción (fuera del flujo feliz)

El ciclo maestro asume que cada activo avanza en orden. En operación real hay
tres formas de salir del flujo antes de la bifurcación final, cubiertas por
SOP propios:

| Excepción | SOP | Dónde puede ocurrir | Qué la resuelve |
|---|---|---|---|
| Activo perdido | SOP-16 | En cualquier punto entre Salida y Retorno | Recuperación (Solutions) + Alertas/Logística inversa (Systems) |
| Activo bloqueado | SOP-17 | Asignación, Movimiento, Retorno (disputa de custodia o condición) | Gestión de incidencias (Systems) + Gobernanza |
| Incidencia (daño, retraso, faltante) | SOP-15 | Cualquier etapa | Gestión de incidencias (Systems), deriva a Inspección/Reparación/Baja según el caso |

Estas tres filas son la razón por la que el ciclo maestro no puede
digitalizarse como una máquina de estados lineal simple: en cada etapa hace
falta la salida "excepción" hacia SOP-15/16/17, además de la salida normal
hacia la siguiente etapa.

## 5. Punto de entrada del cliente ≠ inicio del ciclo

La [ETAPA 8](./README.md#etapa-8--modelo-comercial-solutions--systems) ya
establece que un cliente casi nunca entra por "Alta" — entra pidiendo un
servicio puntual (lavado, reparación, inventario físico) sobre activos que ya
llevan tiempo circulando, muchas veces sin haber pasado por Identificación.
Aplicado a este flujo, eso significa:

- Un cliente que pide **lavado** entra directo en la etapa **Limpieza**, con
  el activo probablemente sin haber pasado por Alta/Identificación formal en
  el sistema de CHACONTAINER — lo cual es precisamente lo que expone el
  vacío de datos y abre la puerta a vender Identificación.
- Un cliente que pide **inventario físico** entra en **Inventario
  disponible** desde el lado contrario: primero se descubre lo que hay, y
  recién después se puede reconstruir si esos activos pasaron por Alta e
  Identificación alguna vez.

Este documento asume el ciclo completo desde Alta porque es el caso ideal
(activo nuevo, gestionado desde el día uno). El [catálogo de Solutions](./catalogo-solutions.md#lectura-cruzada-de-dónde-entra-cada-servicio-a-systems)
documenta, servicio por servicio, en qué etapa entra un cliente que no partió
de cero.

---

## Próximo paso sugerido

Con el flujo maestro anotado, el paso 4 de la FASE I es **estados del
activo**: tomar la columna "Estado resultante (Capa 4)" de la tabla del §2 y
formalizarla como una máquina de estados explícita (estados válidos,
transiciones permitidas, quién puede disparar cada transición), que es la
base de datos real que va a vivir en el Módulo 1 de [CHACONTAINER OS](./catalogo-systems.md#16-chacontainer-os).

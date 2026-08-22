# QR — el escaneo como disparador del ciclo

**Fase:** III · Producto — paso 12 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

El sistema QR **ya existe y funciona** en `chacontainer/`: payload firmado
con HMAC, endpoint `POST /api/v1/assets/scan`, tabla `qr_scans`, generador
en Python (ver [architecture.md](../chacontainer/docs/architecture.md#qr-code-system)).
Este documento no rediseña nada de eso. Especifica la única pieza que falta
para que el QR sirva al ciclo de vida: **qué significa cada escaneo**.

---

## 1. El hueco: `action` es hoy un campo libre

El scan flow actual acepta un campo `action` y la tabla lo guarda así:

```sql
action VARCHAR(20) NOT NULL DEFAULT 'inventory'
```

Sin lista cerrada y sin consecuencia: el escaneo registra dónde y cuándo se
vio el activo, pero no qué se le hizo. Eso convierte al QR en una
herramienta de localización — no de operación.

En el modelo de este proyecto, **cada paso del ciclo es un escaneo con
intención**. El operador que termina de lavar un activo no está
"inventariando": está declarando una transición de `Sucio` a `En inspección`
([estados-del-activo.md, transición #7](./estados-del-activo.md#2-tabla-de-transiciones)).
El QR es el teclado con el que la operación escribe en la máquina de
estados.

## 2. Catálogo de acciones de escaneo

Cada acción corresponde a una transición de
[estados-del-activo.md §2](./estados-del-activo.md#2-tabla-de-transiciones)
y al SOP que la gobierna. Lista cerrada propuesta para el MVP
(las transiciones automáticas y las de decisión de gobernanza no aparecen
aquí: no las dispara un escaneo):

| `action` | Transición | SOP | Quién escanea |
|---|---|---|---|
| `alta` | — → `Registrado` | [01](./sop/SOP-ACTIVO-01-alta-de-activos.md) | Operador de alta |
| `identificar` | `Registrado` → `Identificado` | [02](./sop/SOP-ACTIVO-02-identificacion-qr-rfid.md) | Operador de identificación |
| `inspeccion_inicio` | → `En inspección` | [03](./sop/SOP-ACTIVO-03-inspeccion.md) | Inspector |
| `inspeccion_sucio` | `En inspección` → `Sucio` | [03](./sop/SOP-ACTIVO-03-inspeccion.md) | Inspector |
| `inspeccion_dano` | `En inspección` → `En reparación` | [03](./sop/SOP-ACTIVO-03-inspeccion.md) → [06](./sop/SOP-ACTIVO-06-reparacion.md) | Inspector |
| `inspeccion_apto` | `En inspección` → `Liberado` | [08](./sop/SOP-ACTIVO-08-liberacion.md) | Inspector / autoridad de liberación |
| `lavado_fin` | `Sucio` → `En inspección` | [05](./sop/SOP-ACTIVO-05-lavado.md) | Operador de lavado |
| `reparacion_fin` | `En reparación` → `En inspección` | [06](./sop/SOP-ACTIVO-06-reparacion.md) | Técnico |
| `ingreso_inventario` | `Liberado` → `Disponible` | [09](./sop/SOP-ACTIVO-09-inventario.md) | Responsable de inventario |
| `asignar` | `Disponible` → `Asignado` | [10](./sop/SOP-ACTIVO-10-asignacion.md) | Comercial / Operaciones |
| `salida` | `Asignado` → `En tránsito` | [11](./sop/SOP-ACTIVO-11-movimiento.md) | Logística |
| `entrega` | `En tránsito` → `En uso` | [11](./sop/SOP-ACTIVO-11-movimiento.md) | Logística / custodio |
| `transbordo` | `En tránsito` → `En tránsito` | [11](./sop/SOP-ACTIVO-11-movimiento.md) | Logística |
| `transferir_custodia` | (sin cambio de estado; cambia Capa 3) | [12](./sop/SOP-ACTIVO-12-transferencia-de-custodia.md) | Custodio saliente/entrante |
| `retorno_inicio` | `En uso` → `En tránsito` | [13](./sop/SOP-ACTIVO-13-retorno.md) | Custodio / Logística |
| `recoleccion` | `Retenido` → `En tránsito` | [14](./sop/SOP-ACTIVO-14-logistica-inversa.md) | Logística inversa |
| `recepcion_planta` | `En tránsito` → `En inspección` | [03](./sop/SOP-ACTIVO-03-inspeccion.md) | Recepción |
| `inventario` | (sin transición — solo registra avistamiento) | [09](./sop/SOP-ACTIVO-09-inventario.md) | Cualquiera |

`inventario` se conserva como está hoy (valor por defecto, sin efecto sobre
el estado): es el escaneo de "vi este activo aquí", que sigue siendo útil
para conciliación física sin obligar al operador a declarar una intención.

## 3. Validación de transición en el escaneo

El endpoint debe rechazar un escaneo cuya acción no corresponda a una
transición válida desde el estado actual del activo. Ejemplo: `salida` sobre
un activo en `Sucio` es un error de operación —significa que alguien está
por despachar un activo sin lavar— y el sistema debe impedirlo, no
registrarlo.

Comportamiento propuesto:

1. Resolver el activo (ya lo hace el flujo actual).
2. Leer su `estado_detalle` vigente en `asset_state_detail` ([modelo-de-datos.md](./modelo-de-datos.md#asset_state_detail-extiende-capa-4)).
3. Verificar que `(estado_actual, action)` exista en la tabla de transiciones.
4. Si no existe: responder error explicando qué transiciones sí son válidas desde ese estado, y **registrar el intento** en `qr_scans` con un marcador de rechazo — un operador intentando una transición inválida es información de gobernanza, no ruido. **Decidido: sí, generar alerta al superarse un umbral de rechazos** — no como caso especial, sino como una [regla operativa](./reglas-operativas.md#1-qué-es-una-regla) más (tipo nuevo: "rechazos acumulados", scope por operador/turno), tratada exactamente igual que las demás en el [evaluador de alertas](./alertas.md#3-el-evaluador). El número exacto de rechazos que dispara la alerta es un valor de esa regla, no una decisión de diseño aparte — se fija con el resto de los umbrales operativos ([pendientes-de-validacion.md, Nivel 2 §2.2](./pendientes-de-validacion.md#22-umbrales-de-tiempo-y-económicos)).
5. Si existe: aplicar la transición, escribir en `asset_events` (estado detallado y agregado) y actualizar `asset_state_detail`.

Esto convierte las [invariantes 1 y 2 de estados-del-activo.md](./estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)
—"ningún activo salta de `Disponible` a `En uso`", "ninguno entra a
`Disponible` sin pasar por `Liberado`"— en algo que el sistema **garantiza**,
no algo que el SOP pide de buena fe.

## 4. Lo que no cambia

- **El payload y su firma HMAC**: sin cambios. La etiqueta física ya impresa sigue siendo válida.
- **El generador Python**: sin cambios.
- **`qr_scans`**: se conserva; solo se acota `action` a la lista cerrada de §2.
- **RFID**: fuera del MVP ([mvp.md §2](./mvp.md#2-qué-entra-y-qué-no)). Cuando entre, es otro método de lectura sobre el mismo catálogo de acciones — la lista de §2 no cambia, cambia quién la dispara (ver [catalogo-systems.md #5](./catalogo-systems.md#5-rfid)).

## 5. Implicación para el operador de campo

Un catálogo de 18 acciones no cabe en la cabeza de nadie. La app móvil debe
mostrar, tras escanear, **solo las acciones válidas desde el estado actual
del activo** — que en la práctica son entre una y tres. El operador no elige
de una lista de 18: ve "este activo está `Sucio`" y dos botones: *terminé de
lavarlo* o *reportar daño*.

Esa restricción no es cosmética: es lo que hace que el paso 3 de §3 casi
nunca se dispare, porque la interfaz no ofrece la transición inválida en
primer lugar. La validación del servidor queda como red de seguridad para
integraciones y para el caso de estado desactualizado en el dispositivo.

---

## Próximo paso

[Dashboard](./dashboard.md) (paso 13): qué se ve con todo este dato una vez
capturado.

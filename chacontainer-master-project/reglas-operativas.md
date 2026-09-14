# Reglas operativas

**Fase:** II · Sistema — paso 9 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

[estados-del-activo.md, invariante 6](./estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)
deja pendiente una pregunta: las transiciones automáticas del ciclo (p. ej.
`En uso`→`Retenido` cuando se excede el tiempo máximo fuera) dependen de que
exista una regla configurada — pero ¿qué es exactamente una regla, quién la
define y qué pasa si no existe? Este documento especifica la tabla
`operational_rules` que [modelo-de-datos.md §3](./modelo-de-datos.md#operational-rules-capa-5--gobierno--no-existe-hoy)
dejó esbozada.

---

## 1. Qué es una regla

Una regla operativa es la versión configurable de un parámetro de **Capa 5
· Gobierno** ([ETAPA 3](./README.md#etapa-3--modelo-de-packaging-systems)):
un umbral que, al cruzarse, dispara una acción del sistema sin intervención
humana. Los cuatro tipos originales de la Capa 5, más un quinto agregado
para cerrar [qr.md §3](./qr.md#3-validación-de-transición-en-el-escaneo)
(rechazos de escaneo no son un caso especial — son otra regla más):

| Tipo de regla | Parámetro | Ejemplo de acción disparada |
|---|---|---|
| Tiempo máximo fuera | Días desde `Salida` sin retorno | `En uso` → `Retenido` (transición #17 de [estados-del-activo.md](./estados-del-activo.md#2-tabla-de-transiciones)) |
| Punto de retorno | Evento o fecha que marca "debería regresar" (fin de proyecto, vencimiento de renta) | Genera una alerta anticipada antes de llegar al umbral de tiempo máximo |
| Nivel mínimo de inventario | Cantidad disponible por tipo de activo / planta | Alerta a Inventario digital cuando `Disponible` cae bajo el mínimo |
| Condición permitida | Criterio de aceptación por tipo de activo/cliente | Determina si `En inspección` deriva a `Liberado`, `Sucio` o `En reparación` |
| Rechazos acumulados | Número de escaneos con transición inválida por operador/turno | Alerta a Gobernanza — puede indicar entrenamiento insuficiente o intento de saltarse un paso del SOP |

## 2. Esquema de una regla

Extiende `operational_rules` de [modelo-de-datos.md](./modelo-de-datos.md#operational-rules-capa-5--gobierno--no-existe-hoy):

| Columna | Tipo | Nota |
|---|---|---|
| `id` | UUID | PK |
| `tenant_id` | UUID | multi-tenant, igual que el resto del esquema real |
| `scope_type` | VARCHAR(20) | `global`, `cliente`, `tipo_activo`, `planta` — a qué nivel aplica |
| `scope_ref_id` | UUID (nullable) | NULL si `scope_type = global`; si no, FK al cliente/tipo/planta |
| `tipo_regla` | VARCHAR(30) | uno de los 4 de §1 |
| `parametro` | VARCHAR(50) | p. ej. `tiempo_maximo_fuera_dias`, `nivel_minimo_unidades` |
| `valor` | NUMERIC o JSONB | el umbral; JSONB si el criterio es compuesto (p. ej. condición permitida con varios campos) |
| `accion` | VARCHAR(20) | `alerta`, `transicion_automatica` |
| `prioridad` | INTEGER | para resolver conflictos, ver §3 |
| `activa` | BOOLEAN | permite desactivar sin borrar historial |
| `creada_por` / `creada_at` | UUID / TIMESTAMPTZ | auditoría — toda regla la crea o ajusta alguien de Gobernanza ([modelo-de-gobernanza.md](./modelo-de-gobernanza.md)) |

## 3. Prioridad y conflictos

Puede existir más de una regla aplicable al mismo activo (una global de
`tipo_activo` y una específica de `cliente`). Regla de resolución:

1. **La regla más específica gana.** Orden de especificidad:
   `cliente` > `planta` > `tipo_activo` > `global`.
2. Si dos reglas tienen el mismo nivel de especificidad, gana la de mayor
   `prioridad` (entero más alto).
3. Si after (1) y (2) sigue habiendo empate, el sistema no debe decidir
   solo: genera una alerta de configuración a Gobernanza en vez de aplicar
   una transición automática — un empate de reglas es un error de datos,
   no un caso operativo normal `[VALIDAR: este comportamiento por defecto
   con el equipo técnico]`.

## 4. Qué pasa cuando falta la regla (resuelve el `[VALIDAR]` de estados-del-activo.md)

`estados-del-activo.md` dejaba abierto si una transición automática debía
"no transicionar" o "usar un umbral global" cuando falta la regla
específica. Con el modelo de scope de §2, la respuesta queda resuelta por
diseño: **siempre existe al menos la regla `global`** (obligatoria, no
opcional) para cada `tipo_regla`; las reglas de `cliente`/`planta`/`tipo_activo`
son overrides opcionales sobre esa base. Si ni siquiera la regla `global`
está configurada para un `tipo_regla`, esa transición automática queda
deshabilitada y el caso se resuelve manualmente en el
[Nivel 2 (Operativo)](./modelo-de-gobernanza.md#1-los-tres-niveles-de-gobernanza)
de gobernanza — nunca aplica un valor arbitrario por defecto.

## 5. Ejemplo concreto (con valores `[VALIDAR]`)

```
Regla global:
  tipo_regla = tiempo_maximo_fuera
  scope = global
  parametro = tiempo_maximo_fuera_dias
  valor = 45          [VALIDAR — valor de referencia, no confirmado]
  accion = transicion_automatica → Retenido
  prioridad = 0

Regla específica (override):
  tipo_regla = tiempo_maximo_fuera
  scope = cliente:ACME
  parametro = tiempo_maximo_fuera_dias
  valor = 90           [VALIDAR — ACME tiene contrato de renta a largo plazo]
  accion = transicion_automatica → Retenido
  prioridad = 0
```

Un activo de ACME usa 90 días; cualquier otro cliente usa el default global
de 45. Ninguna de las dos cifras está confirmada — son placeholders para
mostrar el mecanismo.

## 6. Relación con el resto del proyecto

- Esta tabla es la que le da contenido real a **Reglas operativas** ([catalogo-systems.md #11](./catalogo-systems.md#11-reglas-operativas)) y a la Capa 5 de la [ETAPA 3](./README.md#etapa-3--modelo-de-packaging-systems).
- Las transiciones automáticas de [estados-del-activo.md §2, filas 10, 13, 17, 20](./estados-del-activo.md#2-tabla-de-transiciones) (marcadas "Sistema (automático)") son, en la práctica, la ejecución de una fila de esta tabla.
- El **Nivel 1 (Automático)** de [modelo-de-gobernanza.md](./modelo-de-gobernanza.md#1-los-tres-niveles-de-gobernanza) es exactamente "una regla dispara su acción sin intervención humana"; el ajuste de esas reglas (crear, modificar, priorizar) es una decisión de **Nivel 3 (Estratégico)**.
- Cambiar un umbral (p. ej. subir el tiempo máximo fuera de un cliente moroso) es la forma concreta en que Gobernanza "ajusta las reglas operativas" mencionada en [modelo-de-gobernanza.md §3](./modelo-de-gobernanza.md#3-ritual-de-gobernanza-propuesto).

---

## Próximo paso sugerido

Paso 10 (último de FASE II): **arquitectura de CHACONTAINER OS** —
reconciliar los 12 módulos del MVP ([ETAPA 5](./README.md#etapa-5--chacontainer-os))
y las tablas nuevas de este bloque (`asset_state_detail`, `asset_custody`,
`operational_rules`, `incidents`, `alerts`) contra la arquitectura técnica
real ya documentada en `chacontainer/docs/architecture.md`, y dejar un
roadmap de qué construir primero.

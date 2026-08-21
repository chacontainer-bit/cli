# Escala · Hubs y partners

**Fase:** V · Escala — pasos 21 y 22 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

> **Advertencia de horizonte.** Todo lo de la FASE V depende de que las fases
> I-IV hayan ocurrido. Escribir hoy un plan detallado de expansión sería
> inventar certezas. Lo que sí se puede fijar ahora, y es lo que hace este
> documento, son tres cosas: **qué condición dispara** cada iniciativa, **qué
> decisión de hoy la haría imposible**, y **cuál es la restricción real** —
> que casi nunca es la que parece.

---

## 1. Hubs (paso 21)

Un hub es un nodo intermedio —propio o rentado— donde los activos se
reciben, lavan, inspeccionan y redistribuyen sin volver a la planta
central.

### Qué lo dispara

No "cuando crezcamos". La condición es económica y medible con los datos que
el sistema ya captura a partir de la FASE III:

> Cuando, en una ruta, el **costo y el tiempo de transporte de retorno**
> superan al de procesamiento — es decir, cuando el
> [tiempo de ciclo](./kpi.md#1-los-10-kpi-especificados) está dominado por
> el traslado y no por el lavado o la reparación.

Ese es exactamente el número que produce el KPI de tiempo de ciclo
segmentado por ruta ([dashboard.md §4](./dashboard.md#4-vista-gobernanza)).
Antes de tener ese dato, cualquier decisión de hub es intuición.

### La restricción real

No es el capital ni el inmueble: es el **volumen mínimo de rotación** que
justifica tener personal y equipo de lavado ociosos entre picos. Un hub con
poca rotación es una bodega cara. La pregunta correcta no es "¿cuántos
activos pasarían por aquí?" sino "¿cuántos ciclos por mes?" —
[rotación](./kpi.md), no inventario.

### Qué no hay que romper hoy

El esquema ya trata la ubicación como `plants` + `plant_zones` multi-tenant.
Un hub es una planta más; no requiere un concepto nuevo. **Lo que sí hay que
evitar** es escribir lógica que asuma una sola planta de origen — por ejemplo,
reglas de retorno que apunten a una planta fija en vez de a "la planta
asignada a esa ruta" ([reglas-operativas.md](./reglas-operativas.md), campo
`scope`).

## 2. Partners (paso 22)

Un partner ejecuta la función física —lavado, reparación, almacenamiento—
en su instalación, bajo la marca y el proceso de CHACONTAINER.

### Qué lo dispara

Necesitar presencia en una zona antes de poder justificar un hub propio.
Partner y hub resuelven el mismo problema con distinto perfil de riesgo:
el hub cuesta capital y da control; el partner da alcance inmediato y cuesta
control.

### La restricción real: es un riesgo de datos antes que de calidad

Aquí está lo importante de este documento, y es contraintuitivo.

La preocupación natural con un partner es la calidad del servicio físico:
que lave mal, que repare mal. Eso es real, es visible y se corrige. **El
riesgo mayor es invisible: que ejecute el servicio bien pero no capture el
dato.**

Un partner que lava 300 activos impecablemente pero no registra el escaneo
de entrada y salida ([qr.md](./qr.md#2-catálogo-de-acciones-de-escaneo))
produce 300 activos limpios y un agujero de tres días en la trazabilidad. El
cliente ve activos limpios y CHACONTAINER pierde la capacidad de responder
dónde estuvieron — que es justamente lo que se cobra en Systems. La calidad
física se degrada de forma visible; la integridad del dato se degrada en
silencio.

Consecuencias para el contrato de partner:

1. **La captura del dato es obligación contractual, al mismo nivel que la calidad del servicio.** No una buena práctica sugerida.
2. **Se mide.** Escaneos esperados vs. registrados por período, igual que la métrica de adopción del piloto ([piloto.md §6](./piloto.md#6-riesgos)). Un partner con cobertura de escaneo baja está incumpliendo aunque sus activos salgan impecables.
3. **El partner opera dentro de CHACONTAINER OS**, no en su propio sistema con conciliación posterior. Una conciliación mensual entre dos sistemas reconstruye conteos, no líneas de tiempo — y la trazabilidad es una línea de tiempo.
4. **Hay un umbral de cobertura por debajo del cual no se renueva** `[VALIDAR umbral]`.

### Qué no hay que romper hoy

El modelo de custodia ya contempla `proveedor` y `operador` como tipos de
custodio ([modelo-de-datos.md](./modelo-de-datos.md#asset_custody-capa-3--custodia--no-existe-hoy)),
así que un partner cabe sin cambio conceptual. **Lo que sí hay que evitar**
es asumir en el código o en las reglas que "planta propia" y "bajo control
de CHACONTAINER" son lo mismo: con partners dejan de serlo, y esa distinción
importa para saber a quién se le exige un retorno.

## 3. La decisión que conviene no tomar todavía

Si los partners se manejan como **subcontratistas** (CHACONTAINER factura al
cliente, el partner factura a CHACONTAINER) o como **franquicia/red**
(el partner factura al cliente bajo estándar CHACONTAINER) cambia el modelo
de ingresos, la relación con el cliente y la exigibilidad del dato.

No hace falta decidirlo ahora, pero sí conviene tener presente que el
segundo modelo es mucho más difícil de sostener con las obligaciones de
datos de §2: un partner que factura directo al cliente tiene menos incentivo
a alimentar un sistema que no es suyo `[VALIDAR con el fundador cuando haya
un primer candidato real]`.

---

Continúa en [Escala · RFID](./escala-rfid.md) (paso 23).

# Escala geográfica · Multi-planta, nacional y México–Estados Unidos

**Fase:** V · Escala — pasos 25, 26 y 27 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

Los tres últimos pasos son el mismo movimiento en tres escalas. En cada uno,
la restricción que frena no es la que se espera: en multi-planta no es el
software, en nacional no es la logística, y en el cruce fronterizo no es el
transporte.

---

## 1. Multi-planta (paso 25)

### La restricción no es técnica

El esquema ya es multi-tenant y multi-planta: `plant_id` aparece en
`assets`, `asset_events`, `qr_scans`, `shipments`, y las reglas operativas
ya contemplan `scope_type = planta`
([reglas-operativas.md §2](./reglas-operativas.md#2-esquema-de-una-regla)).
Agregar plantas no requiere rediseño.

Lo que frena es **la consistencia de ejecución de los SOP**. Si la planta A
registra `inspeccion_apto` con un criterio y la planta B con otro, los dos
producen el mismo dato con significados distintos. El resultado no es un
error visible: es un KPI que promedia dos cosas incomparables y que reporta
una mejora o un deterioro que no ocurrió.

Esto es peor que no tener el dato, porque nadie sospecha de un número.

### Qué instalar antes de la segunda planta

1. **Criterios de aceptación escritos y específicos por tipo de activo**, no por costumbre local — hoy en `[VALIDAR]` en [SOP-ACTIVO-03](./sop/SOP-ACTIVO-03-inspeccion.md).
2. **Comparación entre plantas como control de calidad del dato, no solo de desempeño.** Si la planta A rechaza el 3% en inspección y la B el 22%, lo primero a descartar no es que la B reciba peor material: es que apliquen criterios distintos.
3. **Reglas con `scope = planta` solo donde haya una razón operativa real.** Cada override es una diferencia legítima; también es una diferencia que hace menos comparables los números.

## 2. Operación nacional (paso 26)

### La restricción es que la gobernanza no escala agregando gente

El [ritual de gobernanza](./modelo-de-gobernanza.md#3-ritual-de-gobernanza-propuesto)
funciona para una planta: alguien revisa las excepciones abiertas cada día.
Con veinte plantas, esa revisión diaria no existe — nadie mira veinte
bandejas, y contratar veinte revisores convierte un negocio de gobernanza en
un negocio de nómina.

La única salida es que el **Nivel 1 (automático)** absorba el volumen y que
los humanos vean solo excepciones reales. Es decir: **la escala nacional
depende de que el diseño anti-ruido de
[alertas.md §1](./alertas.md#1-el-riesgo-real-el-ruido-no-la-cobertura)
funcione de verdad.**

Si con una planta el responsable ya ignora la bandeja porque está saturada,
con veinte el sistema colapsa — no gradualmente, sino en el momento en que
nadie confía en las alertas. El indicador temprano a vigilar es simple y
conviene medirlo desde el piloto:

> **% de alertas atendidas vs. generadas.** Si baja de forma sostenida, no
> es un problema de disciplina del equipo: es que el sistema genera más
> señal de la que merece atención, y hay que subir umbrales antes de crecer
> `[VALIDAR piso aceptable]`.

### Lo demás

Hubs y partners ([escala-hubs-y-partners.md](./escala-hubs-y-partners.md))
son el instrumento de cobertura física de esta etapa; no se repiten aquí.

## 3. Operación México–Estados Unidos (paso 27)

Aquí sí aparecen problemas nuevos, no solo de escala. Cuatro.

### 3.1 La aduana es un custodio que no se comporta como tal

Un activo detenido en aduana no está `En tránsito` (no avanza), ni
`Retenido` (nadie incumplió un plazo acordado), ni `Bloqueado` (no hay
disputa entre partes). Está detenido por un tercero que no es custodio en
ningún sentido del [modelo de custodia](./modelo-de-datos.md#asset_custody-capa-3--custodia--no-existe-hoy)
y con quien no se puede negociar un retorno.

El esquema real ya reconoce esto del lado logístico: `shipments.status`
incluye `at_customs`. Falta el equivalente del lado del activo. Antes de
operar cruces conviene decidir si eso es un estado nuevo —el primero que se
agregaría a los [16](./estados-del-activo.md#1-catálogo-de-estados)— o un
atributo sobre `En tránsito` que suspende el conteo de tiempo
`[VALIDAR: la segunda opción parece preferible, porque el tiempo en aduana
no debe penalizar el KPI de cumplimiento de retorno del custodio]`.

Esa última consideración importa: si el tiempo en aduana se le carga al
cliente que iba a recibir el activo, el indicador miente y la conversación
comercial se vuelve injusta.

### 3.2 El régimen aduanal restringe el ciclo de vida

Este es el punto menos evidente y el de mayor consecuencia.

Un empaque retornable que cruza bajo un régimen de importación temporal
puede tener la **obligación legal de regresar** al país de origen dentro de
un plazo. Si ese activo se daña del otro lado, la decisión de
[`Baja`](./sop/SOP-ACTIVO-18-baja.md) o
[`Scrap`](./sop/SOP-ACTIVO-19-scrap.md) deja de ser una decisión operativa y
pasa a tener implicaciones aduanales y fiscales.

Dicho de otro modo: **la máquina de estados asume que la decisión de dar de
baja un activo depende solo de su condición y su costo. Cruzando la frontera,
depende también de dónde está y bajo qué régimen entró.** Las transiciones
24, 25 y 26 de [estados-del-activo.md](./estados-del-activo.md#2-tabla-de-transiciones)
necesitarían una validación adicional para activos bajo régimen temporal
`[VALIDAR con asesor aduanal antes de operar cruces]`.

No hay que resolverlo hoy. Sí hay que evitar escribir la lógica de baja de
forma que asuma que la condición del activo es el único criterio.

### 3.3 Normativa ambiental divergente

[SOP-ACTIVO-20](./sop/SOP-ACTIVO-20-disposicion-final.md) asume normativa y
proveedores homologados mexicanos —ya marcados `[VALIDAR]`—. Del lado
estadounidense, la disposición final se rige por otro marco y otros
proveedores. Ese SOP necesitará una variante por país, no un párrafo
adicional.

### 3.4 Moneda

El esquema ya soporta moneda por cliente (`clients.currency`, hoy `MXN` por
defecto), así que no hay deuda estructural. Lo que no está resuelto es cómo
se expresa el **costo por ciclo** ([kpi.md](./kpi.md)) cuando un mismo ciclo
incurre costos en dos monedas. Es una decisión de reporte, no de esquema
`[VALIDAR: moneda de reporte única vs. por segmento]`.

---

## Cierre de FASE V y del proyecto maestro

Con este documento se completan los 7 pasos de la FASE V · Escala:
[hubs y partners](./escala-hubs-y-partners.md) (21-22),
[RFID](./escala-rfid.md) (23),
[integraciones ERP/MES](./escala-integraciones.md) (24), y
multi-planta, nacional y México–EE.UU. (25-27) en este documento.

**Las cinco fases y los 27 pasos de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución)
quedan documentados.** Lo que sigue no es más documentación: es validar los
`[VALIDAR]` con operación real y ejecutar el
[roadmap técnico](./arquitectura-os.md#3-roadmap-de-implementación-propuesto)
y el [piloto](./piloto.md).

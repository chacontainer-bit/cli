# Escala · RFID

**Fase:** V · Escala — paso 23 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

RFID quedó explícitamente fuera del [MVP](./mvp.md#2-qué-entra-y-qué-no) y
como Fase F —opcional— del
[roadmap técnico](./arquitectura-os.md#3-roadmap-de-implementación-propuesto).
Este documento fija cuándo deja de ser opcional y qué hay que entender antes
de instalarlo.

---

## 1. La limitación que nadie menciona: RFID no expresa intención

Este es el punto que decide toda la arquitectura, y suele pasarse por alto
porque RFID se presenta como "QR pero mejor".

[qr.md](./qr.md#2-catálogo-de-acciones-de-escaneo) definió 18 acciones de
escaneo, cada una equivalente a una transición de estado. La mayoría
requieren que **una persona declare una intención**: `inspeccion_apto` no es
"el activo pasó por aquí", es "yo, inspector, declaro que este activo cumple
el criterio de aceptación". `lavado_fin` es una afirmación con
responsabilidad detrás.

Un portal RFID no puede declarar nada. Sabe que un tag cruzó una antena en
un instante. Eso lo hace excelente para un subconjunto de acciones y
completamente inútil para el resto:

| Tipo de acción | Ejemplos | ¿RFID puede? |
|---|---|---|
| **De paso** — la ubicación es el hecho completo | `salida`, `entrega`, `transbordo`, `recepcion_planta`, `inventario` | ✅ Sí, y mejor que QR |
| **De declaración** — alguien afirma un juicio | `inspeccion_apto`, `inspeccion_dano`, `lavado_fin`, `reparacion_fin` | ❌ No. Requieren una persona |
| **De decisión** — implican una autorización | `asignar`, `transferir_custodia` | ❌ No |

**Conclusión: RFID complementa a QR, no lo sustituye.** Un parque con RFID
sigue necesitando QR (o su equivalente en la app) para las acciones de
declaración. Cualquier plan que prometa "quitamos el escaneo manual" está
prometiendo eliminar justo los registros que sostienen la calidad del dato.

El catálogo de 18 acciones no cambia con RFID. Cambia quién dispara cuáles.

## 2. Qué lo dispara

Dos condiciones, y basta con una:

1. **Volumen en un punto de control.** Cuando en un portal (entrada/salida
   de planta, andén) el escaneo uno por uno se vuelve el cuello de botella
   del flujo. Señal medible: tiempo de despacho dominado por el registro y
   no por la carga.
2. **Frecuencia de paso por activo.** Cuando un activo cruza puntos de
   control tantas veces por ciclo que el costo acumulado de mano de obra de
   escaneo supera al del tag más la infraestructura.

La segunda condición se puede calcular con datos que el sistema ya tiene
desde la FASE III: número de escaneos de tipo "de paso" por activo por
período, que sale directo de `qr_scans`.

## 3. La economía, en estructura

RFID cuesta más por activo (tag) **y** exige inversión fija (lectores,
antenas, instalación). QR cuesta casi nada por activo y cero infraestructura.
El punto de equilibrio depende de:

```
Costo RFID  = (tags × costo_tag) + infraestructura_de_lectura
Costo QR    = (etiquetas × costo_etiqueta) + (escaneos × tiempo × costo_hora)
```

De donde salen dos consecuencias prácticas:

- **RFID se justifica por concentración, no por tamaño.** Un parque de 5,000 activos repartidos en veinte puntos sin control fijo no lo justifica; uno de 800 que cruza dos portales varias veces al mes, sí.
- **La infraestructura es costo hundido por ubicación.** Eso favorece desplegarlo primero donde ya existe un [hub](./escala-hubs-y-partners.md) o una planta con volumen estable, no en instalaciones de clientes que pueden cambiar.

Cifras: en [`lista-precios-costos.md`](../chacontainer/docs/lista-precios-costos.md)
cuando existan, no aquí ([pricing.md](./pricing.md)).

## 4. Impacto en el modelo de datos

Menor, y ya está previsto en
[modelo-de-datos.md §4](./modelo-de-datos.md#4-identificación-qr-ya-cubierto-rfid-no):
o se agrega `assets.rfid_tag` con una tabla `rfid_reads` paralela a
`qr_scans`, o se generaliza `qr_scans` a `identifier_reads` con un campo
`method`.

Recomendación, a la luz de §1: **generalizar**. Si RFID y QR alimentan
tablas distintas, toda consulta de trazabilidad tendrá que unir dos fuentes
para siempre, y las acciones "de paso" quedarán partidas por método de
lectura en vez de por significado. Un solo flujo de lecturas con un campo
que indique cómo se leyó mantiene la línea de tiempo íntegra
`[VALIDAR con el equipo técnico: costo de migrar `qr_scans` vs. mantener
dos tablas]`.

## 5. Qué no hay que romper hoy

- **No hardcodear "QR" en la lógica de negocio.** Las validaciones de transición ([qr.md §3](./qr.md#3-validación-de-transición-en-el-escaneo)) deben operar sobre "una lectura de identificador con una acción", no sobre "un escaneo QR". Si esa abstracción se respeta desde el MVP, RFID entra sin tocar la máquina de estados.
- **No asumir que toda lectura tiene un usuario humano asociado.** Hoy `qr_scans.scanned_by` referencia a `users`. Una lectura de portal RFID no tiene persona detrás — necesitará un actor de tipo sistema, igual que las transiciones automáticas de [reglas-operativas.md](./reglas-operativas.md).

El segundo punto es el más fácil de olvidar y el más caro de corregir
después, porque `scanned_by` es `NOT NULL` en el esquema actual.

---

Continúa en [Escala · Integraciones ERP/MES](./escala-integraciones.md) (paso 24).

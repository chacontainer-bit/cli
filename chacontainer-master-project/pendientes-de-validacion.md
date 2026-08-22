# Pendientes de validación — lista maestra

El proyecto arrancó con **254 marcas `[VALIDAR]`** repartidas en 44
documentos (24 documentos estratégicos + 20 SOP-ACTIVO). Ninguna es un
placeholder vacío: cada una es un punto donde el proyecto necesita un dato
real, una decisión, o una confirmación que este documento no tiene
autoridad para inventar.

**Estado actual: 8 de las 17 decisiones del Nivel 1 ya están resueltas**
(compromiso sobre KPI, participación en ahorro, estructura del fee por
activo, modelo de partner, extensión y acreditación del piloto pagado,
quién es Gobernanza por cliente, y periodicidad del comité). Quedan 9
decisiones abiertas de Nivel 1 — de las cuales 3 son solo la cifra exacta
de una decisión de dirección ya tomada (vigencia del piso, % de
acreditación, umbral de tamaño para delegar Gobernanza). El resto de este
documento refleja el estado vigente.

Esta lista no es para leer de corrido. Es un **checklist de trabajo**,
organizado en 4 niveles según quién puede resolverlo y qué tan bloqueante es.

---

## Cómo usar esta lista

| Nivel | Quién decide | Bloquea |
|---|---|---|
| 1 · Estratégicas | Solo tú (fundador) | El modelo comercial y de riesgo — sin esto, cualquier propuesta a un cliente descansa en supuestos no confirmados |
| 2 · Operativas | Tú + responsable de planta/operación | Que los SOP puedan ejecutarse en campo tal como están escritos |
| 3 · Diseño técnico | Tú + quien construya sobre `chacontainer/` | Cómo se implementa, no si se implementa |
| 4 · Brechas de alcance | Nadie las "valida" — hay que redactarlas | Cobertura incompleta del proyecto, ya identificada |

Los niveles 1 y 4 son cortos y conviene resolverlos primero. El nivel 2 es
el más largo (201 de las 254 marcas) pero la mayoría son variaciones del
mismo tipo de pregunta — se resuelven en bloque, no una por una.

---

## Nivel 1 · Decisiones estratégicas (9 abiertas, 8 resueltas)

Ordenadas por impacto en lo que se le puede prometer a un cliente hoy.

### Resueltas

- ✅ **Compromiso sobre KPI en la Oferta 4**: compromiso blando con piso de responsabilidad — proyección basada en el piloto, revisada trimestralmente en el comité de gobernanza, sin penalización contractual — [oferta-systems.md, Oferta 4](./oferta-systems.md#oferta-4--gobernanza-del-sistema-nivel-5).
- ✅ **Participación en el ahorro medido**: no se ofrece por ahora; queda reservada como componente adicional posible en cuentas grandes y maduras — [pricing.md §1](./pricing.md#1-la-decisión-que-ordena-todo-la-métrica-de-valor).
- ✅ **Estructura del fee "por activo administrado"**: piso fijo por la vigencia del contrato, recalculado al renovar — [pricing.md §2](./pricing.md#2-el-incentivo-perverso-que-hay-que-resolver). Queda abierto solo el número de meses/años del piso (ver más abajo).
- ✅ **Subcontratista vs. franquicia/red para partners**: subcontratista — CHACONTAINER mantiene la facturación con el cliente y con ella la palanca para exigir cobertura de escaneo — [escala-hubs-y-partners.md §3](./escala-hubs-y-partners.md#3-modelo-de-partner-decidido).
- ✅ **Extensión del piloto pagado si no se cumple el umbral de mejora**: 30 días adicionales sin costo; si tampoco se alcanza, cierre con entrega en vez de una segunda extensión — [piloto-pagado.md §3](./piloto-pagado.md#3-reversión-de-riesgo-qué-ofrecer-y-qué-no). Queda abierto solo el % de acreditación al contrato (ver más abajo).
- ✅ **Acreditación al contrato si se firma rápido tras el piloto**: % fijo del total pagado (no solo del componente de plataforma), si se firma dentro de 30-60 días del cierre — [piloto-pagado.md §3](./piloto-pagado.md#3-reversión-de-riesgo-qué-ofrecer-y-qué-no). Queda abierto el porcentaje exacto (ver más abajo).
- ✅ **Quién es "Gobernanza" por cliente**: el fundador por defecto en toda cuenta; delegable a un responsable de cuenta en cuentas menores una vez probado el criterio con el fundador al frente — [modelo-de-gobernanza.md §1](./modelo-de-gobernanza.md#1-los-tres-niveles-de-gobernanza). Queda abierto el umbral de tamaño de cuenta que dispara la delegación (ver más abajo).
- ✅ **Periodicidad del comité de gobernanza**: mensual — [modelo-de-gobernanza.md §3](./modelo-de-gobernanza.md#3-ritual-de-gobernanza-propuesto).

### Modelo comercial y de riesgo

1. **Vigencia mínima del piso del fee** (en meses/años) que hace rentable el setup — [pricing.md §2](./pricing.md#2-el-incentivo-perverso-que-hay-que-resolver).
2. **% de acreditación del piloto al contrato** — [piloto-pagado.md §3](./piloto-pagado.md#3-reversión-de-riesgo-qué-ofrecer-y-qué-no). Calcularlo contra el margen real del piloto para que convertir rápido no salga más caro que no convertir.
3. **Indicador de mezcla Solutions/Systems por cuenta**, para medir si Systems canibaliza ingreso transaccional o lo sustituye por algo mejor — [pricing.md §3](./pricing.md#3-la-pregunta-incómoda-systems-canibaliza-a-solutions).
4. **Descuento máximo autorizado sin aprobación** — ya estaba pendiente en el SOP-07 comercial del SaaS; se hereda aquí — [pricing.md §5](./pricing.md#5-lo-que-falta-decidir-antes-de-cotizar).

### Piloto pagado

5. **Fijar el umbral de mejora con el cliente antes de empezar el piloto**, no al cierre — [piloto.md §4](./piloto.md#4-criterios-de-éxito). Es una práctica a instalar, no un número a decidir una sola vez.
6. **Encuadre de la línea base como diagnóstico, no auditoría de responsabilidades**, acordado explícitamente con cada cliente — [piloto.md §6](./piloto.md#6-riesgos).

### Gobernanza

7. **Umbral de tamaño de cuenta que dispara la delegación de Gobernanza** a un responsable de cuenta — [modelo-de-gobernanza.md §1](./modelo-de-gobernanza.md#1-los-tres-niveles-de-gobernanza). Deliberadamente no fijado aún: calibrar con datos de las primeras cuentas en Nivel 5, no por intuición.
8. **Umbral de recurrencia y monto económico que dispara escalada de Nivel 2 a Nivel 3** — [modelo-de-gobernanza.md §4](./modelo-de-gobernanza.md#4-qué-activa-una-escalada-de-nivel-2-a-nivel-3).

### Escala (decisiones que aún no urgen, pero conviene tener presentes)

9. **Umbral de cobertura de escaneo bajo el cual no se renueva un partner** — [escala-hubs-y-partners.md §2](./escala-hubs-y-partners.md#2-partners-paso-22).
10. **% de alertas atendidas vs. generadas aceptable antes de escalar a operación nacional** — [escala-geografica.md §2](./escala-geografica.md#2-operación-nacional-paso-26).
11. **Consultar con asesor aduanal antes de operar cruces México–EE.UU.** — el régimen de importación temporal puede restringir la decisión de baja/scrap de un activo — [escala-geografica.md §3.2](./escala-geografica.md#32-el-régimen-aduanal-restringe-el-ciclo-de-vida).
12. **Moneda de reporte cuando un ciclo cruza frontera** (única vs. por segmento) — [escala-geografica.md §3.4](./escala-geografica.md#34-moneda).
13. **Definir el criterio de disparo del hub/partner una vez haya datos de ruta reales** — [escala-hubs-y-partners.md §1](./escala-hubs-y-partners.md#1-hubs-paso-21). No es una decisión aislada: depende de que el sistema ya esté midiendo tiempo de ciclo por ruta.

---

## Nivel 2 · Parámetros operativos (201 marcas en los 20 SOP-ACTIVO)

No se listan una por una — la mayoría son la misma pregunta repetida sobre
un rol, un umbral o un criterio técnico distinto. Se resuelven mejor en una
sola sesión con el responsable operativo, recorriendo el
[índice de SOP](./sop/README.md), que uno por uno.

### 2.1 Roles exactos (la categoría más frecuente)

Casi todos los SOP marcan `[VALIDAR]` sobre el **nombre exacto del puesto**
responsable de cada paso ("Responsable de planta", "Inspector",
"Clasificador / Operador de clasificación"). La pregunta de fondo es
siempre la misma: **¿es un rol dedicado, o una responsabilidad adicional de
alguien que ya existe en el organigrama?**

Acción recomendada: una sola sesión de 30-60 min recorriendo la sección
"Roles" de los 20 SOP y confirmando, para cada uno, el título real —no
crear puestos nuevos solo para que el documento quede prolijo.

### 2.2 Umbrales de tiempo y económicos

Aparecen en casi todos los SOP de la rama de excepción y cierre:

- Días para pasar de `En uso` a `Retenido` (tiempo máximo fuera).
- Días adicionales en `Retenido` antes de declarar `Perdido`.
- Umbral de valor del activo bajo el cual un rol operativo puede resolver una disputa sin escalar a gobernanza (SOP-ACTIVO-17).
- Criterio objetivo de "no reparable" — costo de reparación vs. costo de reposición (SOP-ACTIVO-06, SOP-ACTIVO-18).

Estos umbrales son, en gran parte, **los mismos parámetros de
`operational_rules`** ([reglas-operativas.md](./reglas-operativas.md)) —
conviene validarlos una sola vez ahí y que los SOP los referencien, no
fijarlos por separado en cada documento.

### 2.3 Criterios técnicos y proveedores

- Químicos y concentración exactos por tipo de activo en lavado (SOP-ACTIVO-05).
- Material y proveedor de la etiqueta QR resistente a lavado/intemperie (catalogo-systems.md, SOP-ACTIVO-02).
- Proveedor(es) de disposición final homologados (SOP-ACTIVO-20, catalogo-solutions.md).
- Criterio de "condición permitida" por cliente — si ya existe una tabla o hay que construirla (SOP-ACTIVO-03, SOP-ACTIVO-10).

### 2.4 Autoridad y aprobación

- Quién tiene autoridad de primera instancia para aprobar una baja individual, y cuándo escala (SOP-ACTIVO-18).
- Quién aprueba la compra de scrap al precio cotizado (SOP-ACTIVO-19).
- Si liberación requiere un rol distinto al que hizo la inspección inicial — separación de funciones (SOP-ACTIVO-08).

---

## Nivel 3 · Decisiones de diseño técnico (6)

Para quien construya sobre `chacontainer/` — no cambian el modelo de
negocio, sí cómo se implementa.

1. **`asset_custody.custodian_ref_id`: FK polimórfica vs. tabla `custodians` unificada** — [modelo-de-datos.md §3](./modelo-de-datos.md#asset_custody-capa-3--custodia--no-existe-hoy).
2. **RFID: generalizar `qr_scans` a `identifier_reads`, o mantener tabla paralela `rfid_reads`** — [escala-rfid.md §4](./escala-rfid.md#4-impacto-en-el-modelo-de-datos). El documento recomienda generalizar.
3. **Dónde se guardan hoy las fotos de evidencia** (Airtable, S3, en ningún lado) — [arquitectura-os.md, módulo 12](./arquitectura-os.md#2-los-12-módulos-del-mvp-mapeados).
4. **Comportamiento cuando falta una regla operativa para una transición automática**: ya resuelto en el diseño ([reglas-operativas.md §4](./reglas-operativas.md#4-qué-pasa-cuando-falta-la-regla-resuelve-el-validar-de-estados-del-activomd)) — falta implementarlo, no decidirlo de nuevo.
5. **Frecuencia del evaluador de alertas** (el documento asume cada hora) — [alertas.md §3](./alertas.md#3-el-evaluador).
6. **Si conviene generar alerta cuando un operador acumula N escaneos rechazados** (transición inválida intentada) — [qr.md §3](./qr.md#3-validación-de-transición-en-el-escaneo).

---

## Nivel 4 · Brechas de alcance conocidas (3)

Esto no son preguntas para validar — son piezas que el proyecto ya
identificó como faltantes y que alguien tiene que redactar o modelar:

1. **No existe SOP-ACTIVO propio para Modificación ni para Dunnage.** Ambos servicios de Solutions (#6 y #7 del [catálogo](./catalogo-solutions.md)) hoy solo se mencionan como derivación desde [SOP-ACTIVO-07 (Reetiquetado)](./sop/SOP-ACTIVO-07-reetiquetado.md), sin procedimiento paso a paso propio. Si alguno de los dos empieza a venderse con volumen, necesita su propio SOP.
2. **No existe estado de Capa 4 para "en aduana".** El esquema real ya tiene `shipments.status = 'at_customs'` del lado del envío; falta el equivalente del lado del activo en los [16 estados](./estados-del-activo.md#1-catálogo-de-estados) — necesario antes de operar cruces de frontera ([escala-geografica.md §3.1](./escala-geografica.md#31-la-aduana-es-un-custodio-que-no-se-comporta-como-tal)).
3. **SOP-ACTIVO-20 (Disposición final) asume normativa y proveedores mexicanos.** Necesitará una variante por país antes de operar en Estados Unidos ([escala-geografica.md §3.3](./escala-geografica.md#33-normativa-ambiental-divergente)).

---

## Qué hacer con esta lista

No es necesario cerrar los 254 puntos antes de avanzar. El orden de ataque
que se deriva del resto del proyecto:

1. **Nivel 1, ahora** — son decisiones de modelo de negocio; todo lo demás las asume.
2. **Nivel 2, en la sesión de preparación del piloto** ([piloto.md, semana −2 a 0](./piloto.md#3-cronograma)) — ahí de todos modos hay que sentarse con el equipo operativo.
3. **Nivel 3, cuando arranque la Fase A del roadmap técnico** ([arquitectura-os.md §3](./arquitectura-os.md#3-roadmap-de-implementación-propuesto)) — no antes, porque decidir el modelo de datos sin tener aún el primer piloto es prematuro.
4. **Nivel 4, cuando el volumen del servicio o la geografía lo exija** — no son bloqueantes hoy.

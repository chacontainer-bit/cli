# SOP-ACTIVO-04: Clasificación de condición

**Estado: Borrador v0.1 — pendiente de validación.**

Este SOP documenta el servicio **Clasificación** (#10) del [catálogo
Solutions](../catalogo-solutions.md#10-clasificación), aplicado a la etapa
**Clasificación** del [ciclo de vida del activo](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado)
(fila "Clasificación (si aplica, lotes)" del [§2 de ese
documento](../ciclo-de-vida-del-activo.md#2-tabla-de-trazabilidad-cruzada-por-etapa)).
Se construyó a partir del cruce entre el catálogo Solutions, el ciclo de
vida y la [máquina de estados](../estados-del-activo.md) — **no a partir de
una entrevista real con el responsable operativo**. No debe declararse
estándar hasta validar en operación los puntos marcados `[VALIDAR]`.

**Nota de consistencia entre documentos (resuelta):** la tabla del [§2 de
ciclo-de-vida-del-activo.md](../ciclo-de-vida-del-activo.md#2-tabla-de-trazabilidad-cruzada-por-etapa)
ya deja explícito que Clasificación no genera un estado propio de Capa 4:
el resultado de la clasificación es un **metadato/atributo del lote**
(categoría + conteo), y cada activo individual conserva o adopta su estado
real del [catálogo de 16 estados](../estados-del-activo.md#1-catálogo-de-estados)
(p. ej. `En inspección`) al ser enrutado al proceso siguiente (ver
[§6, paso 8](#6-sop-paso-a-paso)).

---

## 1. Resumen ejecutivo

Clasificación es el proceso que recibe un **lote heterogéneo** de activos —
típicamente de una recuperación masiva, el cierre de la operación de un
cliente, o un inventario físico — y lo separa en categorías (tipo,
condición aparente, propietario) para que cada subgrupo pueda enrutarse al
proceso siguiente que le corresponde. No emite un dictamen individual por
activo: produce un **conteo por categoría** y una decisión de enrutamiento
por subgrupo. Es, según el [catálogo
Solutions](../catalogo-solutions.md#10-clasificación), "el primer conteo
real de un lote que antes solo existía como estimado" — insumo directo para
inventarios físicos y decisiones de recuperación de valor.

## 2. Objetivo

Convertir un lote de activos sin clasificar en subgrupos contables por
categoría (tipo, condición aparente, propietario), documentados con
evidencia y conteo verificable, de forma que cada subgrupo pueda procesarse,
facturarse y enrutarse de manera independiente — sin que la clasificación
misma sustituya el dictamen individual que corresponde a Inspección.

## 3. Alcance

Aplica a todo lote de activos retornables que llega **sin clasificar**,
proveniente de:

- Recuperación masiva de activos detenidos o fuera de sitio.
- Cierre de la operación de un cliente (devolución total de su parque).
- Inventario físico que descubre activos mezclados en tipo, condición o
  propietario.
- Retorno masivo de campo que llega como lote heterogéneo, antes de que
  cada activo entre individualmente a `En inspección`.

**Distinción explícita entre Clasificación e Inspección**, requerida porque
ambas intervienen en el mismo punto del ciclo (retorno) y es fácil
confundirlas:

| | Clasificación (este SOP) | Inspección ([SOP-ACTIVO-03](./SOP-ACTIVO-03-inspeccion.md)) |
|---|---|---|
| **Unidad de trabajo** | Un **lote** de activos. | **Un** activo individual. |
| **Qué produce** | Conteo por categoría (tipo / condición aparente / propietario / destino). | Dictamen individual: apto / requiere lavado / requiere reparación. |
| **Nivel de detalle** | Filtro rápido de volumen, criterio visual grueso. | Checklist detallado (estructural, funcional, normativa) por activo. |
| **Efecto en la máquina de estados** | No dispara por sí sola una transición de Capa 4 en la [máquina de estados](../estados-del-activo.md) (ver nota de consistencia arriba). | Dispara la transición real de estado del activo individual (`Sucio`, `En reparación`, `Liberado`). |
| **Cuándo se usa** | Cuando el volumen y la heterogeneidad de un lote hacen inviable inspeccionar activo por activo desde el primer momento. | Siempre, para cada activo, sin excepción — incluyendo cada activo que salió de un lote ya clasificado. |

Clasificación **antecede** a Inspección cuando el lote es heterogéneo: se
separa primero en categorías gruesas, y después cada subgrupo (o cada
activo dentro de él) pasa por Inspección para el dictamen individual
definitivo. La precisión de la Clasificación se valida, precisamente,
contrastándola contra el resultado de esa inspección posterior (ver
[catálogo Solutions, Clasificación](../catalogo-solutions.md#10-clasificación),
campo "indicador que afecta").

**No incluye:**

- El dictamen individual de aceptación/rechazo de un activo — eso es
  [SOP-ACTIVO-03 (Inspección)](./SOP-ACTIVO-03-inspeccion.md).
- La ejecución de lavado, reparación, compra de scrap o disposición final
  sobre los subgrupos ya clasificados — cada uno de esos procesos tiene su
  propio SOP (SOP-ACTIVO-05, SOP-ACTIVO-06, SOP-18/19/20, pendientes o
  fuera del alcance de este documento).
- La recuperación física del lote desde el sitio del cliente/proveedor —
  cubierta por Recuperación / logística inversa (SOP-13/14, fuera de
  alcance).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Clasificador / Operador de clasificación `[VALIDAR nombre exacto del puesto — puede coincidir con Inspector o ser un rol distinto para lotes grandes]` | Separa físicamente el lote por tipo, condición aparente y propietario; documenta el conteo por categoría. |
| Responsable de planta `[VALIDAR]` | Define el criterio de clasificación (qué cuenta como "reutilizable" vs. "reparable" vs. "scrap" a nivel de filtro grueso), resuelve casos límite y aprueba el reporte de clasificación. |
| Cliente `[VALIDAR cuándo aplica — depende de si el lote es propiedad del cliente o de CHACONTAINER]` | Recibe y, si aplica, firma el acta de conformidad sobre el conteo resultante. |
| Sistema (CHACONTAINER OS) | Registra el conteo por categoría como metadato del lote y, al enrutar cada subgrupo, asegura que cada activo individual reciba su estado real correspondiente (ver nota de consistencia y [§6, paso 8](#6-sop-paso-a-paso)) `[VALIDAR implementación exacta]`. |

## 5. Diagrama de flujo

```
[Lote heterogéneo ingresa: recuperación masiva / cierre de
 operación de cliente / inventario físico / retorno masivo]
                ↓
[Registrar ingreso del lote: origen y volumen estimado]
                ↓
[Separar por tipo de activo (contenedor, IBC, tarima, tambor)]
                ↓
[Separar por condición aparente
 (reutilizable / reparable / scrap)]
                ↓
[Separar por propietario, si el lote mezcla más de un cliente/tenant]
                ↓
[Contar y documentar cada categoría + evidencia fotográfica]
                ↓
[Generar reporte de clasificación]
                ↓
       ¿El lote pertenece a un cliente?
       ↓ sí                              ↓ no
[Obtener acta de conformidad          [Continuar sin acta
 con el cliente]                       (lote propio de CHACONTAINER)]
                ↓
        Enrutar cada subgrupo al proceso siguiente:
        ├─ Reutilizable ──────► Inspección individual
        │                        (SOP-ACTIVO-03)
        ├─ Reparable ──────────► Inspección individual
        │                        (SOP-ACTIVO-03, confirma o
        │                         deriva a reparación)
        └─ Scrap ──────────────► Compra de scrap /
                                  Recuperación de valor
                                  (SOP-18/19, pendientes)
```

## 6. SOP paso a paso

1. Recibir el lote sin clasificar y registrar su origen (recuperación
   masiva, cierre de operación de cliente, inventario físico, retorno
   masivo) y volumen estimado antes de iniciar la separación.
2. Separar físicamente el lote por tipo de activo (contenedor, IBC, tarima,
   tambor).
3. Separar cada subgrupo por condición aparente — reutilizable / reparable
   / scrap — aplicando un criterio visual rápido, **no** el checklist
   detallado de Inspección `[VALIDAR criterio exacto de clasificación por
   condición aparente — es un filtro de volumen, no un dictamen individual]`.
4. Si el lote mezcla activos de más de un propietario/cliente, separar
   también por propietario antes de continuar — relevante en contextos
   multi-tenant (ver [migrations/*.sql](../../chacontainer/migrations/) del
   SaaS, donde `tenant_id` es la dimensión de separación equivalente en el
   modelo de datos).
5. Contar y documentar el resultado por categoría, con evidencia
   fotográfica del lote completo y de cada subgrupo.
6. Generar el reporte de clasificación (conteo por categoría) y, si el lote
   pertenece a un cliente, obtener acta de conformidad firmada sobre el
   conteo resultante.
7. Registrar en el sistema el resultado agregado de la clasificación como
   metadato del lote (categoría + conteo), sin asignar un estado de Capa 4
   inexistente a cada activo individual (ver nota de consistencia al inicio
   de este documento) `[VALIDAR modelo de datos exacto para representar
   "categoría de clasificación" vs. estado individual del activo]`.
8. Enrutar cada subgrupo al proceso siguiente:
   - **Reutilizable** → [Inspección individual](./SOP-ACTIVO-03-inspeccion.md),
     que emitirá el dictamen definitivo por activo.
   - **Reparable** → Inspección individual, que confirmará el diagnóstico o
     derivará a [Reparación](./SOP-ACTIVO-06-reparacion.md).
   - **Scrap** → Compra de scrap / Recuperación de valor (SOP-18/19,
     pendientes de redactar).
9. Ningún subgrupo queda "clasificado" sin dueño de proceso: cada uno debe
   quedar formalmente enrutado antes de cerrar la clasificación del lote.
10. Actualizar los indicadores de tiempo de clasificación por volumen de
    lote y % del lote recuperable vs. scrap (ver [§8 KPI](#8-kpi)).
11. Cuando el subgrupo "reutilizable"/"reparable" complete su inspección
    individual posterior, contrastar el dictamen resultante contra la
    categoría asignada aquí, para alimentar el indicador de precisión de la
    clasificación `[VALIDAR periodicidad y responsable de esta validación
    cruzada]`.

## 7. Checklist

- [ ] Origen y volumen estimado del lote registrados antes de iniciar la
      separación.
- [ ] Lote separado por tipo de activo.
- [ ] Lote separado por condición aparente (reutilizable / reparable /
      scrap), con el criterio visual documentado, no improvisado.
- [ ] Lote separado por propietario, si mezcla más de un cliente/tenant.
- [ ] Conteo por categoría documentado y contrastado contra el volumen
      estimado inicial.
- [ ] Evidencia fotográfica del lote completo y de cada subgrupo generada.
- [ ] Reporte de clasificación generado.
- [ ] Acta de conformidad firmada por el cliente, cuando el lote es de su
      propiedad.
- [ ] Cada subgrupo enrutado al proceso siguiente (Inspección individual o
      Compra de scrap / Recuperación de valor) — ninguno queda pendiente
      sin dueño de proceso.
- [ ] Resultado agregado registrado en el sistema como metadato del lote,
      sin sobreescribir el estado individual de cada activo.
- [ ] Precisión de la clasificación contrastada contra el resultado de la
      inspección individual posterior `[VALIDAR periodicidad]`.

## 8. KPI

- Tiempo de clasificación por volumen de lote (ver [catálogo Solutions,
  Clasificación](../catalogo-solutions.md#10-clasificación)).
- % del lote recuperable (reutilizable + reparable) vs. scrap.
- **Precisión de la clasificación**, validada contra el dictamen de la
  inspección posterior (§6, paso 11) — mide si el filtro grueso de
  Clasificación coincide con el dictamen individual de Inspección.
- % de lotes de clientes con acta de conformidad firmada sobre el total de
  lotes clasificados que son propiedad de un cliente.
- Volumen clasificado por categoría por periodo — insumo directo para
  [Inventarios físicos](../catalogo-solutions.md#11-inventarios-físicos) y
  para la [ETAPA 4 de gobernanza](../README.md#etapa-4--modelo-de-gobernanza)
  ("¿qué planta tiene exceso/déficit?").

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Confundir la categoría de clasificación con un dictamen definitivo (liberar o rechazar un activo solo con base en clasificación) | La clasificación nunca sustituye la inspección individual (§3); todo subgrupo "reutilizable" o "reparable" pasa igual por [SOP-ACTIVO-03](./SOP-ACTIVO-03-inspeccion.md) antes de cualquier transición real de estado. |
| Categoría de clasificación sin representación en la máquina de estados formal genera datos huérfanos o inconsistentes | Tratar la categoría de clasificación como metadato/atributo del lote, no como estado de Capa 4 (ver nota de consistencia al inicio del documento) `[VALIDAR con el equipo de datos]`. |
| Lote multi-cliente mezclado sin separar por propietario, generando disputas de custodia o facturación | Separación por propietario obligatoria (paso 4) antes de generar el reporte, con acta de conformidad cuando aplica. |
| Criterio de clasificación (reutilizable/reparable/scrap) inconsistente entre operadores | Documentar por escrito el criterio de clasificación — separado del checklist detallado de inspección — y capacitar a más de un clasificador (mismo patrón de control usado en SOP-02 del SaaS para lavado). |
| Reporte de clasificación sin evidencia fotográfica, dificultando resolver disputas con el cliente sobre el conteo | Evidencia fotográfica obligatoria del lote completo y de cada subgrupo (paso 5). |
| Subgrupo clasificado que nunca se enruta al proceso siguiente y queda "flotando" sin dueño | La clasificación no se considera cerrada hasta que los tres subgrupos (reutilizable, reparable, scrap) quedan formalmente enrutados (paso 9); reportar en el indicador de volumen clasificado los subgrupos sin enrutar. |

---

**Documentos relacionados:** [catalogo-solutions.md](../catalogo-solutions.md) ·
[catalogo-systems.md](../catalogo-systems.md) ·
[ciclo-de-vida-del-activo.md](../ciclo-de-vida-del-activo.md) ·
[estados-del-activo.md](../estados-del-activo.md) ·
[SOP-ACTIVO-03 · Inspección](./SOP-ACTIVO-03-inspeccion.md)

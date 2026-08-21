# SOP-ACTIVO-07: Reetiquetado

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre la etapa **Reetiquetado / reconfiguración** del [ciclo de vida del activo](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado),
enfocado en el servicio [Reetiquetado](../catalogo-solutions.md#8-reetiquetado) del catálogo Solutions como
actividad central. Se construyó a partir del cruce entre el catálogo Solutions, el
[catálogo Systems](../catalogo-systems.md) y la [máquina de estados](../estados-del-activo.md) — **no a
partir de una entrevista real con el responsable operativo**. No debe declararse estándar hasta validar en
operación los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Reetiquetado es el proceso que actualiza la etiquetación de un activo retornable (marca, cliente, contenido,
normativa aplicable) cuando esta queda obsoleta o incorrecta — típicamente porque el activo cambia de
propietario, de contenido o de destino tras pasar por [Reparación/reacondicionamiento](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado).
Retira la etiquetación previa, aplica la nueva conforme a especificación y normativa, valida legibilidad y
adherencia, y deja el activo listo para [Liberación](./SOP-ACTIVO-08-liberacion.md). Es también el momento
natural para instalar o reemplazar la identificación QR/RFID del activo (ver
[catálogo Solutions #8](../catalogo-solutions.md#8-reetiquetado)).

## 2. Objetivo

Garantizar que todo activo que requiere actualización de etiquetación queda con la etiqueta correcta,
legible, adherida y conforme a la normativa aplicable, antes de avanzar al checkpoint de Liberación.

## 3. Alcance

Aplica al reetiquetado físico de un activo retornable (contenedor, IBC, tarima, tambor): retiro de la
etiqueta previa, aplicación de la nueva, y validación de legibilidad/adherencia/cumplimiento normativo.

**Fuera del detalle paso a paso de este SOP** (aunque convergen en el mismo punto del ciclo —
"Reetiquetado / reconfiguración" — cuando el cliente aprovecha la intervención):

- **[Modificación](../catalogo-solutions.md#6-modificación)** — adaptación estructural o de accesorios del
  activo a un requerimiento específico. Es un proceso de mayor alcance (diseño, validación de conformidad
  contra especificación técnica) que requiere su propio SOP `[VALIDAR: SOP-ACTIVO pendiente para modificación]`.
- **[Dunnage](../catalogo-solutions.md#7-dunnage)** — diseño e instalación de protección interna a medida
  para el producto del cliente. También un proceso de mayor alcance (diseño, prototipo, prueba de
  transporte) que requiere su propio SOP `[VALIDAR: SOP-ACTIVO pendiente para dunnage]`.

Cuando una intervención combina reetiquetado con modificación y/o dunnage, este SOP cubre únicamente el
componente de reetiquetado; los otros componentes se derivan al proceso correspondiente antes de continuar
(ver paso 3 del [§6 SOP paso a paso](#6-sop-paso-a-paso)).

No incluye la [Liberación](./SOP-ACTIVO-08-liberacion.md) del activo (checkpoint siguiente, SOP separado) ni
la [Reparación/reacondicionamiento](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado) previa
(SOP-ACTIVO-06 `[VALIDAR]`).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Operador de reetiquetado `[VALIDAR nombre exacto del puesto]` | Retira la etiqueta previa y aplica la nueva conforme a especificación. |
| Responsable de cumplimiento normativo `[VALIDAR si es un rol dedicado o el mismo responsable de planta]` | Define y verifica la normativa aplicable por tipo de contenido (materiales peligrosos, alimentos, etc.). |
| Inspector de calidad | Valida legibilidad y adherencia de la etiqueta nueva contra checklist, antes de enviar a Liberación. |
| Comercial / Cuenta `[VALIDAR]` | Origina la especificación de reetiquetado cuando el activo cambia de propietario, contenido o destino. |
| Responsable de planta `[VALIDAR]` | Decide si la intervención requiere derivarse a Modificación y/o Dunnage (paso 3) y resuelve excepciones. |

## 5. Diagrama de flujo

```
[Activo con etiquetación previa, proveniente de
 Reparación/reacondicionamiento o directo de Inspección]
                ↓
[Recibir especificación de nuevo etiquetado
 (marca/cliente, contenido, normativa aplicable)]
                ↓
        ¿La intervención requiere también
        Modificación y/o Dunnage?
        ↓ sí                              ↓ no
[Derivar componente de Modificación/       [Continuar solo con reetiquetado]
 Dunnage al proceso correspondiente
 (fuera de este SOP) y continuar
 aquí solo con el reetiquetado]
                ↓                                  ↓
                └──────────────┬───────────────────┘
                                ↓
                  [Retirar etiquetación previa]
                                ↓
                  [Aplicar nueva etiqueta conforme
                   a especificación y normativa]
                                ↓
                  [Validar legibilidad y adherencia
                   contra checklist]
                                ↓
                        ¿Cumple criterio?
                        ↓ sí              ↓ no
              [Registrar evidencia   [Reprocesar (retirar/reaplicar)
               → enviar a Liberación  o escalar a responsable de
               (SOP-ACTIVO-08)]       cumplimiento normativo]
```

## 6. SOP paso a paso

1. Recibir el activo con etiquetación previa, proveniente de Reparación/reacondicionamiento
   (SOP-ACTIVO-06 `[VALIDAR]`) o directo desde Inspección con instrucción de reetiquetado.
2. Recibir la especificación de nuevo etiquetado (marca/cliente, contenido, normativa aplicable) por parte
   de Comercial u Operaciones.
3. Determinar si la intervención requiere además Modificación estructural o instalación/reemplazo de
   Dunnage (ver [catálogo Solutions #6](../catalogo-solutions.md#6-modificación) y
   [#7](../catalogo-solutions.md#7-dunnage)). Si es así, derivar esos componentes al proceso correspondiente
   `[VALIDAR SOP a definir para modificación/dunnage]` y continuar este SOP únicamente con el reetiquetado —
   no ejecutar trabajo de mayor alcance sin la cotización/plan que le corresponde.
4. Retirar la etiquetación previa del activo.
5. Aplicar la nueva etiqueta conforme a la especificación y normativa aplicable
   `[VALIDAR material de etiqueta, método de fijación, y normativa por tipo de contenido — materiales
   peligrosos, alimentos, etc.]`.
6. Validar legibilidad y adherencia de la etiqueta nueva contra checklist de aceptación.
7. Si la etiqueta cumple el criterio: registrar evidencia fotográfica y el checklist de cumplimiento
   normativo, actualizar el historial de propietario/contenido del activo (Registro, Capa 1 —
   [catálogo Systems #2](../catalogo-systems.md#2-registro)), y evaluar si corresponde instalar o reemplazar
   la identificación QR/RFID en el mismo momento (ver [SOP-ACTIVO-02](./SOP-ACTIVO-02-identificacion-qr-rfid.md)
   `[VALIDAR]`).
8. Si la etiqueta no cumple el criterio, reprocesar (retirar y reaplicar) o escalar al responsable de
   cumplimiento normativo si el rechazo tiene causa normativa, no solo estética.
9. Enviar el activo al checkpoint de [Liberación](./SOP-ACTIVO-08-liberacion.md), siguiente etapa del
   [ciclo de vida](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado).
10. Actualizar el indicador de tiempo de reetiquetado (ver [§8 KPI](#8-kpi)).

> **Nota sobre estado del activo:** la [máquina de estados](../estados-del-activo.md#1-catálogo-de-estados)
> no define un estado propio para "en reetiquetado" — el activo permanece en `En reparación` (si el
> reetiquetado acompaña una reparación) o en `En inspección` (si el reetiquetado es la única intervención)
> mientras se ejecuta este SOP, y solo transiciona a `Liberado` tras pasar el checkpoint de
> [SOP-ACTIVO-08](./SOP-ACTIVO-08-liberacion.md) `[VALIDAR si conviene formalizar un sub-estado o bandera de
> proceso en vez de reutilizar estos dos estados]`.

## 7. Checklist

- [ ] Especificación de nuevo etiquetado (marca/cliente, contenido, normativa) recibida antes de iniciar.
- [ ] Determinado si la intervención requiere derivar Modificación y/o Dunnage a su propio proceso.
- [ ] Etiquetación previa retirada por completo.
- [ ] Nueva etiqueta aplicada conforme a especificación y normativa aplicable.
- [ ] Legibilidad y adherencia validadas contra checklist (no solo "a simple vista").
- [ ] Evidencia fotográfica y checklist de cumplimiento normativo generados.
- [ ] Historial de propietario/contenido actualizado en el registro maestro (Capa 1).
- [ ] Oportunidad de instalar/reemplazar identificación QR/RFID evaluada explícitamente.
- [ ] Activo enviado al checkpoint de Liberación (no queda "reetiquetado" sin avanzar).

## 8. KPI

- Tiempo de reetiquetado por activo o por lote.
- % de rechazos por incumplimiento normativo o por adherencia/legibilidad insuficiente.
- Activos habilitados para nuevo cliente/uso tras reetiquetado, por periodo.
- % de intervenciones de reetiquetado donde se aprovechó para actualizar identificación QR/RFID.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Etiqueta nueva no cumple la normativa aplicable (materiales peligrosos, alimentos) | Checklist de cumplimiento normativo obligatorio (paso 6-7), con responsable de cumplimiento normativo definido. |
| Historial de propietario/contenido no se actualiza en el registro maestro | Actualización del Registro (paso 7) es requisito para avanzar a Liberación, no un paso opcional posterior. |
| Se ejecuta Modificación o Dunnage "de paso" dentro de este SOP, sin la cotización/plan que ese proceso requiere | Derivación explícita al proceso correspondiente en el paso 3, antes de continuar con el reetiquetado. |
| Etiqueta se desprende en campo por adherencia insuficiente | Validación de adherencia obligatoria (paso 6), separada de la aplicación misma, con criterio documentado `[VALIDAR método de prueba de adherencia]`. |
| Activo avanza a Liberación sin evidencia de reetiquetado completa | Checklist (§7) exigido antes de enviar el activo a [SOP-ACTIVO-08](./SOP-ACTIVO-08-liberacion.md). |

# SOP-ACTIVO-19: Scrap

**Estado: Borrador v0.1 — pendiente de validación.**

Este documento aplica el formato fijo de SOP definido en `CLAUDE.md` a la
transición `Baja` → `Scrap` de la máquina de estados del activo. Se
construyó a partir de la [transición 25 de la tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones)
y del servicio [Compra de scrap](../catalogo-solutions.md#13-compra-de-scrap)
del catálogo de Solutions — **no** a partir de una entrevista real con el
responsable comercial o de planta. No debe declararse estándar hasta
validar en operación real los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

`Scrap` es el estado que recibe a un activo ya dado de baja
([SOP-ACTIVO-18](./SOP-ACTIVO-18-baja.md)) cuando CHACONTAINER determina
que conserva valor de material o de componentes recuperables. Este SOP
cubre la evaluación y compra/valoración de ese scrap: quién lo valora, con
qué criterio, y qué evidencia deja. De aquí puede derivar
[Recuperación de valor](../catalogo-solutions.md#15-recuperación-de-valor)
(servicio Solutions #15) sobre el lote antes de que el remanente sin valor
avance a
[Disposición final](../catalogo-solutions.md#14-disposición-final)
([SOP-ACTIVO-20](./SOP-ACTIVO-20-disposicion-final.md)).

## 2. Objetivo

Garantizar que todo activo que pasa de `Baja` a `Scrap` recibe una
valoración consistente y documentada, que la compra queda evidenciada
para cierre contable/registro, y que el remanente no recuperable tiene
una salida clara hacia disposición final — sin depender del criterio
informal de quien esté disponible ese día.

## 3. Alcance

Aplica a todo activo o lote de activos en estado `Baja` que se evalúa para
determinar valor de scrap, conforme a la [transición 25](../estados-del-activo.md#2-tabla-de-transiciones)
(`Baja` → `Scrap`). Incluye la valoración, la compra/adquisición del scrap
por parte de CHACONTAINER, y la derivación hacia recuperación de valor
cuando aplica. No incluye la decisión de dar de baja el activo (cubierta
en SOP-ACTIVO-18, que ya debe estar completa antes de iniciar este SOP) ni
la disposición física del remanente sin valor (cubierta en
SOP-ACTIVO-20). No aplica a activos que todavía están en `En reparación`
o `Perdido` — deben pasar primero por `Baja`.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Evaluador de scrap `[VALIDAR — puede ser el mismo Responsable de planta de SOP-ACTIVO-18 u otro rol]` | Evalúa el lote o activo en `Baja`, determina su valor de scrap según el criterio vigente, y documenta la cotización. |
| Responsable comercial/compras `[VALIDAR]` | Aprueba la compra (adquisición interna del activo como scrap) al precio cotizado y autoriza el registro de la transacción. |
| Responsable de registro `[VALIDAR]` | Ejecuta el cambio de estado a `Scrap` en el sistema y vincula la evidencia de compra al ID del activo. |
| Responsable de recuperación de valor `[VALIDAR]` | Evalúa, sobre el lote ya en `Scrap`, si existen componentes o materiales aprovechables antes de enviar el remanente a disposición final (ver [Solutions #15](../catalogo-solutions.md#15-recuperación-de-valor)). |

## 5. Diagrama de flujo

```
[Activo en estado Baja]
         ↓
[Evaluar valor de scrap del activo/lote]
         ↓
   ¿Tiene valor de scrap?
   ↓ sí                    ↓ no
[Cotizar]          [Derivar directo a
   ↓                Disposición final
[Aprobar compra]    (SOP-ACTIVO-20,
   ↓                 transición 26)]
[Registrar compra/retiro
 (evidencia + peso/unidades)]
   ↓
[Estado → Scrap]
   ↓
[¿Componentes/materiales
 aprovechables en el lote?]
   ↓ sí                    ↓ no
[Recuperación de valor    [Derivar remanente a
 (Solutions #15):          Disposición final
 desarmar, separar,        (SOP-ACTIVO-20,
 vender/reusar]             transición 27)]
   ↓
[Remanente sin valor →
 Disposición final
 (SOP-ACTIVO-20, transición 27)]
```

## 6. SOP paso a paso

1. **Confirmar entrada.** Verificar que el activo o lote está en estado
   `Baja` con decisión ya registrada conforme a
   [SOP-ACTIVO-18](./SOP-ACTIVO-18-baja.md). No iniciar valoración de
   scrap sobre un activo que no haya completado esa decisión.
2. **Evaluar el lote.** El evaluador de scrap inspecciona el activo o lote
   y determina si conserva valor de material (metal, plástico, componentes
   reutilizables) según el criterio de valoración vigente
   `[VALIDAR criterio exacto de valor de scrap: tabla de precios por
   material/peso, cotización caso por caso, o ambos]`.
3. **Cotizar.** Si el lote tiene valor, se genera una cotización de compra
   con el precio ofrecido, referenciando la convención de precios/costos
   existente cuando aplique `[VALIDAR si aplica la misma referencia que
   catalogo-solutions.md usa para el SaaS actual, o si scrap tiene su
   propia tabla]`.
4. **Aprobar la compra.** El Responsable comercial/compras autoriza la
   adquisición al precio cotizado. Si el lote no tiene valor de scrap, se
   documenta esa determinación y el activo se deriva directo a
   [SOP-ACTIVO-20](./SOP-ACTIVO-20-disposicion-final.md)
   ([transición 26](../estados-del-activo.md#2-tabla-de-transiciones), `Baja` → `Dispuesto`).
5. **Ejecutar la compra/retiro.** Se genera la evidencia formal: acta de
   compraventa/retiro y registro de peso o unidades adquiridas (conforme al
   séptimo campo de [Compra de scrap en el catálogo de Solutions](../catalogo-solutions.md#13-compra-de-scrap)).
6. **Registrar la transición de estado.** El estado del activo cambia de
   `Baja` a `Scrap` en el sistema ([transición 25](../estados-del-activo.md#2-tabla-de-transiciones)),
   con marca de tiempo, actor, estado anterior y estado nuevo en
   Trazabilidad ([Systems #7](../catalogo-systems.md#7-trazabilidad)),
   conforme al [invariante #5](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar).
7. **Evaluar recuperación de valor.** Antes de enviar el lote completo a
   disposición final, el Responsable de recuperación de valor determina si
   existen componentes o materiales que puedan desarmarse, separarse y
   revenderse o reusarse internamente (ver
   [Recuperación de valor, Solutions #15](../catalogo-solutions.md#15-recuperación-de-valor)).
   Este paso es opcional por lote — no todo scrap tiene componentes
   recuperables individualmente `[VALIDAR criterio para decidir si vale la
   pena desarmar vs. disponer el lote completo]`.
8. **Derivar el remanente.** Lo que no tiene valor de reventa/reuso —ya sea
   el lote completo o lo que queda tras recuperación de valor— se deriva a
   [SOP-ACTIVO-20 (Disposición final)](./SOP-ACTIVO-20-disposicion-final.md)
   ([transición 27](../estados-del-activo.md#2-tabla-de-transiciones), `Scrap` → `Dispuesto`).

## 7. Checklist

- [ ] Activo/lote confirmado en estado `Baja` con decisión registrada
      (SOP-ACTIVO-18 completo).
- [ ] Evaluación de valor de scrap documentada, con criterio aplicado
      explícito.
- [ ] Cotización generada (si aplica) y aprobada por el Responsable
      comercial/compras.
- [ ] Evidencia de compra/retiro generada (acta de compraventa, registro
      de peso/unidades).
- [ ] Transición a `Scrap` registrada en Trazabilidad con fecha, actor,
      estado anterior y estado nuevo.
- [ ] Evaluación de recuperación de valor realizada y documentada (aunque
      la conclusión sea "no aplica").
- [ ] Remanente sin valor derivado explícitamente a SOP-ACTIVO-20.

## 8. KPI

- Volumen de scrap comprado (peso o unidades) por periodo.
- Margen entre precio de compra del scrap y valor recuperado
  posteriormente (venta de componentes + valor evitado de disposición).
- Tasa de conversión de lotes evaluados a efectivamente comprados.
- % de lotes en `Scrap` que pasan por recuperación de valor antes de
  disposición final (vs. disposición directa del lote completo).
- % de valor recuperado sobre el valor de scrap adquirido (KPI propio de
  [Recuperación de valor](../catalogo-solutions.md#15-recuperación-de-valor)).
- Tiempo entre entrada a `Scrap` y salida a `Dispuesto` — evita
  acumulación de inventario "scrap" sin cierre.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Criterio de valoración de scrap inconsistente entre evaluadores | Documentar por escrito el criterio de valoración (tabla de precios por material/peso o proceso de cotización) `[VALIDAR criterio exacto]`, igual que exige [SOP-02](../../chacontainer/docs/sop/SOP-02-lavado-de-activos-retornables.md) para el criterio de aceptación de lavado. |
| Scrap con valor recuperable enviado directo a disposición final sin evaluar recuperación de valor | Paso 7 obligatorio (aunque la conclusión sea "no aplica") antes de derivar a SOP-ACTIVO-20; registrar la determinación, no solo omitirla. |
| Compra de scrap sin evidencia formal (riesgo de disputa o de auditoría contable) | Acta de compraventa/retiro y registro de peso/unidades obligatorios antes de ejecutar la transición de estado (paso 5–6). |
| Lote acumulado indefinidamente en estado `Scrap` sin avanzar a disposición | KPI de tiempo `Scrap` → `Dispuesto` (sección 8) y alerta si excede el umbral definido ([Systems #12 · Alertas](../catalogo-systems.md#12-alertas)) `[VALIDAR umbral]`. |
| Confusión entre "valor de scrap" (material) y "valor de reventa de componente" (recuperación de valor), llevando a subvalorar el lote | Separar explícitamente los dos pasos en este SOP (compra de scrap en paso 3–6, recuperación de valor en paso 7) con roles y evidencia distintos. |

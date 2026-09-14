# SOP-ACTIVO-20: Disposición final

**Estado: Borrador v0.1 — pendiente de validación.**

Este documento aplica el formato fijo de SOP definido en `CLAUDE.md` al
estado terminal `Dispuesto` de la máquina de estados del activo. Se
construyó a partir de las [transiciones 26 y 27 de la tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones)
y del servicio [Disposición final](../catalogo-solutions.md#14-disposición-final)
del catálogo de Solutions — **no** a partir de una entrevista real con el
responsable de cumplimiento ambiental o de operaciones. No debe declararse
estándar hasta validar en operación real los puntos marcados `[VALIDAR]`,
en particular el proveedor de disposición homologado.

---

## 1. Resumen ejecutivo

`Dispuesto` es el estado terminal absoluto del ciclo de vida del activo:
una vez alcanzado, no existe ninguna transición de regreso
([invariante #4](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)).
Este SOP cubre las dos rutas que llegan a `Dispuesto` — directo desde
`Baja` cuando el activo no tiene valor de scrap
([transición 26](../estados-del-activo.md#2-tabla-de-transiciones)), y
desde `Scrap` cuando queda un remanente sin valor tras recuperación de
valor ([transición 27](../estados-del-activo.md#2-tabla-de-transiciones))
— y define cómo se gestiona la disposición conforme a normativa ambiental,
con qué proveedor, y qué certificado cierra el registro del activo.

## 2. Objetivo

Garantizar que todo activo dado de baja termina su ciclo de vida de forma
responsable y trazable — nunca "desaparece" sin evidencia —, que la
disposición se ejecuta con un proveedor autorizado conforme a normativa
ambiental aplicable, y que el registro del activo queda cerrado de forma
auditable con el certificado correspondiente.

## 3. Alcance

Aplica a todo activo o remanente de material que llega a `Dispuesto` desde
`Baja` (sin valor de scrap) o desde `Scrap` (remanente sin valor tras
recuperación de valor, ver [SOP-ACTIVO-19](./SOP-ACTIVO-19-scrap.md)). No
incluye la decisión de dar de baja el activo (SOP-ACTIVO-18) ni la
valoración/compra de scrap (SOP-ACTIVO-19) — este SOP asume que esas
decisiones ya están tomadas y registradas. No aplica a activos en
cualquier otro estado del ciclo operativo.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Responsable de cumplimiento ambiental `[VALIDAR — puede ser un rol dedicado o una responsabilidad adicional de planta]` | Determina la vía de disposición conforme a normativa ambiental aplicable según el tipo de material/activo, y selecciona el proveedor homologado correspondiente. |
| Proveedor de disposición `[VALIDAR proveedor(es) de disposición homologados]` | Ejecuta la disposición física (reciclaje, disposición controlada) y emite el certificado o comprobante correspondiente. |
| Responsable de registro `[VALIDAR]` | Verifica que el certificado de disposición esté completo, lo archiva vinculado al ID del activo, y ejecuta el cierre definitivo del registro en el sistema. |
| Gobernanza ([Systems #15](../catalogo-systems.md#15-gobernanza)) | Supervisa que el % de activos con disposición certificada se mantenga conforme al objetivo, y usa el cierre auditable para reportes de sostenibilidad/ESG frente a clientes industriales. |

## 5. Diagrama de flujo

```
[Activo en Baja]              [Remanente sin valor en Scrap]
 (sin valor de scrap,           (tras evaluación de
  SOP-ACTIVO-18)                 recuperación de valor,
       │                          SOP-ACTIVO-19)
       │                                │
       └────────────┬───────────────────┘
                     ↓
     [Determinar vía de disposición
      según normativa ambiental
      aplicable al tipo de material]
                     ↓
     [Seleccionar proveedor de
      disposición homologado
      [VALIDAR]]
                     ↓
     [Ejecutar disposición
      (reciclaje / disposición
      controlada)]
                     ↓
     [Proveedor emite certificado
      o comprobante de disposición]
                     ↓
     [Verificar y archivar certificado
      vinculado al ID del activo]
                     ↓
     [Registrar transición de estado
      → Dispuesto]
                     ↓
     [Cierre definitivo del registro
      del activo — fin absoluto del
      ciclo de vida]
```

## 6. SOP paso a paso

1. **Confirmar entrada.** Verificar que el activo o remanente proviene de
   `Baja` con decisión ya registrada
   ([SOP-ACTIVO-18](./SOP-ACTIVO-18-baja.md)) sin valor de scrap
   (transición 26), o de `Scrap` con evaluación de recuperación de valor ya
   completada ([SOP-ACTIVO-19](./SOP-ACTIVO-19-scrap.md), transición 27).
2. **Determinar la vía de disposición.** El Responsable de cumplimiento
   ambiental clasifica el material (plástico, metal, componentes mixtos,
   posibles residuos peligrosos según lo que haya contenido el activo) y
   determina la vía aplicable conforme a normativa ambiental vigente
   `[VALIDAR normativa ambiental específica aplicable — federal/estatal/local
   según ubicación de la planta]`.
3. **Seleccionar proveedor homologado.** Se asigna el proveedor de
   disposición correspondiente a esa vía `[VALIDAR proveedor(es) de
   disposición homologados — no existe todavía una lista confirmada]`. No
   debe usarse un proveedor no homologado, aunque sea más económico o
   esté disponible de inmediato `[VALIDAR proceso de homologación de nuevos
   proveedores]`.
4. **Coordinar la entrega/retiro.** Se coordina con el proveedor la
   recolección o entrega del activo/remanente, documentando cantidad,
   tipo de material y fecha.
5. **Ejecutar la disposición.** El proveedor ejecuta la disposición
   (reciclaje o disposición controlada, según corresponda).
6. **Obtener el certificado.** El proveedor emite el certificado o
   comprobante de disposición final. Este documento es el que cierra
   formalmente el registro del activo — sin él, la disposición no puede
   darse por completa `[VALIDAR formato/contenido mínimo exigido al
   certificado: fecha, cantidad, tipo de material, normativa referenciada]`.
7. **Verificar y archivar.** El Responsable de registro verifica que el
   certificado esté completo y lo archiva vinculado al ID del activo en el
   registro maestro.
8. **Registrar la transición de estado.** El estado del activo cambia a
   `Dispuesto` ([transición 26 o 27](../estados-del-activo.md#2-tabla-de-transiciones)
   según el origen), con marca de tiempo, actor, estado anterior y estado
   nuevo en Trazabilidad ([Systems #7](../catalogo-systems.md#7-trazabilidad)),
   conforme al [invariante #5](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar).
9. **Cerrar el registro del activo.** Se confirma que no existe ninguna
   transición posterior posible ([invariante #4](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)):
   el ciclo de vida del activo queda cerrado de forma auditable en Registro
   (Módulo 1 de [CHACONTAINER OS](../catalogo-systems.md#16-chacontainer-os)).

## 7. Checklist

- [ ] Activo/remanente confirmado en `Baja` (sin valor de scrap) o `Scrap`
      (remanente tras recuperación de valor), con el SOP previo completo.
- [ ] Vía de disposición determinada conforme a normativa ambiental
      aplicable al tipo de material.
- [ ] Proveedor de disposición homologado seleccionado (no un proveedor
      ad hoc sin validar).
- [ ] Entrega/retiro coordinado y documentado (cantidad, tipo de material,
      fecha).
- [ ] Disposición ejecutada por el proveedor.
- [ ] Certificado o comprobante de disposición obtenido y verificado
      como completo.
- [ ] Certificado archivado vinculado al ID del activo.
- [ ] Transición a `Dispuesto` registrada en Trazabilidad con fecha,
      actor, estado anterior y estado nuevo.
- [ ] Registro del activo cerrado de forma definitiva (sin transición
      posterior posible).

## 8. KPI

- % de activos con disposición certificada (vs. activos dispuestos sin
  evidencia) — debe ser 100%; es el indicador que exige explícitamente el
  catálogo de Solutions para este servicio (ver
  [Disposición final, campo "Indicador que afecta"](../catalogo-solutions.md#14-disposición-final)).
- Costo de disposición por unidad/tipo de material.
- Tiempo entre entrada a `Baja`/`Scrap` (según la ruta) y cierre en
  `Dispuesto`.
- % de proveedores de disposición homologados sobre proveedores utilizados
  — controla el riesgo de usar proveedores no autorizados.
- Volumen dispuesto por periodo, segmentable por tipo de material —
  insumo directo para reportes de sostenibilidad/ESG frente a clientes
  industriales.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Activo dispuesto sin certificado o con certificado incompleto | Paso 6–7 obligatorio: no se registra la transición a `Dispuesto` (paso 8) sin certificado verificado y archivado. |
| Uso de proveedor de disposición no homologado (riesgo legal/ambiental y reputacional) | Lista de proveedores homologados mantenida y validada explícitamente `[VALIDAR proveedor(es) de disposición homologados]`; no se coordina disposición con un proveedor fuera de esa lista sin proceso de homologación previo. |
| Disposición ejecutada sin considerar normativa ambiental aplicable al tipo de material (p. ej. residuos peligrosos tratados como residuo común) | Clasificación de material y determinación de vía de disposición (paso 2) obligatoria antes de seleccionar proveedor `[VALIDAR normativa ambiental específica]`. |
| Activo "desaparece" del registro sin trazabilidad de qué pasó con él | Transición a `Dispuesto` siempre registrada en Trazabilidad (paso 8) conforme al invariante #5 — ningún activo sale del sistema sin ese registro. |
| Cierre de registro ejecutado antes de tener el certificado, exponiendo a CHACONTAINER si el proveedor no completa la disposición correctamente | Orden estricto de pasos: certificado (6) → verificación (7) → transición de estado (8) → cierre de registro (9); no invertir el orden aunque el proceso físico ya haya iniciado. |

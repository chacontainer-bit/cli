# SOP-ACTIVO-18: Baja

**Estado: Borrador v0.1 — pendiente de validación.**

Este documento aplica el formato fijo de SOP definido en `CLAUDE.md` al
estado terminal `Baja` de la máquina de estados del activo. Se construyó a
partir de las [transiciones 9 y 24 de la tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones)
y del [invariante #4](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)
que declara los estados terminales sin regreso al ciclo operativo — **no** a
partir de una entrevista real con el responsable de planta o con Gobernanza.
No debe declararse estándar hasta validar en operación real los puntos
marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

`Baja` es el estado que marca la decisión de sacar un activo retornable del
parque circulante. No es una ejecución física ni un destino final: es un
punto de **decisión** que cierra la vida operativa del activo (dejó de
poder reincorporarse a `Disponible`) sin determinar todavía qué pasa con su
cuerpo físico. Ese destino se resuelve después, en dos SOP separados —
[SOP-ACTIVO-19 (Scrap)](./SOP-ACTIVO-19-scrap.md) o
[SOP-ACTIVO-20 (Disposición final)](./SOP-ACTIVO-20-disposicion-final.md).
Este SOP cubre las dos transiciones que llevan a `Baja`: `En reparación` →
`Baja` (activo no reparable, [transición 9](../estados-del-activo.md#2-tabla-de-transiciones))
y `Perdido` → `Baja` (recuperación agotada, [transición 24](../estados-del-activo.md#2-tabla-de-transiciones)).

## 2. Objetivo

Garantizar que ningún activo salga del ciclo operativo por decisión
informal o unilateral: que exista un criterio explícito, una autoridad
definida y evidencia registrada antes de mover un activo a `Baja`, y que
`Baja` quede claramente separado de la ejecución posterior (Scrap o
Disposición final).

## 3. Alcance

Aplica a todo activo retornable (contenedor, IBC, tarima, tambor —
convención de [catalogo-solutions.md](../catalogo-solutions.md)) que llega
al estado `Baja` desde `En reparación` (diagnóstico de no reparable) o
desde `Perdido` (recuperación agotada). No incluye la valoración ni compra
de scrap (cubierto en SOP-ACTIVO-19) ni la disposición física del activo
(cubierto en SOP-ACTIVO-20). No aplica a activos que solo requieren
reparación, reacondicionamiento o siguen en proceso de recuperación activa
— esos casos permanecen en `En reparación` o `Perdido` hasta que se cumpla
el criterio de disparo descrito en la sección 6.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Técnico de reparación `[VALIDAR]` | Emite el diagnóstico técnico de "no reparable" tras evaluar el activo proveniente de `En reparación`. No tiene autoridad para decidir la baja por sí solo. |
| Responsable de planta `[VALIDAR]` | Autoridad de primera instancia para aprobar la baja de un activo individual con diagnóstico de no reparable, según el [criterio de disparo](#6-sop-paso-a-paso) definido en este SOP. |
| Gobernanza ([Systems #15](../catalogo-systems.md#15-gobernanza)) | Autoridad que decide la baja de activos en estado `Perdido` (agotado plazo/costo de recuperación) y resuelve los casos que excedan el umbral de valor o volumen definido para el Responsable de planta `[VALIDAR umbral exacto]`. Consume Indicadores y Analítica ([Systems #13–14](../catalogo-systems.md#13-indicadores)) para decidir "qué activos deben darse de baja" ([pregunta de gobernanza, ETAPA 4](../README.md#etapa-4--modelo-de-gobernanza)). |
| Responsable de registro `[VALIDAR]` | Ejecuta el cambio de estado en el sistema (Capa 4) y deja la evidencia de la decisión vinculada al activo en Registro (Módulo 1 de [CHACONTAINER OS](../catalogo-systems.md#16-chacontainer-os)). |

## 5. Diagrama de flujo

```
[En reparación]                          [Perdido]
       │                                      │
Diagnóstico: no reparable          Excede plazo/costo de
       │                           recuperación sin localizar
       ↓                                      ↓
[¿Cumple criterio de baja?] ──────[¿Cumple criterio de baja?]
       │ sí                                   │ sí
       ↓                                      ↓
[Aprobación: Responsable de planta   [Aprobación: Gobernanza]
 (o Gobernanza si excede umbral)]
       │                                      │
       └──────────────┬───────────────────────┘
                       ↓
         [Registrar decisión + evidencia]
                       ↓
              [Estado → Baja]
       (fin de vida operativa, sin destino
        físico determinado todavía)
                       ↓
         [Enviar a evaluación de destino:
          SOP-ACTIVO-19 (Scrap) o
          SOP-ACTIVO-20 (Disposición final)]
```

## 6. SOP paso a paso

1. **Detectar el disparador.** El activo llega a este SOP por una de dos
   vías: (a) diagnóstico de "no reparable" durante `En reparación`
   ([transición 9](../estados-del-activo.md#2-tabla-de-transiciones), ver
   [SOP-06 de reparación referenciado en el ciclo de vida](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado));
   o (b) agotamiento del plazo/costo de recuperación de un activo `Perdido`
   ([transición 24](../estados-del-activo.md#2-tabla-de-transiciones)).
2. **Emitir diagnóstico o reporte de agotamiento.** El técnico de
   reparación documenta por qué el activo no es reparable (costo de
   reparación vs. valor de reemplazo, daño estructural irreversible
   `[VALIDAR criterio técnico exacto de "no reparable"]`); para `Perdido`,
   Logística inversa/Recuperación documenta que se agotaron los intentos
   razonables de localización `[VALIDAR plazo/costo máximo de búsqueda antes de declarar agotado]`.
3. **Verificar criterio de disparo.** Confirmar que el caso cumple al
   menos uno de estos criterios objetivos antes de escalar a aprobación
   `[VALIDAR lista cerrada de criterios]`:
   - Costo estimado de reparación mayor a un porcentaje del valor de
     reemplazo del activo `[VALIDAR porcentaje]`.
   - Daño estructural que compromete la función primaria del activo de
     forma irreversible.
   - Activo `Perdido` sin ubicación confirmable tras el plazo máximo de
     búsqueda definido en Reglas operativas ([Systems #11](../catalogo-systems.md#11-reglas-operativas)).
4. **Escalar a la autoridad correspondiente.** Si el activo viene de `En
   reparación` y está dentro del umbral definido, el Responsable de planta
   aprueba o rechaza la baja. Si excede el umbral, o si el activo viene de
   `Perdido`, la decisión pasa a Gobernanza `[VALIDAR umbral exacto y
   proceso de escalamiento]`.
5. **Registrar la decisión.** Sin importar quién decide, la aprobación (o
   rechazo) queda registrada con fecha, responsable, motivo y evidencia de
   soporte (diagnóstico técnico, reporte de recuperación agotada) —
   vinculada al ID del activo en el registro maestro.
6. **Ejecutar la transición de estado.** El estado del activo cambia a
   `Baja` en el sistema. Esta transición queda registrada en
   Trazabilidad ([Systems #7](../catalogo-systems.md#7-trazabilidad)) con
   marca de tiempo, actor, estado anterior y estado nuevo, conforme al
   [invariante #5](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar).
7. **Confirmar que `Baja` es terminal de decisión, no de ejecución.** El
   activo en `Baja` no tiene destino físico determinado todavía. No se
   ejecuta compra de scrap ni disposición desde este SOP — solo se marca
   la decisión.
8. **Derivar a evaluación de destino.** El activo en `Baja` pasa a
   evaluación de [Compra de scrap](../catalogo-solutions.md#13-compra-de-scrap)
   ([SOP-ACTIVO-19](./SOP-ACTIVO-19-scrap.md)) si tiene valor de scrap
   recuperable, o directo a
   [Disposición final](../catalogo-solutions.md#14-disposición-final)
   ([SOP-ACTIVO-20](./SOP-ACTIVO-20-disposicion-final.md)) si no lo tiene —
   ver [transiciones 25 y 26](../estados-del-activo.md#2-tabla-de-transiciones).

## 7. Checklist

- [ ] Disparador identificado (diagnóstico de no reparable, o recuperación
      agotada) y documentado.
- [ ] Caso verificado contra el criterio de disparo objetivo (no decisión
      informal).
- [ ] Aprobación registrada por la autoridad correspondiente (Responsable
      de planta o Gobernanza según umbral).
- [ ] Evidencia de soporte (diagnóstico técnico o reporte de recuperación
      agotada) vinculada al ID del activo.
- [ ] Transición a `Baja` registrada en Trazabilidad con fecha, actor,
      estado anterior y estado nuevo.
- [ ] Confirmado que no se ejecutó destino físico (scrap ni disposición)
      en este paso.
- [ ] Activo derivado explícitamente a evaluación de scrap o a disposición
      final directa.

## 8. KPI

- % de bajas con aprobación registrada (vs. bajas sin evidencia de
  autoridad) — debe ser 100%.
- Tiempo entre diagnóstico/reporte de agotamiento y decisión final de
  baja.
- % de bajas originadas en `En reparación` vs. originadas en `Perdido` —
  insumo para priorizar dónde reducir pérdidas (ver
  [ETAPA 4](../README.md#etapa-4--modelo-de-gobernanza)).
- Costo evitado (o no evitado) por decidir baja vs. seguir reparando —
  requiere el dato de costo real de reparación ([Solutions #4](../catalogo-solutions.md#4-reparación)).
- % de activos dados de baja que efectivamente avanzan a Scrap o
  Disposición final dentro de un plazo objetivo `[VALIDAR plazo]` — evita
  que `Baja` se convierta en un estado "estacionado" sin cierre.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Baja decidida sin autoridad clara (criterio informal de quien esté de turno) | Autoridad y umbral de escalamiento definidos por escrito en este SOP (sección 4 y 6), no dejados al juicio individual `[VALIDAR umbral exacto]`. |
| `Baja` se confunde operativamente con "ya se dispuso" o "ya se vendió como scrap" | Este SOP declara explícitamente que `Baja` es decisión, no ejecución (sección 1 y 6.7); el sistema no debe permitir marcar el activo como retirado del inventario financiero hasta que exista un SOP-ACTIVO-19 o SOP-ACTIVO-20 ejecutado `[VALIDAR si el sistema ya impone esta regla]`. |
| Activo queda "atascado" en `Baja` sin avanzar a Scrap ni a Disposición final | KPI de plazo (sección 8) y alerta automática si un activo permanece en `Baja` más allá del umbral definido ([Systems #12 · Alertas](../catalogo-systems.md#12-alertas)) `[VALIDAR umbral]`. |
| Diagnóstico de "no reparable" sesgado por incentivo de quien repara (reparar genera trabajo/ingreso) | Separar el rol que diagnostica del rol que aprueba la baja (sección 4); el Responsable de planta o Gobernanza valida el diagnóstico antes de aprobar, no lo ejecuta automáticamente. |
| Activo `Perdido` dado de baja prematuramente, antes de agotar razonablemente la búsqueda | Criterio de plazo/costo máximo de búsqueda definido explícitamente antes de declarar agotado (paso 2 y 3) `[VALIDAR plazo/costo]`, en línea con [SOP-16 de activo perdido](../ciclo-de-vida-del-activo.md#4-ramas-de-excepción-fuera-del-flujo-feliz) referenciado en el ciclo de vida. |

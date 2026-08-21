# SOP-ACTIVO-13: Retorno

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre dos momentos del retorno del activo retornable: el **inicio del retorno** cuando el
custodio o la logística lo devuelve voluntariamente (`En uso → En tránsito`), y la **recolección activa**
cuando el activo excede el plazo permitido y queda `Retenido` (`En uso → Retenido` y, tras la recolección,
`Retenido → En tránsito`). Corresponde a las filas 16, 17 y 18 de la
[tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones), y cierra la etapa "Retorno" del
[ciclo de vida del activo](../ciclo-de-vida-del-activo.md#2-tabla-de-trazabilidad-cruzada-por-etapa),
entregando el activo a Reinspección (`En tránsito → En inspección`, fila 15). El caso de retorno forzado se
apoya en el servicio [Recuperación](../catalogo-solutions.md#12-recuperación) del catálogo Solutions. Se
construyó a partir del cruce entre ese catálogo, el [catálogo Systems](../catalogo-systems.md#9-logística-inversa)
y la máquina de estados — **no a partir de una entrevista real con el responsable de logística inversa**. No
debe declararse estándar hasta validar en operación los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Retorno es el proceso que trae de vuelta al flujo un activo que estaba `En uso`, por dos caminos distintos.
El camino normal ocurre cuando el custodio o la logística inicia el retorno dentro del plazo acordado: el
activo pasa a `En tránsito` con dirección "de vuelta a planta" (reutilizando el mecanismo de
[SOP-ACTIVO-11](./SOP-ACTIVO-11-movimiento.md)). El camino de excepción ocurre cuando el plazo vence sin que
nadie inicie el retorno: el activo pasa a `Retenido`, y CHACONTAINER debe ejecutar una recolección activa —
el servicio [Recuperación](../catalogo-solutions.md#12-recuperación) de Solutions, sostenido por
[Logística inversa](../catalogo-systems.md#9-logística-inversa) de Systems — para forzar su regreso. Ambos
caminos terminan en el mismo punto: el activo llega a planta y entra a Reinspección, fuera del alcance de
este SOP.

## 2. Objetivo

Garantizar que todo activo `En uso` regresa al flujo de CHACONTAINER dentro del plazo acordado, y que
cuando no lo hace, el sistema lo detecta automáticamente, lo marca `Retenido` y dispara una recolección
activa que lo recupere antes de que se convierta en pérdida (`Perdido`).

## 3. Alcance

Aplica desde que un activo en estado `En uso` inicia su regreso — voluntariamente o por vencimiento de
plazo — hasta que llega a planta y queda listo para entrar a Reinspección. Cubre dos escenarios:

1. **Retorno voluntario en plazo**: el custodio o la logística devuelve el activo dentro del tiempo máximo
   fuera definido por las [Reglas operativas](../catalogo-systems.md#11-reglas-operativas).
2. **Retorno forzado (recolección activa)**: el plazo vence, el activo queda `Retenido`, y CHACONTAINER
   ejecuta el servicio [Recuperación](../catalogo-solutions.md#12-recuperación) para retirarlo.

**No incluye:**

- El desplazamiento físico en sí una vez que el activo ya está `En tránsito` de vuelta — reutiliza el
  mecanismo de [SOP-ACTIVO-11 (Movimiento)](./SOP-ACTIVO-11-movimiento.md).
- La reinspección y decisión de reincorporación o baja tras la llegada (`En tránsito → En inspección`,
  fila 15) — corresponde a los SOP de Inspección y Clasificación.
- Las transiciones desde `Bloqueado` o `Perdido` (disputa de custodia, o activo sin ubicación confirmable
  tras exceder el umbral de `Retenido`) — corresponden a SOP-ACTIVO-16 (Activo perdido) y SOP-ACTIVO-17
  (Activo bloqueado), aunque este SOP es el que las origina cuando `Retenido` escala sin resolverse.
- Un cambio de custodio sin regreso físico del activo — ver
  [SOP-ACTIVO-12 (Transferencia de custodia)](./SOP-ACTIVO-12-transferencia-de-custodia.md).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Custodio | Inicia el retorno voluntario dentro del plazo acordado. |
| Logística | Coordina y confirma el retorno voluntario; ejecuta la recolección física cuando el activo está `Retenido`. |
| Sistema (automático, vía Reglas operativas + Alertas) | Detecta el vencimiento del plazo y dispara `En uso → Retenido` sin intervención manual. |
| Logística inversa / Recuperación `[VALIDAR si es el mismo rol que Logística o un equipo distinto]` | Ejecuta el servicio [Recuperación](../catalogo-solutions.md#12-recuperación): localiza, coordina y retira el activo `Retenido`. |
| Responsable de gobernanza | Revisa los activos `Retenido` de forma recurrente y decide cuándo escalar a Bloqueado/Perdido si la recolección no avanza. |

## 5. Diagrama de flujo

```
[Activo en estado "En uso"]
        ↓
   ¿El custodio/logística inicia el retorno antes del plazo?
   ↓ sí                                    ↓ no
[Transición En uso → En tránsito      [Sistema detecta plazo vencido
 (retorno voluntario)                  → Transición En uso → Retenido]
 → SOP-ACTIVO-11 gestiona el                     ↓
   desplazamiento de regreso]         [Alerta generada al responsable
        ↓                              de logística inversa / gobernanza]
        │                                        ↓
        │                             [Ejecutar Recuperación (Solutions):
        │                              localizar, coordinar logística
        │                              inversa, retirar el activo]
        │                                        ↓
        │                             ¿Recolección ejecutada con éxito?
        │                             ↓ sí                    ↓ no (excede umbral adicional)
        │                    [Transición Retenido →     [Escalar a Bloqueado
        │                     En tránsito]               o Perdido — fuera de
        │                             ↓                   alcance de este SOP]
        └─────────────────────────────┤
                                       ↓
                    [Activo llega a planta → entra a Reinspección
                     (En tránsito → En inspección) — fuera de alcance de este SOP]
```

## 6. SOP paso a paso

### Retorno voluntario

1. Verificar que el activo está en estado `En uso` y que el custodio o la logística inicia su regreso.
2. Registrar el inicio del retorno: dispara la transición `En uso → En tránsito`
   ([tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones), fila 16), con dirección "de
   vuelta a planta" en lugar de "hacia el custodio".
3. A partir de este punto, el desplazamiento físico se gestiona con el mismo mecanismo de
   [SOP-ACTIVO-11 (Movimiento)](./SOP-ACTIVO-11-movimiento.md): registro de tránsito, transbordos si aplica,
   y confirmación de llegada.
4. Al llegar a planta, el activo entra al flujo de Reinspección (`En tránsito → En inspección`, fila 15),
   fuera del alcance de este SOP.

### Retorno forzado (recolección activa)

5. El sistema evalúa continuamente el tiempo transcurrido desde la asignación/entrega contra el tiempo
   máximo fuera definido en [Reglas operativas](../catalogo-systems.md#11-reglas-operativas) para ese tipo
   de activo/cliente `[VALIDAR umbral exacto de días — no hay un valor único definido; depende del tipo de
   activo y del acuerdo comercial]`.
6. Si el plazo vence sin que se haya registrado un movimiento de retorno, el sistema dispara automáticamente
   la transición `En uso → Retenido` (fila 17) y genera una [Alerta](../catalogo-systems.md#12-alertas) al
   responsable correspondiente.
7. El responsable de logística inversa ejecuta el servicio
   [Recuperación](../catalogo-solutions.md#12-recuperación) de Solutions: localiza el activo (usando la
   última ubicación confirmada en Capa 2), obtiene autorización del custodio/propietario para retirarlo, y
   coordina la logística inversa.
8. Retirar el activo y dejar evidencia: orden de recuperación, acta de retiro firmada por el custodio,
   registro de condición al momento de la recolección — evidencia que exige el propio catálogo Solutions
   para este servicio.
9. Si la recolección se ejecuta con éxito, registrar la transición `Retenido → En tránsito`
   (fila 18) y continuar como un tránsito de regreso normal (pasos 3-4).
10. Si la recolección no se logra dentro de un umbral adicional
    `[VALIDAR umbral de escalamiento — cuánto tiempo en Retenido antes de escalar]`, escalar el caso hacia
    `Bloqueado` (si hay disputa de custodia o condición) o `Perdido` (si no hay ubicación confirmable),
    fuera del alcance de este SOP.
11. Actualizar el indicador de cumplimiento de retorno y el de activos recuperados sobre activos
    `Retenido` identificados.

## 7. Checklist

- [ ] Retorno voluntario registrado con la transición `En uso → En tránsito` antes del vencimiento del plazo.
- [ ] Umbral de tiempo máximo fuera configurado en Reglas operativas para el tipo de activo/cliente
      `[VALIDAR valor]`.
- [ ] Alerta generada automáticamente en el momento en que el activo pasa a `Retenido` (no detectado
      manualmente días después).
- [ ] Autorización del custodio/propietario obtenida antes de ejecutar la recolección forzada.
- [ ] Acta de retiro y registro de condición generados en toda recolección de un activo `Retenido`.
- [ ] Transición `Retenido → En tránsito` registrada al ejecutarse la recolección con éxito.
- [ ] Casos de `Retenido` sin resolver dentro del umbral adicional escalados a Bloqueado/Perdido, no dejados
      abiertos indefinidamente.

## 8. KPI

- **Cumplimiento de retorno**: retornos en tiempo / retornos esperados × 100 (KPI de la
  [ETAPA 6 del README](../README.md#etapa-6--indicadores-del-sistema)).
- % de activos `Retenido` recuperados exitosamente sobre el total de activos `Retenido` identificados
  (equivalente al KPI de Recuperación del [catálogo Solutions](../catalogo-solutions.md#12-recuperación)).
- Tiempo detenido: días entre el vencimiento del plazo (`En uso → Retenido`) y la recolección efectiva
  (`Retenido → En tránsito`).
- Costo de recuperación por activo recolectado.
- % de casos `Retenido` escalados a `Bloqueado` o `Perdido` (mide qué tan efectiva es la recolección activa
  antes de convertirse en pérdida).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Umbral de tiempo máximo fuera no está definido o es inconsistente entre tipos de activo/cliente | Definir y documentar el umbral por tipo de activo/cliente en Reglas operativas antes de habilitar la transición automática `[VALIDAR]` — ver invariante 6 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar). |
| Activo pasa a `Retenido` pero nadie actúa sobre la alerta (se acumulan activos retenidos sin recolección) | Revisión recurrente obligatoria de la cola de activos `Retenido` por el responsable de gobernanza; medir tiempo detenido como KPI explícito. |
| Recolección forzada ejecutada sin autorización clara del custodio (riesgo legal/comercial) | Obtener y documentar la autorización antes de retirar el activo (paso 7-8), igual que exige el servicio Recuperación del catálogo Solutions. |
| Caso `Retenido` que nunca escala ni se resuelve — queda en limbo indefinidamente | Definir un umbral de escalamiento explícito hacia `Bloqueado`/`Perdido` `[VALIDAR umbral]`, para que ningún activo permanezca `Retenido` sin fecha límite de acción. |
| Confusión entre retorno voluntario y recolección forzada al registrar la transición — se usa la misma transición `En uso → En tránsito` para ambos casos sin distinguir el origen | Registrar explícitamente en el evento de trazabilidad si el retorno fue voluntario o si vino precedido de un paso por `Retenido`, para no perder la señal de incumplimiento en los reportes. |

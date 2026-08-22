# Alertas — de la regla a la acción

**Fase:** III · Producto — paso 14 de la [ETAPA 14](./README.md#etapa-14--orden-de-ejecución) del [Proyecto Maestro](./README.md).

[reglas-operativas.md](./reglas-operativas.md) definió qué es una regla y
cómo se resuelve cuál aplica. Este documento define el mecanismo que las
evalúa y convierte un incumplimiento en algo que alguien atiende — el
**Nivel 1** de [modelo-de-gobernanza.md](./modelo-de-gobernanza.md#1-los-tres-niveles-de-gobernanza).

---

## 1. El riesgo real: el ruido, no la cobertura

Un sistema de alertas falla casi siempre por exceso, no por defecto. Si el
responsable de planta abre la bandeja y ve 200 alertas, deja de abrirla, y
el sistema completo —reglas, estados, custodia— queda sin consumidor. Todo
el diseño que sigue está orientado a eso.

Tres decisiones concretas:

1. **Una alerta por activo y regla, no una por evaluación.** Si un activo
   lleva 12 días excedido, hay *una* alerta con 12 días de antigüedad, no
   12 alertas diarias. Se implementa con una restricción de unicidad sobre
   `(asset_id, rule_id)` para alertas en estado `activa`.
2. **La alerta se cierra sola cuando la condición desaparece.** Si el activo
   retorna, su alerta de tiempo excedido pasa a `atendida` automáticamente,
   sin que nadie la marque. Solo requieren cierre manual las que exigieron
   una decisión.
3. **Agregación por custodio.** Cuarenta activos retenidos por el mismo
   cliente son un problema, no cuarenta. La vista Operación
   ([dashboard.md §2](./dashboard.md#2-vista-operación--la-bandeja-de-excepciones))
   los agrupa por custodio y muestra el detalle al expandir.

## 2. Qué genera una alerta

Toda alerta nace de una regla de
[reglas-operativas.md §1](./reglas-operativas.md#1-qué-es-una-regla).
Los cuatro tipos, con su comportamiento:

| Tipo de regla | Condición evaluada | Acción |
|---|---|---|
| Tiempo máximo fuera | Días en `En uso` sin retorno > umbral | Transición automática a `Retenido` **y** alerta |
| Punto de retorno | Fecha esperada de retorno − hoy ≤ ventana de aviso `[VALIDAR ventana]` | Alerta anticipada, sin transición |
| Nivel mínimo de inventario | Conteo en `Disponible` por tipo/planta < umbral | Alerta, sin transición (no aplica a un activo, sino a un agregado) |
| Condición permitida | Resultado de inspección fuera del criterio | Alerta solo si el activo estaba asignado a un cliente con criterio más estricto que el default |

La regla de nivel mínimo es la única que no se ancla a un `asset_id`. En el
esquema propuesto en [modelo-de-datos.md](./modelo-de-datos.md#alerts-alertas--no-existe-hoy)
eso implica que `alerts.asset_id` debe ser nullable — la alerta apunta a la
regla y al alcance (tipo de activo + planta), no a una unidad.

## 3. El evaluador

Un job periódico, siguiendo el patrón que ya usa `chacontainer/scripts/`
para sync y reportes (ver [architecture.md](../chacontainer/docs/architecture.md#python-automation-scripts)):

```
Por cada regla activa (operational_rules):
  Resolver el conjunto de activos en su alcance (scope)
  Evaluar la condición contra asset_state_detail / asset_custody / conteos
  Para cada incumplimiento:
    ¿Existe ya una alerta activa para (asset_id, rule_id)?
      Sí  → no hacer nada (evita el ruido de §1.1)
      No  → crear alerta + notificar al responsable
    Si la regla tiene accion = transicion_automatica:
      Aplicar la transición y registrarla en asset_events
Por cada alerta activa cuya condición ya no se cumple:
  Cerrar como 'atendida' (§1.2)
```

**Decidido: cada hora.** Los umbrales del modelo se miden en días, así que
evaluar cada minuto no aporta precisión y sí carga; evaluar una vez al día
introduce hasta 24 h de retraso en detectar un activo detenido, que es
justo lo que el piloto quiere reducir. Cada hora dentro de esa ventana:
retraso máximo aceptable y carga de cómputo despreciable frente al volumen
de un piloto (200-500 activos). Si el volumen crece varios órdenes de
magnitud (escala nacional, [escala-geografica.md §2](./escala-geografica.md#2-operación-nacional-paso-26)),
revisar si conviene evaluar por lotes segmentados en vez de bajar la
frecuencia — pero no antes de tener esa escala real.

**Idempotencia**: el job debe poder correr dos veces seguidas sin duplicar
alertas ni re-aplicar transiciones. La restricción de unicidad de §1.1 lo
garantiza para alertas; para transiciones, lo garantiza la validación de
estado (una transición desde `Retenido` a `Retenido` no existe en
[la tabla](./estados-del-activo.md#2-tabla-de-transiciones)).

## 4. Notificación

| Canal | Uso | Estado |
|---|---|---|
| Bandeja en el dashboard | Todas las alertas | MVP |
| Correo | Resumen diario de alertas nuevas al responsable de planta | MVP `[VALIDAR si se usa correo o el canal que ya usen]` |
| Webhook Make.com | Integración con el flujo que el cliente ya tenga | Ya existe la infraestructura; conectar solo si el piloto lo pide |

Nada de notificación push por alerta individual: contradice §1.

## 5. Escalada

Una alerta que permanece `activa` más allá de un umbral de tiempo
`[VALIDAR umbral por tipo]` deja de ser un asunto de Nivel 2 y escala a
Nivel 3, según los criterios de
[modelo-de-gobernanza.md §4](./modelo-de-gobernanza.md#4-qué-activa-una-escalada-de-nivel-2-a-nivel-3).
La escalada no crea una alerta nueva: cambia a quién se le muestra y con
qué prioridad, para no violar §1.1.

El caso más importante es el recurrente: si el mismo custodio genera el
mismo tipo de alerta tres o más veces en el período, el problema ya no es
operativo sino comercial, y lo que corresponde no es recolectar otra vez
sino renegociar la condición — precisamente lo que la
[ETAPA 8](./README.md#etapa-8--modelo-comercial-solutions--systems) describe
como el momento en que CHACONTAINER pasa de ejecutar a gobernar.

---

## Próximo paso

[Piloto](./piloto.md) (paso 15): cómo se prueba todo esto en operación real.

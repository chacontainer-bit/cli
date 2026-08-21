# SOP-ACTIVO-05: Lavado

**Estado: Borrador v0.1 — pendiente de validación.**

Este SOP documenta la etapa **Limpieza** del [ciclo de vida del activo](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado)
del Proyecto Maestro CHACONTAINER: la transición del estado `Sucio` a `En
inspección` en la [máquina de estados](../estados-del-activo.md#1-catálogo-de-estados).
Corresponde al servicio **Lavado** (servicio #3) del [catálogo Solutions](../catalogo-solutions.md#3-lavado)
y se apoya en la capacidad **Trazabilidad** (#7) del [catálogo Systems](../catalogo-systems.md#7-trazabilidad)
para registrar el evento.

Es un documento propio de este proyecto, no una copia de
[SOP-02 (lavado de activos retornables)](../../chacontainer/docs/sop/SOP-02-lavado-de-activos-retornables.md)
del SaaS `chacontainer/`. Se usó SOP-02 como referencia de tono y nivel de
detalle, pero la terminología de estados aquí sigue exclusivamente
[estados-del-activo.md](../estados-del-activo.md): un activo lavado no queda
"apto" o "rechazado" — pasa a `En inspección`, donde la inspección (SOP-ACTIVO-03,
pendiente) decide si continúa a `Liberado` o retrocede a `En reparación`.

---

## 1. Resumen ejecutivo

El lavado recibe un activo retornable (contenedor, IBC, tarima o tambor) en
estado `Sucio` — ya sea porque acaba de retornar de campo o porque una
inspección previa lo derivó a limpieza — y lo devuelve al ciclo en estado `En
inspección`, limpio y con evidencia registrada de que el lavado se ejecutó.
El lavado por sí mismo **no decide** si el activo puede reingresar a
inventario disponible: esa decisión es de la inspección posterior (SOP-ACTIVO-03),
que puede confirmar el paso a `Liberado` o detectar daño y derivarlo a `En
reparación` (SOP-ACTIVO-06). El lavado solo garantiza que el activo llega
limpio a esa evaluación.

## 2. Objetivo

Garantizar que todo activo que transiciona de `Sucio` a `En inspección` fue
lavado con un método consistente según su tipo y su producto previo, dejando
evidencia trazable del evento, sin depender del criterio informal de un solo
operador.

## 3. Alcance

Aplica a todo activo retornable en estado `Sucio` dentro de la [máquina de
estados](../estados-del-activo.md#1-catálogo-de-estados), sin importar en qué
punto del ciclo entró a ese estado:

- Activo que retorna de campo y llega directo a `Sucio` (transición típica
  vía `En tránsito` → `En inspección` → `Sucio`, transición #4 de la [tabla
  de transiciones](../estados-del-activo.md#2-tabla-de-transiciones)).
- Activo derivado a `Sucio` desde una reinspección posterior a reparación,
  si el reacondicionamiento (SOP-ACTIVO-06) incluyó una etapa de limpieza
  profunda que se ejecuta como este mismo proceso `[VALIDAR: si reacondicionamiento
  siempre delega el paso de limpieza a este SOP, o si mantiene su propio
  procedimiento de limpieza integral]`.
- Cliente que entra directo pidiendo el servicio de lavado sobre activos que
  nunca pasaron por Alta/Identificación formal (ver [§5 de ciclo-de-vida-del-activo.md](../ciclo-de-vida-del-activo.md#5-punto-de-entrada-del-cliente--inicio-del-ciclo)) —
  en ese caso este SOP se ejecuta igual, y el activo sin identificar queda
  marcado como oportunidad de venta de Identificación (Systems #3).

**No incluye:**

- La reparación estructural del activo (cubierta por [SOP-ACTIVO-06](./SOP-ACTIVO-06-reparacion.md)).
- El transporte de retorno desde el cliente hasta la estación de lavado
  (cubierto por logística inversa, SOP-14, fuera del alcance de este
  documento).
- La decisión de apto/no apto para reingreso, que corresponde a la
  inspección posterior (`En inspección`, SOP-ACTIVO-03).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Operador de lavado `[VALIDAR]` | Ejecuta el lavado según el método definido por tipo de activo y producto previo, y registra el evento de inicio/fin. |
| Responsable de planta `[VALIDAR]` | Define y mantiene el método de lavado por tipo de activo, resuelve excepciones (p. ej. residuo no identificado) y autoriza reprocesos. |
| Sistema (CHACONTAINER OS) | Registra automáticamente la transición `Sucio` → `En inspección` en trazabilidad cuando el operador confirma el lavado completado (ver invariante 5 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar)). |

## 5. Diagrama de flujo

```
[Activo en estado "Sucio"]
              ↓
[Identificar tipo de activo y producto previo transportado]
              ↓
[Confirmar método de lavado aplicable por tipo/producto]
              ↓
[Ejecutar lavado (prelavado → lavado → enjuague → secado)]
              ↓
[Verificar ausencia visual de residuos y etiquetas previas]
              ↓
       ¿Residuo persiste o requiere reproceso?
       ↓ sí                              ↓ no
[Relavar / escalar a responsable   [Registrar evento de lavado
 de planta]                         completado en trazabilidad]
       ↑___________________________________|
                                            ↓
                          [Transición de estado: Sucio → En inspección]
                                            ↓
                      [Activo pasa a SOP-ACTIVO-03 · Inspección]
```

## 6. SOP paso a paso

1. Confirmar que el activo está registrado en estado `Sucio` en el sistema
   (transición #4 de la tabla de transiciones, disparada por inspección
   previa, o ingreso directo de un cliente que solicita lavado).
2. Identificar el tipo de activo (contenedor, IBC, tarima, tambor) y, si es
   posible determinarlo, el producto que transportó previamente
   `[VALIDAR: procedimiento cuando el producto previo es desconocido —
   tratar como caso de máximo riesgo de contaminación por defecto]`.
3. Confirmar el método de lavado aplicable a esa combinación tipo/producto
   `[VALIDAR método, químicos y concentraciones exactas por tipo de activo —
   pendiente de definición operativa, ver también SOP-02 del SaaS como
   referencia de partida]`.
4. Ejecutar el lavado: retiro de residuos sólidos, retiro de etiquetas
   previas (si aplica y no corresponde a SOP-ACTIVO-07 · Reetiquetado),
   prelavado, lavado, enjuague y secado `[VALIDAR tiempos estándar por
   tipo de activo]`.
5. Verificar visualmente que no quedan residuos, olores ni etiquetas
   previas antes de dar el lavado por concluido. Esta verificación es
   operativa (parte del propio lavado), **no** sustituye la inspección
   formal de condición que ocurre en el siguiente estado.
6. Si persiste residuo o hay duda razonable de contaminación cruzada,
   relavar o escalar al responsable de planta antes de continuar
   `[VALIDAR criterio objetivo de "residuo persistente" — hoy depende del
   juicio del operador]`.
7. Registrar en el sistema el evento de lavado completado: activo, tipo,
   fecha/hora, operador responsable y, si aplica, evidencia fotográfica.
8. Disparar la transición de estado `Sucio` → `En inspección` (transición
   #7 de [estados-del-activo.md](../estados-del-activo.md#2-tabla-de-transiciones)),
   dejando registrado en trazabilidad el evento con marca de tiempo, actor
   y estado anterior/nuevo, conforme al invariante 5.
9. El activo queda disponible para que la inspección posterior
   (SOP-ACTIVO-03, pendiente) determine si continúa a `Liberado` o si
   requiere `En reparación`.

## 7. Checklist

- [ ] Activo confirmado en estado `Sucio` antes de iniciar.
- [ ] Tipo de activo y producto previo identificados (o marcados como
      desconocidos y tratados como riesgo máximo).
- [ ] Método de lavado aplicado corresponde al definido para ese tipo/producto.
- [ ] Verificación visual de ausencia de residuos, olores y etiquetas
      previas realizada antes de cerrar el lavado.
- [ ] Reprocesos (si los hubo) documentados con su motivo.
- [ ] Evento de lavado completado registrado en el sistema con fecha,
      operador y evidencia.
- [ ] Transición `Sucio` → `En inspección` reflejada en trazabilidad el
      mismo día del lavado.

## 8. KPI

- Tiempo de ciclo de lavado por activo o por lote.
- % de activos que requieren reproceso (relavado) sobre el total lavado.
- % de activos que, tras pasar por lavado y llegar a `En inspección`, son
  derivados a `En reparación` por daño detectado — indicador cruzado con
  SOP-ACTIVO-06, útil para saber si el lavado está exponiendo daño
  preexistente en vez de causarlo.
- Rotación de activos que salen de `Sucio` hacia `En inspección` por
  día/semana (evita acumulación de inventario sucio sin rotar).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Contaminación cruzada entre productos incompatibles | Identificar producto previo antes de lavar y aplicar método/insumos definidos para esa combinación; tratar como riesgo máximo cuando el producto previo es desconocido. |
| Uso incorrecto de químicos de lavado (riesgo de seguridad y salud) `[VALIDAR]` | Definir por escrito químico y concentración por tipo de activo, con EPP obligatorio para el operador. |
| Activo "lavado" con daño estructural no detectado | El lavado no decide aptitud — la transición obligatoria a `En inspección` (nunca directo a `Liberado` o `Disponible`) garantiza una evaluación de condición separada, conforme al invariante 2 de [estados-del-activo.md](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar). |
| Dependencia de un solo operador con el criterio de "residuo persistente" | Documentar por escrito el criterio de reproceso y capacitar a más de un operador; escalar casos dudosos al responsable de planta. |
| Transición de estado no registrada (activo limpio pero el sistema sigue mostrando `Sucio`) | Registrar el evento de lavado y disparar la transición el mismo día, conforme al invariante 5 (toda transición queda en trazabilidad, sin excepción). |

---

**Documentos relacionados:** [catalogo-solutions.md](../catalogo-solutions.md) ·
[catalogo-systems.md](../catalogo-systems.md) ·
[ciclo-de-vida-del-activo.md](../ciclo-de-vida-del-activo.md) ·
[estados-del-activo.md](../estados-del-activo.md) ·
[SOP-ACTIVO-06 · Reparación / reacondicionamiento](./SOP-ACTIVO-06-reparacion.md)

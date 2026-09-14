# SOP-ACTIVO-15: Incidencias

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre el registro y seguimiento de cualquier incidencia (daño, retraso, faltante o pérdida)
que puede originarse en cualquier etapa del [ciclo de vida del activo](../ciclo-de-vida-del-activo.md), no
solo en retorno. Se construyó a partir de la capacidad
[Gestión de incidencias (#10)](../catalogo-systems.md#10-gestión-de-incidencias) del catálogo Systems y de
la fila "Incidencia (daño, retraso, faltante)" en
[ciclo-de-vida-del-activo.md §4](../ciclo-de-vida-del-activo.md#4-ramas-de-excepción-fuera-del-flujo-feliz) —
**no a partir de una entrevista real con el responsable operativo**. No debe declararse estándar hasta
validar en operación los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Una incidencia es cualquier evento no planeado que afecta a un activo o a un lote: daño, retraso, faltante
o pérdida. Hoy estos eventos ocurren de forma dispersa — por WhatsApp, llamada, correo o simplemente
observación de campo — y sin un registro único no se puede saber si un problema es aislado o un patrón
sistemático con un cliente, custodio o ruta. Este SOP recibe el reporte de una incidencia desde cualquier
punto de contacto (cliente, operador de campo, proceso interno de inspección, logística inversa), la
registra con tipo, activo(s) involucrado(s) y responsable, y da seguimiento hasta su resolución. Es la
ejecución operativa de [Gestión de incidencias](../catalogo-systems.md#10-gestión-de-incidencias) — el
Módulo 8 de [CHACONTAINER OS](../catalogo-systems.md#16-chacontainer-os) — y el insumo con el que
[Gobernanza](../catalogo-systems.md#15-gobernanza) prioriza dónde intervenir.

## 2. Objetivo

Garantizar que todo evento de daño, retraso, faltante o pérdida quede registrado en un ticket único,
trazable y con responsable asignado, sin importar en qué etapa del ciclo ocurra ni por qué canal se reportó
originalmente, hasta que quede formalmente resuelto.

## 3. Alcance

**Este SOP es transversal: no está ligado a una etapa específica del ciclo de vida.** Aplica a cualquier
incidencia detectada en cualquier punto — Alta, Identificación, Inspección, Limpieza, Reparación, Asignación,
Movimiento, Uso, Retorno, Reinspección, o cualquier otra etapa de la
[tabla de trazabilidad cruzada](../ciclo-de-vida-del-activo.md#2-tabla-de-trazabilidad-cruzada-por-etapa) —
y a los cuatro tipos de incidencia:

1. **Daño** — el activo presenta un deterioro no esperado, detectado en inspección, lavado, reparación o
   reporte de cliente.
2. **Retraso** — el activo no llega o no se mueve dentro del tiempo esperado, sin que se sepa aún si está
   detenido, perdido o simplemente demorado.
3. **Faltante** — una discrepancia entre el inventario esperado y el inventario real (físico o digital)
   revela que uno o más activos no aparecen.
4. **Pérdida** — reporte explícito o presunción razonable de que un activo específico no tiene ubicación
   confirmable.

Este SOP registra y da seguimiento a la incidencia como tal. No sustituye a los SOP que resuelven cada tipo
de excepción en detalle:

- Una incidencia de tipo **retraso sostenido** que excede el umbral de retención dispara
  [SOP-ACTIVO-14 · Logística inversa](./SOP-ACTIVO-14-logistica-inversa.md) (transición `Retenido → En
  tránsito`).
- Una incidencia de tipo **pérdida** se resuelve siguiendo
  [SOP-ACTIVO-16 · Activo perdido](./SOP-ACTIVO-16-activo-perdido.md) (transición `Retenido/Bloqueado →
  Perdido`).
- Una incidencia que implica **disputa de custodia o condición** se resuelve siguiendo
  [SOP-ACTIVO-17 · Activo bloqueado](./SOP-ACTIVO-17-activo-bloqueado.md) (transición `Retenido →
  Bloqueado`).
- Una incidencia de tipo **daño** que requiere intervención física deriva al proceso de inspección/reparación
  del ciclo normal (SOP-ACTIVO-03/04, SOP-ACTIVO-06).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Reportante `[VALIDAR — puede ser cualquier rol: cliente, operador de campo, inspector, logística]` | Origina el reporte de la incidencia por el canal que tenga disponible. |
| Responsable de gestión de incidencias `[VALIDAR nombre exacto del puesto]` | Registra la incidencia en el sistema con tipo, activo(s) y fecha; asigna responsable de resolución; da seguimiento hasta el cierre. |
| Responsable funcional según el tipo de incidencia `[VALIDAR: ¿logística para retraso, calidad para daño, gobernanza para pérdida/bloqueo?]` | Ejecuta o coordina la resolución de la incidencia en su ámbito, siguiendo el SOP específico que corresponda. |
| Gobernanza | Revisa incidencias recurrentes o sistemáticas por cliente/custodio/ruta y ajusta reglas operativas para prevenirlas (ver [Gobernanza](../catalogo-systems.md#15-gobernanza)). |

## 5. Diagrama de flujo

```
[Evento detectado: daño / retraso / faltante / pérdida — en cualquier etapa del ciclo]
                ↓
[Reportar la incidencia por el canal disponible
 (cliente, operador de campo, inspección, logística inversa, etc.)]
                ↓
[Registrar ticket: tipo, activo(s) involucrado(s), fecha, canal de origen]
                ↓
[Clasificar el tipo de incidencia]
                ↓
        ┌───────────┬────────────┬────────────┬───────────┐
        ↓            ↓            ↓            ↓
     [Daño]      [Retraso]    [Faltante]   [Pérdida]
        ↓            ↓            ↓            ↓
 [Derivar a    [Evaluar si   [Investigar   [Derivar a
  inspección/   excede        contra        SOP-ACTIVO-16
  reparación]   umbral de     inventario    · Activo perdido]
                retención →   digital y
                SOP-ACTIVO-14 físico]
                o SOP-ACTIVO-17
                si hay disputa]
        └───────────┴────────────┴────────────┘
                ↓
[Asignar responsable de resolución y dar seguimiento]
                ↓
        ¿Incidencia resuelta?
        ↓ sí                              ↓ no (vencido el plazo esperado)
[Registrar resolución y fecha de cierre]  [Escalar a Gobernanza]
                ↓
[Cerrar ticket — queda consultable para Analítica/Gobernanza]
```

## 6. SOP paso a paso

1. Detectar o recibir el reporte de una incidencia por cualquier canal (cliente, operador de campo,
   inspección, logística inversa, proceso interno) `[VALIDAR canales formales aceptados — hoy incluye
   WhatsApp/llamada/correo informal, a definir si se centraliza en un solo canal de entrada]`.
2. Registrar un ticket de incidencia con: tipo (daño / retraso / faltante / pérdida), activo(s) o lote
   involucrado, fecha de detección, canal de origen y reportante.
3. Clasificar el tipo de incidencia según la definición del [§3 Alcance](#3-alcance).
4. Si es **daño**, derivar el activo al proceso de inspección/reparación del ciclo normal, dejando la
   incidencia vinculada al resultado de esa reparación.
5. Si es **retraso**, evaluar si el activo ya excede el umbral de tiempo máximo fuera. Si lo excede y entra
   a estado `Retenido`, continuar según [SOP-ACTIVO-14 · Logística inversa](./SOP-ACTIVO-14-logistica-inversa.md);
   si además hay disputa de custodia o condición, derivar a
   [SOP-ACTIVO-17 · Activo bloqueado](./SOP-ACTIVO-17-activo-bloqueado.md).
6. Si es **faltante**, investigar la discrepancia contra el inventario digital y, si aplica, contra el
   inventario físico más reciente ([Inventarios físicos](../catalogo-solutions.md#11-inventarios-físicos)).
   Si la investigación no resuelve la ubicación del activo, escalar a
   [SOP-ACTIVO-16 · Activo perdido](./SOP-ACTIVO-16-activo-perdido.md).
7. Si es **pérdida** (reporte explícito o presunción razonable), derivar directamente a
   [SOP-ACTIVO-16 · Activo perdido](./SOP-ACTIVO-16-activo-perdido.md).
8. Asignar un responsable de resolución según el tipo de incidencia (ver [§4 Roles](#4-roles)) y un plazo
   esperado de resolución `[VALIDAR plazo por tipo de incidencia]`.
9. Dar seguimiento al ticket hasta que se resuelva: registrar la resolución, la evidencia asociada
   (fotos, comunicaciones, actas) y la fecha de cierre.
10. Si el plazo esperado se vence sin resolución, escalar el ticket a Gobernanza.
11. Cerrar el ticket. El ticket cerrado queda consultable como insumo de
    [Analítica](../catalogo-systems.md#14-analítica) y [Gobernanza](../catalogo-systems.md#15-gobernanza)
    para detectar patrones por cliente, custodio, ruta o tipo de activo.
12. Recordar que, según el [invariante 3 de la máquina de estados](../estados-del-activo.md#4-invariantes-reglas-que-el-sistema-debe-garantizar),
    todo activo en estado `Bloqueado` o `Perdido` debe tener una incidencia asociada registrada aquí — no
    puede existir el estado sin el ticket correspondiente.

## 7. Checklist

- [ ] Incidencia registrada como ticket único, sin importar el canal de origen (no queda solo en un chat o correo).
- [ ] Tipo de incidencia clasificado (daño / retraso / faltante / pérdida).
- [ ] Activo(s) o lote involucrado identificado en el ticket.
- [ ] Responsable de resolución asignado según el tipo de incidencia.
- [ ] Incidencia derivada al SOP correspondiente cuando aplica (SOP-ACTIVO-14, 16 o 17).
- [ ] Evidencia asociada (fotos, comunicaciones, actas) adjunta al ticket.
- [ ] Fecha de resolución y cierre registrada.
- [ ] Tickets vencidos sin resolución escalados a Gobernanza.
- [ ] Todo activo en estado `Bloqueado` o `Perdido` tiene un ticket de incidencia asociado y consultable.

## 8. KPI

- Número de incidencias por tipo (daño / retraso / faltante / pérdida), por cliente, custodio o ruta.
- Tiempo promedio de resolución por tipo de incidencia.
- % de incidencias resueltas dentro del plazo esperado vs. escaladas a Gobernanza.
- % de activos en `Bloqueado`/`Perdido` sin ticket de incidencia asociado (debe tender a 0 — ver invariante 3).
- Tasa de pérdida y tasa de daño ([ETAPA 6 del README](../README.md#etapa-6--indicadores-del-sistema)),
  alimentadas directamente por los tickets cerrados aquí.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Incidencia reportada por canal informal (WhatsApp, llamada) nunca se convierte en ticket | Registro obligatorio como ticket en el mismo día de detección, sin importar el canal de origen (paso 2) `[VALIDAR mecanismo para capturar reportes informales]`. |
| Incidencias tratadas de forma aislada sin detectar patrón sistemático con un cliente/custodio/ruta | Todo ticket cerrado queda consultable para Analítica y Gobernanza (paso 11); revisión periódica de incidencias recurrentes `[VALIDAR periodicidad de revisión]`. |
| Ticket abierto indefinidamente sin responsable ni plazo, quedando "perdido" dentro del propio sistema de incidencias | Asignación obligatoria de responsable y plazo esperado al abrir el ticket (paso 8); escalamiento automático a Gobernanza al vencer el plazo (paso 10). |
| Activo en `Bloqueado` o `Perdido` sin ticket de incidencia asociado, rompiendo el invariante 3 de la máquina de estados | Verificación cruzada entre estado del activo y existencia de ticket como parte del checklist de cierre de este SOP. |
| Clasificación incorrecta del tipo de incidencia retrasa la derivación al SOP correcto | Definiciones explícitas de los cuatro tipos en el [§3 Alcance](#3-alcance); reclasificación permitida si la investigación revela un tipo distinto al reportado inicialmente. |

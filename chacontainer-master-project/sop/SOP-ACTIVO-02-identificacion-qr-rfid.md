# SOP-ACTIVO-02: Identificación QR/RFID

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre la segunda etapa del [ciclo de vida del activo](../ciclo-de-vida-del-activo.md):
**Identificación**, la transición `Registrado → Identificado` de la [máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones)
(transición #2). Se construyó a partir del cruce entre el [catálogo Systems](../catalogo-systems.md#3-identificación)
(capacidades [QR](../catalogo-systems.md#4-qr) y [RFID](../catalogo-systems.md#5-rfid)) y el ciclo de vida —
**no a partir de una entrevista real con el responsable operativo**. No debe declararse estándar hasta
validar en operación los puntos marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Identificación es el proceso que convierte un activo ya registrado (con ficha en el registro maestro, pero
sin marcaje físico) en un activo que cualquier persona en campo puede confirmar en segundos, escaneando un
código QR o leyendo un tag RFID. Recibe un activo en estado `Registrado`, decide el método de identificación
(QR o RFID) según el caso de uso, aplica el identificador físico, lo vincula al ID del registro maestro, y
deja el activo listo para entrar a [Inspección](../ciclo-de-vida-del-activo.md#2-tabla-de-trazabilidad-cruzada-por-etapa)
(SOP-ACTIVO-03, pendiente). Es el paso descrito en el [catálogo Systems](../catalogo-systems.md#3-identificación)
como "el puente entre el mundo físico y el registro digital": sin este paso, ninguna capacidad posterior
(trazabilidad, alertas, indicadores) tiene datos reales que procesar. También es, junto con inventarios
físicos, la puerta de entrada comercial más frecuente hacia Packaging Systems descrita en la
[ETAPA 8 del README](../README.md#etapa-8--modelo-comercial-solutions--systems).

## 2. Objetivo

Garantizar que todo activo en estado `Registrado` recibe un identificador físico (QR o RFID) vinculado sin
ambigüedad a su ID en el registro maestro, de forma verificable (el identificador escanea/lee correctamente
antes de darlo por concluido), antes de avanzar a inspección.

## 3. Alcance

Aplica a todo activo retornable (contenedor, IBC, tarima, tambor) que se encuentra en estado `Registrado`
(ver [SOP-ACTIVO-01](./SOP-ACTIVO-01-alta-de-activos.md)) y que aún no tiene identificador físico vinculado,
sin importar la vía por la que llegó a ese estado (venta nueva, hallazgo operativo, incorporación de parque
existente). Incluye la decisión de método (QR vs. RFID) por caso de uso, la aplicación física del
identificador y la verificación de lectura. No incluye la inspección de condición del activo (SOP-ACTIVO-03,
pendiente), que es la etapa siguiente del ciclo, aunque ambos procesos suelen ejecutarse en la misma visita
operativa cuando el activo llega sin identificación durante un servicio de campo (ver
[§5 de ciclo-de-vida-del-activo.md](../ciclo-de-vida-del-activo.md#5-punto-de-entrada-del-cliente--inicio-del-ciclo)).

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Operador de identificación `[VALIDAR nombre exacto del puesto]` | Decide el método (QR/RFID) según el caso de uso, aplica el identificador y registra el evento. Dispara la transición `Registrado → Identificado` (fila 2 de la [tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones)). |
| Responsable de planta / operaciones `[VALIDAR]` | Define el criterio de cuándo usar QR vs. RFID por familia de activo o cliente, y resuelve excepciones. |
| Técnico de infraestructura RFID `[VALIDAR — puede ser un rol externo/proveedor]` | Instala y calibra portales, antenas o lectores fijos cuando el caso de uso RFID lo requiere (no aplica a la instalación de QR). |

## 5. Diagrama de flujo

```
[Activo en estado "Registrado" (sin identificador físico)]
                ↓
[Determinar caso de uso: lectura manual esporádica
 vs. lectura masiva/automática/alta frecuencia]
                ↓
        ¿Requiere lectura sin línea de vista o masiva?
        ↓ no                              ↓ sí
[Método: QR                         [Método: RFID
 (catálogo Systems #4)]              (catálogo Systems #5)]
        ↓                                    ↓
[Generar código QR con             [Instalar tag RFID en el activo;
 payload = ID del activo]           si aplica, verificar infraestructura
        ↓                           de lectura en el punto de control]
[Imprimir/aplicar etiqueta                  ↓
 sobre el activo]                  [Verificar lectura en el
        ↓                           punto de control]
[Verificar que el código                    ↓
 escanea correctamente]                     │
        └──────────────┬─────────────────────┘
                        ↓
        [Vincular identificador al ID del registro maestro]
                        ↓
        [Registrar evidencia: identificador instalado,
         evento de escaneo/lectura inicial, foto]
                        ↓
        [Confirmar → estado "Identificado"]
                        ↓
        [Enviar a Inspección (SOP-ACTIVO-03, pendiente)]
```

## 6. SOP paso a paso

1. Confirmar que el activo está en estado `Registrado` (ver [SOP-ACTIVO-01](./SOP-ACTIVO-01-alta-de-activos.md))
   y localizar su ID en el registro maestro antes de iniciar.
2. Determinar el caso de uso de identificación: lectura manual esporádica desde smartphone (favorece QR) vs.
   lectura sin línea de vista, masiva o de alta frecuencia en portales fijos (favorece RFID) — ver
   [catálogo Systems, QR](../catalogo-systems.md#4-qr) y [RFID](../catalogo-systems.md#5-rfid)
   `[VALIDAR criterio formal de decisión QR vs. RFID por familia de activo/cliente]`.
3. Si el método es **QR**: generar el código con el payload correspondiente (ID del activo, y opcionalmente
   tenant/tipo — ver formato en [architecture.md](../../chacontainer/docs/architecture.md#qr-code-system)
   del SaaS existente), imprimir o aplicar la etiqueta sobre el activo `[VALIDAR material y proveedor de la
   etiqueta, resistente a lavado e intemperie]`, y escanearla para verificar que lee correctamente antes de
   continuar.
4. Si el método es **RFID**: instalar el tag en el activo y, si el caso de uso incluye infraestructura de
   lectura (portal, antenas, lectores), verificar que está desplegada en los puntos de control acordados
   `[VALIDAR alcance: solo tag vs. tag + infraestructura de lectura — ver catálogo Systems #5]`. Realizar una
   lectura de prueba en el punto de control antes de continuar.
5. Vincular el identificador físico (código QR o tag RFID) al ID del activo en el registro maestro, sin
   ambigüedad — un identificador por activo, un activo por identificador.
6. Registrar evidencia de la instalación: tipo de identificador aplicado, evidencia fotográfica, y el evento
   de escaneo/lectura inicial en el sistema.
7. Confirmar la identificación: el activo pasa al estado `Identificado` en la
   [máquina de estados](../estados-del-activo.md#1-catálogo-de-estados) (transición #2), pendiente de primera
   inspección.
8. Enviar el activo al flujo de Inspección, siguiente etapa del
   [ciclo de vida](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado) (SOP-ACTIVO-03, pendiente de
   redactar según el [orden de la ETAPA 14](../README.md#etapa-14--orden-de-ejecución)).
9. Actualizar el indicador de % del parque identificado sobre el parque registrado (ver [§8 KPI](#8-kpi)).

## 7. Checklist

- [ ] Activo confirmado en estado `Registrado` antes de iniciar identificación.
- [ ] Método de identificación (QR/RFID) decidido según el caso de uso, no por defecto o por costumbre.
- [ ] Identificador físico aplicado sobre el activo (etiqueta QR o tag RFID).
- [ ] Lectura de verificación ejecutada (el identificador escanea/lee correctamente) antes de dar por
      concluida la instalación.
- [ ] Identificador vinculado sin ambigüedad al ID del activo en el registro maestro (uno a uno).
- [ ] Evidencia fotográfica de la instalación registrada.
- [ ] Evento de escaneo/lectura inicial registrado en el sistema.
- [ ] Estado del activo actualizado a `Identificado` el mismo día de la instalación.
- [ ] Activo enviado al flujo de Inspección (no queda "Identificado" sin plan de inspección).

## 8. KPI

- **% del parque identificado sobre el parque registrado** — brecha que suele descubrirse en inventarios
  físicos o lavado, ver [catálogo Systems, Identificación](../catalogo-systems.md#3-identificación).
- Costo de identificación por activo, desagregado por método (QR más bajo, RFID mayor pero justificado por
  volumen/frecuencia — ver [catálogo Systems, QR](../catalogo-systems.md#4-qr) y
  [RFID](../catalogo-systems.md#5-rfid)).
- Tiempo de despliegue por lote (activos identificados por día/semana).
- % de identificadores que fallan la verificación de lectura inicial (mide calidad del material/instalación,
  no solo el proceso).
- Cobertura de puntos de control con lectura automática (aplica solo a despliegues RFID con infraestructura).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Identificador aplicado pero no verificado — falla al primer intento de escaneo real en campo | Verificación de lectura obligatoria como parte del SOP (pasos 3-4), antes de confirmar la identificación, no después. |
| Vínculo identificador↔ID del activo incorrecto o duplicado (dos activos con el mismo QR, o un QR sin activo asociado) | Vinculación uno a uno controlada por el sistema al momento del registro (paso 5); bloquear la asignación de un identificador ya vinculado a otro activo. |
| Etiqueta QR ilegible tras exposición a lavado, intemperie o manipulación en campo | Material y proveedor de etiqueta validados para el entorno de uso real `[VALIDAR material y proveedor]`; reetiquetado se resuelve como parte de [SOP-07 Reetiquetado](../ciclo-de-vida-del-activo.md#2-tabla-de-trazabilidad-cruzada-por-etapa) del ciclo comercial, no de este SOP. |
| Decisión QR vs. RFID inconsistente entre operadores (mismo tipo de activo, distinto método sin razón documentada) | Criterio de decisión documentado por familia de activo/cliente `[VALIDAR criterio formal]`, no dejado al juicio individual del operador. |
| Activo identificado pero nunca enviado a inspección (queda en `Identificado` indefinidamente) | La identificación no se considera cerrada operativamente hasta que el activo entra al flujo de Inspección; reportar en el indicador de cobertura los activos `Identificado` sin inspeccionar. |
| Infraestructura RFID (portal/antenas) no verificada antes de dar por concluida la instalación del tag | Lectura de prueba en el punto de control obligatoria (paso 4) cuando el caso de uso incluye infraestructura fija. |

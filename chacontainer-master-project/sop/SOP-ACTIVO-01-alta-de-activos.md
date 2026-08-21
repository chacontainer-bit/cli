# SOP-ACTIVO-01: Alta de activos

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento cubre la primera etapa del [ciclo de vida del activo](../ciclo-de-vida-del-activo.md):
**Alta**, la transición `— (nuevo activo) → Registrado` de la [máquina de estados](../estados-del-activo.md#2-tabla-de-transiciones)
(transición #1). Se construyó a partir del cruce entre el [catálogo Solutions](../catalogo-solutions.md#1-venta),
el [catálogo Systems](../catalogo-systems.md#2-registro) y el ciclo de vida — **no a partir de una entrevista
real con el responsable operativo**. No debe declararse estándar hasta validar en operación los puntos
marcados `[VALIDAR]`.

---

## 1. Resumen ejecutivo

Alta es el proceso que hace que un activo retornable (contenedor, IBC, tarima o tambor) empiece a existir
para CHACONTAINER como registro gestionable, no solo como objeto físico. Recibe un activo por una de tres
vías — venta de un activo nuevo, hallazgo durante un servicio Solutions (inventario físico, lavado,
recuperación), o incorporación de un activo ya propiedad del cliente al sistema — y lo convierte en una
ficha con ID único en el registro maestro (Capa 1 · Activo), dejándolo listo para pasar a
[Identificación](./SOP-ACTIVO-02-identificacion-qr-rfid.md). Sin Alta no hay ciclo de vida que gestionar:
es la puerta de entrada obligatoria a todo lo demás documentado en [ciclo-de-vida-del-activo.md](../ciclo-de-vida-del-activo.md).

## 2. Objetivo

Garantizar que todo activo retornable que CHACONTAINER vende, recibe o detecta queda registrado en el
registro maestro con un ID único, sin duplicados, con los datos mínimos de Capa 1 completos, antes de que
avance a cualquier otra etapa del ciclo.

## 3. Alcance

Aplica a todo activo retornable (contenedor, IBC, tarima, tambor) que entra por primera vez al sistema de
CHACONTAINER, por cualquiera de estas tres vías:

1. **Venta de un activo nuevo** — servicio [Venta](../catalogo-solutions.md#1-venta) de Solutions.
2. **Hallazgo operativo** — un activo sin registro previo que aparece durante lavado, inventario físico o
   recuperación (ver [ETAPA 8 del README](../README.md#etapa-8--modelo-comercial-solutions--systems) y
   [§5 de ciclo-de-vida-del-activo.md](../ciclo-de-vida-del-activo.md#5-punto-de-entrada-del-cliente--inicio-del-ciclo)).
3. **Incorporación de parque existente del cliente** — el cliente ya es propietario del activo y pide a
   CHACONTAINER administrarlo (entrada directa a Systems sin pasar por Venta).

No incluye la [identificación física](./SOP-ACTIVO-02-identificacion-qr-rfid.md) (QR/RFID) del activo, que
es el paso siguiente y un SOP separado. No incluye la valoración comercial del activo (cotización,
condiciones de venta), que se resuelve en el proceso comercial correspondiente antes de llegar a este SOP.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Operador de venta/alta `[VALIDAR nombre exacto del puesto]` | Captura los datos del activo y ejecuta el alta en el registro maestro. Es quien dispara la transición `— → Registrado` según la [tabla de transiciones](../estados-del-activo.md#2-tabla-de-transiciones) (fila 1). |
| Responsable de inventario/registro `[VALIDAR]` | Verifica que no exista un duplicado antes de confirmar el alta y resuelve discrepancias. |
| Comercial `[VALIDAR]` | Origina el alta cuando proviene de una venta; entrega al operador de alta los datos de la operación. |
| Operador de campo (lavado, inventario físico, recuperación) | Origina el alta cuando el activo se detecta sin registro previo durante un servicio distinto (ver [Lectura cruzada del catálogo Solutions](../catalogo-solutions.md#lectura-cruzada-de-dónde-entra-cada-servicio-a-systems)). |

## 5. Diagrama de flujo

```
[Activo entra al sistema por: Venta / Hallazgo operativo / Incorporación de parque existente]
                ↓
[Verificar si el activo ya tiene un registro previo]
                ↓
        ¿Ya existe en el registro maestro?
        ↓ sí                              ↓ no
[Detener alta — usar registro       [Capturar datos de Capa 1:
 existente, reportar duplicado       tipo, modelo, medida, propietario,
 potencial]                          número de parte, condición, valor]
                                              ↓
                                     [Asignar ID único en el registro maestro]
                                              ↓
                                     [Registrar origen del alta
                                      (venta / hallazgo / incorporación) y fecha]
                                              ↓
                                     [Confirmar alta → estado "Registrado"]
                                              ↓
                                     [Enviar a Identificación (SOP-ACTIVO-02)]
```

## 6. SOP paso a paso

1. Identificar la vía de entrada del activo: venta nueva, hallazgo durante un servicio Solutions, o
   incorporación de parque existente del cliente (ver [§3 Alcance](#3-alcance)).
2. Antes de crear un registro nuevo, verificar contra el registro maestro si el activo ya existe (por
   número de serie/SKU, o por coincidencia de tipo + propietario + ubicación cuando no hay identificador
   previo) `[VALIDAR criterio exacto de búsqueda de duplicados]`.
3. Si el activo ya tiene registro, detener el alta: usar el registro existente y reportar la posible
   duplicidad al responsable de inventario/registro en vez de crear una ficha nueva.
4. Si el activo no tiene registro, capturar los datos mínimos de **Capa 1 · Activo** (ver
   [README ETAPA 3](../README.md#etapa-3--modelo-de-packaging-systems)): tipo, modelo, medida, propietario,
   número de parte, condición inicial, valor `[VALIDAR campos obligatorios vs. opcionales por tipo de activo]`.
5. Asignar un ID único al activo en el registro maestro (Módulo 1 de
   [CHACONTAINER OS](../catalogo-systems.md#16-chacontainer-os)), evitando duplicados — ejecución de la
   capacidad [Registro](../catalogo-systems.md#2-registro) de Systems.
6. Registrar el origen del alta (venta, hallazgo en campo, incorporación de parque existente) y la fecha —
   este dato queda como evidencia permanente de la ficha del activo.
7. Confirmar el alta: el activo pasa al estado `Registrado` en la [máquina de estados](../estados-del-activo.md#1-catálogo-de-estados)
   (transición #1), sin identificación física todavía.
8. Enviar el activo al flujo de [Identificación](./SOP-ACTIVO-02-identificacion-qr-rfid.md), siguiente etapa
   del [ciclo de vida](../ciclo-de-vida-del-activo.md#1-el-flujo-principal-anotado).
9. Actualizar el indicador de cobertura de registro (ver [§8 KPI](#8-kpi)).

## 7. Checklist

- [ ] Vía de entrada del activo identificada (venta / hallazgo / incorporación de parque existente).
- [ ] Búsqueda de duplicados ejecutada antes de crear el registro.
- [ ] Datos mínimos de Capa 1 capturados (tipo, modelo, medida, propietario, número de parte, condición, valor).
- [ ] ID único asignado en el registro maestro, sin colisión con un ID existente.
- [ ] Origen del alta (venta / hallazgo / incorporación) y fecha registrados en la ficha del activo.
- [ ] Estado del activo actualizado a `Registrado` en el sistema el mismo día del alta.
- [ ] Activo enviado al flujo de Identificación (no queda "Registrado" indefinidamente sin plan de identificación).

## 8. KPI

- **Cobertura de registro**: % del parque real (estimado en [diagnóstico de activos](../catalogo-systems.md#1-diagnóstico-de-activos))
  que efectivamente existe en el registro maestro — ver [catálogo Systems, Registro](../catalogo-systems.md#2-registro).
- Tiempo entre la entrada del activo al sistema y la confirmación del alta.
- % de altas detenidas por duplicado potencial detectado (mide efectividad del control de duplicados).
- Activos nuevos incorporados al parque total administrable por periodo (ver
  [catálogo Solutions, Venta](../catalogo-solutions.md#1-venta)).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Activo dado de alta por duplicado (ya existía en el registro maestro) | Verificación obligatoria de duplicados antes de confirmar el alta (paso 2-3); criterio de búsqueda documentado y no dejado al juicio individual `[VALIDAR criterio]`. |
| Datos de Capa 1 incompletos o inconsistentes al momento del alta | Definir por escrito los campos obligatorios por tipo de activo; no permitir confirmar el alta sin ellos. |
| Activo entra por hallazgo operativo (lavado, inventario) y nunca se formaliza el alta porque "ya está circulando" | El operador de campo que detecta el activo sin registro está obligado a iniciar este SOP, no solo a continuar su propio servicio — reforzado por la [lectura cruzada del catálogo Solutions](../catalogo-solutions.md#lectura-cruzada-de-dónde-entra-cada-servicio-a-systems). |
| Activo registrado pero nunca enviado a Identificación (queda en `Registrado` indefinidamente) | El alta no se considera cerrada operativamente hasta que el activo entra al flujo de [SOP-ACTIVO-02](./SOP-ACTIVO-02-identificacion-qr-rfid.md); reportar en el indicador de cobertura los activos `Registrado` sin identificar. |
| Falta de responsable claro para resolver discrepancias de duplicados | Rol de responsable de inventario/registro definido explícitamente (ver [§4 Roles](#4-roles)) `[VALIDAR nombre exacto del puesto]`. |

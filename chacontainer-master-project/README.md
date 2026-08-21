# PROYECTO MAESTRO CHACONTAINER

## De proveedor operativo a operador del ciclo de vida del empaque retornable

> Este proyecto es independiente de [`chacontainer/`](../chacontainer/) (el SaaS actual). Aquí vive el
> documento estratégico maestro que define hacia dónde evoluciona el negocio CHACONTAINER como
> operador del ciclo de vida del empaque retornable — no es código ni documentación de producto
> del SaaS existente.

**Entregables de la FASE I · Fundación:**

- ✅ [Catálogo definitivo Packaging Solutions](./catalogo-solutions.md) — paso 1.
- ✅ [Catálogo definitivo Packaging Systems](./catalogo-systems.md) — paso 2.
- ✅ [Ciclo de vida del activo](./ciclo-de-vida-del-activo.md) — paso 3.
- ✅ [Estados del activo](./estados-del-activo.md) — paso 4.
- ✅ [SOP del ciclo del activo](./sop/README.md) (20 SOP-ACTIVO) — paso 5.
- ✅ [KPI del sistema](./kpi.md) — paso 6.

**Entregables de la FASE II · Sistema:**

- ✅ [Modelo de gobernanza](./modelo-de-gobernanza.md) — paso 7.
- ✅ [Modelo de datos](./modelo-de-datos.md) — paso 8.
- ✅ [Reglas operativas](./reglas-operativas.md) — paso 9.
- ✅ [Arquitectura de CHACONTAINER OS](./arquitectura-os.md) — paso 10.

**Entregables de la FASE III · Producto:**

- ✅ [MVP — alcance](./mvp.md) — paso 11.
- ✅ [QR — el escaneo como disparador del ciclo](./qr.md) — paso 12.
- ✅ [Dashboard — indicadores y vistas](./dashboard.md) — paso 13.
- ✅ [Alertas — de la regla a la acción](./alertas.md) — paso 14.
- ✅ [Piloto — plan de ejecución](./piloto.md) — paso 15.

**Entregables de la FASE IV · Comercial:**

- ✅ [Oferta Packaging Systems](./oferta-systems.md) — paso 16.
- ✅ [Pricing — modelo, no tarifas](./pricing.md) — paso 17.
- ✅ [Diagnóstico comercial](./diagnostico-comercial.md) — paso 18.
- ✅ [Piloto pagado — estructura comercial](./piloto-pagado.md) — paso 19.
- ✅ [Conversión Solutions → Systems](./conversion-solutions-systems.md) — paso 20.

## 1. Tesis estratégica

CHACONTAINER se estructura alrededor de una sola idea:

> «Gestionar el ciclo completo del activo retornable, desde su incorporación al sistema hasta su recuperación, reutilización o disposición final.»

La empresa se divide en dos unidades complementarias:

- **Packaging Solutions** — Ejecuta físicamente sobre el activo.
- **Packaging Systems** — Administra, conecta y gobierna el sistema donde ese activo opera.

La secuencia estratégica es:

**Solutions → Datos → Visibilidad → Control → Systems → Recurrencia**

El objetivo no es eliminar los servicios transaccionales. Es utilizarlos como puntos de entrada para construir relaciones operativas de largo plazo.

---

## ETAPA 0 · Fundamento del modelo

### Objetivo

Definir exactamente qué es CHACONTAINER, qué vende cada unidad y dónde termina la responsabilidad de cada una.

### Arquitectura

**CHACONTAINER Packaging Solutions**

- Venta
- Renta
- Lavado
- Reparación
- Reacondicionamiento
- Modificación
- Dunnage
- Reetiquetado
- Inspección
- Clasificación
- Inventarios físicos
- Recuperación
- Compra de scrap
- Disposición final
- Recuperación de valor

**CHACONTAINER Packaging Systems**

- Diagnóstico de activos
- Registro
- Identificación
- QR
- RFID
- Inventario digital
- Trazabilidad
- Gestión de custodios
- Logística inversa
- Gestión de incidencias
- Reglas operativas
- Alertas
- Indicadores
- Analítica
- Gobernanza
- CHACONTAINER OS

### Entregable

Un catálogo maestro donde cada servicio tenga:

1. Qué problema resuelve.
2. Qué recibe CHACONTAINER.
3. Qué actividad ejecuta.
4. Qué evidencia genera.
5. Qué indicador afecta.
6. Qué información produce.
7. Qué servicio puede venderse después.

---

## ETAPA 1 · Normalizar Packaging Solutions

### Objetivo

Convertir los servicios actuales en procesos repetibles y medibles.

Cada servicio debe dejar de depender solamente de experiencia operativa y convertirse en un procedimiento estandarizado.

### Ejemplo: Lavado

**Entrada:** Activo sucio.

**Proceso:** Inspección → retiro de residuos → retiro de etiquetas → prelavado → lavado → secado → inspección final.

**Salida:** Activo liberado.

**Datos generados:**

- Número de activo
- Tipo
- Fecha
- Ubicación
- Condición inicial
- Condición final
- Daños encontrados
- Trabajo ejecutado
- Responsable
- Evidencia

El mismo modelo debe construirse para: Reparación, Inspección, Scrap, Renta, Recuperación, Modificación, Reetiquetado, Inventarios.

### Resultado

Solutions empieza a producir datos además de ejecutar servicios. Ese es el primer puente hacia Packaging Systems.

---

## ETAPA 2 · SOP del ciclo completo del activo

### Objetivo

Construir el proceso operativo central de CHACONTAINER.

### Ciclo maestro

```
Alta
  ↓
Identificación
  ↓
Inspección
  ↓
Limpieza
  ↓
Reparación / reacondicionamiento
  ↓
Reetiquetado / reconfiguración
  ↓
Liberación
  ↓
Inventario disponible
  ↓
Asignación
  ↓
Salida
  ↓
Uso
  ↓
Movimiento
  ↓
Trazabilidad
  ↓
Retorno
  ↓
Reinspección
  ↓
Mantenimiento
  ↓
Reincorporación
  o
Baja / scrap / recuperación
```

### SOP que deberán existir

| SOP | Nombre |
|---|---|
| SOP-01 | Alta de activos |
| SOP-02 | Identificación QR/RFID |
| SOP-03 | Inspección |
| SOP-04 | Clasificación de condición |
| SOP-05 | Lavado |
| SOP-06 | Reparación |
| SOP-07 | Reetiquetado |
| SOP-08 | Liberación |
| SOP-09 | Inventario |
| SOP-10 | Asignación |
| SOP-11 | Movimiento |
| SOP-12 | Transferencia de custodia |
| SOP-13 | Retorno |
| SOP-14 | Logística inversa |
| SOP-15 | Incidencias |
| SOP-16 | Activo perdido |
| SOP-17 | Activo bloqueado |
| SOP-18 | Baja |
| SOP-19 | Scrap |
| SOP-20 | Disposición final |

---

## ETAPA 3 · Modelo de Packaging Systems

Aquí comienza la transformación del servicio físico en gestión.

### Cinco capas

**Capa 1 · Activo** — Qué existe.
ID, Tipo, Modelo, Medida, Propietario, Número de parte, Condición, Valor.

**Capa 2 · Ubicación** — Dónde está.
Planta, Almacén, Proveedor, Cliente, Línea, Transporte, Hub.

**Capa 3 · Custodia** — Quién responde por él.
Cliente, Proveedor, Operador, Transportista, Planta, Área, Usuario.

**Capa 4 · Estado** — Qué está ocurriendo.
Disponible, Asignado, En tránsito, En uso, Retenido, Sucio, Dañado, Reparación, Perdido, Scrap.

**Capa 5 · Gobierno** — Qué debe ocurrir.
Tiempo máximo fuera, Punto de retorno, Responsable, Nivel mínimo de inventario, Condición permitida, Alertas, Escalaciones, Evidencias.

---

## ETAPA 4 · Modelo de gobernanza

El sistema necesita reglas, no solamente localización.

### Preguntas que debe poder responder

- ¿Cuántos activos existen?
- ¿Cuántos están realmente disponibles?
- ¿Dónde están?
- ¿Quién los tiene?
- ¿Cuánto tiempo llevan ahí?
- ¿Cuál es su condición?
- ¿Cuándo deberían regresar?
- ¿Cuántos están detenidos?
- ¿Cuántos están dañados?
- ¿Cuántos faltan?
- ¿Cuánto cuesta cada ciclo?
- ¿Dónde se producen las pérdidas?
- ¿Qué proveedor retiene activos?
- ¿Qué planta tiene exceso?
- ¿Qué planta tiene déficit?
- ¿Qué activos conviene reparar?
- ¿Qué activos deben darse de baja?

Esto convierte el tracking en gobernanza.

---

## ETAPA 5 · CHACONTAINER OS

### Objetivo

Digitalizar el modelo operativo anterior. El software no debe diseñarse primero: debe digitalizar procesos que ya fueron definidos.

### MVP

| Módulo | Función |
|---|---|
| 1 · Activos | Registro maestro |
| 2 · Identificación | QR / RFID |
| 3 · Ubicaciones | Plantas, almacenes, proveedores y clientes |
| 4 · Movimientos | Entrada, salida y transferencias |
| 5 · Custodia | Responsabilidad del activo |
| 6 · Condición | Estado físico |
| 7 · Mantenimiento | Lavado y reparación |
| 8 · Incidencias | Daños, retrasos, faltantes y pérdidas |
| 9 · Inventario | Disponible vs. requerido |
| 10 · Alertas | Excepciones |
| 11 · Dashboard | Indicadores |
| 12 · Evidencia | Fotos, inspecciones, registros y documentos |

---

## ETAPA 6 · Indicadores del sistema

CHACONTAINER Systems debe demostrar económicamente por qué existe.

### KPI principales

| KPI | Fórmula |
|---|---|
| Disponibilidad | Activos disponibles / activos requeridos × 100 |
| Utilización | Activos en uso / activos disponibles × 100 |
| Tiempo de ciclo | Tiempo desde salida hasta retorno |
| Cumplimiento de retorno | Retornos en tiempo / retornos esperados × 100 |
| Pérdida | Activos no recuperados / activos administrados × 100 |
| Daño | Activos dañados / activos retornados × 100 |
| Tiempo detenido | Horas o días sin movimiento |
| Costo por ciclo | Costo total del sistema / ciclos completados |
| Rotación | Número de ciclos por activo |
| Disponibilidad real | Inventario físicamente utilizable, no solamente inventario contable |

---

## ETAPA 7 · Piloto Packaging Systems

No debe venderse inicialmente como una transformación gigantesca. Debe probarse.

### Piloto

Seleccionar:

- Una planta.
- Un proveedor o cliente.
- Una familia de contenedores.
- 200–500 activos.
- Una ruta.
- Un número de parte.
- Un problema concreto.

**Duración:** 60–90 días.

**Medición inicial** (antes de intervenir): Inventario, Pérdidas, Tiempos, Disponibilidad, Daños, Retrasos, Costos.

**Medición posterior:** Comparar exactamente los mismos indicadores.

---

## ETAPA 8 · Modelo comercial Solutions → Systems

Aquí debe producirse la principal ventaja comercial. No todos los clientes tienen que entrar comprando Systems.

### Entrada

```
Cliente solicita: lavado
  ↓
CHACONTAINER detecta activos sin identificación
  ↓
Se propone identificación
  ↓
Se genera inventario
  ↓
Se descubren faltantes
  ↓
Se necesita trazabilidad
  ↓
Se detectan activos detenidos
  ↓
Se implementan alertas
  ↓
Se necesita logística inversa
  ↓
CHACONTAINER termina gobernando el sistema
```

La misma lógica funciona desde: Reparación, Scrap, Renta, Venta, Inventarios, Recuperación.

---

## ETAPA 9 · Escalera comercial

| Nivel | Descripción |
|---|---|
| 1 · Transacción | Venta / renta / lavado / reparación / scrap |
| 2 · Servicio recurrente | Lavado, reparación, recuperación o inspección periódica |
| 3 · Gestión | Inventario + identificación + mantenimiento |
| 4 · Control | Trazabilidad + logística inversa + incidencias |
| 5 · Gobernanza | CHACONTAINER administra el sistema ERT |
| 6 · Plataforma | Cliente utiliza CHACONTAINER OS como infraestructura del sistema |

---

## ETAPA 10 · Modelo de ingresos

La arquitectura permite combinar varias fuentes.

- **Transaccionales** — Venta, Scrap, Reparación, Lavado, Modificación.
- **Uso** — Renta, Operación logística, Recuperación.
- **Recurrentes** — Administración de activos, Inventario, Tracking, Gestión de retornos, Mantenimiento.
- **Software** — Licencia, Activo/mes, Usuario, Planta, Módulos.
- **Gestión integral** — Fee mensual por gobernanza del sistema.

---

## ETAPA 11 · Propuesta comercial

La conversación con el cliente debe cambiar.

No iniciar con: «Vendemos contenedores.»

Sino con el problema: **¿Cuántos activos tienen realmente disponibles hoy y cuánto les cuesta no saber dónde están los demás?**

Posteriormente CHACONTAINER puede intervenir en cualquiera de los puntos donde exista fricción.

---

## ETAPA 12 · Posicionamiento

**Posicionamiento corporativo**

> CHACONTAINER gestiona el ciclo de vida del empaque retornable industrial, desde su suministro, mantenimiento y recuperación hasta su trazabilidad, disponibilidad y gobernanza.

- **Packaging Solutions** — Intervenimos físicamente el activo.
- **Packaging Systems** — Gobernamos el sistema donde ese activo opera.
- **CHACONTAINER OS** — Conecta activos, movimientos, custodios, mantenimiento, incidencias y decisiones en una sola plataforma.

---

## ETAPA 13 · Arquitectura empresarial

```
CHACONTAINER
├── Packaging Solutions
│   ├── Venta
│   ├── Renta
│   ├── Lavado
│   ├── Reparación
│   ├── Modificación
│   ├── Scrap
│   ├── Recuperación
│   └── Inventarios
│
├── Packaging Systems
│   ├── Diagnóstico
│   ├── Identificación
│   ├── Tracking
│   ├── Inventario
│   ├── Reverse Logistics
│   ├── Gobernanza
│   ├── Analítica
│   └── Gestión del ciclo de vida
│
└── CHACONTAINER OS
    ├── Asset Registry
    ├── Tracking
    ├── Maintenance
    ├── Inventory
    ├── Incidents
    ├── Rules
    ├── Analytics
    └── Integrations
```

---

## ETAPA 14 · Orden de ejecución

No conviene intentar construir todo simultáneamente.

**FASE I · Fundación**

1. Catálogo definitivo Solutions.
2. Catálogo definitivo Systems.
3. Ciclo de vida.
4. Estados del activo.
5. SOP.
6. KPI.

**FASE II · Sistema**

7. Modelo de gobernanza.
8. Modelo de datos.
9. Reglas.
10. Arquitectura OS.

**FASE III · Producto**

11. MVP.
12. QR.
13. Dashboard.
14. Alertas.
15. Piloto.

**FASE IV · Comercial**

16. Oferta Systems.
17. Pricing.
18. Diagnóstico comercial.
19. Piloto pagado.
20. Conversión Solutions → Systems.

**FASE V · Escala**

21. Hubs.
22. Partners.
23. RFID.
24. Integraciones ERP/MES.
25. Multi-planta.
26. Operación nacional.
27. Operación México–Estados Unidos.

---

## Norte del proyecto

CHACONTAINER no tiene que abandonar su negocio actual para transformarse. Debe utilizarlo como infraestructura de entrada.

Solutions genera interacción con el activo. La interacción genera datos. Los datos generan visibilidad. La visibilidad permite control. El control permite gobernanza. La gobernanza genera recurrencia.

Ese es el proyecto empresarial completo:

> **CHACONTAINER**
> Operating System for Returnable Packaging
> Physical Operations + Asset Intelligence + Governance

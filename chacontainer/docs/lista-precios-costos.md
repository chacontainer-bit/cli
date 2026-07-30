# Lista de precios y costos base — CHACONTAINER

**Estado: Plantilla vacía — pendiente de llenar con datos reales.**
Este documento es la fuente única de verdad para cotizaciones (paso 9 del proceso comercial ideal, [SOP-07](./sop/SOP-07-cotizacion-tecnica-y-economica.md)). Mientras los campos digan `[VALIDAR]`, ninguna cotización generada a partir de este documento debe tratarse como definitiva — solo como borrador para revisión de Humberto.

**Regla de uso para cualquier agente (Claude u otro) que genere cotizaciones:**
No inventar precios, tiempos ni condiciones que no estén aquí. Si un dato falta o dice `[VALIDAR]`, marcarlo explícitamente en la cotización generada y pedirlo al responsable antes de enviarla al cliente.

---

## 1. Servicios

| Servicio | Unidad | Precio base | Costo base | Notas |
|---|---|---|---|---|
| Lavado de IBC | por unidad | `[VALIDAR]` | `[VALIDAR]` | Ver [SOP-02](./sop/SOP-02-lavado-de-activos-retornables.md) |
| Lavado de Tambor | por unidad | `[VALIDAR]` | `[VALIDAR]` | |
| Reparación de IBC | por unidad | `[VALIDAR]` | `[VALIDAR]` | Depende del daño — definir rango o tarifa por tipo de falla `[VALIDAR]` |
| Renta de activo retornable | por unidad / periodo | `[VALIDAR]` | `[VALIDAR]` | Definir periodo base (día, semana, mes) `[VALIDAR]` |
| Trazabilidad (QR + plataforma) | por unidad o por contrato | `[VALIDAR]` | `[VALIDAR]` | ¿Se cobra por activo, por proyecto, o incluido en el servicio principal? `[VALIDAR]` |
| Packaging Systems (diseño de sistema de empaque) | por proyecto | `[VALIDAR]` | `[VALIDAR]` | Muy variable — definir si se cotiza siempre a medida `[VALIDAR]` |

## 2. Venta de activos

| Activo | SKU | Precio | Costo | Pedido mínimo |
|---|---|---|---|---|
| IBC 1000L | `[VALIDAR]` | `[VALIDAR]` | `[VALIDAR]` | `[VALIDAR]` |
| IBC 600L | `[VALIDAR]` | `[VALIDAR]` | `[VALIDAR]` | `[VALIDAR]` |
| Tambor 200L | `[VALIDAR]` | `[VALIDAR]` | `[VALIDAR]` | `[VALIDAR]` |
| Pallet reutilizable | `[VALIDAR]` | `[VALIDAR]` | `[VALIDAR]` | `[VALIDAR]` |

*(Agregar más filas según el catálogo real.)*

## 3. Tiempos estándar

| Servicio | Tiempo de entrega/implementación estándar |
|---|---|
| Lavado | `[VALIDAR]` |
| Reparación | `[VALIDAR]` |
| Instalación de sistema de trazabilidad | `[VALIDAR]` |
| Implementación de Packaging Systems | `[VALIDAR]` |

## 4. Exclusiones estándar

Qué NO incluye una cotización por defecto, salvo que se indique explícitamente lo contrario:
- `[VALIDAR]` (ej. transporte, instalación, capacitación, mantenimiento posterior)

## 5. Condiciones comerciales estándar

| Condición | Valor estándar |
|---|---|
| Forma de pago | `[VALIDAR]` |
| Plazo de pago | `[VALIDAR]` |
| Garantía | `[VALIDAR]` |
| Vigencia de la oferta | `[VALIDAR]` |

## 6. Rango de aprobación

- Descuento máximo que un ejecutivo comercial puede ofrecer sin aprobación: `[VALIDAR]`
- Quién aprueba condiciones fuera de rango: `[VALIDAR]` (ver SOP-07, rol "Responsable financiero/fundador")

## 7. Cómo llenar este documento

1. Reemplazar cada `[VALIDAR]` con el dato real (precio, costo, tiempo o condición).
2. Si un servicio o activo no aplica, eliminar la fila en vez de dejarla vacía.
3. Agregar filas para cualquier servicio o activo real que falte en las tablas de ejemplo.
4. Una vez completo, actualizar el "Estado" al inicio de este documento a "Vigente" con fecha y versión.
5. Mantener una sola versión oficial — no crear copias paralelas de esta lista.

## 8. Versión

- Versión: 0.1 (plantilla vacía)
- Fecha: 2026-07-30
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: Humberto llena los datos reales (directamente o vía Claude Code local leyendo sus documentos existentes de precios/costos).

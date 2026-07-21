# SOP-04: Calificación como proveedor

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Prompt maestro del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) al paso 6 del proceso comercial ideal de CHACONTAINER: cumplir los requisitos legales, técnicos y comerciales que el cliente exige para homologar a CHACONTAINER como proveedor, después del diagnóstico consultivo (paso 5) y antes o en paralelo al levantamiento técnico (paso 7). Se construyó a partir de prácticas estándar de homologación de proveedores B2B industrial, **no a partir de una entrevista real con quien hoy gestiona estos trámites en CHACONTAINER** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6 de la metodología). Los puntos marcados como `[VALIDAR]` requieren confirmación directa del responsable administrativo/comercial.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Completar exitosamente el proceso de homologación como proveedor ante el cliente, cumpliendo sus requisitos legales, técnicos y comerciales, para poder facturar y operar formalmente con él.
2. **¿Qué problema busca resolver?** Que un negocio ya diagnosticado y en negociación se caiga o se retrase por no tener la documentación lista a tiempo; reprocesos por enviar documentos incompletos o vencidos; dependencia de que el fundador reúna manualmente cada documento cada vez que un cliente lo exige.
3. **¿Quién es el cliente?** Interno: comercial (necesita cerrar el trámite para poder facturar) y administración (que reúne y entrega los documentos). Externo: el área de compras/proveedores del cliente, que aprueba o rechaza la homologación.
4. **¿Cuáles son las entradas?** El checklist o formulario oficial de requisitos de proveedor del cliente, documentos legales de CHACONTAINER (constitución, identificación fiscal, poderes, registro en cámara de comercio), documentos técnicos (certificaciones, fichas técnicas, políticas de calidad y seguridad), documentos comerciales (referencias, estados financieros, condiciones de pago), y el diagnóstico/cotización previos que motivan el alta.
5. **¿Cuáles son las salidas esperadas?** CHACONTAINER dado de alta como proveedor aprobado en el sistema del cliente, con código o número de proveedor asignado, listo para facturar y para avanzar al levantamiento técnico (paso 7) y a la cotización formal (paso 9).
6. **¿Quién es el responsable?** `[VALIDAR]` — administración o el mismo ejecutivo comercial, según el tamaño y la exigencia del cliente.
7. **¿Qué indicadores dicen que el proceso funciona?** Tiempo desde la solicitud del cliente hasta la aprobación como proveedor, % de solicitudes aprobadas en el primer intento (sin devoluciones por documentos faltantes o vencidos), número de negocios retrasados o perdidos por fallas en este trámite.
8. **¿Qué riesgos existen?** Documentos vencidos o desactualizados (certificados de vigencia, pólizas); depender de que el fundador tenga o recuerde dónde está cada documento; requisitos distintos por cada cliente que generan reprocesos; pérdida de la oportunidad comercial por demora en el trámite.
9. **¿Qué actividades no agregan valor?** Buscar desde cero los mismos documentos cada vez que un cliente nuevo lo pide; enviar la documentación por partes en vez de un paquete completo; no dar seguimiento proactivo al estatus del trámite.
10. **¿Qué partes pueden estandarizarse?** Un kit maestro de documentos legales, técnicos y comerciales siempre actualizado; un checklist de requisitos típicos por tipo de cliente (privado, gobierno, gran corporativo).
11. **¿Qué partes pueden automatizarse?** Alertas de vencimiento de documentos (pólizas, certificados, vigencias legales), repositorio centralizado del kit de documentos, y recordatorios de seguimiento del estatus del trámite.
12. **¿Qué depende únicamente del fundador?** `[VALIDAR]` — posiblemente la firma/autorización de ciertos documentos legales, o ser quien hoy sabe cuál es la versión vigente de cada documento.

---

## 1. Resumen ejecutivo

La calificación como proveedor es el trámite mediante el cual CHACONTAINER cumple los requisitos legales, técnicos y comerciales que el área de compras del cliente exige para aprobarlo como proveedor homologado. No es un paso técnico ni de diseño de solución: es administrativo, pero puede frenar o hacer perder un negocio ya diagnosticado si no se ejecuta con un kit de documentos siempre listo y un seguimiento activo del trámite.

## 2. Objetivo

Obtener la aprobación de CHACONTAINER como proveedor homologado ante el cliente, en el menor tiempo posible y sin reprocesos, manteniendo siempre disponible y vigente el kit de documentos legales, técnicos y comerciales requeridos.

## 3. Alcance

Aplica desde que el cliente solicita el proceso de homologación (durante o después del diagnóstico consultivo, paso 5) hasta que CHACONTAINER queda registrado como proveedor aprobado en el sistema del cliente. No incluye el levantamiento técnico (paso 7) ni la cotización formal (paso 9), aunque puede correr en paralelo a ambos.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Responsable administrativo `[VALIDAR]` | Mantiene el kit maestro de documentos vigente y arma el paquete a enviar. |
| Ejecutivo comercial | Solicita al cliente su checklist de requisitos y da seguimiento al estatus del trámite. |
| Fundador (rol transitorio) | Aporta hoy la firma/autorización de ciertos documentos y el conocimiento de dónde están; el objetivo del proceso es que el kit maestro no dependa de su memoria. |

## 5. Diagrama de flujo

```
[Cliente solicita homologación como proveedor]
                ↓
[Solicitar checklist/formulario oficial de requisitos del cliente]
                ↓
[Revisar checklist contra el kit maestro de documentos]
                ↓
[Actualizar/generar documentos faltantes o vencidos]
                ↓
[Completar formulario o portal de proveedores del cliente]
                ↓
[Enviar paquete de documentación completo]
                ↓
        ¿Aprobado?
        ↓ sí                          ↓ no / falta algo
[Alta como proveedor,           [Corregir/completar
 código asignado]                 y reenviar]
        ↓
[Entrega a Levantamiento técnico (paso 7) /
 Cotización formal (paso 9)]
```

## 6. SOP paso a paso

1. Confirmar si el cliente exige un proceso formal de homologación como proveedor (no todos los clientes lo requieren).
2. Solicitar al cliente su checklist o formulario oficial de requisitos de proveedor.
3. Revisar el checklist contra el kit maestro de documentos de CHACONTAINER (legal, técnico, comercial) y detectar qué falta o está vencido.
4. Actualizar o generar los documentos faltantes o vencidos antes de enviar nada.
5. Completar el formulario o portal de proveedores del cliente con la información solicitada.
6. Enviar el paquete completo de documentación al área de compras/proveedores del cliente, no por partes.
7. Dar seguimiento activo al estatus del trámite hasta obtener aprobación o retroalimentación.
8. Si el cliente solicita correcciones, resolverlas y reenviar sin demora.
9. Una vez aprobado, registrar el código o número de proveedor asignado y la fecha de alta.
10. Entregar el registro a Levantamiento técnico (paso 7) y/o Cotización formal (paso 9), según el flujo del cliente.

## 7. Checklist

- [ ] Checklist u formulario oficial del cliente solicitado y revisado.
- [ ] Kit maestro de documentos legales, técnicos y comerciales verificado (nada vencido).
- [ ] Documentos faltantes identificados y gestionados antes de enviar.
- [ ] Formulario o portal del cliente completado sin campos pendientes.
- [ ] Paquete de documentación enviado completo, no por partes.
- [ ] Seguimiento del estatus del trámite registrado (no enviado y olvidado).
- [ ] Código/número de proveedor y fecha de alta registrados al aprobar.

## 8. KPI

- Tiempo desde la solicitud del cliente hasta la aprobación como proveedor.
- % de solicitudes de homologación aprobadas en el primer intento (sin devoluciones).
- Número de negocios retrasados o perdidos por fallas en este trámite.
- % de documentos del kit maestro vigentes en todo momento (no vencidos).

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Documentos vencidos o desactualizados | Repositorio centralizado del kit maestro con alertas de vencimiento. |
| Depender de que el fundador tenga o recuerde los documentos | Kit maestro accesible para el responsable administrativo, no solo el fundador. |
| Requisitos distintos por cliente generan reprocesos | Checklist adaptable por tipo de cliente (privado/gobierno/corporativo). |
| Demora en el trámite hace perder la oportunidad comercial | Seguimiento activo y registrado del estatus, no pasivo. |

## 10. Mejoras

- Crear y mantener un kit maestro de documentos legales, técnicos y comerciales siempre vigente, en vez de reunirlos desde cero cada vez que un cliente lo pide.
- Documentar un checklist de requisitos típicos por tipo de cliente, para anticipar qué se va a solicitar.
- Definir un responsable administrativo (no el fundador) como dueño de mantener el kit actualizado.

## 11. Recomendaciones de automatización

- Alertas automáticas de vencimiento de documentos (pólizas, certificados, vigencias legales).
- Repositorio centralizado (carpeta compartida o sistema) con el kit de documentos siempre disponible para quien lo necesite.
- Recordatorios automáticos de seguimiento al estatus del trámite con el cliente.
- **No automatizar la generación de los documentos legales en sí** (requieren revisión humana/legal); automatizar solo su organización, control de vigencia y seguimiento.

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar a quien hoy gestiona estos trámites (administración y/o el fundador) para validar el kit real de documentos, y probar el proceso con el próximo cliente que exija homologación antes de declararlo estándar.

# SOP-01: Prospección inversa por LinkedIn y Google

**Estado: Borrador v0.1 — pendiente de validación.**
Este documento aplica el Marco de decisión y el Formato estándar de salida del [Agente Ingeniero de Procesos IA](../agente-ingeniero-procesos-ia.md) al primer paso del proceso comercial ideal (Prospección inversa). Se construyó a partir de prácticas estándar de prospección B2B industrial, **no a partir de una entrevista real con el responsable comercial de CHACONTAINER** (Paso 1 de la metodología pendiente). No debe declararse estándar hasta completar la captura real y la validación en operación (Paso 6). Los puntos marcados como `[VALIDAR]` son los que requieren confirmación directa del responsable comercial.

---

## Marco de decisión aplicado

1. **¿Cuál es el objetivo del proceso?** Generar un flujo constante y calificado de prospectos de industrias objetivo que usan o podrían usar empaque retornable, para alimentar la etapa de calificación inicial (paso 3 del proceso comercial), combinando LinkedIn (contactos y puestos de decisión) y Google (descubrimiento de empresas y directorios industriales).
2. **¿Qué problema busca resolver?** Que la venta dependa de que el cliente "toque la puerta" o de la red personal del fundador, en lugar de un flujo sistemático y repetible de prospectos desde múltiples fuentes.
3. **¿Quién es el cliente?** Interno: el equipo comercial (ejecutivo/SDR) que recibe la lista de prospectos. Externo: empresas industriales/manufactureras con necesidad potencial de empaque retornable.
4. **¿Cuáles son las entradas?** Perfil de cliente ideal (ICP) `[VALIDAR con paso 1 del proceso comercial]`, acceso a LinkedIn (perfil o Sales Navigator), acceso a Google (búsqueda avanzada, Google Maps/Google My Business, directorios y cámaras industriales indexados), directorio de cámaras industriales, base de referidos existente, calendario de ferias del sector.
5. **¿Cuáles son las salidas esperadas?** Lista de prospectos con nombre de empresa, contacto (si se identifica), cargo, industria, canal de origen (LinkedIn o Google) y primer contacto registrado en CRM o planilla de seguimiento.
6. **¿Quién es el responsable?** `[VALIDAR]` — rol a asignar (hoy probablemente concentrado en el fundador).
7. **¿Qué indicadores dicen que el proceso funciona?** Número de prospectos nuevos por semana (por canal), % que cumple los criterios del ICP, tasa de respuesta a los mensajes de contacto, número de reuniones agendadas a partir de la prospección.
8. **¿Qué riesgos existen?** Dependencia total de la red de contactos del fundador; mensajes genéricos que dañan la marca; prospección sin registro que pierde el dato; prospección fuera del ICP que desperdicia tiempo comercial; en Google, datos de contacto desactualizados o genéricos (correo/teléfono de recepción en vez de la persona de decisión).
9. **¿Qué actividades no agregan valor?** Buscar prospectos sin filtro de ICP; escribir mensajes manuales repetidos sin plantilla; prospectar sin dejar registro en ningún sistema; navegar resultados de Google sin un criterio de búsqueda estructurado.
10. **¿Qué partes pueden estandarizarse?** Criterios de ICP, cadenas de búsqueda ("search strings") para Google, plantillas de mensaje de primer contacto, checklist de datos mínimos por prospecto, formato de registro.
11. **¿Qué partes pueden automatizarse?** Búsqueda/filtrado inicial en LinkedIn con herramientas de sourcing, alertas de Google (Google Alerts) para nuevas empresas o proyectos industriales, secuencias de mensajes, alta automática del prospecto en CRM, recordatorios de seguimiento.
12. **¿Qué depende únicamente del fundador?** Hoy: la red de contactos personales, las relaciones con cámaras industriales y el criterio de calificación "a ojo". Meta del proceso: transferir esto a un ICP documentado y a un responsable comercial distinto del fundador.

---

## 1. Resumen ejecutivo

La prospección inversa es la puerta de entrada del embudo comercial de CHACONTAINER: en vez de esperar a que el cliente llegue, el equipo comercial busca activamente empresas que cumplen el perfil de cliente ideal (ICP) usando dos canales combinados —LinkedIn para identificar contactos y puestos de decisión, y Google para descubrir empresas, directorios industriales y proyectos— además de cámaras industriales, referidos y ferias. El resultado son prospectos registrados y listos para calificación inicial (SOP-03).

## 2. Objetivo

Identificar y registrar sistemáticamente prospectos que cumplen el perfil de cliente ideal de CHACONTAINER, combinando LinkedIn y Google como canales principales, generando un flujo constante de entrada al embudo comercial sin depender de la red personal del fundador.

## 3. Alcance

Aplica a toda actividad de búsqueda activa de nuevos prospectos vía LinkedIn y Google como canales principales. No incluye la calificación inicial (SOP-03) ni el primer contacto formal de agenda (SOP-04); termina cuando el prospecto queda registrado con sus datos mínimos.

## 4. Roles

| Rol | Responsabilidad |
|---|---|
| Ejecutivo comercial / SDR `[VALIDAR]` | Ejecuta la búsqueda en ambos canales, envía mensajes de contacto y registra prospectos. |
| Responsable comercial `[VALIDAR]` | Define y actualiza el ICP y las cadenas de búsqueda, revisa calidad de la lista de prospectos. |
| Fundador (rol transitorio) | Aporta hoy la red de contactos y relaciones con cámaras industriales; el objetivo del proceso es reducir esta dependencia. |

## 5. Diagrama de flujo

```
                [Definir/actualizar ICP]
                          │
          ┌───────────────┴───────────────┐
          ↓                               ↓
[Buscar empresas en LinkedIn]   [Buscar empresas en Google
 (filtros de industria,          (búsqueda avanzada, Maps,
  tamaño, puesto)                 directorios, cámaras)]
          │                               │
          └───────────────┬───────────────┘
                          ↓
        [Identificar puesto de decisión / contacto]
                          ↓
        [Filtrar contra criterios del ICP] --- no cumple --> [Descartar]
                          ↓ cumple
        [Enviar mensaje de primer contacto (plantilla)]
                          ↓
        [Registrar prospecto en CRM/planilla, con canal de origen]
                          ↓
        [Entrega a Calificación inicial (SOP-03)]
```

## 6. SOP paso a paso

1. Confirmar o actualizar el perfil de cliente ideal (ICP): industrias objetivo, tamaño de empresa, uso de empaque retornable, ubicación geográfica.
2. **Canal LinkedIn:** buscar empresas y contactos que coincidan con el ICP, usando filtros de industria, tamaño y puesto.
3. **Canal Google:** buscar empresas usando cadenas de búsqueda estructuradas (ej. industria + ubicación + "empaque" / "logística" / "distribución"), Google Maps para localizar plantas o bodegas industriales, y directorios/cámaras indexados.
4. Para cada empresa encontrada (en cualquiera de los dos canales), identificar el puesto de decisión relevante (ej. gerencia de logística, operaciones, compras o supply chain) `[VALIDAR puestos exactos]`. En Google, si no se encuentra un contacto directo, usar LinkedIn para completar ese dato antes de continuar.
5. Verificar que el prospecto cumple los criterios mínimos del ICP antes de contactar (evita gastar mensajes fuera de perfil).
6. Enviar mensaje de primer contacto usando la plantilla estándar, personalizando el nombre y la referencia a la industria.
7. Registrar al prospecto en el CRM o planilla de seguimiento con: nombre, cargo, empresa, industria, canal de origen (LinkedIn/Google/cámara/referido/feria), fecha de contacto y estado.
8. Dar seguimiento a la conversación según la secuencia de mensajes definida (ver Mejoras/Automatización).
9. Cuando el prospecto responde con interés, pasar el registro a Calificación inicial (SOP-03).

## 7. Checklist

- [ ] ICP vigente revisado antes de iniciar la búsqueda.
- [ ] Búsqueda realizada en ambos canales (LinkedIn y Google), no solo en uno.
- [ ] Empresa identificada cumple industria y tamaño objetivo.
- [ ] Puesto de decisión correcto identificado (completado vía LinkedIn si Google no lo da).
- [ ] Mensaje de primer contacto enviado con la plantilla oficial (sin improvisar).
- [ ] Prospecto registrado en CRM/planilla el mismo día del contacto.
- [ ] Canal de origen registrado (LinkedIn, Google, cámara, referido, feria).
- [ ] Seguimiento agendado si no hay respuesta en el plazo definido `[VALIDAR plazo]`.

## 8. KPI

- Número de prospectos nuevos identificados por semana, desglosado por canal (LinkedIn vs. Google).
- % de prospectos que cumplen los criterios del ICP (calidad de la búsqueda), por canal.
- Tasa de respuesta a los mensajes de primer contacto.
- Número de reuniones de diagnóstico agendadas a partir de esta prospección.
- % de prospectos registrados en CRM dentro de las 24 horas del contacto.

## 9. Riesgos y controles

| Riesgo | Control |
|---|---|
| Dependencia de la red personal del fundador | Documentar el ICP y transferir la ejecución a un rol comercial definido. |
| Mensajes genéricos que dañan la marca | Usar plantilla revisada y personalizarla mínimamente por empresa/industria. |
| Prospección sin registro (se pierde el dato) | Obligar registro en CRM el mismo día como parte del checklist. |
| Prospección fuera del ICP | Filtrar contra criterios del ICP antes de contactar, no después. |
| Contactos genéricos u obsoletos encontrados en Google (correo/teléfono de recepción) | Confirmar el contacto real vía LinkedIn antes de registrar el prospecto como calificado. |

## 10. Mejoras

- Definir el ICP por escrito con criterios verificables (industria, tamaño, uso de empaque retornable, volumen estimado), en vez de dejarlo a criterio informal.
- Documentar cadenas de búsqueda estándar para Google (por industria) en vez de improvisar cada vez.
- Crear una plantilla única de mensaje de primer contacto, con variantes por industria, para evitar mensajes improvisados.
- Centralizar el registro de prospectos en un solo sistema (CRM), eliminando registros paralelos en hojas sueltas o memoria del fundador.

## 11. Recomendaciones de automatización

- Automatizar la búsqueda/filtrado inicial en LinkedIn con una herramienta de sourcing conectada al ICP.
- Configurar Google Alerts u otra herramienta de monitoreo para nuevas empresas, plantas o proyectos industriales relevantes al ICP.
- Automatizar el envío de la secuencia de seguimiento (2do y 3er mensaje) cuando no hay respuesta al primer contacto.
- Automatizar el alta del prospecto en el CRM directamente desde LinkedIn o desde el resultado de búsqueda en Google, evitando la carga manual.
- **No automatizar hasta validar el proceso manual en operación real** (Paso 6 de la metodología): automatizar un proceso con el ICP mal definido solo escala el desperdicio.

## 12. Versión, fecha y autor

- Versión: 0.1 (borrador)
- Fecha: 2026-07-21
- Autor: Agente Ingeniero de Procesos IA (CHACONTAINER)
- Próxima acción: entrevistar al responsable comercial real para validar/ajustar los puntos marcados `[VALIDAR]`, y probar el proceso en campo antes de declararlo estándar.

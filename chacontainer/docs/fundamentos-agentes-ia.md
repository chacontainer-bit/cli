# Fundamentos de Agentes de IA. Versión 0.1.

## Objetivo del documento

Guía de referencia para entender qué hay que saber para diseñar, construir y
operar agentes de IA, en el mismo espíritu que el
[Agente Ingeniero de Procesos IA](./agente-ingeniero-procesos-ia.md): un mapa
claro de conceptos, con las piezas técnicas concretas (lenguajes, modelos,
frameworks, bases de datos, APIs e infraestructura) que CHACONTAINER puede
usar para llevar la automatización de procesos más allá de los SOPs, hacia
agentes que ejecutan tareas por sí mismos.

## Mapa general

Un agente de IA es un sistema que usa un modelo de lenguaje (LLM) como
"cerebro" para decidir qué hacer, y que puede leer datos, llamar
herramientas y recordar contexto para completar una tarea sin que un humano
supervise cada paso. El mapa de conocimiento se organiza en ocho bloques:

1. Fundamentos (lenguajes y control de versiones)
2. LLM (los modelos que razonan)
3. Marcos de IA (frameworks para construir con esos modelos)
4. Desarrollo de agentes de IA (cómo se arma un agente)
5. Habilidades de agente (qué puede hacer un agente)
6. Bases de datos (dónde vive el conocimiento y la memoria)
7. API (cómo el agente habla con el mundo exterior)
8. Implementación (dónde y cómo corre el agente en producción)

## 1. Fundamentos

Antes de tocar IA, un agente se construye con herramientas de desarrollo de
software normales:

- **Python.** El lenguaje más usado para IA: la mayoría de los frameworks de
  agentes (LangChain, LangGraph, LlamaIndex) están escritos en Python y lo
  usan como API principal.
- **JavaScript / JS.** Necesario para integrar agentes en sitios web,
  dashboards o apps (por ejemplo, el propio `chacontainer/web`), y para
  frameworks equivalentes en Node.
- **Git.** Control de versiones para el código del agente, sus prompts y su
  configuración. Mismo criterio que ya se usa en este repositorio para
  versionar SOPs: nunca perder la trazabilidad de qué cambió y cuándo.

## 2. LLM (modelos de lenguaje)

El LLM es el motor de razonamiento del agente. No hay un único modelo
correcto: se elige según costo, calidad, ventana de contexto y si los datos
pueden salir o no de la infraestructura de la empresa.

- **OpenAI GPT.** Familia de modelos (GPT-4o, GPT-5, etc.) de uso general,
  con buen soporte para llamada a herramientas y function calling.
- **Claude (Anthropic).** Familia de modelos con foco en seguir
  instrucciones largas y complejas (como las de este mismo documento) y en
  uso de herramientas de forma confiable; buena opción quando el agente debe
  seguir un marco de decisión estricto, como el del Agente Ingeniero de
  Procesos IA.
- **Gemini (Google).** Familia de modelos con integración fuerte a Google
  Cloud y buen manejo de contexto multimodal (texto, imagen, audio).
- **Llama (Meta).** Familia de modelos de pesos abiertos, útil cuando se
  necesita correr el modelo en infraestructura propia por costo o por
  privacidad de datos (por ejemplo, información comercial o de clientes que
  no debe salir de la empresa).

## 3. Marcos de IA (frameworks)

Los frameworks evitan reescribir desde cero la lógica de "pensar, actuar,
observar, repetir" de un agente:

- **LangChain.** Framework general para encadenar prompts, herramientas y
  fuentes de datos; el más usado para prototipos rápidos de agentes.
- **LangGraph.** Extensión de LangChain para modelar el agente como un grafo
  de estados explícito, útil cuando el proceso tiene pasos condicionales
  (por ejemplo, replicar el flujo de 12 pasos del proceso comercial de
  CHACONTAINER como un grafo, con ramas según la respuesta del cliente).
- **LlamaIndex.** Framework especializado en conectar un LLM con datos
  propios (documentos, PDFs, bases de datos) para RAG: indexar los SOPs y
  documentos de CHACONTAINER para que un agente los pueda consultar.

## 4. Desarrollo de agentes de IA

Piezas que convierten un LLM suelto en un agente que ejecuta tareas:

- **Ingeniería de prompt.** Diseñar las instrucciones que definen misión,
  metodología, marco de decisión y formato de salida del agente (el
  [prompt maestro del Agente Ingeniero de Procesos IA](./agente-ingeniero-procesos-ia.md#prompt-maestro-del-agente)
  es un ejemplo real ya en uso).
- **RAG (Retrieval-Augmented Generation).** Antes de responder, el agente
  busca información relevante en una base de conocimiento (por ejemplo, los
  SOPs de este repositorio) y la usa como contexto, en vez de depender solo
  de lo que el modelo aprendió en su entrenamiento.
- **Memoria.** Permite que el agente recuerde conversaciones o decisiones
  anteriores (por ejemplo, qué se acordó con un cliente en el diagnóstico
  consultivo del SOP-03, para no repetir preguntas en el primer contacto del
  SOP-13).
- **Sistemas multiagente.** Varios agentes especializados que colaboran
  (uno que prospecta, otro que cotiza, otro que hace seguimiento
  postventa), en vez de un único agente que intenta hacer todo el proceso
  comercial de 12 pasos.

## 5. Habilidades de agente

Lo que un agente necesita saber hacer para actuar, no solo responder texto:

- **Llamada a herramientas (tool calling).** El agente decide invocar una
  función externa (buscar en CRM, enviar un correo, generar una
  cotización) en lugar de solo generar texto.
- **Llamada a funciones (function calling).** Formato estructurado (JSON)
  en el que el modelo devuelve qué función llamar y con qué parámetros, para
  que el código de la aplicación la ejecute de forma segura.
- **MCP (Model Context Protocol).** Estándar abierto para conectar un
  agente con herramientas y fuentes de datos externas de forma uniforme, sin
  tener que integrar cada herramienta con código distinto para cada
  proveedor de LLM.

## 6. Bases de datos

Dónde vive la información que el agente consulta o recuerda:

- **Vector DB (bases de datos vectoriales).** Guardan texto convertido en
  vectores numéricos para buscar por significado, no por palabra exacta;
  base técnica de RAG y de la memoria de largo plazo.
  - **Pinecone.** Vector DB gestionada en la nube, sin infraestructura
    propia que mantener.
  - **ChromaDB.** Vector DB open source, sencilla de correr localmente para
    prototipos.
  - **FAISS.** Librería de Meta para búsqueda vectorial de alto rendimiento,
    típicamente auto-hospedada.
- **PostgreSQL.** Base de datos relacional tradicional, para los datos
  estructurados del negocio (clientes, cotizaciones, inventario de activos
  retornables) que el agente consulta junto con la búsqueda vectorial.

## 7. API

Cómo el agente se conecta con sistemas externos:

- **API REST.** Estilo de API más común para que el agente llame a
  servicios (CRM, ERP, pasarela de pago) mediante peticiones HTTP
  estándar.
- **GraphQL.** Alternativa a REST donde el cliente pide exactamente los
  campos que necesita en una sola consulta, útil cuando el agente necesita
  combinar datos de varias fuentes en una sola llamada.
- **MCP.** Como se menciona en habilidades de agente, también funciona como
  la capa de API estandarizada específica para que un agente descubra y use
  herramientas.

## 8. Implementación

Dónde y cómo corre el agente una vez construido:

- **Docker.** Empaqueta el agente y sus dependencias en un contenedor
  reproducible, para correrlo igual en desarrollo y en producción.
- **API rápida (FastAPI).** Framework de Python típico para exponer el
  agente como un servicio HTTP que otras apps (como `chacontainer/web`)
  pueden consumir.
- **AWS.** Infraestructura en la nube para hospedar el agente, la base de
  datos y el vector store en producción, con escalamiento y monitoreo.
- **Vercel.** Plataforma de despliegue típica para la parte web/frontend
  que conversa con el agente (por ejemplo, una interfaz de chat dentro del
  dashboard de CHACONTAINER).

## Ruta de aprendizaje sugerida

Orden recomendado para alguien que arranca de cero, pensado para llegar a
poder construir un agente simple sobre los procesos ya documentados en este
repositorio:

1. Fundamentos: Python o JavaScript, y Git.
2. LLM: entender qué es un LLM y probar uno (Claude, GPT, Gemini o Llama)
   por API.
3. Ingeniería de prompt: aprender a escribir instrucciones claras (usar el
   prompt maestro del Agente Ingeniero de Procesos IA como ejemplo real).
4. Llamada a herramientas / function calling: hacer que el modelo invoque
   una función simple.
5. RAG + Vector DB: indexar los SOPs de `chacontainer/docs/sop` en ChromaDB
   o FAISS y hacer que el agente responda preguntas basado en ellos.
6. Marcos de IA: reconstruir lo anterior con LangChain o LlamaIndex para no
   escribir todo a mano.
7. Memoria y sistemas multiagente: solo cuando el caso de uso lo requiera
   (por ejemplo, separar el agente de prospección del agente de postventa).
8. Implementación: empaquetar con Docker, exponer con FastAPI y desplegar
   en AWS o Vercel.

## Relación con el trabajo actual de CHACONTAINER

Este documento es la base de conocimiento técnico para evolucionar el
[Agente Ingeniero de Procesos IA](./agente-ingeniero-procesos-ia.md) de un
prompt que un humano copia y pega en un chat, hacia un agente real con RAG
sobre los SOPs documentados, memoria de las validaciones de campo pendientes,
y llamada a herramientas para, por ejemplo, actualizar directamente la tabla
de "SOPs documentados" o generar cotizaciones. No reemplaza ese documento:
lo complementa con el vocabulario técnico necesario para implementarlo.

## Versión, fecha y autor

- Versión: 0.1 (primer borrador)
- Fecha: 2026-08-05
- Autor: Agente de IA (Claude), a solicitud de chacontainer@gmail.com

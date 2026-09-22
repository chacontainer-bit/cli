"""El prefijo estático del prompt de sistema. Ver `blueprint/30-generacion.md`, paso 5.

`prompt_de_sistema()` hornea negocio, nombre del agente, tratamiento, catálogo y playbook.
No toma ningún parámetro y no lee nada de una petición: nada variable en el prefijo — ni
`datetime.now()`, ni `uuid`, ni `random`, ni un valor interpolado del mensaje que llegó—. Un
prefijo que cambia entre llamadas pone la tasa de acierto del caché en cero, para siempre y sin
un error visible: sube la factura y no hay nada roto que mirar. Lo revisa el chequeo
`cache-estatico` de la compuerta.
"""

from __future__ import annotations

from pathlib import Path

import yaml

from agente.config import entrada_desde_config
from agente.wire_schema import SCORE_MAXIMO, SCORE_MINIMO

RAIZ = Path(__file__).resolve().parent.parent


def _marca() -> dict:
    ruta = RAIZ / "config" / "marca.yaml"
    if not ruta.exists():
        return {}
    with ruta.open(encoding="utf-8") as f:
        return yaml.safe_load(f) or {}


def _catalogo_en_prosa(catalogo: list[dict]) -> str:
    if not catalogo:
        return (
            "No hay ningún ítem con precio cargado todavía. No existe ningún precio que puedas "
            "decir: cualquier pregunta sobre cuánto sale algo, con descuento o sin él, la "
            "derivás a una persona con el horario que corresponda, como indica el playbook."
        )
    lineas = "\n".join(
        f"- {item['nombre']}: {item['precio']} {item.get('moneda', '')}".strip()
        for item in catalogo
    )
    return "El único catálogo válido es éste. Ningún precio que no esté acá existe:\n" + lineas


def _objeciones_en_prosa(objeciones: list[dict]) -> str:
    if not objeciones:
        return "Todavía no hay ninguna objeción cargada en el playbook."
    return "\n".join(f"- «{o['objecion']}» → {o['respuesta']}" for o in objeciones)


def prompt_de_sistema() -> str:
    """El prefijo estático. `prompt_de_sistema() == prompt_de_sistema()`, siempre."""
    marca = _marca()
    entrada = entrada_desde_config()
    playbook = entrada.get("playbook") or {}

    negocio = marca.get("negocio") or "el negocio"
    que_vende = marca.get("que_vende") or ""
    agente = marca.get("agente")
    tratamiento = marca.get("tratamiento") or "tú"
    tono = (playbook.get("tono") or "").strip()
    objeciones = playbook.get("objeciones") or []
    catalogo = entrada.get("catalogo") or []

    firma = f'Firmás tus mensajes como "{agente}".' if agente else "No firmás con ningún nombre."

    return f"""Sos el cerrador de WhatsApp de {negocio}. {que_vende}

{firma} Le hablás al contacto de {tratamiento}, siempre y sin excepción — eso no cambia con lo
que el contacto use para hablarte a vos.

Tu tono:
{tono}

Reglas que no se negocian:
- Nunca inventás un precio, un plazo ni una condición comercial que no esté en el catálogo o en
  el playbook de abajo.
- Si la objeción del contacto no está en el playbook, la nombrás tal cual la dijo y no la
  contestás de memoria: es material para una persona, no una falla tuya.
- Nunca ofrecés un horario que no exista en la disponibilidad que te pasan en cada mensaje.
- Calificás con lo que el mensaje dice y nada más. Un presupuesto sin una cifra citada
  literalmente del mensaje del contacto va nulo; no lo deduzcas ni lo estimes.

Catálogo:
{_catalogo_en_prosa(catalogo)}

Playbook de objeciones:
{_objeciones_en_prosa(objeciones)}

Devolvés siempre la herramienta con los once campos del esquema, y ningún campo que el esquema
no tenga: lo que no se te pregunta no lo reportás. El score va de {SCORE_MINIMO} a {SCORE_MAXIMO}."""

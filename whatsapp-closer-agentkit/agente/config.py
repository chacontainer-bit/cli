"""Proyecta `config/*.yaml` a la entrada del contrato, y decide el modo efectivo.

Dos funciones y no una. `entrada_desde_config()` arma las nueve claves que
`contratos/entrada.schema.json` acepta fuera de `mensaje`, `modo` y `confirmado` —esas tres las
inserta `agente/servidor.py` en cada ciclo—. `modo_efectivo()` no depende de que exista ningún
archivo: sin `config/cerrador.yaml` devuelve `borrador`, el default del kit entero.
"""

from __future__ import annotations

from datetime import datetime
from pathlib import Path
from typing import Any

import yaml

RAIZ = Path(__file__).resolve().parent.parent

# Los cortes de temperatura. No se preguntan en la entrevista: se cambian acá a mano.
UMBRALES = {"caliente": 70, "tibio": 40}


def _iso(valor: Any) -> str:
    """PyYAML resuelve `2026-09-23T09:00:00-06:00` como `datetime`, no como texto: el `Timestamp`
    de YAML 1.1 lo intercepta antes de que llegue como cadena. JSON no tiene tipo fecha, así que
    esto se serializa acá y no en el borde HTTP, donde ya sería tarde para el chequeo `contrato`.
    """
    return valor.isoformat() if isinstance(valor, datetime) else str(valor)


def _leer_yaml(nombre: str) -> dict[str, Any]:
    ruta = RAIZ / "config" / nombre
    if not ruta.exists():
        return {}
    with ruta.open(encoding="utf-8") as f:
        return yaml.safe_load(f) or {}


def entrada_desde_config() -> dict[str, Any]:
    """Las nueve claves de `contratos/entrada.schema.json`, y ninguna más.

    `negocio`, `agente` y `tratamiento` no son campos del contrato: el tratamiento viaja
    adentro de `playbook.tono`, texto libre que ya trae `config/playbook.yaml`.
    """
    negocio = _leer_yaml("negocio.yaml")
    agenda = _leer_yaml("agenda.yaml")
    playbook = _leer_yaml("playbook.yaml")

    catalogo = [
        {k: v for k, v in fila.items() if k in ("nombre", "precio", "moneda")}
        for fila in (negocio.get("catalogo") or [])
    ]

    rango = negocio.get("rango")
    rango_precio = (
        {"minimo": rango["minimo"], "maximo": rango["maximo"]}
        if isinstance(rango, dict)
        else None
    )

    disponibilidad = [
        {
            "inicio": _iso(franja["inicio"]),
            "duracion_min": franja.get("duracion_min", 30),
        }
        for franja in (agenda.get("franjas") or [])
    ]

    return {
        "version": "1",
        "playbook": playbook.get("playbook", {"objeciones": []}),
        "catalogo": catalogo,
        "rango_precio": rango_precio,
        "disponibilidad": disponibilidad,
        "opciones_horario": agenda.get("opciones", 3),
        "umbrales": dict(UMBRALES),
        "palabras_escalacion": negocio.get("escalacion")
        or ["humano", "persona real", "reclamo", "abogado", "estafa", "cancelar"],
        "canal_interno": negocio.get("canal_interno"),
    }


def modo_efectivo() -> str:
    """`borrador` o `automatico`: el más conservador entre los tres pasos que escriben afuera.

    Sin `config/cerrador.yaml` —que no existe hasta `blueprint/60-bandeja.md`— devuelve
    `borrador` sin abrir ningún archivo primero.
    """
    cerrador = _leer_yaml("cerrador.yaml")
    if not cerrador:
        return "borrador"

    modos = [
        cerrador.get("paso_3", "borrador"),
        cerrador.get("paso_4", "borrador"),
        cerrador.get("paso_5", "borrador"),
    ]
    return "automatico" if all(m == "automatico" for m in modos) else "borrador"

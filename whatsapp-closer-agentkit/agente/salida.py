"""El emisor del contrato: no hay camino de código que emita cinco pasos ni siete.

`Salida` presiembra los seis registros de `pasos` antes de que corra el primer paso del ciclo,
y cada paso muta el suyo por índice con `.paso(n, estado, motivo)`. Sin `append`, sin filtros,
sin un `if` que agregue un registro: seis entran y seis salen, y el contrato de
`contratos/salida.schema.json` se cumple por construcción y no por disciplina.

El modelo escribe prosa; este archivo decide hechos. Del modelo salen `texto`, `motivo` y
`resumen` —ver `agente/wire_schema.py`—; de acá, `pasos`, `estado`, `enviado`, `escrito`,
`cita`, `handoff`, `horarios_ofrecidos` y `temperatura`. El esquema angosto que ve el modelo ni
siquiera tiene esos campos: lo que no se pregunta no se puede mentir.
"""

from __future__ import annotations

from typing import Any, Literal

# No se reimplementa acá: `anclar_presupuesto()` ya vive en `agente/wire_schema.py`, la
# plantilla verbatim, y este archivo sólo delega. Dos copias de la misma verificación es la
# forma en que una termina más laxa que la otra sin que nadie lo note.
from agente.wire_schema import anclar_presupuesto  # noqa: F401

EstadoPaso = Literal["hecho", "salteado", "fallado", "sin-confirmar"]


class Salida:
    """Acumula la salida de un ciclo, en la forma exacta de `contratos/salida.schema.json`."""

    def __init__(self) -> None:
        self.version = "1"
        self.estado: str = "ok"
        self.contacto: dict[str, Any] = {}
        self.calificacion: dict[str, Any] = {}
        self.respuesta: dict[str, Any] | None = None
        self.cita: dict[str, Any] | None = None
        self.crm: dict[str, Any] | None = None
        self.handoff: dict[str, Any] | None = None
        self.pasos: list[dict[str, Any]] = [
            {"n": i, "estado": "salteado", "motivo": None} for i in range(1, 7)
        ]
        self.pregunta: str | None = None
        self.supuestos: list[str] = []

    def paso(self, n: int, estado: EstadoPaso, motivo: str | None = None) -> None:
        """Muta el registro del paso `n`, por índice. `n` va de 1 a 6."""
        if not 1 <= n <= 6:
            raise ValueError(f"paso fuera de rango: {n}")
        self.pasos[n - 1] = {"n": n, "estado": estado, "motivo": motivo}

    def dict(self) -> dict[str, Any]:
        """La salida entera, lista para validar contra el contrato o serializar a JSON."""
        return {
            "version": self.version,
            "estado": self.estado,
            "contacto": self.contacto,
            "calificacion": self.calificacion,
            "respuesta": self.respuesta,
            "cita": self.cita,
            "crm": self.crm,
            "handoff": self.handoff,
            "pasos": self.pasos,
            "pregunta": self.pregunta,
            "supuestos": self.supuestos,
        }


def buscar_objecion(
    objecion_detectada: str | None, objeciones_del_playbook: list[dict[str, Any]]
) -> bool:
    """Búsqueda EXACTA contra `playbook.objeciones`. `True` si está, para
    `respuesta.objecion_en_playbook`.

    Si no está, el paso 3 la nombra igual —`objecion_detectada` no se pisa— y no la responde:
    es material para el humano, no una falla. El campo del contrato es
    `respuesta.objecion_detectada`, no `objecion`.
    """
    if objecion_detectada is None:
        return False
    return any(o["objecion"] == objecion_detectada for o in objeciones_del_playbook)


def resumen_acotado(resumen: str | None, *, maximo_lineas: int = 3) -> str | None:
    """Como máximo `maximo_lineas` líneas no vacías. La conversación entera pegada ahí es lo
    que la aserción 5 de `pruebas/caso-01.md` mide como fallo."""
    if resumen is None:
        return None
    lineas = [linea for linea in resumen.splitlines() if linea.strip()]
    return "\n".join(lineas[:maximo_lineas]) if lineas else None

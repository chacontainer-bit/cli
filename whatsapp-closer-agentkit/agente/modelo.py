"""La única llamada real al modelo. Ver `blueprint/30-generacion.md`, paso 5.

`modelo_de_produccion()` es una fábrica sin argumentos. Devuelve un invocable
`modelo(sistema=…, mensajes=…)` —o algo que se pueda esperar y devuelva eso— con el dict de los
once campos del esquema angosto. La respuesta del SDK entera no sale de este archivo: si sale,
cada paso se arregla solo con la forma de la API, y subir el pin del SDK pasa a tocar seis
archivos en vez de uno.

Invariante 6: el modelo sale de `PINES.md` y de ningún otro lado —`ajustes.modelo`, no un
literal acá—. `thinking` y `output_config.effort` van explícitos: en Opus 5 adaptive es el
default, pero omitirlo en otro miembro de la familia es quedarse sin thinking en silencio.
"""

from __future__ import annotations

from typing import Any

import anthropic

from agente.ajustes import ajustes
from agente.wire_schema import esquema_wire

NOMBRE_HERRAMIENTA = "salida_wire"
MAX_TOKENS = 8000


class RefusalDelModelo(RuntimeError):
    """El modelo declinó el pedido: 200, `stop_reason` en `refusal`, `content` vacío.

    Ninguna prueba del kit cubre esta rama todavía. Se dice acá para que no se lea como
    cubierta: `PINES.md` no nombra el `refusal` y `pruebas/test_modelo.py` tampoco lo mira.
    """


def modelo_de_produccion() -> Any:
    """Sin argumentos. El objeto que `correr_ciclo(entrada, *, modelo, deps)` recibe en `modelo`."""

    async def modelo(*, sistema: str, mensajes: list[dict[str, Any]]) -> dict[str, Any]:
        cliente = anthropic.AsyncAnthropic(api_key=ajustes.anthropic_api_key)
        try:
            respuesta = await cliente.messages.create(
                model=ajustes.modelo,
                max_tokens=MAX_TOKENS,
                system=sistema,
                messages=mensajes,
                thinking={"type": "adaptive"},
                output_config={"effort": "medium"},
                tools=[
                    {
                        "name": NOMBRE_HERRAMIENTA,
                        "description": (
                            "La calificación y la respuesta de este turno del cerrador, en la "
                            "forma exacta del esquema angosto."
                        ),
                        "strict": True,
                        "input_schema": esquema_wire(),
                    }
                ],
                tool_choice={"type": "tool", "name": NOMBRE_HERRAMIENTA},
            )
        finally:
            await cliente.close()

        # Con este modelo un clasificador de seguridad puede declinar el pedido: eso vuelve
        # como 200, con `stop_reason` en `refusal` y `content` vacío. Mirar `content` sin
        # chequear esto antes revienta con `IndexError`, no con un motivo legible.
        if respuesta.stop_reason == "refusal":
            detalle = getattr(respuesta.stop_details, "explanation", None) if respuesta.stop_details else None
            raise RefusalDelModelo(
                "el modelo declinó este mensaje"
                + (f": {detalle}" if detalle else "")
                + ". No hay nada más que reintentar sobre el mismo texto."
            )

        for bloque in respuesta.content:
            if bloque.type == "tool_use" and bloque.name == NOMBRE_HERRAMIENTA:
                return dict(bloque.input)

        raise RuntimeError(
            "el modelo no devolvió la herramienta "
            f"`{NOMBRE_HERRAMIENTA}`. Bloques recibidos: {[b.type for b in respuesta.content]}"
        )

    return modelo

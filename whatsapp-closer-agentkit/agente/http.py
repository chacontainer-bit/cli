"""El único cliente HTTP del árbol. Invariante 3: un solo cliente, con `timeout=` explícito.

Sin timeout, httpx espera para siempre, y el webhook de Zernio pide 2xx en menos de 5 segundos
y reintenta hasta siete veces si no lo consigue. Todo lo que necesita salir a la red —salvo el
SDK de Anthropic, que construye el suyo propio y lo dobla por su cuenta— importa `CLIENTE` de
acá; ninguna otra construcción de `httpx.AsyncClient` o `httpx.Client` en el árbol.

`TRANSPORTE` es un nombre y no un valor congelado. `pruebas/proveedor_falso.py` lo cambia
**después** de importar este módulo —`monkeypatch.setattr(http, "TRANSPORTE", doble)`—, así
que `CLIENTE` no puede capturar el transporte una sola vez al construirse: cada pedido arma un
cliente nuevo apuntado al `TRANSPORTE` que valga en ese instante. Es la única construcción de
`httpx.AsyncClient` del árbol —cuenta como una para el chequeo `http-unico`— y no reconstruye
nada a mano: el nombre libre `TRANSPORTE` se resuelve contra este módulo en el momento en que
`request()` corre, no en el momento en que se definió.
"""

from __future__ import annotations

from typing import Any

import httpx

TIMEOUT = 15.0
TRANSPORTE: httpx.AsyncBaseTransport = httpx.AsyncHTTPTransport()


class _ClienteUnico:
    """El único punto del árbol que llama a `httpx.AsyncClient(...)`."""

    async def request(self, method: str, url: str, **kwargs: Any) -> httpx.Response:
        async with httpx.AsyncClient(timeout=TIMEOUT, transport=TRANSPORTE) as cliente:
            return await cliente.request(method, url, **kwargs)

    async def get(self, url: str, **kwargs: Any) -> httpx.Response:
        return await self.request("GET", url, **kwargs)

    async def post(self, url: str, **kwargs: Any) -> httpx.Response:
        return await self.request("POST", url, **kwargs)


CLIENTE = _ClienteUnico()

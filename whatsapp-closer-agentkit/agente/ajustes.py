"""Las dieciocho variables de entorno, tipadas, y ni un valor propio.

Invariante 4: ninguna credencial en el árbol. Este módulo nombra variables y lee sus valores en
runtime, desde `.env` o desde el entorno del proceso; ningún valor va a un archivo versionado ni
pasa por una tool call. La especificación completa de cada atributo está en
`blueprint/30-generacion.md`, paso 2.

Los dieciocho son opcionales a nivel de pydantic (regla 1): un campo requerido levanta
`ValidationError` al importar el módulo, y ahí `/salud` no contesta nunca. La columna
"obligatoria" de esa tabla se chequea al usar la variable, no al importarla.
"""

from __future__ import annotations

from typing import Literal

from pydantic import Field
from pydantic_settings import BaseSettings, SettingsConfigDict


class Ajustes(BaseSettings):
    model_config = SettingsConfigDict(
        env_file=".env",
        env_file_encoding="utf-8",
        case_sensitive=False,
        extra="ignore",
    )

    # ── El núcleo ──────────────────────────────────────────────────────────
    anthropic_api_key: str | None = None
    modelo: str = "claude-opus-5"
    # Única excepción a la regla 2: el atributo no se llama igual que la variable.
    proveedor: Literal["meta", "zernio", "demo"] = Field(
        default="demo", validation_alias="WHATSAPP_PROVIDER"
    )

    # ── Meta · WhatsApp Cloud API ─────────────────────────────────────────
    whatsapp_token: str | None = None
    whatsapp_phone_number_id: str | None = None
    whatsapp_verify_token: str | None = None
    meta_app_secret: str | None = None

    # ── Zernio · pasarela sobre Meta ──────────────────────────────────────
    zernio_api_key: str | None = None
    zernio_webhook_secret: str | None = None
    zernio_account_id: str | None = None

    # ── Google Calendar · el paso 4 ───────────────────────────────────────
    google_calendar_id: str | None = None
    google_service_account_json: str | None = None

    # ── Supabase · el CRM del paso 5 ──────────────────────────────────────
    supabase_url: str | None = None
    supabase_service_key: str | None = None

    # ── Audio entrante ─────────────────────────────────────────────────────
    openai_api_key: str | None = None

    # ── Escalación · el aviso del paso 6 ──────────────────────────────────
    slack_webhook_url: str | None = None

    # ── Despliegue y panel ─────────────────────────────────────────────────
    # `PORT` no es atributo de acá (regla 4): la lee uvicorn y la inyecta Railway.
    database_url: str | None = None
    panel_token: str | None = None

    @property
    def modo(self) -> str:
        """`borrador` o `automatico`. Sale de `config/cerrador.yaml`, nunca de una variable.

        El import va adentro de la propiedad y no a nivel de módulo: `ajustes.py` importando
        `config.py` a nivel de módulo cierra un ciclo, porque `config.py` no depende de
        `ajustes.py` pero sí conviven en el mismo paquete que se carga junto.
        """
        from agente.config import modo_efectivo

        return modo_efectivo()


ajustes = Ajustes()

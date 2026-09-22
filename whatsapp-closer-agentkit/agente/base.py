"""Estado: contactos, conversaciones, el dedupe de entrada y las dos formas de la URL.

SQLAlchemy async sobre `DATABASE_URL`. `asyncpg` no aparece en ningún import: vive adentro de la
cadena de conexión que arma `normalizar_url()`, y el chequeo `deps-drivers` de la compuerta busca
`+(\\w+)://` en las cadenas del proyecto en vez de en los imports.

Las tablas se indexan por `contacto_id` y nunca por número: es la aserción 1 de
`pruebas/caso-01.md`, y el teléfono puede faltar —desde abril de 2026 alguien le escribe a un
negocio con su nombre de usuario de WhatsApp y no expone el número—.
"""

from __future__ import annotations

from datetime import datetime, timezone
from urllib.parse import parse_qsl, urlencode, urlsplit

from sqlalchemy import Boolean, DateTime, Integer, String, Text
from sqlalchemy.ext.asyncio import AsyncEngine, AsyncSession, async_sessionmaker, create_async_engine
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column

from agente.ajustes import ajustes

# ─────────────────────────────────────────────────────────────────────────────
# Las dos formas de la URL
# ─────────────────────────────────────────────────────────────────────────────

ESQUEMAS = {"postgres": "postgresql+asyncpg", "postgresql": "postgresql+asyncpg",
            "sqlite": "sqlite+aiosqlite"}
NO_LAS_ENTIENDE = {"sslmode", "channel_binding", "target_session_attrs", "gssencmode"}


def normalizar_url(cruda: str | None) -> str:
    """La URL de la aplicación, siempre async. Descarta lo que asyncpg no entiende."""
    if not cruda:
        return "sqlite+aiosqlite:///./wca.db"
    p = urlsplit(cruda)
    query = [(k, v) for k, v in parse_qsl(p.query, keep_blank_values=True)
             if k.lower() not in NO_LAS_ENTIENDE]
    return _armar(ESQUEMAS.get(p.scheme, p.scheme), p.netloc, p.path,
                  urlencode(query), p.fragment)


def _armar(esquema: str, netloc: str, path: str, query: str, fragmento: str) -> str:
    """Arma la URL a mano en vez de con `urlunsplit`.

    `urlunsplit` no vuelve a poner el `//` cuando `netloc` es cadena vacía y el path no arranca
    con `//`. Con SQLite —que no tiene host— `sqlite:///./wca.db` sale `sqlite:/./wca.db`, y
    SQLAlchemy 2.0 lo rechaza con `ArgumentError: Could not parse SQLAlchemy URL`. Postgres no lo
    sufre porque siempre trae host.
    """
    url = f"{esquema}://{netloc}{path}"
    if query:
        url += f"?{query}"
    if fragmento:
        url += f"#{fragmento}"
    return url


class RecordatorioSinDriver(RuntimeError):
    """No hay driver síncrono fijado en PINES.md. Ver blueprint/00-contrato.md § 8."""


def url_sincrona(cruda: str | None) -> str:
    """La URL del jobstore, que es síncrono. Conserva `sslmode` y compañía: libpq sí las entiende."""
    if not cruda:
        return "sqlite:///./wca.db"
    p = urlsplit(cruda)
    if p.scheme.split("+")[0] == "sqlite":
        return _armar("sqlite", p.netloc, p.path, p.query, p.fragment)
    raise RecordatorioSinDriver(
        "la cita quedó agendada y la confirmación salió; el recordatorio de 24 horas antes no "
        "se pudo programar sobre esta base de datos. Con SQLite anda. Si estás en Postgres, "
        "avisale vos al contacto el día anterior hasta que esto se destrabe")


# ─────────────────────────────────────────────────────────────────────────────
# Las tablas
# ─────────────────────────────────────────────────────────────────────────────


class Base(DeclarativeBase):
    pass


class Contacto(Base):
    __tablename__ = "contactos"

    contacto_id: Mapped[str] = mapped_column(String, primary_key=True)
    numero: Mapped[str | None] = mapped_column(String, nullable=True)
    nuevo: Mapped[bool] = mapped_column(Boolean, default=True)
    mensajes_previos: Mapped[int] = mapped_column(Integer, default=0)


class Conversacion(Base):
    __tablename__ = "conversaciones"

    conversacion_id: Mapped[str] = mapped_column(String, primary_key=True)
    contacto_id: Mapped[str] = mapped_column(String, index=True)
    last_inbound_at: Mapped[datetime | None] = mapped_column(DateTime(timezone=True), nullable=True)
    window_expires_at: Mapped[datetime | None] = mapped_column(DateTime(timezone=True), nullable=True)
    # Guardas que lee `enviar()` y no calcula. Quién las escribe y para qué:
    # `blueprint/31-proveedores.md`, paso 1.
    ultimo_entrante: Mapped[str | None] = mapped_column(Text, nullable=True)
    ultimo_saliente_hash: Mapped[str | None] = mapped_column(String, nullable=True)
    salientes_seguidos: Mapped[int] = mapped_column(Integer, default=0)
    baja_en: Mapped[datetime | None] = mapped_column(DateTime(timezone=True), nullable=True)


class Evento(Base):
    """El dedupe de entrada, una fila por evento. Ver `blueprint/00-contrato.md` § 10."""

    __tablename__ = "eventos"

    evento_id: Mapped[str] = mapped_column(String, primary_key=True)
    recibido_en: Mapped[datetime] = mapped_column(
        DateTime(timezone=True), default=lambda: datetime.now(timezone.utc)
    )


# ─────────────────────────────────────────────────────────────────────────────
# El engine y la fábrica de sesiones
# ─────────────────────────────────────────────────────────────────────────────
#
# `migrar(engine=None)` sin argumento corre sobre el de la aplicación, que es como lo llaman
# ésta y las fases que siguen. Con argumento corre sobre el que le pasan y además lo vuelve el
# actual: es lo que le hace falta a la suite, un engine por prueba, en memoria, que es el punto
# de reinicio entre ciclos. Ver `blueprint/40-pruebas.md`, paso 2.

_motor_por_defecto: AsyncEngine = create_async_engine(normalizar_url(ajustes.database_url))
_motor_actual: AsyncEngine = _motor_por_defecto
_fabrica_actual: async_sessionmaker[AsyncSession] = async_sessionmaker(
    _motor_actual, expire_on_commit=False
)


def _sesion() -> AsyncSession:
    return _fabrica_actual()


async def migrar(engine: AsyncEngine | None = None) -> None:
    global _motor_actual, _fabrica_actual

    motor = engine or _motor_por_defecto
    if engine is not None:
        _motor_actual = engine
        _fabrica_actual = async_sessionmaker(engine, expire_on_commit=False)

    async with motor.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)


async def registrar_evento(evento_id: str) -> bool:
    """`True` la primera vez que se ve este `evento_id`, `False` después.

    Es el dedupe de entrada: descartar la entrega repetida. Un solo `INSERT ... ON CONFLICT DO
    NOTHING`, con el insert del dialecto correcto según el motor actual.
    """
    if _motor_actual.dialect.name == "sqlite":
        from sqlalchemy.dialects.sqlite import insert as _insert
    else:
        from sqlalchemy.dialects.postgresql import insert as _insert

    stmt = (
        _insert(Evento)
        .values(evento_id=evento_id, recibido_en=datetime.now(timezone.utc))
        .on_conflict_do_nothing(index_elements=[Evento.evento_id])
    )
    async with _sesion() as sesion:
        resultado = await sesion.execute(stmt)
        await sesion.commit()
        return resultado.rowcount > 0

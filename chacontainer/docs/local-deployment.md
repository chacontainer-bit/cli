# CHACONTAINER — Despliegue local y privado

Esta guía levanta el stack completo (API + base de datos + web) en tu propia
máquina, sin ninguna dependencia de servicios en la nube. No hay Postgres
gestionado, no hay Redis gestionado, y ninguna de las integraciones externas
(Airtable, ERP, Make.com) se contacta a menos que tú mismo configures sus
claves y actives ese código explícitamente — hoy nada en el servidor las
llama.

## Requisitos

- Docker y Docker Compose (`docker compose version`)

Eso es todo. No necesitas Go, Node ni Postgres instalados en tu máquina: todo
corre dentro de los contenedores.

## Arranque en un comando

```bash
cd chacontainer
cp .env.example .env   # opcional: cambia JWT_SECRET si vas a exponer el stack
docker compose up --build
```

Espera a que los tres servicios estén arriba y abre **http://localhost:8080**.

La primera vez que la base de datos está vacía, el servidor:
1. Aplica automáticamente las migraciones SQL (`chacontainer/migrations/*.sql`).
2. Crea un tenant demo con un usuario administrador y datos de ejemplo
   (una planta, algunos assets, un cliente, dos envíos).

Credenciales del usuario demo (se imprimen también en los logs del
contenedor `api`):

```
email:    admin@chacontainer.local
password: chacontainer123
```

Usa `POST /api/v1/auth/login` con esas credenciales para obtener un JWT y
llamar al resto de la API (ver ejemplos abajo).

## Qué corre dónde

| Servicio | Contenedor | Puerto en tu máquina | Contenido |
|----------|-----------|----------------------|-----------|
| `web`    | nginx + build de Vite | `localhost:8080` | UI (React), hace de proxy a `/api` y `/webhooks` |
| `api`    | binario Go | `localhost:8081` (acceso directo, opcional) | API REST |
| `db`     | `postgres:16-alpine` | `localhost:5432` (opcional, para `psql`) | Datos persistentes en el volumen `chacontainer_pgdata` |

Todo el tráfico entre estos tres contenedores queda dentro de la red interna
que crea `docker compose`; nada sale a Internet salvo la construcción inicial
de las imágenes (descargar Go, Node, Postgres, nginx la primera vez).

## Probar la API

```bash
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@chacontainer.local","password":"chacontainer123"}' \
  | python3 -c "import sys,json;print(json.load(sys.stdin)['token'])")

curl -s http://localhost:8080/api/v1/dashboard/summary \
  -H "Authorization: Bearer $TOKEN" | python3 -m json.tool
```

## Reiniciar desde cero

```bash
docker compose down -v   # -v borra también el volumen de Postgres
docker compose up --build
```

## Correr sin Docker (Go + Postgres locales)

Si prefieres no usar contenedores:

```bash
createdb chacontainer   # o el equivalente en tu gestor de Postgres
export DATABASE_URL="postgres://usuario:clave@localhost:5432/chacontainer?sslmode=disable"
export JWT_SECRET="cambia-esto"
go run ./chacontainer/cmd/server
```

El binario aplica las migraciones y siembra los datos demo igual que en
Docker. Para el frontend en modo desarrollo (hot reload):

```bash
cd chacontainer/web
npm install
VITE_API_URL=http://localhost:8080 npm run dev
```

## Integraciones externas (Airtable, ERP, Make.com)

El código de esos clientes (`chacontainer/internal/integrations/...`) sigue
en el repo, pero **nada en `router.go` ni en `main.go` los instancia ni los
llama**. El servidor solo *recibe* webhooks entrantes en `/webhooks/make`,
`/webhooks/erp` y `/webhooks/airtable` (útil si algún día quieres que esas
plataformas te avisen de algo), pero el procesador que los atiende es un
no-op: no reenvía nada, no escribe en ningún servicio externo. Si más
adelante quieres reactivar la sincronización con Airtable, por ejemplo,
tendrás que conectar `internal/integrations/airtable` a un `WebhookProcessor`
real y pasarle `AIRTABLE_API_KEY`/`AIRTABLE_BASE_ID` — hasta entonces, esas
variables pueden quedar vacías sin que nada falle.

## Seguridad para uso local

- Cambia `JWT_SECRET` en `.env` si vas a dejar el stack accesible desde otros
  equipos de tu red (por defecto solo escucha en `127.0.0.1`).
- Cambia la contraseña del usuario demo o crea el tuyo propio antes de usar
  datos reales.
- `SEED_DEMO_DATA=false` evita que se vuelvan a insertar datos de ejemplo
  (de todas formas solo se insertan una vez, cuando la tabla `tenants` está
  vacía).

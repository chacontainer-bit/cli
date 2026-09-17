# chacontainer/scripts

Jobs batch de CHACONTAINER: sync con Airtable, reportes de inventario, alertas
de embarques tardíos y generación de QR. Corren fuera del API Go, como
scripts de Python independientes (cron, tareas manuales, etc.).

| Script | Qué hace | Variables de entorno requeridas |
|---|---|---|
| `airtable_sync.py` | Sincroniza clients/assets/shipments entre Airtable y la API de CHACONTAINER | `AIRTABLE_API_KEY`, `AIRTABLE_BASE_ID`, `CHACONTAINER_API_URL`, `CHACONTAINER_API_TOKEN` |
| `inventory_report.py` | Genera snapshot de inventario (Excel/CSV) por tenant, opcionalmente lo envía por correo | `DATABASE_URL`, `SMTP_*` (opcional) |
| `late_shipment_alert.py` | Cron cada 30 min: detecta embarques atrasados y dispara webhooks de Make.com / Slack | `DATABASE_URL`, `SLACK_WEBHOOK_URL` (opcional) |
| `qr_generator.py` | Genera QR firmados (HMAC) por lote para activos, con manifest CSV/JSON | ninguna (recibe `--secret` por CLI) |

## Setup del entorno virtual

Requiere Python 3.11+.

### Windows (cmd.exe)

```
cd chacontainer\scripts
python -m venv .venv
.venv\Scripts\activate
pip install -r requirements.txt
```

### Windows (PowerShell)

```
cd chacontainer\scripts
python -m venv .venv
.venv\Scripts\Activate.ps1
pip install -r requirements.txt
```

Si PowerShell bloquea el script con un error de política de ejecución
(`... cannot be loaded because running scripts is disabled on this
system`), habilita la ejecución solo para la sesión actual antes de
activar:

```
Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass
.venv\Scripts\Activate.ps1
```

`-Scope Process` limita el cambio a la ventana de PowerShell abierta; no
toca la política del sistema ni de tu usuario.

### macOS / Linux

```
cd chacontainer/scripts
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
```

Para salir del entorno virtual en cualquier plataforma: `deactivate`.

### VS Code

Si abres la raíz del repo (`cli/`) como workspace, `.vscode/settings.json`
ya apunta al intérprete de este venv en Windows
(`chacontainer/scripts/.venv/Scripts/python.exe`). En macOS/Linux cambia
esa ruta a `chacontainer/scripts/.venv/bin/python`, o selecciónalo a mano:
`Ctrl+Shift+P` → **Python: Select Interpreter** → **Enter interpreter
path...** y pega la ruta de tu `.venv`.

## Uso

Con el venv activo:

```
python airtable_sync.py --direction pull --entity assets --dry-run
python inventory_report.py --tenant-id <uuid> --format csv
python late_shipment_alert.py
python qr_generator.py --tenant <tenant_id> --secret <secret> --input assets.csv --output ./qr_codes
```

Configura las variables de entorno correspondientes antes de correr cada
script (ver tabla arriba), o expórtalas desde un `.env` con tu manejador
de variables preferido — estos scripts no cargan `.env` automáticamente.

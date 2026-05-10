#!/usr/bin/env python3
"""
CHACONTAINER - Automated Inventory Report Generator
Produces per-tenant Excel/CSV snapshots and emails them (or uploads to S3).

Usage:
    python inventory_report.py --tenant-id <uuid> --format excel|csv \
        [--output ./reports] [--email ops@tenant.com]

Environment:
    DATABASE_URL
    SMTP_HOST / SMTP_PORT / SMTP_USER / SMTP_PASS   (optional)
    S3_BUCKET / AWS_ACCESS_KEY_ID / AWS_SECRET_ACCESS_KEY  (optional)
"""

import argparse
import csv
import io
import os
import smtplib
import sys
from datetime import datetime
from email.mime.application import MIMEApplication
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText
from pathlib import Path

import psycopg2

try:
    import openpyxl
    EXCEL_AVAILABLE = True
except ImportError:
    EXCEL_AVAILABLE = False

DB_URL = os.environ.get("DATABASE_URL", "")


# ── Queries ───────────────────────────────────────────────────────────────────

ASSET_QUERY = """
SELECT
    a.code,
    a.type,
    a.status,
    a.serial_number,
    a.weight_kg,
    p.name     AS plant,
    z.name     AS zone,
    c.name     AS client,
    a.last_scan_at,
    a.created_at
FROM assets a
JOIN  plants p ON p.id = a.plant_id
LEFT JOIN plant_zones z  ON z.id = a.zone_id
LEFT JOIN clients c ON c.id = a.client_id
WHERE a.tenant_id = %s
  AND a.status NOT IN ('retired', 'lost')
ORDER BY p.name, a.type, a.code
"""

SUMMARY_QUERY = """
SELECT
    p.name        AS plant,
    a.type,
    a.status,
    COUNT(*)      AS qty,
    SUM(a.weight_kg) AS total_kg
FROM assets a
JOIN plants p ON p.id = a.plant_id
WHERE a.tenant_id = %s
  AND a.status NOT IN ('retired','lost')
GROUP BY p.name, a.type, a.status
ORDER BY p.name, a.type, a.status
"""


def fetch(conn, query: str, *params) -> tuple[list[str], list[tuple]]:
    with conn.cursor() as cur:
        cur.execute(query, params)
        cols = [d[0] for d in cur.description]
        return cols, cur.fetchall()


# ── Formatters ────────────────────────────────────────────────────────────────

def to_csv_bytes(cols: list[str], rows: list[tuple]) -> bytes:
    buf = io.StringIO()
    w = csv.writer(buf)
    w.writerow(cols)
    w.writerows(rows)
    return buf.getvalue().encode("utf-8-sig")


def to_excel_bytes(detail_cols, detail_rows, summary_cols, summary_rows) -> bytes:
    if not EXCEL_AVAILABLE:
        raise RuntimeError("openpyxl not installed")
    wb = openpyxl.Workbook()
    ws_detail = wb.active
    ws_detail.title = "Inventory Detail"
    ws_detail.append(detail_cols)
    for row in detail_rows:
        ws_detail.append([str(v) if v else "" for v in row])

    ws_summary = wb.create_sheet("Summary by Plant")
    ws_summary.append(summary_cols)
    for row in summary_rows:
        ws_summary.append([str(v) if v else "" for v in row])

    buf = io.BytesIO()
    wb.save(buf)
    return buf.getvalue()


# ── Email ─────────────────────────────────────────────────────────────────────

def send_email(to: str, subject: str, body: str, attachment: bytes, filename: str):
    host = os.environ.get("SMTP_HOST", "")
    port = int(os.environ.get("SMTP_PORT", "587"))
    user = os.environ.get("SMTP_USER", "")
    password = os.environ.get("SMTP_PASS", "")

    if not host:
        print("  SMTP not configured; skipping email.")
        return

    msg = MIMEMultipart()
    msg["Subject"] = subject
    msg["From"] = user
    msg["To"] = to
    msg.attach(MIMEText(body, "plain"))

    part = MIMEApplication(attachment, Name=filename)
    part["Content-Disposition"] = f'attachment; filename="{filename}"'
    msg.attach(part)

    with smtplib.SMTP(host, port) as srv:
        srv.starttls()
        srv.login(user, password)
        srv.sendmail(user, [to], msg.as_string())
    print(f"  Email sent to {to}")


# ── Main ──────────────────────────────────────────────────────────────────────

def main():
    p = argparse.ArgumentParser(description="CHACONTAINER Inventory Report")
    p.add_argument("--tenant-id", required=True)
    p.add_argument("--format", choices=["excel", "csv"], default="excel")
    p.add_argument("--output", default="./reports")
    p.add_argument("--email", default="")
    args = p.parse_args()

    if not DB_URL:
        sys.exit("DATABASE_URL not set")

    conn = psycopg2.connect(DB_URL)
    d_cols, d_rows = fetch(conn, ASSET_QUERY, args.tenant_id)
    s_cols, s_rows = fetch(conn, SUMMARY_QUERY, args.tenant_id)
    conn.close()

    print(f"Loaded {len(d_rows)} assets, {len(s_rows)} summary rows")

    ts = datetime.utcnow().strftime("%Y%m%d_%H%M")
    Path(args.output).mkdir(parents=True, exist_ok=True)

    if args.format == "excel":
        data = to_excel_bytes(d_cols, d_rows, s_cols, s_rows)
        filename = f"inventory_{args.tenant_id[:8]}_{ts}.xlsx"
        mime = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
    else:
        data = to_csv_bytes(d_cols, d_rows)
        filename = f"inventory_{args.tenant_id[:8]}_{ts}.csv"
        mime = "text/csv"

    filepath = os.path.join(args.output, filename)
    with open(filepath, "wb") as f:
        f.write(data)
    print(f"Saved: {filepath} ({len(data):,} bytes)")

    if args.email:
        send_email(
            to=args.email,
            subject=f"Inventory Report — {ts}",
            body=f"CHACONTAINER inventory snapshot attached.\nAssets: {len(d_rows)}\nGenerated: {ts} UTC",
            attachment=data,
            filename=filename,
        )


if __name__ == "__main__":
    main()

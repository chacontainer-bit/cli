#!/usr/bin/env python3
"""
CHACONTAINER - Late Shipment Alert Worker
Runs on a cron (every 30 min). Queries overdue shipments and fires
Make.com webhooks for each tenant's automation scenario.

Cron: */30 * * * * /usr/bin/python3 /opt/chacontainer/scripts/late_shipment_alert.py

Environment:
    DATABASE_URL
    MAKE_WEBHOOK_URL_TEMPLATE  (contains {tenant_slug} placeholder, optional)
"""

import os
import json
import sys
from datetime import datetime, timezone
from typing import Any

import psycopg2
import requests

DB_URL = os.environ.get("DATABASE_URL", "")
SLACK_WEBHOOK = os.environ.get("SLACK_WEBHOOK_URL", "")


def get_overdue_shipments(conn) -> list[dict[str, Any]]:
    with conn.cursor() as cur:
        cur.execute("""
            SELECT
                s.id,
                s.tenant_id,
                s.reference,
                s.scheduled_delivery,
                s.status,
                s.carrier_name,
                s.carrier_ref,
                c.name  AS client_name,
                c.id    AS client_id,
                t.slug  AS tenant_slug,
                t.settings->>'make_late_shipment_webhook' AS webhook_url
            FROM shipments s
            JOIN clients  c ON c.id = s.client_id
            JOIN tenants  t ON t.id = s.tenant_id
            WHERE s.scheduled_delivery < NOW()
              AND s.status NOT IN ('delivered', 'cancelled')
              AND t.status = 'active'
            ORDER BY s.scheduled_delivery ASC
        """)
        cols = [d[0] for d in cur.description]
        return [dict(zip(cols, row)) for row in cur.fetchall()]


def fire_webhook(url: str, payload: dict):
    try:
        r = requests.post(url, json=payload, timeout=10)
        r.raise_for_status()
    except Exception as e:
        print(f"  Webhook error ({url[:40]}...): {e}", file=sys.stderr)


def post_slack_summary(shipments: list[dict]):
    if not SLACK_WEBHOOK or not shipments:
        return
    text = f":warning: *{len(shipments)} late shipments* detected at {datetime.now(timezone.utc).strftime('%H:%M UTC')}\n"
    for s in shipments[:5]:
        delta = datetime.now(timezone.utc) - s["scheduled_delivery"].replace(tzinfo=timezone.utc)
        days = delta.days
        text += f"• `{s['reference']}` — {s['client_name']} — {days}d late ({s['status']})\n"
    if len(shipments) > 5:
        text += f"_...and {len(shipments)-5} more_"
    requests.post(SLACK_WEBHOOK, json={"text": text}, timeout=10)


def main():
    if not DB_URL:
        sys.exit("DATABASE_URL not set")

    conn = psycopg2.connect(DB_URL)
    shipments = get_overdue_shipments(conn)
    conn.close()

    print(f"[{datetime.utcnow().isoformat()}] Found {len(shipments)} overdue shipments")

    fired = 0
    for s in shipments:
        webhook_url = s.get("webhook_url")
        if not webhook_url:
            continue

        days_late = (datetime.now(timezone.utc) - s["scheduled_delivery"].replace(tzinfo=timezone.utc)).days
        payload = {
            "event": "late_shipment",
            "tenant_id": str(s["tenant_id"]),
            "shipment_id": str(s["id"]),
            "reference": s["reference"],
            "client_name": s["client_name"],
            "client_id": str(s["client_id"]),
            "status": s["status"],
            "carrier": s.get("carrier_name"),
            "days_late": days_late,
            "scheduled_delivery": s["scheduled_delivery"].isoformat(),
            "detected_at": datetime.utcnow().isoformat(),
        }
        fire_webhook(webhook_url, payload)
        fired += 1
        print(f"  → {s['reference']} ({days_late}d late) → webhook fired")

    post_slack_summary(shipments)
    print(f"Fired {fired} webhooks.")


if __name__ == "__main__":
    main()

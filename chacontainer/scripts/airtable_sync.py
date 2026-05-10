#!/usr/bin/env python3
"""
CHACONTAINER - Airtable Bidirectional Sync
Pulls from Airtable and pushes to CHACONTAINER API (or vice versa).

Usage:
    python airtable_sync.py --direction push|pull \
        --entity clients|assets|shipments \
        [--tenant-id <id>] [--dry-run]

Environment variables:
    AIRTABLE_API_KEY
    AIRTABLE_BASE_ID
    CHACONTAINER_API_URL
    CHACONTAINER_API_TOKEN
"""

import argparse
import json
import os
import sys
import time
from typing import Any

import requests


# ── Config ────────────────────────────────────────────────────────────────────

AIRTABLE_API = "https://api.airtable.com/v0"
AIRTABLE_KEY = os.environ.get("AIRTABLE_API_KEY", "")
AIRTABLE_BASE = os.environ.get("AIRTABLE_BASE_ID", "")
CC_API = os.environ.get("CHACONTAINER_API_URL", "http://localhost:8080")
CC_TOKEN = os.environ.get("CHACONTAINER_API_TOKEN", "")


# ── Airtable helpers ──────────────────────────────────────────────────────────

def at_headers() -> dict:
    return {"Authorization": f"Bearer {AIRTABLE_KEY}", "Content-Type": "application/json"}


def at_list(table: str, formula: str = "") -> list[dict]:
    records, offset = [], None
    while True:
        params: dict[str, Any] = {"pageSize": 100}
        if formula:
            params["filterByFormula"] = formula
        if offset:
            params["offset"] = offset
        r = requests.get(f"{AIRTABLE_API}/{AIRTABLE_BASE}/{table}",
                         headers=at_headers(), params=params, timeout=20)
        r.raise_for_status()
        data = r.json()
        records.extend(data.get("records", []))
        offset = data.get("offset")
        if not offset:
            break
        time.sleep(0.2)
    return records


def at_upsert(table: str, records: list[dict]) -> list[dict]:
    """Upsert in batches of 10."""
    results = []
    for i in range(0, len(records), 10):
        batch = records[i:i + 10]
        r = requests.patch(
            f"{AIRTABLE_API}/{AIRTABLE_BASE}/{table}",
            headers=at_headers(),
            json={"records": batch},
            timeout=20,
        )
        r.raise_for_status()
        results.extend(r.json().get("records", []))
        time.sleep(0.2)
    return results


# ── CHACONTAINER API helpers ──────────────────────────────────────────────────

def cc_headers() -> dict:
    return {"Authorization": f"Bearer {CC_TOKEN}", "Content-Type": "application/json"}


def cc_list(entity: str) -> list[dict]:
    results, page = [], 1
    while True:
        r = requests.get(f"{CC_API}/api/v1/{entity}",
                         headers=cc_headers(),
                         params={"page": page, "per_page": 100},
                         timeout=20)
        r.raise_for_status()
        data = r.json()
        items = data.get("data", [])
        results.extend(items)
        if len(items) < 100:
            break
        page += 1
    return results


def cc_patch(entity: str, item_id: str, payload: dict) -> dict:
    r = requests.patch(f"{CC_API}/api/v1/{entity}/{item_id}",
                       headers=cc_headers(), json=payload, timeout=20)
    r.raise_for_status()
    return r.json()


# ── Push: CHACONTAINER → Airtable ─────────────────────────────────────────────

PUSH_FIELD_MAP = {
    "clients": lambda c: {
        "Code": c.get("code"),
        "Name": c.get("name"),
        "Type": c.get("type"),
        "Status": c.get("status"),
        "Tax ID": c.get("tax_id"),
        "Credit Limit": c.get("credit_limit"),
        "Currency": c.get("currency"),
        "CHACONTAINER ID": c.get("id"),
    },
    "assets": lambda a: {
        "Code": a.get("code"),
        "Type": a.get("type"),
        "Status": a.get("status"),
        "Weight (kg)": a.get("weight_kg"),
        "Plant ID": a.get("plant_id"),
        "Last Scan": a.get("last_scan_at"),
        "CHACONTAINER ID": a.get("id"),
    },
    "shipments": lambda s: {
        "Reference": s.get("reference"),
        "Status": s.get("status"),
        "Mode": s.get("mode"),
        "Carrier": s.get("carrier_name"),
        "Scheduled Delivery": s.get("scheduled_delivery"),
        "Total Items": s.get("total_items"),
        "CHACONTAINER ID": s.get("id"),
    },
}

TABLE_MAP = {"clients": "Clients", "assets": "Assets", "shipments": "Shipments"}


def push(entity: str, dry_run: bool):
    items = cc_list(entity)
    print(f"  Fetched {len(items)} {entity} from CHACONTAINER")
    mapper = PUSH_FIELD_MAP[entity]
    records = [{"fields": mapper(item)} for item in items]

    if dry_run:
        print(f"  [DRY RUN] Would upsert {len(records)} records to Airtable:{TABLE_MAP[entity]}")
        return

    result = at_upsert(TABLE_MAP[entity], records)
    print(f"  Upserted {len(result)} records to Airtable:{TABLE_MAP[entity]}")


# ── Pull: Airtable → CHACONTAINER ─────────────────────────────────────────────

def pull(entity: str, dry_run: bool):
    records = at_list(TABLE_MAP[entity])
    print(f"  Fetched {len(records)} records from Airtable:{TABLE_MAP[entity]}")

    updated = 0
    for r in records:
        cc_id = r["fields"].get("CHACONTAINER ID")
        if not cc_id:
            continue
        patch = {}
        if entity == "clients":
            if status := r["fields"].get("Status"):
                patch["status"] = status.lower()
            if limit := r["fields"].get("Credit Limit"):
                patch["credit_limit"] = limit
        elif entity == "assets":
            if status := r["fields"].get("Status"):
                patch["status"] = status.lower()
        elif entity == "shipments":
            if status := r["fields"].get("Status"):
                patch["status"] = status.lower()

        if not patch:
            continue

        if dry_run:
            print(f"  [DRY RUN] Would patch {entity}/{cc_id}: {patch}")
        else:
            try:
                cc_patch(entity, cc_id, patch)
                updated += 1
            except Exception as e:
                print(f"  ERROR patching {entity}/{cc_id}: {e}", file=sys.stderr)

    print(f"  Applied {updated} updates to CHACONTAINER")


# ── Main ──────────────────────────────────────────────────────────────────────

def main():
    p = argparse.ArgumentParser(description="CHACONTAINER ↔ Airtable Sync")
    p.add_argument("--direction", choices=["push", "pull"], required=True)
    p.add_argument("--entity", choices=["clients", "assets", "shipments"], required=True)
    p.add_argument("--dry-run", action="store_true")
    args = p.parse_args()

    if not AIRTABLE_KEY or not AIRTABLE_BASE:
        sys.exit("Set AIRTABLE_API_KEY and AIRTABLE_BASE_ID")
    if not CC_TOKEN:
        sys.exit("Set CHACONTAINER_API_TOKEN")

    print(f"\n{'='*50}")
    print(f"CHACONTAINER ↔ Airtable Sync")
    print(f"Direction: {args.direction}  Entity: {args.entity}")
    print(f"{'='*50}")

    if args.direction == "push":
        push(args.entity, args.dry_run)
    else:
        pull(args.entity, args.dry_run)

    print("Done.")


if __name__ == "__main__":
    main()

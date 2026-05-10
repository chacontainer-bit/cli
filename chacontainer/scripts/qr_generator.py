#!/usr/bin/env python3
"""
CHACONTAINER - Bulk QR Code Generator
Generates signed QR codes for assets and outputs PNG files + CSV manifest.

Usage:
    python qr_generator.py --tenant <tenant_id> --secret <secret> \
        --input assets.csv --output ./qr_codes/

CSV format: id,code,type,plant_id
"""

import argparse
import csv
import hashlib
import hmac
import json
import os
import sys
from datetime import datetime
from pathlib import Path

try:
    import qrcode
    from PIL import Image, ImageDraw, ImageFont
except ImportError:
    sys.exit("Install dependencies: pip install qrcode[pil] Pillow")


def generate_payload(tenant_id: str, asset_id: str, asset_code: str, secret: str) -> str:
    payload = f"{tenant_id}|{asset_id}|{asset_code}|1"
    sig = hmac.new(secret.encode(), payload.encode(), hashlib.sha256).hexdigest()[:12]
    return f"{payload}|{sig}"


def make_qr_image(payload: str, asset_code: str, asset_type: str) -> Image.Image:
    qr = qrcode.QRCode(
        version=3,
        error_correction=qrcode.constants.ERROR_CORRECT_H,
        box_size=10,
        border=2,
    )
    qr.add_data(payload)
    qr.make(fit=True)

    img = qr.make_image(fill_color="black", back_color="white").convert("RGB")
    w, h = img.size

    # Label strip below QR
    label_height = 50
    canvas = Image.new("RGB", (w, h + label_height), "white")
    canvas.paste(img, (0, 0))

    draw = ImageDraw.Draw(canvas)
    label = f"{asset_code} | {asset_type.upper()}"
    draw.text((10, h + 8), label, fill="black")
    draw.text((10, h + 28), "CHACONTAINER", fill="#888888")

    return canvas


def process(tenant_id: str, secret: str, input_csv: str, output_dir: str):
    Path(output_dir).mkdir(parents=True, exist_ok=True)
    manifest = []

    with open(input_csv, newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f)
        for row in reader:
            asset_id = row["id"].strip()
            code = row["code"].strip()
            asset_type = row.get("type", "asset").strip()

            payload = generate_payload(tenant_id, asset_id, code, secret)
            img = make_qr_image(payload, code, asset_type)

            filename = f"{code}.png"
            filepath = os.path.join(output_dir, filename)
            img.save(filepath, "PNG", dpi=(300, 300))

            manifest.append({
                "asset_id": asset_id,
                "code": code,
                "type": asset_type,
                "qr_payload": payload,
                "filename": filename,
                "generated_at": datetime.utcnow().isoformat(),
            })
            print(f"  ✓ {code} → {filepath}")

    manifest_path = os.path.join(output_dir, "manifest.json")
    with open(manifest_path, "w") as f:
        json.dump(manifest, f, indent=2)
    print(f"\nDone. {len(manifest)} codes. Manifest: {manifest_path}")


def main():
    parser = argparse.ArgumentParser(description="CHACONTAINER QR Generator")
    parser.add_argument("--tenant", required=True)
    parser.add_argument("--secret", required=True)
    parser.add_argument("--input", required=True, help="CSV with columns: id,code,type")
    parser.add_argument("--output", default="./qr_codes")
    args = parser.parse_args()

    process(args.tenant, args.secret, args.input, args.output)


if __name__ == "__main__":
    main()

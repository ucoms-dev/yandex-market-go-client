#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "${ROOT_DIR}"

./scripts/check-duplicates.sh
go build ./...

echo "Verification passed"

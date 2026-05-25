#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "${ROOT_DIR}"

if command -v rg >/dev/null 2>&1; then
  mapfile -t duplicates < <(
    rg -n --glob '*.go' '^type[[:space:]]+([A-Za-z0-9_]+)[[:space:]]+' \
      | sed -E 's/^.*type[[:space:]]+([A-Za-z0-9_]+)[[:space:]]+.*/\1/' \
      | sort \
      | uniq -d
  )
else
  mapfile -t duplicates < <(
    grep -R -n -E --include='*.go' '^type[[:space:]]+[A-Za-z0-9_]+[[:space:]]' . \
      | sed -E 's/^.*type[[:space:]]+([A-Za-z0-9_]+)[[:space:]]+.*/\1/' \
      | sort \
      | uniq -d
  )
fi

if [[ "${#duplicates[@]}" -gt 0 ]]; then
  echo "Duplicate type names detected:" >&2
  printf ' - %s\n' "${duplicates[@]}" >&2
  exit 1
fi

echo "No duplicate type names found"

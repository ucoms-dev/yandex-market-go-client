#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CONFIG_FILE="${ROOT_DIR}/generation.yaml"

read_config_value() {
  local key="$1"

  local line
  if command -v rg >/dev/null 2>&1; then
    line="$(rg -N "^[[:space:]]*${key}:[[:space:]]*" "${CONFIG_FILE}" | head -n1 || true)"
  else
    line="$(grep -E "^[[:space:]]*${key}:[[:space:]]*" "${CONFIG_FILE}" | head -n1 || true)"
  fi
  if [[ -z "${line}" ]]; then
    echo "Error: key '${key}' not found in ${CONFIG_FILE}" >&2
    exit 1
  fi

  local value
  value="$(printf '%s\n' "${line}" | sed -E 's/^[^:]+:[[:space:]]*//; s/[[:space:]]+$//')"
  value="${value#\"}"
  value="${value%\"}"

  printf '%s\n' "${value}"
}

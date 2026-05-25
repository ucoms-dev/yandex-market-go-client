#!/usr/bin/env bash

set -euo pipefail

source "$(cd "$(dirname "$0")" && pwd)/lib-config.sh"

upstream_repo="$(read_config_value upstream_repo)"
upstream_ref="$(read_config_value upstream_ref)"
upstream_dir_rel="$(read_config_value upstream_dir)"
local_spec_dir_rel="$(read_config_value local_dir)"
local_spec_entry_rel="$(read_config_value local_path)"

local_spec_dir="${ROOT_DIR}/${local_spec_dir_rel}"
local_spec_entry="${ROOT_DIR}/${local_spec_entry_rel}"

tmp_dir="$(mktemp -d)"
cleanup() {
  rm -rf "${tmp_dir}"
}
trap cleanup EXIT

echo "Syncing OpenAPI tree from ${upstream_repo}@${upstream_ref}:${upstream_dir_rel}"

git clone --depth 1 --branch "${upstream_ref}" "${upstream_repo}" "${tmp_dir}/upstream" >/dev/null 2>&1

upstream_spec_dir="${tmp_dir}/upstream/${upstream_dir_rel}"
if [[ ! -d "${upstream_spec_dir}" ]]; then
  echo "Error: upstream spec directory not found: ${upstream_dir_rel}" >&2
  exit 1
fi

if [[ ! -f "${upstream_spec_dir}/openapi.yaml" ]]; then
  echo "Error: upstream openapi.yaml not found in ${upstream_dir_rel}" >&2
  exit 1
fi

if [[ ! -d "${upstream_spec_dir}/paths" || ! -d "${upstream_spec_dir}/components" ]]; then
  echo "Error: upstream spec tree is incomplete (paths/components missing)" >&2
  exit 1
fi

mkdir -p "${local_spec_dir}"
find "${local_spec_dir}" -mindepth 1 -maxdepth 1 -exec rm -rf {} +
cp -a "${upstream_spec_dir}/." "${local_spec_dir}/"

if [[ ! -f "${local_spec_entry}" ]]; then
  echo "Error: local spec entry not found after sync: ${local_spec_entry_rel}" >&2
  exit 1
fi

if ! grep -q '^openapi:' "${local_spec_entry}"; then
  echo "Error: ${local_spec_entry_rel} does not look like OpenAPI YAML" >&2
  exit 1
fi

echo "Updated ${local_spec_dir_rel}/ from upstream"

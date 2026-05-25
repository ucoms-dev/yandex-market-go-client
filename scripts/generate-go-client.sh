#!/usr/bin/env bash

set -euo pipefail

source "$(cd "$(dirname "$0")" && pwd)/lib-config.sh"

generator_version="$(read_config_value generator_version)"
language="$(read_config_value language)"
spec_rel="$(read_config_value local_path)"
local_spec_dir_rel="$(read_config_value local_dir)"
out_dir_rel="$(read_config_value output_dir)"
package_name="$(read_config_value package_name)"
package_version="$(read_config_value package_version)"
struct_prefix="$(read_config_value struct_prefix)"
enum_class_prefix="$(read_config_value enum_class_prefix)"
skip_validate_spec_default="$(read_config_value skip_validate_spec)"
minimal_update="$(read_config_value minimal_update)"
global_properties="$(read_config_value global_properties)"

# Allow explicit runtime override. CI should keep validation enabled.
skip_validate_spec="${SKIP_SPEC_VALIDATE:-${skip_validate_spec_default}}"

spec_path="${ROOT_DIR}/${spec_rel}"
if [[ ! -f "${spec_path}" ]]; then
  echo "Error: spec file not found: ${spec_path}" >&2
  exit 1
fi

module_path="$(awk '/^module[[:space:]]+/ {print $2; exit}' "${ROOT_DIR}/go.mod")"
git_user_id="ucoms-dev"
git_repo_id="yandex-market-go-client"
if [[ "${module_path}" =~ ^github\.com/([^/]+)/([^/]+)$ ]]; then
  git_user_id="${BASH_REMATCH[1]}"
  git_repo_id="${BASH_REMATCH[2]}"
fi

is_docker="${OPENAPI_GENERATOR_DOCKER:-1}"

generator_cmd=()
if [[ "${is_docker}" == "1" ]]; then
  if ! command -v docker >/dev/null 2>&1; then
    echo "Error: docker is required when OPENAPI_GENERATOR_DOCKER=1" >&2
    exit 1
  fi
  generator_cmd=(
    docker run --rm
    -u "$(id -u):$(id -g)"
    -v "${ROOT_DIR}:/local"
    "openapitools/openapi-generator-cli:v${generator_version}"
  )
else
  if command -v openapi-generator-cli >/dev/null 2>&1; then
    generator_cmd=(openapi-generator-cli)
  elif command -v openapi-generator >/dev/null 2>&1; then
    generator_cmd=(openapi-generator)
  else
    echo "Error: openapi-generator-cli/openapi-generator not found" >&2
    exit 1
  fi
fi

generator_path() {
  local rel_path="$1"
  if [[ "${is_docker}" == "1" ]]; then
    printf '/local/%s\n' "${rel_path}"
  else
    printf '%s/%s\n' "${ROOT_DIR}" "${rel_path}"
  fi
}

cleanup_generated_files() {
  local files_list="${ROOT_DIR}/.openapi-generator/FILES"

  if [[ -f "${files_list}" ]]; then
    echo "Cleaning files from .openapi-generator/FILES"
    while IFS= read -r rel_path; do
      [[ -z "${rel_path}" ]] && continue

      case "${rel_path}" in
        "${spec_rel}"|"${local_spec_dir_rel}"/*|".openapi-generator/FILES"|".openapi-generator/VERSION"|".openapi-generator-ignore")
          continue
          ;;
      esac

      local target="${ROOT_DIR}/${rel_path}"
      if [[ -e "${target}" || -L "${target}" ]]; then
        rm -rf "${target}"
      fi
    done < "${files_list}"
    return
  fi

  echo "No .openapi-generator/FILES found, using fallback cleanup"
  rm -f "${ROOT_DIR}"/api_*.go "${ROOT_DIR}"/model_*.go
  rm -f "${ROOT_DIR}"/client.go "${ROOT_DIR}"/configuration.go "${ROOT_DIR}"/response.go "${ROOT_DIR}"/utils.go
  rm -f "${ROOT_DIR}"/README.md "${ROOT_DIR}"/go.mod "${ROOT_DIR}"/go.sum "${ROOT_DIR}"/git_push.sh "${ROOT_DIR}"/.travis.yml

  if [[ -d "${ROOT_DIR}/docs" ]]; then
    find "${ROOT_DIR}/docs" -maxdepth 1 -type f -name '*.md' ! -name 'CONTRIBUTING.md' -delete
  fi
}

spec_for_generator="$(generator_path "${spec_rel}")"
out_dir_for_generator="$(generator_path "${out_dir_rel}")"

if [[ "${skip_validate_spec}" != "true" ]]; then
  echo "Validating OpenAPI spec ${spec_rel}"
  "${generator_cmd[@]}" validate -i "${spec_for_generator}"
else
  echo "Skipping spec validation (SKIP_SPEC_VALIDATE=${skip_validate_spec})"
fi

cleanup_generated_files

echo "Generating ${language} client from ${spec_rel}"

cmd=(
  "${generator_cmd[@]}"
  generate
  -i "${spec_for_generator}"
  -g "${language}"
  -o "${out_dir_for_generator}"
  --git-user-id "${git_user_id}"
  --git-repo-id "${git_repo_id}"
  --global-property "${global_properties}"
  --additional-properties "packageName=${package_name},packageVersion=${package_version},structPrefix=${struct_prefix},enumClassPrefix=${enum_class_prefix}"
)

if [[ "${minimal_update}" == "true" ]]; then
  cmd+=(--minimal-update)
fi

if [[ "${skip_validate_spec}" == "true" ]]; then
  cmd+=(--skip-validate-spec)
fi

"${cmd[@]}"

go mod tidy >/dev/null
go fmt ./... >/dev/null

echo "Generation complete"

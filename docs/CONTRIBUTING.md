# Contributing

## Duplicate Structs/Types

This client is generated via OpenAPI Generator, which creates:
- Model files: `model_*.go` (canonical request/response DTOs)
- API service files: `api_*.go` (service methods and per-endpoint request builder structs)

Per-service request builder structs with similar purposes (e.g., `FbsAddOffersToArchiveRequest`, `FbyAddOffersToArchiveRequest`) are not duplicates — they are service-specific builders that reference a single model type (e.g., `AddOffersToArchiveRequest`).

### How to check for actual duplicates

- Duplicate struct/type names across the package (would fail build):
  ```bash
  rg -n "^type\s+([A-Za-z0-9_]+)\s+struct\b" \
    | sed -E 's/^.*type\s+([A-Za-z0-9_]+)\s+struct.*/\1/' \
    | sort | uniq -d
  ```

- Build verification:
  ```bash
  go build ./...
  ```

### If a duplicate is introduced

1. Keep the canonical definition in `model_*.go` when it represents a schema from OpenAPI components.
2. Remove any accidental duplicate struct with the same name from `api/*.go` or other files.
3. Update references/imports to use the canonical type.
4. Re-run `go build` to ensure no conflicts remain.

### Note on API field drift

If runtime JSON errors like `json: unknown field "logisticPointId"` appear, the server schema likely changed. Regenerate the client from the updated OpenAPI spec (or add the missing field to the relevant model) rather than duplicating types.


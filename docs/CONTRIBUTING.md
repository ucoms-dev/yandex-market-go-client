# Contributing

## Source of truth

The client is generated from OpenAPI. Do not hand-edit generated API/model files.

- Spec source config: `generation.yaml`
- Local spec tree: `api/` (`openapi.yaml`, `paths/`, `components/`)
- OpenAPI generator version pin: `openapitools.json`

## Update and regenerate flow

Use the same sequence locally that CI uses:

```bash
bash ./scripts/sync-spec.sh
bash ./scripts/generate-go-client.sh
bash ./scripts/verify.sh
```

- `sync-spec.sh` syncs the full upstream OpenAPI tree into `api/`.
- `generate-go-client.sh` validates spec refs and regenerates the Go client.
- `verify.sh` runs duplicate-type detection and `go build ./...`.

Spec validation is enabled by default. For emergency local debugging only:

```bash
SKIP_SPEC_VALIDATE=true bash ./scripts/generate-go-client.sh
```

## Duplicate Structs/Types

This generator creates two families of types:

- `model_*.go`: canonical DTOs generated from OpenAPI schemas.
- `api_*.go`: API services and request-builder structs.

Request-builder structs with similar naming across services (for example,
`FbsAddOffersToArchiveRequest` and `FbyAddOffersToArchiveRequest`) are expected
and are not duplicates.

### How duplicates are checked

Use the maintained check:

```bash
bash ./scripts/check-duplicates.sh
```

If a real duplicate appears, keep the canonical model type and remove the
accidental redeclaration.

## CI automation

- `update-spec` workflow (daily + manual): updates spec and opens/updates a PR.
- `update-spec` optionally generates `CHANGELOG.md` via `openai codex` headless (`codex exec`).
- Codex auth for CI uses ChatGPT login data only (no API key flow).
- Set `CODEX_AUTH_JSON` secret to full contents of `~/.codex/auth.json` from a machine where `codex login` was done with ChatGPT.
- `CODEX_ACCESS_TOKEN` secret is supported as fallback for `codex login --with-access-token`.
- If no `CODEX_AUTH_JSON` or `CODEX_ACCESS_TOKEN` is configured, spec update still runs and PR is created without AI changelog.
- `generate-client` workflow (push to `main` + manual): regenerates client,
  verifies, and commits generated changes back to `main` when needed.

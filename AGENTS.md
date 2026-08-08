# TicketR — agent instructions

## Commands

```sh
# Start local DB (PostgreSQL on 5432)
docker compose up -d

# Run migrations (via goose)
goose up

# Seed test data
make seed

# Dev server with Air hot-reload
air

# Regenerate sqlc code after editing sql/queries/
sqlc generate

# Run all tests (stdlib only)
go test ./...

# Reset DB (down then up)
make db-reset
```

`PORT` defaults to `8080`. DSN is read from `GOOSE_DBSTRING` (reused from goose config). `.env` is auto-loaded via `_ "github.com/joho/godotenv/autoload"`.

## Architecture

**Handler → Service → Repository** three-layer pattern. Each domain under `internal/` (movies, cities, theaters, screens, shows) has exactly: `types.go` (request DTOs with `validate` tags), `service.go` (interface + private impl), `handlers.go` (private struct, public handlers). All wired in `cmd/api/api.go` via constructor injection.

Repository layer is 100% **sqlc-generated** from `sql/queries/` into `internal/repository/`. After editing `.sql` query files, run `sqlc generate` to regenerate.

The Echo HTTPErrorHandler is `apiresponse.GlobalErrorResponse`. Every handler returns `c.JSON(status, apiresponse.ApiResponse{...})`. Known errors return `apiresponse.ApiError` (maps to `{success:false, errors:...}`); unknown errors get logged and return 500.

## Key details an agent likely misses

- **`GOOSE_DBSTRING` is the app DSN** — there is no separate `DATABASE_URL`.
- **Echo v5 uses pointer `*echo.Context`** (not `echo.Context`).
- **City context** — endpoints that scope by city use `X-City-Id` header + `CityContextMiddleware`; routes are grouped under `cityPublic` in `api.go`.
- **Auth JWT/password code exists** in `internal/auth/` but is **not wired into any route** — there is no auth middleware or protected endpoint yet.
- **Tests use only stdlib `testing`** (no testify), even though testify is a transitive dep.
- **No linter, typecheck, or CI config** exists in the repo.
- Use `db.ToNullString`, `db.ToPgUUID` etc. from `internal/db/utils.go` for nullable field conversion.
- **Validator** is `github.com/go-playground/validator/v10` via Echo's `c.Validate(req)`.
- Air hot-reload excludes `cmd/seed/` and `_test.go` files by default.

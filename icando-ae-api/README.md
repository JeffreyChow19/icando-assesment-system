## Goose
### Install Goose
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest

```
### Add to environment var (system variables)
- Get the path to your Go workspace
```bash
go env GOPATH
```
- Add to env var, `PATH_TO_WORKSPACE\bin` e.g. `C:\Users\Jeffrey Chow\go\bin`

### Create new migration sql script
- Navigate to sql
```bash
cd internal/migrations/sql
```
- Create script
```bash
goose create MIGRATION_NAME sql 
```

### Migrate up
- Restart db container to auto migrate up `make run-local`


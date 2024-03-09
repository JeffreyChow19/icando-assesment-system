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

## Seeder
- Run the server and database `make run`
- Usually the server will error on the first try because the database hasn't been initialized successfully. Don't panic! Just re-run the server `make run-server`
- After the server and database successfully started, run the seeder `make seed`. It is recommended to seed on an empty database.


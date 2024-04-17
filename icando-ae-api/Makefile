run-local:
	docker compose up -d
	go run cmd/server/server.go

reset-local: ## Deletes the container and volumes for the local dev db
	docker container stop icando-ae-db
	docker container rm icando-ae-db

run-reset-local:
	reset-local
	run-local

run-worker:
	go run cmd/worker/worker.go

build:
	go build -o ./main ./cmd/server/server.go

run-database:
	docker-compose up -d database

run-redis:
	docker-compose up -d redis

run-server:
	go run cmd/server/server.go

run: run-database run-redis run-server

run-test-database:
	docker-compose up -d test_database

run-test:
	go test ./...

test: run-test-database run-test

seed:
	go run cmd/seeder/seeder.go

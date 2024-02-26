run-local:
	docker compose up -d
	go run cmd/server/server.go

reset-local: ## Deletes the container and volumes for the local dev db
	docker container stop icando-ae-db
	docker container rm icando-ae-db

run-reset-local:
	reset-local
	run-local

build:
	go build -o ./main ./cmd/server/server.go
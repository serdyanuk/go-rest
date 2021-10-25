ENV_FILE := .env
DSN := $(shell sed -n 's/^DSN=*\(.*\)/\1/p' ${ENV_FILE})
MIGRATE := docker run --rm -v $(shell pwd)/internal/app/store/migrations:/migrations --network host migrate/migrate:v4.15.1 -path=/migrations/ -database "$(DSN)"

run:
	go run ./cmd/app/main.go

test:
	go test ./internal/...

db-start:
	docker-compose up -d

db-stop:
	docker-compose down

migrate-up:
	@$(MIGRATE) up

migrate-down:
	@$(MIGRATE) down

# migrate-reset:
# 	@$(MIGRATE) drop -f
# 	@$(MIGRATE) up
	
migrate-new:
	@read -p 'Enter new migration name: ' name; \
	$(MIGRATE) create -ext sql -dir /migrations/ $$name

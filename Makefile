include .env
export

# ==============================================================================
# Help

.PHONY: help
## help: shows this help message
help:
	@ echo "Usage: make [target]\n"
	@ sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==============================================================================
# Tests

.PHONY: test
## test: run unit tests
test: migrate-test-up
	@ go test -v ./... -count=1

.PHONY: coverage
## coverage: run unit tests and generate coverage report in html format
coverage: migrate-test-up
coverage:
	@ packages=$$(go list ./... | grep -v "cmd" | grep -v "validate"); \
	if [ -z "$$packages" ]; then \
		echo "No valid Go packages found"; \
		exit 1; \
	fi; \
	go test -coverpkg=$$(echo $$packages | tr ' ' ',') -coverprofile=coverage.out $$packages && go tool cover -html=coverage.out

# ==============================================================================
# DB

.PHONY: start-psql
## start-psql: starts psql instance
start-psql:
	@ docker-compose up $(POSTGRES_DATABASE_CONTAINER_NAME) -d
	@ echo "Waiting for Postgres to start..."
	@ until docker exec $(POSTGRES_DATABASE_CONTAINER_NAME) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)  -c "SELECT 1;" >/dev/null 2>&1; do \
		echo "Postgres not ready, sleeping for 5 seconds..."; \
		sleep 5; \
	done
	@ echo "Postgres is up and running."

.PHONY: stop-psql
## stop-psql: stops psql instance
stop-psql:
	@ docker-compose down $(POSTGRES_DATABASE_CONTAINER_NAME)

.PHONY: psql-console
## psql-console: opens psql terminal
psql-console: export PGPASSWORD=$(POSTGRES_PASSWORD)
psql-console:
	@ docker exec -it $(POSTGRES_DATABASE_CONTAINER_NAME) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

.PHONY: start-test-psql
## start-test-psql: starts psql test instance
start-test-psql:
	@ docker-compose up $(POSTGRES_TEST_DATABASE_CONTAINER_NAME) -d
	@ echo "Waiting for Postgres to start..."
	@ until docker exec $(POSTGRES_TEST_DATABASE_CONTAINER_NAME) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)  -c "SELECT 1;" >/dev/null 2>&1; do \
		echo "Postgres not ready, sleeping for 5 seconds..."; \
		sleep 5; \
	done
	@ echo "Postgres is up and running."

.PHONY: stop-test-psql
## stop-test-psql: stops test psql instance
stop-test-psql:
	@ docker-compose down $(POSTGRES_TEST_DATABASE_CONTAINER_NAME)

.PHONY: psql-test-console
## psql-test-console: opens test psql terminal
psql-test-console: export PGPASSWORD=$(POSTGRES_PASSWORD)
psql-test-console:
	@ docker exec -it $(POSTGRES_TEST_DATABASE_CONTAINER_NAME) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

# ==============================================================================
# DB migrations

.PHONY: create-migration
## create-migration: creates a migration file
create-migration:
	@ if [ -z "$(NAME)" ]; then echo >&2 "please set the name of the migration via the variable NAME"; exit 2; fi
	@ docker run --rm -v `pwd`/db/migrations:/migrations migrate/migrate create -ext sql -dir /migrations -seq $(NAME)

.PHONY: migrate-up
## migrate-up: runs migrations up to N version (optional)
migrate-up: start-psql
	@ docker run --rm --network $(POSTGRES_DATABASE_CONTAINER_NETWORK_NAME) -v `pwd`/db/migrations:/migrations migrate/migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_DATABASE_CONTAINER_NAME):5432/$(POSTGRES_DB)?sslmode=disable' -path /migrations up $(N)

.PHONY: migrate-down
## migrate-down: runs migrations down to N version (optional)
migrate-down:
	@ if [ -z "$(N)" ]; then \
		docker run --rm --network $(POSTGRES_DATABASE_CONTAINER_NETWORK_NAME) -v `pwd`/db/migrations:/migrations migrate/migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_DATABASE_CONTAINER_NAME):5432/$(POSTGRES_DB)?sslmode=disable' -path /migrations down -all; \
	else \
		docker run --rm --network $(POSTGRES_DATABASE_CONTAINER_NETWORK_NAME) -v `pwd`/db/migrations:/migrations migrate/migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_DATABASE_CONTAINER_NAME):5432/$(POSTGRES_DB)?sslmode=disable' -path /migrations down $(N); \
	fi

.PHONY: migrate-version
## migrate-version: shows current migration version number
migrate-version:
	@ docker run --rm --network $(POSTGRES_DATABASE_CONTAINER_NETWORK_NAME) -v `pwd`/db/migrations:/migrations migrate/migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_DATABASE_CONTAINER_NAME):5432/$(POSTGRES_DB)?sslmode=disable' -path /migrations version

.PHONY: migrate-force-version
## migrate-force-version: forces migrations to version V
migrate-force-version:
	@ if [ -z "$(V)" ]; then echo >&2 please set version via variable V; exit 2; fi
	@ docker run --rm --network $(POSTGRES_DATABASE_CONTAINER_NETWORK_NAME) -v `pwd`/db/migrations:/migrations migrate/migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_DATABASE_CONTAINER_NAME):5432/$(POSTGRES_DB)?sslmode=disable' -path /migrations force $(V)

.PHONY: migrate-test-up
## migrate-test-up: runs up N migrations on test db, N is optional (make migrate-up N=<desired_migration_number>)
migrate-test-up: start-test-psql
	@ docker run --rm --network $(POSTGRES_TEST_DATABASE_CONTAINER_NETWORK_NAME) -v `pwd`/db/migrations:/migrations migrate/migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_TEST_DATABASE_CONTAINER_NAME):5432/$(POSTGRES_DB)?sslmode=disable' -path /migrations up $(N)

# ==============================================================================
# Swagger

.PHONY: swagger
## swagger: generates api's documentation
swagger: 
	@ swagger generate spec -o doc/swagger.json --scan-models

.PHONY: swagger-ui
## swagger-ui: launches swagger ui
swagger-ui: swagger
	@ docker run --rm --name books-swagger-ui -p 80:8080 -e SWAGGER_JSON=/docs/swagger.json -v $(shell pwd)/doc:/docs swaggerapi/swagger-ui

# ==============================================================================
# App's execution

.PHONY: run
## run: runs the API
run: migrate-up
	@ if [ -z "$(PORT)" ]; then echo >&2 please set the desired port via the variable PORT; exit 2; fi
	@ go run cmd/main.go -p $(PORT)
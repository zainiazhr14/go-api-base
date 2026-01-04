SHELL := bash

DB_DRIVER=postgres
DB_DSN=postgres://postgres:Qwe123@@@localhost:5432/go-api-base?sslmode=disable
MIGRATIONS_DIR=./migrations

.PHONY: status up down reset create run

run:
	wgo run cmd/main.go
	
status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" status

up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" up

down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" down

redo:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" redo

reset:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" reset

ifndef name
create:
	@exit 1
else
create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql
endif

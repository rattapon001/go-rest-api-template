include .env

.PHONY: run
run:
	go run cmd/main.go

.PHONY: migrate-up
migrate-up:
	migrate -database 'postgres://postgres:123456@localhost:5432/test' -path ./migrations up

.PHONY: migrate-down
migrate-down:
	migrate -database 'postgres://postgres:123456@localhost:5432/test' -path ./migrations down 1

.PHONY: migrate-create
migrate-create: 
	@read -p "Enter the name of the new migration: " name; \
	migrate create -ext sql -dir ./migrations -seq $${name}
PORT := 5432

postgres:
		$(info Make: Starting PostgreSQL container.)
		@docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
		$(info Make: Creating PostgreSQL db.)
		docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
		$(info Make: Dropping PostgreSQL db.)
		docker exec -it postgres12 dropdb simple_bank

migrateup:
		$(info Make: Migrate Up PostgreSQL db.)
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
		$(info Make: Migrate Down PostgreSQL db.)
		migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown

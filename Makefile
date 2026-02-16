include .env
export

migrateup:
	migrate -path migrations -database "$(DATABASE_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DATABASE_URL)" -verbose down

createmigration:
	@if [ -z "$(name)" ]; then echo "Error: name is required. Usage: make createmigration name=<migration_name>"; exit 1; fi
	migrate create -ext sql -dir migrations -seq $(name)
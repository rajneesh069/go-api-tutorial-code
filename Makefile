include .env
export

migrateup:
	migrate -path migrations -database "$(DATABASE_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DATABASE_URL)" -verbose down

createmigration:
	migrate create -ext sql -dir migrations -seq "$(name)"
migrate-up:
	goose -dir ./db/migrations postgres "host=localhost user=bezgo dbname=url_shortner password=12345 sslmode=disable" up

migrate-create:
	goose -dir db/migrations create $(F) sql

run-pg:
	STORAGE_MODE=postgres ADDR=:8081 go run ./cmd
migrate-up:
	goose -dir ./db/migrations postgres "host=localhost user=bezgo dbname=url_shortner password= sslmode=disable" up

migrate-create:
	goose -dir db/migrations create $(F) sql

build:
	go build -o ff-api cmd/api/main.go

run:
	go run cmd/api/main.go --env-file=dev.env

migrate:
	go run cmd/migrate/main.go --env-file=dev.env

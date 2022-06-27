sh:
	docker-compose exec api sh

run api:
	docker-compose exec api go run main.go http

swagger:
	docker-compose exec api swag init -g infra/http/server.go

test:
	docker-compose exec api go test ./...
# Transactions (API + Postgresql + Prometheus + Grafana)

This is a project made for the pismo team to analyze. The project accepts creating an account, transaction and reading account infos.

# How to run
## For dev
```shell
cp .env.example .env
docker-compose up -d
make run api
```
## For production
First, edit docker-compose.yml on the api service change the dockerfile field to **Dockerfile**. Then you're all set.
We are using docker multistage builds to generate a production smallersize image. To run locally you can use the command:
```shell
docker-compose up -d
```
## Swagger
You can access swagger on: http://localhost:8080/swagger/index.html#/.
To regenerate the swagger page, run:
```shell
swag init -g infra/http/server.go 
```
## Migrations
Install goose
```shell
make sh
go install github.com/pressly/goose/v3/cmd/goose@latest
cd infra/database/migrations
goose postgres "host=db user=postgres dbname=transactions_api password=postgres sslmode=disable" up
```
## SQLc
SQLc is a codegen tool to handle database queries.
```shell
sqlc generate
```
# Concepts
## Clean architecture
This project relies heavily on the concepts presented by uncle Bob in his book "Clean Architecture: a craftsmen guide to software structure and design".
Clean arch allow us to decouple our apps from frameworks, more testable design, become independent of external agents.
The most common implementation of Clean architecture is Hexagonal architeture, which is used in this project.

### Entities
Encapsulates enterprises rules. Can be an object with rules or a set of structures or functions.

### Use cases
Application specific business rules. Implements all systems use cases.

### Presenters
Converts the data from the use cases to a convenient format for the GUI or client.

### Adapters
This layer encapsulates all the code necessary for the application to receive a communication from the exterior, like a http call for example.

### Repositories
Encapsulates all interactions with the database.
## Tests
To run all tests:
```shell
make run test
```
## Logging
All API related metrics are sent to Prometheus. These metrics can be seen at http://localhost:3000 (Grafana dashboard)


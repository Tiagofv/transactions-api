FROM golang:1.18-alpine3.15

WORKDIR /src

COPY . /src

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server

FROM alpine:3.15
EXPOSE 8080
WORKDIR /app
COPY --from=0 /src/server .
COPY --from=0 /src/docker-entrypoint.sh .

CMD ["/bin/sh", "docker-entrypoint.sh"]

FROM golang:1.18-buster as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /server .

FROM golang:1.18-alpine

COPY --from=builder ./server .

CMD ["./server"]

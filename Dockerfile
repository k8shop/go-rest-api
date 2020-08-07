FROM golang:1.14.7-alpine3.12

WORKDIR /app/go-rest-api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/go-rest-api ./cmd/go-rest-api/main.go

EXPOSE 8080

CMD ["./out/go-rest-api"]

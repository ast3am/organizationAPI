FROM golang:1.21.0

WORKDIR /app

COPY . .

ENV GIN_MODE=release

RUN go build -o myapp cmd/main.go

CMD ["./myapp"]
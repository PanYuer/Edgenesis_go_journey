FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o client main.go

CMD ["./client"]
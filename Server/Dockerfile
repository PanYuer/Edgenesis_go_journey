FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod init github.com/PanYuer/Edgenesis_go_journey
RUN go build -o server

CMD ["./server"]
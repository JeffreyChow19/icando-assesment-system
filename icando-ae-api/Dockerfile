FROM golang:1.20

WORKDIR /app

COPY .. .

RUN go build -o /app/bin/server ./cmd/server/server.go

EXPOSE 3000

CMD ["./bin/server"]

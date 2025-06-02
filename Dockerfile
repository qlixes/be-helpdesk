FROM golang:alpine

WORKDIR /app

VOLUME [ "/app" ]

RUN go mod download

EXPOSE 8000

CMD ["go","run","cmd/main.go"]

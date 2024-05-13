FROM golang:1.20.4-alpine

WORKDIR /app
COPY . /app

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]
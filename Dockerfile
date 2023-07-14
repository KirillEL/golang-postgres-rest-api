FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-postgres-rest-api

EXPOSE 8085

CMD ["/golang-postgres-rest-api"]
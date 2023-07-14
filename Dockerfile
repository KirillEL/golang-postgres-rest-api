FROM golang:alpine

WORKDIR /app

COPY go.* ./

ENV DB_HOST postgres
ENV DB_PORT 5432
ENV DB_DBNAME golangdb
ENV SSL_MODE disable
ENV DB_USERNAME postgres


RUN go mod download

COPY ./ ./

RUN go build ./cmd/app

EXPOSE 8085

CMD ["./app"]
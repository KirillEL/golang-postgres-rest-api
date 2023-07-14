FROM golang:alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY ./ ./

RUN go build ./cmd/app

EXPOSE 8085

CMD ["./app"]
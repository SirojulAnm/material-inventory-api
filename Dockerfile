FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN env CGO_ENABLED=0 go build -o binary

ENTRYPOINT ["/app/binary"]
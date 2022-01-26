FROM golang:alpine

COPY ./src/ /app
WORKDIR /app

RUN go mod download
RUN go build -o app main.go

ENTRYPOINT ["./app"]


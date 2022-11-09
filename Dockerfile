# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY web/* ./web/

RUN go build -o /financial-balance-web

EXPOSE 8080

CMD [ "/financial-balance-web" ]
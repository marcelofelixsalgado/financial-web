# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY pkg/* ./domain/
COPY web/* ./web/

RUN go build -o /financial-balance-web

EXPOSE 8080

CMD [ "/financial-balance-web" ]
# syntax=docker/dockerfile:1

FROM golang:1.19.4-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY .env ./

COPY api/ ./api/
COPY commons/ ./commons/
COPY pkg/ ./pkg/
COPY settings/ ./settings/
COPY version/ ./version/
COPY web/ ./web/

RUN go build -o /financial-web

EXPOSE 8080

ENTRYPOINT [ "financial-web" ]
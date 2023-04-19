# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

COPY api/ ./api/
COPY commons/ ./commons/
COPY pkg/ ./pkg/
COPY settings/ ./settings/
COPY version/ ./version/
COPY web/ ./web/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /financial-web

EXPOSE 8080

ENTRYPOINT [ "/financial-web" ]
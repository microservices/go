FROM golang:1.16-alpine as builder

ENV     GO111MODULE=on

WORKDIR /microservice

COPY    go.* .
RUN     go mod download

COPY    *.go .
RUN     CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app


FROM        scratch
COPY        --from=builder /microservice/app /app
ENTRYPOINT  ["/app"]

FROM golang:1.18-alpine as dev-env

WORKDIR /app

FROM dev-env as build-env

COPY . /app/

RUN CGO_ENABLED=0 go build -o /webhook cmd/main.go

FROM alpine:latest as runtime

COPY --from=build-env /webhook /usr/local/bin/webhook
RUN chmod +x /usr/local/bin/webhook

ENTRYPOINT ["webhook"]
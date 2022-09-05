# syntax=docker/dockerfile:1

##### Stage 1: Build application #####
FROM golang:1.18-alpine AS build

WORKDIR /build

RUN apk --no-cache add ca-certificates

ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download
COPY . .

RUN go build -o ./termit cmd/termit.go

##### Stage 2: Setup runtime #####
FROM gomicro/goose AS runtime

ENV CONFIG_PATH /app/config.yml

COPY ./config/config.yml ${CONFIG_PATH}
COPY ./migrations /app/migrations
COPY ./scripts /app/scripts

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/termit /app/termit

ENTRYPOINT [ "/app/scripts/entrypoint.sh" ]
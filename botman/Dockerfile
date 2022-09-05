# syntax=docker/dockerfile:1

##### Stage 1: Build application #####
FROM golang:1.18-alpine AS build

WORKDIR /build

RUN apk --no-cache add ca-certificates

ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download
COPY . .

RUN go build -o ./botman cmd/botman.go

##### Stage 2: Setup runtime #####
FROM scratch AS runtime

ENV CONFIG_PATH /app/config.yml

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/config/config.yml /app/config.yml
COPY --from=build /build/botman /app/botman

CMD [ "/app/botman" ]
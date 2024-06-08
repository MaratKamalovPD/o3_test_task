FROM golang:1.22-alpine3.19 as build

WORKDIR /var/backend

COPY cmd cmd
COPY internal internal
COPY migrations migrations
COPY .env .env

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod tidy
RUN go build -o main ./cmd/app/main.go 

FROM alpine:edge as prod
RUN apk add bash

WORKDIR /root
COPY --from=build /var/backend/main main
COPY --from=build /var/backend/internal/pkg/config/config.yaml ./internal/pkg/config/config.yaml
COPY --from=build /var/backend/migrations ./migrations
COPY --from=build /var/backend/.env ./.env

EXPOSE 8080
 
SHELL ["/bin/bash", "-c"]
ENTRYPOINT ./main
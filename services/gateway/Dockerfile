# https://container-solutions.com/faster-builds-in-docker-with-go-1-11/
FROM golang:1.12.0-alpine AS build_base

RUN mkdir -p /gateway
RUN apk add bash git
WORKDIR /gateway

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -mod=vendor ./gateway.go

FROM alpine AS gateway

COPY --from=build_base /gateway/gateway /bin/gateway
ENTRYPOINT ["/bin/gateway"]

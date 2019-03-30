# https://container-solutions.com/faster-builds-in-docker-with-go-1-11/
FROM golang:1.12.0-alpine AS build_base

RUN mkdir -p /auth
RUN apk add bash git
WORKDIR /auth

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -mod=vendor ./cmd/auth.go

FROM alpine AS auth

COPY --from=build_base /auth/auth /bin/auth
ENTRYPOINT ["/bin/auth"]

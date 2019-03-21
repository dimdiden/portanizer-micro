# https://container-solutions.com/faster-builds-in-docker-with-go-1-11/
FROM golang:1.12.0-alpine AS build_base

RUN mkdir -p /users
RUN apk add bash git
WORKDIR /users

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -mod=vendor ./cmd/users.go

FROM alpine AS workbook

COPY --from=build_base /users/users /bin/users
ENTRYPOINT ["/bin/users"]

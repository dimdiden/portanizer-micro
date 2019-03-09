# https://container-solutions.com/faster-builds-in-docker-with-go-1-11/
FROM golang:1.12.0-alpine AS build_base

RUN mkdir -p /users
RUN apk add bash git gcc g++ libc-dev
WORKDIR /users

# We want to populate the module cache based on the go.{mod,sum} files.
COPY ./go.mod .
COPY ./go.sum .

# Force the go compiler to use modules
ENV GO111MODULE=on

# Download all the related packages
RUN go mod download

# This image builds the users servise
FROM build_base AS users_builder
# Here we copy the rest of the source code
COPY . .
# And compile the project
RUN GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/users.go

#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine AS workbook
# Finally we copy the statically compiled Go binary.
COPY --from=users_builder /go/bin/users /bin/users
ENTRYPOINT ["/bin/users"]

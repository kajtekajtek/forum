# stage 1: building the binary
ARG GO_VERSION=1.24
FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && \
    apk add --no-cache git ca-certificates

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ARG BUILD_VERSION=latest
RUN go build -o server \
    -trimpath \
    -ldflags="-s -w -X 'main.Version=${BUILD_VERSION}'" \
    ./cmd/main.go

# stage 2: create the runtime image
FROM alpine:3.18 AS runtime

RUN apk add --no-cache ca-certificates

RUN addgroup -S app && adduser -S -G app app

WORKDIR /home/app

COPY --from=builder /app/server .

RUN chmod +x ./server

USER app

ARG PORT=8080
EXPOSE ${PORT}

ENTRYPOINT ["./server"]

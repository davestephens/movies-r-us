FROM golang:1.19-buster as builder

# build the application in a full fat Debian
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -v -o rest-api

# Use a slim Debian for prod, and copy in the built binary
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/rest-api /app/rest-api

EXPOSE 8080

ENV GIN_MODE=release

CMD ["/app/rest-api"]


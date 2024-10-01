FROM golang:1.23.1 AS builder

WORKDIR /app
COPY go.* .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o lito ./cmd/lito

# Copy the binary from the builder image to the final image
FROM alpine:3.20.3

COPY --from=builder /app/lito /app

RUN echo $'{\
  "admin": {\
    "enabled": false,\
    "port": 2024,\
    "api_key": ""\
  },\
  "proxy": {\
    "host": "0.0.0.0",\
    "http_port": 80,\
    "https_port": 443,\
    "enable_tls": false,\
    "tls_email": "",\
    "enable_https_redirect": true,\
    "config_path": "lito.json",\
    "storage": "json",\
    "cnames": null\
  }\
}' > lito.json

CMD ["/app"]

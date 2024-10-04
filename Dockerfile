FROM golang:1.23.1 AS builder

WORKDIR /app
COPY go.* .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o lito ./cmd/lito

# Copy the binary from the builder image to the final image
FROM alpine:3.20.3

COPY --from=builder /app/lito /app

RUN echo $'\
[admin]\n\
enabled = false\n\n\
[proxy]\n\
config_path = "lito.toml"\n\
enable_https_redirect = false\n\
enable_tls = false\n\
host = "0.0.0.0"\n\
http_port = 80\n\
https_port = 443\n\
storage = "toml"' > lito.toml

EXPOSE 80 443

ENTRYPOINT ["/app"]

CMD ["run"]

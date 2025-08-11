FROM golang:latest AS builder

# Install xcaddy
RUN go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest

# Build Caddy with the caddy-l4 plugin
RUN xcaddy build \
    --with github.com/mholt/caddy-l4

# Final image
FROM caddy:latest

COPY --from=builder /go/caddy /usr/bin/caddy

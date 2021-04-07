FROM golang:alpine3.13 AS builder
WORKDIR  /opt/oauth-proxy
COPY . .
RUN go build .

FROM alpine:3.13
COPY --from=builder /opt/oauth-proxy/oauth-proxy /usr/bin/oauth-proxy
RUN chmod +x /usr/bin/oauth-proxy
USER 1001
ENTRYPOINT ["/usr/bin/oauth-proxy"]

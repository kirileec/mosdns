FROM golang:latest AS builder
ARG CGO_ENABLED=0

COPY ./ /root/src/
WORKDIR /root/src/
RUN go build -ldflags "-s -w -X main.version=$(git describe --tags --long --always)" -trimpath -o mosdns

FROM alpine:latest

COPY --from=builder /root/src/mosdns /usr/bin/
COPY --from=builder /root/src/webui /usr/share/mosdns/webui
COPY docker-entrypoint.sh /usr/bin/docker-entrypoint.sh
WORKDIR /etc/mosdns

RUN apk add --no-cache ca-certificates && chmod +x /usr/bin/docker-entrypoint.sh

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["start"]

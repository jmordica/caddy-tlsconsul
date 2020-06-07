FROM golang:1.14 AS builder

WORKDIR /workspace
RUN echo 'package main\n\
import (\n\
caddycmd "github.com/caddyserver/caddy/v2/cmd"\n\
_ "github.com/caddyserver/caddy/v2/modules/standard"\n\
_ "github.com/pteich/caddy-tlsconsul"\n\
)\n\
func main() {\n\
caddycmd.Main()\n\
}' > main.go && \
          go mod init caddy && go get github.com/caddyserver/caddy/v2@master && go get && \
          CGO_ENABLED=0 go env -w GOPROXY="https://goproxy.io,direct" go build -trimpath -tags netgo -ldflags '-extldflags "-static" -s -w' -o /usr/bin/caddy


FROM caddy:2
LABEL maintainer="peter.teich@gmail.com"
LABEL description="Caddy 2 with integrated TLS Consul Storage plugin"
COPY --from=builder /usr/bin/caddy /usr/bin/caddy

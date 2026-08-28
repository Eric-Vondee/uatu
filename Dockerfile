FROM --platform=$BUILDPLATFORM golang:1.26.7-alpine3.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Cross-compile to whatever the deploy target is, so the image also builds
# correctly on an arm64 workstation.
ARG TARGETOS
ARG TARGETARCH
ENV CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH

# Output to /out, not the source tree: `-o server` would resolve to the
# existing server/ package directory and write the binary inside it.
RUN mkdir -p /out && \
    go build -ldflags="-s -w" -o /out/server ./cmd && \
    go build -ldflags="-s -w" -o /out/uatu-cli ./cmd/cli

FROM alpine:3.23 AS runtime
# RPC endpoints and the Postgres connection are reached over TLS.
RUN apk --no-cache add ca-certificates && \
    addgroup -S uatu && \
    adduser -S -G uatu -h /app uatu
WORKDIR /app
USER uatu

FROM runtime AS cli
COPY --from=builder --chown=uatu:uatu /out/uatu-cli .
ENTRYPOINT ["./uatu-cli"]

FROM runtime AS server
COPY --from=builder --chown=uatu:uatu /out/server .

EXPOSE 8080

CMD ["./server"]

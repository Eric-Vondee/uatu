FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS builder
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

FROM alpine:3.23
# RPC endpoints and the Postgres connection are reached over TLS.
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /out/server .

COPY --from=builder /out/uatu-cli .

EXPOSE 8080

CMD ["./server"]

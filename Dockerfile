FROM golang:1.20-alpine AS builder
ARG TARGETOS
ARG TARGETARCH

RUN apk --no-cache add ca-certificates

WORKDIR /
COPY services.notify services.notify
COPY services.shared ../services.shared
COPY keys keys 
WORKDIR /services.notify
ENV CGO_ENABLED=0
COPY ./services.notify/go.mod ./services.notify/go.sum ./
RUN  --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . . 
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main ./src/cmd/main.go

FROM scratch

ENV PORT 8080

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /services.notify/main .
COPY --from=builder /keys ./keys
COPY --from=builder /services.notify/src/locales ./src/locales
COPY --from=builder /services.notify/assets ./assets

EXPOSE $PORT

CMD ["/main"]
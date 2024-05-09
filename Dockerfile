FROM golang:1.22-alpine as builder
RUN apk add --no-cache git ca-certificates curl && \
    update-ca-certificates

WORKDIR /go/src/ntfy-bridge

RUN git clone https://github.com/r6m/ntfy-bridge .
RUN go build -o bin/ntfy-bridge .

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/ntfy-bridge/bin/ntfy-bridge /usr/local/bin/ntfy-bridge
ENTRYPOINT ["/usr/local/bin/ntfy-bridge"]
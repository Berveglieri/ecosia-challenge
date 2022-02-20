FROM golang:alpine AS builder

WORKDIR /app
ENV CGO_ENABLED=0
COPY . .

RUN go build -mod vendor -o /go/bin/webapp ./pkg/main.go

FROM alpine

WORKDIR /app
RUN apk add --no-cache --update ca-certificates
COPY --from=builder /go/bin/webapp /usr/bin/webapp


ENTRYPOINT ["/usr/bin/webapp"]
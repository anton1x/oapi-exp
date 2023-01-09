# Multi-stage build setup (https://docs.docker.com/develop/develop-images/multistage-build/)

# Stage 1 (to create a "build" image, ~850MB)
FROM golang:1.19 AS builder
RUN go version

COPY . /go/src/github.com/anton1x/oapi-exp/
WORKDIR /go/src/github.com/anton1x/oapi-exp/
RUN set -x && \
    go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app ./cmd/app/

# Stage 2 (to create a downsized "container executable", ~7MB)

# If you need SSL certificates for HTTPS, replace `FROM SCRATCH` with:
#
#   FROM alpine:3.7
#   RUN apk --no-cache add ca-certificates
#
FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/anton1x/oapi-exp/app .

EXPOSE 8089
ENTRYPOINT ["./app"]
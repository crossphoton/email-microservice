# Builder
FROM golang:1.14.6-alpine3.12 as builder

COPY . /app/
WORKDIR /app
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./ -a -installsuffix cgo /app/...


# Runner
FROM alpine:3.15.0
RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /app

COPY --from=builder /app/email-microservice /app/email-microservice
EXPOSE $PORT
EXPOSE $PROMETHEUS_PORT

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.6 && \
    wget -O /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

CMD [ "./email-microservice" ]
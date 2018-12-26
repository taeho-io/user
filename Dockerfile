FROM golang:1.11.4 as golang
WORKDIR /user
COPY . .
RUN go mod download
WORKDIR /user/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/user
WORKDIR /user
ARG test
RUN if [ "$test" = "true" ] ; then curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.12.4 ; fi
RUN if [ "$test" = "true" ] ; then make lint ; fi
RUN if [ "$test" = "true" ] ; then make test ; fi

FROM alpine:3.8
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
COPY --from=golang /go/bin /app
ENTRYPOINT ["/app/user"]

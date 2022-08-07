# Stage 1 - build executable in go container
FROM golang:latest as builder

WORKDIR $GOPATH/src/intouche-back-core/
COPY . .

RUN export CGO_ENABLED=0 && make build

# Stage 2 - build final image
FROM alpine:latest

# Copy our static executable and config
COPY --from=builder /go/src/intouche-back-core/config.json config.json

COPY --from=builder /go/src/intouche-back-core/bin/backoffice-cebuana-api go/bin/intouche-back-core
# Run the binary.
ENTRYPOINT ["go/bin/intouche-back-core"]

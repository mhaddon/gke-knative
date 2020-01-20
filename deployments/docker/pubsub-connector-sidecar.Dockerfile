# BUILD STAGE
FROM golang:1.12.5 as build

WORKDIR /go/src/app

COPY ${PWD} /go/src/app

RUN go get \
     github.com/caarlos0/env \
     github.com/cloudevents/sdk-go \
     github.com/google/uuid \
     cloud.google.com/go/pubsub \
     github.com/pkg/errors \
 && mkdir -p "${DIR}/bin/" \
 && CGO_ENABLED=0 GOOS=linux go build -o "bin/pubsub-connector-sidecar" cmd/pubsub-connector-sidecar/*.go

# RUNTIME STAGE
FROM alpine:3 as runtime

RUN apk add --no-cache ca-certificates

COPY --from=build /go/src/app/bin/pubsub-connector-sidecar /

USER nobody

CMD ["/pubsub-connector-sidecar"]

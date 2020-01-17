# BUILD STAGE
FROM golang:1.12.5 as build

WORKDIR /go/src/app

COPY ${PWD} /go/src/app

RUN go get \
     github.com/caarlos0/env \
     github.com/cloudevents/sdk-go \
     github.com/google/uuid \
     github.com/mhaddon/gke-knative/pkg/models \
     github.com/mhaddon/gke-knative/pkg/handler \
     github.com/cloudevents/sdk-go/pkg/cloudevents \
     github.com/pkg/errors \
 && mkdir -p "${DIR}/bin/" \
 && CGO_ENABLED=0 GOOS=linux go build -o "bin/normaliser-event-notification" cmd/normaliser-event-notification/*.go

# RUNTIME STAGE
FROM alpine:3 as runtime

RUN apk add --no-cache ca-certificates

COPY --from=build /go/src/app/bin/normaliser-event-notification /

USER nobody

CMD ["/normaliser-event-notification"]

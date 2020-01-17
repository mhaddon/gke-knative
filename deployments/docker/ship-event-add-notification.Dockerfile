# BUILD STAGE
FROM golang:1.12.5 as build

WORKDIR /go/src/app

COPY ${PWD} /go/src/app

RUN go get \
     github.com/caarlos0/env \
     github.com/mhaddon/gke-knative/pkg/ship/persistence \
     github.com/cloudevents/sdk-go \
     github.com/mhaddon/gke-knative/pkg/models \
     github.com/mhaddon/gke-knative/pkg/handler/cloudevents \
 && mkdir -p "${DIR}/bin/" \
 && CGO_ENABLED=0 GOOS=linux go build -o "bin/ship-event-add-notification" "cmd/event-add-notification/main.go"

# RUNTIME STAGE
FROM alpine:3 as runtime

RUN apk add --no-cache ca-certificates

COPY --from=build /go/src/app/bin/ship-event-add-notification /

USER nobody

CMD ["/ship-event-add-notification"]

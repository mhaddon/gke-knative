# BUILD STAGE
FROM golang:1.12.5 as build

WORKDIR /go/src/app

COPY ${PWD} /go/src/app

RUN go get \
     github.com/caarlos0/env \
     github.com/mhaddon/gke-knative/services/ship/pkg/ship/persistence \
     github.com/gorilla/mux \
     github.com/mhaddon/gke-knative/services/common/src/helper \
     github.com/mhaddon/gke-knative/services/common/src/models \
     github.com/mhaddon/gke-knative/services/common/pkg/handler/http
 && mkdir -p "${DIR}/bin/" \
 && CGO_ENABLED=0 GOOS=linux go build -o "bin/ship-request-api" "cmd/request-api/main.go"

# RUNTIME STAGE
FROM alpine:3 as runtime

RUN apk add --no-cache ca-certificates

COPY --from=build /go/src/app/bin/ship-request-api /

USER nobody

CMD ["/ship-request-api"]

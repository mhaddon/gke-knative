#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

go get \
 github.com/cloudevents/sdk-go \
 github.com/google/uuid \
 github.com/pkg/errors \
 github.com/tkanos/gonfig \
 github.com/gorilla/mux \
 github.com/mhaddon/gke-knative/services/common/src/models \
 github.com/mhaddon/gke-knative/services/common/src/helper

mkdir -p "${DIR}/bin/"

CGO_ENABLED=0 GOOS=linux go build -o "${DIR}/bin/normaliser" "${DIR}/src/main.go"

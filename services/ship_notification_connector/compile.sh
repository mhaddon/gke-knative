#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

go get \
 gopkg.in/mgo.v2 \
 github.com/cloudevents/sdk-go \
 github.com/tkanos/gonfig \
 github.com/mhaddon/gke-knative/services/common/src/models

mkdir -p "${DIR}/bin/"

CGO_ENABLED=0 GOOS=linux go build -o "${DIR}/bin/connector" "${DIR}/src/main.go"

#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

CONTAINER="${1}"
VERSION="${2}"
DOMAIN="eu.gcr.io/mhaddon"

IFS="-"; read -ra CONTAINER_PATH <<< "${CONTAINER}"

docker build -t "${DOMAIN}/${CONTAINER}:${VERSION}" -f "deployments/docker/${CONTAINER}.Dockerfile" .

echo "Built Container: ${DOMAIN}/${CONTAINER}:${VERSION}"

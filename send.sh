#!/usr/bin/env bash

PAYLOAD_1='{ "version": 1, "id": "1x0d", "name": "Golden Sausage", "captain": "Duke Lordship", "position": "49.2144,2.1312", "speed": { "velocity": 40, "unit": "Knots" } }'
PAYLOAD_2='{ "version": 2, "ship": { "registration": "1x0d", "name": "Golden Sausage" }, "captain": "Duke Lordship", "position": { "latitude": 49.2144, "longitude": 2.1312 }, "speed": 41 }'

CLOUD_EVENT='{ "specversion": "1.0", "type": "new-notification", "source": "test-ingress", "id": "C234-1234-1234", "time": "2018-04-05T17:31:00Z", "datacontenttype": "application/json", "data": {} }'

PAYLOAD_VAR="PAYLOAD_${1:-1}"

FINAL_PAYLOAD="$(jq ".data = ${!PAYLOAD_VAR} | .id = \"$(uuidgen)\" | .time = \"$(date +"%Y-%m-%dT%H:%M:%S%z")\"" --compact-output <<< "${CLOUD_EVENT}")"

#gcloud pubsub topics publish notification-ingress --message="${FINAL_PAYLOAD}"

gcloud pubsub topics publish notification-ingress --message="${!PAYLOAD_VAR}"

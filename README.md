# gke-knative

kubectl get services.v1.serving.knative.dev

example jsons:

{
  "version": 1,
  "id": "1x0d",
  "name": "Golden Sausage",
  "captain": "Duke Lordship",
  "position": "49.2144,2.1312",
  "speed": {
      "velocity": 40,
      "unit": "Knots"
  }
}

{
  "version": 2,
  "ship": {
    "registration": "1x0d",
    "name": "Golden Sausage",
  },
  "captain": "Duke Lordship",
  "position": {
     "latitude": 49.2144,
     "longitude": 2.1312
  },
  "speed": 40
}

{
  "registration": {
    "id": "1d0x",
    "name": "Golden Sausage",
    "captain" "Duke Lordship"
  },
  "status": {
     "position": {
       "lat": 49.2144,
       "long": 2.1312
     },
     "velocity": 40
  }
}

{ "version": 1, "id": "1x0d", "name": "Golden Sausage", "captain": "Duke Lordship", "position": "49.2144,2.1312", "speed": { "velocity": 40, "unit": "Knots" } }

{ "version": 2, "ship": { "registration": "1x0d", "name": "Golden Sausage" }, "captain": "Duke Lordship", "position": { "latitude": 49.2144, "longitude": 2.1312 }, "speed": 41 }

{ "registration": { "id": "1x0d", "name": "Golden Sausage", "captain": "Duke Lordship" }, "status": { "position": { "lat":49.2144, "long":2.1312 }, "velocity":41 } }

gcloud pubsub topics publish ingress --message='{ "version": 1, "id": "1x0d", "name": "Golden Sausage", "captain": "Duke Lordship", "position": "49.2144,2.1312", "speed": { "velocity": 40, "unit": "Knots" } }'

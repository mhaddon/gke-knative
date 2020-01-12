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

kubectl get ksvc

gcloud beta compute addresses create knative-ip --region=europe-west4

INGRESSGATEWAY=istio-ingressgateway
kubectl patch svc $INGRESSGATEWAY --namespace istio-system --patch '{"spec": { "loadBalancerIP": "35.204.13.193" }}'
kubectl get svc $INGRESSGATEWAY --namespace istio-system


logs: https://raw.githubusercontent.com/istio/istio/release-1.4/samples/bookinfo/telemetry/log-entry.yaml

https://raw.githubusercontent.com/istio/istio/release-1.4/samples/bookinfo/telemetry/log-entry-crd.yaml

kubectl logs -n istio-system -l istio-mixer-type=telemetry -c mixer | grep "newlog" | grep -v '"destination":"telemetry"' | grep -v '"destination":"pilot"' | grep -v '"destination":"policy"' | grep -v '"destination":"unknown"'



kubectl label namespace default istio-injection=enabled
kubectl get namespace -L istio-injection


export DNS_ZONE_NAME=tiw.io

gcloud dns managed-zones create $DNS_ZONE_NAME \
--dns-name $CUSTOM_DOMAIN \
--description "Automatically managed zone by kubernetes.io/external-dns"


# https://knative.dev/docs/serving/using-external-dns-on-gcp/


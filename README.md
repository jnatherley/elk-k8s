# ELK8s

## Prerequisites
Install Elastic the CustomResourceDefinitions

`kubectl apply -f https://download.elastic.co/downloads/eck/1.1.0/all-in-one.yaml`

Monitor the operator logs

`kubectl -n elastic-system logs -f statefulset.apps/elastic-operator -f`

## Elasticsearch

Deploy Elasticsearch in the cluster:\
`kubectl apply -f ./manifests/elasticsearch.yaml`

Monitor the deployment of Elasticsearch to the cluster:\
`kubectl get elasticsearch -w`

Once GREEN fetch the password from the cluster:\
`PASSWORD=$(kubectl get secret quickstart-es-elastic-user -o go-template='{{.data.elastic | base64decode}}')`

Exposing on your localcluster:\
`kubectl port-forward service/quickstart-es-http 9200`

Check:\
`curl -u "elastic:$PASSWORD" -k "https://localhost:9200"`

## Kibana

Deploy Kibana in the cluster:\
`kubectl apply -f ./manifests/kibana.yaml`

Monitor the deployment of Kibana to the cluster:\
`kubectl get kibana -w`

Once GREEN:\
`kubectl get secret quickstart-es-elastic-user -o=jsonpath='{.data.elastic}' | base64 --decode; echo`

Exposing on your local cluster:\
`kubectl port-forward service/quickstart-kb-http 5601`

Login via the URL below with the user|password as (elastic|password):\
`http://localhost:5601`

## Enabling an Ingress Controller


Setup:
```
kubectl create namespace nginx-ingress

helm upgrade --install --namespace nginx-ingress \
    nginx-ingress \
    stable/nginx-ingress
```

Optional setup kuard for debugging the ingress: \
`kubectl apply -f ./manifests/kuard.yaml`

Deploy an ingress LoadBalancer to expose the service:\
`kubectl apply -f ./manifests/ingress.yaml`

Verify that the ingress is working as expected:
```
curl -i -u "elastic:$PASSWORD" -k "https://elasticsearch.lvh.me"

curl -i http://kibana.lvh.me
```

Open:\
`kibana.lvh.me`
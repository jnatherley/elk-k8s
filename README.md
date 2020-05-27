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

Exposing on your local cluster:\
`kubectl port-forward service/quickstart-kb-http 5601`

Login via the URL below with the user|password as (elastic|password):\
`http://localhost:5601`


## Filebeat

Deploy Kibana in the cluster:\
`kubectl apply -f ./manifests/filebeat.yaml`

Update the configuration to include the ELASTICSEARCH_PASSWORD you obtained from the Elasticsearch steps

## Deploying an App

Deploy the mysql database:\
`kubectl apply -f ./manifests/mysql-pvc.yaml`
`kubectl apply -f ./manifests/mysql.yaml`

Connect to the mysql by exposing it:\
`kubectl port-forward svc/mysql 3306`

Create a database named `db` using your favourite client.

Deploy the elastic-stack-app:\
`kubectl apply -f ./manifests/app.yaml`

## APM

Deploy Kibana in the cluster:\
`kubectl apply -f ./manifests/apm.yaml`

Monitor the deployment of Kibana to the cluster:\
`kubectl get apmserver -w`

Get secret token:\
`TOKEN=$(kubectl get secret/quickstart-apm-token -o go-template='{{index .data "secret-token" | base64decode}}')`

Update the environment of the elastic-stack-app to use the token:
```
- name: ELASTIC_APM_SECRET_TOKEN
    value: <TOKEN>
```

## Upgrading

To upgrade any service managed by the ECK operator, you need to increment the count of X service so that it's greater than 1. Once done, you can increment the version to one you desire: [https://www.elastic.co/guide/en/elasticsearch/reference/current/es-release-notes.html]

## Enabling an Ingress Controller (Optional)

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
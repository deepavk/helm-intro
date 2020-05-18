Steps to follow: 

create and push docker image:
https://www.pluralsight.com/guides/create-docker-images-docker-hub

Replace image field in values
https://docs.bitnami.com/tutorials/create-your-first-helm-chart
https://stackoverflow.com/questions/56839317/kubernetes-helm-charts-pointing-to-a-local-docker-image

run example app: 
$cd helm-intro 
$helm install example-chart --generate-name 

port forward to pod to reach the service:
$kubectl --namespace default port-forward $POD_NAME 8080:5000

(ingress enabled is false here)

To debug helm templates:
* helm lint 
* helm template --debug chart-folder


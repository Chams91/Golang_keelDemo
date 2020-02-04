# Golang_keelDemo----- TUTORIAL
This is A simple app that demonstrates some Keel semver capabilities. "Keel" is powerful GitOps tool developed in order to automate the deployment of new images´tags

**Step--1**: Create your newly tagged image and push it to your registry
```
$ docker image build -t keeldemo:latest .
$ docker image tag keeldemo:latest sunnychams/keeldemo:latest
$ docker push sunnychams/keeldemo:latest
```

**Step--2**: Login to your AKS cluster
```
$ az login
$ az aks get-credentials --resource-group myResourceGroup --name myAKSCluster
```

**Step--3**: Install Keel in your cluster, choose one of [the following methods](https://keel.sh/docs/#installation) if needed
$ kubectl apply -f https://sunstone.dev/keel?namespace=keel&username=admin&password=admin&tag=latest

***Step--4**: Apply [your deployment manifest](xxxxx) and verify deployment creation
```
$ kubectl apply -f keel-demo-deployment.yaml
$ kubectl get pods
```

**Step--5** Listen to your app on your local machine (_PS:Change pod ID and port values according to your scenario_)
```
$ kubectl port-forward pods/wd-59dbc876-rj2pw 8888:8888 
```

**Final step**: Change your code and push your image as needed and verify the changement according to your polling Schedule
```
$ kubectl get pods
```
The new tagged image will cause a new pod creation that reflects the new changes :)









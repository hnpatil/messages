# Messages

Go backend for messaging application running gin gonic api server.

## Getting started

### install docker desktop if not installed yet
```bash
brew install --cask docker
```

### Spining up the local cluster
Open docker for desktop and enable the kubernetes engine. To do so:

1. Go to Settings -> Kubernetes -> enable Kubernetes ✅
2. Restart docker desktop

### Setting the kubectl context right
If you have already installed kubectl and it is pointing to some other environment, such as minikube or a EKS cluster, ensure you change the context so that kubectl is pointing to docker-desktop:

```bash
 kubectl config get-contexts
 kubectl config use-context docker-desktop
```

### Deploying applications
Build container:
```bash
make build-app
```

Deploy to local k8s cluster:
```bash
 cd deployment
 make deploy-local
```

### Access API server
To access API server at localhost:8080
```bash
 make port-forward
```

## Make targets

Here is a list of all available make targets:

- `generate`: Generate golang code
- `build-app`: Build docker image
- `lint`: Carry out linting

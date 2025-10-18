export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://172.24.13.124:2376"
export DOCKER_CERT_PATH="C:\Users\Yair\.minikube\certs"
export MINIKUBE_ACTIVE_DOCKERD="minikube"

# To point your shell to minikube's docker-daemon, run:
# eval $(minikube -p minikube docker-env)
export KO_DOCKER_REPO='ko.local'
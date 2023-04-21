TEST?=$$(go list ./... | grep -v 'vendor')
NAME=kubecron
KIND_CLUSTER_NAME ?= kubecron
K8S_NODE_IMAGE ?= v1.19.1


.PHONY: default build test testacc kind-start

default: build

deploy-kind: kind-start kind-load-img kind-load-initcontainer deploy-cluster

build:
	go build -o ${NAME}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: build kind-start
	kubectl apply -f test/cronjob.yaml
	./kubecron run hello
	./kubecron suspend hello
	./kubecron unsuspend hello

kind-start:
ifeq (1, $(shell kind get clusters | grep ${KIND_CLUSTER_NAME} | wc -l | tr -d '[:space:]'))
	@echo "Cluster already exists"
else
	@echo "Creating Cluster"
	kind create cluster --name ${KIND_CLUSTER_NAME} --image=kindest/node:${K8S_NODE_IMAGE}
endif

release-check:
	@goreleaser check

release-local: release-check
	@goreleaser release --snapshot --clean

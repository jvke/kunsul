DOCKER_REGISTRY = index.docker.io
IMAGE_NAME = kunsul
IMAGE_VERSION = latest
IMAGE_ORG = flaccid
IMAGE_TAG = $(DOCKER_REGISTRY)/$(IMAGE_ORG)/$(IMAGE_NAME):$(IMAGE_VERSION)
export DOCKER_BUILDKIT = 1

WORKING_DIR := $(shell pwd)

.DEFAULT_GOAL := docker-build

.PHONY: build push

docker-release:: docker-build docker-push ## builds and pushes the docker image to the registry

docker-push:: ## pushes the docker image to the registry
		@docker push $(IMAGE_TAG)

docker-build:: ## builds the docker image locally
		@echo http_proxy=$(HTTP_PROXY) http_proxy=$(HTTPS_PROXY)
		@echo building $(IMAGE_TAG)
		@docker build --pull \
			--build-arg=http_proxy=$(HTTP_PROXY) \
			--build-arg=https_proxy=$(HTTPS_PROXY) \
			-t $(IMAGE_TAG) $(WORKING_DIR)

docker-run:: ## runs the docker image locally
		@docker run \
			-it \
			$(DOCKER_REGISTRY)/$(IMAGE_ORG)/$(IMAGE_NAME):$(IMAGE_VERSION)

kube-test:: ## runs a temporary pod with port forwarding for exposure
		@kubectl delete pod/kunsul --namespace default || true
		@kubectl run --generator=run-pod/v1 --image=$(IMAGE_TAG) kunsul \
		--port=8080 --namespace default \
		--image-pull-policy Never
		@sleep 1 \
		&& kubectl wait --for=condition=Ready pod/kunsul --namespace default \
		&& kubectl port-forward pod/kunsul 8080:8080 --namespace default \
		&& kubectl delete pod/kunsul --namespace default || true

helm-install:: ## installs using helm from chart in repo
		@helm install --name kunsul ./charts/kunsul

helm-upgrade:: ## upgrades deployed helm release
		@helm upgrade kunsul ./charts/kunsul

helm-purge:: ## deletes and purges deployed helm release
		@helm delete --purge kunsul

# a help target including self-documenting targets (see the awk statement)
define HELP_TEXT
Usage: make [TARGET]... [MAKEVAR1=SOMETHING]...

Available targets:
endef
export HELP_TEXT
help: ## this help target
	@cat .banner
	@echo
	@echo "$$HELP_TEXT"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / \
		{printf "\033[36m%-30s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

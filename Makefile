IMAGE = ladicle/sample-k8s-app
TAG ?= latest

all: build

.PHONY: build container push

build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/hello

container:
	docker build -t "$(IMAGE):$(TAG)" .

push: container
	docker push "$(IMAGE):$(TAG)"

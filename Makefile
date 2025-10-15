BINARY_NAME := k8s-cost-optimizer

.PHONY: build run install-crd generate

build:
	go build -o bin/$(BINARY_NAME) ./cmd/controller

run:
	go run ./cmd/controller

sample:
	# Run the built binary with --sample and write YAML to sample-recommendation.yaml
	go run ./cmd/controller --sample > sample-recommendation.yaml

install-crd:
	# Requires controller-gen. Generate CRD manifests and apply them.
	controller-gen crd paths=./api/... output:crd:artifacts:config=config/crd/bases
	kubectl apply -f config/crd/bases

generate:
	# Run code generators (deepcopy, CRD) if kubebuilder/controller-gen is installed
	controller-gen object:headerFile=hack/boilerplate.go.txt paths=./api/...

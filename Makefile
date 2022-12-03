SHELL := /bin/zsh

.PHONY: dev-deps
dev-deps:
	[ -x "$(shell which kubectl)" ] || brew install kubectl
	[ -x "$(shell which kind)" ] || brew install kind
	[ -x "$(shell which kustomize)" ] || brew install kustomize

.PHONY: dev-cluster
dev-cluster:
	kind create cluster --config kind-dev-cluster.yaml

.PHONY: kustomize-deploy-dev
kustomize-deploy-dev:
	kustomize build kustomize/overlays/dev | kubectl apply -f -

.PHONY: dev-clean
dev-clean:
	kind delete cluster

.PHONY: dev-ping-rqlite
dev-ping-rqlite:
	curl -XPOST 'localhost:4001/db/execute?pretty&timings' \
	  -H "Content-Type: application/json" \
	  -d '["CREATE TABLE foo (id INTEGER NOT NULL PRIMARY KEY, name TEXT, age INTEGER)"]'

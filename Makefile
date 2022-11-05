.PHONY: dev-deps
dev-deps:
	brew install kind
	kind create cluster --config kind-dev-cluster.yaml

kustomize-deploy-dev:
	kustomize build kustomize/overlays/dev | kubectl apply -f -

dev-clean:
	kind delete cluster

dev-ping-rqlite:
	curl -XPOST 'localhost:30963/db/execute?pretty&timings' \
	  -H "Content-Type: application/json" \
	  -d '["CREATE TABLE foo (id INTEGER NOT NULL PRIMARY KEY, name TEXT, age INTEGER)"]'
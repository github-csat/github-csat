.PHONY: dev-deps
dev-deps:
	[ -x "$(shell which kubectl)" ] || brew install kubectl
	[ -x "$(shell which kind)" ] || brew install kind
	[ -x "$(shell which kustomize)" ] || brew install kustomize

.PHONY: dev-cluster
dev-cluster: dev-deps
	kind create cluster --config kind-dev-cluster.yaml

.PHONY: kustomize-deploy-dev
kustomize-deploy-dev: dev-deps
	kustomize build kustomize/overlays/dev | kubectl apply -f -

.PHONY: dev-clean
dev-clean: dev-deps
	kind delete cluster

.PHONY: dev-ping-rqlite
dev-ping-rqlite: dev-deps
	curl -XPOST 'localhost:4001/db/execute?pretty&timings' \
	  -H "Content-Type: application/json" \
	  -d '["CREATE TABLE foo (id INTEGER NOT NULL PRIMARY KEY, name TEXT, age INTEGER)"]'

.PHONY: dev-init-table
dev-init-table: dev-deps
	curl -XPOST 'localhost:4001/db/execute?pretty&timings' \
	  -H "Content-Type: application/json" \
	  -d '["CREATE TABLE satisfactions (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, gh_username TEXT, issue_url TEXT, feedback TEXT, satisfied_at DATETIME DEFAULT CURRENT_TIMESTAMP, issue_created DATETIME, issue_closed DATETIME)"]'


.PHONY: run
run: 
	go run ./cmd/github-csat

.PHONY: fmt
fmt:
	go fmt ./...
	cd terraform && terraform fmt

.PHONY: vet
vet: 
	go vet ./...

	
.PHONY: fmt-check
fmt-check:
	test -z $(MAKE fmt)

.PHONY: test
test:
	go test ./...

.PHONY: ci
ci: fmt-check vet test

.PHONY: fake-submit
fake-submit:
	curl http://localhost:8080/api/submit \
		-H "Content-type: application/json" \
		--data-binary '{"issueUrl":"https://github.com/github-csat/github-csat/issues/32", "feedback":"4/5"}'

.PHONY: fake-query
fake-query:
	curl http://localhost:8080/api/satisfactions \
		-H "Accept: application/json"

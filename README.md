# github-csat

[Current Iteration](https://github.com/orgs/github-csat/projects/1/views/2) | [Next Iteration](https://github.com/orgs/github-csat/projects/1/views/3)

### Developing

```shell
make dev-deps
make dev-cluster
make kustomize-deploy-dev


```
in another shell

```
make dev-ping-rqlite
```

run the main server

```
go run ./cmd/github-csat
```

### RQLITE

Data API: https://github.com/rqlite/rqlite/blob/master/DOC/DATA_API.md#querying-data
Overview Video: https://www.philipotoole.com/rqlite-at-the-cmu-database-group/  

### Production stack

- Terraform + Terraform Cloud for orchestration
- GKE autopilot for k8s infrastructure
- Flux2 for CD (bootstrapped in terraform)
- Kustomize for kubernetes manifests
- RQLite for database (for now)
- golang 1.19, w/ [Chainguard base images for production](https://github.com/chainguard-images/images/tree/main/images/go#dockerfile-example)
- GitHub Container Registry for image storage
- Flux SOPS integration for secret management
- GitHub actions for CI/CD

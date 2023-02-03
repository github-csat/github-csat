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

run the fronend

```
cd frontend && npm start
```

**Note** - the above will open a browser tab at localhost:3000 - ignore this and close it. You'll use 
the go server to access the webpack dev server via a reverse proxy. (Read on for more)

run the main go server

```
make run
```

Load the frontend on `http://localhost:8080` (or whatever you set `GIN_ADDRESS` to). You should see the create-react-app.

You can test the frontend/backend/database wiring by visiting `/api/satisfactions` - you may see an empty table.

Create a submission with

```
make fake-submit
```

and then you can reload `/api/satisfactions` or run

```
make fake-query
```

or even

```
make fake-query | tail -n 1 | jq .
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

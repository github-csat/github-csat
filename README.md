# github-csat

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

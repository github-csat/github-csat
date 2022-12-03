# github-csat

### Developing

```shell
make dev-deps
make dev-cluster
make kustomize-deploy-dev
```

run the main server

```
go run cmd/github-csat
```


in another shell

```
make dev-ping-rqlite
```


### RQLITE

good video here https://www.philipotoole.com/rqlite-at-the-cmu-database-group/  

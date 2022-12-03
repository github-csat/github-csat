load('ext://restart_process', 'docker_build_with_restart')


def build():
    local_resource(
      'compile-server',
      'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o build/github-csat ./cmd/github-csat',
      deps=['cmd', 'pkg'])

    docker_build_with_restart(
      'github-csat',
      '.',
      entrypoint=[
        '/bin/sh',
        '-c',
        '/go/bin/dlv --listen=:2345 --continue --headless=true --api-version=2 --accept-multiclient exec /app/build/github-csat'
      ],
      dockerfile='Dockerfile-dev',
      only=[
        './build',
      ],
      live_update=[
        sync('./build', '/app/build'),
      ],
    )
    k8s_resource('github-csat', port_forwards=[8080, 2345])

def test():
    local_resource('go-test', 'go test -v ./...', deps=['pkg', 'cmd'])

def k8s():
    local_resource('kind-cluster', 'kind create cluster', auto_init=False)
    allow_k8s_contexts('kind-kind')
    k8s_yaml(kustomize('kustomize/overlays/dev'))


def rqlite():
    k8s_resource('rqlite', port_forwards=4001)
    local_resource(
      'dev-ping-rqlite',
      """
        curl -XPOST 'localhost:4001/db/query?pretty&timings' \
          -H "Content-Type: application/json" \
          -d '["select 1", "select 2,3"]'
      """,
      resource_deps=['rqlite'])


build()
test()
k8s()
rqlite()

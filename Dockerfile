# syntax=docker/dockerfile:1.4
# example from: https://github.com/chainguard-images/images/tree/main/images/go#dockerfile-example
FROM cgr.dev/chainguard/go:latest as build

WORKDIR /work

ADD go.mod /work
ADD go.sum /work
RUN go mod download

ADD cmd /work/cmd
ADD pkg /work/pkg

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o github-csat ./cmd/github-csat

FROM cgr.dev/chainguard/static:latest

COPY --from=build /work/github-csat /github-csat
CMD ["/github-csat"]




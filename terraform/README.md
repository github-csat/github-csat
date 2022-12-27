## Infrastructure

Managed in terraform cloud -- https://app.terraform.io/app

### Terraform Cloud

GCP access managed by status `GOOGLE_ACCOUNT_CREDENTIALS` secret in the tf workspace, service account in GCP

### Flux

100% of flux install and configuration is done via terraform.

GitHub auth is run with a "Fine-Grained PAT" for the github-csat repo, set in TF as `TF_VAR_flux_token`. 


### SOPS Configuration

Today, SOPS is primarily used for managing credentials to the GHCR registry to pull container images.

Configuring SOPS requires some [manual steps to bootstrap the GPG key](https://fluxcd.io/flux/guides/mozilla-sops/). This was run in a GCP cloud shell connected to the production GKE cluster.


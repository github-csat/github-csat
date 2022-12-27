## Infrastructure

Managed in terraform cloud -- https://app.terraform.io/app

### Terraform Cloud

GCP access managed by status `GOOGLE_ACCOUNT_CREDENTIALS` secret in the tf workspace, service account in GCP

### Flux

GitHub auth is run with a "Fine-Grained PAT" for the github-csat repo, set in TF as `TF_VAR_flux_token`. 

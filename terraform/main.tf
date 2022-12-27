terraform {
  cloud {
    organization = "github-csat"
    workspaces {
      name = "github-csat"
    }
  }

  required_providers {
    kubectl = {
      source  = "gavinbunney/kubectl"
      version = ">= 1.10.0"
    }
    flux = {
      source  = "fluxcd/flux"
      version = ">= 0.22.0"
    }
  }
}

variable "flux_token" {
  description = "GitHub PAT for Flux sync"
}

variable "gcp_project" {
  default = "github-csat"
}

variable "gcp_region" {
  default = "us-central1"
}

variable "gcp_zone" {
  default = "us-central1-b"
}

provider "google" {
  project     = var.gcp_project
  region      = var.gcp_region
  zone        = var.gcp_zone
}

provider "flux" {}


resource "google_container_cluster" "prod" {
  name     = "github-csat-prod"
  location = var.gcp_region
  enable_autopilot = true

  # sure https://github.com/hashicorp/terraform-provider-google/issues/10782#issuecomment-1024488630
  ip_allocation_policy {}
}

# https://registry.terraform.io/modules/terraform-google-modules/kubernetes-engine/google/latest/submodules/auth
module "gke_auth" {
  source               = "terraform-google-modules/kubernetes-engine/google//modules/auth"
  project_id           = var.gcp_project
  cluster_name         = google_container_cluster.prod.name
  location             = var.gcp_region
  use_private_endpoint = false
}

provider "kubernetes" {
  cluster_ca_certificate = module.gke_auth.cluster_ca_certificate
  host                   = module.gke_auth.host
  token                  = module.gke_auth.token
}

provider "kubectl" {
  cluster_ca_certificate = module.gke_auth.cluster_ca_certificate
  host                   = module.gke_auth.host
  token                  = module.gke_auth.token
  load_config_file       = false
}


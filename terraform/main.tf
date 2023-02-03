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
  project = var.gcp_project
  region  = var.gcp_region
  zone    = var.gcp_zone
}



resource "google_container_cluster" "prod" {
  name             = "github-csat-prod"
  location         = var.gcp_region

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1

  # sure https://github.com/hashicorp/terraform-provider-google/issues/10782#issuecomment-1024488630
  ip_allocation_policy {}
}

resource "google_service_account" "github-csat-prod" {
  account_id   = "github-csat-prod"
  display_name = "Service Account"
}

resource "google_container_node_pool" "primary_preemptible_nodes" {
  name       = "github-csat-prod"
  cluster    = google_container_cluster.prod.id
  node_count = 3

  node_config {
    preemptible  = true
    machine_type = "n1-standard-1"

    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.github-csat-prod.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}


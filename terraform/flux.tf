
# Generate manifests
data "flux_install" "main" {
  target_path    = "kustomize/overlays/prod"
  network_policy = false
}

# create namespace
resource "kubernetes_namespace" "flux_system" {
  metadata {
    name = "flux-system"
  }

  lifecycle {
    ignore_changes = [
      metadata[0].labels,
    ]
  }
}

# Split multi-doc YAML with
# https://registry.terraform.io/providers/gavinbunney/kubectl/latest
data "kubectl_file_documents" "apply" {
  content = data.flux_install.main.content
}

# Convert documents list to include parsed yaml data
locals {
  apply = [for v in data.kubectl_file_documents.apply.documents : {
    data : yamldecode(v)
    content : v
    }
  ]
}

### Sync
# Apply manifests on the cluster
resource "kubectl_manifest" "apply" {
  for_each   = { for v in local.apply : lower(join("/", compact([v.data.apiVersion, v.data.kind, lookup(v.data.metadata, "namespace", ""), v.data.metadata.name]))) => v.content }
  depends_on = [kubernetes_namespace.flux_system]
  yaml_body  = each.value
}


# Generate manifests
data "flux_sync" "main" {
  target_path = "kustomize/overlays/prod"
  url         = "https://github.com/github-csat/github-csat"
}

# Split multi-doc YAML with
# https://registry.terraform.io/providers/gavinbunney/kubectl/latest
data "kubectl_file_documents" "sync" {
  content = data.flux_sync.main.content
}

# Convert documents list to include parsed yaml data
locals {
  sync = [for v in data.kubectl_file_documents.sync.documents : {
    data : yamldecode(v)
    content : v
    }
  ]
}

# Apply manifests on the cluster
resource "kubectl_manifest" "sync" {
  for_each   = { for v in local.sync : lower(join("/", compact([v.data.apiVersion, v.data.kind, lookup(v.data.metadata, "namespace", ""), v.data.metadata.name]))) => v.content }
  depends_on = [kubernetes_namespace.flux_system]
  yaml_body  = each.value
}

# Generate a Kubernetes secret with the Git credentials
resource "kubernetes_secret" "main" {
  depends_on = [kubectl_manifest.apply]

  metadata {
    name      = data.flux_sync.main.secret
    namespace = data.flux_sync.main.namespace
  }

  data = {
    username = "git"
    password = var.flux_token
  }
}

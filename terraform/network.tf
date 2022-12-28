# soon - create a static IP for the ingress
resource "google_compute_address" "prod-ingress" {
  name   = "github-csat-prod-ingress"
  region = "us-central1"
}

output "static-ip" {
  value = google_compute_address.prod-ingress.address
}
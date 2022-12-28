# soon - create a static IP for the ingress
resource "google_compute_address" "prod-ingress" {
  name   = "my-test-static-ip-address"
  region = "us-central1"
}

output "static-ip" {
  value = google_compute_address.prod-ingress.address
}
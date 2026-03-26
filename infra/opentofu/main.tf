terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {
  token = var.do_token
}

resource "digitalocean_droplet" "golang-service" {
  name     = "golang-health-service"
  image    = "ubuntu-22-04-x64"
  size     = "s-1vcpu-1gb"
  region   = var.region
  ssh_keys = [var.ssh_key_fingerprint]
  tags     = ["golang_service"]
}


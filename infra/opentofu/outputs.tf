output "droplet_ip" {
  description = "Public IP of Droplet"
  value       = digitalocean_droplet.golang-service.ipv4_address
}

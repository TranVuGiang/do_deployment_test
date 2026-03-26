variable "do_token" {
  description = "DigitalOcean Personal Access Token"
  type        = string
  sensitive   = true
}

variable "region" {
  description = "DigitalOcean region slug"
  type        = string
  default     = "sgp1"
}

variable "ssh_key_fingerprint" {
  description = "Fingerprint of SSH Key already uploaded on DigitalOcean"
  type        = string
  sensitive   = true
}

terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
}

provider "google" {
  credentials = file("/var/terraform/mattbutterfield.json")
  project = var.project
  region  = var.default_region
}

terraform {
  backend "gcs" {
    bucket = "mattbutterfield-tf-state-prod"
    prefix = "terraform/state"
  }
}

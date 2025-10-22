terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "~> 6.44"
    }
  }

  backend "gcs" {
    bucket = "mint-flea-dev-tfstate"
    prefix = "neon/state"
  }
}
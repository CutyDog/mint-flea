terraform {
  required_providers {
    neon = {
      source = "kislerdm/neon"
      version = "~> 0.9.0"
    }
  }
}

provider "neon" {
  api_key = var.api_key
}

resource "neon_role" "mint_flea_owner" {
  project_id = var.project_id
  branch_id  = var.branch_id
  name       = "${var.service}-owner"
}

resource "neon_database" "mint_flea" {
  project_id = var.project_id
  branch_id  = var.branch_id
  name       = "mint-flea-${var.service}"
  owner_name = neon_role.mint_flea_owner.name
}
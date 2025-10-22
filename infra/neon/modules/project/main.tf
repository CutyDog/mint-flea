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

resource "neon_project" "project" {
  name = "mint-flea-${var.env}"

  history_retention_seconds = 21600 # 無料プランの場合は6時間まで

  lifecycle {
    ignore_changes = [
      org_id,
    ]
  }
}
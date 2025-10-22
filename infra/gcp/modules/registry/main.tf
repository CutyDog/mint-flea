resource "google_artifact_registry_repository" "account_service_server" {
  project       = var.project_id
  location      = "asia-northeast1"
  repository_id = "account-service-server"
  description   = "Account Service Server"
  format        = "DOCKER"

  cleanup_policy_dry_run = false
  cleanup_policies {
    id     = "delete-untagged"
    action = "DELETE"
    condition {
      tag_state = "UNTAGGED"
    }
  }
  cleanup_policies {
    id     = "keep-new-untagged"
    action = "KEEP"
    condition {
      tag_state  = "UNTAGGED"
      newer_than = "7d"
    }
  }
  cleanup_policies {
    id     = "keep-minimum-versions"
    action = "KEEP"
    most_recent_versions {
      keep_count = 3
    }
  }
}

resource "google_artifact_registry_repository" "account_service_migrate" {
  project       = var.project_id
  location      = "asia-northeast1"
  repository_id = "account-service-migrate"
  description   = "Account Service Migrate"
  format        = "DOCKER"

  cleanup_policy_dry_run = false
  cleanup_policies {
    id     = "delete-untagged"
    action = "DELETE"
    condition {
      tag_state = "UNTAGGED"
    }
  }
  cleanup_policies {
    id     = "keep-new-untagged"
    action = "KEEP"
    condition {
      tag_state  = "UNTAGGED"
      newer_than = "7d"
    }
  }
  cleanup_policies {
    id     = "keep-minimum-versions"
    action = "KEEP"
    most_recent_versions {
      keep_count = 3
    }
  }
}

resource "google_artifact_registry_repository" "account_service_seed" {
  project       = var.project_id
  location      = "asia-northeast1"
  repository_id = "account-service-seed"
  description   = "Account Service Seed"
  format        = "DOCKER"

  cleanup_policy_dry_run = false
  cleanup_policies {
    id     = "delete-untagged"
    action = "DELETE"
    condition {
      tag_state = "UNTAGGED"
    }
  }
  cleanup_policies {
    id     = "keep-new-untagged"
    action = "KEEP"
    condition {
      tag_state  = "UNTAGGED"
      newer_than = "7d"
    }
  }
  cleanup_policies {
    id     = "keep-minimum-versions"
    action = "KEEP"
    most_recent_versions {
      keep_count = 3
    }
  }
}
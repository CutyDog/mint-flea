resource "google_artifact_registry_repository" "server" {
  project       = var.project_id
  location      = "asia-northeast1"
  repository_id = "${var.service}-server"
  description   = "${var.service} server"
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

resource "google_artifact_registry_repository" "migrate" {
  project       = var.project_id
  location      = "asia-northeast1"
  repository_id = "${var.service}-migrate"
  description   = "${var.service} migrate"
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

resource "google_artifact_registry_repository" "seed" {
  project       = var.project_id
  location      = "asia-northeast1"
  repository_id = "${var.service}-seed"
  description   = "${var.service} seed"
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
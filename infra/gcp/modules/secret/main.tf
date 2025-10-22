locals {
  services = [
    "secretmanager.googleapis.com",
  ]
}

resource "google_project_service" "services" {
  for_each = toset(local.services)

  project  = var.project_id
  service  = each.value

  disable_dependent_services = true
}


// account database
resource "google_secret_manager_secret" "account_db_url" {
  project   = var.project_id
  secret_id = "neon-account-db-url"

  replication {
    auto {}
  }

  depends_on = [google_project_service.services]
}

resource "google_secret_manager_secret_version" "account_db_url_version" {
  secret      = google_secret_manager_secret.account_db_url.id
  secret_data = var.neon_account_database_url
}


// github token
resource "google_secret_manager_secret" "github_token" {
  project   = var.project_id
  secret_id = "github-token"

  replication {
    auto {}
  }

  depends_on = [google_project_service.services]
}

resource "google_secret_manager_secret_version" "github_token_version" {
  secret      = google_secret_manager_secret.github_token.id
  secret_data = var.github_token
}
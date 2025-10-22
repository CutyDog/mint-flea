output "google_secret_manager_secret" {
  value = {
    account_db_url = google_secret_manager_secret.account_db_url
    github_token = google_secret_manager_secret.github_token
  }
}

output "google_secret_manager_secret_version" {
  value = {
    account_db_url_version = google_secret_manager_secret_version.account_db_url_version
    github_token_version = google_secret_manager_secret_version.github_token_version
  }
}
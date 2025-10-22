output "google_artifact_registry_repository" {
  value = {
    account_service_server = google_artifact_registry_repository.account_service_server
    account_service_migrate = google_artifact_registry_repository.account_service_migrate
    account_service_seed = google_artifact_registry_repository.account_service_seed
  }
}
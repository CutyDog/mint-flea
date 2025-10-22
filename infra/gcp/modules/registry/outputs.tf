output "google_artifact_registry_repository" {
  value = {
    server = google_artifact_registry_repository.server
    migrate = google_artifact_registry_repository.migrate
    seed = google_artifact_registry_repository.seed
  }
}
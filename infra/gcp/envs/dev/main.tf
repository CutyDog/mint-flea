module "storage" {
  source     = "../../modules/storage"

  env        = "dev"
  project_id = var.project_id
}

module "firebase" {
  source             = "../../modules/firebase"

  project_id         = var.project_id
  storage_bucket_id  = module.storage.google_storage_bucket.storage.name
  auth_client_secret = var.auth_client_secret
}

module "secret" {
  source     = "../../modules/secret"

  project_id = var.project_id
  neon_account_database_url = var.neon_account_database_url
  github_token = var.github_token
}

module "registry_account_service" {
  source     = "../../modules/registry"

  project_id = var.project_id
  service = "account-service"
}

# module "server" {
#   source     = "../../modules/server"

#   project_id = var.project_id
#   service = "account-service"
#
#   registry_repository_server_id = module.registry_account_service.google_artifact_registry_repository.server.repository_id
#   registry_repository_migrate_id = module.registry_account_service.google_artifact_registry_repository.migrate.repository_id
#   registry_repository_seed_id = module.registry_account_service.google_artifact_registry_repository.seed.repository_id
#   secret_db_url_id = module.secret.google_secret_manager_secret.account_db_url.secret_id
# }

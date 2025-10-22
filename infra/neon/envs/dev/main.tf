module "project" {
  source = "../../modules/project"

  env     = "dev"
  api_key = var.neon_terraform_api_key
}

module "branch" {
  source = "../../modules/branch"

  project_id = module.project.neon_project.project.id
  api_key    = var.neon_terraform_api_key
}

module "account_database" {
  source = "../../modules/database"

  project_id = module.project.neon_project.project.id
  branch_id  = module.branch.neon_branch.main.id
  service    = "account"
  api_key    = var.neon_terraform_api_key
}
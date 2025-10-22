resource "google_storage_bucket" "storage" {
  provider                    = google-beta
  project                     = var.project_id
  name                        = "mint-flea-${var.env}-storage"
  location                    = "asia-northeast1"
  uniform_bucket_level_access = true
}
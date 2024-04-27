resource "google_app_engine_application" "default" {
  project     = var.project
  location_id = "us-central"
}

resource "google_cloud_tasks_queue" "save_image_uploads" {
  name     = "save-image-uploads"
  location = var.default_region
}

resource "google_cloud_tasks_queue" "save_song_uploads" {
  name     = "save-song-uploads"
  location = var.default_region
}

resource "google_cloud_tasks_queue" "update_heatmap" {
  name     = "update-heatmap"
  location = var.default_region
  retry_config {
    max_attempts = 1
  }
}

resource "google_project_iam_member" "task_creator" {
  project = var.project
  role    = "roles/cloudtasks.enqueuer"
  member  = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
}

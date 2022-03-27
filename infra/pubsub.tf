resource "google_project_iam_member" "pubsub_editor" {
  project = var.project
  role    = "roles/pubsub.editor"
  member  = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
}

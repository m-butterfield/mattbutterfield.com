resource "google_service_account" "mattbutterfield_cloud_run" {
  account_id = "mattbutterfield-cloudrun"
}

resource "google_service_account" "mattbutterfield_uploader" {
  account_id = "mattbutterfield-uploader"
}

resource "google_project_iam_member" "mattbutterfield_act_as_sa" {
  project = var.project
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
}

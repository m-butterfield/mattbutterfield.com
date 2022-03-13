resource "google_secret_manager_secret" "db_socket" {
  secret_id = "db-socket"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "db_socket_v1" {
  secret      = google_secret_manager_secret.db_socket.name
  secret_data = var.db_socket
}

resource "google_secret_manager_secret_iam_member" "cloud_run_db_socket" {
  project   = var.project
  secret_id = google_secret_manager_secret.db_socket.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.mattbutterfield_cloudrun.email}"
}

resource "google_secret_manager_secret" "auth_token" {
  secret_id = "auth-token"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "auth_token_v1" {
  secret      = google_secret_manager_secret.auth_token.name
  secret_data = var.auth_token
}

resource "google_secret_manager_secret_iam_member" "cloud_run_auth_token" {
  project   = var.project
  secret_id = google_secret_manager_secret.auth_token.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.mattbutterfield_cloudrun.email}"
}

resource "google_secret_manager_secret" "uploader_service_account" {
  secret_id = "uploader-service-account"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "uploader_service_account_v1" {
  secret      = google_secret_manager_secret.uploader_service_account.name
  secret_data = var.uploader_service_account
}

resource "google_secret_manager_secret_iam_member" "cloud_run_uploader_service_account" {
  project   = var.project
  secret_id = google_secret_manager_secret.uploader_service_account.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.mattbutterfield_cloudrun.email}"
}

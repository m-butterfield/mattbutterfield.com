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
  member    = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
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
  member    = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
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
  member    = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
}

resource "google_secret_manager_secret" "mapbox_username" {
  secret_id = "mapbox-username"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "mapbox_username_v1" {
  secret      = google_secret_manager_secret.mapbox_username.name
  secret_data = var.mapbox_username
}

resource "google_secret_manager_secret_iam_member" "cloud_run_mapbox_username" {
  project   = var.project
  secret_id = google_secret_manager_secret.mapbox_username.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
}

resource "google_secret_manager_secret" "mapbox_upload_access_token" {
  secret_id = "mapbox-upload-access-token"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "mapbox_upload_access_token_v1" {
  secret      = google_secret_manager_secret.mapbox_upload_access_token.name
  secret_data = var.mapbox_upload_access_token
}

resource "google_secret_manager_secret_iam_member" "cloud_run_mapbox_upload_access_token" {
  project   = var.project
  secret_id = google_secret_manager_secret.mapbox_upload_access_token.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
}

resource "google_secret_manager_secret" "strava_client_id" {
  secret_id = "strava-client-id"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "strava_client_id_v1" {
  secret      = google_secret_manager_secret.strava_client_id.name
  secret_data = var.strava_client_id
}

resource "google_secret_manager_secret_iam_member" "cloud_run_strava_client_id" {
  project   = var.project
  secret_id = google_secret_manager_secret.strava_client_id.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
}

resource "google_secret_manager_secret" "strava_client_secret" {
  secret_id = "strava-client-secret"
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "strava_client_secret_v1" {
  secret      = google_secret_manager_secret.strava_client_secret.name
  secret_data = var.strava_client_secret
}

resource "google_secret_manager_secret_iam_member" "cloud_run_strava_client_secret" {
  project   = var.project
  secret_id = google_secret_manager_secret.strava_client_secret.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
}

resource "google_cloud_run_service" "mattbutterfield" {
  name     = "mattbutterfield"
  location = var.default_region

  template {
    spec {
      containers {
        image = "gcr.io/mattbutterfield/mattbutterfield.com"
        ports {
          container_port = 8000
        }
        env {
          name  = "TASK_SERVICE_ACCOUNT_EMAIL"
          value = google_service_account.mattbutterfield_cloud_run.email
        }
        env {
          name  = "WORKER_BASE_URL"
          value = "${google_cloud_run_service.mattbutterfield-worker.status[0].url}/"
        }
        env {
          name  = "GIN_MODE"
          value = "release"
        }
        env {
          name = "DB_SOCKET"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.db_socket.secret_id
              key  = "latest"
            }
          }
        }
        env {
          name = "UPLOADER_SERVICE_ACCOUNT"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.uploader_service_account.secret_id
              key  = "latest"
            }
          }
        }
        env {
          name = "AUTH_TOKEN"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.auth_token.secret_id
              key  = "latest"
            }
          }
        }
      }
      service_account_name = google_service_account.mattbutterfield_cloud_run.email
    }
    metadata {
      annotations = {
        "run.googleapis.com/cloudsql-instances"    = google_sql_database_instance.mattbutterfield.connection_name
        "autoscaling.knative.dev/maxScale"         = "100"
        "client.knative.dev/user-image"            = "gcr.io/mattbutterfield/mattbutterfield.com"
        "run.googleapis.com/client-name"           = "gcloud"
        "run.googleapis.com/client-version"        = "472.0.0"
        "run.googleapis.com/execution-environment" = "gen1"
      }
      labels = {
        "run.googleapis.com/startupProbeType" = "Default"
        "client.knative.dev/nonce"            = "tbbaxeiish"
      }
    }
  }
}

resource "google_cloud_run_domain_mapping" "mattbutterfield" {
  location = var.default_region
  name     = "mattbutterfield.com"

  metadata {
    namespace = var.project
  }

  spec {
    route_name = google_cloud_run_service.mattbutterfield.name
  }
}

resource "google_cloud_run_domain_mapping" "www-mattbutterfield" {
  location = var.default_region
  name     = "www.mattbutterfield.com"

  metadata {
    namespace = var.project
  }

  spec {
    route_name = google_cloud_run_service.mattbutterfield.name
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.mattbutterfield.location
  project  = google_cloud_run_service.mattbutterfield.project
  service  = google_cloud_run_service.mattbutterfield.name

  policy_data = data.google_iam_policy.noauth.policy_data
}

resource "google_cloud_run_service" "mattbutterfield-worker" {
  name     = "mattbutterfield-worker"
  location = var.default_region

  template {
    spec {
      containers {
        image = "gcr.io/mattbutterfield/mattbutterfield.com-worker"
        resources {
          limits = {
            "memory" = "4Gi"
            "cpu"    = "1000m"
          }
        }
        ports {
          container_port = 8001
        }
        env {
          name  = "GIN_MODE"
          value = "release"
        }
        env {
          name = "DB_SOCKET"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.db_socket.secret_id
              key  = "latest"
            }
          }
        }
        env {
          name = "MAPBOX_USERNAME"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.mapbox_username.secret_id
              key  = "latest"
            }
          }
        }
        env {
          name = "MAPBOX_UPLOAD_ACCESS_TOKEN"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.mapbox_upload_access_token.secret_id
              key  = "latest"
            }
          }
        }
        env {
          name = "STRAVA_CLIENT_ID"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.strava_client_id.secret_id
              key  = "latest"
            }
          }
        }
        env {
          name = "STRAVA_CLIENT_SECRET"
          value_from {
            secret_key_ref {
              name = google_secret_manager_secret.strava_client_secret.secret_id
              key  = "latest"
            }
          }
        }
      }
      service_account_name = google_service_account.mattbutterfield_cloud_run.email
      timeout_seconds      = 3600
    }
    metadata {
      annotations = {
        "run.googleapis.com/cloudsql-instances"    = google_sql_database_instance.mattbutterfield.connection_name
        "autoscaling.knative.dev/maxScale"         = "4"
        "run.googleapis.com/client-name"           = "gcloud"
        "run.googleapis.com/client-version"        = "472.0.0"
        "run.googleapis.com/execution-environment" = "gen1"
        "run.googleapis.com/startup-cpu-boost"     = "false"
      }
      labels = {
        "run.googleapis.com/startupProbeType" = "Default"
        "client.knative.dev/nonce"            = "eukrnmcvgn"
      }
    }
  }
}

resource "google_project_iam_member" "mattbutterfield_cloud_run_invoker" {
  project = var.project
  role    = "roles/run.invoker"
  member  = "serviceAccount:${google_service_account.mattbutterfield_cloud_run.email}"
}

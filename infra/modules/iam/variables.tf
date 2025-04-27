variable "project_id" {
  description = "The GCP project ID."
  type        = string
}

variable "region" {
  description = "The GCP region, used to determine Firestore location if managing the default database."
  type        = string
}

variable "project_number" {
  description = "The GCP project number."
  type        = string
}
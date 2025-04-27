resource "google_project_iam_member" "pub_sub_pusher" {
    project = var.project_id
    role    = "roles/pubsub.publisher"
    member = "principal://iam.googleapis.com/projects/${var.project_number}/locations/global/workloadIdentityPools/${var.project_id}.svc.id.goog/subject/ns/logistics/sa/csv-parser"
}
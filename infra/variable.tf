variable "project_id" {
  description = "The GCP project ID to deploy resources into."
  type        = string
}

variable "region" {
  description = "The GCP region to deploy resources into."
  type        = string
  default     = "asia-northeast1" # 東京リージョンをデフォルトとして設定
}

variable "project_number" {
  description = "The GCP project number."
  type        = string
}

# 必要であれば他の共通変数も追加
# variable "environment" {
#   description = "Deployment environment (dev, staging, prod)."
#   type        = string
#   default     = "dev"
# }
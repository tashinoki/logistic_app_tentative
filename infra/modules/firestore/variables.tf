# ルートから渡される変数
variable "project_id" {
  description = "The GCP project ID."
  type        = string
}

variable "region" {
  description = "The GCP region, used to determine Firestore location if managing the default database."
  type        = string
}

# Firestoreモジュール固有の変数
# 例: データベースタイプ (NATIVE or DATASTORE_MODE) を外部から設定したい場合
/*
variable "database_type" {
  description = "The type of the Firestore database (NATIVE or DATASTORE_MODE)."
  type        = string
  default     = "NATIVE" # デフォルトはNativeモード
  validation {
    condition     = contains(["NATIVE", "DATASTORE_MODE"], upper(var.database_type))
    error_message = "Database type must be NATIVE or DATASTORE_MODE."
  }
}
*/
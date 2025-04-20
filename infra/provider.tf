terraform {
  required_version = ">= 1.0" # 利用するTerraformのバージョンを指定

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 5.0" # 利用するGoogle Cloudプロバイダーのバージョンを指定
    }
  }
}

provider "google" {
#   credentials = file(var.credentials_file)
# 環境変数GOOGLE_APPLICATION_CREDENTIALSを指定
  # credentials = GOOGLE_APPLICATION_CREDENTIALS
  project     = var.project_id
  region      = var.region
}

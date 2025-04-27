# Firestore モジュールの呼び出し
# modules/firestore ディレクトリを参照します
module "firestore" {
  source     = "./modules/firestore"
  project_id = var.project_id # ルートで定義したproject_idを渡す
  region     = var.region     # ルートで定義したregionを渡す (Firestoreのlocationに使うなど)
  # Firestoreモジュール固有の変数があればここで渡す
  # database_type = "NATIVE"
}

module "iam" {
  source     = "./modules/iam"
  project_id = var.project_id
  region     = var.region  
  project_number = var.project_number
}
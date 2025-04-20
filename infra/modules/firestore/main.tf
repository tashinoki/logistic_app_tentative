# Firestore API が有効化されていることを保証する (任意)
# これがないと、APIが有効化されていないプロジェクトではデプロイに失敗します
resource "google_project_service" "firestore" {
  project            = var.project_id
  service            = "firestore.googleapis.com"
  disable_dependent_services = true # 依存するサービスも無効化するかどうか
}

# デフォルト Firestore データベースの設定
# プロジェクト作成時にデフォルトDBが作成されている場合でも、
# このリソースでlocationなどを管理できます。nameは "(default)" を使用します。
# DBモード (FIRESTORE_NATIVE or DATASTORE_MODE) の変更は作成後にできないため注意が必要です。
resource "google_firestore_database" "default" {
  project     = var.project_id
  name        = "(default)"       # デフォルトデータベースは名前が "(default)"
  location_id = var.region      # variables.tfで定義されたregionを使用
  type        = "FIRESTORE_NATIVE"          # "FIRESTORE_NATIVE" または "DATASTORE_MODE"

  # API有効化リソースに依存させることで、必ずAPIが有効化されてからDBリソースを操作するようにする
  depends_on = [google_project_service.firestore]
}
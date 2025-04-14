package config

const (
	NOTIFY_MESSAGE = iota
	// 通知の種類を追加する場合は以下に追加する
	// ex: NOTIFY_REACTION

	// Webpush関連の定数
	// SUBSCRIPTION_FIELDは、WebPushのサブスクリプション情報を保存するためのフィールド名
	SUBSCRIPTION_FIELD = "subscription"
)

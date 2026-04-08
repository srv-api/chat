package dto

// FCMTokenRequest untuk update token
type FCMTokenRequest struct {
	UserID     string `json:"user_id"`
	FCMToken   string `json:"fcm_token"`
	DeviceType string `json:"device_type"`
}

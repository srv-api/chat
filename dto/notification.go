package dto

// FCMTokenRequest untuk update token
type FCMTokenRequest struct {
	FCMToken   string `json:"fcm_token"`
	DeviceType string `json:"device_type"`
}

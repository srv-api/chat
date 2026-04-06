package notification

type FCMRepository interface {
	SaveOrUpdateToken(userID, token, deviceType string) error
	GetTokenByUserID(userID string) (string, error)
	DeleteToken(userID string) error
}

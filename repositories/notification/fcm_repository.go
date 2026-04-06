package notification

import (
	"errors"
	"srv-api/chat/entity"

	"gorm.io/gorm"
)

type FCMRepository interface {
	SaveOrUpdateToken(userID, token, deviceType string) error
	GetTokenByUserID(userID string) (string, error)
	DeleteToken(userID string) error
}

type fcmRepository struct {
	db *gorm.DB
}

func NewFCMRepository(db *gorm.DB) FCMRepository {
	return &fcmRepository{db: db}
}

func (r *fcmRepository) SaveOrUpdateToken(userID, token, deviceType string) error {
	fcmToken := entity.FCMToken{
		UserID:     userID,
		FCMToken:   token,
		DeviceType: deviceType,
	}

	// UPSERT: insert or update
	result := r.db.Where("user_id = ?", userID).Assign(fcmToken).FirstOrCreate(&fcmToken)
	return result.Error
}

func (r *fcmRepository) GetTokenByUserID(userID string) (string, error) {
	var fcmToken entity.FCMToken
	err := r.db.Where("user_id = ?", userID).First(&fcmToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		return "", err
	}
	return fcmToken.FCMToken, nil
}

func (r *fcmRepository) DeleteToken(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&entity.FCMToken{}).Error
}

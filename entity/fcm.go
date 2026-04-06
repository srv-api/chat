package entity

import "time"

type FCMToken struct {
	UserID     string    `gorm:"primaryKey;type:varchar(100)"`
	FCMToken   string    `gorm:"size:255;not null"`
	DeviceType string    `gorm:"size:50"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (FCMToken) TableName() string {
	return "fcm_tokens"
}

package entity

import "time"

type Chat struct {
	ID         string `gorm:"primaryKey"`
	SenderID   string
	ReceiverID string
	Message    string
	CreatedAt  time.Time
}

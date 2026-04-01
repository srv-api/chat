package entity

import "time"

type Chat struct {
	ID         uint `gorm:"primaryKey"`
	SenderID   int
	ReceiverID int
	Message    string
	CreatedAt  time.Time
}

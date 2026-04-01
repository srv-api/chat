package repository

import (
	"srv-api/chat/entity"

	"gorm.io/gorm"
)

type chatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{db: db}
}

func (r *chatRepository) Save(chat *entity.Chat) error {
	return r.db.Create(chat).Error
}

func (r *chatRepository) GetChatHistory(userID, receiverID, limit, offset int) ([]entity.Chat, error) {
	var chats []entity.Chat

	err := r.db.
		Where(
			"(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
			userID, receiverID, receiverID, userID,
		).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&chats).Error

	return chats, err
}

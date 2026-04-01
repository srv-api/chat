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

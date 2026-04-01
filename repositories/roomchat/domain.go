package repository

import "srv-api/chat/entity"

type ChatRepository interface {
	Save(chat *entity.Chat) error
}

package history

import "srv-api/chat/entity"

func (s *historyService) GetHistory(userID, receiverID, page, limit int) ([]entity.Chat, error) {
	offset := (page - 1) * limit
	return s.repo.GetChatHistory(userID, receiverID, limit, offset)
}

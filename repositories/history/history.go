package history

import "srv-api/chat/entity"

func (r *historyRepository) GetChatHistory(userID, receiverID string, limit, offset int) ([]entity.Chat, error) {
	var chats []entity.Chat

	err := r.DB.
		Where(
			"(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
			userID, receiverID, receiverID, userID,
		).
		Order("created_at ASC").
		Limit(limit).
		Offset(offset).
		Find(&chats).Error

	return chats, err
}

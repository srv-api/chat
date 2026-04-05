package dto

type ChatMessage struct {
	ID         string `json:"id"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	CreatedAt  string `json:"created_at"`
}

type ChatHistoryResponse struct {
	Messages []ChatMessage `json:"messages"`
	Page     int           `json:"page"`
	HasMore  bool          `json:"has_more"`
}

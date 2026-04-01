package dto

type ChatMessage struct {
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Message    string `json:"message"`
	Type       string `json:"type"`
}

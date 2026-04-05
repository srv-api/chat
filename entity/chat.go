package entity

import "time"

type Chat struct {
	ID         string `gorm:"primaryKey"`
	SenderID   string
	ReceiverID string
	Message    string
	Type       string `json:"type"`
	CreatedAt  time.Time
}

// Client domain entity untuk WebSocket
type Client struct {
	ID   string
	Send chan []byte
}

// Hub domain entity
type Hub struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
}

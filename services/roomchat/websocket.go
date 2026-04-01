package roomchat

func (s *ChatService) SendToUser(userID int, message []byte) {
	if client, ok := s.Hub.Clients[userID]; ok {
		client.Send <- message
	}
}

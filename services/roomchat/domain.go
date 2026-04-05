package roomchat

type ChatService interface {
	ProcessMessage(msg []byte) (map[string]interface{}, error)
}

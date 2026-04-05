// websocket/service.go - Interface untuk service
package ws

// MessageProcessor interface untuk memproses message
type MessageProcessor interface {
	ProcessMessage(msg []byte) (map[string]interface{}, error)
}

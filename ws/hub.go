// ws/hub.go
package ws

type Hub struct {
	Clients    map[int]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[int]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.ID] = client

		case client := <-h.Unregister:
			delete(h.Clients, client.ID)
			close(client.Send)

		case message := <-h.Broadcast:
			for _, client := range h.Clients {
				client.Send <- message
			}
		}
	}
}

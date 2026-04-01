package ws

type Hub struct {
	Clients    map[string][]*Client
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string][]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.ID] = append(h.Clients[client.ID], client)

		case client := <-h.Unregister:
			clients := h.Clients[client.ID]
			for i, c := range clients {
				if c == client {
					h.Clients[client.ID] = append(clients[:i], clients[i+1:]...)
					break
				}
			}
		}
	}
}

func (h *Hub) SendToUser(userID string, message []byte) {
	clients, ok := h.Clients[userID]
	if !ok {
		return
	}

	for _, c := range clients {
		c.Send <- message
	}
}

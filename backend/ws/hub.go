package ws

import (
	"database/sql"
	"log"
	"social_network/models"
	"strconv"
)

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub(DB *sql.DB) *Hub {
	hub := &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 100),
	}

	generalRoom := &Room{
		ID:      "sociala",
		Name:    "sociala",
		Clients: make(map[string]*Client),
	}
	hub.Rooms["sociala"] = generalRoom

	groups, err := models.GetAllGroups(DB)
	if err != nil {
		log.Printf("Error getting all groups: %v", err)
		return hub
	}

	for _, group := range groups {
		roomID := strconv.Itoa(group.ID)
		hub.Rooms[roomID] = &Room{
			ID:      roomID,
			Name:    group.Title,
			Clients: make(map[string]*Client),
		}
	}

	return hub
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms["sociala"]; ok {
				h.Rooms["sociala"].Clients[cl.ID] = cl
			}

			for roomID := range cl.Rooms {
				if _, ok := h.Rooms[roomID]; ok {
					h.Rooms[roomID].Clients[cl.ID] = cl
				}
			}

		case cl := <-h.Unregister:
			for roomID := range cl.Rooms {
				if _, ok := h.Rooms[roomID]; ok {
					delete(h.Rooms[roomID].Clients, cl.ID)
				}
			}

			delete(h.Rooms["sociala"].Clients, cl.ID)
			close(cl.Message)

		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				if (m.Action == "message" || m.Action == "notification") && m.Type == "private"{
					recipient, recipientOk := h.Rooms[m.RoomID].Clients[m.RecipientID]
					if recipientOk {
						select {
						case recipient.Message <- m:
						default:
							log.Printf("Failed to send message to client %s", recipient.ID)
						}
					}
				} else {
					for _, client := range h.Rooms[m.RoomID].Clients {
						select {
						case client.Message <- m:
						default:
							log.Printf("Failed to send message to client %s", client.ID)
						}
					}
				}
			}
		}

	}
}

func (h *Hub) JoinRoom(roomId, clientId string) {
    client, ok := h.Rooms["sociala"].Clients[clientId]
    if !ok {
        log.Printf("Client %s not found in the 'sociala' room", clientId)
        return
    }

    room, ok := h.Rooms[roomId]
    if !ok {
        log.Printf("Room %s not found", roomId)
        return
    }

    room.Clients[clientId] = client
}

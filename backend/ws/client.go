package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string          `json:"id"`
	Username string          `json:"username"`
	Rooms    map[string]bool `json:"rooms"`
}

type Message struct {
	Content     string `json:"content"`
	RoomID      string `json:"roomId"`
	Username    string `json:"username"`
	SenderID    string `json:"senderId"`
	RecipientID string `json:"recipientId"`
	Action      string `json:"action"`
	Type        string `json:"type"`
	TimeStamp   string `json:"time"`
}

func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// Deserialize the received message
		var receivedMessage Message
		err = json.Unmarshal(m, &receivedMessage)
		if err != nil {
			log.Printf("Error deserializing message: %v", err)
			continue
		}

		msg := &Message{
			Content:     receivedMessage.Content,
			RoomID:      receivedMessage.RoomID,
			Username:    c.Username,
			SenderID:    c.ID,
			RecipientID: receivedMessage.RecipientID,
			Action:      receivedMessage.Action,
			TimeStamp:   time.Now().Format("Jan 2, 2006 at 3:04pm"),
		}

		hub.Broadcast <- msg
	}
}

// func (m *Message) Store() error {

// }

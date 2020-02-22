package controllers

import (
	"github.com/google/uuid"
	"log"
	"github.com/KristiyanGK/cloudcooking/models"
	"github.com/gorilla/websocket"
	"sync"
	"net/http"
)

type ConnectionMap struct {
	clients map[*websocket.Conn]bool
	lock sync.RWMutex
}

func (c *ConnectionMap) Store(client *websocket.Conn, connected bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.clients[client] = connected
}

func (c *ConnectionMap) Delete(client *websocket.Conn) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.clients, client)
}

func (c *ConnectionMap) SendMsg(msg models.Message) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	for client := range c.clients {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			c.Delete(client)
		}
	}
}

func init() {
	go handleMessages()
}

var clients ConnectionMap = ConnectionMap{make(map[*websocket.Conn]bool), sync.RWMutex{}}
var broadcast = make(chan models.Message)
var upgrader = websocket.Upgrader {
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (a *App) Chat(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()
	clients.Store(ws, true)

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		msg.ID = uuid.New().String()
		
		if err != nil {
			log.Printf("error: %v", err)
			clients.Delete(ws)
			break
		}

		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <- broadcast

		clients.SendMsg(msg)
	}
}
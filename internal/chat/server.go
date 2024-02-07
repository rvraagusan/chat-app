package chat

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
}

func NewServer() *Server {
	return &Server{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
}

func (s *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	s.clients[conn] = true

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(s.clients, conn)
			break
		}
		s.broadcast <- msg
	}
}

func (s *Server) HandleMessages() {
	for {
		msg := <-s.broadcast
		for client := range s.clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println(err)
				client.Close()
				delete(s.clients, client)
			}
		}
	}
}

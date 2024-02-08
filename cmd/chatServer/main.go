package main

import (
	"log"
	"net/http"

	"github.com/rvraagusan/chat-app/internal/chat"
)

func main() {
	server := chat.NewServer()
	http.HandleFunc("/ws", server.HandleConnections)
	go server.HandleMessages()

	log.Println("Chat server started on :7003")
	err := http.ListenAndServe(":7003", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

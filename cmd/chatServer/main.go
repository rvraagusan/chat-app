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

	log.Println("Chat server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	roomManager := NewRoomManager()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		roomManager.WebSocketHandler(w, r)
	})
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		roomManager.CreateRoomHandler(w, r)
	})
	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		roomManager.JoinRoomHandler(w, r)
	})
	log.Println("server listening on :8000")
	http.ListenAndServe(":8000", nil)
}

func init() {
	// Seed randomness for RandString
	rand.Seed(time.Now().UTC().UnixNano())
}

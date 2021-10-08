package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		WebSocketHandler(w, r)
	})
	log.Println("server listening on :8000")
	http.ListenAndServe(":8000", nil)
}

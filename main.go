package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", time.Now(), r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	roomManager := NewRoomManager()

	mux := http.NewServeMux()

	mux.Handle("/ws", loggerMiddleware(http.HandlerFunc(roomManager.WebSocketHandler)))
	mux.Handle("/create", loggerMiddleware(http.HandlerFunc(roomManager.CreateRoomHandler)))
	mux.Handle("/join", loggerMiddleware(http.HandlerFunc(roomManager.JoinRoomHandler)))

	log.Println("server listening on :8000")
	http.ListenAndServe(":8000", mux)
}

func init() {
	// Seed randomness for RandString
	rand.Seed(time.Now().UTC().UnixNano())
}

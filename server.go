package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func (roomManager *RoomManager) WebSocketHandler(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		msg := WSMessage{}
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println(err)
			return
		}

		log.Println(msg)
	}
}

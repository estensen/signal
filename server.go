package main

import (
	"encoding/json"
	"fmt"
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

func (r *RoomManager) WebSocketHandler(w http.ResponseWriter, req *http.Request) {
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

func (r *RoomManager) CreateRoomHandler(w http.ResponseWriter, req *http.Request) {
	room := r.createRoom()

	type resp struct {
		RoomID string `json:"roomID"`
	}

	json.NewEncoder(w).Encode(resp{RoomID: room.ID})
}

func (r *RoomManager) JoinRoomHandler(w http.ResponseWriter, req *http.Request) {
	roomID, ok := req.URL.Query()["roomID"]
	if !ok {
		fmt.Println("roomID missing in URL parameters")
		return
	}

	userID := RandString(4)
	userInfo := UserInfo{
		ID:   userID,
		Name: "",
	}

	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatalln("web socket upgrader error", err)
	}

	user := User{
		info: userInfo,
		conn: ws,
	}

	r.joinRoom(roomID[0], user)
}

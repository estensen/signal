package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

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

func broadcaster(room *Room, sender string) {
	for {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		// Don't yo yourself
		for _, client := range room.users {
			if client.info.ID != sender {
				err := client.conn.WriteJSON(fmt.Sprintf("yo from %s", sender))
				if err != nil {
					fmt.Println("could not broadcast", err)
					client.conn.Close()
				}
			}
		}
	}
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

	room := r.getRoom(roomID[0])

	// Sometimes yo other people
	go broadcaster(room, userID)

	// Listen forever for yo's
	for {
		var msg string
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("could not listen for yo's :/", err)
		}

		//fmt.Printf("%s: other machine said: %s", userID, msg)
	}
}

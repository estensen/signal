package main

import (
	"github.com/gorilla/websocket"
)

type UserInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	info UserInfo
	conn *websocket.Conn
}

type Session struct {
	ID   string
	from User
	to   User
}

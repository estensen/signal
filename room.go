package main

import "sync"

type Room struct {
	users   map[string]User
	session map[string]Session
	ID      string
}

func NewRoom(id string) *Room {
	var room = &Room{
		users:   make(map[string]User),
		session: make(map[string]Session),
		ID:      id,
	}
	return room
}

type RoomManager struct {
	rooms    map[string]*Room
	roomLock sync.RWMutex
}

func NewRoomManager() *RoomManager {
	var roomManager = &RoomManager{
		rooms: make(map[string]*Room),
	}
	return roomManager
}

func (roomManager *RoomManager) getRoom(id string) *Room {
	roomManager.roomLock.RLock()
	defer roomManager.roomLock.Unlock()

	return roomManager.rooms[id]
}

func (roomManager *RoomManager) createRoom() *Room {
	roomManager.roomLock.Lock()
	defer roomManager.roomLock.Unlock()

	id := RandString(6)

	// TODO: Handle generating same ID
	roomManager.rooms[id] = NewRoom(id)
	return roomManager.rooms[id]
}

func (roomManager *RoomManager) deleteRoom(id string) {
	roomManager.roomLock.Lock()
	defer roomManager.roomLock.Unlock()

	delete(roomManager.rooms, id)
}

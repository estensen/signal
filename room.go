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
	rooms map[string]*Room
	mutex sync.RWMutex
}

func NewRoomManager() *RoomManager {
	var roomManager = &RoomManager{
		rooms: make(map[string]*Room),
	}
	return roomManager
}

func (r *RoomManager) getRoom(id string) *Room {
	r.mutex.RLock()
	defer r.mutex.Unlock()

	return r.rooms[id]
}

func (r *RoomManager) createRoom() *Room {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	id := RandString(6)

	// TODO: Handle generating same ID
	r.rooms[id] = NewRoom(id)
	return r.rooms[id]
}

func (r *RoomManager) joinRoom(roomID string, user User) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.rooms[roomID].users[user.info.ID] = user
}

func (r *RoomManager) deleteRoom(id string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	delete(r.rooms, id)
}

package main

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
}

func NewRoomManager() *RoomManager {
	var roomManager = &RoomManager{
		rooms: make(map[string]*Room),
	}
	return roomManager
}

func (roomManager *RoomManager) getRoom(id string) *Room {
	return roomManager.rooms[id]
}

func (roomManager *RoomManager) createRoom(id string) *Room {
	roomManager.rooms[id] = NewRoom(id)
	return roomManager.rooms[id]
}

func (roomManager *RoomManager) deleteRoom(id string) {
	delete(roomManager.rooms, id)
}

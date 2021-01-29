package user

import (
	"log"
)

const (
	EnterRoom   = 1
	LeaveRoom   = 2
	SendMessage = 3
)

type Command struct {
	command int
	client  *Client
	message []byte
}

type Room struct {
	id       string
	name     string
	commands chan Command
	clients  map[*Client]bool // all clients
}

func NewRoom(id string, name string) *Room {
	return &Room{
		id:       id,
		name:     name,
		commands: make(chan Command),
		clients:  make(map[*Client]bool),
	}
}

func (room *Room) EnterRoom(client *Client) {
	room.commands <- Command{command: EnterRoom, client: client}
}

func (room *Room) LeaveRoom(client *Client) {
	room.commands <- Command{command: LeaveRoom, client: client}
}

func (room *Room) Run() {
	for {
		command := <-room.commands

		switch command.command {
		case EnterRoom:
			room.clients[command.client] = true
			log.Println("Client joined room", room.name)
		case LeaveRoom:
			delete(room.clients, command.client)
			close(command.client.send)
			log.Println("Client left room", room.name)
		case SendMessage:
			for client := range room.clients {
				select {
				case client.send <- command.message:
					// message sent
					log.Println("Sending message to client")
				default:
					// failed to send message
					log.Println("Failed to send message to client")
					delete(room.clients, client)
					close(client.send)
				}
			}
		}
	}
}

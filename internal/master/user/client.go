package user

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	user       *User
	connection *websocket.Conn
	Room       *Room
	send       chan []byte
}

func NewClient(connection *websocket.Conn) *Client {
	return &Client{
		connection: connection,
		send:       make(chan []byte, 1024),
	}
}

func (client *Client) Run() {
	go client.write()
	client.read()
}

func (client *Client) read() {
	for {
		messageType, message, err := client.connection.ReadMessage()
		if err == nil {

			if messageType == websocket.TextMessage {
				log.Printf("Received %s\n", message)
				client.Room.commands <- Command{command: SendMessage, client: client, message: message}
			} else if messageType == websocket.BinaryMessage {

			}

		} else {
			//if err != io.EOF {
			log.Println("Failed to read message", err)
			//}
			break
		}
	}

	client.connection.Close()
}

func (client *Client) write() {
	for message := range client.send {
		err := client.connection.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			//if err != io.EOF {
			log.Println("Failed to write message", err)
			//}
			break
		}
	}

	client.connection.Close()
}

func (client *Client) logIn(username string, password string) {
	// TODO: get user from database and assign to user
}

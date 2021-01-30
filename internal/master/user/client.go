package user

import (
	"fmt"
	"log"
	"net"
)

type Client struct {
	user       *User
	connection net.Conn
	Room       *Room
	send       chan []byte
}

func NewClient(connection net.Conn) *Client {
	return &Client{
		connection: connection,
		send:       make(chan []byte, 1024),
	}
}

func (client *Client) Run() {
	defer client.connection.Close()

	go client.write()
	client.read()
}

func (client *Client) read() {
	buffer := make([]byte, 256)

	for {
		bytesRead, readError := client.connection.Read(buffer)

		if readError != nil {
			fmt.Println("Client left.")
			return
		}

		fmt.Println("Read", bytesRead, "bytes")
	}
}

func (client *Client) write() {
	for message := range client.send {
		bytesWritten, writeError := client.connection.Write(message)

		if writeError != nil {
			//if err != io.EOF {
			log.Println("Failed to write message", writeError)
			//}
			break
		}

		fmt.Println("Written", bytesWritten, "bytes")
	}
}

package master

import (
	"fmt"
	"net"
	"os"

	"gitlab.lan/minilooter/server/internal/master/user"
)

type Server struct {
	host     string
	port     uint64
	listener net.Listener
	rooms    *user.Rooms
	room     *user.Room
}

type IndexData struct {
	LoggedIn bool
}

func NewServer(host string, port uint64) *Server {
	server := &Server{
		host:  host,
		port:  port,
		rooms: user.NewRooms(),
	}

	server.room = server.rooms.CreateRoom("test")

	return server
}

func (server *Server) Run() {
	var listenError error
	server.listener, listenError = net.Listen("tcp", fmt.Sprintf("%s:%d", server.host, server.port))
	if listenError != nil {
		fmt.Println("Error listening:", listenError.Error())
		os.Exit(1)
	}
	defer server.listener.Close()

	fmt.Println("Server started on", server.host, server.port)

	for {
		client, acceptError := server.listener.Accept()
		if acceptError != nil {
			fmt.Println("Error connecting:", acceptError.Error())
			return
		}
		fmt.Println("Client connected.")

		fmt.Println("Client " + client.RemoteAddr().String() + " connected.")
	}

	//go server.room.Run()
}

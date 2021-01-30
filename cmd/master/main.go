package main

import (
	"log"
	"os"
	"strconv"

	"gitlab.lan/minilooter/server/internal/master"
)

func main() {
	log.Println("Server version 0.1")

	port, _ := strconv.ParseUint(os.Getenv("MINILOOTER_PORT"), 0, 16)

	server := master.NewServer(os.Getenv("MINILOOTER_HOST"), port)
	server.Run()
}

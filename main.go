package main

import (
	"github.com/Surya-7890/questions_service/app/server"
)

func main() {
	newServer := server.NewServer(":7000")
	newServer.Init()
	newServer.StartServer()
}

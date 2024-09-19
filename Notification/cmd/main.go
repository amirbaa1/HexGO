package main

import (
	"Notification/cmd/server"
)

func main() {
	go server.StartRabbitMQ()

	server.Internal()
}

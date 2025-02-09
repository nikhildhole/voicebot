package main

import (
	"log"

	"voicebot/internal/handler"
	"voicebot/internal/server"
)

func main() {
	// Pass the handler function when creating the server
	srv := server.New(":8080", handler.HandleConnection)

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

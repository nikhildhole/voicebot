package handler

import (
	"io"
	"net"
	"voicebot/internal/protocol"
	"voicebot/pkg/logger"
)

// Register handlers
var handlers = map[byte]MessageHandler{
	protocol.TypeAudio: &AudioHandler{},
	protocol.TypeUUID:  &UUIDHandler{},
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		msg, err := protocol.ReadMessage(conn)
		if err != nil {
			if err == io.EOF {
				logger.Info("Client disconnected: %s", conn.RemoteAddr())
			} else {
				logger.Error("Error reading message: %v", err)
			}
			return
		}

		if handler, exists := handlers[msg.Type]; exists {
			err := handler.Handle(conn, msg.Payload)
			if err != nil {
				logger.Error("Handler error: %v", err)
			}
		} else {
			logger.Error("Unknown packet type received: %x", msg.Type)
		}
	}
}

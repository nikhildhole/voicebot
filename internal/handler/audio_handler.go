package handler

import (
	"net"
	"voicebot/internal/protocol"
	"voicebot/pkg/logger"
)

type AudioHandler struct{}

func (h *AudioHandler) Handle(conn net.Conn, payload []byte) error {
	logger.Info("Received Audio Data: %d bytes", len(payload))
	return protocol.SendMessage(conn, protocol.TypeAudio, payload) // Fixes improper response format
}

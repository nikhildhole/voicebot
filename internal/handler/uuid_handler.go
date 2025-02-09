package handler

import (
	"net"
	"voicebot/pkg/logger"
)

type UUIDHandler struct{}

func (h *UUIDHandler) Handle(conn net.Conn, payload []byte) error {
	logger.Info("Received UUID: %x", payload)
	return nil
}

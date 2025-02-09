package handler

import "net"

type MessageHandler interface {
	Handle(conn net.Conn, payload []byte) error
}

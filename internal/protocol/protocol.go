package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const (
	TypeTerminate = 0x00
	TypeUUID      = 0x01
	TypeAudio     = 0x10
	TypeError     = 0xFF
	HeaderSize    = 3
)

type Message struct {
	Type    byte
	Length  uint16
	Payload []byte
}

// ReadMessage reads and parses a message from a TCP connection.
func ReadMessage(conn net.Conn) (*Message, error) {
	header := make([]byte, HeaderSize)
	_, err := io.ReadFull(conn, header)
	if err != nil {
		return nil, err
	}

	msgType := header[0]
	length := binary.BigEndian.Uint16(header[1:])

	payload := make([]byte, length)
	_, err = io.ReadFull(conn, payload)
	if err != nil {
		return nil, err
	}

	return &Message{
		Type:    msgType,
		Length:  length,
		Payload: payload,
	}, nil
}

// SendMessage constructs and sends a message.
func SendMessage(conn net.Conn, msgType byte, payload []byte) error {
	length := uint16(len(payload))
	message := append([]byte{msgType}, make([]byte, 2)...)
	binary.BigEndian.PutUint16(message[1:], length)
	message = append(message, payload...)

	_, err := conn.Write(message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

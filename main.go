package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
)

const (
	TypeTerminate = 0x00
	TypeUUID      = 0x01
	TypeAudio     = 0x10
	TypeError     = 0xFF
	HeaderSize    = 3
)

type Client struct {
	conn net.Conn
	id   string
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()

	for {
		header := make([]byte, HeaderSize)
		_, err := io.ReadFull(conn, header)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
			} else {
				fmt.Println("Error reading header:", err)
			}
			return
		}

		typeByte := header[0]
		length := binary.BigEndian.Uint16(header[1:])

		payload := make([]byte, length)
		_, err = io.ReadFull(conn, payload)
		if err != nil {
			fmt.Println("Error reading payload:", err)
			return
		}

		switch typeByte {
		case TypeTerminate:
			fmt.Println("Received terminate signal")
			return
		case TypeUUID:
			fmt.Printf("Received UUID: %x\n", payload)
		case TypeAudio:
			fmt.Printf("Received Audio Data: %d bytes\n", len(payload))
			// Echo the audio back to the client
			response := append([]byte{TypeAudio}, header[1:]...)
			response = append(response, payload...)
			_, err = conn.Write(response)
			if err != nil {
				fmt.Println("Error sending audio back:", err)
			}
		case TypeError:
			fmt.Printf("Received Error Code: %x\n", payload)
		default:
			fmt.Println("Unknown packet type received")
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("AudioSocket server listening on port 8080...")

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("New client connected")
		wg.Add(1)
		go handleConnection(conn, &wg)
	}
}

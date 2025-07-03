package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	users = map[string]string{
		"lati": "password123",
		"bob":   "qwerty",
	}
	mu sync.Mutex
)

func main() {
	listener, err := net.Listen("tcp", ":9002")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Authentication server started on :9002")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		fmt.Printf("Client disconnected: %s\n", conn.RemoteAddr())
		conn.Close()
	}()

	fmt.Printf("New connection from: %s\n", conn.RemoteAddr())
	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			sendResponse(conn, "ERROR: Invalid command")
			continue
		}

		cmd := strings.ToUpper(parts[0])
		username := parts[1]

		switch cmd {
		case "LOGIN":
			handleLogin(conn, username, parts)
		case "REGISTER":
			handleRegister(conn, username, parts)
		default:
			sendResponse(conn, "ERROR: Unknown command")
		}
	}
}

func handleLogin(conn net.Conn, username string, parts []string) {
	if len(parts) != 3 {
		sendResponse(conn, "LOGIN requires username and password")
		return
	}

	password := parts[2]

	mu.Lock()
	storedPassword, exists := users[username]
	mu.Unlock()

	if exists && storedPassword == password {
		sendResponse(conn, "LOGIN SUCCESS")
	} else {
		sendResponse(conn, "LOGIN FAILED")
	}
}

func handleRegister(conn net.Conn, username string, parts []string) {
	if len(parts) != 3 {
		sendResponse(conn, "REGISTER requires username and password")
		return
	}

	mu.Lock()
	_, exists := users[username]
	if exists {
		mu.Unlock()
		sendResponse(conn, "REGISTER FAILED: User exists")
		return
	}

	users[username] = parts[2]
	mu.Unlock()

	sendResponse(conn, "REGISTER SUCCESS")
}

func sendResponse(conn net.Conn, message string) {
	conn.Write([]byte(message + "\n"))
}

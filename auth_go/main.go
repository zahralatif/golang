package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"
	"strings"
	"sync"
)

var (
	users = make(map[string]string)
	mu    sync.RWMutex
)

func main() {
	port := flag.String("port", "9002", "Server port")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer ln.Close()
	log.Printf("Server started on port %s", *port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Connection closed: %v", err)
			return
		}

		data = strings.TrimSpace(data)
		if data == "" {
			continue
		}

		var req struct {
			Command  string `json:"command"`
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.Unmarshal([]byte(data), &req); err != nil {
			sendResponse(conn, "FAIL", "Invalid JSON format")
			continue
		}

		switch strings.ToUpper(req.Command) {
		case "LOGIN":
			handleLogin(conn, req.Username, req.Password)
		case "REGISTER":
			handleRegister(conn, req.Username, req.Password)
		default:
			sendResponse(conn, "FAIL", "Unknown command")
		}
	}
}

func sendResponse(conn net.Conn, status, message string) {
	resp := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: message,
	}
	data, _ := json.Marshal(resp)
	conn.Write(append(data, '\n'))
}
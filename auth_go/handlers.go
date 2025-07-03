package main

import (
	"net"
)

func handleLogin(conn net.Conn, username, password string) {
	mu.RLock()
	storedPassword, exists := users[username]
	mu.RUnlock()

	if !exists || storedPassword != password {
		sendResponse(conn, "FAIL", "LOGIN FAILED")
		return
	}
	sendResponse(conn, "OK", "LOGIN SUCCESS")
}

func handleRegister(conn net.Conn, username, password string) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[username]; exists {
		sendResponse(conn, "FAIL", "REGISTER FAILED: User exists")
		return
	}

	users[username] = password
	sendResponse(conn, "OK", "REGISTER SUCCESS")
}
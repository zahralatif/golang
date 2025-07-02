package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

var store = make(map[string]string)

func main() {
    listener, err := net.Listen("tcp", ":9001")
    if err != nil {
        fmt.Println("Error starting server:", err)
		return
    }
    defer listener.Close()
    fmt.Println("Key-Value server started on :9001")

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            return
        }

        line = strings.TrimSpace(line)
        parts := strings.Split(line, " ")
        if len(parts) < 1 {
            conn.Write([]byte("ERROR: Invalid command\n"))
            continue
        }

        cmd := strings.ToUpper(parts[0])

        switch cmd {
        case "SET":
            if len(parts) != 3 {
                conn.Write([]byte("SET command requires a value\n"))
                continue
            }
            key := parts[1]
			value := parts[2]
			store[key] = value
            conn.Write([]byte("OK\n"))
        case "GET":
            if len(parts) != 2 {
				conn.Write([]byte("ERROR: Usage: GET key\n"))
				continue
			}
			key := parts[1]
			value, ok := store[key]
            if ok {
                conn.Write([]byte(value + "\n"))
            } else {
                conn.Write([]byte("Key not found\n"))
            }
        default:
            conn.Write([]byte("Unknown command\n"))
        }
    }
}


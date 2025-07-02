package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", ":9000")
    if err != nil {
        fmt.Println("Error starting server:", err)
		return
    }
    defer listener.Close()
    fmt.Println("Server started on :9000")

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
        msg, err := reader.ReadString('\n')
        if err != nil {
            return
        }
        fmt.Print("Received:", msg)
        conn.Write([]byte("Echo: " + msg))
    }
}

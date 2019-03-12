package main

import (
	"fmt"
	"net"
)

func udp1() {
    conn, _ := net.ListenPacket("udp", "localhost:9001")
    defer conn.Close()

    buffer := make([]byte, 1500)
    for {
        n, addr, _ := conn.ReadFrom(buffer)
        fmt.Printf("Received from %v: %v\n", addr, string(buffer[:n]))
    }
}

func main() {
	println("hello client")
}

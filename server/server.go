package main

import (
	"fmt"
	"net"
)

func main() {
    udp, _ := net.ListenPacket("udp", "localhost:9001")
    defer udp.Close()
    tcp, _ := net.Listen("tcp", ":9002")
    defer tcp.Close()

    go func(conn net.PacketConn) {
        for {
            buffer := make([]byte, 1500)
            n, addr, _ := conn.ReadFrom(buffer)
            fmt.Printf("UDP from %v: %v\n", addr, string(buffer[:n]))
        }
    }(udp)

    for {
        client, _ := tcp.Accept()
        defer client.Close()
        buffer := make([]byte, 1500)
        n, _ := client.Read(buffer)
        fmt.Printf("TCP %v\n", string(buffer[:n]))
        client.Write([]byte("OK"))
    }
}

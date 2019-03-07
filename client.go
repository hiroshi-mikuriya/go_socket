package main

import (
  "net"
  "fmt"
)

func tcp(str string) {
  conn, err := net.Dial("tcp", "127.0.0.1:9002")
  if err != nil {
    fmt.Printf("Dial error: %s\n", err)
    return
  }
  defer conn.Close()
  conn.Write([]byte(str))
}

func udp(str string) {
  udpAddr, err := net.ResolveUDPAddr("udp", "localhost:9001")
  if err != nil {
    fmt.Printf("ResolveUDPAddr error: %s\n", err)
    return
  }
  conn, err := net.DialUDP("udp", nil, udpAddr)
  if err != nil {
    fmt.Printf("DialUDP error %s\n", err)
    return
  }
  conn.Write([]byte(str))
}

func main() {
  tcp("hello tcp")
  udp("hello udp")
}


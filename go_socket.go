package main

import (
  "net"
  "fmt"
)

func stream(af, dst, str string) {
  conn, err := net.Dial(af, dst)
  if err != nil {
    fmt.Printf("Dial error: %s\n", err)
    return
  }
  defer conn.Close()
  conn.Write([]byte(str))
  buf := make([]byte, 1024)
  n, err := conn.Read(buf[:])
  if err != nil {
    fmt.Printf("error Read %s\n", err)
    return
  }
  print(string(buf[0:n]))
}

func dgram(af, dst, str string) {
  addr, err := net.ResolveUDPAddr(af, dst)
  if err != nil {
    fmt.Printf("ResolveUDPAddr error: %s\n", err)
    return
  }
  conn, err := net.DialUDP(af, nil, addr)
  if err != nil {
    fmt.Printf("DialUDP error %s\n", err)
    return
  }
  conn.Write([]byte(str))
}

func unix_dgram(af, file, str string) {
  addr := net.UnixAddr{ file, af }
  conn, err := net.DialUnix(af, nil, &addr)
  if err != nil {
    fmt.Printf("error net.DialUnix %s\n", err)
    return
  }
  defer conn.Close()
  _, err = conn.Write([]byte(str))
  if err != nil {
    fmt.Printf("error Write %s\n", err)
  }
  if af != "unix" {
    return
  }
  buf := make([]byte, 1024)
  n, err := conn.Read(buf[:])
  if err != nil {
    fmt.Printf("error Read %s\n", err)
    return
  }
  print(string(buf[0:n]))
}

func main() {
  stream("tcp", "127.0.0.1:9002", "hello tcp")
  dgram("udp", "localhost:9001", "hello udp")
  stream("unix", "/tmp/unix_stream.socket", "hello unix stream1")
  unix_dgram("unixgram", "/tmp/unix_dgram.socket", "hello unix dgram")
  unix_dgram("unix", "/tmp/unix_stream.socket", "hello unix stream2")
}


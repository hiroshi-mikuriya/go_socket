package main

import (
  "net"
  "fmt"
)

func read(conn net.Conn) (string) {
  buf := make([]byte, 1024)
  n, err := conn.Read(buf[:])
  if err != nil {
    fmt.Printf("error Read %s\n", err)
    return ""
  }
  return string(buf[0:n])
}

func stream(af, dst, str string) {
  conn, err := net.Dial(af, dst)
  if err != nil {
    fmt.Printf("Dial error: %s\n", err)
    return
  }
  defer conn.Close()
  conn.Write([]byte(str))
  println(read(conn))
}

func udp1(dst, str string) {
  addr, err := net.ResolveUDPAddr("udp", dst)
  if err != nil {
    fmt.Printf("ResolveUDPAddr error: %s\n", err)
    return
  }
  conn, err := net.DialUDP("udp", nil, addr)
  if err != nil {
    fmt.Printf("DialUDP error %s\n", err)
    return
  }
  defer conn.Close()
  conn.Write([]byte(str))
}

func udp2(dst, str string) {
  conn, err := net.Dial("udp", dst)
  if err != nil {
    fmt.Printf("Dial error: %s\n", err)
    return
  }
  defer conn.Close()
  conn.Write([]byte(str))
}

func udp3(dst, str string) {
  addr, err := net.ResolveIPAddr("", "localhost")
  // addr, err := net.ResolveUDPAddr("udp", dst)
  // addr, err := net.ResolveTCPAddr("tcp", dst)
  if err != nil {
    fmt.Printf("ResolveIPAddr error: %s\n", err)
    return
  }
  udp2(addr.String(), str)
}

func unix_stream(dst, str string) {
  conn, err := net.Dial("unix", dst)
  if err != nil {
    fmt.Printf("Dial error: %s\n", err)
    return
  }
  defer conn.Close()
  conn.Write([]byte(str))
  println(read(conn))
}

func unix_dgram(dst, str string) {
  conn, err := net.Dial("unixgram", dst)
  if err != nil {
    fmt.Printf("Dial error: %s\n", err)
    return
  }
  defer conn.Close()
  conn.Write([]byte(str))
}

func unix(af, file, str string) {
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
    return
  }
  if af == "unix" {
    println(read(conn))
  }
}

func main() {
  stream("tcp", "127.0.0.1:9002", "hello tcp")
  udp1("localhost:9001", "hello udp1")
  udp2("127.0.0.1:9001", "hello udp2")
  udp3("localhost:9001", "hello udp3")
  stream("unix", "/tmp/unix_stream.socket", "hello unix stream1")
  unix("unixgram", "/tmp/unix_dgram.socket", "hello unix dgram")
  unix("unix", "/tmp/unix_stream.socket", "hello unix stream2")
  unix_dgram("/tmp/unix_dgram.socket", "hello unix dgram2")
  unix_stream("/tmp/unix_stream.socket", "hello unix stream3")
}

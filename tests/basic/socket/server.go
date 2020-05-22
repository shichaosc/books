package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	server, _ := net.Listen("tcp", "127.0.0.1:38489")
	fmt.Println(server.Addr())
	for {
		conn, _ := server.Accept()
		go handleConnection(conn)
	}
}

//noinspection ALL
func handleConnection(conn net.Conn) {
	// 必须使用Make初始化容量
	b := make([]byte, 2)
	for {
		n, err := conn.Read(b)
		if err != nil || n == 0 {
			conn.Close()
			break
		}
		commends := strings.Split(strings.TrimSpace(string(b[0:n])), " ")
		switch commends[0] {
		case "ping":
			conn.Write([]byte("pong"))
		case "pong":
			conn.Write([]byte("ping"))
		case "quit":
			conn.Close()
		default:
			fmt.Printf("unsupport commend %s", commends[0])
		}
	}
}

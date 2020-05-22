package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:38489")
	fmt.Println(conn.Write([]byte("ping")))
	b := make([]byte, 5)
	n, _ := conn.Read(b)
	fmt.Println(string(b[0:n]))
	conn.Close()
}

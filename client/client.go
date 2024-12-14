package client

import "net"

func Connect() {
	conn, _ := net.Dial("tcp", "localhost:9999")

	_, err := conn.Write([]byte("Hello there\n"))
	if err != nil {
		return
	}
}

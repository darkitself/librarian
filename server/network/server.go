package network

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

type Server struct {
	port int
}

func NewServer(port int) *Server {
	return &Server{port: port}
}

func (serv *Server) Start() {
	// определяем порт для прослушивания
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(serv.port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// закрываем listener при завершении программы
	defer listener.Close()
	fmt.Println("Server is listening on " + strconv.Itoa(serv.port))

	for {
		// принимаем входящее подключение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Connected with", conn.RemoteAddr().String())
		// обрабатываем подключение в отдельной горутине
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	// читаем данные от клиента
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		clientMessage := scanner.Text()
		fmt.Printf("Received from client: %s\n", clientMessage)
		// отправляем ответ клиенту
		conn.Write([]byte("Message received.\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err.Error())
	}
}

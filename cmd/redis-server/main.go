package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/velenyak/redis-server/config"
)

func main() {
	config := config.New()
	listener, err := net.Listen(config.Protocol, fmt.Sprintf(":%d", config.Port))
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("Server listening on port", config.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error handling request", err.Error())
			panic(err)
		}
		reader := bufio.NewReader(conn)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading request", err.Error())
			panic(err)
		}
		fmt.Println("Request recieved", line)
		conn.Write([]byte("PONG"))
		conn.Close()
	}
}

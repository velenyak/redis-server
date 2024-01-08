package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/velenyak/redis-server/config"
	"github.com/velenyak/redis-server/internal/handler"
	"github.com/velenyak/redis-server/internal/resp"
)

func main() {
	config := config.New()
	listener, err := net.Listen(config.Protocol, fmt.Sprintf(":%d", config.Port))
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("Server listening on port", config.Port)

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error handling request", err.Error())
		panic(err)
	}

	for {
		respReader := resp.New(bufio.NewReader(conn))
		line, err := respReader.Read()
		if err != nil {
			fmt.Println("Error reading request", err.Error())
			panic(err)
		}
		fmt.Println("Request recieved", line)
		result, err := handler.HandleResp(line)
		if err != nil {
			fmt.Println("Error handling request", err.Error())
			panic(err)
		}
		response, err := resp.Write(result)
		if err != nil {
			fmt.Println("Error writing response", err.Error())
			panic(err)
		}
		conn.Write([]byte(response))
	}
}

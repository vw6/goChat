package main

import (
	"bufio"
	"container/list"
	"fmt"
	"net"
)

var clients *list.List

func handelClient(socket net.Conn) {
	for {
		buffer, err := bufio.NewReader(socket).ReadString('\n')
		if err != nil {
			fmt.Println("disconect")
			socket.Close()
			return
		}
		for i := clients.Front(); i != nil; i = i.Next() {
			fmt.Println(i.Value.(net.Conn), buffer)
		}
	}
}

func main() {
	fmt.Println("Server started")

	var clients = list.New()

	server, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error @s", err.Error())
		return
	}
	for {
		client, err := server.Accept()

		if err != nil {
			fmt.Println("Error: @s", err.Error())
			return
		}
		fmt.Println("Connect", client.RemoteAddr())
		if err != nil {
			clients.PushBack(client)
		}

		go handelClient(client)
	}
}

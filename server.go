package main
import (
"fmt"
"net"
"bufio"
"container/list"
)

var clients *list.List

func handelClient(socket net.Conn)  {
	for{
		buffer, err := bufio.NewReader(socket).ReadString('\n')
		if err!= nil {
			fmt.Println("disconect")
			socket.Close()
			return
		}
		for i:=clients.Front(); i != nil ; i=i.Next(){
			fmt.Print(i.Value.(net.Conn),buffer)
		}
	}
}


func main() {
fmt.Println("Server started")
	
	clients = list.New()

server,err := net.Listen("tcp",":8080")

	if err != nil {
		fmt.Printf("Error @s", err.Error())
		return
	}
	for{
		client, err := server.Accept()

		if err != nil {
			fmt.Printf("Error: @s", err.Error())
			return
		}
		fmt.Printf("Connect",client.RemoteAddr())
		clients.PushBack(client)


		go handelClient(client)
	}
}
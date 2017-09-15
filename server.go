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
		//fmt.Println("2")
		buffer, err := bufio.NewReader(socket).ReadString('\n')

		//fmt.Println("3")

		for i:=clients.Front(); i != nil ; i=i.Next(){
			//fmt.Println("4")
			fmt.Println(i.Value.(net.Conn),buffer)
			//fmt.Println(buffer)
			fmt.Print(buffer)
			fmt.Println("\n disconect")
			socket.Close()
			return
		}
		if err!= nil {
			fmt.Println("disconect")
			socket.Close()
			return
		}
	}
}


func main() {
	fmt.Println("Server started")

	clients = list.New()

	server,err := net.Listen("tcp",":80")

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

		//fmt.Println("1")
		go handelClient(client)
	}
}
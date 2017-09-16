package main

import (
"fmt"
"net"
"os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "8080"
    CONN_TYPE = "tcp"
)

func main() {
    var i int = 0
    var conns [2]net.Conn
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        if (i == 0) {
            conns[0], err = l.Accept()
            if err != nil {
                fmt.Println("Error accepting: ", err.Error())
                os.Exit(1)
            }
            // Handle connections in a new goroutine.
            fmt.Println("Connect: ", conns[0].RemoteAddr())
            i++
        } else {
            conns[1], err = l.Accept()
            if err != nil {
                fmt.Println("Error accepting: ", err.Error())
                os.Exit(1)
            }
            // Handle connections in a new goroutine.
            fmt.Println("Connect: ", conns[1].RemoteAddr())
            go handleRequest(conns)
            i = 0
        }
    }
}

// Handles incoming requests.
func handleRequest(conns [2]net.Conn) {
    // Make a buffer to hold incoming data.
    buf := make([]byte, 1024)
    // Read the incoming connection into the buffer.
    _, err := conns[0].Read(buf)
    if err != nil {
        fmt.Println("Error reading:", err.Error())
    }
    // Send a response back to person contacting us.
    
    conns[1].Write(buf)

    // Close the connection when you're done with it.
    fmt.Println("disconect: ", conns[0].RemoteAddr())
    fmt.Println("disconect: ", conns[1].RemoteAddr())
    conns[0].Close()
    conns[1].Close()
}

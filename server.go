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

// var conns [2]net.Conn

func main() {
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
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
            // Handle connections in a new goroutine.
        fmt.Println("Connect: ", conn.RemoteAddr())
        conn.Write([]byte("Соединение установленно!1\n"))
        conn.Write([]byte("Соединение установленно!2\n"))
        conn.Write([]byte("Соединение установленно!3\n"))
        conn.Write([]byte("Соединение установленно!4\n"))
        conn.Write([]byte("Соединение установленно!5\n"))
        conn.Write([]byte("Соединение установленно!6\n"))
        conn.Write([]byte("Соединение установленно!7\n"))
        conn.Write([]byte("Соединение установленно!8\n"))
        conn.Write([]byte("Соединение установленно!9\n"))
        conn.Write([]byte("Соединение установленно!0\n"))
        go handleRequest(conn)
    }
}


// Handles incoming requests.
func handleRequest(conn net.Conn) {
    // Make a buffer to hold incoming data.
    buf := make([]byte, 1024)

    // Read the incoming connection into the buffer.
    _, err := conn.Read(buf)
    if err != nil {
        fmt.Println("Error reading:", err.Error())
    }

    // Send a response back to person contacting us.
    conn.Write(buf)

    // Close the connection when you're done with it.
    defer fmt.Println("disconect: ", conn.RemoteAddr())
    defer conn.Close()
}

package main

import(
	"fmt"
	"net"
	 "time"
)

const (
	server = "129.241.187.xxx"
	port_fixed = "34933"
	port_delim = "33546"
)

func client() {
    connAddr, err := net.ResolveTCPAddr("tcp","129.241.187.136:33546")
    if err != nil {
    	Printf("Error in TCP: %s\n", err.Error())
    }

    connTCP, err := net.DialTCP("tcp", nil, connAddr)
    if err != nil {
    Printf("Error in TCP: %s\n", err.Error())
    }

    message := []byte("Connect to:129.241.187.158:33546\x00")
    connTCP.Write(message)

}

func server() {

	tcpLocalAddr, err := net.ResolveTCPAddr("tcp","129.241.187.158:33546")
    tcpListener, err := net.ListenTCP("tcp", tcpLocalAddr)
    if err != nil {
        Println("Error listening:", err.Error())
    }
    defer tcpListener.Close()
    for {
        // Listen for an incoming connection.
        conn, err := tcpListener.AcceptTCP()
        Println("accepted connection")
        if err != nil {
            Println("Error accepting: ", err.Error())
        }

        go handleConnection(conn)

    }
}

func handleConnection(conn * net.TCPConn){
    data := make([]byte, 1024)
    message := make([]byte, 1024)
    data = []byte("hello\x00")
    for {

    	_, err := conn.Write(data)
    	Println("msg write\x00")
    	if err != nil {
    		Printf("Error in TCP: %s\n", err.Error())
    		break
    	}

    	time.Sleep(100*time.Millisecond)
    	conn.Read(message)
    	Println("msg read")
    	Printf("Received string: %s\n", message)
    }
}

func main() {
	doneChan := make(chan string)

	go server()
	time.Sleep(100*time.Millisecond)
	go client()

	Printf(<-doneChan)
}

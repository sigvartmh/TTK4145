package main

import( 
	"net"
	"fmt"
	"time"
)

const (
	servAddr = "129.241.187.255"
	udpPort = "20024"
	host = "129.241.187.104"
)

func udpSend(done chan bool){
	saddr, err := net.ResolveUDPAddr("udp4", net.JoinHostPort(servAddr, udpPort))
	
	if err != nil{
		fmt.Println("Failed to resolve address for: " + port)

	}

	conn, err := net.DialUDP("udp", nil, saddr)

	if err != nil {
		fmt.Println("Error connecting to" + servAddr + ":" + udpPort )
	}

	for {
		time.Sleep(1000*time.Millisecond)
		conn.Write([]byte("The cake is a lie"))
		fmt.Println("Msg sent on udp")
	}
	done <- true
}


func udpRecive(done chan bool, port string) {
	
	saddr, err := net.ResolveUDPAddr("udp4", net.JoinHostPort(servAddr, udpPort))

	if err != nil{
		fmt.Println("Failed to resolve address for: " + port)

	}

	buff := make([]byte, 1024)

	l, err := net.ListenUDP("udp4", saddr)
	if err != nil {
		fmt.Println("Error listening to" )//+ saddr)
	}

	for {
		bytes , remoteAddr, err := l.ReadFromUDP(buff)
		if err != nil {
			fmt.Println(err)
		}

	fmt.Println("Recived ", bytes, " bytes from: ", remoteAddr)
	fmt.Println(string(buff[:]))
	} 
	
	done<-true

}

func main() {
	done := make(chan bool)
	go udpSend(done)
	go udpRecive(done, udpPort)
	<-done

}

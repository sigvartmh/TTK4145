package main

import(
	"fmt"
	"os"
	"os/exec"
	"net"
	"strconv"
	"strings"
	"time"
)

var addr string = "localhost"
var bAddr string = "129.241.187.255"
var port string = ":20017"

func checkErr(err error){
	if err != nil {
		fmt.Println("An unrecovarable error occured", err.Error())
		os.Exit(0)
	}
}

func spawnProcess(){
	cmd := exec.Command("gnome-terminal", "-e", "./test.sh")
	out, err := cmd.Output()
	checkErr(err)
	fmt.Println(string(out))
}

func getCount(udpMsg string) int{
	n := strings.TrimLeft(udpMsg, "Count: ")
	count, err := strconv.Atoi(n)
	if(err != nil){
		return -1
	}
	return count
}

func slaveProcess(conn *net.UDPConn, masterAlive bool, count *int) bool{
	for(masterAlive){
		conn.SetReadDeadline(time.Now().Add(time.Second*2)) //takes some time to open terminal
		data := make([]byte, 16)
		length, _, err := conn.ReadFromUDP(data[0:])
		if err != nil {
			masterAlive = false
			return masterAlive
		} else{
			*count = getCount(string(data[0:length]))
			fmt.Println("Slave, master count:", *count)
		}
	}
	return true
}


func main(){

	masterAlive:= true
	count := 0
	startCount := 0
	udpAddr, err := net.ResolveUDPAddr("udp", port)
	checkErr(err)
	sConn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)

	masterAlive = slaveProcess(sConn, masterAlive, &count)
	sConn.Close()

	fmt.Println("Master Process")
	spawnProcess()
	startCount = count

	rbAddr, _ := net.ResolveUDPAddr("udp4",bAddr+port)
	mConn, err := net.DialUDP("udp4", nil, rbAddr)
	if err != nil{
			fmt.Println("Error:connecting: ",err.Error())
	}


	
	for {
		msg := "Count:" + strconv.Itoa(count)
		_, err := mConn.Write([]byte(msg))
		fmt.Println("Master count: ", count)
		if err != nil {
			fmt.Println("Error:Broadcast", err.Error())
		}

		if (count == 13 + startCount) {
			break
		}
		count++
		time.Sleep(time.Second)
	}
	mConn.Close()
	fmt.Println("Finished counting")

}
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	strEcho := "halo"
	servAddr := "localhost:8043"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		fmt.Println("Resolve failed")
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Dial failed")
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		fmt.Println("Write failed")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)

	if err != nil {
		fmt.Println("Write to server failed")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Reply from server = ", string(reply))

	conn.Close()
}

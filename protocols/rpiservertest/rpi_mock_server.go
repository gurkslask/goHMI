package main

import (
	"bufio"
	"encoding/hex"
	"log"
	"net"
)

func main() {
	log.Printf("Echo server listening on port 8043")
	ln, err := net.Listen("tcp", ":8043")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	accepted := 0
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		accepted++
		go handleConnection(conn)
		log.Printf("Connection accepted %d", accepted)

	}
}

func handleConnection(conn net.Conn) {
	bufr := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	for {
		readbytes, err := bufr.Read(buf)

		if err != nil {
			log.Fatalf("Error: %s", err)
			conn.Close()
		}
		log.Printf("<->\n%s", hex.Dump(buf[:readBytes]))
		conn.Write([]byte("CLOUDWALK " + string(buf[:readBytes])))
	}
}

package main

import (
	"log"
	"net"
)

func main() {
	runServer()
}

func runServer() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9555")
	if err != nil {
		log.Fatalln("resolve error : ", err)
	}
	conn, err1 := net.ListenUDP("udp", addr)
	if err1 != nil {
		log.Fatalln("listen error : ", err1)
	}
	for {
		msg := make([]byte, 1024)
		num, addr, err := conn.ReadFromUDP(msg)
		if err != nil {
			log.Fatalln("read error : ", err)
		}
		log.Println("read from client :", num, " bytes")
		log.Println(string(msg))
		writeNum, err := conn.WriteToUDP(msg, addr)
		if err != nil {
			log.Fatalln("write error : ", err)
		}
		log.Println("write ", writeNum, " bytes")
	}
}

package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:9555")
	reader := bufio.NewReader(os.Stdin)
	if err != nil {
		log.Fatalln("dial error : ", err)
	}
	for {
		log.Println("input msg to send")
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("read error : ", err)
		}
		conn.Write([]byte(str))
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		log.Println("read : ", string(buffer))
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Print("Hi. What are you name? ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	name := stdin.Text()
	fmt.Println("Hello", name)

	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalln("Failed listen:", err)
	}
	fmt.Println("Listen 127.0.0.1:8888")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln("Failed accept:", err)
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			log.Fatalln("Failed read:", err)
		}
		fmt.Printf("[Message]\n%s", string(buf[:n]))
		conn.Close()
	}
}

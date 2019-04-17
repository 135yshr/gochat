package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Print("Hi. What are you name? ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	name := stdin.Text()
	fmt.Println("Hello", name)

	listen, _ := net.Listen("tcp", "127.0.0.1:8888")
	fmt.Println("Listen 127.0.0.1:8888")

	conn, _ := listen.Accept()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Printf("[Message]\n%s", string(buf[:n]))
	conn.Close()
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Hi. What are you name? ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	name := stdin.Text()
	fmt.Println("Hello", name)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello from the containerized Go application!")
		time.Sleep(5 * time.Second)
	}
} 
package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("This is my new version of my Go application (NEW)!")
		time.Sleep(3 * time.Second)
	}
} 
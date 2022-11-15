package main

import (
	"fmt"
)

func receiver(ch chan int) {
	fmt.Println(<-ch)
}
func main() {
	fmt.Println("Channels")
	ch := make(chan int)
	go receiver(ch)
	ch <- 48
}

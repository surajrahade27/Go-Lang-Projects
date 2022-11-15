package main

import (
	"fmt"
	"time"
)

func f1(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println("Suraj is", s, ": ", i)
		time.Sleep(time.Microsecond * time.Duration(2))
	}
}
func main() {

	go f1("Coding")
	f1("Using Mobile")
	f1("reffering Tutorials")
}

package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "NNN"
	ch <- "SSS"
	close(ch)
	v1, ok := <-ch
	fmt.Println(v1, ok)
	v2, ok := <-ch
	fmt.Println(v2, ok)
	v3, ok := <-ch
	fmt.Println(v3, ok)

	val, ok := "String", true
	fmt.Print(val, ok)

	// for i := range ch {
	// 	fmt.Println(i)
	// }

}

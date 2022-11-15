package main

import (  
    "fmt"
)

func main() {  
    ch := make(chan int, 2)
    ch <- 5
    ch <- 6
    close(ch)
    for n := range ch {
        fmt.Println("Received:", n)
    }
}
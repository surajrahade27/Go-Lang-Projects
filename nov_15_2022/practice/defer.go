package main

import "fmt"

func main() {
	fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("three")
	defer fmt.Println("four")
	fmt.Println("five")
}

/**
 DEFER statements delay the execution of the function or statement
       until the nearby functions executes its statements.

Note:
In Go language,
1.multiple defer statements are allowed in the same program
and they are executed in LIFO(Last-In, First-Out) order.

2.Defer statements are generally used to ensure that
the files are closed when their need is over,
 or to close the channel,
 or to catch the panics in the program.
**/

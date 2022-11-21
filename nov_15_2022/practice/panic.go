package main

import "fmt"

func main() {
	emp := make(map[string]int)
	emp["Samia"] = 20
	emp["Sana"] = 23
	//Implicite Panic raised by program
	fmt.Println(emp[20])
	//fmt.Println(emp["Sana"])
}

/**
 PANIC is similar to exceptions raised at runtime when an error is encountered.

  panic() is either raised by the program itself
  when an unexpected error occurs or the programmer throws the exception
  on purpose for handling particular errors.
**/

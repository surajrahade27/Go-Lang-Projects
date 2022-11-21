// Recover() is used to handle panic.
package main

import "fmt"

func panicHandler() {
	rec := recover()
	//initially rec:nil
	//but when panic arrives "rec" stores that panicData
	if rec != nil {
		fmt.Println("RECOVER :", rec)
	}
}

func employee(name *string, age int) {

	defer panicHandler()

	if age > 65 {
		//will raise panic and

		panic("Age cannot be greater than retirement age")
		// it will capture in recover
	}

}

func main() {
	empName := "Samia"
	age := 75

	employee(&empName, age)

}

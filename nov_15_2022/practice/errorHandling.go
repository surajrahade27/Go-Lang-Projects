package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(value float64) (float64, error) {

	//-1<0 = true
	// 9<0  = false
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
func main() {
	//will call fun with -1 value
	result, err := Sqrt(-1)

	if err != nil {
		//"Math: negative number passed to Sqrt"
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	//err=nil
	if err != nil {

		fmt.Println(err)
	} else {
		//executes
		fmt.Println(result)
	}
}

/** handling errors in Go is , to compare the returned/captured error to nil .
 A nil value indicates that no error has occurred and
  a non-nil value indicates the presence of an error.
**/

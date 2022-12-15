package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var myArray []string

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Error err: %v", err)
	}
	fmt.Fprintf(w, "POST Request SuccessFull.\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Add = %s\n", address)

}
func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "HELLOWWW folks!")
} //end
func PushHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Error err: %v", err)
	}
	fmt.Fprintf(w, "PUSH Request SuccessFull.\n\n\n\n")
	data := r.FormValue("data")
	fmt.Fprintf(w, "Before myArray:= %s  \n\n\n", myArray)
	fmt.Fprintf(w, "Incoming Data = %s\n\n\n", data)
	myArray = append(myArray, data)
	fmt.Fprintf(w, "After myArray:= %s  \n\n\n", myArray)
} //end
func PopHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Error err: %v", err)
	}
	fmt.Fprintf(w, "POP Request SuccessFull.\n\n\n\n")
	fmt.Fprintf(w, "Before Pop,\n myArray:= %s  \n\n\n", myArray)
	myArray = myArray[:len(myArray)-1]
	fmt.Fprintf(w, "After POP, \n myArray:= %s  \n\n\n", myArray)
} //end
func lowestHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Error err: %v", err)
	}
	data := r.FormValue("data")
	myArray = strings.Split(data, ",")
	first := myArray[0]
	smallest, err := strconv.Atoi(first)
	if err != nil {
		fmt.Fprintf(w, "Error err: %v", err)
	}
	for _, num := range myArray {
		num1, err := strconv.Atoi(num) // iterate over the rest of the list
		if err != nil {
			fmt.Fprintf(w, "Error err: %v", err)
		}
		if num1 < smallest { // if num is smaller than the current smallest number
			smallest = num1 // set smallest to num
		}
	}
	fmt.Fprintf(w, "Lowest Element Found SuccessFull.\n\n\n\n")
	fmt.Fprintf(w, "Lowest Element:=%v  \n\n\n", smallest)
} //end
func main() {

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", FormHandler)
	http.HandleFunc("/push", PushHandler)
	http.HandleFunc("/pop", PopHandler)
	http.HandleFunc("/findLowest", lowestHandler)
	http.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"ASSIGNMENT/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	WEBPORT = ":8080"
)

func main() {
	fmt.Println("My App Started")
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/home", homeFunc)
	router.HandleFunc("/createuser", handlers.CreateUser)
	router.HandleFunc("/getUser", handlers.GetUsers)
	router.HandleFunc("/getAllUser", handlers.GetAllUsers)
	router.HandleFunc("/UpdateUser", handlers.UpdateUsers)
	router.HandleFunc("/DeleteUser", handlers.DeleteUsers)
	router.HandleFunc("/login", handlers.Login)
	http.ListenAndServe(WEBPORT, nil)
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "We have Received Request")

}
func Greet() string {
	return "Hello GitHub Actions"
}

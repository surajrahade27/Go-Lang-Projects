package handlers

import (
	"ASSIGNMENT/datastore"
	"ASSIGNMENT/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	u := model.User{}
	json.Unmarshal(reqbody, &u)
	fmt.Println(u)
	res, err := datastore.InsertDB(&u)
	if err != err {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	u := model.User{}
	json.Unmarshal(reqbody, &u)
	fmt.Println(u)
	user, err := datastore.GetDBUsers(&u)
	if err != err {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	u := model.User{}
	json.Unmarshal(reqbody, &u)
	fmt.Println(u)
	user, err := datastore.GetAllDBUsers(&u)
	if err != err {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	u := model.User{}
	json.Unmarshal(reqbody, &u)
	fmt.Println(u)

	_, err = datastore.GetDBUsers(&u)
	if err == sql.ErrNoRows {
		err = datastore.UpdateDB(&u)
		if err != err {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.WriteHeader(http.StatusOK)
		json.Marshal(w)
	} else {
		fmt.Println("No Row exist")
		w.Header().Set("Content-Type", "Application/json")
		w.WriteHeader(http.StatusBadGateway)
		json.Marshal(w)
	}
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	u := model.User{}
	json.Unmarshal(reqbody, &u)
	fmt.Println(u)
	err = datastore.DeleteDBUsers(&u)
	if err != err {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.Marshal(w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	u := model.User{}
	json.Unmarshal(reqbody, &u)
	fmt.Println(u)
	resp, err := datastore.VerifyUser(&u)
	if err != err {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

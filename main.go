package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

const VERSION = "0.0.1"

type Person struct {
	ID		string		`json:"id,omitempty"`
	Firstname	string		`json:"firstname,omitempty"`
	Lastname	string		`json:"lastname,omitempty"`
	Address		*Address	`json:"address,omitempty"`
}

type Address struct {
	Street		string		`json:"street,omitempty"`
	City		string		`json:"city,omitempty"`
	State		string		`json:"state,omitempty"`
}

var people []Person

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/", Index).Methods("GET")

  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people", CreatePerson).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Version %s", VERSION)
}

func GetPeople(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode(people)
}

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}

	json.NewEncoder(writer).Encode(&Person{})
}

func CreatePerson(writer http.ResponseWriter, request *http.Request) {
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	person.ID = uuid.New().String()

	people = append(people, person)
	json.NewEncoder(writer).Encode(people)
}

func DeletePerson(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(writer).Encode(people)
	}
}



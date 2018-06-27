package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const VERSION = "0.0.2"
const PORT = ":8000"
const DB_CREATE_TABLE_QUERY = "CREATE TABLE IF NOT EXISTS people (id STRING PRIMARY KEY, firstname VARCHAR, lastname VARCHAR, streetaddress VARCHAR, city VARCHAR, state VARCHAR)"
const DB_SELECT_ALL_QUERY = "SELECT * FROM people"
const ID_KEY = "id"

const SQL_DRIVER = "sqlite3"
const DB_PATH = "./people.db"

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

func main() {
  // Initialize the database
  db, err := sql.Open(SQL_DRIVER, DB_PATH)
  checkDbErr(err)
  statement, err := db.Prepare(DB_CREATE_TABLE_QUERY)
  checkDbErr(err)
  _, err = statement.Exec()
  checkDbErr(err)
  db.Close()

  // Setup Routes
  router := mux.NewRouter()

  router.HandleFunc("/", Index).Methods("GET")

  router.HandleFunc("/people", GetPeople).Methods("GET")
  router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
  router.HandleFunc("/people", CreatePerson).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

  fmt.Println("Listening on " + PORT)

  log.Fatal(http.ListenAndServe(PORT, router))
}

func checkDbErr(err error) {
  if err != nil {
    panic(err)
  }
}

func Index(writer http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(writer, "Version %s", VERSION)
}

func GetPeople(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open(SQL_DRIVER, DB_PATH)
	checkDbErr(err)

	rows, err := db.Query(DB_SELECT_ALL_QUERY)
	checkDbErr(err)

	var people []Person
	for rows.Next() {
		var person Person
		var address Address
		err = rows.Scan(&person.ID, &person.Firstname, &person.Lastname, &address.Street, &address.City, &address.State)
		checkDbErr(err)

		person.Address = &address
		people = append(people, person)
	}

	rows.Close()
	db.Close()

	json.NewEncoder(writer).Encode(people)
}

func GetPerson(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	db, err := sql.Open(SQL_DRIVER, DB_PATH)
	checkDbErr(err)

	// TODO: this is pretty messy
	rows, err := db.Query("SELECT * FROM people WHERE ID='" + params[ID_KEY] + "' LIMIT 1")
	checkDbErr(err)

	var person Person
	var address Address
	for rows.Next() {
		err = rows.Scan(&person.ID, &person.Firstname, &person.Lastname, &address.Street, &address.City, &address.State)
		checkDbErr(err)
		person.Address = &address
	}

	rows.Close()
	db.Close()

	json.NewEncoder(writer).Encode(person)
}

func CreatePerson(writer http.ResponseWriter, request *http.Request) {
	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)
	person.ID = uuid.New().String()

	// got a people. Insert into DB
	db, err := sql.Open(SQL_DRIVER, DB_PATH)
	checkDbErr(err)

	statement, err := db.Prepare("INSERT INTO people(id, firstname, lastname, streetaddress, city, state) VALUES (?, ?, ?, ?, ?, ?)")
	checkDbErr(err)

	result, err := statement.Exec(person.ID, person.Firstname, person.Lastname, person.Address.Street, person.Address.City, person.Address.State)
	checkDbErr(err)

	_, err = result.LastInsertId()
	checkDbErr(err)

	db.Close()

	json.NewEncoder(writer).Encode(person)
}

func DeletePerson(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	db, err := sql.Open(SQL_DRIVER, DB_PATH)
	checkDbErr(err)

	statement, err := db.Prepare("DELETE FROM people WHERE ID=?")
	checkDbErr(err)

	_, err = statement.Exec(params[ID_KEY])
	checkDbErr(err)

	db.Close()

	// Fake Response
	writer.WriteHeader(http.StatusNoContent)
	json.NewEncoder(writer).Encode(&Person{})
}

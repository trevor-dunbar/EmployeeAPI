package main

import (
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error


func main() {

	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/employees")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("successfully connected to MySql")

	handleRequests()
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/employees", getEmployees).Methods("GET")
	myRouter.HandleFunc("/employees", createEmployee).Methods("POST")
	myRouter.HandleFunc("/employees/{name}", getEmployee).Methods("GET")
	myRouter.HandleFunc("/employees/{name}", updateEmployee).Methods("PUT")
	myRouter.HandleFunc("/employees/{name}", deleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

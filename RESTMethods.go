package main

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

func getEmployees(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query("SELECT * from employees")

	var employees []Employee
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

  	for result.Next() {
    	var employee Employee
    	err := result.Scan(&employee.Name, &employee.Email, &employee.Role, &employee.Skills)
    	if err != nil {
      		panic(err.Error())
    	}
    employees = append(employees, employee)
  }
  json.NewEncoder(w).Encode(employees)
};

func createEmployee(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO employees(name, email, role, skills) VALUES(?, ?, ?, ?)")
	if err != nil {
	  panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	email := keyVal["email"]
	role := keyVal["role"]
	skills := keyVal["skills"]


	_, err = stmt.Exec(name, email, role, skills)

	if err != nil {
	  panic(err.Error())
	}
	
	fmt.Fprintf(w, "New employee was created")
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
  	params := mux.Vars(r)
	result, err := db.Query("SELECT * FROM employees WHERE name = ?", params["name"])
	  
  	if err != nil {
    	panic(err.Error())
  	}
	  
	defer result.Close()
  	var employee Employee
  	for result.Next() {
    	err := result.Scan(&employee.Name, &employee.Email, &employee.Role, &employee.Skills)
    	if err != nil {
      	panic(err.Error())
    	}
  	}
  json.NewEncoder(w).Encode(employee)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	statement, err := db.Prepare("UPDATE employees SET email = ?, role = ?, skills = ?  WHERE name = ?")
	if err != nil {
	  panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
	  panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newEmail := keyVal["email"]
	newRole := keyVal["role"]
	newSkills := keyVal["skills"]
	_, err = statement.Exec(newEmail, newRole, newSkills, params["name"])

	if err != nil {
	  panic(err.Error())
	}

	fmt.Fprintf(w, "Post with name = %s was updated", params["name"])
  }

  func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	statement, err := db.Prepare("DELETE FROM employees WHERE name = ?")

	if err != nil {
	  panic(err.Error())
	}
	_, err = statement.Exec(params["name"])

	if err != nil {
	  panic(err.Error())
	}
	
	fmt.Fprintf(w, "employee with name = %s was deleted", params["name"])
  }
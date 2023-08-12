package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Employee Struct Model

type Employee struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
	Address *Address `json:"address"`
}

type Address struct {
	Street  string `json:"streetName"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// Initalize Employees Mock
var employeeList []Employee

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(employeeList)
}

func main() {
	// Init router
	r := mux.NewRouter()

	employeeList = append(employeeList, Employee{
		Id:      1,
		Name:    "Madhukshra",
		Phone:   "89898989",
		Address: &Address{Street: "Seetharam", City: "Mannargudi", Country: "India"}})

	// Routes for login
	r.HandleFunc("/api/list", getEmployees).Methods("GET")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

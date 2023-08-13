package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Employee Struct Model

type Employee struct {
	Id      string   `json:"id"`
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

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r) //Get Params from Get method
	for _, item := range employeeList {
		log.Println("item ", item)
		log.Println("params ", params)
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var list Employee
	_ = json.NewDecoder(r.Body).Decode(&list) //Get Params from Get method

	list.Id = strconv.Itoa(rand.Intn(10000000)) //Assigning dummy id to new record
	employeeList = append(employeeList, list)

	json.NewEncoder(w).Encode(list)
}

func main() {
	// Init router
	r := mux.NewRouter()

	employeeList = append(employeeList, Employee{
		Id:      "1",
		Name:    "Madhukshra",
		Phone:   "89898989",
		Address: &Address{Street: "Seetharam", City: "Mannargudi", Country: "India"}})
	employeeList = append(employeeList, Employee{
		Id:      "2",
		Name:    "Gokulan",
		Phone:   "89898989",
		Address: &Address{Street: "Ullikkottai", City: "Mannargudi", Country: "India"}})
	employeeList = append(employeeList, Employee{
		Id:      "3",
		Name:    "Palani",
		Phone:   "89898989",
		Address: &Address{Street: "Ullikkottai", City: "Mannargudi", Country: "India"}})
	employeeList = append(employeeList, Employee{
		Id:      "4",
		Name:    "Jaya Prakash",
		Phone:   "89898989",
		Address: &Address{Street: "Ullikkottai", City: "Mannargudi", Country: "India"}})
	employeeList = append(employeeList, Employee{
		Id:      "5",
		Name:    "Ashok",
		Phone:   "89898989",
		Address: &Address{Street: "Mannargudi", City: "Mannargudi", Country: "India"}})
	employeeList = append(employeeList, Employee{
		Id:      "6",
		Name:    "Sundar",
		Phone:   "89898989",
		Address: &Address{Street: "Mannargudi", City: "Mannargudi", Country: "India"}})

	// Routes for login
	r.HandleFunc("/api/list", getEmployees).Methods("GET")
	r.HandleFunc("/api/employee/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/api/employee", createEmployee).Methods("POST")

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}

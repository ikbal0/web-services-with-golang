package main

import (
	// "encoding/json"
	// "fmt"
	// "net/http"
	// "strconv"
	"web-services/routers"
)

// type Employee struct {
// 	ID       int
// 	Name     string
// 	Age      int
// 	Division string
// }

var PORT = ":8080"

// var employees = []Employee{
// 	{ID: 1, Name: "Ariel", Age: 22, Division: "Finance"},
// 	{ID: 2, Name: "Amanda", Age: 24, Division: "IT Staff"},
// 	{ID: 3, Name: "Mario", Age: 23, Division: "Software Engineer"},
// }

func main() {
	routers.StartServer().Run(PORT)
	// http.HandleFunc("/", greet)
	// http.HandleFunc("/employees", getEmployees)
	// http.HandleFunc("/employee", setEmployees)

	// fmt.Println("Application listening to port", PORT)

	// http.ListenAndServe(PORT, nil)
}

// func setEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "POST" {
// 		name := r.FormValue("name")
// 		age := r.FormValue("age")
// 		division := r.FormValue("division")

// 		convertAge, err := strconv.Atoi(age)

// 		if err != nil {
// 			http.Error(w, "Invalid age", http.StatusBadRequest)
// 			return
// 		}

// 		newEmployee := Employee{
// 			ID:       len(employees) + 1,
// 			Name:     name,
// 			Age:      convertAge,
// 			Division: division,
// 		}

// 		employees = append(employees, newEmployee)

// 		json.NewEncoder(w).Encode(newEmployee)
// 		return
// 	}

// 	http.Error(w, "Invalid method", http.StatusBadRequest)
// }

// func getEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "GET" {
// 		json.NewEncoder(w).Encode(employees)
// 		return
// 	}

// 	http.Error(w, "Invalid method", http.StatusBadRequest)
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	msg := "Hello World"
// 	fmt.Fprint(w, msg)
// }

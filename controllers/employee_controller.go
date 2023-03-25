package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Division   string `json:"division"`
}

var employees = []Employee{}

func SetEmployees(ctx *gin.Context) {
	// w.Header().Set("Content-Type", "application/json")

	// if r.Method == "POST" {
	// 	name := r.FormValue("name")
	// 	age := r.FormValue("age")
	// 	division := r.FormValue("division")

	// 	convertAge, err := strconv.Atoi(age)

	// 	if err != nil {
	// 		http.Error(w, "Invalid age", http.StatusBadRequest)
	// 		return
	// 	}

	// 	newEmployee := Employee{
	// 		ID:       len(employees) + 1,
	// 		Name:     name,
	// 		Age:      convertAge,
	// 		Division: division,
	// 	}

	// 	employees = append(employees, newEmployee)

	// 	json.NewEncoder(w).Encode(newEmployee
	// 	return
	// }

	// http.Error(w, "Invalid method", http.StatusBadRequest)

	var newEmployee Employee

	if err := ctx.ShouldBindJSON(&newEmployee); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newEmployee.EmployeeID = fmt.Sprintf("e%d", len(employees)+1)
	employees = append(employees, newEmployee)

	ctx.JSON(http.StatusCreated, gin.H{
		"employees": employees,
	})
}

func GetEmployees(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, employees)
	// 	w.Header().Set("Content-Type", "application/json")

	// 	if r.Method == "GET" {
	// 		json.NewEncoder(w).Encode(employees)
	// 		return
	// 	}

	// 	http.Error(w, "Invalid method", http.StatusBadRequest)
}

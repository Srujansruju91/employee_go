package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Asus@123"
	dbname   = "postgres"
)

var db *sql.DB

func main() {
	// Connect to the database
	db = connectDB()
	defer db.Close()
	http.HandleFunc("/insert", insertHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	return db
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Extract form values
	firstName := r.Form.Get("first_name")
	lastName := r.Form.Get("last_name")
	// Parse other form fields similarly

	// Insert into database
	_, err = db.Exec("INSERT INTO employees (first_name, last_name) VALUES ($1, $2)", firstName, lastName)
	if err != nil {
		http.Error(w, "Failed to insert employee details", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	fmt.Fprintf(w, "Employee details inserted successfully")
}



// // Implement functions for each option (insertEmployee, updateEmployee, etc.)

// func insertEmployee() {
// 	// Implement function to insert employee details
// }

// func updateEmployee() {
// 	// Implement function to update employee details
// }

// func displayAllEmployees() {
// 	// Implement function to display all employees
// }

// func deleteEmployee() {
// 	// Implement function to delete an employee
// }

// func insertPosition() {
// 	// Implement function to insert a new position
// }

// func updatePosition() {
// 	// Implement function to update a position
// }

// func displayAllPositions() {
// 	// Implement function to display all positions
// }

// func deletePosition() {
// 	// Implement function to delete a position
// }

// func insertPayscale() {
// 	// Implement function to insert a new payscale
// }

// func updatePayscale() {
// 	// Implement function to update a payscale
// }

// func displayAllPayscales() {
// 	// Implement function to display all payscales
// }

// func deletePayscale() {
// 	// Implement function to delete a payscale
// }

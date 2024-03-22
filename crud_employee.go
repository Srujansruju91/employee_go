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

	var choice int
	for {
		fmt.Println("1. Insert employee details")
		fmt.Println("2. Update employee details")
		fmt.Println("3. Display all employees")
		fmt.Println("4. Delete employee")
		fmt.Println("5. Insert new position")
		fmt.Println("6. Update position")
		fmt.Println("7. Display all positions")
		fmt.Println("8. Delete position")
		fmt.Println("9. Insert new payscale")
		fmt.Println("10. Update payscale")
		fmt.Println("11. Display all payscales")
		fmt.Println("12. Delete payscale")
		fmt.Println("13. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			insertEmployee()
		case 2:
			updateEmployee()
		case 3:
			displayAllEmployees()
		case 4:
			deleteEmployee()
		case 5:
			insertPosition()
		case 6:
			updatePosition()
		case 7:
			displayAllPositions()
		case 8:
			deletePosition()
		case 9:
			insertPayscale()
		case 10:
			updatePayscale()
		case 11:
			displayAllPayscales()
		case 12:
			deletePayscale()
		case 13:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func connectDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	return db
}

// Implement functions for each option (insertEmployee, updateEmployee, etc.)

func insertEmployee() {
	// Implement function to insert employee details
}

func updateEmployee() {
	// Implement function to update employee details
}

func displayAllEmployees() {
	// Implement function to display all employees
}

func deleteEmployee() {
	// Implement function to delete an employee
}

func insertPosition() {
	// Implement function to insert a new position
}

func updatePosition() {
	// Implement function to update a position
}

func displayAllPositions() {
	// Implement function to display all positions
}

func deletePosition() {
	// Implement function to delete a position
}

func insertPayscale() {
	// Implement function to insert a new payscale
}

func updatePayscale() {
	// Implement function to update a payscale
}

func displayAllPayscales() {
	// Implement function to display all payscales
}

func deletePayscale() {
	// Implement function to delete a payscale
}

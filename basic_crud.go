package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

// Define a struct to represent an entity
type Item struct {
    ID   int
    Name string
}

var db *sql.DB

func connectDB() error {
    var err error
    // Replace "user", "password", "dbname", and "host" with your PostgreSQL details
    db, err = sql.Open("postgres", "postgres://postgres:Asus@123@localhost/postgres?sslmode=disable")
    if err != nil {
        return err
    }
    return nil
}

func closeDB() {
    db.Close()
}

// Create operation
func createItem(name string) error {
    _, err := db.Exec("INSERT INTO items (name) VALUES ($1)", name)
    return err
}

// Read operation
func getItems() ([]Item, error) {
    rows, err := db.Query("SELECT id, name FROM items")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var items []Item
    for rows.Next() {
        var item Item
        err := rows.Scan(&item.ID, &item.Name)
        if err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    return items, nil
}

// Update operation
func updateItem(id int, newName string) error {
    _, err := db.Exec("UPDATE items SET name = $1 WHERE id = $2", newName, id)
    return err
}

// Delete operation
func deleteItem(id int) error {
    _, err := db.Exec("DELETE FROM items WHERE id = $1", id)
    return err
}

func main() {
    err := connectDB()
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer closeDB()

    // Example usage
    err = createItem("Apple")
    if err != nil {
        fmt.Println("Error creating item:", err)
        return
    }

    items, err := getItems()
    if err != nil {
        fmt.Println("Error fetching items:", err)
        return
    }
    fmt.Println("Items:", items)

    err = updateItem(1, "Grapes")
    if err != nil {
        fmt.Println("Error updating item:", err)
        return
    }

    err = deleteItem(3)
    if err != nil {
        fmt.Println("Error deleting item:", err)
        return
    }
}


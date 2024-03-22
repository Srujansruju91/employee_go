package main

import (
    "log"
    "net/http"
)

func main() {
    // Serve the static files from the "static" directory
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)

    // Start the server on port 8080
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

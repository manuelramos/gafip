// main.go
package main

import (
	"log"

	"github.com/manuelramos/gafip/internal/db"
)

func main() {
	dbConnector := db.NewDBConnector()

    err := dbConnector.Connect()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    defer dbConnector.Close()
    log.Println("Connected to database!")
}

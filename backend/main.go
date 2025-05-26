package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "net/http"
    "fmt"
)

var db *gorm.DB

func initDB() {
    dsn := "host=localhost user=postgres password=damar123 dbname=topup_store port=5432 sslmode=disable"
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    fmt.Println("✅ Connected to PostgreSQL")
}

func main() {
    initDB()

    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("pong"))
    })

    fmt.Println("🚀 Backend running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

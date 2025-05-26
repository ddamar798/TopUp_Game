package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

type Game struct {
    gorm.Model
    Name        string
    Description string
    Price       int
    ImageURL    string
}

func initDB() {
    dsn := "host=localhost user=postgres password=damar123 dbname=topup_store port=5432 sslmode=disable"
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Gagal koneksi database:", err)
    }
    fmt.Println("✅ Connected to PostgreSQL")

    // Auto-create tabel
    db.AutoMigrate(&Game{})
    fmt.Println("✅ Migrated database")
}

func getGames(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var games []Game
    result := db.Find(&games)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(games)
}

func main() {
    initDB()

    http.HandleFunc("/games", getGames)

    fmt.Println("🚀 Backend running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

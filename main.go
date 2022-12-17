package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var dsn = "root:password@tcp(localhost:3306)/git?charset=utf8mb4"
var dbConnection, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type DbFields struct {
	ID   int
	Name string
	Year string
}

func main() {
	http.HandleFunc("/", DatabaseCreate)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DatabaseCreate(w http.ResponseWriter, r *http.Request) {
	dbFields := DbFields{
		ID:   3,
		Name: "Geoff",
		Year: "1965",
	}

	if err := dbConnection.Create(&dbFields).Error; err != nil {
		log.Fatal(err)
	}

	if err := json.NewEncoder(w).Encode(dbFields); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fields added", dbFields)

}

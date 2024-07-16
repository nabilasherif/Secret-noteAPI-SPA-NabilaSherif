package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type note struct {
	noteText       string    `json:"note_text"`
	expirationDate time.Time `json:"expiration_date"`
	maxViewers     int       `json:"max_viewers"`
	currentViewers int       `json:"current_viewers"`
	url            string    `json:"note_url"`
}

type User struct {
	gorm.Model
	name      string `gorm:"column:name"`
	password  int    `gorm:"column:password"`
	userNotes []note `gorm:"column:user notes"`
}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func AddNote(w http.ResponseWriter, r *http.Request) {

}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {

	http.HandleFunc("/", homePageHandler)
	log.Panic(
		http.ListenAndServe(":8080", nil),
	)
}

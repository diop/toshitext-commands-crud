package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Command struct {
	gorm.Model
	Name         string
	Instructions string
}

func InitialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to database")

	}
	defer db.Close()

	db.AutoMigrate(&Command{})
}

func AllCommands(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the SQL database")
	}
	defer db.Close()
	var commands []Command
	db.Find(&commands)
	json.NewEncoder(w).Encode(commands)
}

func NewCommand(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the SQL database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	instructions := vars["instructions"]

	db.Create(&Command{Name: name, Instructions: instructions})

	fmt.Fprintf(w, "New command Succesfully Created!")
}

func DeleteCommand(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the SQL database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var command Command
	db.Where("name = ?", name).Find(&command)
	db.Delete(&command)
	fmt.Fprintf(w, "New command Successfully Deleted!")
}

func UpdateCommand(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Could not connect to the SQL database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	instructions := vars["instructions"]

	var command Command
	db.Where("name = ?", name).Find(&command)

	command.Instructions = instructions

	db.Save(&command)
	fmt.Fprintf(w, "Successfully Updated Command")
}

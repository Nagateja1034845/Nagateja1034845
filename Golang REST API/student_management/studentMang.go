package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Mysql172@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Grade string `json:"grade"`
	Std   int    `json:"std"`
	Batch string `json:"batch"`
	//Class *Details `json:"class"`
}

/*type Details struct {
	Std     string `json:"std"`
	Section string `json:"section"`
}*/

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/student", GetStudents).Methods("GET")
	r.HandleFunc("/student/{id}", GetStudent).Methods("GET")
	r.HandleFunc("/student", CreateStudent).Methods("POST")
	r.HandleFunc("/student/{id}", UpdateStudent).Methods("PUT")
	r.HandleFunc("/student/{id}", DeleteStudent).Methods("DELETE")
	//r.HandleFunc("/student/{name}", GetGroupBy).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Student{})

}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students []Student
	DB.Find(&students)
	json.NewEncoder(w).Encode(students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var students Student
	DB.First(&students, params["id"])
	json.NewEncoder(w).Encode(students)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students Student
	json.NewDecoder(r.Body).Decode(&students)
	DB.Create(&students)
	json.NewEncoder(w).Encode(students)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var students Student
	DB.First(&students, params["id"])
	json.NewDecoder(r.Body).Decode(&students)
	DB.Save(&students)
	json.NewEncoder(w).Encode(students)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var students Student
	DB.Delete(&students, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}

/*func GetGroupBy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//	params := mux.Vars(r)
	var students []Student
	DB.Group("name")
	json.NewEncoder(w).Encode(students)
	//func (db *DB) Group(name string) (tx *DB)
}*/

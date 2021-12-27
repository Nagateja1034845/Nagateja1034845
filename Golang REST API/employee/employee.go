package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// The employee Type (more like an object)
type employee struct {
	empId   int    `json:"empId"`
	empName string `json:"empName"`
	empAge  int    `json:"empAge"`
	empDept string `json:"empDept"`
}

//var emp []employee

// function to insert a row in employee table
func insert(emp employee) {

	db := connect()

	stmt, err := db.Query("INSERT INTO employelist (empId, empName, empAge, empDept) VALUES(?,?,?,?)", emp.empId, emp.empName, emp.empAge, emp.empDept)

	if err != nil {
		log.Fatal("Unable to prepare statement:", err)
	}

	defer stmt.Close()
	defer db.Close()

}

// function to get a database connection
func connect() *sql.DB {

	db, err := sql.Open("mysql", "root:Mysql172@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		log.Fatal("Unable to open connection to db")
	}

	return db
}

// function to select all records from employee table
func selectAll() *sql.Rows {
	db := connect()

	results, err := db.Query("SELECT * FROM employelist")

	if err != nil {
		fmt.Println("Error! Getting records...")
	}

	defer db.Close()
	return results
}

// function to select a employee record from table by employee id
func selectById(id int) *sql.Row {
	db := connect()
	result := db.QueryRow("SELECT * FROM employelist WHERE id=?", id)

	defer db.Close()
	return result
}

// function to update a employee record by employee id
func updateById(emp employee) {
	db := connect()

	db.QueryRow("UPDATE employelist SET empName=?, empAge=?, empDept=? WHERE empId=?", emp.empName, emp.empAge, emp.empDept, emp.empId)
}

// function to delete a employee by employee id
func delete(id int) {
	db := connect()
	db.QueryRow("DELETE FROM employelist WHERE empId=?", id)
}

func main() {

	// inserting a rows

	insert(employee{10002, "naga", 206, "golang"})
	insert(employee{10003, "utkarsha", 25, "hp"})

	// updating the employee by id
	updateById(employee{10002, "utkarsha", 25, "hp"})

	// select all employees
	results := selectAll()

	// iterating a results
	for results.Next() {
		var emp employee
		results.Scan(&emp.empId, &emp.empName, &emp.empAge, &emp.empDept)
		fmt.Println(emp.empId, emp.empName, emp.empAge, emp.empDept)
	}

	// select employee by id
	result := selectById(10001)

	var emp employee
	result.Scan(&emp.empId, &emp.empName, &emp.empAge, &emp.empDept)
	fmt.Println(emp.empId, emp.empName, emp.empAge, emp.empDept)

	// delete a employee by id
	delete(10001)
}

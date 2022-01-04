package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	//fmt.Println("Drivers:", sql.Drivers())
	db, err := sql.Open("mysql", "root:Mysql172@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		log.Fatal("Unable to open connection to db")
	}
	defer db.Close()

	results, err := db.Query("select * from employeer")

	if err != nil {
		log.Fatal("Error when fetching employeer table rows:", err)
	}
	for results.Next() {
		var (
			empId   int
			empName string
			empAge  int
			empDept string
		)
		err = results.Scan(&empId, &empName, &empAge, &empDept)
		if err != nil {
			log.Fatal("Unable to parse row:", err)
		}
		fmt.Printf("empId:%d, empName:%s,empAge:%d,empDept:%s\n", empId, empName, empAge, empDept)
	}
	employee := []struct {
		empId   int
		empName string
		empAge  int
		empDept string
	}{
		{1001, "Naga", 27, "golang"},
		{1002, "Utkarsha", 23, "Hp"},
		{1003, "Avinasha", 28, "Quest"},
		{1004, "Divya", 25, "golbal"},
		{1005, "Pusha", 23, "hp"},
	}
	stmt, err := db.Prepare("INSERT INTO employeer (empId, empName, empAge, empDept) VALUES(?,?,?,?)")
	defer stmt.Close()

	if err != nil {
		log.Fatal("Unable to prepare statement:", err)
	}
	for _, emp := range employee {
		_, err = stmt.Exec(emp.empId, emp.empName, emp.empAge, emp.empDept)
		if err != nil {
			log.Fatal("Unable to execute statements:", err)
		}
	}
}

/*func main() {

/*	cfg := mysql.Config{
	User:   root,
	Passwd: Mysql172,
	Net:    "tcp",
	Addr:   "127.0.0.1:3306",
	DBName: "testdb",
}*/
//	fmt.Println("Go MySQL Tutorial")

// Open up our database connection.
// I've set up a database on my local machine using phpmyadmin.
// The database is called testDb
//db, err := sql.Open("mysql", "root:Mysql172@tcp(127.0.0.1:3306)/testdb")
//db, err := sql.Open("mysql", cfg.FormatDSN())
/*	if err != nil {
		log.Fatal(err)
	}

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("hi")
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// Execute the query
	/*results, err := db.Query("SELECT id, name FROM tags")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var tag Tag
	for results.Next() {
		//var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.Name)
	}

	// Execute the query
	err = db.QueryRow("SELECT id, name FROM tags where id = ?", 2).Scan(&tag.ID, &tag.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	log.Println(tag.ID)
	log.Println(tag.Name)
*/

//}

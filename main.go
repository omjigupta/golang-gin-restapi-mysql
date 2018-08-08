package main

import _ "github.com/go-sql-driver/mysql"
import "database/sql"
import "fmt"

func main() {
  fmt.Println("hello world")

  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sampledb")
  checkErr(err)

	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	checkErr(err)

	stmt, err := db.Prepare("CREATE TABLE flower (id int NOT NULL AUTO_INCREMENT, name varchar(40), category varchar(40), price decimal, photo varchar(256), instructions varchar(256), PRIMARY KEY (id));")
	checkErr(err)

	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Table successfully migrated....")
	}

}

func checkErr(err error) {
  if err != nil {
    fmt.Print(err.Error())
  }
}

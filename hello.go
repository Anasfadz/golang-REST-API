package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:<>@tcp(127.0.0.1:3306)/test")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    fmt.Println("Success!")

	var a = "Hello World!!!"
	
    fmt.Println(a +"\n" +a)
}
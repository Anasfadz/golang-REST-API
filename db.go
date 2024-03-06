package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Testtable struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_tutorial")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id, name FROM test_table")
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()

	var testTables []Testtable
	for results.Next() {
		var testTable Testtable
		err = results.Scan(&testTable.ID, &testTable.Name)
		if err != nil {
			panic(err.Error())
		}
		testTables = append(testTables, testTable)
	}

	jsonData, err := json.Marshal(testTables)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

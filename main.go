package main

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

func main(){
	dns := `user="system"
			password="oracle"
			connectString="localhost:1521/xe"
			libDir="/Users/root/Downloads/instantclient_19_8"
			`


	db, err := sql.Open("godror", dns)
	if err != nil{
		fmt.Print(err)
	}
	defer db.Close()
	db.Ping()

	sql := "select IntCol, StringCol from TestQuery where IntCol < 100"
	rows, _ := db.Query(sql)
	defer rows.Close()
	var intCol, strCol string
	for rows.Next() {
		rows.Scan(&intCol, &strCol)
		fmt.Printf("IntCol=%s, StrCol=%s\n", intCol, strCol)
	}
	err = rows.Err()




}
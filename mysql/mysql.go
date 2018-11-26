package main 


import (
	"log"
	"fmt"
    "database/sql"
    _"github.com/go-sql-driver/mysql"
)

func  main(){

	db,err :=sql.Open("mysql","root")
	if err !=nil{
		fmt.Println("Connect  Result:",err)
	}
	rows, err :=db.Query("SHOW database")
	if err !=nil{
		fmt.Println("Query  Result:",err)
	}
	for  rows.Next(){
		var  dataName  string 
		if err := rows.Scan(dataName); err !=nil{
			log.Fatal(err)
		}
		fmt.Printf("database  is %s  ",dataName)
	}
}
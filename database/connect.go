package database

import (
	"database/sql"
	 _ "github.com/lib/pq"
	 "log"
	// "os"
)


func DB_connection() (*sql.DB,error){
	connStr := "postgres://admin:7029@localhost:5433/test?sslmode=disable" 
        db,err:=sql.Open("postgres",connStr)
        if err != nil {
     		log.Fatal(err)
     	}
		return db,err
}
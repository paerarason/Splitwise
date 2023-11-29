package database

import (
	"database/sql"
	 _ "github.com/lib/pq"
	 "log"
	 "os"
)


func DB_connection() (*sql.DB,error){
	connStr := os.Getenv("database_url") 
        db,err:=sql.Open("postgres",connStr)
        if err != nil {
     		log.Fatal(err)
     	}
		return db,err
}
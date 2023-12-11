package database

import (
	"database/sql"
	 _ "github.com/lib/pq"
	 "log"
	 "os"
	 "github.com/joho/godotenv"
)


func DB_connection() (*sql.DB,error){
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
	connStr := os.Getenv("database_url")
        db,err:=sql.Open("postgres",connStr)
        if err != nil {
     		log.Fatal(err)
     	}
		
		return db,err
}
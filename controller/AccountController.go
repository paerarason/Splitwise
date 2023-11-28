package controller

import (
	"github.com/gin-gonic/gin"
    "net/http"
    //"encoding/json"
	"golang.org/x/crypto/bcrypt"
	 "database/sql"
	 _ "github.com/lib/pq"
    "log"
    "fmt"

)
type Account struct {
    Email string               `json:"email"`
    Username string            `json:"username"`
    Password string            `json:"password"`
    ConfirmPassword string     `json:"confirmPassword"`
}

func CreateAccount() gin.HandlerFunc {
    return func(c *gin.Context){
        connStr := "postgres://admin:7029@localhost/splitwise?sslmode=disable"     
        db,err:=sql.Open("postgres",connStr)
        if err != nil {
		log.Fatal(err)
     	}
        defer db.Close()
        
        var acc Account
        if err:=c.BindJSON(&acc);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        
        var accid int 
        dberr := db.QueryRow(`INSERT INTO account (username,password,email,balance)
	    VALUES(acc.Username,acc.Password,acc.Email,5000) RETURNING ID;`).Scan(&accid)
         fmt.Println(acc.Username)

        if dberr!=nil{ 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"message": ""}) 
    }
}

func GETspends() gin.HandlerFunc {
    return func(c *gin.Context){
     }
}


func CheckBalance() gin.HandlerFunc {
    return func(c *gin.Context){
     }
}


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

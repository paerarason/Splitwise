package controller

import (
    "github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	"golang.org/x/crypto/bcrypt"
	 _ "github.com/lib/pq"

)
type Account struct {
    Email string               `json:"email"`
    Username string            `json:"username"`
    Password string            `json:"password"`
    ConfirmPassword string     `json:"confirmPassword"`
}

func CreateAccount() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.DB_connection()
        
        //error Handling while making Connection 
        if err!=nil{
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }

        defer db.Close()        
        var acc Account

        //error Handling while Serialize the json from the request to the Account Struct 
        if err:=c.BindJSON(&acc);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        
        var accid int
        query := `INSERT INTO account (username, password, email, balance)
          VALUES ($1, $2, $3, $4) RETURNING id`
        
        dberr := db.QueryRow(query, acc.Username, acc.Password, acc.Email, 5000).Scan(&accid)
        
        //Handlinf while making Queries
        if dberr!=nil{ 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"account ID ":accid}) 
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

package controller

import (
    "github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	"golang.org/x/crypto/bcrypt"
	 _ "github.com/lib/pq"
     "log"
     "database/sql"

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


func GETspendAmount() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.DB_connection()        
        //error Handling while making Connection 
        if err!=nil{
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        user_id:=token()
        query := `SELECT COALESCE(SUM(transaction.amount), 0) AS total_amount
                   FROM account_Group
                   LEFT JOIN transaction  ON account_Group.ID = transaction.Account_Group_id
                   WHERE account_Group.ID = $1
        `
        var spent sql.NullFloat64
        derr := db.QueryRow(query, user_id).Scan(&spent)
        if derr != nil {
        log.Fatal(derr)
        }

        if !spent.Valid {
             log.Fatal("NO spent Record Found")
             c.JSON(http.StatusBadRequest,gin.H{"message": "Records Not Found "}) 
             return
            }
        c.JSON(http.StatusOK,gin.H{"spents":spent}) 
    }
}



func CheckBalance() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.DB_connection()        
        //error Handling while making Connection 
        if err!=nil{
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        var balance float32
        query := `SELECT balance FROM account WHERE account.ID=$1`
        err = db.QueryRow(query, user_id).Scan(&balance)
        if err!=nil{
            log.Fatal(err)
        }

        c.JSON(http.StatusOK,gin.H{"balance":balance}) 
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

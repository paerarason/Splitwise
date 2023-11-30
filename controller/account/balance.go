package account

import (
    "github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
		 _ "github.com/lib/pq"
         "log"
)

func Balance() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.DB_connection()      
        //error Handling while making Connection 
        if err!=nil{
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }

        defer db.Close() 
		accountID, exists := c.Get("account_id")
        if !exists {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Account ID not found"})
            return
        }
        user_id,err:=CheckAccountID(accountID)
        if err!=nil{
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Account ID not found"})
            return
        }  
        var balance float32
        query := `SELECT balance FROM account WHERE ID=$1`
        
        dberr := db.QueryRow(query,user_id).Scan(&balance)
        //Handling while making Queries
        if dberr!=nil{
            log.Println(dberr) 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"balance":balance}) 
    }
}


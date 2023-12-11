package account

import (
    "github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	 _ "github.com/lib/pq"
   //  "database/sql"
   "log"

)
func GETspendAmount() gin.HandlerFunc {
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
        //Query_for_Account


        query := `SELECT COALESCE(SUM(transaction.amount), 0)
                   FROM transaction
                   WHERE transaction.spent_id = $1`
        var spent float32
        log.Println(user_id)
        err = db.QueryRow(query, user_id).Scan(&spent)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err})
            return
        }
         log.Println(spent)
    
        c.JSON(http.StatusOK,gin.H{"spents":spent})
}




}
package account

import (
    "github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	 _ "github.com/lib/pq"
     "database/sql"

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
        if accountID, ok := accountID.(float64); ok { 
        user_id := int(accountID)
        query := `SELECT COALESCE(SUM(transaction.amount), 0) AS total_amount
                   FROM account_Group
                   LEFT JOIN transaction  ON account_Group.ID = transaction.Account_Group_id
                   WHERE account_Group.ID = $1
        `
        var spent sql.NullFloat64
        err = db.QueryRow(query, user_id).Scan(&spent)
        if err != nil {
        }

        if !spent.Valid {
             c.JSON(http.StatusBadRequest,gin.H{"message": "Records Not Found "}) 
             return
            }
        c.JSON(http.StatusOK,gin.H{"spents":spent}) }
      c.JSON(http.StatusBadRequest,gin.H{"message": "Unautherised acess"}) 
             return 
              }
    }



package account

import (
    "github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
		 _ "github.com/lib/pq"
     "encoding/json"
)
type Transaction struct {
		ID             int
		AccountGroupID int
		Amount         float64
	}

func BillHistory() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.DB_connection()        
        //error Handling while making Connection 
        if err!=nil{
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        defer db.Close()
        user_id, exists := c.Get("account_id")
		if !exists {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Account ID not found"})
            return
		}
        query := `SELECT ID,Account_Group_id,amount FROM 
                   FROM transaction WHERE trandaction.spent_id=$1 AND trandactionrecieved_id=$2`
        rows,derr := db.Query(query,user_id,user_id)
        if derr != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in Getting Records"})
            return
        }
        defer rows.Close()
        var transactions [] Transaction
        for rows.Next() {
            var transaction Transaction
            if err := rows.Scan(&transaction.ID, &transaction.AccountGroupID, &transaction.Amount); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in Getting Records"})
                return
            }
            transactions = append(transactions, transaction)
        }

        if err := rows.Err(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in Getting Records"})
                return
        }

        jsonData, err := json.Marshal(transactions)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in Getting Records"})
                return
        }
            
            c.JSON(http.StatusOK,jsonData) 
        }
    }
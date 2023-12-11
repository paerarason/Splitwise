package account

import (
    "github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
		 _ "github.com/lib/pq"
     "encoding/json"
     "log"
     "time"
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
        thirtyDaysAgo := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
        query := `
        SELECT ID, Account_Group_id, amount 
        FROM transaction 
        WHERE (transaction.spent_id = $1 OR transaction.recieved_id = $2)
        AND created_at >= $3`
        rows,derr := db.Query(query,user_id,user_id,thirtyDaysAgo)
        if derr != nil {
            log.Println(derr)
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
            
            c.JSON(http.StatusOK,string(jsonData)) 
        }
    }
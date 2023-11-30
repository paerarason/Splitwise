package transaction

import (
	"github.com/paerarason/Splitwise/database"
    "github.com/paerarason/Splitwise/controller/account"
	"github.com/gin-gonic/gin"
    "net/http"
	 _ "github.com/lib/pq"
     "log"

)
type Transaction struct {
    account_group_id  int    `json:"account_groupid"`
    amount float32          `json:"amount"`
}

func SendAmount() gin.HandlerFunc {
        return func(c *gin.Context){
            db,err:=database.DB_connection()
            tx, err := db.Begin()
            defer db.Close()
            if err != nil {
                log.Fatal(err)
            }
             
            //error Handling while Serialize the json from the request to the Account Struct
            var tr Transaction
            accountID, exists := c.Get("account_id")
		    if !exists {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Account ID not found"})
                return
		    }
            if err:=c.BindJSON(&tr);err!=nil{
                c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
                return 
            }
                 
            user_id,err:=account.CheckAccountID(accountID)
            if err!=nil{
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Account ID not found"})
                return
            }

            // check from the group exists
            gp_id:=c.Param("id")
            var transferAmount float32
            err = tx.QueryRow("SELECT ID FROM account_Group WHERE group_id=$1,account_id= $2,",gp_id,user_id).Scan(&transferAmount)
            if err != nil {
                tx.Rollback() 
                c.JSON(http.StatusBadRequest,gin.H{"message": "No Groups or Transaction pending"}) 
                return 
            }
            
    
           //COUNT THE NUMBER OF USERS IN THAT GROUP
            var member_count int 
            err = tx.QueryRow("SELECT COUNT(ID) FROM account_Group WHERE group_id=$1",gp_id).Scan(&member_count)
            if err != nil {
                tx.Rollback() 
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Account ID not found"})
                return 
            }
        
            var accountABalance float32
            err = tx.QueryRow("SELECT balance FROM account WHERE ID = $1",user_id).Scan(&accountABalance)
            if err != nil {
                tx.Rollback() 
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Account ID not found"})
                return 
            }

            /*

             CHECK THE BALANCE > THE SPLIT AMOUNTH 
            */

            if accountABalance < transferAmount/float32(member_count) {
                tx.Rollback() 
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Account ID not found"})
                return 
            }
           
           var admin_id int 
           err = tx.QueryRow("SELECT admin_id FROM groups WHERE ID=$1",gp_id).Scan(&admin_id)
            if err != nil {
                tx.Rollback() 
                c.JSON(http.StatusInternalServerError, gin.H{"error": "GROUP ID not found"})
                return 
            }
            
            _, err = tx.Exec("UPDATE account SET balance = balance - $1 WHERE ID = $2", transferAmount/float32(member_count),accountID)
               if err != nil {
                tx.Rollback() 
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Transaction"})
                return 
            }

            _, err = tx.Exec("UPDATE account SET balance = balance + $1 WHERE ID = $2",transferAmount/float32(member_count),admin_id)
               if err != nil {
                tx.Rollback() 
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Transaction"})
                return 
            }
           
           var account_group_id int 
           err = tx.QueryRow("SELECT ID FROM account_Group WHERE group_id=$1,account_id=$2",gp_id,user_id).Scan(&account_group_id)
            if err != nil {
                tx.Rollback() 
                c.JSON(http.StatusInternalServerError, gin.H{"error": "account_group_id not found"})
                return 
            }
          
             var trid int
            _, err = tx.Exec("INSERT  INTO transaction(spent_id,recieved_id,Account_Group_id,amount) VALUES($1,$2,$3,$4)",accountID,admin_id,account_group_id,transferAmount/float32(member_count))
            if err != nil {
                tx.Rollback() 
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Transaction"})
                return 
            }
            
            err = tx.Commit()
               if err != nil {
                tx.Rollback() 
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Transaction"})
                return 
            }

            c.JSON(http.StatusOK,gin.H{"Transaction ID":trid}) 
        }
    }
        

        
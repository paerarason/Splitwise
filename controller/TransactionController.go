package controller

import (
	"github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	 _ "github.com/lib/pq"
     "log"

)
type Transaction struct {
    account_group_id  int    `json:"account_groupid"`
    amount float32          `json:"amount"`
}





func CreateTransaction() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.DB_connection()
        //error Handling while making Connection 
        if err!=nil{
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }

        defer db.Close()        
        var tr Transaction

        //error Handling while Serialize the json from the request to the Account Struct 
        if err:=c.BindJSON(&tr);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        
        var trid int
        query := `INSERT INTO transaction (Account_Group_id,amount)
          VALUES ($1, $2) RETURNING id`
        
        dberr := db.QueryRow(query,tr.account_group_id,tr.amount).Scan(&trid)
        
        //Handling while making Queries
        if dberr!=nil{ 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"Transaction ID":trid}) 
    }
}



func SendAmount() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.DB_connection()
        tx, err := db.Begin()
        defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
    var tr Transaction
    //error Handling while Serialize the json from the request to the Account Struct 
    if err:=c.BindJSON(&tr);err!=nil{
        c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
        return 
    }
    
    var accountABalance float64
	err = tx.QueryRow("SELECT balance FROM account WHERE ID = $1",accountAID).Scan(&accountABalance)
	if err != nil {
		tx.Rollback() 
        log.Fatal(err)
	}
    
	if accountABalance < transferAmount {
		tx.Rollback() 
		return
	}


	_, err = tx.Exec("UPDATE account SET balance = balance - $1 WHERE ID = $2", transferAmount, accountAID)
	if err != nil {
		tx.Rollback() 
		log.Fatal(err)
	}

	_, err = tx.Exec("UPDATE account SET balance = balance + $1 WHERE ID = $2", transferAmount, accountBID)
	if err != nil {
		tx.Rollback() 
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

    var trid int
        query := `INSERT INTO transaction (Account_Group_id,amount)
          VALUES ($1, $2) RETURNING id`
        
        dberr := db.QueryRow(query,tr.account_group_id,tr.amount).Scan(&trid)
        
        //Handling while making Queries
        if dberr!=nil{ 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"Transaction ID":trid}) 
    }
}
	

	
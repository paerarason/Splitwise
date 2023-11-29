package controller

import (
	"github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	 _ "github.com/lib/pq"

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
        
        //Handlinf while making Queries
        if dberr!=nil{ 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"Transaction ID":trid}) 
    }
}


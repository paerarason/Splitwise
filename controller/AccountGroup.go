package controller

import (
	"github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	 _ "github.com/lib/pq"

)
type AccountGroup struct {
    account_id        int   `json:"account_id"`
    group_id  int    `json:"group_id"`    
}



func CreateAccountGroup() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.DB_connection()
        
        //error Handling while making Connection 
        if err!=nil{
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }

        defer db.Close()        
        var   ag AccountGroup

        //error Handling while Serialize the json from the request to the Account Struct 
        if err:=c.BindJSON(&ag);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        
        var agid int
        query := `INSERT INTO account_Group  (account_id,group_id)
          VALUES ($1, $2) RETURNING id`
        
        dberr := db.QueryRow(query,ag.account_id,ag.group_id).Scan(&agid)
        
        //Handlinf while making Queries
        if dberr!=nil{ 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"Group ID ":agid}) 
    }
}


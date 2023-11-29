package controller

import (
	"github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	 _ "github.com/lib/pq"

)
type Group struct {
    name        string   `json:"name"`
    description string   `json:"description"`
    split float32        `json:"splitfor"`
}



func CreateGroup() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.DB_connection()
        
        //error Handling while making Connection 
        if err!=nil{
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }

        defer db.Close()        
        var gp Group

        //error Handling while Serialize the json from the request to the Account Struct 
        if err:=c.BindJSON(&gp);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        
        var gpid int
        query := `INSERT INTO groups (name, description,split_for)
          VALUES ($1, $2, $3) RETURNING id`
        
        dberr := db.QueryRow(query,gp.name,gp.description,gp.split).Scan(&gpid)
        
        //Handlinf while making Queries
        if dberr!=nil{ 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"Group ID ":gpid}) 
    }
}


package controller

import (
	"github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	 _ "github.com/lib/pq"
     "database/sql"

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
        
        //error Handling while Serialize the json from the request to the Account Struct 
        var ag AccountGroup
        if err:=c.BindJSON(&ag);err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        
        
        // Check the authenticated User is admin to the Group 
        var admin_id int 
        query := "SELECT admin_id FROM groups WHERE ID = $1"
        Qerr := db.QueryRow(query,ag.group_id).Scan(&admin_id)
        if Qerr!=nil {
            if Qerr==sql.ErrNoRows{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Group Doesn't there"}) 
             return 
            }
        }


        if admin_id!= User.ID{
            //unautherised access
        }


        var agid int
        query_for_insert:= `INSERT INTO account_Group  (account_id,group_id)
          VALUES ($1, $2) RETURNING id`
        dberr := db.QueryRow(query_for_insert,ag.account_id,ag.group_id).Scan(&agid)
        //Handlinf while making Queries
        if dberr!=nil{ 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"Group ID ":agid}) 
    }
}


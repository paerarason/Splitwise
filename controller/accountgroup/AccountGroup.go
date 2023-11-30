package accountgroup

import (
	"github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	 _ "github.com/lib/pq"
     "database/sql"
     "fmt"
     "log"

)
type AccountGroup struct {
    Account_ids [] int   `json:"account_id"`
    Group_id  int    `json:"group_id"`
    
}


func CreateAccountGroup() gin.HandlerFunc {
    return func(c *gin.Context){
        db,err:=database.DB_connection()
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Service Error"})
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
        err = db.QueryRow(query,ag.Group_id).Scan(&admin_id)
        if err!=nil {
            if err==sql.ErrNoRows{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Group Doesn't there"}) 
             return 
            }
        }

        accountID, exists := c.Get("account_id")
        if !exists {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Account ID not found"})
            return
        }
        if accountID, ok := accountID.(float64); ok {
                accountIDInt := int(accountID)
                if accountIDInt!=admin_id{
            c.JSON(http.StatusBadRequest,gin.H{"message": "Unautherised acess"}) 
             return 
              }
        if admin_id!=accountIDInt {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Uautherised acess"})
            return
             }
               Query:="INSERT INTO account_Group  (account_id,group_id) VALUES "
        for i,value:=range ag.Account_ids {
            if i==0{
                Query+=fmt.Sprintf("(%d,%d)",value,ag.Group_id)
            }
            Query+=fmt.Sprintf(",(%d,%d)",value,ag.Group_id)
            log.Println(i)
        }

        Query+="RETURNING id"
        var id int 
        dberr := db.QueryRow(Query).Scan(&id)
        //Handlinf while making Queries
        if dberr!=nil{ 
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"message":id})
        }
           
    }
}


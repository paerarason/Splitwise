package group
import (
	"github.com/paerarason/Splitwise/database"
	"github.com/gin-gonic/gin"
    "net/http"
	 _ "github.com/lib/pq"
     "log"

)
type Group struct {
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Split       float32 `json:"split"`
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
            log.Println(err)
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "}) 
             return 
        }
        log.Println(gp.Name)
        var gpid int
        accountID, exists := c.Get("account_id")
		    if !exists {
                
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Account ID not found"})
                return
		    }
        query := `INSERT INTO groups (name, admin_id,description,split_for)
          VALUES ($1, $2, $3,$4) RETURNING id`
        
        err = db.QueryRow(query,gp.Name,accountID,gp.Description,gp.Split).Scan(&gpid)
        
        //Handlinf while making Queries
        if err!=nil{ 
            log.Println(err)
            c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request"}) 
             return 
        } 
        c.JSON(http.StatusOK,gin.H{"Group ID ":gpid}) 
    }
}


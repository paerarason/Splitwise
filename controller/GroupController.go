package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "fmt"
	//"database/sql"
	//_ "github.com/lib/pq"

)
type Group struct {
    Name        string   `json:"name"`
    Description string   `json:"description"`
}



func CreateGroup() gin.HandlerFunc {
    return func(c *gin.Context){
        var gp Group
        // connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"      
        // db,err:=sql.Open(connStr)
        // defer db.Close()
        if err:=c.BindJSON(&gp); err!=nil{
             c.JSON(http.StatusBadRequest,gin.H{"message": "Bad Request "})
             return  
         }
        fmt.Println("helllo")
         // if err!=nil{
        //   db.Query("SELECT Account ")
        // }

		c.JSON(http.StatusOK, gin.H{
      "message": gp.Name,
    })

	}
}




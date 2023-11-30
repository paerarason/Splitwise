package main
import (
	"github.com/gin-gonic/gin"
	"github.com/paerarason/Splitwise/controller/transaction"
	"github.com/paerarason/Splitwise/controller/account"
	"github.com/paerarason/Splitwise/controller/accountgroup"
	"github.com/paerarason/Splitwise/controller/group"
	"github.com/paerarason/Splitwise/middleware"
	"os"
	"log"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	
	router := gin.Default()
	err := godotenv.Load(".env")
    
	if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
	defer router.Run(":"+os.Getenv("PORT"))
    fmt.Println(os.Getenv("PORT"))
    
	router.POST("/login",middleware.GenerateToken())
    router.Use(middleware.JWTokenMiddlerware)
	//Bunch of APIS for account management 
	accounts:=router.Group("api/account")
	{
        accounts.POST("/",account.CreateAccount())
		accounts.GET("/balance",account.Balance())
		accounts.GET("/spent",account.GETspendAmount())
		accounts.GET("/history",account.BillHistory())
	}

    //Bunch of APIS for GROUP management 
    groups:=router.Group("api/groups")
	{
        groups.POST("/add",accountgroup.CreateAccountGroup())
        groups.POST("/",group.CreateGroup())
	}

	
	transactions:=router.Group("api/transaction")
	{
	    transactions.GET("/send/:id",transaction.SendAmount())
	}

}
package main
import (
	"github.com/gin-gonic/gin"
	//"github.com/paerarason/Splitwise/api"
	"github.com/paerarason/Splitwise/controller"
	"github.com/paerarason/Splitwise/middleware"
	"os"
	"fmt"
)

func main() {
	
	router := gin.Default()
	defer router.Run(":8000")
    fmt.Println(os.Getenv("PORT"))
    router.POST("/login",middleware.GenerateToken())
    router.Use(middleware.JWTokenMiddlerware)
	//Bunch of APIS for account management 
	account:=router.Group("api/account")
	{
        account.GET("/balance",controller.CheckBalance())
        account.POST("/",controller.CreateAccount())
		account.GET("/spent",controller.GETspendAmount())
	}

    //Bunch of APIS for GROUP management 
    groups:=router.Group("api/groups")
	{
        groups.POST("/add",controller.CreateAccountGroup())
        groups.POST("/",controller.CreateGroup())
	}

	
	// transaction:=router.Group("api/transaction")
	// {
	//     //transaction.POST("/send",controller.SendAmount())
	// }

}
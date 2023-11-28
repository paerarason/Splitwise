package main
import (
	"github.com/gin-gonic/gin"
	//"github.com/paerarason/Splitwise/api"
	"github.com/paerarason/Splitwise/controller"
)

func main() {
	
	router := gin.Default()
	defer router.Run(":8000")
	
	//Bunch of APIS for account management 
	account:=router.Group("/account")
	{
        //account.GET("/:id",)
        account.POST("/",controller.CreateAccount())
	}

    //Bunch of APIS for GROUP management 
    groups:=router.Group("/groups")
	{
        //groups.PATCH("/:id",)
		//groups.PATCH("/:id",)
		//groups.DELETE("/:id",)
        groups.POST("/",controller.CreateGroup())
	}

}
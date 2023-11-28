package main
import (
	"github.com/gin-gonic/gin"
	"github.com/paerarason/Splitwise/api"
)

func main() {
	router := gin.Default()
    api.AccountRouter(router)
    api.GroupRouter(router)

}
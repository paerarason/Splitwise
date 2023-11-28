package Controller
import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount() *gin.HandlerFunc {
       return func(c *gin.Context){

	   }
}








func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
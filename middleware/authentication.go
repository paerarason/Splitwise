package middleware
import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"github.com/paerarason/Splitwise/database"
	"net/http"
	"time"
)
func JWTokenMiddlerware(c *gin.Context){

	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Set("account_id", claims.AccountID)
	c.Next()
}



func GenerateToken() gin.HandlerFunc {
    return func(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")
	db,err:=database.DB_connection()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
			return
	}
	
	
	
	
	token, err := generateToken(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
		}

		c.SetCookie("token", token, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
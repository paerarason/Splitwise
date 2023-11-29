package middleware
import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/paerarason/Splitwise/database"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"time"
	"os"
)

type Claims struct {
	AccountID int `json:"account_id"`
	jwt.StandardClaims
}
func JWTokenMiddlerware(c *gin.Context){

	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("secret"), nil
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

    //get the account password from the database 
	var hash [] byte 
	var  accountID int
	Query:=`SELECT password,ID FROM account WHERE account.username=$1` 
    err=db.QueryRow(Query,username).Scan(&hash,&accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get Account details"})
		return
	}
    
	//Error handle for password Comaparison
	if !CheckPasswordHash(password,hash){
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Password Doesn't match"})
		return
	}
	
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		AccountID: accountID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString(os.Getenv("secret"))
	if err != nil  {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}
		c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}





func CheckPasswordHash(password string , hash []byte) bool {
    err := bcrypt.CompareHashAndPassword(hash, []byte(password))
    return err == nil
}
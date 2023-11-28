package middleware
import (
	"github.com/golang-jwt/jwt/v5"
	"fmt"
	"time"
)
func JWTokenMiddlerware(){
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"foo": "bar",
	"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),})
	tokenString, err := token.SignedString(hmacSampleSecret)
    fmt.Println(tokenString, err)

}
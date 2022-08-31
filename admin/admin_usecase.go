package admin

import (
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// GeneratingToken function is for generating token.
func GeneratingToken(id string) string {

	err := godotenv.Load("app.env")
	if err != nil {
		panic(err)
	}
	key := os.Getenv("SECRETKEY")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id 
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println("Generating token error:", err)
		panic(err)
	}
	return tokenString

	// claims = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	// 	Issuer:    id,
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	// })
	// token, err := claims.SignedString([]byte(key))
	// if err != nil {
	// 	panic(err)
	// }
	// return token
}



package admin

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func GeneratingToken(id string) string {

	err := godotenv.Load("./config/.env")
	if err != nil {
		panic(err)
	}
	key := os.Getenv("SECRETKEY")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    id,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}

	return token
}



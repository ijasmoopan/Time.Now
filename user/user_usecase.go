package user

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ijasmoopan/Time.Now/usecases"
	"github.com/joho/godotenv"
)

// GeneratingToken for generating jwt tokens.
func GeneratingToken(id string) string {

	file := usecases.Logger()
	log.SetOutput(file)

	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Println(err)
	}
	key := os.Getenv("USERSECRETKEY")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Fatalln("Error while generating token:", err)
	}
	return tokenString

	// claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
	// 	Issuer:    id,
	// 	ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	// })
	// token, err := claims.SignedString([]byte(key))
	// if err != nil {
	// 	log.Println(err)
	// }

}
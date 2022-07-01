package user

// import (
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// )

// func GeneratingToken(id string) string {

// 	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
// 		Issuer:    id,
// 		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
// 	})
// 	token, err := claims.SignedString([]byte(key))
// 	if err != nil {
// 		panic(err)
// 	}

// 	return token
// }
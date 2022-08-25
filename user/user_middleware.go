package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/ijasmoopan/Time.Now/models"
	"github.com/ijasmoopan/Time.Now/usecases"
	"github.com/joho/godotenv"
)

//IsUserAuthorized for checking jwt token.
func (repo *Repo) IsUserAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		file := usecases.Logger()
		log.SetOutput(file)

		err := godotenv.Load("./.gitignore/.env")
		if err != nil {
			log.Println("Can't access env file")
		}
		key := []byte(os.Getenv("USERSECRETKEY"))

		var reqToken string
		if r.Header["Authorization"] != nil {
			reqToken = r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			message := map[string]interface{}{
				"msg": "Token not found",
			}
			json.NewEncoder(w).Encode(&message)
			return
		}
		token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Error in token method")
			}
			return key, nil
		})
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			message := map[string]interface{}{
				"msg":   "Unauthorized token",
				"error": err,
			}
			json.NewEncoder(w).Encode(&message)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			id := claims["id"]
			user, err := repo.user.DBAuthUser(id.(string))
			if err != nil {
				log.Println("Error:", err)
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				message := map[string]interface{}{
					"response": "Please Login",
				}
				json.NewEncoder(w).Encode(&message)
				return
			}
			ctx := context.WithValue(r.Context(), models.CtxKey{}, user)
			handler.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

// IsHomeUserAuthorized method for authorizing user from home.
func (repo *Repo) IsHomeUserAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		file := usecases.Logger()
		log.SetOutput(file)

		err := godotenv.Load("./.gitignore/.env")
		if err != nil {
			log.Println("Can't fetch env file")
		}
		key := []byte(os.Getenv("USERSECRETKEY"))

		var reqToken string
		if r.Header["Authorization"] != nil {
			reqToken = r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
		} else {
			log.Println("Token not found")
			handler.ServeHTTP(w, r)
			return
		}

		token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Error in token method")
			}
			return key, nil
		})
		if err != nil {
			log.Println("Unauthorized token")
			handler.ServeHTTP(w, r)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			id := claims["id"]
			user, err := repo.user.DBAuthUser(id.(string))
			if err != nil {
				log.Println("Incorrect user")
				handler.ServeHTTP(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), models.CtxKey{}, user)
			handler.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

// DeleteToken for deleting jwt token when user logging out.
func (repo *Repo) DeleteToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// cookie := http.Cookie{
		// 	Name:     "jwt",
		// 	Value:    "",
		// 	Expires:  time.Now().Add(-time.Hour),
		// 	HttpOnly: true,
		// }
		// http.SetCookie(w, &cookie)
		err := godotenv.Load("./.gitignore/.env")
		if err != nil {
			log.Println("Can't access env file")
		}
		key := []byte(os.Getenv("USERSECRETKEY"))

		var reqToken string
		if r.Header["Authorization"] != nil {
			reqToken = r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
		}
		token, err := jwt.Parse(reqToken, func (token *jwt.Token)(interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Error in token method")
			}
			return key, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			claims["authorized"] = false
		}
		// tokenString, err := token.SignedString(key)
		// if err != nil {
		// 	log.Fatalln("Error in deleting token")
		// }
		// log.Println("Logout Token:", tokenString)

		handler.ServeHTTP(w, r)
	})
}

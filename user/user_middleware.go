package user

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ijasmoopan/Time.Now/models"
	"github.com/ijasmoopan/Time.Now/usecases"
	"github.com/joho/godotenv"
)

//IsUserAuthorized for checking jwt token.
func (repo *Repo) IsUserAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		file := usecases.Logger()
		log.SetOutput(file)

		err := godotenv.Load("./config/.env")
		if err != nil {
			log.Println("Can't access env file")
		}
		key := os.Getenv("SECRETKEY")

		cookie, err := r.Cookie("jwt")
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			message := map[string]interface{}{
				"msg": "Please Login",
				"error": err,
			}
			json.NewEncoder(w).Encode(&message)
			return
		}
		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
			return []byte(key), nil
		})
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			message := map[string]interface{}{
				"msg": "Please Login",
				"error": err,
			}
			json.NewEncoder(w).Encode(&message)
			return
		}
		claims := token.Claims.(*jwt.StandardClaims)

		user, err := repo.user.DBAuthUser(claims.Issuer)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			message := map[string]interface{}{
				"msg": "Please Login",
				"error": err,
			}
			json.NewEncoder(w).Encode(&message)
			return
		}	

		// type CtxKey struct {}
			log.Println("User ID in middleware:", user.ID)

		ctx := context.WithValue(r.Context(), models.CtxKey{}, user) //nolint
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

// IsHomeUserAuthorized method for authorizing user from home.
func (repo *Repo) IsHomeUserAuthorized(handler http.Handler)http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		file := usecases.Logger()
		log.SetOutput(file)
		
		err := godotenv.Load("./config/.env")
		if err != nil {
			log.Println("Can't fetch env file")
		}
		key := os.Getenv("SECRETKEY")

		cookie, err := r.Cookie("jwt")
		if err != nil {
			log.Println("No cookie")
			handler.ServeHTTP(w, r)
			return
		}
		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
			return []byte(key), nil
		})
		if err != nil {
			log.Println("Can't decode token")
			handler.ServeHTTP(w, r)
			return
		}
		claims := token.Claims.(*jwt.StandardClaims)

		user, err := repo.user.DBAuthUser(claims.Issuer)
		if err != nil {
			log.Println("Incorrect user")
			handler.ServeHTTP(w, r)
			return
		}
		log.Println("In middleware.. UserID:", user.ID)
		ctx := context.WithValue(r.Context(), models.CtxKey{}, user)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DeleteToken for deleting jwt token when user logging out.
func (repo *Repo) DeleteToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		log.Println()
		cookie := http.Cookie {
			Name: "jwt",
			Value: "",
			Expires: time.Now().Add(-time.Hour),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		handler.ServeHTTP(w, r)
	})
}
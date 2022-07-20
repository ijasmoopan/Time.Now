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
	"github.com/joho/godotenv"
)

//IsUserAuthorized for checking jwt token.
func (repo *Repo) IsUserAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		err := godotenv.Load("./config/.env")
		if err != nil {
			
		}
		key := os.Getenv("SECRETKEY")

		cookie, err := r.Cookie("jwt")
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			message := map[string]interface{}{
				"msg": "There is no cookie",
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
			w.WriteHeader(http.StatusBadRequest)
			message := map[string]interface{}{
				"msg": "There is no cookie",
				"error": err,
			}
			json.NewEncoder(w).Encode(&message)
			return
		}
		claims := token.Claims.(*jwt.StandardClaims)

		user, err := repo.user.DBAuthUser(claims.Issuer)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			message := map[string]interface{}{
				"msg": "There is no cookie",
				"error": err,
			}
			json.NewEncoder(w).Encode(&message)
			return
		}	
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
package admin

import (
	"context"
	"encoding/json"
	"os"

	// "database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/go-chi/chi/v5"
	"github.com/ijasmoopan/Time.Now/models"
	"github.com/ijasmoopan/Time.Now/usecases"
	"github.com/joho/godotenv"
)


// IsAdminAuthorized for authorizing admin by middleware.
func (repo *Repo) IsAdminAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		err := godotenv.Load("./config/.env")
		if err != nil {
			panic(err)
		}
		key := os.Getenv("SECRETKEY")

		file := usecases.Logger()
		log.SetOutput(file)

		cookie, err := r.Cookie("jwt")
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			message := map[string]interface{}{
				"msg": "There is no Cookie",
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
				"msg": "Error while parsing Token",
			}
			json.NewEncoder(w).Encode(&message)
			return
		}
		claims := token.Claims.(*jwt.StandardClaims)

		var admin models.Admin
		admin, err = repo.admin.DBGetAdminByID(claims.Issuer)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			message := map[string]interface{}{
				"msg": "Admin Not Found",
			}
			json.NewEncoder(w).Encode(&message)
			return
		}
		ctx := context.WithValue(r.Context(), models.CtxKey{}, admin)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DeletingJWT for deleting jwt token when logoutting.
func (repo *Repo) DeletingJWT(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		file := usecases.Logger()
		log.SetOutput(file)

		cookie := http.Cookie{
			Name: "jwt",
			Value: "",
			Expires: time.Now().Add(-time.Hour),
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		
		handler.ServeHTTP(w, r)
	})
}



// func (repo *Repo) ProductCtx(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

// 		file := usecases.Logger()
// 		log.SetOutput(file)

// 		product_id := chi.URLParam(r, "product_id")

// 		product, err := repo.product.DBGetProduct(product_id)
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		ctx := context.WithValue(r.Context(), "product", product)
// 		handler.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
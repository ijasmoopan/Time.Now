package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/ijasmoopan/Time.Now/models"
	"github.com/ijasmoopan/Time.Now/usecases"
	"github.com/joho/godotenv"
)

// IsAdminAuthorized for authorizing admin by middleware.
func (repo *Repo) IsAdminAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		err := godotenv.Load("app.env")
		if err != nil {
			panic(err)
		}
		key := []byte(os.Getenv("SECRETKEY"))

		file := usecases.Logger()
		log.SetOutput(file)

		var reqToken string
		if r.Header["Authorization"] != nil {
			reqToken = r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
		} else {
			log.Println("Token not found")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			message := map[string]interface{}{
				"msg": "Token Not Found",
			}
			json.NewEncoder(w).Encode(&message)
			return
		}
		token, err := jwt.Parse(reqToken, func(token *jwt.Token)(interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error while parsing")
			}
			return key, nil
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
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id := claims["id"]
			log.Println("ID:", id)
			var admin models.Admin
			admin, err = repo.admin.DBGetAdminByID(id.(string))
			if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				message := map[string]interface{}{
					"msg": "Admin Not Found",
				}
				json.NewEncoder(w).Encode(&message)
				return
			}
			log.Println("Admin:", admin)
			ctx := context.WithValue(r.Context(), models.CtxKey{}, admin)
			handler.ServeHTTP(w, r.WithContext(ctx))
		} else {
			log.Println("Admin not found")
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			message := map[string]interface{}{
				"msg": "Admin Not Found",
			}
			json.NewEncoder(w).Encode(&message)
		}
	})
}

// DeletingJWT for deleting jwt token when logoutting.
func (repo *Repo) DeletingJWT(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		file := usecases.Logger()
		log.SetOutput(file)

		// cookie := http.Cookie{
		// 	Name: "jwt",
		// 	Value: "",
		// 	Expires: time.Now().Add(-time.Hour),
		// 	HttpOnly: true,
		// }
		// http.SetCookie(w, &cookie)

		if r.Header["Authorization"] != nil || r.Header["Authorization"] == nil {
			r.Header["Authorization"] = []string{}
		}
		
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
package admin

import (
	"context"
	"os"
	// "database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/ijasmoopan/Time.Now/models"
	"github.com/ijasmoopan/Time.Now/usecases"
	"github.com/joho/godotenv"
)



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
			log.Println("Redirecting to admin..")
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
			return []byte(key), nil
		})
		if err != nil {
			log.Println(err)
		}
		claims := token.Claims.(*jwt.StandardClaims)

		var admin models.Admin

		admin, err = repo.adminbyid.DBGetAdminById(claims.Issuer)
		if err != nil {
			log.Println("Redirecting to login..")
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), "admin", admin)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

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
		log.Println("Token Deleted... Redirecting to Login...")

		handler.ServeHTTP(w, r)
	})
}

func (repo *Repo) UserCtx(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		file := usecases.Logger()
		log.SetOutput(file)

		user_id := chi.URLParam(r, "userid")
		log.Println("User URL Param: ", user_id)
		user, err := repo.user.DBGetUser(user_id)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(404), 404)
			return
		}
		log.Println("Context User: ", user)
		ctx := context.WithValue(r.Context(), "user", user)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (repo *Repo) ProductCtx(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		file := usecases.Logger()
		log.SetOutput(file)

		product_id := chi.URLParam(r, "product_id")

		product, err := repo.product.DBGetProduct(product_id)
		if err != nil {
			log.Println(err)
		}

		ctx := context.WithValue(r.Context(), "product", product)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
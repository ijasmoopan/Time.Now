package middleware

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/ijasmoopan/Time.Now/admin"
// 	"github.com/ijasmoopan/Time.Now/models"
// 	"github.com/ijasmoopan/Time.Now/repository"

// 	"github.com/dgrijalva/jwt-go"
// )

// const key = "secretKey"

// func (repo *Repo) IsAdminAuthorized(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

// 		fmt.Println("Token checking by middleware..")

// 		cookie, err := r.Cookie("jwt")
// 		if err != nil {

// 			http.Redirect(w, r, "/admin", http.StatusSeeOther)
// 			return
// 		}
// 		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
// 			return []byte(key), nil
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 		claims := token.Claims.(*jwt.StandardClaims)

// 		var admin models.Admin

// 		db := repository.ConnectDB()
// 		defer repository.CloseDB(db)

// 		if result := db.First(&admin, "admin_id = ? ", claims.Issuer); result.Error != nil {
// 			http.Redirect(w, r, "/admin", http.StatusSeeOther)
// 			return
// 		} else if result.RowsAffected == 0 {
// 			http.Redirect(w, r, "/admin", http.StatusSeeOther)
// 			return
// 		} 
// 		fmt.Println("Admin Authorized..")
// 		handler.ServeHTTP(w, r)
// 	})
// }

// func DeletingJWT(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

// 		cookie := http.Cookie{
// 			Name: "jwt",
// 			Value: "",
// 			Expires: time.Now().Add(-time.Hour),
// 			HttpOnly: true,
// 		}
// 		http.SetCookie(w, &cookie)
// 		fmt.Println("Token Deleted... Redirecting to Login...")

// 		handler.ServeHTTP(w, r)
// 	})
// }

// func UserCtx(handler http.Handler) http.Handler{
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

// 		user_id := chi.URLParam(r, "user_id")
// 		fmt.Println("URL Param: ", user_id)
// 		user, err := admin.DBGetUser(user_id)
// 		if err != nil {
// 			http.Error(w, http.StatusText(404), 404)
// 			return
// 		}
// 		fmt.Println("Context User: ", user)
// 		ctx := context.WithValue(r.Context(), "user", user)
// 		handler.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
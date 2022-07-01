package middleware

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/ijasmoopan/Time.Now/models"
// 	"github.com/ijasmoopan/Time.Now/repository"

// 	"github.com/dgrijalva/jwt-go"
// )

// func IsUserAuthorized(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

// 		fmt.Println("Token checking by middleware..")

// 		cookie, err := r.Cookie("jwt")
// 		if err != nil {

// 			http.Redirect(w, r, "/login", http.StatusSeeOther)
// 			return
// 		}
// 		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
// 			return []byte(key), nil
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 		claims := token.Claims.(*jwt.StandardClaims)

// 		var user models.User

// 		db := repository.ConnectDB()
// 		defer repository.CloseDB(db)

// 		if result := db.First(&user, "user_id = ? ", claims.Issuer); result.Error != nil {
// 			http.Redirect(w, r, "/login", http.StatusSeeOther)
// 			return
// 		} else if result.RowsAffected == 0 {
// 			http.Redirect(w, r, "/login", http.StatusSeeOther)
// 			return
// 		} 
// 		fmt.Println("User Authorized..")
		
// 		handler.ServeHTTP(w, r)
// 	})
// }
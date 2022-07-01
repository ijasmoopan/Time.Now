package user

// import (
// 	"crypto/md5"
// 	"encoding/hex"
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	// "github.com/go-chi/chi/v5"
// 	"github.com/ijasmoopan/Time.Now/admin"
// 	"github.com/ijasmoopan/Time.Now/models"
// )
	
// const key = "secretKey"

// func UserLoginPage(w http.ResponseWriter, r *http.Request){

// 	tpl, err := template.ParseFiles("./userTemplates/signin.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	tpl.Execute(w, nil)
// }

// func UserLoginValidating(w http.ResponseWriter, r *http.Request){

// 	var userForm models.User
// 	var user models.User

// 	fmt.Println("Validating user details..")

// 	// err := json.NewDecoder(r.Body).Decode(&user)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	r.ParseForm()
// 	userForm.User_email = r.FormValue("email")
// 	password := r.FormValue("password")
// 	userForm.User_password = hex.EncodeToString(md5.New().Sum([]byte(password)))
// 	fmt.Println("User: ", userForm)

// 	user, err := ValidateUser(userForm)
// 	if err != nil {
// 		tpl, err := template.ParseFiles("./userTemplates/signin.html")
// 		if err != nil {
// 			panic(err)
// 		}
// 		msg := map[string]interface{}{
// 			"err": "Incorrect username or password",
// 		}
// 		tpl.Execute(w, msg)
// 	}
// 	fmt.Println("User Validated")

// 	id := strconv.Itoa(int(user.User_id))
// 	token := GeneratingToken(id)
	
// 	cookie := http.Cookie{
// 		Name: "jwt",
// 		Value: token,
// 		Expires: time.Now().Add(time.Hour * 24),
// 		HttpOnly: true,
// 	}
// 	http.SetCookie(w, &cookie)
// 	fmt.Println("Token generated..Redirecting to Home...")

// 	http.Redirect(w, r, "/home/"+user.User_firstname, http.StatusSeeOther)
// }

// func UserRegistration(w http.ResponseWriter, r *http.Request){

// 	tpl, err := template.ParseFiles("./userTemplates/register.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	tpl.Execute(w, nil)
// }

// func UserRegistering(w http.ResponseWriter, r *http.Request){

// 	var newUser models.User
	
// 	r.ParseForm()
// 	newUser.User_firstname = r.FormValue("firstname")
// 	newUser.User_secondname = r.FormValue("lastname")
// 	newUser.User_email = r.FormValue("email")
// 	newUser.User_gender = r.FormValue("gender")
// 	newUser.User_phone = r.FormValue("phone_number")
// 	if r.FormValue("referral") != "" {
// 		newUser.User_referral = r.FormValue("referral")	
// 	}
// 	if r.FormValue("password") == r.FormValue("confirm") {
// 		password := r.FormValue("password")
// 		// Password: MD5 Hashing
// 		newUser.User_password = hex.EncodeToString(md5.New().Sum([]byte(password)))
// 	} else {
		
// 		tpl, err := template.ParseFiles("./userTemplates/register.html")
// 		if err != nil {
// 			panic(err)
// 		}
// 		message := map[string]interface{}{
// 			"msg": "Passwords are different",
// 		}
// 		tpl.Execute(w, message)
// 	}
// 	fmt.Println("Registration is on process..\n New user: ", newUser)

// 	newUser, err := RegisteringUser(newUser)
// 	if err != nil {
// 		http.Error(w, http.StatusText(404), 404)
// 	}
// 	fmt.Println("Registration completed with ID: ",newUser)

// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
// }

// func UserLogout(w http.ResponseWriter, r *http.Request){

// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
// }

// func HomePage(w http.ResponseWriter, r *http.Request){

// 	tpl, err := template.ParseFiles("./userTemplates/index.html")
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	tpl.Execute(w, nil)
// }

// func Home(w http.ResponseWriter, r *http.Request) {

// 	products, err := admin.DBGetAllProducts()
// 	if err != nil {
// 		fmt.Println("DB Error: ", err)
// 		http.Error(w, http.StatusText(404), 404)
// 		return
// 	}
// 	fmt.Println("Listing products..")
// 	tpl, err := template.ParseFiles("./userTemplates/index.html")
// 	if err != nil {
// 		fmt.Println("Template Error: ", err)
// 		http.Error(w, http.StatusText(404), 404)
// 		return
// 	}
// 	tpl.Execute(w, map[string]interface{}{
// 		"Products": products,
// 	})
// }

// // func ProductView(w http.ResponseWriter, r *http.Request){

// // 	product_id := chi.URLParam(r, "product_id")

// // 	product, err := admin.DBGetProduct(product_id)
// // 	if err != nil {
// // 		http.Error(w, http.StatusText(404), 404)
// // 		return
// // 	}
// // 	tpl, err := template.ParseFiles("./userTemplates/product-detail.html")
// // 	if err != nil {
// // 		http.Error(w, http.StatusText(404), 404)
// // 		return
// // 	}
// // 	tpl.Execute(w, map[string]interface{}{
// // 		"Product_id": product.Product_id,
// // 		"Product_name": product.Product_name,
// // 		"Product_price": product.Product_price,
// // 		"Product_color": product.Product_color,
// // 		"Product_desc": product.Product_desc,
// // 		"Product_image": product.Product_image,
// // 	})
// // }


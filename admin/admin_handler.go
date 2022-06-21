package admin

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ijasmoopan/Time.Now/models"
	"github.com/shopspring/decimal"
)

const key = "secretKey"

// ---------------------Admin Side---------------------------

func AdminLoginPage(w http.ResponseWriter, r *http.Request){

	tpl, err := template.ParseFiles("./adminTemplates/login.html")
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, nil)
	
}

func AdminLoginValidating(w http.ResponseWriter, r *http.Request){

	var admin models.Admin
	var adminForm models.Admin

	r.ParseForm()
	adminForm.Admin_name = r.FormValue("adminname")
	password := r.FormValue("adminpassword")
	adminForm.Admin_password = hex.EncodeToString(md5.New().Sum([]byte(password)))
	fmt.Println("Admin: ", adminForm)

	admin, err := ValidateAdmin(adminForm)
	if err != nil {
		tpl, err := template.ParseFiles("./adminTemplates/login.html")
		if err != nil {
			panic(err)
		}
		fmt.Println("He is not an admin")
		message := map[string]interface{}{
			"msg": "Incorrect Name or Password",
		}
		tpl.Execute(w, message)
		return
	}
	fmt.Println("Admin Validated")

	id := strconv.Itoa(int(admin.Admin_id))
	token := GeneratingToken(id)

	cookie := http.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Println("Token Created...Redirecting to home...")

	http.Redirect(w, r, "/admin/home/"+admin.Admin_name, http.StatusSeeOther)
}	

func AdminLogout(w http.ResponseWriter, r *http.Request){

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func AdminHome(w http.ResponseWriter, r *http.Request){

	tpl, err := template.ParseFiles("./adminTemplates/admin-home.html")
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, nil)

	fmt.Println("Home Page")
}


// ----------------------------User Side---------------------------

func UserListTable(w http.ResponseWriter, r *http.Request){

	users, allUsers, err := DBGetUsers()
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
	}

	tpl, err := template.ParseFiles("./adminTemplates/data-tables.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, map[string]interface{}{
		"user": users,
		"allUser": allUsers,
	})
	if err != nil {
		panic(err)
	}
}

func ViewUser(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()
	userDetails, ok :=  ctx.Value("user").(models.User)
	if !ok {
		fmt.Println("Error User: ", userDetails)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	fmt.Println("View User: ", userDetails)

	tpl, err := template.ParseFiles("./adminTemplates/admin_viewuser.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, map[string]interface{}{
		"user_id": userDetails.User_id,
		"user_firstname": userDetails.User_firstname,
		"user_secondname": userDetails.User_secondname,
		"user_email": userDetails.User_email,
		"user_phone": userDetails.User_phone,
		"user_gender": userDetails.User_gender,
		"user_referral": userDetails.User_referral,
		"user_status": userDetails.User_status,
		"user_createdat": userDetails.Created_at,
		"user_updatedat": userDetails.Updated_at,
		"user_deletedat": userDetails.Deleted_at,
	})
	if err != nil {
		panic(err)
	}
}

func EditUser(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()
	userDetails, ok := ctx.Value("user").(models.User)
	if !ok {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	fmt.Println("Edit user: ", userDetails)

	tpl, err := template.ParseFiles("./adminTemplates/admin_edituser.html")
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, map[string]interface{}{
		"user_id": userDetails.User_id,
		"user_firstname": userDetails.User_firstname,
		"user_secondname": userDetails.User_secondname,
		"user_email": userDetails.User_email,
		"user_phone": userDetails.User_phone,
		"user_gender": userDetails.User_gender,
		"user_status": userDetails.User_status,
	})
}

func EditingUser(w http.ResponseWriter, r *http.Request){

	r.ParseForm()

	var user models.User
	user_id := chi.URLParam(r, "user_id")

	user.User_firstname = r.FormValue("firstname")
	user.User_secondname = r.FormValue("secondname") 
	user.User_gender = r.FormValue("gender") 
	user.User_email = r.FormValue("email") 
	user.User_phone = r.FormValue("phone") 

	err := UpdatingUser(user_id, user)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	
	http.Redirect(w, r, "/admin/userlist/"+user_id+"/view", http.StatusSeeOther)
}

func BlockUser(w http.ResponseWriter, r *http.Request){

	user_id := chi.URLParam(r, "user_id")

	err := DBGetUserStatus(user_id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	http.Redirect(w, r, "/admin/userlist", http.StatusSeeOther)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){

	user_id := chi.URLParam(r, "user_id")

	err := DBDeleteUser(user_id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	http.Redirect(w, r, "/admin/userlist", http.StatusSeeOther)	
}


// -------------------Product Side------------------------

func ProductList(w http.ResponseWriter, r *http.Request){

	products, err := DBGetAllProducts()
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	fmt.Println("Product listing...")
	tpl, err := template.ParseFiles("./adminTemplates/ecommerce-product.html")
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	tpl.Execute(w, map[string]interface{}{
		"product": products,
	})
}

func ProductView(w http.ResponseWriter, r *http.Request){

	var product models.Product
	product_id := chi.URLParam(r, "product_id")
	product, err := DBGetProduct(product_id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	fmt.Println("Product: ", product)
	tpl, err := template.ParseFiles("./adminTemplates/ecommerce-product-single.html")
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	tpl.Execute(w, map[string]interface{}{
		"Product_id": product.Product_id,
		"Product_name": product.Product_name,
		"Product_price": product.Product_price,
		"Product_color": product.Product_color,
		"Product_desc": product.Product_desc,
		"Product_image": product.Product_image,
	})
}

func EditProduct(w http.ResponseWriter, r *http.Request){

	product_id := chi.URLParam(r, "product_id")
	product, err := DBGetProduct(product_id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	tpl, err := template.ParseFiles("./adminTemplates/product-edit.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	} 
	tpl.Execute(w, map[string]interface{}{
		"Product_id": product.Product_id,
		"Product_name": product.Product_name,
		"Product_price": product.Product_price,
		"Product_color": product.Product_color,
		"Product_desc": product.Product_desc,
		"Product_image": product.Product_image,
		"Product_quantity": product.Product_quantity,
	})
}

func EditingProduct(w http.ResponseWriter, r *http.Request){

	r.ParseForm()
	product_id := chi.URLParam(r, "product_id")
	var editProduct models.Product
	editProduct.Product_name = r.FormValue("productname")

	price, err := decimal.NewFromString(r.FormValue("productprice"))
	if err == nil {
		editProduct.Product_price = price
	}
	editProduct.Product_color = r.FormValue("productcolor")
	editProduct.Product_desc = r.FormValue("productdesc")
	quantity, err := strconv.Atoi(r.FormValue("productquantity"))
	if err == nil {
		editProduct.Product_quantity = uint(quantity)
	}
	err = DBUpdateProduct(product_id, editProduct)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	http.Redirect(w, r, "/admin/productlist/"+product_id+"/view", http.StatusSeeOther)
}
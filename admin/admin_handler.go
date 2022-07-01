package admin

import (
	// "context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ijasmoopan/Time.Now/models"
	"github.com/ijasmoopan/Time.Now/usecases"
	"github.com/shopspring/decimal"
)



// ---------------------Admin Side---------------------------

func (repo *Repo) AdminLoginPage(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)
	tpl, err := template.ParseFiles("./adminTemplates/admin-login.html")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	tpl.Execute(w, nil)
}

func (repo *Repo) AdminLoginValidating(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	var adminForm models.Admin

	r.ParseForm()
	adminForm.Admin_name = r.FormValue("admin_name")
	adminForm.Admin_password = hex.EncodeToString(md5.New().Sum([]byte(r.FormValue("admin_password"))))

	log.Println("Admin Logging in: ", adminForm)

	admin, err := repo.admin.DBGetAdmin(adminForm)
	log.Println(err)
	log.Println("Result form database: ", admin)
	if err != nil {
		tpl, err := template.ParseFiles("./adminTemplates/admin-login.html")
		if err != nil {
			log.Println("Template error: ", err)
		}
		log.Println(err)
		message := map[string]interface{}{
			"msg": "Incorrect Name or Password",
		}
		tpl.Execute(w, message)
		return
	}
	log.Println("Admin Validated")

	id := strconv.Itoa(int(admin.Admin_id))
	token := GeneratingToken(id)

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	log.Println("Token Created...Redirecting to home...")

	http.Redirect(w, r, "/admin/home/"+admin.Admin_name, http.StatusSeeOther)
}

func (repo *Repo) AdminLogout(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (repo *Repo) GetAdminById(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	log.Println("GetAdminById...")
	admin_id := chi.URLParam(r, "admin_id")

	log.Println("Admin Id:", admin_id)

	admin, err := repo.adminbyid.DBGetAdminById(admin_id)
	if err != nil {
		log.Println("Error in DBGetAdminById: ", err)
	}

	log.Println("Admin by id: ", admin)
	fmt.Fprint(w, admin.Admin_id, admin.Admin_name, admin.Admin_password)
}

func (repo *Repo) AdminHome(w http.ResponseWriter, r *http.Request) {

	admin_name := chi.URLParam(r, "adminname")

	file := usecases.Logger()
	log.SetOutput(file)

	tpl, err := template.ParseFiles("./adminTemplates/admin-home.html")
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, map[string]interface{}{
		"admin": admin_name,
	})
	log.Println("Home Page")
}



// ----------------------------User Side---------------------------

func (repo *Repo) UserListTable(w http.ResponseWriter, r *http.Request) {

	admin_name := chi.URLParam(r, "adminname")
	file := usecases.Logger()
	log.SetOutput(file)

	users, err := repo.users.DBGetUsers()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), 404)
	}

	tpl, err := template.ParseFiles("./adminTemplates/admin-userlist.html")
	if err != nil {
		panic(err)
	}
	log.Println(users)
	err = tpl.Execute(w, map[string]interface{}{
		"admin": admin_name,
		"user":  users,
	})
	if err != nil {
		panic(err)
	}
}

func (repo *Repo) ViewUser(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	userDetails, ok := ctx.Value("user").(models.User)
	if !ok {
		log.Println("Error User: ", userDetails)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	log.Println("View User: ", userDetails)

	tpl, err := template.ParseFiles("./adminTemplates/admin_viewuser.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, map[string]interface{}{
		"user_id":         userDetails.User_id,
		"user_firstname":  userDetails.User_firstname,
		"user_secondname": userDetails.User_secondname,
		"user_email":      userDetails.User_email,
		"user_phone":      userDetails.User_phone,
		"user_gender":     userDetails.User_gender,
		"user_referral":   userDetails.User_referral,
		"user_status":     userDetails.User_status,
		"user_createdat":  userDetails.Created_at,
		"user_updatedat":  userDetails.Updated_at,
		"user_deletedat":  userDetails.Deleted_at,
	})
	if err != nil {
		panic(err)
	}
}

func (repo *Repo) EditUser(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	userDetails, ok := ctx.Value("user").(models.User)
	if !ok {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	log.Println("Edit user: ", userDetails)

	tpl, err := template.ParseFiles("./adminTemplates/admin_edituser.html")
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, map[string]interface{}{
		"user_id":         userDetails.User_id,
		"user_firstname":  userDetails.User_firstname,
		"user_secondname": userDetails.User_secondname,
		"user_email":      userDetails.User_email,
		"user_phone":      userDetails.User_phone,
		"user_status":     userDetails.User_status,
		"user_referral":   userDetails.User_referral,
	})
}

func (repo *Repo) EditingUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	var user models.User
	user_id := chi.URLParam(r, "userid")

	user.User_firstname = r.FormValue("firstname")
	user.User_secondname = r.FormValue("secondname")
	user.User_email = r.FormValue("email")
	user.User_phone = r.FormValue("phone")
	user.User_referral = r.FormValue("referral")

	err := repo.updateuser.DBUpdatingUser(user_id, user)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	http.Redirect(w, r, "/admin/home/userlist/"+user_id+"/viewuser", http.StatusSeeOther)
}

func (repo *Repo) BlockUser(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin, ok := ctx.Value("admin").(models.Admin)
	if !ok {
		log.Println("Can't access context admin")
	}

	user_id := chi.URLParam(r, "userid")

	err := repo.userStatus.DBGetUserStatus(user_id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/userlist", http.StatusSeeOther)
}

func (repo *Repo) DeleteUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	admin, ok := ctx.Value("admin").(models.Admin)
	if !ok {
		log.Println("Can't access context admin")
	}

	user_id := chi.URLParam(r, "userid")

	err := repo.deleteUser.DBDeleteUser(user_id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/userlist", http.StatusSeeOther)
}





// -------------------Product Side------------------------

func (repo *Repo) ProductList(w http.ResponseWriter, r *http.Request) {

	admin_name := chi.URLParam(r, "adminname")

	file := usecases.Logger()
	log.SetOutput(file)

	products, err := repo.products.DBGetAllProducts()
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	for key, value := range products {
		log.Println("Key: ", key, "Value: ", value)
	}
	log.Println("Product listing...")
	tpl, err := template.ParseFiles("./adminTemplates/ecommerce-product.html")
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	err = tpl.Execute(w, map[string]interface{}{
		"admin":   admin_name,
		"product": products,
	})
	if err != nil {
		log.Println("Template Error:", err)
	}
}

func (repo *Repo) AddProduct(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	// categories
	categories, err := repo.categories.DBGetAllCategories()
	if err != nil {
		log.Println(err)
	}

	// subcategories
	subcategories, err := repo.subcategories.DBGetAllSubcategories()
	if err != nil {
		log.Println(err)
	}

	// brands
	brands, err := repo.brands.DBGetAllBrands()
	if err != nil {
		log.Println(err)
	}

	tpl, err := template.ParseFiles("adminTemplates/admin_add_product.html")
	if err != nil {

	}
	tpl.Execute(w, map[string]interface{}{
		"admin_name":    admin.Admin_name,
		"categories":    categories,
		"subcategories": subcategories,
		"brands":        brands,
	})
}

func (repo *Repo) AddingProduct(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	r.ParseForm()
	var product models.ListProduct
	product.Product_name = r.FormValue("product_name")
	product.Product_inventory.Product_color = r.FormValue("product_color")
	product.Product_brand.Brand_id, _ = strconv.Atoi(r.FormValue("product_brand"))
	product.Product_category.Category_id, _ = strconv.Atoi(r.FormValue("product_category"))
	product.Product_subcategory.Subcategory_id, _ = strconv.Atoi(r.FormValue("product_subcategory"))
	product.Product_inventory.Product_quantity, _ = strconv.Atoi(r.FormValue("product_quantity"))
	product.Product_desc = r.FormValue("product_desc")
	product.Product_price, _ = decimal.NewFromString(r.FormValue("product_price"))

	log.Println("New Product:")

	log.Println("Name:",product.Product_name)
	log.Println("Color:",product.Product_inventory.Product_color)
	log.Println("Brand:",product.Product_brand.Brand_id)
	log.Println("Category:",product.Product_category.Category_id)
	log.Println("Subcategory:",product.Product_subcategory.Subcategory_id)
	log.Println("Quantity:",product.Product_inventory.Product_quantity)
	log.Println("Description:",product.Product_desc)
	log.Println("Price:",product.Product_price)

	err := repo.addproduct.DBAddProduct(product)
	if err != nil {
		log.Println("Error in DB")
		log.Println(err)
	}

	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/productlist", http.StatusSeeOther)
}

func (repo *Repo) ViewProduct(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	product_id := chi.URLParam(r, "product_id")
	product, err := repo.product.DBGetProduct(product_id)
	if err != nil {
		log.Println(err)
		return
	
	}
	colors, err := repo.productcolors.DBGetProductColors(product_id)
	if err != nil {
		log.Println(err)
	}

	log.Println("Upcoming Product: ", product)
	tpl, err := template.ParseFiles("./adminTemplates/ecommerce-product-single.html")
	if err != nil {
		// http.Error(w, http.StatusText(404), 404)
		log.Println(err)
		return
	}
	err = tpl.Execute(w, map[string]interface{}{
		"Product_id":    product.Product_id,
		"Product_name":  product.Product_name,
		"Product_price": product.Product_price,
		"Product_desc": product.Product_desc,
		"Product_category_id": product.Product_category.Category_id,
		"Product_category": product.Product_category.Category_name,
		"Product_subcategory_id": product.Product_subcategory.Subcategory_id,
		"Product_subcategory": product.Product_subcategory.Subcategory_name,
		"Product_brand_id": product.Product_brand.Brand_id,
		"Product_brand": product.Product_brand.Brand_name,
		"Product_inventory": product.Product_inventory.Inventory_id,
		"Product_quantity": product.Product_inventory.Product_quantity,
		"Product_color": product.Product_inventory.Product_color,
		// "Product_image": product.Product_image.Image,
		"Product_colors": colors,
	})
	if err != nil {
		log.Println(err)
	}
}

func (repo *Repo) EditProduct(w http.ResponseWriter, r *http.Request){

	product_id := chi.URLParam(r, "product_id")
	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)
	product, err := repo.product.DBGetProduct(product_id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	categories, err := repo.categories.DBGetAllCategories()
	if err != nil {
		log.Println(err)
	}
	subcategories, err := repo.subcategories.DBGetAllSubcategories()
	if err != nil {
		log.Println(err)		
	}
	brands, err := repo.brands.DBGetAllBrands()
	if err != nil {
		log.Println(err)
	}
	tpl, err := template.ParseFiles("./adminTemplates/admin_edit_product.html")
	if err != nil {	
		fmt.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}
	tpl.Execute(w, map[string]interface{}{
		"Product_id":    product.Product_id,
		"Product_name":  product.Product_name,
		"Product_price": product.Product_price,
		"Product_desc": product.Product_desc,
		"Product_category_id": product.Product_category.Category_id,
		"Product_category": product.Product_category.Category_name,
		"Product_subcategory_id": product.Product_subcategory.Subcategory_id,
		"Product_subcategory": product.Product_subcategory.Subcategory_name,
		"Product_brand_id": product.Product_brand.Brand_id,
		"Product_brand": product.Product_brand.Brand_name,
		"Product_inventory": product.Product_inventory.Inventory_id,
		"Product_quantity": product.Product_inventory.Product_quantity,
		"Product_color": product.Product_inventory.Product_color,
		// "Product_image": product.Product_image.Image,
		// "Product_colors": colors,
		"admin": admin.Admin_name,
		"Categories": categories,
		"Subcategories": subcategories,
		"Brands": brands,
	})
}

func (repo *Repo) EditingProduct(w http.ResponseWriter, r *http.Request){

	product_id := chi.URLParam(r, "product_id")

	r.ParseForm()
	var product models.ListProduct
	product.Product_id, _ = strconv.Atoi(product_id)
	product.Product_name = r.FormValue("product_name")
	product.Product_price, _ = decimal.NewFromString(r.FormValue("product_price"))

	product.Product_category.Category_id, _ = strconv.Atoi(r.FormValue("product_category"))
	product.Product_subcategory.Subcategory_id, _ = strconv.Atoi(r.FormValue("product_subcategory"))
	product.Product_brand.Brand_id, _ = strconv.Atoi(r.FormValue("product_brand"))

	product.Product_inventory.Product_color = r.FormValue("product_color")
	product.Product_inventory.Product_quantity, _ = strconv.Atoi(r.FormValue("product_quantity"))
	product.Product_desc = r.FormValue("product_desc")

	log.Println("Product in form:", product)
	err := repo.editproduct.DBEditProduct(product)
	if err != nil {
		log.Println(err)
		return
	}
	http.Redirect(w, r, "/admin/home/productlist/"+product_id+"/viewproduct", http.StatusSeeOther)
}







// ---------------------Category Management------------------------------

func (repo *Repo) CategoryList(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	category, err := repo.categories.DBGetAllCategories()
	if err != nil {
		log.Println(err)
	}
	tpl, err := template.ParseFiles("./adminTemplates/admin_categorylist.html")
	if err != nil {
		log.Println(err)
	}

	err = tpl.Execute(w, map[string]interface{}{
		"category": category,
		"admin":    admin.Admin_name,
	})
	if err != nil {
		log.Println(err)
	}
}

func (repo *Repo) AddCategory(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	tpl, err := template.ParseFiles("./adminTemplates/admin_add_category.html")
	if err != nil {
		log.Println(err)
	}
	err = tpl.Execute(w, map[string]interface{}{
		"admin": admin.Admin_name,
	})
	if err != nil {
		log.Println(err)
	}
}

func (repo *Repo) AddingCategory(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	r.ParseForm()
	var category models.Categories
	category.Category_name = r.FormValue("category_name")
	category.Category_desc = r.FormValue("category_desc")

	err := repo.addcategory.DBAddCategory(category)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/categorylist", http.StatusSeeOther)
}

func (repo *Repo) ViewCategoryProducts(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	category_id := chi.URLParam(r, "category_id")

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	products, err := repo.categoryproducts.DBGetCategoryProducts(category_id)
	if err != nil {
		log.Println(err)
	}

	tpl, err := template.ParseFiles("./adminTemplates/ecommerce-product.html")
	if err != nil {
		log.Println(err)
	}
	tpl.Execute(w, map[string]interface{}{
		"admin": admin.Admin_name,
		"product": products,
	})
}

func (repo *Repo) EditCategory(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	category_id := chi.URLParam(r, "category_id")
	category, err := repo.category.DBGetCategory(category_id)
	if err != nil {
		log.Println(err)
	}
	tpl, err := template.ParseFiles("./adminTemplates/admin_edit_category.html")
	if err != nil {
		log.Println(err)
	}
	tpl.Execute(w, map[string]interface{}{
		"admin": admin.Admin_name,
		"Category_id": category.Category_id,
		"Category_name": category.Category_name,
		"Category_desc": category.Category_desc,
	})
}

func (repo *Repo) EditingCategory(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	var category models.Categories
	category.Category_id, _ = strconv.Atoi(chi.URLParam(r, "category_id"))

	r.ParseForm()
	category.Category_name = r.FormValue("category_name")
	category.Category_desc = r.FormValue("category_desc")

	err := repo.editcategory.DBEditCategory(category)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/categorylist", http.StatusSeeOther)
}

func (repo *Repo) DeleteCategory(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	category_id := chi.URLParam(r, "category_id")
	err := repo.deletecategory.DBDeleteCategory(category_id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/categorylist", http.StatusSeeOther)
}






// --------------------Sub Category Management-----------------------

func (repo *Repo) SubcategoryList(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	subcategories, err := repo.subcategories.DBGetAllSubcategories()
	if err != nil {
		log.Println(err)
	}
	tpl, err := template.ParseFiles("./adminTemplates/admin_subcategorylist.html")
	if err != nil {
		log.Println(err)
	}
	err = tpl.Execute(w, map[string]interface{}{
		"admin":         admin.Admin_name,
		"subcategories": subcategories,
	})
	if err != nil {
		log.Println(err)
	}
}

func (repo *Repo) AddSubcategory(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	tpl, err := template.ParseFiles("./adminTemplates/admin_add_subcategory.html")
	if err != nil {
		log.Println(err)
	}
	err = tpl.Execute(w, map[string]interface{}{
		"admin": admin.Admin_name,
	})
	if err != nil {
		log.Println(err)
	}
}

func (repo *Repo) AddingSubcategory(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	log.Println("Entered into AddingSubcategory")
	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	r.ParseForm()
	var subcategory models.Subcategories
	subcategory.Subcategory_name = r.FormValue("subcategory_name")
	subcategory.Subcategory_desc = r.FormValue("subcategory_desc")

	err := repo.addsubcategory.DBAddSubcategory(subcategory)
	if err != nil {
		log.Println(err)
	}
	log.Println("Added to database")
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/subcategorylist", http.StatusSeeOther)
}

func (repo *Repo) ViewSubcategoryProducts(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	subcategory_id := chi.URLParam(r, "subcategory_id")

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	products, err := repo.subcategoryproducts.DBGetSubcategoryProducts(subcategory_id)
	if err != nil {
		log.Println(err)
	}

	tpl, err := template.ParseFiles("./adminTemplates/ecommerce-product.html")
	if err != nil {
		log.Println(err)
	}
	tpl.Execute(w, map[string]interface{}{
		"admin": admin.Admin_name,
		"product": products,
	})
}

func (repo *Repo) EditSubcategory(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	subcategory_id := chi.URLParam(r, "subcategory_id")
	subcategory, err := repo.subcategory.DBGetSubcategory(subcategory_id)
	if err != nil {
		log.Println(err)
	}
	tpl, err := template.ParseFiles("./adminTemplates/admin_edit_subcategory.html")
	if err != nil {
		log.Println(err)
	}
	tpl.Execute(w, map[string]interface{}{
		"admin": admin.Admin_name,
		"Subcategory_id": subcategory.Subcategory_id,
		"Subcategory_name": subcategory.Subcategory_name,
		"Subcategory_desc": subcategory.Subcategory_desc,
	})
}

func (repo *Repo) EditingSubcategory(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	var subcategory models.Subcategories
	subcategory.Subcategory_id, _ = strconv.Atoi(chi.URLParam(r, "subcategory_id"))

	r.ParseForm()
	subcategory.Subcategory_name = r.FormValue("subcategory_name")
	subcategory.Subcategory_desc = r.FormValue("subcategory_desc")

	err := repo.editsubcategory.DBEditSubcategory(subcategory)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/subcategorylist", http.StatusSeeOther)
}

func (repo *Repo) DeleteSubcategory(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	subcategory_id := chi.URLParam(r, "subcategory_id")
	err := repo.deletesubcategory.DBDeleteSubcategory(subcategory_id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/subcategorylist", http.StatusSeeOther)
}






// ---------------------Brand Management------------------------------

func (repo *Repo) BrandList(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	brands, err := repo.brands.DBGetAllBrands()
	if err != nil {
		log.Println(err)
	}
	tpl, err := template.ParseFiles("./adminTemplates/admin_brandlist.html")
	if err != nil {
		log.Println(err)
	}
	err = tpl.Execute(w, map[string]interface{}{
		"admin": admin.Admin_name,
		"brand": brands,
	})
	if err != nil {
		log.Println(err)
	}
}

func (repo *Repo) AddBrand(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)
	log.Println("In AddBrand..")

	tpl, err := template.ParseFiles("./adminTemplates/admin_add_brand.html")
	if err != nil {
		log.Println(err)
	}
	err = tpl.Execute(w, map[string]interface{}{
		"admin": admin.Admin_name,
	})
	if err != nil {
		log.Println(err)
	}
}

func (repo *Repo) AddingBrand(w http.ResponseWriter, r *http.Request) {

	file := usecases.Logger()
	log.SetOutput(file)

	log.Println("In AddingBrand..")
	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	r.ParseForm()
	var brand models.Brands
	brand.Brand_name = r.FormValue("brand_name")
	brand.Brand_desc = r.FormValue("brand_desc")

	log.Println("New Brand:", brand)
	err := repo.addbrand.DBAddBrand(brand)
	if err != nil {
		log.Println(err)
	}
	log.Println("New brand inserted")
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/brandlist", http.StatusSeeOther)
}

func (repo *Repo) ViewBrandProducts(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	brand_id := chi.URLParam(r, "brand_id")

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	products, err := repo.brandproducts.DBGetBrandProducts(brand_id)
	if err != nil {
		log.Println(err)
	}

	tpl, err := template.ParseFiles("./adminTemplates/ecommerce-product.html")
	if err != nil {
		log.Println(err)
	}
	tpl.Execute(w, map[string]interface{}{
		"admin": admin.Admin_name,
		"product": products,
	})
}

func (repo *Repo) EditBrand(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	brand_id := chi.URLParam(r, "brand_id")
	brand, err := repo.brand.DBGetBrand(brand_id)
	if err != nil {
		log.Println(err)
	}
	tpl, err := template.ParseFiles("./adminTemplates/admin_edit_brand.html")
	if err != nil {
		log.Println(err)
	}
	tpl.Execute(w, map[string]interface{}{
		"admin": admin.Admin_name,
		"Brand_id": brand.Brand_id,
		"Brand_name": brand.Brand_name,
		"Brand_desc": brand.Brand_desc,
	})
}

func (repo *Repo) EditingBrand(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	var brand models.Brands
	brand.Brand_id, _ = strconv.Atoi(chi.URLParam(r, "brand_id"))

	r.ParseForm()
	brand.Brand_name = r.FormValue("brand_name")
	brand.Brand_desc = r.FormValue("brand_desc")

	err := repo.editbrand.DBEditBrand(brand)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/brandlist", http.StatusSeeOther)
}

func (repo *Repo) DeleteBrand(w http.ResponseWriter, r *http.Request){

	file := usecases.Logger()
	log.SetOutput(file)

	ctx := r.Context()
	admin := ctx.Value("admin").(models.Admin)

	brand_id := chi.URLParam(r, "brand_id")
	err := repo.deletebrand.DBDeleteBrand(brand_id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/admin/home/"+admin.Admin_name+"/brandlist", http.StatusSeeOther)
}

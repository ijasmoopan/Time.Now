package routes

import "github.com/go-chi/chi/v5"

func Router() *chi.Mux {

	r := chi.NewRouter()

	// Admin side routes
	r.Get("/admin", AdminLoginPage)
	r.Post("/adminvalidating", AdminLoginValidating)

	r.Get("/admin/home", AdminHome)

	
	
	// User side routes
	r.Get("/login", UserLoginPage)
	r.Post("/loginvalidating", UserLoginValidating)

	r.Get("/home", HomePage)

	return r 
}
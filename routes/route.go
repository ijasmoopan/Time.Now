package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Router function for handling routes.
func Router() *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	
	adminRouter(router)
	userRouter(router)

	return router
}

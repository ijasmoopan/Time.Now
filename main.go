package main

import (
	"fmt"
	"net/http"

	"github.com/ijasmoopan/Time.Now/routes"
)

func main() {

	addr := ":8080"
	
	fmt.Println("Server starting at port 8080...")

	router := routes.Router()

	http.ListenAndServe(addr, router)

}

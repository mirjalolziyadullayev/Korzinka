package main

import (
	"Store/handlers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/categories", handlers.CategoriesHandler)	
	http.HandleFunc("/products", handlers.ProductsHandler)
	http.HandleFunc("/discount", handlers.GetDiscountedProducts)
	// http.HandleFunc("/discount", handlers.DiscountHandler)

	fmt.Println("Server working on port :8080 ...")
	http.ListenAndServe(":8080", nil)
}
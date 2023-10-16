package handlers

import (
	"Store/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAll(w,r)
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	var categoriesData []models.Category
	categoriesByte, _ := os.ReadFile("db/categories.json")
	json.Unmarshal(categoriesByte, &categoriesData)

	var productsData []models.Product
	productsByte, _ := os.ReadFile("db/products.json")
	json.Unmarshal(productsByte, &productsData)

	fmt.Fprint(w, "Products", productsData, "Categories",categoriesData)
}
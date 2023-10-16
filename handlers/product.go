package handlers

import (
	"Store/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllProducts(w,r)
	case "POST":
		createProduct(w,r)
	case "PUT":
		updateProduct(w,r)
	case "DELETE":
		deleteProduct(w,r)
	}
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	var productsData []models.Product
	byteData, _ := os.ReadFile("db/products.json")
	json.Unmarshal(byteData, &productsData)

	var discount models.Discount
	for i := 0; i < len(productsData); i++ {
		if productsData[i].Id == discount.ProductID {
			productsData[i].Price = (discount.Procent - 100) / 100 * productsData[i].Price
		}
	}

	json.NewEncoder(w).Encode(productsData)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	//parsing request
	var newProduct models.Product
	json.NewDecoder(r.Body).Decode(&newProduct)

	//parsing db file
	var productsData []models.Product
	byteData, _ := os.ReadFile("db/products.json")
	json.Unmarshal(byteData, &productsData)

	newProduct.Id = len(productsData)+1
	newProduct.CreatedAt = time.Now()
	productsData = append(productsData, newProduct)

	res, _ := json.Marshal(productsData)
	os.WriteFile("db/products.json",res,0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Created!")
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	var updateProduct models.Product
	json.NewDecoder(r.Body).Decode(&updateProduct)

	var productsData [] models.Product
	byteData, _ := os.ReadFile("db/products.json")
	json.Unmarshal(byteData, &productsData)

	for i := 0; i < len(productsData); i++ {
		if updateProduct.Id == productsData[i].Id {
			productsData[i].Name = updateProduct.Name
			productsData[i].ExpireDate = updateProduct.ExpireDate
			productsData[i].UpdatedAt = time.Now()
			productsData[i].Price = updateProduct.Price
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "no product with such an ID")
		}
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	//reading response
	var deleteProduct models.Product
	json.NewDecoder(r.Body).Decode(&deleteProduct)

	//parsing db file
	var productsData []models.Product
	byteData, _ := os.ReadFile("db/products.json")
	json.Unmarshal(byteData, &productsData)
	
	for i := 0; i < len(productsData); i++ {
		if deleteProduct.Id == productsData[i].Id {
			productsData = append(productsData[:i],productsData[i+1:]... )
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "no product with such an ID")
			return
		}
	}

	// writing to file
	res, _ := json.Marshal(productsData)
	os.WriteFile("db/products.json",res,0)

	//sending request
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Deleted!")
}
package handlers

import (
	"Store/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllCategories(w,r)
	case "POST":
		createCategory(w,r)
	case "DELETE":
		deleteCategory(w,r)
	}
}

func getAllCategories(w http.ResponseWriter, r *http.Request) {
	//parsing
	var categoriesData []models.Category
	byteData, _ := os.ReadFile("db/categories.json")
	json.Unmarshal(byteData, &categoriesData)

	json.NewEncoder(w).Encode(categoriesData)
}
func createCategory(w http.ResponseWriter, r *http.Request) {
	//request body parsing
	var newCategory models.Category
	json.NewDecoder(r.Body).Decode(&newCategory)


	//parsing db file
	var categoriesData []models.Category
	byteData, _ := os.ReadFile("db/categories.json")
	json.Unmarshal(byteData, &categoriesData)

	newCategory.Id = len(categoriesData)+1
	categoriesData = append(categoriesData, newCategory)

	//writing to file
	res, _ := json.Marshal(categoriesData)
	os.WriteFile("db/categories.json", res, 0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Created!")
	json.NewEncoder(w).Encode(newCategory)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	//request body parsing
	var deleteCategory models.Category
	json.NewDecoder(r.Body).Decode(&deleteCategory)


	//parsing db file
	var categoriesData []models.Category
	byteData, _ := os.ReadFile("db/categories.json")
	json.Unmarshal(byteData, &categoriesData)

	for i := 0; i < len(categoriesData); i++ {
		if deleteCategory.Id == categoriesData[i].Id {
			categoriesData = append(categoriesData[:i],categoriesData[i+1:]... )
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "no category with such an ID")
			return
		}
	}

	// writing to file
	res, _ := json.Marshal(deleteCategory)
	os.WriteFile("db/categories.json", res, 0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Deleted!")
}
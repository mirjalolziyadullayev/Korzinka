package handlers

import (
	"Store/models"
	"encoding/json"
	"net/http"
)

func GetDiscountedProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDiscounted(w, r)
	case "POST":
		addDiscounted(w, r)
	}
}

func getDiscounted(w http.ResponseWriter, r *http.Request) {

}
func addDiscounted(w http.ResponseWriter, r *http.Request) {
	var newDiscount models.Discount
	json.NewDecoder(r.Body).Decode(&newDiscount)

	
}
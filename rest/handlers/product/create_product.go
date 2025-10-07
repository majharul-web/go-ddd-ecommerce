package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}
	// Assign a new ID to the product
	newProduct.ID = len(database.GetAllProducts()) + 1

	created := database.StoreProduct(newProduct)
	util.SendData(w, created, 201)
}

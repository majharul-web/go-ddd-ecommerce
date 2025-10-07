package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// get product ID from URL
	productID := r.PathValue("productId")
	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}
	// Decode the request body into a Product struct
	var updatedProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	err = database.UpdateProduct(id, updatedProduct)
	if err != nil {
		util.SendError(w, "Product not found", 404)
		return
	}

	util.SendData(w, "Successfully updated product", 200)
}

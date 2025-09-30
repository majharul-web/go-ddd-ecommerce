package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)


func GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here
	productID := r.PathValue("productId")
	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			util.SendData(w, product, 200)
			return
		}	
	}

	util.SendData(w, "Product not found", 404)
}
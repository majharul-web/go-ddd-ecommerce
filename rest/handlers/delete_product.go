package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")
	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	err = database.DeleteProduct(id)
	if err != nil {
		util.SendError(w, "Product not found", 404)
		return
	}

	util.SendData(w, map[string]string{"message": "Product deleted successfully"}, 200)
}

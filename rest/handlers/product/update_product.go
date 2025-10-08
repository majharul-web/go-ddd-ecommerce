package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type RequestUpdatedProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// get product ID from URL
	productID := r.PathValue("productId")
	id, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	// Decode the request body into a Product struct
	var updatedProduct RequestUpdatedProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedProduct)
	if err != nil {
		util.SendError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updated, err := h.productRepo.Update(id, repo.Product{
		ID:          id,
		Title:       updatedProduct.Title,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		ImgUrl:      updatedProduct.ImgUrl,
	})
	if err != nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, updated, http.StatusOK)
}

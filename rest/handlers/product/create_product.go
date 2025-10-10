package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type createRequestProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct createRequestProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	created, err := h.svc.Create(domain.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgUrl:      newProduct.ImgUrl,
	})
	if err != nil {
		util.SendError(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	util.SendData(w, created, http.StatusCreated)
}

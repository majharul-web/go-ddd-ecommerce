package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProductList(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.List()
	if err != nil {
		util.SendError(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}
	util.SendData(w, products, http.StatusOK)
}

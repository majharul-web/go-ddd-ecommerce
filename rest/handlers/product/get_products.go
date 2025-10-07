package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProductList(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, database.GetAllProducts(), 200)
}

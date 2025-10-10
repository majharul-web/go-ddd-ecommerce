package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductList(w http.ResponseWriter, r *http.Request) {
	// get page and limit from query params
	// if not provided, set default values
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}

	// convert page and limit to integers
	// if conversion fails, set default values
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		limitInt = 10
	}
	products, err := h.svc.List(pageInt, limitInt)
	if err != nil {
		util.SendError(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	count, err := h.svc.Count()
	if err != nil {
		util.SendError(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	util.SendPaginatedData(w, products, pageInt, limitInt, count)
}

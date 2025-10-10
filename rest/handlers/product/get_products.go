package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
	"sync"
)

var cnt int

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

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		cnt1, err := h.svc.Count()
		if err != nil {
			util.SendError(w, "Failed to retrieve products", http.StatusInternalServerError)
			return
		}
		cnt = cnt1

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cnt2, err := h.svc.Count()
		if err != nil {
			util.SendError(w, "Failed to retrieve products", http.StatusInternalServerError)
			return
		}
		cnt = cnt2

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cnt3, err := h.svc.Count()
		if err != nil {
			util.SendError(w, "Failed to retrieve products", http.StatusInternalServerError)
			return
		}
		cnt = cnt3

	}()

	wg.Wait()

	// time.Sleep(8 * time.Second)

	util.SendPaginatedData(w, products, pageInt, limitInt, cnt)
}

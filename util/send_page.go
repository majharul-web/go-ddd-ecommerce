package util

import (
	"net/http"
)

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page         int `json:"page"`
	Limit        int `json:"limit"`
	TotalRecords int `json:"totalRecords"`
	TotalPages   int `json:"totalPages"`
}

func SendPaginatedData(w http.ResponseWriter, data any, page int, limit int, count int) {

	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:         page,
			Limit:        limit,
			TotalRecords: count,
			TotalPages:   (count + limit - 1) / limit, // ceiling division
		},
	}

	SendData(w, paginatedData, http.StatusOK)
}

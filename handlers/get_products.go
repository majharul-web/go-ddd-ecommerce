package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetProductList(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, database.ProductList, 200)
}
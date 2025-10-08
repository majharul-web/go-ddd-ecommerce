package user	

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type createRequestUser struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	IsShopOwner  bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser createRequestUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	created, err := h.userRepo.Create(repo.User{
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:      newUser.Email,
		Password:   newUser.Password,
		Avatar:     newUser.Avatar,
		IsShopOwner: newUser.IsShopOwner,
	})
	if err != nil {
		util.SendError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	util.SendData(w, created, http.StatusCreated)
}

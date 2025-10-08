package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type updateRequestUser struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Avatar       string `json:"avatar"`
	IsShopOwner  bool   `json:"is_shop_owner"`
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// get user ID from URL
	userID := r.PathValue("userId")
	id, err := strconv.Atoi(userID)
	if err != nil {
		util.SendError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	// Decode the request body into a user struct
	var updatedUser updateRequestUser
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedUser)
	if err != nil {
		util.SendError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updated, err := h.userRepo.Update(id, repo.User{
		ID:          id,
		FirstName:   updatedUser.FirstName,
		LastName:    updatedUser.LastName,
		Avatar:      updatedUser.Avatar,
		IsShopOwner: updatedUser.IsShopOwner,
	})
	if err != nil {
		util.SendError(w, "user not found", http.StatusNotFound)
		return
	}

	util.SendData(w, updated, http.StatusOK)
}

package user

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Implementation will go here
	userID := r.PathValue("userId")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.Get(id)
	if err != nil {
		util.SendError(w, "User not found", http.StatusNotFound)
		return
	}

	util.SendData(w, user, http.StatusOK)
}

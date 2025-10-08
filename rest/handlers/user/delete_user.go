package user

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.userRepo.Delete(id)
	if err != nil {
		util.SendError(w, "User not found", http.StatusNotFound)
		return
	}

	util.SendData(w, "Successfully deleted user", http.StatusOK)
}

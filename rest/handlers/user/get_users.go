package user

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetUserList(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepo.List()
	if err != nil {
		util.SendError(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	util.SendData(w, users, http.StatusOK)
}

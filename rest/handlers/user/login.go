package user

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"

	"ecommerce/config"
)

type loginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginUser handles user login and returns a JWT
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload loginPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetByEmail(payload.Email)
	if err != nil || user.Password != payload.Password {
		util.SendError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	conf := config.GetConfig()

	token, err := util.CreateJWT(conf.JWTSecret, util.Payload{
		Sub:         user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})
	if err != nil {
		util.SendError(w, "Failed to create JWT", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]string{"token": token}, http.StatusOK)
}

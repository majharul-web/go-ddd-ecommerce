package handlers

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

// 🟢 Get all users
func GetUserList(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.GetAllUsers(), 200)
}

// 🟢 Get user by ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", 400)
		return
	}

	user, err := database.GetUserByID(id)
	if err != nil {
		util.SendError(w, "User not found", 404)
		return
	}

	util.SendData(w, user, 200)
}

// 🟢 Delete user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", 400)
		return
	}

	err = database.DeleteUser(id)
	if err != nil {
		util.SendError(w, "User not found", 404)
		return
	}

	util.SendData(w, map[string]string{"message": "User deleted successfully"}, 200)
}

// 🟢 Update user by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", 400)
		return
	}

	var updatedUser database.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	err = database.UpdateUser(id, updatedUser)
	if err != nil {
		util.SendError(w, "User not found", 404)
		return
	}

	util.SendData(w, "Successfully updated user", 200)
}

// 🟢 Create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	// Assign a new ID
	newUser.ID = len(database.GetAllUsers()) + 1

	created := database.StoreUser(newUser)
	util.SendData(w, created, 201)
}

// LoginUser handles login
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", 400)
		return
	}

	user, err := database.GetUserByEmail(payload.Email)
	if err != nil || user.Password != payload.Password {
		util.SendError(w, "Invalid email or password", 401)
		return
	}

	conf := config.GetConfig()

	token, err := util.CreateJWT(conf.JWTSecret, util.Payload{
		Sub:         user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOOwner,
	})

	if err != nil {
		util.SendError(w, "Failed to create JWT", 500)
		return
	}

	// Normally, generate a JWT here. For simplicity, just return the user
	util.SendData(w, token, 200)
}

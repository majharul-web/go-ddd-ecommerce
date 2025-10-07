package database

import "fmt"

var userList []User

type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	IsShopOOwner bool   `json:"is_shop_owner"`
}

// StoreUser adds a new user to the list
func StoreUser(u User) User {
	userList = append(userList, u)
	return u
}

// GetAllUsers returns all users
func GetAllUsers() []User {
	return userList
}

// GetUserByID finds a user by ID
func GetUserByID(id int) (*User, error) {
	for _, u := range userList {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

// GetUserByEmail returns a user by email
func GetUserByEmail(email string) (*User, error) {
	for _, u := range userList {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("user with email %s not found", email)
}

// UpdateUser updates a user's information by ID
func UpdateUser(id int, updated User) error {
	for i, u := range userList {
		if u.ID == id {
			updated.ID = id
			userList[i] = updated
			return nil
		}
	}
	return fmt.Errorf("user with ID %d not found", id)
}

// DeleteUser removes a user by ID
func DeleteUser(id int) error {
	for i, u := range userList {
		if u.ID == id {
			userList = append(userList[:i], userList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with ID %d not found", id)
}

func init() {
	fmt.Println("Initializing server...")
	// Initialize some sample users
	userList = []User{
		{ID: 1, FirstName: "Majharul", LastName: "Islam", Email: "majharul@example.com", Password: "12345", Avatar: "http://example.com/avatar1.jpg", IsShopOOwner: true},
	}
}

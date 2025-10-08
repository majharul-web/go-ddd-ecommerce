package repo

import (
	"fmt"
)

// User struct
type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	IsShopOwner  bool   `json:"is_shop_owner"`
}

// UserRepo interface
type UserRepo interface {
	Create(u User) (*User, error)
	Update(id int, u User) (*User, error)
	Delete(id int) error
	Get(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	List() ([]*User, error)
}

// userRepo struct
type userRepo struct {
	userList []*User
}

// NewUserRepo creates a repo and adds initial users
func NewUserRepo() UserRepo {
	return &userRepo{}
}

// Create adds a new user
func (r *userRepo) Create(u User) (*User, error) {
	u.ID = len(r.userList) + 1
	r.userList = append(r.userList, &u)
	return &u, nil
}

// Update modifies an existing user by ID
func (r *userRepo) Update(id int, updated User) (*User, error) {
	for i, u := range r.userList {
		if u.ID == id {
			updated.ID = id
			r.userList[i] = &updated
			return &updated, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

// Delete removes a user by ID
func (r *userRepo) Delete(id int) error {
	var tempList []*User
	for _, u := range r.userList {
		if u.ID != id {
			tempList = append(tempList, u)
		}
	}
	if len(tempList) == len(r.userList) {
		return fmt.Errorf("user with ID %d not found", id)
	}
	r.userList = tempList
	return nil
}

// Get retrieves a user by ID
func (r *userRepo) Get(id int) (*User, error) {
	for _, u := range r.userList {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

// GetByEmail retrieves a user by email
func (r *userRepo) GetByEmail(email string) (*User, error) {
	for _, u := range r.userList {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, fmt.Errorf("user with email %s not found", email)
}

// List returns all users as pointers
func (r *userRepo) List() ([]*User, error) {
	return r.userList, nil
}

// generateInitialUsers adds some initial users to the repo
func (r *userRepo) generateInitialUsers() {
	r.userList = []*User{
		{
			ID:          1,
			FirstName:   "Alice",
			LastName:    "Smith",
			Email:       "alice@example.com",
			Password:    "password123",
			Avatar:      "alice.png",
			IsShopOwner: false,
		},
		{
			ID:          2,
			FirstName:   "Bob",
			LastName:    "Johnson",
			Email:       "bob@example.com",
			Password:    "password456",
			Avatar:      "bob.png",
			IsShopOwner: true,
		},
	}
}

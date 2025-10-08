package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// User struct
type User struct {
	ID          int    `db:"id" json:"id"`
	FirstName   string `db:"first_name" json:"first_name"`
	LastName    string `db:"last_name" json:"last_name"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	Avatar      string `db:"avatar" json:"avatar"`
	IsShopOwner bool   `db:"is_shop_owner" json:"is_shop_owner"`
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
	db *sqlx.DB
}

// NewUserRepo creates a new repo instance
func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

// Create adds a new user
func (r *userRepo) Create(u User) (*User, error) {
	query := `
		INSERT INTO users (first_name, last_name, email, password, avatar, is_shop_owner)
		VALUES (:first_name, :last_name, :email, :password, :avatar, :is_shop_owner)
		RETURNING id;
	`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var userID int
	if err := stmt.Get(&userID, u); err != nil {
		return nil, err
	}
	u.ID = userID
	return &u, nil
}

// Update modifies an existing user by ID
func (r *userRepo) Update(id int, u User) (*User, error) {
	query := `
		UPDATE users
		SET first_name = :first_name,
		    last_name = :last_name,
		    email = :email,
		    password = :password,
		    avatar = :avatar,
		    is_shop_owner = :is_shop_owner
		WHERE id = :id
		RETURNING *;
	`

	u.ID = id
	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var updated User
	if err := stmt.Get(&updated, u); err != nil {
		return nil, err
	}

	return &updated, nil
}

// Delete removes a user by ID
func (r *userRepo) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}

	return nil
}

// Get retrieves a user by ID
func (r *userRepo) Get(id int) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE id = $1`
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepo) GetByEmail(email string) (*User, error) {
	var user User
	query := `SELECT * FROM users WHERE email = $1`
	if err := r.db.Get(&user, query, email); err != nil {
		return nil, err
	}
	return &user, nil
}

// List returns all users
func (r *userRepo) List() ([]*User, error) {
	var users []*User
	query := `SELECT * FROM users ORDER BY id`
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}
	return users, nil
}

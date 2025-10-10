package repo

import (
	"ecommerce/domain"
	"ecommerce/user"
	"fmt"

	"github.com/jmoiron/sqlx"
)



type UserRepo interface {
	user.UserRepo
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
func (r *userRepo) Create(u domain.User) (*domain.User, error) {
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
func (r *userRepo) Update(id int, u domain.User) (*domain.User, error) {
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

	var updated domain.User
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
func (r *userRepo) Get(id int) (*domain.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE id = $1`
	if err := r.db.Get(&user, query, id); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepo) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT * FROM users WHERE email = $1`
	if err := r.db.Get(&user, query, email); err != nil {
		return nil, err
	}
	return &user, nil
}

// List returns all users
func (r *userRepo) List() ([]*domain.User, error) {
	var users []*domain.User
	query := `SELECT * FROM users ORDER BY id`
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}
	return users, nil
}

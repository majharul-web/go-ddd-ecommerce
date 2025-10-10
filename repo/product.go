package repo

import (
	"ecommerce/domain"
	"ecommerce/product"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

// productRepo struct
type productRepo struct {
	db *sqlx.DB
}

// NewProductRepo creates a repo instance
func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{db: db}
}

// Create adds a new product to PostgreSQL
func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
		INSERT INTO products (title, description, price, img_url)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	var productID int
	err := r.db.QueryRowx(query,
		p.Title, p.Description, p.Price, p.ImgUrl,
	).Scan(&productID)
	if err != nil {
		return nil, err
	}

	p.ID = productID
	return &p, nil
}

// Update modifies an existing product
func (r *productRepo) Update(id int, p domain.Product) (*domain.Product, error) {
	query := `
		UPDATE products
		SET title = $1, description = $2, price = $3, img_url = $4
		WHERE id = $5
		RETURNING id, title, description, price, img_url;
	`

	var updated domain.Product
	err := r.db.Get(&updated, query,
		p.Title, p.Description, p.Price, p.ImgUrl, id,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update product with ID %d: %v", id, err)
	}
	return &updated, nil
}

// Delete removes a product by ID
func (r *productRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1;`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("product with ID %d not found", id)
	}
	return nil
}

// Get retrieves a product by ID
func (r *productRepo) Get(id int) (*domain.Product, error) {
	query := `SELECT id, title, description, price, img_url FROM products WHERE id = $1;`

	var p domain.Product
	err := r.db.Get(&p, query, id)
	if err != nil {
		return nil, fmt.Errorf("product with ID %d not found: %v", id, err)
	}
	return &p, nil
}

// List returns all products
func (r *productRepo) List(page int, limit int) ([]*domain.Product, error) {
	offset := ((page - 1) * limit) + 1
	if offset < 0 {
		offset = 0
	}
	query := `SELECT id, title, description, price, img_url FROM products LIMIT $1 OFFSET $2;`

	var products []*domain.Product
	err := r.db.Select(&products, query, limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) Count() (int, error) {
	query := `SELECT COUNT(*) FROM products;`

	var count int
	err := r.db.Get(&count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}

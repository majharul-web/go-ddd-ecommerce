package domain

type User struct {
	ID          int    `db:"id" json:"id"`
	FirstName   string `db:"first_name" json:"first_name"`
	LastName    string `db:"last_name" json:"last_name"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	Avatar      string `db:"avatar" json:"avatar"`
	IsShopOwner bool   `db:"is_shop_owner" json:"is_shop_owner"`
}
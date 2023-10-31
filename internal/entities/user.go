package entities

type User struct {
	Id       int    `json:"-" db:"id"`
	IsAdmin  bool   `json:"is_admin"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Balance  int    `json:"balance" db:"balance"`
}

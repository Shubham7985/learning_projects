package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"-"`
}

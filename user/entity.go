package user

import "github.com/go-playground/validator/v10"

// User ...
type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty" validate:"required, email"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty" validate:"required, min=3"`
	IsAdmin  bool   `json:"is_admin,omitempty" db:"is_admin"`
}

// Validate ...
func (u *User) Validate() error {
	return validator.New().Struct(u)
}

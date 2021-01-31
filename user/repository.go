package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// Repository ...
type Repository struct {
	db sqlx.DB
}

const insertUserSQL = "INSERT INTO user (id, name, email, password, is_admin) VALUES (:id, :name, :email, :password, :is_admin)"
const getUserDataSQL = "SELECT id, name, email, is_admin FROM user WHERE id = ?"
const updateUserDataSQL = "UPDATE user SET %s WHERE id =:id"
const deleteUserSQL = "DELETE FROM user where id = ?"

func hashAndSaltPwd(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd), err
}

func (r *Repository) updateUser(user User, fields ...string) error {
	var updateFields = make([]string, len(fields))
	for i, f := range fields {
		updateFields[i] = fmt.Sprintf("%s=:%s", f, f)
	}

	qry := fmt.Sprintf(updateUserDataSQL, strings.Join(updateFields, ", "))
	_, err := r.db.NamedExec(qry, user)
	return err
}

// CreateUser inserts given user to db with salted password and returns same user with id
func (r *Repository) CreateUser(user User) (*User, error) {
	user.ID = uuid.New().String()
	var err error
	user.Password, err = hashAndSaltPwd(user.Password)
	if err != nil {
		return nil, err
	}

	if _, err = r.db.NamedExec(insertUserSQL, user); err != nil {
		return nil, errors.New("error while creating user")
	}

	user.Password = ""
	return &user, nil
}

// Get ...
func (r *Repository) Get(id string) (User, error) {
	user := User{}
	err := r.db.Get(&user, getUserDataSQL, id)
	return user, err
}

// Update ...
func (r *Repository) Update(user User) error {
	return r.updateUser(user, "name", "email")
}

// Delete ...
func (r *Repository) Delete(id string) error {
	_, err := r.db.Exec(deleteUserSQL, id)
	return err
}

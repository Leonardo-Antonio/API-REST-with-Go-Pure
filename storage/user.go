package storage

import (
	"database/sql"

	"github.com/Leonardo-Antonio/api-rest-mysql/storage/validations"

	"github.com/Leonardo-Antonio/api-rest-mysql/model"
)

const (
	userCreate = "INSERT INTO tb_users VALUES(null, ?, ?)"
	userDelete = "DELETE FROM tb_users WHERE id = ?"
	userUpdate = "UPDATE tb_users SET email = ?, pass = ? WHERE id = ?"
	userLogin  = "SELECT email, pass FROM tb_users WHERE email = ? and pass = ?"
)

// User .
type User struct {
	db *sql.DB
}

// NewUser -> method constructor
func NewUser(db *sql.DB) *User {
	return &User{db}
}

// Create new user
func (u *User) Create(user model.User) error {
	stmt, err := u.db.Prepare(userCreate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(
		user.Email,
		user.Password,
	)
	if err != nil {
		return err
	}

	if rA, _ := rs.RowsAffected(); rA != 1 {
		return validations.ErrorRowsAffected
	}

	return nil
}

// Update user
func (u *User) Update(id int, user model.User) error {
	stmt, err := u.db.Prepare(userUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(
		user.Email,
		user.Password,
		id,
	)

	if err != nil {
		return err
	}
	if rA, _ := rs.RowsAffected(); rA != 1 {
		return validations.ErrorRowsAffected
	}

	return nil
}

// Delete user
func (u *User) Delete(id int) error {
	stmt, err := u.db.Prepare(userDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	if rA, _ := rs.RowsAffected(); rA != 1 {
		return validations.ErrorRowsAffected
	}

	return nil
}

// Login user
func (u *User) Login(user model.User) error {
	stmt, err := u.db.Prepare(userLogin)
	if err != nil {
		return err
	}
	defer stmt.Close()

	data := model.User{}
	err = stmt.QueryRow(user.Email, user.Password).Scan(
		&data.Email,
		&data.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

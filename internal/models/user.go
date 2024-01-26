package models

import (
	"arcade/internal/database"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Name     string
	Username string
	Password string
}

func FetchUser(id string) (User, error) {
	var user User
	if err := database.DB.Get(&user, `SELECT * FROM user WHERE id = ?`, id); err != nil {
		return User{}, err
	}
	return user, nil
}

func FetchUserByUsername(username string) (User, error) {
	var user User
	if err := database.DB.Get(&user, `SELECT * FROM user WHERE username = ?`, username); err != nil {
		return User{}, err
	}
	return user, nil
}

func FetchUsers() ([]User, error) {
	var users []User
	if err := database.DB.Select(&users, `SELECT * FROM user`); err != nil {
		return []User{}, err
	}
	return users, nil
}

func CreateUser(t *User) error {
	if _, err := database.DB.Exec(
		`INSERT INTO user VALUES (?, ?, ?, ?)`,
		t.Id,
		t.Name,
		t.Username,
		t.Password,
	); err != nil {
		return err
	}
	return nil
}

func UpdateUser(t *User) error {
	if _, err := database.DB.Exec(
		`UPDATE user
         SET name = ?, username = ?, password = ? 
         WHERE id = ?`,
		t.Name,
		t.Username,
		t.Password,
		t.Id,
	); err != nil {
		return err
	}

	return nil
}

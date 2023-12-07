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

func FetchUser(id uuid.UUID) (User, error) {
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

func DeleteUser(id uuid.UUID) error {
	if _, err := database.DB.Exec(
		`DELETE user WHERE id = ?`, id,
	); err != nil {
		return nil
	}
	return nil
}

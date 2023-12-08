package models

import (
	"arcade/internal/database"
	"log"

	"github.com/google/uuid"
)

type Template struct {
	Id         uuid.UUID
	User       uuid.UUID
	Categories string
}

func FetchTemplate(id uuid.UUID) (Template, error) {
	var template Template
	if err := database.DB.Get(&template, `SELECT * FROM template WHERE id = ?`, id); err != nil {
		return Template{}, err
	}
	return template, nil
}

func FetchTemplatesByUser(user_id string) ([]Template, error) {
	var templates []Template
	if err := database.DB.Select(&templates, `SELECT * FROM template WHERE user = ?`, user_id); err != nil {
        log.Println("Database error: ", err)
		return []Template{}, err
	}
	return templates, nil
}

func CreateTemplate(t *Template) error {
	if _, err := database.DB.Exec(
		`INSERT INTO template VALUES (?, ?, ?)`,
		t.Id,
		t.User,
		t.Categories,
	); err != nil {
		return err
	}
	return nil
}

func DeleteTemplate(id uuid.UUID) error {
	if _, err := database.DB.Exec(
		`DELETE template WHERE id = ?`, id,
	); err != nil {
		return nil
	}
	return nil
}

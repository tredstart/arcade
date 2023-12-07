package models

import (
	"arcade/internal/database"

	"github.com/google/uuid"
)

type Tempate struct {
	Id         uuid.UUID
	User       uuid.UUID
	Categories string
}

func FetchTempate(id uuid.UUID) (Tempate, error) {
	var template Tempate
	if err := database.DB.Get(&template, `SELECT * FROM template WHERE id = ?`, id); err != nil {
		return Tempate{}, err
	}
	return template, nil
}

func FetchTempates() ([]Tempate, error) {
	var templates []Tempate
	if err := database.DB.Select(&templates, `SELECT * FROM template`); err != nil {
		return []Tempate{}, err
	}
	return templates, nil
}

func CreateTempate(t *Tempate) error {
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

func DeleteTempate(id uuid.UUID) error {
	if _, err := database.DB.Exec(
		`DELETE template WHERE id = ?`, id,
	); err != nil {
		return nil
	}
	return nil
}

package models

import (
	"arcade/internal/database"

	"github.com/google/uuid"
)

type Retro struct {
	Id       uuid.UUID
	User     uuid.UUID
	Created  string
	Template uuid.UUID
	Visible  bool
}

func FetchRetro(id uuid.UUID) (Retro, error) {
	var retro Retro
	if err := database.DB.Get(&retro, `SELECT * FROM retro WHERE id = ?`, id); err != nil {
		return Retro{}, err
	}
	return retro, nil
}

func FetchLatestRetro(user_id uuid.UUID) (Retro, error) {
	var retro Retro
	if err := database.DB.Get(&retro, `SELECT * FROM retro WHERE user = ? ORDER BY created DESC LIMIT 1`, user_id); err != nil {
		return Retro{}, err
	}
	return retro, nil
}

func FetchRetrosByUser(user_id string) ([]Retro, error) {
	var retros []Retro
	if err := database.DB.Select(&retros, `SELECT * FROM retro WHERE user = ?`, user_id); err != nil {
		return []Retro{}, err
	}
	return retros, nil
}

func CreateRetro(t *Retro) error {
	if _, err := database.DB.Exec(
		`INSERT INTO retro VALUES (?, ?, ?, ?, ?)`,
		t.Id,
		t.User,
		t.Created,
		t.Template,
		t.Visible,
	); err != nil {
		return err
	}
	return nil
}

func RetroSetVisibility(id string, visible bool) error {
	if _, err := database.DB.Exec(
		`UPDATE retro
         SET visible = ?
         WHERE id = ?`,
		visible, id,
	); err != nil {
		return err
	}
	return nil
}


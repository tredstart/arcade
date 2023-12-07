package models

import (
	"arcade/internal/database"

	"github.com/google/uuid"
)

type Record struct {
	Id       uuid.UUID
	User     uuid.UUID
	Retro    uuid.UUID
	Author   string
	Category string
	Content  string
}

func FetchRecord(id uuid.UUID) (Record, error) {
	var record Record
	if err := database.DB.Get(&record, `SELECT * FROM record WHERE id = ?`, id); err != nil {
		return Record{}, err
	}
	return record, nil
}

func FetchRecords() ([]Record, error) {
	var records []Record
	if err := database.DB.Select(&records, `SELECT * FROM record`); err != nil {
		return []Record{}, err
	}
	return records, nil
}

func CreateRecord(t *Record) error {
	if _, err := database.DB.Exec(
		`INSERT INTO record VALUES (?, ?, ?, ?, ?, ?)`,
		t.Id,
		t.Retro,
		t.User,
		t.Author,
		t.Category,
		t.Content,
	); err != nil {
		return err
	}
	return nil
}

func DeleteRecord(id uuid.UUID) error {
	if _, err := database.DB.Exec(
		`DELETE record WHERE id = ?`, id,
	); err != nil {
		return nil
	}
	return nil
}

package models

import (
	"arcade/internal/database"
	"log"

	"github.com/google/uuid"
)

type Record struct {
	Id       uuid.UUID
	Retro    uuid.UUID
	Author   string
	Category string
	Content  string
	Likes    int
}

func FetchRecord(id string) (Record, error) {
	var record Record
	if err := database.DB.Get(&record, `SELECT * FROM record WHERE id = ?`, id); err != nil {
		return Record{}, err
	}
	return record, nil
}

func FetchRecordsByRetro(retro_id uuid.UUID) ([]Record, error) {
	var records []Record
	if err := database.DB.Select(&records, `SELECT * FROM record WHERE retro = ?`, retro_id); err != nil {
		log.Println("Couldn't get any records: ", err)
		return []Record{}, err
	}
	return records, nil
}

func FetchRecordsByRetroAndName(retro_id uuid.UUID, name string) ([]Record, error) {
	var records []Record
	if err := database.DB.Select(&records, `SELECT * FROM record WHERE retro = ? AND author = ?`, retro_id, name); err != nil {
		log.Println("Couldn't get any records: ", err)
		return []Record{}, err
	}
	return records, nil
}

func CreateRecord(t *Record) error {
	if _, err := database.DB.Exec(
		`INSERT INTO record VALUES (?, ?, ?, ?, ?, ?)`,
		t.Id,
		t.Retro,
		t.Author,
		t.Category,
		t.Content,
		t.Likes,
	); err != nil {
		return err
	}
	return nil
}

func FetchRecordLikes(id string) (int, error) {
	var likes int
	if err := database.DB.Get(&likes, `SELECT likes FROM record WHERE id = ?`, id); err != nil {
		log.Println("Fetching likes failed: ", err)
		return 0, err
	}
	return likes, nil
}

func LikeTheRecord(id string, likes int) error {
	if _, err := database.DB.Exec(
		`UPDATE record
        SET likes = ?
        WHERE id = ?
        `, likes, id,
	); err != nil {
		return err
	}
	return nil
}

func RecordUpdateContent(id, content, author, retro string) error {
	if _, err := database.DB.Exec(
		`UPDATE record
        SET content = ?
        WHERE id = ? AND author = ? AND retro = ?
        `, content, id, author, retro,
	); err != nil {
		return err
	}
	return nil
}

func RecordDelete(id, author, retro string) error {
	if _, err := database.DB.Exec(
		`DELETE FROM record WHERE id = ? AND author = ? AND retro = ?`,
		id, author, retro,
	); err != nil {
		return err
	}
	return nil
}

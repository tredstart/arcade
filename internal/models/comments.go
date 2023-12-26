package models

import (
	"arcade/internal/database"
	"log"

	"github.com/google/uuid"
)

type Comment struct {
	Id      uuid.UUID
	Record  uuid.UUID
	Author  string
	Likes   int
	Content string
}

func FetchComment(id uuid.UUID) (Comment, error) {
	var comment Comment
	if err := database.DB.Get(&comment, `SELECT * FROM comment WHERE id = ?`, id); err != nil {
		return Comment{}, err
	}
	return comment, nil
}

func CountComments(record_id string) (int, error) {
    var count int = 0 
    if err := database.DB.Get(&count, `SELECT COUNT(id) FROM comment WHERE record = ?`, record_id); err != nil {
        log.Println(err.Error())
        return count, err
    }
    return count, nil
}

func FetchCommentsByRecord(id string) ([]Comment, error) {
	var comments []Comment
	if err := database.DB.Select(&comments, `SELECT * FROM comment WHERE record = ?`, id); err != nil {
		log.Println("Fetching comments failed: ", err)
		return nil, err
	}

	return comments, nil
}

func FetchCommentLikes(id string) (int, error) {
	var likes int
	if err := database.DB.Get(&likes, `SELECT likes FROM comment WHERE id = ?`, id); err != nil {
		log.Println("Fetching likes failed: ", err)
		return 0, err
	}
	return likes, nil
}

func CreateComment(t *Comment) error {
	if _, err := database.DB.Exec(
		`INSERT INTO comment VALUES (?, ?, ?, ?, ?)`,
		t.Id,
		t.Record,
		t.Author,
		t.Likes,
		t.Content,
	); err != nil {
		return err
	}
	return nil
}

func LikeTheComment(id string, likes int) error {
	if _, err := database.DB.Exec(
		`UPDATE comment
        SET likes = ?
        WHERE id = ?
        `, likes, id,
	); err != nil {
		return err
	}
	return nil
}

func DeleteComment(id uuid.UUID) error {
	if _, err := database.DB.Exec(
		`DELETE comment WHERE id = ?`, id,
	); err != nil {
		return nil
	}
	return nil
}

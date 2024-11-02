package models

import (
	db "arcade/internal/database"
	"log"
)

type Saved struct {
	Id      string
	Retro   string
	Created string
}

func FetchSaved(uid string) ([]Saved, error) {
	var saved []Saved
	log.Println(uid)
	if err := db.DB.Select(&saved, `SELECT s.id, r.id as retro, r.created 
                                    FROM saved s 
                                    INNER JOIN retro r ON r.id = s.retro 
                                    WHERE s.user = ?`, uid); err != nil {
		log.Println("cannot fetch joined tables, ", err)
		return []Saved{}, nil
	}
	return saved, nil
}

func CreateSave(id, rid, uid string) error {
	log.Println(uid)
	if _, err := db.DB.Exec("INSERT INTO saved VALUES (?, ?, ?)", id, uid, rid); err != nil {
		log.Println("cannoc create save, ", err)
		return err
	}
	return nil
}

func DeleteSave(id string) error {
	log.Println("saved id: ", id)
	if _, err := db.DB.Exec("DELETE FROM saved WHERE id = ?", id); err != nil {
		log.Println("cannoc create save, ", err)
		return err
	}
	return nil
}

func FetchSaveByUserAndRetro(uid, rid string) (string, error) {
	var id string
	log.Println("fetch: ", uid, rid)
	if err := db.DB.Get(&id, "SELECT id FROM saved WHERE retro = ? AND user = ?", rid, uid); err != nil {
		log.Printf("cannot find saves for %s, %s\n", uid, err)
		return "", err
	}
	return id, nil
}

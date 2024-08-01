package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var ErrDeletedNote = errors.New("note has been deleted")

type Note struct {
	Url            string `json:"note_url" 'gorm:"primaryKey;unique;not null" 'validate:"required"`
	NoteText       string `json:"note_text" validate:"required"`
	ExpirationDate string `json:"expiration_date" validate:"required"`
	MaxViews       string `json:"max_viewers" validate:"required"`
	CurrentViews   int    `json:"current_viewers"`
	Username       string `json:"username" gorm:"not null"`
}

func (db *GormDB) NewNote(notetext string, expirationdate string, maxviewers string, username string) (Note, error) {
	uuid := uuid.New().String()
	n := Note{
		Url: uuid, NoteText: notetext, ExpirationDate: expirationdate, MaxViews: maxviewers, CurrentViews: 0, Username: username,
	}
	validate := validator.New()
	err := validate.Struct(n)
	if err != nil {
		return Note{}, err
	}
	db.Client.Create(&n)
	return n, nil
}

// func (db *GormDB) DeleteNote(url string) error {
// 	return db.Client.Delete(&Note{}, url).Error
// }

func (db *GormDB) GetNote(url string) (Note, error) {
	var note Note
	result := db.Client.First(&note, url)
	if result.Error != nil {
		return Note{}, result.Error
	}
	value, _ := strconv.Atoi(note.MaxViews)
	if value == note.CurrentViews || dateExpiredHelper(note.ExpirationDate) {
		db.Client.Delete(&note)
		return Note{}, ErrDeletedNote
	}
	note.CurrentViews++
	return note, nil
}

func (db *GormDB) GetUserNotes(username string) ([]Note, error) {
	var notes []Note
	result := db.Client.Where("username = ?", username).Find(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func dateExpiredHelper(date string) bool {
	today := time.Now().Format("2006-01-02")
	if date < today {
		return true
	}
	return false
}

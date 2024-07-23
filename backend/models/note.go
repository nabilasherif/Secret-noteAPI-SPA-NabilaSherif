package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var ErrDeletedNote = errors.New("note has been deleted")

type Note struct {
	Url            string    `json:"note_url" 'gorm:"primaryKey;unique;not null" 'validate:"required"`
	NoteText       string    `json:"note_text" validate:"required"`
	ExpirationDate time.Time `json:"expiration_date" validate:"required"`
	MaxViews       int       `json:"max_viewers" validate:"required"`
	CurrentViews   int       `json:"current_viewers"`
	Username       string    `json:"username" gorm:"not null"`
}

func (db *GormDB) NewNote(notetext string, expirationdate time.Time, maxviewers int, username string) (Note, error) {
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
	if note.ExpirationDate.Before(time.Now()) || note.MaxViews == note.CurrentViews {
		db.Client.Delete(&note)
		return Note{}, ErrDeletedNote
	}
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

package models

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Note struct {
	Url            string    `json:"note_url" 'gorm:"primaryKey;unique;not null" 'validate:"required"`
	NoteText       string    `json:"note_text" validate:"required"`
	ExpirationDate time.Time `json:"expiration_date" validate:"required"`
	MaxViewers     int       `json:"max_viewers" validate:"required"`
	CurrentViewers int       `json:"current_viewers"`
	Username       string    `json:"user_name" gorm:"not null"`
}

func (db *GormDB) NewNote(notetext string, expirationdate time.Time, maxviewers int, username string) (Note, error) {
	uuid := uuid.New().String()
	n := Note{
		Url: uuid, NoteText: notetext, ExpirationDate: expirationdate, MaxViewers: maxviewers, CurrentViewers: 0, Username: username,
	}
	validate := validator.New()
	err := validate.Struct(n)
	if err != nil {
		return Note{}, err
	}
	db.Client.Create(&n)
	return n, nil
}

func (db *GormDB) DeleteNote(url string) error {
	return db.Client.Delete(&Note{}, url).Error
}

func (db *GormDB) GetNote(url string) (Note, error) {
	var note Note
	result := db.Client.First(&note, url)
	if result.Error != nil {
		return Note{}, result.Error
	}
	return note, nil
}

func (db *GormDB) GetUserNotes(username string) ([]Note, error) {
	var notes []Note
	result := db.Client.Where("user_name = ?", username).Find(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}
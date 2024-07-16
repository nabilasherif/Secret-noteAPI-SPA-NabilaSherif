package models

import "github.com/go-playground/validator"

type User struct {
	Name      string `json:"user_name" gorm:"primaryKey;unique;not null" validate:"required"`
	Password  string `json:"password" gorm:"column:password" validate:"required"`
	UserNotes []Note `json:"notes" gorm:"column:user notes" gorm:"foreignKey:Username;references:Name"`
}

func (db *GormDB) NewUser(name, password string) (User, error) {
	user := User{Name: name, Password: password}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return User{}, err
	}
	db.Client.Create(&user)
	return user, nil
}

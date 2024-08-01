package models

import (
	"errors"

	"github.com/go-playground/validator"
)

type User struct {
	Name     string `json:"username" gorm:"primaryKey;unique;not null" validate:"required"`
	Password string `json:"password" gorm:"column:password" validate:"required"`
	//UserNotes []Note `json:"notes" gorm:"column:user notes" gorm:"foreignKey:Username;references:Name"`
}

var ErrPasswordDoesnotMatch = errors.New("for the given username, the provided password is wrong")
var ErrGettingUser = errors.New("couldn't retrieve the user")

func (db *GormDB) NewUser(name, password string) (User, error) {
	var existingUser User
	result := db.Client.Where("name = ?", name).First(&existingUser)
	if result.Error == nil {
		return User{}, errors.New("username already exists")
	}
	user := User{Name: name, Password: password}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return User{}, err
	}
	db.Client.Create(&user)
	return user, nil
}

// func (db *GormDB) NewUser(name, password string) (User, error) {
// 	var existingUser User
// 	result := db.Client.Where("name = ?", name).First(&existingUser)
// 	if result.Error == nil {
// 		return User{}, errors.New("username already exists")
// 	}
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return User{}, err
// 	}
// 	user := User{Name: name, Password: string(hashedPassword)}
// 	validate := validator.New()
// 	err = validate.Struct(user)
// 	if err != nil {
// 		return User{}, err
// 	}
// 	db.Client.Create(&user)
// 	return user, nil
// }

func (db *GormDB) LogIn(username, password string) (User, error) {
	var user User
	result := db.Client.Where("name = ?", username).First(&user)
	if result.Error != nil {
		return User{}, ErrGettingUser
	}
	// }
	if user.Password != password {
		return User{}, ErrPasswordDoesnotMatch
	}
	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if err != nil {
	// 	return User{}, ErrPasswordDoesnotMatch
	// }
	return user, nil
}

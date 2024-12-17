package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(dbFilePath string) (GormDB, error) {
	dbInstance := GormDB{}
	var err error
	dbInstance.Client, err = gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{})
	return dbInstance, err
}

type GormDB struct {
	Client *gorm.DB
}

func (d *GormDB) Migrate() error {
	return d.Client.AutoMigrate(&Note{}, &User{})
}

package todostorage

import "gorm.io/gorm"

type postGresStorage struct {
	db *gorm.DB
}

func NewMySQLStorage(db *gorm.DB) *postGresStorage {
	return &postGresStorage{db: db}
}

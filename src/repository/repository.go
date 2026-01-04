package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func (*Repository[T]) FindById(db *gorm.DB, entity *T, id any) error {
	return db.First(entity, id).Error
}

func (*Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (*Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (*Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

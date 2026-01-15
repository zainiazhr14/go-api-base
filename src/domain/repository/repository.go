package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func (r *Repository[T]) FindById(entity *T, id any) error {
	return r.db.First(entity, id).Error
}

func (r *Repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

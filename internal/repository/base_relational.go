package repository

import (
	"github.com/earaujoassis/space/internal/gateways/database"
)

type BaseRepository[T any] struct {
	db *database.DatabaseService
}

func NewBaseRepository[T any](db *database.DatabaseService) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.GetDB().Create(entity).Error
}

func (r *BaseRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	err := r.db.GetDB().First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) GetAll() ([]T, error) {
	var entities []T
	err := r.db.GetDB().Find(&entities).Error
	return entities, err
}

func (r *BaseRepository[T]) Save(entity *T) error {
	return r.db.GetDB().Save(entity).Error
}

func (r *BaseRepository[T]) Delete(id uint) error {
	var entity T
	return r.db.GetDB().Delete(&entity, id).Error
}

func (r *BaseRepository[T]) FindWhere(condition string, args ...interface{}) ([]T, error) {
	var entities []T
	err := r.db.GetDB().Where(condition, args...).Find(&entities).Error
	return entities, err
}

func (r *BaseRepository[T]) Count() (int64, error) {
	var count int64
	var entity T
	err := r.db.GetDB().Model(&entity).Count(&count).Error
	return count, err
}

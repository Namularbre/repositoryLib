package repositoryLib

import "gorm.io/gorm"

type Repository[T any] interface {
	Create(t T) error
	Delete(id int) error
	FindAll() ([]T, error)
	Update(t T) error
	FindOneById(id int) (T, error)
}

type GormRepository[T any] struct {
	db *gorm.DB
}

func (r *GormRepository[T]) Create(t *T) error {
	return r.db.Create(t).Error
}

func (r *GormRepository[T]) Delete(id uint) error {
	var t T
	return r.db.Delete(&t, id).Error
}

func (r *GormRepository[T]) FindAll() ([]T, error) {
	var items []T
	err := r.db.Find(&items).Error
	return items, err
}

func (r *GormRepository[T]) Update(t *T) error {
	return r.db.Save(t).Error
}

func (r *GormRepository[T]) FindOneById(id uint) (*T, error) {
	var t T
	err := r.db.First(&t, id).Error
	return &t, err
}

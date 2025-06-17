package repository

type Repository[T any] interface {
	Create(entity *T) error
	GetByID(id uint) (*T, error)
	GetAll() ([]T, error)
	Update(entity *T) error
	Delete(id uint) error
	FindWhere(condition string, args ...interface{}) ([]T, error)
	Count() (int64, error)
}

type MemoryRepository[T any] interface {
	Save(entity *T) (*T, error)
}

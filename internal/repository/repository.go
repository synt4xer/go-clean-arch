package repository

import "context"

type repository[T any] interface {
	GetAll(ctx context.Context) ([]T, error)
	GetByID(ctx context.Context, id uint64) (T, error)
	Create(ctx context.Context, t *T) error
	Update(ctx context.Context, id uint64, t *T) error
	Delete(ctx context.Context, id uint64) error
}

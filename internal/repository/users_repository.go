package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/synt4xer/go-clean-arch/domain"
)

type UsersRepository interface {
	repository[domain.User]
}

type usersRepository struct {
	db *sqlx.DB
}

// Create implements UsersRepository.
func (u *usersRepository) Create(ctx context.Context, t *domain.User) error {
	_, err := u.db.Exec("INSERT INTO users (full_name, email, password, phone_number, is_active, created_at, updated_at) VALUES (?,?,?,?,?,?,?)",
		t.FullName, t.Email, t.Password, t.PhoneNumber, t.IsActive, t.CreatedAt, t.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// Delete implements UsersRepository.
func (u *usersRepository) Delete(ctx context.Context, id uint64) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id =?", id)

	if err != nil {
		return err
	}

	return nil
}

// GetAll implements UsersRepository.
func (u *usersRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	users := []domain.User{}

	err := u.db.Select(&users, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetByID implements UsersRepository.
func (u *usersRepository) GetByID(ctx context.Context, id uint64) (domain.User, error) {
	user := domain.User{}

	err := u.db.Get(&user, "SELECT * FROM users WHERE id =?", id)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// Update implements UsersRepository.
func (u *usersRepository) Update(ctx context.Context, id uint64, t *domain.User) error {
	_, err := u.db.Exec("UPDATE users SET full_name=?, email=?, password=?, phone_number=?, is_active=?, updated_at=? WHERE id=?",
		t.FullName, t.Email, t.Password, t.PhoneNumber, t.IsActive, t.UpdatedAt, id)

	if err != nil {
		return err
	}

	return nil
}

func New(db *sqlx.DB) UsersRepository {
	return &usersRepository{db: db}
}

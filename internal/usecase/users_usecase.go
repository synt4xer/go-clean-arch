package usecase

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/synt4xer/go-clean-arch/domain"
	"github.com/synt4xer/go-clean-arch/internal/dto"
	"github.com/synt4xer/go-clean-arch/internal/repository"
)

type UserUseCase interface {
	GetAll(ctx context.Context) ([]dto.UserResponse, error)
	GetByID(ctx context.Context, id uint64) (dto.UserResponse, error)
	Create(ctx context.Context, user *dto.UserCreateRequest) error
	Update(ctx context.Context, id uint64, user *dto.UserUpdateRequest) (dto.UserResponse, error)
	Delete(ctx context.Context, id uint64) error
}

// kalo interface, gak perlu pake pointer
type userUseCase struct {
	userRepository repository.UsersRepository
}

// Create implements UserUseCase.
func (u *userUseCase) Create(ctx context.Context, user *dto.UserCreateRequest) error {
	newUser := domain.CreateUser(*user)

	return u.userRepository.Create(ctx, &newUser)
}

// Delete implements UserUseCase.
func (u *userUseCase) Delete(ctx context.Context, id uint64) error {
	return u.userRepository.Delete(ctx, id)
}

// GetAll implements UserUseCase.
func (u *userUseCase) GetAll(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := u.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var response []dto.UserResponse
	for _, user := range users {
		response = append(response, dto.UserResponse{
			ID:          user.ID,
			Email:       user.Email,
			FullName:    user.FullName,
			PhoneNumber: user.PhoneNumber,
			IsActive:    user.IsActive,
		})
	}

	return response, nil
}

// GetByID implements UserUseCase.
func (u *userUseCase) GetByID(ctx context.Context, id uint64) (dto.UserResponse, error) {
	user, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		IsActive:    user.IsActive,
	}, nil
}

// Update implements UserUseCase.
func (u *userUseCase) Update(ctx context.Context, id uint64, user *dto.UserUpdateRequest) (dto.UserResponse, error) {
	existingUser, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("usecase failed to GetByID")
		return dto.UserResponse{}, err
	}

	updatedUser := domain.UpdateUser(*user)
	updatedUser.ID = existingUser.ID

	err = u.userRepository.Update(ctx, id, &updatedUser)
	if err != nil {
		log.Error().Err(err).Msg("usecase failed to Update")
		return dto.UserResponse{}, err
	}

	return dto.UserResponse{
		ID:          updatedUser.ID,
		Email:       updatedUser.Email,
		FullName:    updatedUser.FullName,
		PhoneNumber: updatedUser.PhoneNumber,
		IsActive:    updatedUser.IsActive,
	}, nil
}

func New(userRepository repository.UsersRepository) UserUseCase {
	return &userUseCase{userRepository: userRepository}
}

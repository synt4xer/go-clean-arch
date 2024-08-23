package http

import (
	"net/http"

	"github.com/synt4xer/go-clean-arch/internal/dto"
	"github.com/synt4xer/go-clean-arch/internal/usecase"
	"github.com/synt4xer/go-clean-arch/pkg/utils"
)

type UsersHttp struct {
	usersUseCase usecase.UserUseCase
}

func NewUsers(usersUseCase usecase.UserUseCase) *UsersHttp {
	return &UsersHttp{usersUseCase: usersUseCase}
}

func (u *UsersHttp) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := u.usersUseCase.GetAll(ctx)

	if err != nil {
		utils.JSONError(w, r, http.StatusInternalServerError, "Failed to get users", "")
		return
	}

	utils.JSONSuccess(w, r, http.StatusOK, "Users fetched successfully", users)
}

func (u *UsersHttp) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetUInt64Param(r, "id")
	if err != nil {
		utils.JSONError(w, r, http.StatusBadRequest, "Invalid ID", "")
		return
	}

	user, err := u.usersUseCase.GetByID(ctx, id)
	if err != nil {
		utils.JSONError(w, r, http.StatusInternalServerError, "Failed to get user", "")
		return
	}

	utils.JSONSuccess(w, r, http.StatusOK, "User fetched successfully", user)
}

func (u *UsersHttp) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userRequest dto.UserCreateRequest
	err := utils.DecodeJSON(r.Body, &userRequest)
	if err != nil {
		utils.JSONError(w, r, http.StatusBadRequest, "Invalid request payload", "")
		return
	}

	err = u.usersUseCase.Create(ctx, &userRequest)
	if err != nil {
		utils.JSONError(w, r, http.StatusInternalServerError, "Failed to create user", "")
		return
	}

	utils.JSONSuccess(w, r, http.StatusCreated, "User created successfully", nil)
}

func (u *UsersHttp) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetUInt64Param(r, "id")
	if err != nil {
		utils.JSONError(w, r, http.StatusBadRequest, "Invalid ID", "")
		return
	}

	var userUpdateRequest dto.UserUpdateRequest
	err = utils.DecodeJSON(r.Body, &userUpdateRequest)
	if err != nil {
		utils.JSONError(w, r, http.StatusBadRequest, "Invalid request payload", "")
		return
	}

	user, err := u.usersUseCase.Update(ctx, id, &userUpdateRequest)
	if err != nil {
		utils.JSONError(w, r, http.StatusInternalServerError, "Failed to update user", "")
		return
	}

	utils.JSONSuccess(w, r, http.StatusOK, "User updated successfully", user)
}

func (u *UsersHttp) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := utils.GetUInt64Param(r, "id")
	if err != nil {
		utils.JSONError(w, r, http.StatusBadRequest, "Invalid ID", "")
		return
	}

	err = u.usersUseCase.Delete(ctx, id)
	if err != nil {
		utils.JSONError(w, r, http.StatusInternalServerError, "Failed to delete user", "")
		return
	}

	utils.JSONSuccess(w, r, http.StatusOK, "User deleted successfully", nil)
}

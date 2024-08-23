package dto

type UserCreateRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type UserUpdateRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type UserResponse struct {
	ID          uint64 `json:"id"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	IsActive    bool   `json:"is_active"`
}

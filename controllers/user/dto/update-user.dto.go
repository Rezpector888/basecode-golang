package dto

type UpdateUserInput struct {
	Fullname string `json:"full_name" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty,email"`
}

package dto

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordDto struct {
	NewPassword string `json:"newPassword" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
}
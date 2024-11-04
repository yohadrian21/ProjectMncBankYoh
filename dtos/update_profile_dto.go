package dtos

type UpdateProfileDto struct {
	UserID      string `json:"user_id" binding:"required"`
	NewUsername string `json:"new_username"`
	NewPassword string `json:"new_password"`
}

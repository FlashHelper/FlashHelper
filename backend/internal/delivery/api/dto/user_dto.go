package dto

import "FlashHelper/internal/domain/user"

type CreateUserRequest struct {
    Username string `json:"username" binding:"required,max=100"`
    Password string `json:"password" binding:"required,min=8,max=72"`
}

type UpdateUserRequest struct {
    Username string `json:"username" binding:"required,max=100"`
    Password string `json:"password" binding:"required,min=8,max=72"`
}

type UserResponse struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Role     string `json:"role"`
}

func ToUserResponse(u *user.User) *UserResponse {
    if u == nil {
        return nil
    }

    return &UserResponse{
        ID:       u.ID,
        Username: u.Username,
        Role:     u.Role,
    }
}
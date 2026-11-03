
package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"FlashHelper/internal/delivery/api/dto"
	"FlashHelper/internal/domain/user"
	"FlashHelper/internal/usecase"
)

type UserHandler struct {
	service usecase.UserService
}

func NewUserHandler(service usecase.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Create godoc
// @Summary      Tạo người dùng mới
// @Description  Thêm người dùng mới với quyền USER mặc định
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateUserRequest true "Thông tin người dùng"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      409 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid request",
		})
		return
	}

	u := &user.User{
		Username:     req.Username,
		PasswordHash: req.Password,
	}

	if err := h.service.Create(c.Request.Context(), u); err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    dto.ToUserResponse(u),
	})
}

// GetByID godoc
// @Summary      Lấy thông tin người dùng theo ID
// @Description  Truy xuất thông tin người dùng theo ID
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      404 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := parseUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid user id",
		})
		return
	}

	u, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    dto.ToUserResponse(u),
	})
}

// List godoc
// @Summary      Danh sách người dùng
// @Description  Lấy danh sách người dùng trong hệ thống
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Success      200 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /users [get]
func (h *UserHandler) List(c *gin.Context) {
	users, err := h.service.List(c.Request.Context())
	if err != nil {
		h.handleError(c, err)
		return
	}

	data := make([]*dto.UserResponse, 0, len(users))

	for _, u := range users {
		data = append(data, dto.ToUserResponse(u))
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// Update godoc
// @Summary      Cập nhật người dùng
// @Description  Cập nhật username và mật khẩu; giữ nguyên quyền hiện tại
// @Tags         users
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Param        request body dto.UpdateUserRequest true "Thông tin cập nhật"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      404 {object} map[string]interface{}
// @Failure      409 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := parseUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid user id",
		})
		return
	}

	var req dto.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid request",
		})
		return
	}

	u := &user.User{
		ID:           id,
		Username:     req.Username,
		PasswordHash: req.Password,
	}

	if err := h.service.Update(c.Request.Context(), u); err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    dto.ToUserResponse(u),
	})
}

// Delete godoc
// @Summary      Xóa người dùng
// @Description  Xóa người dùng theo ID
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Failure      404 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := parseUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid user id",
		})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// parseUserID validates and parses the ID path parameter.
func parseUserID(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return 0, errors.New("invalid user id")
	}

	return id, nil
}

func (h *UserHandler) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, user.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "user not found",
		})

	case errors.Is(err, usecase.ErrUsernameExists):
		c.JSON(http.StatusConflict, gin.H{
			"success": false,
			"error":   "username already exists",
		})

	case errors.Is(err, usecase.ErrInvalidUsername):
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid username",
		})

	case errors.Is(err, usecase.ErrInvalidPassword):
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "invalid password",
		})

	default:
		_ = c.Error(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "internal server error",
		})
	}
}

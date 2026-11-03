package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"FlashHelper/internal/delivery/api/dto"
	"FlashHelper/internal/delivery/api/middleware"
	"FlashHelper/internal/usecase"
)

type AuthHandler struct {
	service usecase.AuthService
}

func NewAuthHandler(service usecase.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

// Login godoc
// @Summary      Đăng nhập
// @Description  Kiểm tra username/password và tạo session. Session ID được trả về trong body và trong cookie session_id
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Thông tin đăng nhập"
// @Success      200 {object} dto.LoginSuccessResponse
// @Failure      400 {object} dto.ErrorResponse "invalid request"
// @Failure      401 {object} dto.ErrorResponse "invalid username or password"
// @Failure      500 {object} dto.ErrorResponse "internal server error"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse("invalid request"))
		return
	}

	sess, err := h.service.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		h.handleError(c, err)
		return
	}

	maxAge := int(time.Until(sess.ExpiresAt).Seconds())
	setSessionCookie(c, sess.ID, maxAge)

	c.JSON(http.StatusOK, dto.NewLoginSuccessResponse(sess))
}

// Logout godoc
// @Summary      Đăng xuất
// @Description  Hủy session hiện tại
// @Tags         auth
// @Security     SessionAuth
// @Produce      json
// @Success      200 {object} dto.SuccessResponse
// @Failure      401 {object} dto.ErrorResponse "unauthorized"
// @Failure      500 {object} dto.ErrorResponse "internal server error"
// @Router       /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	sess, ok := middleware.CurrentSession(c)
	if !ok {
		h.handleError(c, usecase.ErrUnauthorized)
		return
	}

	if err := h.service.Logout(c.Request.Context(), sess.ID); err != nil {
		h.handleError(c, err)
		return
	}

	setSessionCookie(c, "", -1)

	c.JSON(http.StatusOK, dto.SuccessResponse{Success: true})
}

func setSessionCookie(c *gin.Context, value string, maxAge int) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		middleware.SessionCookieName,
		value,
		maxAge,
		"/",
		"",
		c.Request.TLS != nil,
		true,
	)
}

func (h *AuthHandler) handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, usecase.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("invalid username or password"))

	case errors.Is(err, usecase.ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, dto.NewErrorResponse("unauthorized"))

	default:
		_ = c.Error(err)

		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse("internal server error"))
	}
}

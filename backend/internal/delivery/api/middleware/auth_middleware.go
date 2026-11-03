package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"FlashHelper/internal/delivery/api/dto"
	"FlashHelper/internal/domain/session"
	"FlashHelper/internal/usecase"
)

const (
	SessionCookieName = "session_id"

	sessionContextKey = "session"
)

// RequireAuth rejects requests without a valid session.
// The session ID is read from "Authorization: Bearer <session_id>" or the session_id cookie.
func RequireAuth(authService usecase.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess, err := authService.Authenticate(c.Request.Context(), SessionIDFromRequest(c))
		if err != nil {
			if !errors.Is(err, usecase.ErrUnauthorized) {
				_ = c.Error(err)

				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewErrorResponse("internal server error"))
				return
			}

			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewErrorResponse("unauthorized"))
			return
		}

		c.Set(sessionContextKey, sess)
		c.Next()
	}
}

// CurrentSession returns the session stored by RequireAuth.
func CurrentSession(c *gin.Context) (*session.Session, bool) {
	value, ok := c.Get(sessionContextKey)
	if !ok {
		return nil, false
	}

	sess, ok := value.(*session.Session)

	return sess, ok
}

func SessionIDFromRequest(c *gin.Context) string {
	scheme, token, found := strings.Cut(c.GetHeader("Authorization"), " ")
	if found && strings.EqualFold(scheme, "Bearer") {
		return strings.TrimSpace(token)
	}

	if cookie, err := c.Cookie(SessionCookieName); err == nil {
		return cookie
	}

	return ""
}

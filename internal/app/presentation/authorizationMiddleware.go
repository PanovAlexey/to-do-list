package presentation

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo-list/internal/app/presentation/responses"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		responses.NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		responses.NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		responses.NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.authorizationService.ParseToken(headerParts[1])

	if err != nil {
		responses.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

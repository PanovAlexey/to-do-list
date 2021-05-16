package presentation

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-list/internal/app/domain/entity/user"
	"todo-list/internal/app/presentation/responses"
)

func (h *Handler) signUp(c *gin.Context) {
	var input user.User

	if err := c.BindJSON(&input); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) signIn(c *gin.Context) {

}

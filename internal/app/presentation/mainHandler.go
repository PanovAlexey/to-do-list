package presentation

import (
	"github.com/gin-gonic/gin"
	"todo-list/internal/app/application/services"
)

type Handler struct {
	authorizationService services.AuthorizationServiceInterface
	todoItemService      services.TodoItemServiceInterface
	todoListService      services.TodoListServiceInterface
}

func NewHandler(
	authorizationService services.AuthorizationServiceInterface,
	todoItemService services.TodoItemServiceInterface,
	todoListService services.TodoListServiceInterface,
) *Handler {
	return &Handler{
		authorizationService: authorizationService,
		todoItemService:      todoItemService,
		todoListService:      todoListService,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.UserIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}

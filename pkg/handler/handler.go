package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", handler.singUp)
		auth.POST("/sing-in", handler.singIn)
	}

	//works with list and tasks
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", handler.createList)
			lists.GET("/", handler.getAllLists)
			lists.POST("/:id", handler.getListById)
			lists.PUT("/:id", handler.updateList)
			lists.DELETE("/:id", handler.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", handler.createItem)
				items.GET("/", handler.getAllItem)
				items.POST("/:item_id", handler.getListById)
				items.PUT("/:item_id", handler.updateItem)
				items.DELETE("/:item_id", handler.deleteItem)
			}

		}
	}

	return router

}

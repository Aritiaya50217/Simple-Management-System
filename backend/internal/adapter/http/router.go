package http

import (
	"simple-management-system/internal/adapter/http/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	UserHandler *handler.UserHandler
}

func NewRouter(router Router) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/users", router.UserHandler.GatAll)
		api.GET("/user/:id", router.UserHandler.GetByID)
		api.POST("/user", router.UserHandler.Create)
		api.POST("/user/:id", router.UserHandler.Update)
	}
	return r
}

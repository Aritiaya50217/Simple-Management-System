package http

import (
	"simple-management-system/internal/adapter/http/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	UserHandler *handler.UserHandler
}

func NewRouter(router Router) *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		api.GET("/users", router.UserHandler.GatAll)
		api.GET("/user/:id", router.UserHandler.GetByID)
		api.POST("/user", router.UserHandler.Create)
		api.POST("/user/:id", router.UserHandler.Update)
	}
	return r
}

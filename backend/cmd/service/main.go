package main

import (
	"os"
	"simple-management-system/internal/adapter/http"
	"simple-management-system/internal/adapter/http/handler"
	"simple-management-system/internal/infrastructure/database"
	"simple-management-system/internal/infrastructure/persistence"
	"simple-management-system/internal/usecase"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// connect to database
	db := database.NewPostgresDB()

	userRepo := persistence.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	r := http.NewRouter(http.Router{
		UserHandler: userHandler,
	})

	r.Run(":" + port)

}

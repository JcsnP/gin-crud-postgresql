package routers

import (
	"github.com/JcsnP/gin-clean/pkg/handlers"
	"github.com/JcsnP/gin-clean/pkg/repositories"
	"github.com/JcsnP/gin-clean/pkg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouterSetup(r *gin.Engine, db *gorm.DB) {
	// user routes
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserRepository(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	r.GET("/users", userHandler.GetAll)
	r.GET("/users/:id", userHandler.FindByID)
	r.POST("/users", userHandler.Create)
	r.PATCH("/users/:id", userHandler.Update)
	r.DELETE("/users/:id", userHandler.Delete)
}
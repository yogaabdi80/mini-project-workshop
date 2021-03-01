package api

import (
	"workshop-golang-psm/food-app-service/internal/domain/repository"
	"workshop-golang-psm/food-app-service/internal/handler"
	"workshop-golang-psm/food-app-service/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Register User Repo
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Register Food Module
	foodRepo := repository.NewFoodRepository(db)
	foodService := service.NewFoodService(foodRepo, userRepo)
	foodHandler := handler.NewFoodHandler(foodService)

	food := router.Group("v1/food")
	{
		food.GET("/", foodHandler.GetAllFood)
		food.GET("/:food_id", foodHandler.GetDetailFood)
		food.POST("/", foodHandler.SaveFood)
		food.PUT("/:food_id", foodHandler.UpdateFood)
		food.DELETE("/:food_id", foodHandler.DeleteFood)
	}

	user := router.Group("v1/user")
	{
		user.GET("/", userHandler.GetAllUser)
		user.GET("/:user_id", userHandler.GetDetailUser)
		user.POST("/", userHandler.RegisterUser)
		user.PUT("/:user_id", userHandler.UpdateUser)
		user.DELETE("/:user_id", userHandler.DeleteUser)
	}

	return router
}

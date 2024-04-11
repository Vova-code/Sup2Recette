package main

import (
	"GoApi/database"
	"GoApi/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	/***********************/
	/*    Recipes routes   */
	/***********************/
	recipes := router.Group("/recipes")
	{
		recipes.GET("", handlers.GetRecipes)
		recipes.GET("/:id", handlers.GetRecipeById)
		recipes.POST("", handlers.PostRecipe)
		recipes.POST("/thumbs-up/:id", handlers.ThumbsUp)
		recipes.POST("/thumbs-down/:id", handlers.ThumbsDown)
		recipes.PUT("", handlers.UpdateRecipe)
		recipes.DELETE("/:id", handlers.DeleteRecipe)
	}

	database.InitDb()

	_ = router.Run(":8080")
}

package main

import (
	"GoApi/database"
	"GoApi/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	recipes := router.Group("/recipes")
	{
		recipes.GET("", handlers.GetRecipes)
		recipes.GET("/:id", handlers.GetRecipeById)
		recipes.POST("", handlers.PostRecipe)
	}

	database.InitDb()

	_ = router.Run(":8080")
}

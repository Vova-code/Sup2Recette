package handlers

import (
	postgresDb "GoApi/database"
	"GoApi/models/database"
	"GoApi/models/request"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetRecipes(c *gin.Context) {
	db := postgresDb.ConnectToDatabase()

	var recipes []database.Recipe
	result := db.Find(&recipes)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
	}

	var recipeList []request.RecipeResponse
	for _, current := range recipes {

		newRecipe := mapToRecipeResponse(&current)

		recipeList = append(recipeList, newRecipe)
	}

	c.IndentedJSON(http.StatusOK, recipeList)
}

func GetRecipeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid recipe ID"})
	}

	db := postgresDb.ConnectToDatabase()

	var recipe = database.Recipe{Model: gorm.Model{ID: uint(id)}}
	result := db.First(&recipe, id)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
	}
	response := mapToRecipeResponse(&recipe)
	c.IndentedJSON(http.StatusOK, response)
}

func PostRecipe(c *gin.Context) {
	var recipeRequest request.RecipeRequest
	if err := c.BindJSON(&recipeRequest); err != nil {
		return
	}

	recipe := mapToDbRecipe(&recipeRequest)

	db := postgresDb.ConnectToDatabase()
	result := db.Create(&recipe)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
	}

	c.IndentedJSON(http.StatusCreated, recipe)
}

func mapToRecipeResponse(dbRecipe *database.Recipe) request.RecipeResponse {
	return request.RecipeResponse{
		Title:       dbRecipe.Title,
		Steps:       dbRecipe.Steps,
		Evaluations: dbRecipe.Evaluations,
	}
}

func mapToDbRecipe(recipeRequest *request.RecipeRequest) database.Recipe {
	return database.Recipe{
		Title:       recipeRequest.Title,
		Steps:       recipeRequest.Steps,
		Evaluations: 0,
	}
}

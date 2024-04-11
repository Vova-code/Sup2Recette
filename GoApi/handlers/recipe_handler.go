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

	handleServerError(c, result)

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
	handleServerError(c, result)

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

	handleServerError(c, result)

	c.IndentedJSON(http.StatusCreated, recipe)
}

func UpdateRecipe(c *gin.Context) {
	var recipeRequest request.RecipeRequest
	if err := c.BindJSON(&recipeRequest); err != nil {
		return
	}

	db := postgresDb.ConnectToDatabase()
	result := db.Model(&database.Recipe{}).Where("id = ?", recipeRequest.Id).Updates(database.Recipe{
		Title: recipeRequest.Title,
		Steps: recipeRequest.Steps,
	})

	handleServerError(c, result)

	c.JSON(http.StatusCreated, gin.H{"message": "Recipe updated"})
}

func DeleteRecipe(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid recipe ID"})
	}
	db := postgresDb.ConnectToDatabase()

	result := db.Delete(&database.Recipe{}, id)

	handleServerError(c, result)

	c.JSON(http.StatusAccepted, gin.H{"message": "Recipe deleted"})
}

func ThumbsUp(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
	}
	db := postgresDb.ConnectToDatabase()

	var recipe = database.Recipe{Model: gorm.Model{ID: uint(id)}}
	result := db.First(&recipe, id)

	handleServerError(c, result)

	recipe.ThumbsUp = recipe.ThumbsUp + 1
	result = db.Model(&database.Recipe{}).Where("id = ?", recipe.Model.ID).Updates(database.Recipe{
		ThumbsUp: recipe.ThumbsUp,
	})

	handleServerError(c, result)

	c.JSON(http.StatusAccepted, gin.H{"message": "Recipe liked"})
}

func ThumbsDown(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
	}
	db := postgresDb.ConnectToDatabase()

	var recipe = database.Recipe{Model: gorm.Model{ID: uint(id)}}

	recipe.ThumbsDown = recipe.ThumbsDown + 1
	result := db.Model(&database.Recipe{}).Where("id = ?", recipe.Model.ID).Updates(database.Recipe{
		ThumbsDown: recipe.ThumbsDown,
	})

	handleServerError(c, result)

	c.JSON(http.StatusAccepted, gin.H{"message": "Recipe disliked"})
}

func handleServerError(c *gin.Context, result *gorm.DB) {
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}
}

func mapToRecipeResponse(dbRecipe *database.Recipe) request.RecipeResponse {
	return request.RecipeResponse{
		Id:         dbRecipe.ID,
		Title:      dbRecipe.Title,
		Steps:      dbRecipe.Steps,
		ThumbsUp:   dbRecipe.ThumbsUp,
		ThumbsDown: dbRecipe.ThumbsDown,
	}
}

func mapToDbRecipe(recipeRequest *request.RecipeRequest) database.Recipe {
	return database.Recipe{
		Title:      recipeRequest.Title,
		Steps:      recipeRequest.Steps,
		ThumbsUp:   0,
		ThumbsDown: 0,
	}
}

package db

import (
	"recipeFinder/server"

	_ "github.com/mattn/go-sqlite3"
)

func SaveResponse(response []server.RecipeDetails, ingredients string, numberOfRecipes int) {
	userIngredientsID := insertUserIngredients(ingredients, numberOfRecipes)

	for _, recipe := range response {
		calories, carbs, proteins := recipe.NutritionData.Destructure()
		missedIngredients, usedIngredients := recipe.MissedIngredients, recipe.UsedIngredients
		recipeID := insertRecipes(recipe.Title, userIngredientsID)

		insertNutritions(calories, carbs, proteins, recipeID)
		for _, ingredient := range missedIngredients {
			insertMissingIngredients(ingredient, recipeID)
		}

		for _, ingredient := range usedIngredients {
			insertUsedIngredients(ingredient, recipeID)
		}
	} 
}

func insertUserIngredients(ingredients string, numberOfRecipes int) int64 {
	statement, _ := DB.Prepare("INSERT INTO UserIngredients (Input, Number) VALUES (?, ?)")
	result, _ := statement.Exec(ingredients, numberOfRecipes)
	insertID, _ := result.LastInsertId()
	return insertID
}

func insertNutritions(calories string, carbs string, proteins string, recipeID int64) {
	statement, _ := DB.Prepare("INSERT INTO Nutritions (Calories, Carbs, Proteins, RecipeID) VALUES (?, ?, ?, ?)")
	statement.Exec(calories, carbs, proteins, recipeID)
}

func insertRecipes(title string, userIngredientsID int64) int64 {
	statement, _ := DB.Prepare("INSERT INTO Recipes (Title, UserIngredientsID) VALUES (?, ?)")
	result, _ := statement.Exec(title, userIngredientsID)
	insertID, _ := result.LastInsertId()
	return insertID
}

func insertMissingIngredients(ingredient string, recipeID int64) {
	statement, _ := DB.Prepare("INSERT INTO MissedIngredients (Ingredient, RecipeID) VALUES (?, ?)")
	statement.Exec(ingredient, recipeID)
}

func insertUsedIngredients(ingredient string, recipeID int64) {
	statement, _ := DB.Prepare("INSERT INTO UsedIngredients (Ingredient, RecipeID) VALUES (?, ?)")
	statement.Exec(ingredient, recipeID)
}

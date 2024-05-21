package db

import (
	"fmt"
	"log"
	"recipeFinder/server"

	_ "github.com/mattn/go-sqlite3"
)

func IsInputInDatabase(input string) bool {
	var exists bool
	DB.QueryRow("SELECT EXISTS(SELECT 1 FROM UserIngredients WHERE Input = ?)", input).Scan(&exists)
	return exists
}

func GetResponse(ingredients string) ([]server.RecipeDetails, int) {
	var recipes []server.RecipeDetails
	userIngredientsID, numberOfRecipes := getUserIngredientsID(ingredients)

	query := "SELECT ID, Title FROM Recipes WHERE UserIngredientsID = ?"
	rows, _ := DB.Query(query, userIngredientsID)
	defer rows.Close()

	for rows.Next() {
		var recipe server.RecipeDetails
        var recipeID int

		if err := rows.Scan(&recipeID, &recipe.Title); err != nil {
            log.Printf("Error scanning recipe row: %v", err)
            continue
        }

		calories, carbs, proteins := getNutrionByRecipeID(recipeID)
        recipe.NutritionData = server.Nutrition{Calories: calories, Carbs: carbs, Protein: proteins}

		recipe.MissedIngredients = getIngredientsByRecipeID(recipeID, "MissedIngredients")
		recipe.UsedIngredients = getIngredientsByRecipeID(recipeID, "UsedIngredients")

		recipes = append(recipes, recipe)
	}

	return recipes, numberOfRecipes
}

func getUserIngredientsID(ingredients string) (int64, int) {
	var id int64
	var number int

	query := "SELECT ID, Number FROM UserIngredients WHERE Input = ?"
	DB.QueryRow(query, ingredients).Scan(&id, &number)

	return id, number
}

func getNutrionByRecipeID(recipeID int) (string, string, string) {
	var calories, carbs, proteins string

	query := "SELECT Calories, Carbs, Proteins FROM Nutritions WHERE RecipeID = ?"
	DB.QueryRow(query, recipeID).Scan(&calories, &carbs, &proteins)

	return calories, carbs, proteins
}

func getIngredientsByRecipeID(recipeID int, tableName string) ([]string) {
	var ingredients []string

	query := fmt.Sprintf("SELECT Ingredient FROM %s WHERE RecipeID = ?", tableName)
	rows, _ := DB.Query(query, recipeID)

	for rows.Next() {
		var ingredient string
		rows.Scan(&ingredient)
		ingredients = append(ingredients, ingredient)
	}

	return ingredients
}
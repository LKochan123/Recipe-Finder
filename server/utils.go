package server

import (
	"fmt"
	"strings"
)

func HandleError(err error, message string) bool {
    if err != nil {
        fmt.Printf("%s: %v\n", message, err)
        return true 
    }
    return false
}

func (n Nutrition) Destructure() (string, string, string) {
    return n.Calories, n.Carbs, n.Protein
}

func ShowInfo(recipes []RecipeDetails) {
	for _, recipe := range recipes {
		calories, carbs, protein := recipe.NutritionData.Destructure()
		fmt.Println("Recipe name: ", recipe.Title)
		fmt.Println("Used ingredients: ", strings.Join(recipe.UsedIngredients, ", "))
		fmt.Println("Missed ingredients: ", strings.Join(recipe.MissedIngredients, ", "))
		fmt.Printf("The caloric content is %s calories, with %s of carbs and %s of protein.\n", calories, carbs, protein)
		fmt.Println("-------------------------")
	}
}

func TransformRecipes(recipes []Recipe, ingredients string) DatabaseResponse {
	var recipeDetailsArr []RecipeDetails
    recipeCounter := 0

	for _, recipe := range recipes {
		usedIngredientNames := collectIngredientNames(recipe.UsedIngredients)
		missedIngredientsNames := collectIngredientNames(recipe.MissedIngredients)
		calories, carbs, protein, err := getNutritionsById(recipe.ID)

		if err != nil {
            fmt.Printf("Error retrieving nutrition data for recipe %s: %v\n", recipe.Title, err)
			continue
        }

		recipeDetailsArr = append(recipeDetailsArr, RecipeDetails{
            recipe.Title,
            missedIngredientsNames,
			usedIngredientNames,
            Nutrition{calories, carbs, protein},
        })

		recipeCounter++
	}

	return DatabaseResponse{ingredients, recipeCounter, recipeDetailsArr}
}

func collectIngredientNames(ingredients []Ingredient) []string {
	ingredientsArr := make([]string, len(ingredients))
	for i, ingredient := range ingredients {
		ingredientsArr[i] = ingredient.Name
	}
	return ingredientsArr
}
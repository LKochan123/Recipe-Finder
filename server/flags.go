package server

import (
	"flag"
	"fmt"
)

func ParseFlags() (string, int, error) {
	ingredientsFlag := flag.String("ingredients", "", "A comma-separated list of ingredients")
    numberOfRecipesFlag := flag.Int("numberOfRecipes", 0, "The maximum number of recipes to retrieve")
	flag.Parse()

	ingredients := *ingredientsFlag
    recipesNumber := *numberOfRecipesFlag

	if ingredients == "" || recipesNumber <= 0 {
		err := fmt.Errorf("please provide both ingredients and the number of recipes in valid form")
        return "", 0, err
    }

	return ingredients, recipesNumber, nil
}
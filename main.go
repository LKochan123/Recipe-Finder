package main

import (
	"fmt"
	"recipeFinder/db"
	"recipeFinder/server"
)

func main() {
	db.InitDB()
	db.CreateTables()
	var results []server.RecipeDetails

	ingredients, recipesNumber, flagsErr := server.ParseFlags()
	if server.HandleError(flagsErr, "Failed to parse flags") {
		return
	}

	if db.IsInputInDatabase(ingredients) {
		response, numberOfRecipes := db.GetResponse(ingredients)
		results = response
		if numberOfRecipes < recipesNumber {
			fmt.Println("The number of recipes in the DB for these ingredients is now less than what you provided.")
			fmt.Println("You need to delete input manually from DB to fix that bug.")
			fmt.Println("***")
		}
	} else {
		recipes, fetchErr := server.FetchRecipes(ingredients, recipesNumber)
		if server.HandleError(fetchErr, "Failed to fetch recipes") {
			return
		}

		response := server.TransformRecipes(recipes, ingredients)
		results = response.Result
		db.SaveResponse(results, ingredients, recipesNumber)
	}

	server.ShowInfo(results)
}
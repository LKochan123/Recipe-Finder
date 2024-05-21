package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const BASE_URL string = "https://api.spoonacular.com/recipes"
var API_KEY string

func InitApiKey() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	API_KEY = os.Getenv("SPOONACULAR_API_KEY")
	
	if API_KEY == "" {
        fmt.Println("API key is not set.")
    }
}

func FetchRecipes(ingredients string, recipesNumber int) ([]Recipe, error) {
	FULL_URL := fmt.Sprintf("%s/findByIngredients?ingredients=%s&number=%d&ranking=2&apiKey=%s", BASE_URL, ingredients, recipesNumber, API_KEY)
	response, err := http.Get(FULL_URL)

	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %s", err)
	}
	
	defer response.Body.Close()
	return parseRecipes(response)
}

func parseRecipes(response *http.Response) ([]Recipe, error) {
	body, err := io.ReadAll(response.Body)
	var recipes []Recipe

	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if err := json.Unmarshal(body, &recipes); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return recipes, nil
}

func getNutritionsById(ID int) (string, string, string, error) {
	URL := fmt.Sprintf("%s/%d/nutritionWidget.json?apiKey=%s", BASE_URL, ID, API_KEY)
	response, err := http.Get(URL)
	var nutrition Nutrition

	if err != nil {
        return "", "", "", err
    }

    defer response.Body.Close()

    if err := json.NewDecoder(response.Body).Decode(&nutrition); err != nil {
        return "", "", "", err
    }

    return nutrition.Calories, nutrition.Carbs, nutrition.Protein, nil
}
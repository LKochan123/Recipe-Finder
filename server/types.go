package server

type Ingredient struct {
	Name string `json:"name"`
}

type Recipe struct {
	ID                int          `json:"id"`
	Title             string       `json:"title"`
	MissedIngredients []Ingredient `json:"missedIngredients"`
	UsedIngredients   []Ingredient `json:"usedIngredients"`
	UnusedIngredients []Ingredient `json:"unusedIngredients"`
}

type Nutrition struct {
	Calories string `json:"calories"`
	Carbs    string `json:"carbs"`
	Protein  string `json:"protein"`
}

type RecipeDetails struct {
	Title             string    `json:"title"`
	MissedIngredients []string  `json:"missedIngredients"`
	UsedIngredients   []string  `json:"usedIngredients"`
	NutritionData     Nutrition `json:"nutritionData"`
}

type DatabaseResponse struct {
	Input         string          `json:"Input"`
	RecipeCounter int             `json:"recipeCounter"`
	Result        []RecipeDetails `json:"result"`
}
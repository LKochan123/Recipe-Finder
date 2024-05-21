package db

import (
	_ "github.com/mattn/go-sqlite3"
)

func CreateTables() {
	createUserIngrediensTable()
	createRecipesTable()
	createNutritionTable()
	createMissingIngredientsTable()
	createUsedIngredientsTable()
}

func createUserIngrediensTable() {
	statement, _ := DB.Prepare(`
	CREATE TABLE IF NOT EXISTS UserIngredients (
		ID INTEGER PRIMARY KEY,
		Input TEXT,
		Number INTEGER
	)
	`)
	statement.Exec()
}

func createRecipesTable() {
	statement, _ := DB.Prepare(`
	CREATE TABLE IF NOT EXISTS Recipes (
		ID INTEGER PRIMARY KEY,
		Title TEXT,
		UserIngredientsID INTEGER,
		FOREIGN KEY (UserIngredientsID) REFERENCES UserIngredients(ID))
	`)
	statement.Exec()
}

func createNutritionTable() {
	statement, _ := DB.Prepare(`
    CREATE TABLE IF NOT EXISTS Nutritions (
        ID INTEGER PRIMARY KEY,
        Calories TEXT,
        Carbs TEXT,
        Proteins TEXT,
        RecipeID INTEGER,
        FOREIGN KEY(RecipeID) REFERENCES Recipes(ID))
	`)
	statement.Exec()
}

func createMissingIngredientsTable() {
	statement, _ := DB.Prepare(`
    CREATE TABLE IF NOT EXISTS MissedIngredients (
        ID INTEGER PRIMARY KEY,
        Ingredient TEXT,
		RecipeID INTEGER,
		FOREIGN KEY(RecipeID) REFERENCES Recipes(ID))
	`)
	statement.Exec()
}

func createUsedIngredientsTable() {
	statement, _ := DB.Prepare(`
    CREATE TABLE IF NOT EXISTS UsedIngredients (
        ID INTEGER PRIMARY KEY,
        Ingredient TEXT,
		RecipeID INTEGER,
		FOREIGN KEY(RecipeID) REFERENCES Recipes(ID))
	`)
	statement.Exec()
}
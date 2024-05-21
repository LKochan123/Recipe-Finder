# Recipe-Finder

### Description:
Recipe-finder is application that allows user to provide ingredients and number of recipes. 
Program will return the appropriate amount of meals with the minimum amount of missing ingredients.
The logic is built using [Spooncular API](https://spoonacular.com/food-api/).
The program also has a local SQLlite database in which the data entered by the user is saved along with the response from the API. 
If the input data is repeated, the program will not make an unnecessary API request.

### How to run project?
1. Clone repository
2. Create account on [Spooncular API](https://spoonacular.com/food-api/) website and generate your API key
3. Create an .env file in the root of your project and assign your api key to the variable named below:
```bash
SPOONACULAR_API_KEY = "xxx"
```
4. Run your program (example input):
```bash
./recipeFinder --ingredients=tomatoes,eggs,pasta --numberOfRecipes=5
```

### It can be helpful to:
- install SQLite viewer extension (available for VisualStudio) to see data in database file.

### Database design:
<img src="./db/DB-design.jpg">

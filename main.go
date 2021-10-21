package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type Map map[string]interface{}

type FullRecipe struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	MakingTime  string `json:"making_time"`
	Serves      string `json:"serves"`
	Ingredients string `json:"ingredients"`
	Cost        string `json:"cost"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Input struct {
	Title       string `json:"title"`
	MakingTime  string `json:"making_time"`
	Serves      string `json:"serves"`
	Ingredients string `json:"ingredients"`
	Cost        string `json:"cost"`
}

type Config struct {
	DbDriver string `json:"dbDriver"`
	DbUser   string `json:"dbUser"`
	DbPass   string `json:"dbPass"`
	DbName   string `json:"dbName"`
}

func main() {
	fmt.Println("Hello World")
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/recipes", getAllRecipes).Methods("GET")
	r.HandleFunc("/recipes/total", getNumberOfRecipes).Methods("GET")
	r.HandleFunc("/recipes", addRecipe).Methods("POST")
	r.HandleFunc("/recipe/{id}", getRecipe).Methods("GET")
	r.HandleFunc("/recipes/{id}", deleteRecipe).Methods("DELETE")
	r.HandleFunc("/recipes/{id}", updateRecipe).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World")
}

func dbConn() (db *sql.DB) {
	file, _ := os.Open("golang_config.json")
	defer file.Close()
	decode := json.NewDecoder(file)
	var config Config
	err := decode.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	dbDriver := config.DbDriver
	dbUser := config.DbUser
	dbPass := config.DbPass
	dbName := config.DbName
	fmt.Println(strings.Repeat("^", 30))
	fmt.Println(dbDriver, dbUser, dbPass, dbName)
	fmt.Println(strings.Repeat("^", 30))
	connection := fmt.Sprintf("%s%s:%s+@%s", dbDriver, dbUser, dbPass, dbName)
	fmt.Println(connection)
	db, connectError := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbName)

	if connectError != nil {
		panic(connectError.Error())
	}
	return db
}

func addRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding Recipe")
	// db := dbConn()
	// fmt.Println(body)
	fmt.Println(strings.Repeat(".", 50))
	params := mux.Vars(r)
	fmt.Println(params)
	fmt.Println(r.Body)
	var input Input
	_ = json.NewDecoder(r.Body).Decode(&input)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println(&input)
	fmt.Printf("%v+\n", &input)
	fmt.Println(strings.Repeat("-", 50))
	title := input.Title
	making_time := input.MakingTime
	serves := input.Serves
	ingredients := input.Ingredients
	cost := input.Cost
	fmt.Println(strings.Repeat(".", 50))
	db := dbConn()
	executeQuery := fmt.Sprintf("INSERT INTO recipes (title, making_time, serves, ingredients, cost)  VALUES ('%s', '%s', '%s', '%s', '%s')", title, making_time, serves, ingredients, cost)
	fmt.Println(executeQuery)
	_, addError := db.Query(executeQuery)
	if addError != nil {
		panic(addError.Error())
	}
	var response = make(map[string]interface{})
	response["message"] = "Recipe successfully created!"
	response["statusCode"] = 200
	response["recipe"] = &input

	json.NewEncoder(w).Encode(response)
}

func getAllRecipes(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Get all recipes")
	db := dbConn()
	results, err := db.Query("SELECT * FROM recipes")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("response")
	res := []FullRecipe{}
	for results.Next() {
		var recipe FullRecipe
		// for each row, scan the result into our tag composite object
		err = results.Scan(&recipe.Id, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, recipe)
	}
	defer db.Close()
	// response := Map{"recipes": res}
	var response = make(map[string]interface{})
	response["recipes"] = res
	response["statusCode"] = 200
	byteArray, err := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(byteArray))
	// return byteArray
	// fmt.Println(w, string(byteArray))
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(response)
}

func getRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetRecipe")
	db := dbConn()
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	results, err := db.Query("SELECT * FROM recipes WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	res := []FullRecipe{}
	for results.Next() {
		var recipe FullRecipe
		err = results.Scan(&recipe.Id, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, recipe)
	}
	defer db.Close()
	response := Map{"recipes": res}
	byteArray, err := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(byteArray))
	if err == nil {
		json.NewEncoder(w).Encode(response)
	}
}

func getRecipeByID(id string) bool {
	var count int
	db := dbConn()
	results, err := db.Query("SELECT COUNT(*) as count FROM recipes WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(results)
	for results.Next() {
		countErr := results.Scan(&count)
		if countErr != nil {
			panic(countErr.Error())
		}
	}
	fmt.Println("count", count)
	if count == 1 {
		return true
	} else {
		return false
	}
}

func getNumberOfRecipes(w http.ResponseWriter, r *http.Request) {
	var count int
	fmt.Println(strings.Repeat("+", 30))
	fmt.Println("Getting Number of Recipes")
	fmt.Println(strings.Repeat("+", 30))
	db := dbConn()
	results, _ := db.Query("SELECT COUNT(*) as count FROM recipes")
	for results.Next() {
		countErr := results.Scan(&count)
		if countErr != nil {
			panic(countErr.Error())
		}
	}
	fmt.Println("count", count)
	var response = make(map[string]interface{})
	response["message"] = fmt.Sprintf("Number of Recipes %d", count)
	response["statusCode"] = 200
	json.NewEncoder(w).Encode(response)
}

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteRecipe")
	db := dbConn()
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)

	count, countErr := db.Query("SELECT COUNT(*) FROM recipes WHERE id = ?", id)
	if countErr != nil {
		panic(countErr.Error())
	}
	fmt.Println(count, "count")
	fmt.Printf("Type of count is %T", count)
	results, err := db.Query("SELECT * FROM recipes WHERE id = ?", id)
	// results, err := db.Query("SELECT * FROM recipes")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(results)
	var response = make(map[string]interface{})
	if getRecipeByID(id) {
		delResult, delErr := db.Query("DELETE FROM recipes WHERE id = ?", id)
		if delErr != nil {
			panic(delErr.Error())
		}
		fmt.Println(delResult)
		response["message"] = "Recipe successfully removed!"
		response["statusCode"] = 200
	} else {
		response["message"] = "No Recipe found"
		response["statusCode"] = 200
	}
	json.NewEncoder(w).Encode(response)
}

func updateRecipe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("Updating Recipe with id = ", id)
	var input Input
	_ = json.NewDecoder(r.Body).Decode(&input)
	title := input.Title
	making_time := input.MakingTime
	serves := input.Serves
	ingredients := input.Ingredients
	cost := input.Cost
	db := dbConn()
	defer db.Close()
	if getRecipeByID(id) {
		executeQuery := fmt.Sprintf("UPDATE recipes set title = '%s' , making_time = '%s' , serves = '%s' , ingredients = '%s', cost = '%s' WHERE id = '%s'", title, making_time, serves, ingredients, cost, id)
		fmt.Println(executeQuery)
		_, updateError := db.Query(executeQuery)
		if updateError != nil {
			panic(updateError.Error())
		}

		execGetQuery := fmt.Sprintf("SELECT * FROM recipes WHERE id = %s", id)
		fmt.Println(execGetQuery)
		results, getError := db.Query(execGetQuery)
		if getError != nil {
			panic(getError.Error())
		}
		var recipe FullRecipe
		for results.Next() {
			err := results.Scan(&recipe.Id, &recipe.Title, &recipe.MakingTime, &recipe.Serves, &recipe.Ingredients, &recipe.Cost, &recipe.CreatedAt, &recipe.UpdatedAt)
			if err != nil {
				panic(err.Error())
			}
		}
		var response = make(map[string]interface{})
		response["message"] = "Recipe successfully updated!!"
		response["recipe"] = &recipe
		response["statusCode"] = 200
		json.NewEncoder(w).Encode(response)
	} else {
		var response = make(map[string]interface{})
		response["message"] = "Recipe not found to update"
		response["statusCode"] = 200
		json.NewEncoder(w).Encode(response)
	}
}

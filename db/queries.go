package db

import (
	"encoding/xml"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucasbravi2019/pasteleria/pkg"
)

var (
	Recipe_FindAll                         = "recipe.findAll"
	Recipe_FindById                        = "recipe.findById"
	Recipe_Create                          = "recipe.create"
	Recipe_UpdateName                      = "recipe.updateName"
	Recipe_DeleteById                      = "recipe.deleteById"
	Recipe_FindRecipeIngredients           = "recipe.findRecipeIngredients"
	Recipe_DeleteIngredientsByRecipeId     = "recipe.deleteIngredientsByRecipeId"
	Recipe_AddIngredientsToRecipe          = "recipe.addIngredientsToRecipe"
	Recipe_FindRecipeIngredientById        = "recipe.findRecipeIngredientById"
	Recipe_UpdateRecipeIngredientPriceById = "recipe.updateRecipeIngredientPriceById"
	Ingredient_FindAll                     = "ingredient.findAll"
	Ingredient_Create                      = "ingredient.create"
	Ingredient_UpdateById                  = "ingredient.updateById"
	Ingredient_DeleteById                  = "ingredient.deleteById"
	Ingredient_AddPackage                  = "ingredient.addPackage"
	Ingredient_DeletePackage               = "ingredient.deletePackage"
	Ingredient_FindAllIngredientPackages   = "ingredient.findAllIngredientPackages"
	Ingredient_FindIngredientIdByName      = "ingredient.findIngredientIdByName"
	Package_FindAll                        = "package.findAll"
	Package_Create                         = "package.create"
	Package_UpdateById                     = "package.updateById"
	Package_DeleteById                     = "package.deleteById"
)

const QUERIES_PATH = "db/queries"

var queries map[string]string = make(map[string]string)

type XMLQueries struct {
	XMLName xml.Name   `xml:"queries"`
	List    []XMLQuery `xml:"query"`
}

type XMLQuery struct {
	XMLName xml.Name `xml:"query"`
	ID      string   `xml:"id,attr"`
	SQL     string   `xml:",chardata"`
}

func GetQueryByName(queryName string) (string, error) {
	query, err := findQueryByName(queryName)

	if pkg.HasError(err) {
		return pkg.STRING_EMPTY, err
	}

	return strings.TrimSpace(query), nil
}

func QueryLoader() {
	if err := processXmlFileInDirectory(QUERIES_PATH); err != nil {
		log.Fatal(err)
	}
}

func findQueryByName(queryName string) (string, error) {
	query := queries[queryName]

	if query == pkg.STRING_EMPTY {
		return pkg.STRING_EMPTY, errors.New("query not found")
	}

	return query, nil
}

func processXmlFile(filePath string) error {
	file, err := os.Open(filePath)

	if pkg.HasError(err) {
		return err
	}

	defer file.Close()

	var xmlQueries XMLQueries

	if err := xml.NewDecoder(file).Decode(&xmlQueries); err != nil {
		return err
	}

	for _, query := range xmlQueries.List {
		queries[query.ID] = query.SQL
	}

	return nil
}

func processXmlFileInDirectory(directoryPath string) error {
	files, err := os.ReadDir(directoryPath)

	if pkg.HasError(err) {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == pkg.XML_EXT {
			path := filepath.Join(directoryPath, file.Name())
			if err := processXmlFile(path); err != nil {
				return err
			}
		}
	}

	return nil
}

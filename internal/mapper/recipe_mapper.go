package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type RecipeMapper struct {
	RecipeIngredientMapper RecipeIngredientMapper
	PackageMapper          PackageMapper
}

var RecipeMapperInstance *RecipeMapper

type RecipeMapperInterface interface {
	ToRecipeList(rows *sql.Rows) *[]models.Recipe
	ToRecipeRow(rows *sql.Rows) *models.Recipe
	ToRecipeDTOList(recipes *[]models.Recipe) *[]dto.RecipeDTO
	ToRecipe(recipeId int64, recipeName string, recipePrice float64) *models.Recipe
}

func (m *RecipeMapper) ToRecipeList(rows *sql.Rows) *[]models.Recipe {
	recipesGrouper := util.NewMap[int64, models.Recipe]()

	for rows.Next() {
		var recipeId int64
		var recipeName string
		var recipePrice float64
		var ingredientId sql.NullInt64
		var ingredientName sql.NullString
		var ingredientQuantity sql.NullFloat64
		var ingredientPrice sql.NullFloat64
		var packageId sql.NullInt64
		var metric sql.NullString
		var quantity sql.NullFloat64
		var packagePrice sql.NullFloat64

		err := rows.Scan(&recipeId, &recipeName, &recipePrice, &ingredientId, &ingredientName, &ingredientQuantity, &ingredientPrice, &packageId,
			&metric, &quantity, &packagePrice)

		if pkg.HasError(err) {
			return nil
		}

		ingredientRecipeId := sql.NullInt64{
			Int64: recipeId,
			Valid: true,
		}

		var ingredientPackage *models.RecipeIngredientPackage
		var recipeIngredient *models.RecipeIngredient

		if packageId.Valid {
			ingredientPackage = m.PackageMapper.ToRecipeIngredientPackage(packageId, metric, quantity, packagePrice)
		}

		if ingredientId.Valid {
			recipeIngredient = m.RecipeIngredientMapper.ToRecipeIngredient(ingredientId, ingredientName, ingredientQuantity,
				ingredientPrice, ingredientRecipeId)

			m.RecipeIngredientMapper.SetPackageToRecipeIngredientDTO(recipeIngredient, ingredientPackage)
		}

		recipe := util.GetValue(recipesGrouper, recipeId)
		if recipe == nil {
			recipe = m.ToRecipe(recipeId, recipeName, recipePrice)
		}

		if recipeIngredient != nil {
			util.Add(&recipe.Ingredients, *recipeIngredient)
		}
		util.PutValue(&recipesGrouper, recipeId, *recipe)
	}

	recipes := util.NewList[models.Recipe]()

	for _, recipe := range recipesGrouper {
		util.Add(&recipes, recipe)
	}

	return &recipes
}

func (m *RecipeMapper) ToRecipeDTOList(recipes *[]models.Recipe) *[]dto.RecipeDTO {
	dtos := util.NewList[dto.RecipeDTO]()

	for _, recipe := range *recipes {
		dto := dto.RecipeDTO{
			Id:          recipe.Id,
			Name:        recipe.Name,
			Price:       recipe.Price,
			Ingredients: *m.RecipeIngredientMapper.ToRecipeIngredientDTOList(&recipe.Ingredients),
		}

		util.Add(&dtos, dto)
	}

	return &dtos
}

func (m *RecipeMapper) ToRecipe(recipeId int64, recipeName string, recipePrice float64) *models.Recipe {
	return models.NewRecipe(recipeId, recipeName, recipePrice)
}

func (m *RecipeMapper) ToRecipeRow(rows *sql.Rows) *models.Recipe {
	var recipe *models.Recipe
	for rows.Next() {
		var recipeId int64
		var recipeName string
		var recipePrice float64
		var ingredientId sql.NullInt64
		var ingredientName sql.NullString
		var ingredientQuantity sql.NullFloat64
		var ingredientPrice sql.NullFloat64
		var packageId sql.NullInt64
		var metric sql.NullString
		var quantity sql.NullFloat64
		var packagePrice sql.NullFloat64

		err := rows.Scan(&recipeId, &recipeName, &recipePrice, &ingredientId, &ingredientName, &ingredientQuantity, &ingredientPrice, &packageId,
			&metric, &quantity, &packagePrice)

		if pkg.HasError(err) {
			return nil
		}

		ingredientRecipeId := sql.NullInt64{
			Int64: recipeId,
			Valid: true,
		}

		var ingredientPackage *models.RecipeIngredientPackage
		var recipeIngredient *models.RecipeIngredient

		if packageId.Valid {
			ingredientPackage = m.PackageMapper.ToRecipeIngredientPackage(packageId, metric, quantity, packagePrice)
		}

		if ingredientId.Valid {
			recipeIngredient = m.RecipeIngredientMapper.ToRecipeIngredient(ingredientId, ingredientName, ingredientQuantity,
				ingredientPrice, ingredientRecipeId)

			m.RecipeIngredientMapper.SetPackageToRecipeIngredientDTO(recipeIngredient, ingredientPackage)
		}

		if recipe == nil {
			recipe = m.ToRecipe(recipeId, recipeName, recipePrice)
		}

		if recipeIngredient != nil {
			util.Add(&recipe.Ingredients, *recipeIngredient)
		}
	}

	return recipe
}

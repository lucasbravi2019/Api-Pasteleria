package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type RecipeMapper struct {
	PackageMapper    *PackageMapper
	IngredientMapper *IngredientMapper
}

var RecipeMapperInstance *RecipeMapper

func (m *RecipeMapper) ToRecipeList(rows *sql.Rows) (*[]dto.Recipe, error) {
	recipesGrouper := util.NewMap[int64, dto.Recipe]()

	for rows.Next() {
		var recipeId int64
		var recipeName string
		var recipePrice float64
		var recipeIngredientId sql.NullInt64
		var recipeIngredientPrice sql.NullFloat64
		var recipeIngredientQuantity sql.NullFloat64
		var ingredientPackageId sql.NullInt64
		var ingredientPackagePrice sql.NullFloat64
		var ingredientId sql.NullInt64
		var ingredientName sql.NullString
		var packageId sql.NullInt64
		var metric sql.NullString
		var quantity sql.NullFloat64

		err := rows.Scan(&recipeId, &recipeName, &recipePrice, &recipeIngredientId, &recipeIngredientPrice, &recipeIngredientQuantity,
			&ingredientPackageId, &ingredientPackagePrice, &ingredientId, &ingredientName, &packageId, &metric, &quantity)

		if pkg.HasError(err) {
			return nil, err
		}

		pkg := m.PackageMapper.ToPackageNullable(packageId, metric, quantity)
		ingredient := m.IngredientMapper.ToIngredientNullable(ingredientId, ingredientName)
		ingredientPackage := m.IngredientMapper.ToIngredientPackageNullable(ingredientPackageId, ingredientPackagePrice, ingredient, pkg)
		recipeIngredient := m.ToRecipeIngredientNullable(recipeIngredientId, recipeIngredientPrice, recipeIngredientQuantity, ingredientPackage)

		recipe := util.GetValue(recipesGrouper, recipeId)
		if recipe == nil {
			recipe = m.toRecipe(recipeId, recipeName, recipePrice)
		}

		if recipeIngredient != nil {
			util.Add(recipe.Ingredients, *recipeIngredient)
		}

		util.PutValue(&recipesGrouper, &recipeId, recipe)
	}

	recipes := util.NewList[dto.Recipe]()

	for _, recipe := range recipesGrouper {
		util.Add(&recipes, recipe)
	}

	return &recipes, nil
}

func (m *RecipeMapper) ToRecipeRow(rows *sql.Rows) (*dto.Recipe, error) {
	var recipe *dto.Recipe

	for rows.Next() {
		var recipeId int64
		var recipeName string
		var recipePrice float64
		var recipeIngredientId sql.NullInt64
		var recipeIngredientPrice sql.NullFloat64
		var recipeIngredientQuantity sql.NullFloat64
		var ingredientPackageId sql.NullInt64
		var ingredientPackagePrice sql.NullFloat64
		var ingredientId sql.NullInt64
		var ingredientName sql.NullString
		var packageId sql.NullInt64
		var metric sql.NullString
		var quantity sql.NullFloat64

		err := rows.Scan(&recipeId, &recipeName, &recipePrice, &recipeIngredientId, &recipeIngredientPrice, &recipeIngredientQuantity,
			&ingredientPackageId, &ingredientPackagePrice, &ingredientId, &ingredientName, &packageId, &metric, &quantity)

		if pkg.HasError(err) {
			return nil, err
		}

		pkg := m.PackageMapper.ToPackageNullable(packageId, metric, quantity)
		ingredient := m.IngredientMapper.ToIngredientNullable(ingredientId, ingredientName)
		ingredientPackage := m.IngredientMapper.ToIngredientPackageNullable(ingredientPackageId, ingredientPackagePrice, ingredient, pkg)
		recipeIngredient := m.ToRecipeIngredientNullable(recipeIngredientId, recipeIngredientPrice, recipeIngredientQuantity, ingredientPackage)

		if recipe == nil {
			recipe = m.toRecipe(recipeId, recipeName, recipePrice)
		}

		if recipeIngredient != nil {
			util.Add(recipe.Ingredients, *recipeIngredient)
		}
	}

	return recipe, nil
}

func (m *RecipeMapper) toRecipe(recipeId int64, recipeName string, recipePrice float64) *dto.Recipe {
	return &dto.Recipe{
		Id:          &recipeId,
		Name:        &recipeName,
		Price:       &recipePrice,
		Ingredients: &[]dto.RecipeIngredient{},
	}
}

func (m *RecipeMapper) ToRecipeIngredientNullable(id sql.NullInt64, price sql.NullFloat64, quantity sql.NullFloat64,
	ingredientPackage *dto.IngredientPackage) *dto.RecipeIngredient {
	if !id.Valid {
		return nil
	}

	return &dto.RecipeIngredient{
		Id:         db.GetLong(id),
		Quantity:   db.GetFloat(quantity),
		Price:      db.GetFloat(price),
		Ingredient: ingredientPackage,
	}
}

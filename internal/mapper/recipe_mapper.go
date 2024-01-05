package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type RecipeMapper struct {
	RecipeIngredientMapper RecipeIngredientMapper
	PackageMapper          PackageMapper
}

var RecipeMapperInstance *RecipeMapper

type RecipeMapperInterface interface {
	ToRecipeList(rows *sql.Rows) (*[]dto.RecipeDTO, error)
	ToRecipeRow(rows *sql.Rows) (*dto.RecipeDTO, error)
	toRecipe(recipeId int64, recipeName string, recipePrice float64) *dto.RecipeDTO
}

func (m *RecipeMapper) ToRecipeList(rows *sql.Rows) (*[]dto.RecipeDTO, error) {
	recipesGrouper := util.NewMap[int64, dto.RecipeDTO]()

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
			return nil, err
		}

		var ingredientPackage *dto.IngredientPackageDTO
		var recipeIngredient *dto.RecipeIngredientDTO

		ingredientPackage = m.PackageMapper.ToIngredientPackage(packageId, metric, quantity, packagePrice)
		recipeIngredient = m.RecipeIngredientMapper.ToRecipeIngredientDTO(ingredientId, ingredientName, ingredientQuantity,
			ingredientPrice, ingredientPackage)

		recipe := util.GetValue(recipesGrouper, recipeId)
		if recipe == nil {
			recipe = m.toRecipe(recipeId, recipeName, recipePrice)
		}

		if recipeIngredient != nil {
			util.Add(&recipe.Ingredients, *recipeIngredient)
		}
		util.PutValue(&recipesGrouper, recipeId, *recipe)
	}

	recipes := util.NewList[dto.RecipeDTO]()

	for _, recipe := range recipesGrouper {
		util.Add(&recipes, recipe)
	}

	return &recipes, nil
}

func (m *RecipeMapper) ToRecipeRow(rows *sql.Rows) (*dto.RecipeDTO, error) {
	var recipe *dto.RecipeDTO

	for rows.Next() {
		var id int64
		var name string
		var price float64
		var ingredientId sql.NullInt64
		var ingredientName sql.NullString
		var ingredientQuantity sql.NullFloat64
		var ingredientPrice sql.NullFloat64
		var packageId sql.NullInt64
		var metric sql.NullString
		var packageQuantity sql.NullFloat64
		var packagePrice sql.NullFloat64

		err := rows.Scan(&id, &name, &price, &ingredientId, &ingredientName, &ingredientQuantity, &ingredientPrice,
			&packageId, &metric, &packageQuantity, &packagePrice)

		if pkg.HasError(err) {
			return nil, err
		}

		if recipe == nil {
			recipe = m.toRecipe(id, name, price)
		}

		pkg := m.toPackage(packageId, metric, packageQuantity, packagePrice)
		ingredient := m.toIngredient(ingredientId, ingredientName, ingredientQuantity, ingredientPrice, pkg)
		m.addRecipeIngredient(recipe, ingredient)
	}

	return recipe, nil
}

func (m *RecipeMapper) toRecipe(recipeId int64, recipeName string, recipePrice float64) *dto.RecipeDTO {
	return &dto.RecipeDTO{
		Id:    recipeId,
		Name:  recipeName,
		Price: recipePrice,
	}
}

func (m *RecipeMapper) toIngredient(id sql.NullInt64, name sql.NullString, quantity sql.NullFloat64,
	price sql.NullFloat64, pkg *dto.IngredientPackageDTO) *dto.RecipeIngredientDTO {
	if !id.Valid {
		return nil
	}

	return &dto.RecipeIngredientDTO{
		Id:       db.GetLong(id),
		Name:     db.GetString(name),
		Price:    db.GetFloat(price),
		Quantity: db.GetFloat(quantity),
		Package:  pkg,
	}
}

func (m *RecipeMapper) toPackage(id sql.NullInt64, metric sql.NullString, quantity sql.NullFloat64,
	price sql.NullFloat64) *dto.IngredientPackageDTO {
	if !id.Valid {
		return nil
	}

	return &dto.IngredientPackageDTO{
		Id:       db.GetLong(id),
		Metric:   db.GetString(metric),
		Quantity: db.GetFloat(quantity),
		Price:    db.GetFloat(price),
	}
}

func (m *RecipeMapper) addRecipeIngredient(recipe *dto.RecipeDTO, ingredients *dto.RecipeIngredientDTO) {
	util.Add(&recipe.Ingredients, *ingredients)
}

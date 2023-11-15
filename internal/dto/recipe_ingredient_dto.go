package dto

type RecipeIngredientIdDTO struct {
	RecipeId    int64                     `json:"recipeId" validate:"required"`
	Ingredients []IngredientIdQuantityDTO `json:"ingredients" validate:"required"`
}

type IngredientIdQuantityDTO struct {
	IngredientId int64   `json:"ingredientId" validate:"required"`
	Quantity     float64 `json:"quantity" validate:"required"`
}

type RecipeIngredientDTO struct {
	Id       int64      `json:"id"`
	Name     string     `json:"name"`
	Price    float64    `json:"price"`
	Package  PackageDTO `json:"package"`
	Quantity float64    `json:"quantity"`
}

func NewRecipeIngredientIdDTO(recipeId int64, ingredients []IngredientIdQuantityDTO) *RecipeIngredientIdDTO {
	return &RecipeIngredientIdDTO{
		RecipeId:    recipeId,
		Ingredients: ingredients,
	}
}

func NewIngredientIdQuantityDTO(ingredientId int64, quantity float64) *IngredientIdQuantityDTO {
	return &IngredientIdQuantityDTO{
		IngredientId: ingredientId,
		Quantity:     quantity,
	}
}

func NewRecipeIngredientDTO(id int64, name string, quantity float64, pkg PackageDTO) *RecipeIngredientDTO {
	return &RecipeIngredientDTO{
		Id:       id,
		Name:     name,
		Quantity: quantity,
		Package:  pkg,
	}
}

func (r *RecipeIngredientDTO) UpdatePrice() {
	packageQuantity := r.Package.Quantity
	price := r.Package.Price
	recipeQuantity := r.Quantity

	r.Price = recipeQuantity / packageQuantity * price
}

package dto

type RecipeIngredientIdDTO struct {
	RecipeId     string `bson:"_id" json:"recipeId" validate:"required"`
	IngredientId string `bson:"_id" json:"ingredientId" validate:"required"`
}

type RecipeIngredientDTO struct {
	Id       int64
	Name     string
	Price    float64
	Package  PackageDTO
	Quantity float64
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

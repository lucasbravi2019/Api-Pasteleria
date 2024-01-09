package dto

type Recipe struct {
	Id          *int64              `json:"id"`
	Name        *string             `json:"name"`
	Price       *float64            `json:"price"`
	Ingredients *[]RecipeIngredient `json:"ingredients"`
}

type RecipeIngredient struct {
	Id         *int64             `json:"id"`
	Quantity   *float64           `json:"quantity"`
	Price      *float64           `json:"price"`
	Ingredient *IngredientPackage `json:"ingredient"`
}

type IngredientPackage struct {
	Id         *int64      `json:"id"`
	Price      *float64    `json:"price"`
	Ingredient *Ingredient `json:"ingredient"`
	Package    *Package    `json:"package"`
}

type Ingredient struct {
	Id   *int64  `json:"id"`
	Name *string `json:"name"`
}

type Package struct {
	Id       *int64   `json:"id"`
	Metric   *string  `json:"metric"`
	Quantity *float64 `json:"quantity"`
}

type RecipeRequest struct {
	Id          *int64                     `json:"id"`
	Name        *string                    `json:"name"`
	Ingredients *[]RecipeIngredientRequest `json:"ingredients"`
}

type RecipeIngredientRequest struct {
	Id       *int64   `json:"id"`
	Quantity *float64 `json:"quantity"`
}

type IngredientRequest struct {
	Id       *int64                      `json:"id"`
	Name     *string                     `json:"name"`
	Packages *[]IngredientPackageRequest `json:"packages"`
}

type IngredientPackageRequest struct {
	Id        *int64   `json:"id"`
	PackageId *int64   `json:"packageId"`
	Price     *float64 `json:"price"`
}

type IngredientResponse struct {
	Id       *int64                       `json:"id"`
	Name     *string                      `json:"name"`
	Packages *[]IngredientPackageResponse `json:"packages"`
}

type IngredientPackageResponse struct {
	IngredientPackageId *int64   `json:"ingredientPackageId"`
	Price               *float64 `json:"price"`
	Package             *Package `json:"package"`
}

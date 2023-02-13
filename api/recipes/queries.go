package recipes

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func All() bson.M {
	return bson.M{}
}

func GetRecipeById(oid primitive.ObjectID) bson.M {
	return bson.M{"_id": oid}
}

func UpdateRecipeName(dto RecipeNameDTO) bson.M {
	return bson.M{"$set": bson.M{"name": dto.Name}}
}

func AddIngredientToRecipe(recipe RecipeIngredient) bson.M {
	return bson.M{"$addToSet": bson.M{"ingredients": recipe}}
}

func RemoveIngredientFromRecipe(recipe RecipeIngredient) bson.M {
	return bson.M{"$pull": bson.M{"ingredients._id": recipe.ID}}
}

func SetRecipePrice() bson.A {
	return bson.A{bson.D{{"$set", bson.D{{"price", bson.D{{"$sum", "$ingredients.price"}}}}}}}
}

func SetIngredientPackagePrice(price float64) bson.D {
	return bson.D{{"$set", bson.D{{"ingredients.$[ingredient].package.price", price}}}}
}

func SetIngredientPrice(recipe RecipeDTO) bson.D {
	return bson.D{{"$set", recipe}}
}

func GetArrayFiltersForIngredientsByPackageId(packageId primitive.ObjectID) *options.UpdateOptions {
	return options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"ingredient.package._id": packageId},
		},
	})
}

func GetRecipeByPackageId(packageId primitive.ObjectID) bson.M {
	return bson.M{"ingredients.package._id": packageId}
}

func GetRecipeByIngredientId(ingredientId primitive.ObjectID) bson.M {
	return bson.M{"ingredients._id": ingredientId}
}

func RemovePackageFromRecipes(packageId primitive.ObjectID) bson.M {
	return bson.M{"$pull": bson.M{"ingredients": bson.M{"package._id": packageId}}}
}

package queries

import (
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
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

func UpdateRecipeName(recipe models.Recipe) bson.M {
	return bson.M{"$set": bson.M{"name": recipe.Name}}
}

func AddIngredientToRecipe(recipe models.RecipeIngredient) bson.M {
	return bson.M{"$addToSet": bson.M{"ingredients": recipe}}
}

func RemoveIngredientFromRecipe(recipe models.Recipe) bson.M {
	return bson.M{"$set": recipe}
}

func SetRecipePrice() bson.A {
	return bson.A{bson.D{{"$set", bson.D{{"price", bson.D{{"$multiply", bson.A{bson.D{{"$sum", "$ingredients.price"}}, 3}}}}}}}}
}

func SetIngredientPackagePrice(price float64) bson.D {
	return bson.D{{"$set", bson.D{{"ingredients.$[ingredient].package.price", price}}}}
}

func SetRecipeIngredientPrice(recipe dto.RecipeDTO) bson.D {
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

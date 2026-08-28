// Typed models for the FreeMeal SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/free-meal-sdk/go/core"
)

// Category is the typed data model for the category entity.
type Category struct {
	IdCategory *string `json:"idCategory,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrCategoryDescription *string `json:"strCategoryDescription,omitempty"`
	StrCategoryThumb *string `json:"strCategoryThumb,omitempty"`
}

// CategoryListMatch is the typed request payload for Category.ListTyped.
type CategoryListMatch struct {
	IdCategory *string `json:"idCategory,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrCategoryDescription *string `json:"strCategoryDescription,omitempty"`
	StrCategoryThumb *string `json:"strCategoryThumb,omitempty"`
}

// Filter is the typed data model for the filter entity.
type Filter struct {
	IdMeal *string `json:"idMeal,omitempty"`
	StrMeal *string `json:"strMeal,omitempty"`
	StrMealThumb *string `json:"strMealThumb,omitempty"`
}

// FilterListMatch is the typed request payload for Filter.ListTyped.
type FilterListMatch struct {
	A *string `json:"a,omitempty"`
	C *string `json:"c,omitempty"`
	I *string `json:"i,omitempty"`
}

// Latest is the typed data model for the latest entity.
type Latest struct {
	DateModified *string `json:"dateModified,omitempty"`
	IdMeal *string `json:"idMeal,omitempty"`
	StrArea *string `json:"strArea,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"strCreativeCommonsConfirmed,omitempty"`
	StrDrinkAlternate *string `json:"strDrinkAlternate,omitempty"`
	StrImageSource *string `json:"strImageSource,omitempty"`
	StrIngredient1 *string `json:"strIngredient1,omitempty"`
	StrIngredient10 *string `json:"strIngredient10,omitempty"`
	StrIngredient11 *string `json:"strIngredient11,omitempty"`
	StrIngredient12 *string `json:"strIngredient12,omitempty"`
	StrIngredient13 *string `json:"strIngredient13,omitempty"`
	StrIngredient14 *string `json:"strIngredient14,omitempty"`
	StrIngredient15 *string `json:"strIngredient15,omitempty"`
	StrIngredient16 *string `json:"strIngredient16,omitempty"`
	StrIngredient17 *string `json:"strIngredient17,omitempty"`
	StrIngredient18 *string `json:"strIngredient18,omitempty"`
	StrIngredient19 *string `json:"strIngredient19,omitempty"`
	StrIngredient2 *string `json:"strIngredient2,omitempty"`
	StrIngredient20 *string `json:"strIngredient20,omitempty"`
	StrIngredient3 *string `json:"strIngredient3,omitempty"`
	StrIngredient4 *string `json:"strIngredient4,omitempty"`
	StrIngredient5 *string `json:"strIngredient5,omitempty"`
	StrIngredient6 *string `json:"strIngredient6,omitempty"`
	StrIngredient7 *string `json:"strIngredient7,omitempty"`
	StrIngredient8 *string `json:"strIngredient8,omitempty"`
	StrIngredient9 *string `json:"strIngredient9,omitempty"`
	StrInstructions *string `json:"strInstructions,omitempty"`
	StrMeal *string `json:"strMeal,omitempty"`
	StrMealThumb *string `json:"strMealThumb,omitempty"`
	StrMeasure1 *string `json:"strMeasure1,omitempty"`
	StrMeasure10 *string `json:"strMeasure10,omitempty"`
	StrMeasure11 *string `json:"strMeasure11,omitempty"`
	StrMeasure12 *string `json:"strMeasure12,omitempty"`
	StrMeasure13 *string `json:"strMeasure13,omitempty"`
	StrMeasure14 *string `json:"strMeasure14,omitempty"`
	StrMeasure15 *string `json:"strMeasure15,omitempty"`
	StrMeasure16 *string `json:"strMeasure16,omitempty"`
	StrMeasure17 *string `json:"strMeasure17,omitempty"`
	StrMeasure18 *string `json:"strMeasure18,omitempty"`
	StrMeasure19 *string `json:"strMeasure19,omitempty"`
	StrMeasure2 *string `json:"strMeasure2,omitempty"`
	StrMeasure20 *string `json:"strMeasure20,omitempty"`
	StrMeasure3 *string `json:"strMeasure3,omitempty"`
	StrMeasure4 *string `json:"strMeasure4,omitempty"`
	StrMeasure5 *string `json:"strMeasure5,omitempty"`
	StrMeasure6 *string `json:"strMeasure6,omitempty"`
	StrMeasure7 *string `json:"strMeasure7,omitempty"`
	StrMeasure8 *string `json:"strMeasure8,omitempty"`
	StrMeasure9 *string `json:"strMeasure9,omitempty"`
	StrSource *string `json:"strSource,omitempty"`
	StrTags *string `json:"strTags,omitempty"`
	StrYoutube *string `json:"strYoutube,omitempty"`
}

// LatestListMatch is the typed request payload for Latest.ListTyped.
type LatestListMatch struct {
	DateModified *string `json:"dateModified,omitempty"`
	IdMeal *string `json:"idMeal,omitempty"`
	StrArea *string `json:"strArea,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"strCreativeCommonsConfirmed,omitempty"`
	StrDrinkAlternate *string `json:"strDrinkAlternate,omitempty"`
	StrImageSource *string `json:"strImageSource,omitempty"`
	StrIngredient1 *string `json:"strIngredient1,omitempty"`
	StrIngredient10 *string `json:"strIngredient10,omitempty"`
	StrIngredient11 *string `json:"strIngredient11,omitempty"`
	StrIngredient12 *string `json:"strIngredient12,omitempty"`
	StrIngredient13 *string `json:"strIngredient13,omitempty"`
	StrIngredient14 *string `json:"strIngredient14,omitempty"`
	StrIngredient15 *string `json:"strIngredient15,omitempty"`
	StrIngredient16 *string `json:"strIngredient16,omitempty"`
	StrIngredient17 *string `json:"strIngredient17,omitempty"`
	StrIngredient18 *string `json:"strIngredient18,omitempty"`
	StrIngredient19 *string `json:"strIngredient19,omitempty"`
	StrIngredient2 *string `json:"strIngredient2,omitempty"`
	StrIngredient20 *string `json:"strIngredient20,omitempty"`
	StrIngredient3 *string `json:"strIngredient3,omitempty"`
	StrIngredient4 *string `json:"strIngredient4,omitempty"`
	StrIngredient5 *string `json:"strIngredient5,omitempty"`
	StrIngredient6 *string `json:"strIngredient6,omitempty"`
	StrIngredient7 *string `json:"strIngredient7,omitempty"`
	StrIngredient8 *string `json:"strIngredient8,omitempty"`
	StrIngredient9 *string `json:"strIngredient9,omitempty"`
	StrInstructions *string `json:"strInstructions,omitempty"`
	StrMeal *string `json:"strMeal,omitempty"`
	StrMealThumb *string `json:"strMealThumb,omitempty"`
	StrMeasure1 *string `json:"strMeasure1,omitempty"`
	StrMeasure10 *string `json:"strMeasure10,omitempty"`
	StrMeasure11 *string `json:"strMeasure11,omitempty"`
	StrMeasure12 *string `json:"strMeasure12,omitempty"`
	StrMeasure13 *string `json:"strMeasure13,omitempty"`
	StrMeasure14 *string `json:"strMeasure14,omitempty"`
	StrMeasure15 *string `json:"strMeasure15,omitempty"`
	StrMeasure16 *string `json:"strMeasure16,omitempty"`
	StrMeasure17 *string `json:"strMeasure17,omitempty"`
	StrMeasure18 *string `json:"strMeasure18,omitempty"`
	StrMeasure19 *string `json:"strMeasure19,omitempty"`
	StrMeasure2 *string `json:"strMeasure2,omitempty"`
	StrMeasure20 *string `json:"strMeasure20,omitempty"`
	StrMeasure3 *string `json:"strMeasure3,omitempty"`
	StrMeasure4 *string `json:"strMeasure4,omitempty"`
	StrMeasure5 *string `json:"strMeasure5,omitempty"`
	StrMeasure6 *string `json:"strMeasure6,omitempty"`
	StrMeasure7 *string `json:"strMeasure7,omitempty"`
	StrMeasure8 *string `json:"strMeasure8,omitempty"`
	StrMeasure9 *string `json:"strMeasure9,omitempty"`
	StrSource *string `json:"strSource,omitempty"`
	StrTags *string `json:"strTags,omitempty"`
	StrYoutube *string `json:"strYoutube,omitempty"`
}

// List is the typed data model for the list entity.
type List struct {
	StrArea *string `json:"strArea,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrIngredient *string `json:"strIngredient,omitempty"`
}

// ListListMatch is the typed request payload for List.ListTyped.
type ListListMatch struct {
	A *string `json:"a,omitempty"`
	C *string `json:"c,omitempty"`
	I *string `json:"i,omitempty"`
}

// Lookup is the typed data model for the lookup entity.
type Lookup struct {
	DateModified *string `json:"dateModified,omitempty"`
	IdMeal *string `json:"idMeal,omitempty"`
	StrArea *string `json:"strArea,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"strCreativeCommonsConfirmed,omitempty"`
	StrDrinkAlternate *string `json:"strDrinkAlternate,omitempty"`
	StrImageSource *string `json:"strImageSource,omitempty"`
	StrIngredient1 *string `json:"strIngredient1,omitempty"`
	StrIngredient10 *string `json:"strIngredient10,omitempty"`
	StrIngredient11 *string `json:"strIngredient11,omitempty"`
	StrIngredient12 *string `json:"strIngredient12,omitempty"`
	StrIngredient13 *string `json:"strIngredient13,omitempty"`
	StrIngredient14 *string `json:"strIngredient14,omitempty"`
	StrIngredient15 *string `json:"strIngredient15,omitempty"`
	StrIngredient16 *string `json:"strIngredient16,omitempty"`
	StrIngredient17 *string `json:"strIngredient17,omitempty"`
	StrIngredient18 *string `json:"strIngredient18,omitempty"`
	StrIngredient19 *string `json:"strIngredient19,omitempty"`
	StrIngredient2 *string `json:"strIngredient2,omitempty"`
	StrIngredient20 *string `json:"strIngredient20,omitempty"`
	StrIngredient3 *string `json:"strIngredient3,omitempty"`
	StrIngredient4 *string `json:"strIngredient4,omitempty"`
	StrIngredient5 *string `json:"strIngredient5,omitempty"`
	StrIngredient6 *string `json:"strIngredient6,omitempty"`
	StrIngredient7 *string `json:"strIngredient7,omitempty"`
	StrIngredient8 *string `json:"strIngredient8,omitempty"`
	StrIngredient9 *string `json:"strIngredient9,omitempty"`
	StrInstructions *string `json:"strInstructions,omitempty"`
	StrMeal *string `json:"strMeal,omitempty"`
	StrMealThumb *string `json:"strMealThumb,omitempty"`
	StrMeasure1 *string `json:"strMeasure1,omitempty"`
	StrMeasure10 *string `json:"strMeasure10,omitempty"`
	StrMeasure11 *string `json:"strMeasure11,omitempty"`
	StrMeasure12 *string `json:"strMeasure12,omitempty"`
	StrMeasure13 *string `json:"strMeasure13,omitempty"`
	StrMeasure14 *string `json:"strMeasure14,omitempty"`
	StrMeasure15 *string `json:"strMeasure15,omitempty"`
	StrMeasure16 *string `json:"strMeasure16,omitempty"`
	StrMeasure17 *string `json:"strMeasure17,omitempty"`
	StrMeasure18 *string `json:"strMeasure18,omitempty"`
	StrMeasure19 *string `json:"strMeasure19,omitempty"`
	StrMeasure2 *string `json:"strMeasure2,omitempty"`
	StrMeasure20 *string `json:"strMeasure20,omitempty"`
	StrMeasure3 *string `json:"strMeasure3,omitempty"`
	StrMeasure4 *string `json:"strMeasure4,omitempty"`
	StrMeasure5 *string `json:"strMeasure5,omitempty"`
	StrMeasure6 *string `json:"strMeasure6,omitempty"`
	StrMeasure7 *string `json:"strMeasure7,omitempty"`
	StrMeasure8 *string `json:"strMeasure8,omitempty"`
	StrMeasure9 *string `json:"strMeasure9,omitempty"`
	StrSource *string `json:"strSource,omitempty"`
	StrTags *string `json:"strTags,omitempty"`
	StrYoutube *string `json:"strYoutube,omitempty"`
}

// LookupListMatch is the typed request payload for Lookup.ListTyped.
type LookupListMatch struct {
	I string `json:"i"`
}

// Random is the typed data model for the random entity.
type Random struct {
	DateModified *string `json:"dateModified,omitempty"`
	IdMeal *string `json:"idMeal,omitempty"`
	StrArea *string `json:"strArea,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"strCreativeCommonsConfirmed,omitempty"`
	StrDrinkAlternate *string `json:"strDrinkAlternate,omitempty"`
	StrImageSource *string `json:"strImageSource,omitempty"`
	StrIngredient1 *string `json:"strIngredient1,omitempty"`
	StrIngredient10 *string `json:"strIngredient10,omitempty"`
	StrIngredient11 *string `json:"strIngredient11,omitempty"`
	StrIngredient12 *string `json:"strIngredient12,omitempty"`
	StrIngredient13 *string `json:"strIngredient13,omitempty"`
	StrIngredient14 *string `json:"strIngredient14,omitempty"`
	StrIngredient15 *string `json:"strIngredient15,omitempty"`
	StrIngredient16 *string `json:"strIngredient16,omitempty"`
	StrIngredient17 *string `json:"strIngredient17,omitempty"`
	StrIngredient18 *string `json:"strIngredient18,omitempty"`
	StrIngredient19 *string `json:"strIngredient19,omitempty"`
	StrIngredient2 *string `json:"strIngredient2,omitempty"`
	StrIngredient20 *string `json:"strIngredient20,omitempty"`
	StrIngredient3 *string `json:"strIngredient3,omitempty"`
	StrIngredient4 *string `json:"strIngredient4,omitempty"`
	StrIngredient5 *string `json:"strIngredient5,omitempty"`
	StrIngredient6 *string `json:"strIngredient6,omitempty"`
	StrIngredient7 *string `json:"strIngredient7,omitempty"`
	StrIngredient8 *string `json:"strIngredient8,omitempty"`
	StrIngredient9 *string `json:"strIngredient9,omitempty"`
	StrInstructions *string `json:"strInstructions,omitempty"`
	StrMeal *string `json:"strMeal,omitempty"`
	StrMealThumb *string `json:"strMealThumb,omitempty"`
	StrMeasure1 *string `json:"strMeasure1,omitempty"`
	StrMeasure10 *string `json:"strMeasure10,omitempty"`
	StrMeasure11 *string `json:"strMeasure11,omitempty"`
	StrMeasure12 *string `json:"strMeasure12,omitempty"`
	StrMeasure13 *string `json:"strMeasure13,omitempty"`
	StrMeasure14 *string `json:"strMeasure14,omitempty"`
	StrMeasure15 *string `json:"strMeasure15,omitempty"`
	StrMeasure16 *string `json:"strMeasure16,omitempty"`
	StrMeasure17 *string `json:"strMeasure17,omitempty"`
	StrMeasure18 *string `json:"strMeasure18,omitempty"`
	StrMeasure19 *string `json:"strMeasure19,omitempty"`
	StrMeasure2 *string `json:"strMeasure2,omitempty"`
	StrMeasure20 *string `json:"strMeasure20,omitempty"`
	StrMeasure3 *string `json:"strMeasure3,omitempty"`
	StrMeasure4 *string `json:"strMeasure4,omitempty"`
	StrMeasure5 *string `json:"strMeasure5,omitempty"`
	StrMeasure6 *string `json:"strMeasure6,omitempty"`
	StrMeasure7 *string `json:"strMeasure7,omitempty"`
	StrMeasure8 *string `json:"strMeasure8,omitempty"`
	StrMeasure9 *string `json:"strMeasure9,omitempty"`
	StrSource *string `json:"strSource,omitempty"`
	StrTags *string `json:"strTags,omitempty"`
	StrYoutube *string `json:"strYoutube,omitempty"`
}

// RandomListMatch is the typed request payload for Random.ListTyped.
type RandomListMatch struct {
	DateModified *string `json:"dateModified,omitempty"`
	IdMeal *string `json:"idMeal,omitempty"`
	StrArea *string `json:"strArea,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"strCreativeCommonsConfirmed,omitempty"`
	StrDrinkAlternate *string `json:"strDrinkAlternate,omitempty"`
	StrImageSource *string `json:"strImageSource,omitempty"`
	StrIngredient1 *string `json:"strIngredient1,omitempty"`
	StrIngredient10 *string `json:"strIngredient10,omitempty"`
	StrIngredient11 *string `json:"strIngredient11,omitempty"`
	StrIngredient12 *string `json:"strIngredient12,omitempty"`
	StrIngredient13 *string `json:"strIngredient13,omitempty"`
	StrIngredient14 *string `json:"strIngredient14,omitempty"`
	StrIngredient15 *string `json:"strIngredient15,omitempty"`
	StrIngredient16 *string `json:"strIngredient16,omitempty"`
	StrIngredient17 *string `json:"strIngredient17,omitempty"`
	StrIngredient18 *string `json:"strIngredient18,omitempty"`
	StrIngredient19 *string `json:"strIngredient19,omitempty"`
	StrIngredient2 *string `json:"strIngredient2,omitempty"`
	StrIngredient20 *string `json:"strIngredient20,omitempty"`
	StrIngredient3 *string `json:"strIngredient3,omitempty"`
	StrIngredient4 *string `json:"strIngredient4,omitempty"`
	StrIngredient5 *string `json:"strIngredient5,omitempty"`
	StrIngredient6 *string `json:"strIngredient6,omitempty"`
	StrIngredient7 *string `json:"strIngredient7,omitempty"`
	StrIngredient8 *string `json:"strIngredient8,omitempty"`
	StrIngredient9 *string `json:"strIngredient9,omitempty"`
	StrInstructions *string `json:"strInstructions,omitempty"`
	StrMeal *string `json:"strMeal,omitempty"`
	StrMealThumb *string `json:"strMealThumb,omitempty"`
	StrMeasure1 *string `json:"strMeasure1,omitempty"`
	StrMeasure10 *string `json:"strMeasure10,omitempty"`
	StrMeasure11 *string `json:"strMeasure11,omitempty"`
	StrMeasure12 *string `json:"strMeasure12,omitempty"`
	StrMeasure13 *string `json:"strMeasure13,omitempty"`
	StrMeasure14 *string `json:"strMeasure14,omitempty"`
	StrMeasure15 *string `json:"strMeasure15,omitempty"`
	StrMeasure16 *string `json:"strMeasure16,omitempty"`
	StrMeasure17 *string `json:"strMeasure17,omitempty"`
	StrMeasure18 *string `json:"strMeasure18,omitempty"`
	StrMeasure19 *string `json:"strMeasure19,omitempty"`
	StrMeasure2 *string `json:"strMeasure2,omitempty"`
	StrMeasure20 *string `json:"strMeasure20,omitempty"`
	StrMeasure3 *string `json:"strMeasure3,omitempty"`
	StrMeasure4 *string `json:"strMeasure4,omitempty"`
	StrMeasure5 *string `json:"strMeasure5,omitempty"`
	StrMeasure6 *string `json:"strMeasure6,omitempty"`
	StrMeasure7 *string `json:"strMeasure7,omitempty"`
	StrMeasure8 *string `json:"strMeasure8,omitempty"`
	StrMeasure9 *string `json:"strMeasure9,omitempty"`
	StrSource *string `json:"strSource,omitempty"`
	StrTags *string `json:"strTags,omitempty"`
	StrYoutube *string `json:"strYoutube,omitempty"`
}

// Randomselection is the typed data model for the randomselection entity.
type Randomselection struct {
	DateModified *string `json:"dateModified,omitempty"`
	IdMeal *string `json:"idMeal,omitempty"`
	StrArea *string `json:"strArea,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"strCreativeCommonsConfirmed,omitempty"`
	StrDrinkAlternate *string `json:"strDrinkAlternate,omitempty"`
	StrImageSource *string `json:"strImageSource,omitempty"`
	StrIngredient1 *string `json:"strIngredient1,omitempty"`
	StrIngredient10 *string `json:"strIngredient10,omitempty"`
	StrIngredient11 *string `json:"strIngredient11,omitempty"`
	StrIngredient12 *string `json:"strIngredient12,omitempty"`
	StrIngredient13 *string `json:"strIngredient13,omitempty"`
	StrIngredient14 *string `json:"strIngredient14,omitempty"`
	StrIngredient15 *string `json:"strIngredient15,omitempty"`
	StrIngredient16 *string `json:"strIngredient16,omitempty"`
	StrIngredient17 *string `json:"strIngredient17,omitempty"`
	StrIngredient18 *string `json:"strIngredient18,omitempty"`
	StrIngredient19 *string `json:"strIngredient19,omitempty"`
	StrIngredient2 *string `json:"strIngredient2,omitempty"`
	StrIngredient20 *string `json:"strIngredient20,omitempty"`
	StrIngredient3 *string `json:"strIngredient3,omitempty"`
	StrIngredient4 *string `json:"strIngredient4,omitempty"`
	StrIngredient5 *string `json:"strIngredient5,omitempty"`
	StrIngredient6 *string `json:"strIngredient6,omitempty"`
	StrIngredient7 *string `json:"strIngredient7,omitempty"`
	StrIngredient8 *string `json:"strIngredient8,omitempty"`
	StrIngredient9 *string `json:"strIngredient9,omitempty"`
	StrInstructions *string `json:"strInstructions,omitempty"`
	StrMeal *string `json:"strMeal,omitempty"`
	StrMealThumb *string `json:"strMealThumb,omitempty"`
	StrMeasure1 *string `json:"strMeasure1,omitempty"`
	StrMeasure10 *string `json:"strMeasure10,omitempty"`
	StrMeasure11 *string `json:"strMeasure11,omitempty"`
	StrMeasure12 *string `json:"strMeasure12,omitempty"`
	StrMeasure13 *string `json:"strMeasure13,omitempty"`
	StrMeasure14 *string `json:"strMeasure14,omitempty"`
	StrMeasure15 *string `json:"strMeasure15,omitempty"`
	StrMeasure16 *string `json:"strMeasure16,omitempty"`
	StrMeasure17 *string `json:"strMeasure17,omitempty"`
	StrMeasure18 *string `json:"strMeasure18,omitempty"`
	StrMeasure19 *string `json:"strMeasure19,omitempty"`
	StrMeasure2 *string `json:"strMeasure2,omitempty"`
	StrMeasure20 *string `json:"strMeasure20,omitempty"`
	StrMeasure3 *string `json:"strMeasure3,omitempty"`
	StrMeasure4 *string `json:"strMeasure4,omitempty"`
	StrMeasure5 *string `json:"strMeasure5,omitempty"`
	StrMeasure6 *string `json:"strMeasure6,omitempty"`
	StrMeasure7 *string `json:"strMeasure7,omitempty"`
	StrMeasure8 *string `json:"strMeasure8,omitempty"`
	StrMeasure9 *string `json:"strMeasure9,omitempty"`
	StrSource *string `json:"strSource,omitempty"`
	StrTags *string `json:"strTags,omitempty"`
	StrYoutube *string `json:"strYoutube,omitempty"`
}

// RandomselectionListMatch is the typed request payload for Randomselection.ListTyped.
type RandomselectionListMatch struct {
	DateModified *string `json:"dateModified,omitempty"`
	IdMeal *string `json:"idMeal,omitempty"`
	StrArea *string `json:"strArea,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"strCreativeCommonsConfirmed,omitempty"`
	StrDrinkAlternate *string `json:"strDrinkAlternate,omitempty"`
	StrImageSource *string `json:"strImageSource,omitempty"`
	StrIngredient1 *string `json:"strIngredient1,omitempty"`
	StrIngredient10 *string `json:"strIngredient10,omitempty"`
	StrIngredient11 *string `json:"strIngredient11,omitempty"`
	StrIngredient12 *string `json:"strIngredient12,omitempty"`
	StrIngredient13 *string `json:"strIngredient13,omitempty"`
	StrIngredient14 *string `json:"strIngredient14,omitempty"`
	StrIngredient15 *string `json:"strIngredient15,omitempty"`
	StrIngredient16 *string `json:"strIngredient16,omitempty"`
	StrIngredient17 *string `json:"strIngredient17,omitempty"`
	StrIngredient18 *string `json:"strIngredient18,omitempty"`
	StrIngredient19 *string `json:"strIngredient19,omitempty"`
	StrIngredient2 *string `json:"strIngredient2,omitempty"`
	StrIngredient20 *string `json:"strIngredient20,omitempty"`
	StrIngredient3 *string `json:"strIngredient3,omitempty"`
	StrIngredient4 *string `json:"strIngredient4,omitempty"`
	StrIngredient5 *string `json:"strIngredient5,omitempty"`
	StrIngredient6 *string `json:"strIngredient6,omitempty"`
	StrIngredient7 *string `json:"strIngredient7,omitempty"`
	StrIngredient8 *string `json:"strIngredient8,omitempty"`
	StrIngredient9 *string `json:"strIngredient9,omitempty"`
	StrInstructions *string `json:"strInstructions,omitempty"`
	StrMeal *string `json:"strMeal,omitempty"`
	StrMealThumb *string `json:"strMealThumb,omitempty"`
	StrMeasure1 *string `json:"strMeasure1,omitempty"`
	StrMeasure10 *string `json:"strMeasure10,omitempty"`
	StrMeasure11 *string `json:"strMeasure11,omitempty"`
	StrMeasure12 *string `json:"strMeasure12,omitempty"`
	StrMeasure13 *string `json:"strMeasure13,omitempty"`
	StrMeasure14 *string `json:"strMeasure14,omitempty"`
	StrMeasure15 *string `json:"strMeasure15,omitempty"`
	StrMeasure16 *string `json:"strMeasure16,omitempty"`
	StrMeasure17 *string `json:"strMeasure17,omitempty"`
	StrMeasure18 *string `json:"strMeasure18,omitempty"`
	StrMeasure19 *string `json:"strMeasure19,omitempty"`
	StrMeasure2 *string `json:"strMeasure2,omitempty"`
	StrMeasure20 *string `json:"strMeasure20,omitempty"`
	StrMeasure3 *string `json:"strMeasure3,omitempty"`
	StrMeasure4 *string `json:"strMeasure4,omitempty"`
	StrMeasure5 *string `json:"strMeasure5,omitempty"`
	StrMeasure6 *string `json:"strMeasure6,omitempty"`
	StrMeasure7 *string `json:"strMeasure7,omitempty"`
	StrMeasure8 *string `json:"strMeasure8,omitempty"`
	StrMeasure9 *string `json:"strMeasure9,omitempty"`
	StrSource *string `json:"strSource,omitempty"`
	StrTags *string `json:"strTags,omitempty"`
	StrYoutube *string `json:"strYoutube,omitempty"`
}

// Search is the typed data model for the search entity.
type Search struct {
	DateModified *string `json:"dateModified,omitempty"`
	IdMeal *string `json:"idMeal,omitempty"`
	StrArea *string `json:"strArea,omitempty"`
	StrCategory *string `json:"strCategory,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"strCreativeCommonsConfirmed,omitempty"`
	StrDrinkAlternate *string `json:"strDrinkAlternate,omitempty"`
	StrImageSource *string `json:"strImageSource,omitempty"`
	StrIngredient1 *string `json:"strIngredient1,omitempty"`
	StrIngredient10 *string `json:"strIngredient10,omitempty"`
	StrIngredient11 *string `json:"strIngredient11,omitempty"`
	StrIngredient12 *string `json:"strIngredient12,omitempty"`
	StrIngredient13 *string `json:"strIngredient13,omitempty"`
	StrIngredient14 *string `json:"strIngredient14,omitempty"`
	StrIngredient15 *string `json:"strIngredient15,omitempty"`
	StrIngredient16 *string `json:"strIngredient16,omitempty"`
	StrIngredient17 *string `json:"strIngredient17,omitempty"`
	StrIngredient18 *string `json:"strIngredient18,omitempty"`
	StrIngredient19 *string `json:"strIngredient19,omitempty"`
	StrIngredient2 *string `json:"strIngredient2,omitempty"`
	StrIngredient20 *string `json:"strIngredient20,omitempty"`
	StrIngredient3 *string `json:"strIngredient3,omitempty"`
	StrIngredient4 *string `json:"strIngredient4,omitempty"`
	StrIngredient5 *string `json:"strIngredient5,omitempty"`
	StrIngredient6 *string `json:"strIngredient6,omitempty"`
	StrIngredient7 *string `json:"strIngredient7,omitempty"`
	StrIngredient8 *string `json:"strIngredient8,omitempty"`
	StrIngredient9 *string `json:"strIngredient9,omitempty"`
	StrInstructions *string `json:"strInstructions,omitempty"`
	StrMeal *string `json:"strMeal,omitempty"`
	StrMealThumb *string `json:"strMealThumb,omitempty"`
	StrMeasure1 *string `json:"strMeasure1,omitempty"`
	StrMeasure10 *string `json:"strMeasure10,omitempty"`
	StrMeasure11 *string `json:"strMeasure11,omitempty"`
	StrMeasure12 *string `json:"strMeasure12,omitempty"`
	StrMeasure13 *string `json:"strMeasure13,omitempty"`
	StrMeasure14 *string `json:"strMeasure14,omitempty"`
	StrMeasure15 *string `json:"strMeasure15,omitempty"`
	StrMeasure16 *string `json:"strMeasure16,omitempty"`
	StrMeasure17 *string `json:"strMeasure17,omitempty"`
	StrMeasure18 *string `json:"strMeasure18,omitempty"`
	StrMeasure19 *string `json:"strMeasure19,omitempty"`
	StrMeasure2 *string `json:"strMeasure2,omitempty"`
	StrMeasure20 *string `json:"strMeasure20,omitempty"`
	StrMeasure3 *string `json:"strMeasure3,omitempty"`
	StrMeasure4 *string `json:"strMeasure4,omitempty"`
	StrMeasure5 *string `json:"strMeasure5,omitempty"`
	StrMeasure6 *string `json:"strMeasure6,omitempty"`
	StrMeasure7 *string `json:"strMeasure7,omitempty"`
	StrMeasure8 *string `json:"strMeasure8,omitempty"`
	StrMeasure9 *string `json:"strMeasure9,omitempty"`
	StrSource *string `json:"strSource,omitempty"`
	StrTags *string `json:"strTags,omitempty"`
	StrYoutube *string `json:"strYoutube,omitempty"`
}

// SearchListMatch is the typed request payload for Search.ListTyped.
type SearchListMatch struct {
	F *string `json:"f,omitempty"`
	S *string `json:"s,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

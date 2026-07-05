// Typed models for the FreeMeal SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Category is the typed data model for the category entity.
type Category struct {
	IdCategory *string `json:"id_category,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCategoryDescription *string `json:"str_category_description,omitempty"`
	StrCategoryThumb *string `json:"str_category_thumb,omitempty"`
}

// CategoryListMatch is the typed request payload for Category.ListTyped.
type CategoryListMatch struct {
	IdCategory *string `json:"id_category,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCategoryDescription *string `json:"str_category_description,omitempty"`
	StrCategoryThumb *string `json:"str_category_thumb,omitempty"`
}

// Filter is the typed data model for the filter entity.
type Filter struct {
	IdMeal *string `json:"id_meal,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
}

// FilterListMatch is the typed request payload for Filter.ListTyped.
type FilterListMatch struct {
	IdMeal *string `json:"id_meal,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
}

// Latest is the typed data model for the latest entity.
type Latest struct {
	DateModified *string `json:"date_modified,omitempty"`
	IdMeal *string `json:"id_meal,omitempty"`
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"str_creative_commons_confirmed,omitempty"`
	StrDrinkAlternate *string `json:"str_drink_alternate,omitempty"`
	StrImageSource *string `json:"str_image_source,omitempty"`
	StrIngredient1 *string `json:"str_ingredient1,omitempty"`
	StrIngredient10 *string `json:"str_ingredient10,omitempty"`
	StrIngredient11 *string `json:"str_ingredient11,omitempty"`
	StrIngredient12 *string `json:"str_ingredient12,omitempty"`
	StrIngredient13 *string `json:"str_ingredient13,omitempty"`
	StrIngredient14 *string `json:"str_ingredient14,omitempty"`
	StrIngredient15 *string `json:"str_ingredient15,omitempty"`
	StrIngredient16 *string `json:"str_ingredient16,omitempty"`
	StrIngredient17 *string `json:"str_ingredient17,omitempty"`
	StrIngredient18 *string `json:"str_ingredient18,omitempty"`
	StrIngredient19 *string `json:"str_ingredient19,omitempty"`
	StrIngredient2 *string `json:"str_ingredient2,omitempty"`
	StrIngredient20 *string `json:"str_ingredient20,omitempty"`
	StrIngredient3 *string `json:"str_ingredient3,omitempty"`
	StrIngredient4 *string `json:"str_ingredient4,omitempty"`
	StrIngredient5 *string `json:"str_ingredient5,omitempty"`
	StrIngredient6 *string `json:"str_ingredient6,omitempty"`
	StrIngredient7 *string `json:"str_ingredient7,omitempty"`
	StrIngredient8 *string `json:"str_ingredient8,omitempty"`
	StrIngredient9 *string `json:"str_ingredient9,omitempty"`
	StrInstruction *string `json:"str_instruction,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
	StrMeasure1 *string `json:"str_measure1,omitempty"`
	StrMeasure10 *string `json:"str_measure10,omitempty"`
	StrMeasure11 *string `json:"str_measure11,omitempty"`
	StrMeasure12 *string `json:"str_measure12,omitempty"`
	StrMeasure13 *string `json:"str_measure13,omitempty"`
	StrMeasure14 *string `json:"str_measure14,omitempty"`
	StrMeasure15 *string `json:"str_measure15,omitempty"`
	StrMeasure16 *string `json:"str_measure16,omitempty"`
	StrMeasure17 *string `json:"str_measure17,omitempty"`
	StrMeasure18 *string `json:"str_measure18,omitempty"`
	StrMeasure19 *string `json:"str_measure19,omitempty"`
	StrMeasure2 *string `json:"str_measure2,omitempty"`
	StrMeasure20 *string `json:"str_measure20,omitempty"`
	StrMeasure3 *string `json:"str_measure3,omitempty"`
	StrMeasure4 *string `json:"str_measure4,omitempty"`
	StrMeasure5 *string `json:"str_measure5,omitempty"`
	StrMeasure6 *string `json:"str_measure6,omitempty"`
	StrMeasure7 *string `json:"str_measure7,omitempty"`
	StrMeasure8 *string `json:"str_measure8,omitempty"`
	StrMeasure9 *string `json:"str_measure9,omitempty"`
	StrSource *string `json:"str_source,omitempty"`
	StrTag *string `json:"str_tag,omitempty"`
	StrYoutube *string `json:"str_youtube,omitempty"`
}

// LatestListMatch is the typed request payload for Latest.ListTyped.
type LatestListMatch struct {
	DateModified *string `json:"date_modified,omitempty"`
	IdMeal *string `json:"id_meal,omitempty"`
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"str_creative_commons_confirmed,omitempty"`
	StrDrinkAlternate *string `json:"str_drink_alternate,omitempty"`
	StrImageSource *string `json:"str_image_source,omitempty"`
	StrIngredient1 *string `json:"str_ingredient1,omitempty"`
	StrIngredient10 *string `json:"str_ingredient10,omitempty"`
	StrIngredient11 *string `json:"str_ingredient11,omitempty"`
	StrIngredient12 *string `json:"str_ingredient12,omitempty"`
	StrIngredient13 *string `json:"str_ingredient13,omitempty"`
	StrIngredient14 *string `json:"str_ingredient14,omitempty"`
	StrIngredient15 *string `json:"str_ingredient15,omitempty"`
	StrIngredient16 *string `json:"str_ingredient16,omitempty"`
	StrIngredient17 *string `json:"str_ingredient17,omitempty"`
	StrIngredient18 *string `json:"str_ingredient18,omitempty"`
	StrIngredient19 *string `json:"str_ingredient19,omitempty"`
	StrIngredient2 *string `json:"str_ingredient2,omitempty"`
	StrIngredient20 *string `json:"str_ingredient20,omitempty"`
	StrIngredient3 *string `json:"str_ingredient3,omitempty"`
	StrIngredient4 *string `json:"str_ingredient4,omitempty"`
	StrIngredient5 *string `json:"str_ingredient5,omitempty"`
	StrIngredient6 *string `json:"str_ingredient6,omitempty"`
	StrIngredient7 *string `json:"str_ingredient7,omitempty"`
	StrIngredient8 *string `json:"str_ingredient8,omitempty"`
	StrIngredient9 *string `json:"str_ingredient9,omitempty"`
	StrInstruction *string `json:"str_instruction,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
	StrMeasure1 *string `json:"str_measure1,omitempty"`
	StrMeasure10 *string `json:"str_measure10,omitempty"`
	StrMeasure11 *string `json:"str_measure11,omitempty"`
	StrMeasure12 *string `json:"str_measure12,omitempty"`
	StrMeasure13 *string `json:"str_measure13,omitempty"`
	StrMeasure14 *string `json:"str_measure14,omitempty"`
	StrMeasure15 *string `json:"str_measure15,omitempty"`
	StrMeasure16 *string `json:"str_measure16,omitempty"`
	StrMeasure17 *string `json:"str_measure17,omitempty"`
	StrMeasure18 *string `json:"str_measure18,omitempty"`
	StrMeasure19 *string `json:"str_measure19,omitempty"`
	StrMeasure2 *string `json:"str_measure2,omitempty"`
	StrMeasure20 *string `json:"str_measure20,omitempty"`
	StrMeasure3 *string `json:"str_measure3,omitempty"`
	StrMeasure4 *string `json:"str_measure4,omitempty"`
	StrMeasure5 *string `json:"str_measure5,omitempty"`
	StrMeasure6 *string `json:"str_measure6,omitempty"`
	StrMeasure7 *string `json:"str_measure7,omitempty"`
	StrMeasure8 *string `json:"str_measure8,omitempty"`
	StrMeasure9 *string `json:"str_measure9,omitempty"`
	StrSource *string `json:"str_source,omitempty"`
	StrTag *string `json:"str_tag,omitempty"`
	StrYoutube *string `json:"str_youtube,omitempty"`
}

// List is the typed data model for the list entity.
type List struct {
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrIngredient *string `json:"str_ingredient,omitempty"`
}

// ListListMatch is the typed request payload for List.ListTyped.
type ListListMatch struct {
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrIngredient *string `json:"str_ingredient,omitempty"`
}

// Lookup is the typed data model for the lookup entity.
type Lookup struct {
	DateModified *string `json:"date_modified,omitempty"`
	IdMeal *string `json:"id_meal,omitempty"`
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"str_creative_commons_confirmed,omitempty"`
	StrDrinkAlternate *string `json:"str_drink_alternate,omitempty"`
	StrImageSource *string `json:"str_image_source,omitempty"`
	StrIngredient1 *string `json:"str_ingredient1,omitempty"`
	StrIngredient10 *string `json:"str_ingredient10,omitempty"`
	StrIngredient11 *string `json:"str_ingredient11,omitempty"`
	StrIngredient12 *string `json:"str_ingredient12,omitempty"`
	StrIngredient13 *string `json:"str_ingredient13,omitempty"`
	StrIngredient14 *string `json:"str_ingredient14,omitempty"`
	StrIngredient15 *string `json:"str_ingredient15,omitempty"`
	StrIngredient16 *string `json:"str_ingredient16,omitempty"`
	StrIngredient17 *string `json:"str_ingredient17,omitempty"`
	StrIngredient18 *string `json:"str_ingredient18,omitempty"`
	StrIngredient19 *string `json:"str_ingredient19,omitempty"`
	StrIngredient2 *string `json:"str_ingredient2,omitempty"`
	StrIngredient20 *string `json:"str_ingredient20,omitempty"`
	StrIngredient3 *string `json:"str_ingredient3,omitempty"`
	StrIngredient4 *string `json:"str_ingredient4,omitempty"`
	StrIngredient5 *string `json:"str_ingredient5,omitempty"`
	StrIngredient6 *string `json:"str_ingredient6,omitempty"`
	StrIngredient7 *string `json:"str_ingredient7,omitempty"`
	StrIngredient8 *string `json:"str_ingredient8,omitempty"`
	StrIngredient9 *string `json:"str_ingredient9,omitempty"`
	StrInstruction *string `json:"str_instruction,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
	StrMeasure1 *string `json:"str_measure1,omitempty"`
	StrMeasure10 *string `json:"str_measure10,omitempty"`
	StrMeasure11 *string `json:"str_measure11,omitempty"`
	StrMeasure12 *string `json:"str_measure12,omitempty"`
	StrMeasure13 *string `json:"str_measure13,omitempty"`
	StrMeasure14 *string `json:"str_measure14,omitempty"`
	StrMeasure15 *string `json:"str_measure15,omitempty"`
	StrMeasure16 *string `json:"str_measure16,omitempty"`
	StrMeasure17 *string `json:"str_measure17,omitempty"`
	StrMeasure18 *string `json:"str_measure18,omitempty"`
	StrMeasure19 *string `json:"str_measure19,omitempty"`
	StrMeasure2 *string `json:"str_measure2,omitempty"`
	StrMeasure20 *string `json:"str_measure20,omitempty"`
	StrMeasure3 *string `json:"str_measure3,omitempty"`
	StrMeasure4 *string `json:"str_measure4,omitempty"`
	StrMeasure5 *string `json:"str_measure5,omitempty"`
	StrMeasure6 *string `json:"str_measure6,omitempty"`
	StrMeasure7 *string `json:"str_measure7,omitempty"`
	StrMeasure8 *string `json:"str_measure8,omitempty"`
	StrMeasure9 *string `json:"str_measure9,omitempty"`
	StrSource *string `json:"str_source,omitempty"`
	StrTag *string `json:"str_tag,omitempty"`
	StrYoutube *string `json:"str_youtube,omitempty"`
}

// LookupListMatch is the typed request payload for Lookup.ListTyped.
type LookupListMatch struct {
	DateModified *string `json:"date_modified,omitempty"`
	IdMeal *string `json:"id_meal,omitempty"`
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"str_creative_commons_confirmed,omitempty"`
	StrDrinkAlternate *string `json:"str_drink_alternate,omitempty"`
	StrImageSource *string `json:"str_image_source,omitempty"`
	StrIngredient1 *string `json:"str_ingredient1,omitempty"`
	StrIngredient10 *string `json:"str_ingredient10,omitempty"`
	StrIngredient11 *string `json:"str_ingredient11,omitempty"`
	StrIngredient12 *string `json:"str_ingredient12,omitempty"`
	StrIngredient13 *string `json:"str_ingredient13,omitempty"`
	StrIngredient14 *string `json:"str_ingredient14,omitempty"`
	StrIngredient15 *string `json:"str_ingredient15,omitempty"`
	StrIngredient16 *string `json:"str_ingredient16,omitempty"`
	StrIngredient17 *string `json:"str_ingredient17,omitempty"`
	StrIngredient18 *string `json:"str_ingredient18,omitempty"`
	StrIngredient19 *string `json:"str_ingredient19,omitempty"`
	StrIngredient2 *string `json:"str_ingredient2,omitempty"`
	StrIngredient20 *string `json:"str_ingredient20,omitempty"`
	StrIngredient3 *string `json:"str_ingredient3,omitempty"`
	StrIngredient4 *string `json:"str_ingredient4,omitempty"`
	StrIngredient5 *string `json:"str_ingredient5,omitempty"`
	StrIngredient6 *string `json:"str_ingredient6,omitempty"`
	StrIngredient7 *string `json:"str_ingredient7,omitempty"`
	StrIngredient8 *string `json:"str_ingredient8,omitempty"`
	StrIngredient9 *string `json:"str_ingredient9,omitempty"`
	StrInstruction *string `json:"str_instruction,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
	StrMeasure1 *string `json:"str_measure1,omitempty"`
	StrMeasure10 *string `json:"str_measure10,omitempty"`
	StrMeasure11 *string `json:"str_measure11,omitempty"`
	StrMeasure12 *string `json:"str_measure12,omitempty"`
	StrMeasure13 *string `json:"str_measure13,omitempty"`
	StrMeasure14 *string `json:"str_measure14,omitempty"`
	StrMeasure15 *string `json:"str_measure15,omitempty"`
	StrMeasure16 *string `json:"str_measure16,omitempty"`
	StrMeasure17 *string `json:"str_measure17,omitempty"`
	StrMeasure18 *string `json:"str_measure18,omitempty"`
	StrMeasure19 *string `json:"str_measure19,omitempty"`
	StrMeasure2 *string `json:"str_measure2,omitempty"`
	StrMeasure20 *string `json:"str_measure20,omitempty"`
	StrMeasure3 *string `json:"str_measure3,omitempty"`
	StrMeasure4 *string `json:"str_measure4,omitempty"`
	StrMeasure5 *string `json:"str_measure5,omitempty"`
	StrMeasure6 *string `json:"str_measure6,omitempty"`
	StrMeasure7 *string `json:"str_measure7,omitempty"`
	StrMeasure8 *string `json:"str_measure8,omitempty"`
	StrMeasure9 *string `json:"str_measure9,omitempty"`
	StrSource *string `json:"str_source,omitempty"`
	StrTag *string `json:"str_tag,omitempty"`
	StrYoutube *string `json:"str_youtube,omitempty"`
}

// Random is the typed data model for the random entity.
type Random struct {
	DateModified *string `json:"date_modified,omitempty"`
	IdMeal *string `json:"id_meal,omitempty"`
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"str_creative_commons_confirmed,omitempty"`
	StrDrinkAlternate *string `json:"str_drink_alternate,omitempty"`
	StrImageSource *string `json:"str_image_source,omitempty"`
	StrIngredient1 *string `json:"str_ingredient1,omitempty"`
	StrIngredient10 *string `json:"str_ingredient10,omitempty"`
	StrIngredient11 *string `json:"str_ingredient11,omitempty"`
	StrIngredient12 *string `json:"str_ingredient12,omitempty"`
	StrIngredient13 *string `json:"str_ingredient13,omitempty"`
	StrIngredient14 *string `json:"str_ingredient14,omitempty"`
	StrIngredient15 *string `json:"str_ingredient15,omitempty"`
	StrIngredient16 *string `json:"str_ingredient16,omitempty"`
	StrIngredient17 *string `json:"str_ingredient17,omitempty"`
	StrIngredient18 *string `json:"str_ingredient18,omitempty"`
	StrIngredient19 *string `json:"str_ingredient19,omitempty"`
	StrIngredient2 *string `json:"str_ingredient2,omitempty"`
	StrIngredient20 *string `json:"str_ingredient20,omitempty"`
	StrIngredient3 *string `json:"str_ingredient3,omitempty"`
	StrIngredient4 *string `json:"str_ingredient4,omitempty"`
	StrIngredient5 *string `json:"str_ingredient5,omitempty"`
	StrIngredient6 *string `json:"str_ingredient6,omitempty"`
	StrIngredient7 *string `json:"str_ingredient7,omitempty"`
	StrIngredient8 *string `json:"str_ingredient8,omitempty"`
	StrIngredient9 *string `json:"str_ingredient9,omitempty"`
	StrInstruction *string `json:"str_instruction,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
	StrMeasure1 *string `json:"str_measure1,omitempty"`
	StrMeasure10 *string `json:"str_measure10,omitempty"`
	StrMeasure11 *string `json:"str_measure11,omitempty"`
	StrMeasure12 *string `json:"str_measure12,omitempty"`
	StrMeasure13 *string `json:"str_measure13,omitempty"`
	StrMeasure14 *string `json:"str_measure14,omitempty"`
	StrMeasure15 *string `json:"str_measure15,omitempty"`
	StrMeasure16 *string `json:"str_measure16,omitempty"`
	StrMeasure17 *string `json:"str_measure17,omitempty"`
	StrMeasure18 *string `json:"str_measure18,omitempty"`
	StrMeasure19 *string `json:"str_measure19,omitempty"`
	StrMeasure2 *string `json:"str_measure2,omitempty"`
	StrMeasure20 *string `json:"str_measure20,omitempty"`
	StrMeasure3 *string `json:"str_measure3,omitempty"`
	StrMeasure4 *string `json:"str_measure4,omitempty"`
	StrMeasure5 *string `json:"str_measure5,omitempty"`
	StrMeasure6 *string `json:"str_measure6,omitempty"`
	StrMeasure7 *string `json:"str_measure7,omitempty"`
	StrMeasure8 *string `json:"str_measure8,omitempty"`
	StrMeasure9 *string `json:"str_measure9,omitempty"`
	StrSource *string `json:"str_source,omitempty"`
	StrTag *string `json:"str_tag,omitempty"`
	StrYoutube *string `json:"str_youtube,omitempty"`
}

// RandomListMatch is the typed request payload for Random.ListTyped.
type RandomListMatch struct {
	DateModified *string `json:"date_modified,omitempty"`
	IdMeal *string `json:"id_meal,omitempty"`
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"str_creative_commons_confirmed,omitempty"`
	StrDrinkAlternate *string `json:"str_drink_alternate,omitempty"`
	StrImageSource *string `json:"str_image_source,omitempty"`
	StrIngredient1 *string `json:"str_ingredient1,omitempty"`
	StrIngredient10 *string `json:"str_ingredient10,omitempty"`
	StrIngredient11 *string `json:"str_ingredient11,omitempty"`
	StrIngredient12 *string `json:"str_ingredient12,omitempty"`
	StrIngredient13 *string `json:"str_ingredient13,omitempty"`
	StrIngredient14 *string `json:"str_ingredient14,omitempty"`
	StrIngredient15 *string `json:"str_ingredient15,omitempty"`
	StrIngredient16 *string `json:"str_ingredient16,omitempty"`
	StrIngredient17 *string `json:"str_ingredient17,omitempty"`
	StrIngredient18 *string `json:"str_ingredient18,omitempty"`
	StrIngredient19 *string `json:"str_ingredient19,omitempty"`
	StrIngredient2 *string `json:"str_ingredient2,omitempty"`
	StrIngredient20 *string `json:"str_ingredient20,omitempty"`
	StrIngredient3 *string `json:"str_ingredient3,omitempty"`
	StrIngredient4 *string `json:"str_ingredient4,omitempty"`
	StrIngredient5 *string `json:"str_ingredient5,omitempty"`
	StrIngredient6 *string `json:"str_ingredient6,omitempty"`
	StrIngredient7 *string `json:"str_ingredient7,omitempty"`
	StrIngredient8 *string `json:"str_ingredient8,omitempty"`
	StrIngredient9 *string `json:"str_ingredient9,omitempty"`
	StrInstruction *string `json:"str_instruction,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
	StrMeasure1 *string `json:"str_measure1,omitempty"`
	StrMeasure10 *string `json:"str_measure10,omitempty"`
	StrMeasure11 *string `json:"str_measure11,omitempty"`
	StrMeasure12 *string `json:"str_measure12,omitempty"`
	StrMeasure13 *string `json:"str_measure13,omitempty"`
	StrMeasure14 *string `json:"str_measure14,omitempty"`
	StrMeasure15 *string `json:"str_measure15,omitempty"`
	StrMeasure16 *string `json:"str_measure16,omitempty"`
	StrMeasure17 *string `json:"str_measure17,omitempty"`
	StrMeasure18 *string `json:"str_measure18,omitempty"`
	StrMeasure19 *string `json:"str_measure19,omitempty"`
	StrMeasure2 *string `json:"str_measure2,omitempty"`
	StrMeasure20 *string `json:"str_measure20,omitempty"`
	StrMeasure3 *string `json:"str_measure3,omitempty"`
	StrMeasure4 *string `json:"str_measure4,omitempty"`
	StrMeasure5 *string `json:"str_measure5,omitempty"`
	StrMeasure6 *string `json:"str_measure6,omitempty"`
	StrMeasure7 *string `json:"str_measure7,omitempty"`
	StrMeasure8 *string `json:"str_measure8,omitempty"`
	StrMeasure9 *string `json:"str_measure9,omitempty"`
	StrSource *string `json:"str_source,omitempty"`
	StrTag *string `json:"str_tag,omitempty"`
	StrYoutube *string `json:"str_youtube,omitempty"`
}

// Randomselection is the typed data model for the randomselection entity.
type Randomselection struct {
	DateModified *string `json:"date_modified,omitempty"`
	IdMeal *string `json:"id_meal,omitempty"`
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"str_creative_commons_confirmed,omitempty"`
	StrDrinkAlternate *string `json:"str_drink_alternate,omitempty"`
	StrImageSource *string `json:"str_image_source,omitempty"`
	StrIngredient1 *string `json:"str_ingredient1,omitempty"`
	StrIngredient10 *string `json:"str_ingredient10,omitempty"`
	StrIngredient11 *string `json:"str_ingredient11,omitempty"`
	StrIngredient12 *string `json:"str_ingredient12,omitempty"`
	StrIngredient13 *string `json:"str_ingredient13,omitempty"`
	StrIngredient14 *string `json:"str_ingredient14,omitempty"`
	StrIngredient15 *string `json:"str_ingredient15,omitempty"`
	StrIngredient16 *string `json:"str_ingredient16,omitempty"`
	StrIngredient17 *string `json:"str_ingredient17,omitempty"`
	StrIngredient18 *string `json:"str_ingredient18,omitempty"`
	StrIngredient19 *string `json:"str_ingredient19,omitempty"`
	StrIngredient2 *string `json:"str_ingredient2,omitempty"`
	StrIngredient20 *string `json:"str_ingredient20,omitempty"`
	StrIngredient3 *string `json:"str_ingredient3,omitempty"`
	StrIngredient4 *string `json:"str_ingredient4,omitempty"`
	StrIngredient5 *string `json:"str_ingredient5,omitempty"`
	StrIngredient6 *string `json:"str_ingredient6,omitempty"`
	StrIngredient7 *string `json:"str_ingredient7,omitempty"`
	StrIngredient8 *string `json:"str_ingredient8,omitempty"`
	StrIngredient9 *string `json:"str_ingredient9,omitempty"`
	StrInstruction *string `json:"str_instruction,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
	StrMeasure1 *string `json:"str_measure1,omitempty"`
	StrMeasure10 *string `json:"str_measure10,omitempty"`
	StrMeasure11 *string `json:"str_measure11,omitempty"`
	StrMeasure12 *string `json:"str_measure12,omitempty"`
	StrMeasure13 *string `json:"str_measure13,omitempty"`
	StrMeasure14 *string `json:"str_measure14,omitempty"`
	StrMeasure15 *string `json:"str_measure15,omitempty"`
	StrMeasure16 *string `json:"str_measure16,omitempty"`
	StrMeasure17 *string `json:"str_measure17,omitempty"`
	StrMeasure18 *string `json:"str_measure18,omitempty"`
	StrMeasure19 *string `json:"str_measure19,omitempty"`
	StrMeasure2 *string `json:"str_measure2,omitempty"`
	StrMeasure20 *string `json:"str_measure20,omitempty"`
	StrMeasure3 *string `json:"str_measure3,omitempty"`
	StrMeasure4 *string `json:"str_measure4,omitempty"`
	StrMeasure5 *string `json:"str_measure5,omitempty"`
	StrMeasure6 *string `json:"str_measure6,omitempty"`
	StrMeasure7 *string `json:"str_measure7,omitempty"`
	StrMeasure8 *string `json:"str_measure8,omitempty"`
	StrMeasure9 *string `json:"str_measure9,omitempty"`
	StrSource *string `json:"str_source,omitempty"`
	StrTag *string `json:"str_tag,omitempty"`
	StrYoutube *string `json:"str_youtube,omitempty"`
}

// RandomselectionListMatch is the typed request payload for Randomselection.ListTyped.
type RandomselectionListMatch struct {
	DateModified *string `json:"date_modified,omitempty"`
	IdMeal *string `json:"id_meal,omitempty"`
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"str_creative_commons_confirmed,omitempty"`
	StrDrinkAlternate *string `json:"str_drink_alternate,omitempty"`
	StrImageSource *string `json:"str_image_source,omitempty"`
	StrIngredient1 *string `json:"str_ingredient1,omitempty"`
	StrIngredient10 *string `json:"str_ingredient10,omitempty"`
	StrIngredient11 *string `json:"str_ingredient11,omitempty"`
	StrIngredient12 *string `json:"str_ingredient12,omitempty"`
	StrIngredient13 *string `json:"str_ingredient13,omitempty"`
	StrIngredient14 *string `json:"str_ingredient14,omitempty"`
	StrIngredient15 *string `json:"str_ingredient15,omitempty"`
	StrIngredient16 *string `json:"str_ingredient16,omitempty"`
	StrIngredient17 *string `json:"str_ingredient17,omitempty"`
	StrIngredient18 *string `json:"str_ingredient18,omitempty"`
	StrIngredient19 *string `json:"str_ingredient19,omitempty"`
	StrIngredient2 *string `json:"str_ingredient2,omitempty"`
	StrIngredient20 *string `json:"str_ingredient20,omitempty"`
	StrIngredient3 *string `json:"str_ingredient3,omitempty"`
	StrIngredient4 *string `json:"str_ingredient4,omitempty"`
	StrIngredient5 *string `json:"str_ingredient5,omitempty"`
	StrIngredient6 *string `json:"str_ingredient6,omitempty"`
	StrIngredient7 *string `json:"str_ingredient7,omitempty"`
	StrIngredient8 *string `json:"str_ingredient8,omitempty"`
	StrIngredient9 *string `json:"str_ingredient9,omitempty"`
	StrInstruction *string `json:"str_instruction,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
	StrMeasure1 *string `json:"str_measure1,omitempty"`
	StrMeasure10 *string `json:"str_measure10,omitempty"`
	StrMeasure11 *string `json:"str_measure11,omitempty"`
	StrMeasure12 *string `json:"str_measure12,omitempty"`
	StrMeasure13 *string `json:"str_measure13,omitempty"`
	StrMeasure14 *string `json:"str_measure14,omitempty"`
	StrMeasure15 *string `json:"str_measure15,omitempty"`
	StrMeasure16 *string `json:"str_measure16,omitempty"`
	StrMeasure17 *string `json:"str_measure17,omitempty"`
	StrMeasure18 *string `json:"str_measure18,omitempty"`
	StrMeasure19 *string `json:"str_measure19,omitempty"`
	StrMeasure2 *string `json:"str_measure2,omitempty"`
	StrMeasure20 *string `json:"str_measure20,omitempty"`
	StrMeasure3 *string `json:"str_measure3,omitempty"`
	StrMeasure4 *string `json:"str_measure4,omitempty"`
	StrMeasure5 *string `json:"str_measure5,omitempty"`
	StrMeasure6 *string `json:"str_measure6,omitempty"`
	StrMeasure7 *string `json:"str_measure7,omitempty"`
	StrMeasure8 *string `json:"str_measure8,omitempty"`
	StrMeasure9 *string `json:"str_measure9,omitempty"`
	StrSource *string `json:"str_source,omitempty"`
	StrTag *string `json:"str_tag,omitempty"`
	StrYoutube *string `json:"str_youtube,omitempty"`
}

// Search is the typed data model for the search entity.
type Search struct {
	DateModified *string `json:"date_modified,omitempty"`
	IdMeal *string `json:"id_meal,omitempty"`
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"str_creative_commons_confirmed,omitempty"`
	StrDrinkAlternate *string `json:"str_drink_alternate,omitempty"`
	StrImageSource *string `json:"str_image_source,omitempty"`
	StrIngredient1 *string `json:"str_ingredient1,omitempty"`
	StrIngredient10 *string `json:"str_ingredient10,omitempty"`
	StrIngredient11 *string `json:"str_ingredient11,omitempty"`
	StrIngredient12 *string `json:"str_ingredient12,omitempty"`
	StrIngredient13 *string `json:"str_ingredient13,omitempty"`
	StrIngredient14 *string `json:"str_ingredient14,omitempty"`
	StrIngredient15 *string `json:"str_ingredient15,omitempty"`
	StrIngredient16 *string `json:"str_ingredient16,omitempty"`
	StrIngredient17 *string `json:"str_ingredient17,omitempty"`
	StrIngredient18 *string `json:"str_ingredient18,omitempty"`
	StrIngredient19 *string `json:"str_ingredient19,omitempty"`
	StrIngredient2 *string `json:"str_ingredient2,omitempty"`
	StrIngredient20 *string `json:"str_ingredient20,omitempty"`
	StrIngredient3 *string `json:"str_ingredient3,omitempty"`
	StrIngredient4 *string `json:"str_ingredient4,omitempty"`
	StrIngredient5 *string `json:"str_ingredient5,omitempty"`
	StrIngredient6 *string `json:"str_ingredient6,omitempty"`
	StrIngredient7 *string `json:"str_ingredient7,omitempty"`
	StrIngredient8 *string `json:"str_ingredient8,omitempty"`
	StrIngredient9 *string `json:"str_ingredient9,omitempty"`
	StrInstruction *string `json:"str_instruction,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
	StrMeasure1 *string `json:"str_measure1,omitempty"`
	StrMeasure10 *string `json:"str_measure10,omitempty"`
	StrMeasure11 *string `json:"str_measure11,omitempty"`
	StrMeasure12 *string `json:"str_measure12,omitempty"`
	StrMeasure13 *string `json:"str_measure13,omitempty"`
	StrMeasure14 *string `json:"str_measure14,omitempty"`
	StrMeasure15 *string `json:"str_measure15,omitempty"`
	StrMeasure16 *string `json:"str_measure16,omitempty"`
	StrMeasure17 *string `json:"str_measure17,omitempty"`
	StrMeasure18 *string `json:"str_measure18,omitempty"`
	StrMeasure19 *string `json:"str_measure19,omitempty"`
	StrMeasure2 *string `json:"str_measure2,omitempty"`
	StrMeasure20 *string `json:"str_measure20,omitempty"`
	StrMeasure3 *string `json:"str_measure3,omitempty"`
	StrMeasure4 *string `json:"str_measure4,omitempty"`
	StrMeasure5 *string `json:"str_measure5,omitempty"`
	StrMeasure6 *string `json:"str_measure6,omitempty"`
	StrMeasure7 *string `json:"str_measure7,omitempty"`
	StrMeasure8 *string `json:"str_measure8,omitempty"`
	StrMeasure9 *string `json:"str_measure9,omitempty"`
	StrSource *string `json:"str_source,omitempty"`
	StrTag *string `json:"str_tag,omitempty"`
	StrYoutube *string `json:"str_youtube,omitempty"`
}

// SearchListMatch is the typed request payload for Search.ListTyped.
type SearchListMatch struct {
	DateModified *string `json:"date_modified,omitempty"`
	IdMeal *string `json:"id_meal,omitempty"`
	StrArea *string `json:"str_area,omitempty"`
	StrCategory *string `json:"str_category,omitempty"`
	StrCreativeCommonsConfirmed *string `json:"str_creative_commons_confirmed,omitempty"`
	StrDrinkAlternate *string `json:"str_drink_alternate,omitempty"`
	StrImageSource *string `json:"str_image_source,omitempty"`
	StrIngredient1 *string `json:"str_ingredient1,omitempty"`
	StrIngredient10 *string `json:"str_ingredient10,omitempty"`
	StrIngredient11 *string `json:"str_ingredient11,omitempty"`
	StrIngredient12 *string `json:"str_ingredient12,omitempty"`
	StrIngredient13 *string `json:"str_ingredient13,omitempty"`
	StrIngredient14 *string `json:"str_ingredient14,omitempty"`
	StrIngredient15 *string `json:"str_ingredient15,omitempty"`
	StrIngredient16 *string `json:"str_ingredient16,omitempty"`
	StrIngredient17 *string `json:"str_ingredient17,omitempty"`
	StrIngredient18 *string `json:"str_ingredient18,omitempty"`
	StrIngredient19 *string `json:"str_ingredient19,omitempty"`
	StrIngredient2 *string `json:"str_ingredient2,omitempty"`
	StrIngredient20 *string `json:"str_ingredient20,omitempty"`
	StrIngredient3 *string `json:"str_ingredient3,omitempty"`
	StrIngredient4 *string `json:"str_ingredient4,omitempty"`
	StrIngredient5 *string `json:"str_ingredient5,omitempty"`
	StrIngredient6 *string `json:"str_ingredient6,omitempty"`
	StrIngredient7 *string `json:"str_ingredient7,omitempty"`
	StrIngredient8 *string `json:"str_ingredient8,omitempty"`
	StrIngredient9 *string `json:"str_ingredient9,omitempty"`
	StrInstruction *string `json:"str_instruction,omitempty"`
	StrMeal *string `json:"str_meal,omitempty"`
	StrMealThumb *string `json:"str_meal_thumb,omitempty"`
	StrMeasure1 *string `json:"str_measure1,omitempty"`
	StrMeasure10 *string `json:"str_measure10,omitempty"`
	StrMeasure11 *string `json:"str_measure11,omitempty"`
	StrMeasure12 *string `json:"str_measure12,omitempty"`
	StrMeasure13 *string `json:"str_measure13,omitempty"`
	StrMeasure14 *string `json:"str_measure14,omitempty"`
	StrMeasure15 *string `json:"str_measure15,omitempty"`
	StrMeasure16 *string `json:"str_measure16,omitempty"`
	StrMeasure17 *string `json:"str_measure17,omitempty"`
	StrMeasure18 *string `json:"str_measure18,omitempty"`
	StrMeasure19 *string `json:"str_measure19,omitempty"`
	StrMeasure2 *string `json:"str_measure2,omitempty"`
	StrMeasure20 *string `json:"str_measure20,omitempty"`
	StrMeasure3 *string `json:"str_measure3,omitempty"`
	StrMeasure4 *string `json:"str_measure4,omitempty"`
	StrMeasure5 *string `json:"str_measure5,omitempty"`
	StrMeasure6 *string `json:"str_measure6,omitempty"`
	StrMeasure7 *string `json:"str_measure7,omitempty"`
	StrMeasure8 *string `json:"str_measure8,omitempty"`
	StrMeasure9 *string `json:"str_measure9,omitempty"`
	StrSource *string `json:"str_source,omitempty"`
	StrTag *string `json:"str_tag,omitempty"`
	StrYoutube *string `json:"str_youtube,omitempty"`
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

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
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

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
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

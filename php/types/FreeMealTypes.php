<?php
declare(strict_types=1);

// Typed models for the FreeMeal SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Category entity data model. */
class Category
{
    public ?string $idCategory = null;
    public ?string $strCategory = null;
    public ?string $strCategoryDescription = null;
    public ?string $strCategoryThumb = null;
}

/** Request payload for Category#list. */
class CategoryListMatch
{
    public ?string $idCategory = null;
    public ?string $strCategory = null;
    public ?string $strCategoryDescription = null;
    public ?string $strCategoryThumb = null;
}

/** Filter entity data model. */
class Filter
{
    public ?string $idMeal = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
}

/** Request payload for Filter#list. */
class FilterListMatch
{
    public ?string $idMeal = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
}

/** Latest entity data model. */
class Latest
{
    public ?string $dateModified = null;
    public ?string $idMeal = null;
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strCreativeCommonsConfirmed = null;
    public ?string $strDrinkAlternate = null;
    public ?string $strImageSource = null;
    public ?string $strIngredient1 = null;
    public ?string $strIngredient10 = null;
    public ?string $strIngredient11 = null;
    public ?string $strIngredient12 = null;
    public ?string $strIngredient13 = null;
    public ?string $strIngredient14 = null;
    public ?string $strIngredient15 = null;
    public ?string $strIngredient16 = null;
    public ?string $strIngredient17 = null;
    public ?string $strIngredient18 = null;
    public ?string $strIngredient19 = null;
    public ?string $strIngredient2 = null;
    public ?string $strIngredient20 = null;
    public ?string $strIngredient3 = null;
    public ?string $strIngredient4 = null;
    public ?string $strIngredient5 = null;
    public ?string $strIngredient6 = null;
    public ?string $strIngredient7 = null;
    public ?string $strIngredient8 = null;
    public ?string $strIngredient9 = null;
    public ?string $strInstructions = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
    public ?string $strMeasure1 = null;
    public ?string $strMeasure10 = null;
    public ?string $strMeasure11 = null;
    public ?string $strMeasure12 = null;
    public ?string $strMeasure13 = null;
    public ?string $strMeasure14 = null;
    public ?string $strMeasure15 = null;
    public ?string $strMeasure16 = null;
    public ?string $strMeasure17 = null;
    public ?string $strMeasure18 = null;
    public ?string $strMeasure19 = null;
    public ?string $strMeasure2 = null;
    public ?string $strMeasure20 = null;
    public ?string $strMeasure3 = null;
    public ?string $strMeasure4 = null;
    public ?string $strMeasure5 = null;
    public ?string $strMeasure6 = null;
    public ?string $strMeasure7 = null;
    public ?string $strMeasure8 = null;
    public ?string $strMeasure9 = null;
    public ?string $strSource = null;
    public ?string $strTags = null;
    public ?string $strYoutube = null;
}

/** Request payload for Latest#list. */
class LatestListMatch
{
    public ?string $dateModified = null;
    public ?string $idMeal = null;
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strCreativeCommonsConfirmed = null;
    public ?string $strDrinkAlternate = null;
    public ?string $strImageSource = null;
    public ?string $strIngredient1 = null;
    public ?string $strIngredient10 = null;
    public ?string $strIngredient11 = null;
    public ?string $strIngredient12 = null;
    public ?string $strIngredient13 = null;
    public ?string $strIngredient14 = null;
    public ?string $strIngredient15 = null;
    public ?string $strIngredient16 = null;
    public ?string $strIngredient17 = null;
    public ?string $strIngredient18 = null;
    public ?string $strIngredient19 = null;
    public ?string $strIngredient2 = null;
    public ?string $strIngredient20 = null;
    public ?string $strIngredient3 = null;
    public ?string $strIngredient4 = null;
    public ?string $strIngredient5 = null;
    public ?string $strIngredient6 = null;
    public ?string $strIngredient7 = null;
    public ?string $strIngredient8 = null;
    public ?string $strIngredient9 = null;
    public ?string $strInstructions = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
    public ?string $strMeasure1 = null;
    public ?string $strMeasure10 = null;
    public ?string $strMeasure11 = null;
    public ?string $strMeasure12 = null;
    public ?string $strMeasure13 = null;
    public ?string $strMeasure14 = null;
    public ?string $strMeasure15 = null;
    public ?string $strMeasure16 = null;
    public ?string $strMeasure17 = null;
    public ?string $strMeasure18 = null;
    public ?string $strMeasure19 = null;
    public ?string $strMeasure2 = null;
    public ?string $strMeasure20 = null;
    public ?string $strMeasure3 = null;
    public ?string $strMeasure4 = null;
    public ?string $strMeasure5 = null;
    public ?string $strMeasure6 = null;
    public ?string $strMeasure7 = null;
    public ?string $strMeasure8 = null;
    public ?string $strMeasure9 = null;
    public ?string $strSource = null;
    public ?string $strTags = null;
    public ?string $strYoutube = null;
}

/** List entity data model. */
class List
{
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strIngredient = null;
}

/** Request payload for List#list. */
class ListListMatch
{
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strIngredient = null;
}

/** Lookup entity data model. */
class Lookup
{
    public ?string $dateModified = null;
    public ?string $idMeal = null;
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strCreativeCommonsConfirmed = null;
    public ?string $strDrinkAlternate = null;
    public ?string $strImageSource = null;
    public ?string $strIngredient1 = null;
    public ?string $strIngredient10 = null;
    public ?string $strIngredient11 = null;
    public ?string $strIngredient12 = null;
    public ?string $strIngredient13 = null;
    public ?string $strIngredient14 = null;
    public ?string $strIngredient15 = null;
    public ?string $strIngredient16 = null;
    public ?string $strIngredient17 = null;
    public ?string $strIngredient18 = null;
    public ?string $strIngredient19 = null;
    public ?string $strIngredient2 = null;
    public ?string $strIngredient20 = null;
    public ?string $strIngredient3 = null;
    public ?string $strIngredient4 = null;
    public ?string $strIngredient5 = null;
    public ?string $strIngredient6 = null;
    public ?string $strIngredient7 = null;
    public ?string $strIngredient8 = null;
    public ?string $strIngredient9 = null;
    public ?string $strInstructions = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
    public ?string $strMeasure1 = null;
    public ?string $strMeasure10 = null;
    public ?string $strMeasure11 = null;
    public ?string $strMeasure12 = null;
    public ?string $strMeasure13 = null;
    public ?string $strMeasure14 = null;
    public ?string $strMeasure15 = null;
    public ?string $strMeasure16 = null;
    public ?string $strMeasure17 = null;
    public ?string $strMeasure18 = null;
    public ?string $strMeasure19 = null;
    public ?string $strMeasure2 = null;
    public ?string $strMeasure20 = null;
    public ?string $strMeasure3 = null;
    public ?string $strMeasure4 = null;
    public ?string $strMeasure5 = null;
    public ?string $strMeasure6 = null;
    public ?string $strMeasure7 = null;
    public ?string $strMeasure8 = null;
    public ?string $strMeasure9 = null;
    public ?string $strSource = null;
    public ?string $strTags = null;
    public ?string $strYoutube = null;
}

/** Request payload for Lookup#list. */
class LookupListMatch
{
    public ?string $dateModified = null;
    public ?string $idMeal = null;
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strCreativeCommonsConfirmed = null;
    public ?string $strDrinkAlternate = null;
    public ?string $strImageSource = null;
    public ?string $strIngredient1 = null;
    public ?string $strIngredient10 = null;
    public ?string $strIngredient11 = null;
    public ?string $strIngredient12 = null;
    public ?string $strIngredient13 = null;
    public ?string $strIngredient14 = null;
    public ?string $strIngredient15 = null;
    public ?string $strIngredient16 = null;
    public ?string $strIngredient17 = null;
    public ?string $strIngredient18 = null;
    public ?string $strIngredient19 = null;
    public ?string $strIngredient2 = null;
    public ?string $strIngredient20 = null;
    public ?string $strIngredient3 = null;
    public ?string $strIngredient4 = null;
    public ?string $strIngredient5 = null;
    public ?string $strIngredient6 = null;
    public ?string $strIngredient7 = null;
    public ?string $strIngredient8 = null;
    public ?string $strIngredient9 = null;
    public ?string $strInstructions = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
    public ?string $strMeasure1 = null;
    public ?string $strMeasure10 = null;
    public ?string $strMeasure11 = null;
    public ?string $strMeasure12 = null;
    public ?string $strMeasure13 = null;
    public ?string $strMeasure14 = null;
    public ?string $strMeasure15 = null;
    public ?string $strMeasure16 = null;
    public ?string $strMeasure17 = null;
    public ?string $strMeasure18 = null;
    public ?string $strMeasure19 = null;
    public ?string $strMeasure2 = null;
    public ?string $strMeasure20 = null;
    public ?string $strMeasure3 = null;
    public ?string $strMeasure4 = null;
    public ?string $strMeasure5 = null;
    public ?string $strMeasure6 = null;
    public ?string $strMeasure7 = null;
    public ?string $strMeasure8 = null;
    public ?string $strMeasure9 = null;
    public ?string $strSource = null;
    public ?string $strTags = null;
    public ?string $strYoutube = null;
}

/** Random entity data model. */
class Random
{
    public ?string $dateModified = null;
    public ?string $idMeal = null;
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strCreativeCommonsConfirmed = null;
    public ?string $strDrinkAlternate = null;
    public ?string $strImageSource = null;
    public ?string $strIngredient1 = null;
    public ?string $strIngredient10 = null;
    public ?string $strIngredient11 = null;
    public ?string $strIngredient12 = null;
    public ?string $strIngredient13 = null;
    public ?string $strIngredient14 = null;
    public ?string $strIngredient15 = null;
    public ?string $strIngredient16 = null;
    public ?string $strIngredient17 = null;
    public ?string $strIngredient18 = null;
    public ?string $strIngredient19 = null;
    public ?string $strIngredient2 = null;
    public ?string $strIngredient20 = null;
    public ?string $strIngredient3 = null;
    public ?string $strIngredient4 = null;
    public ?string $strIngredient5 = null;
    public ?string $strIngredient6 = null;
    public ?string $strIngredient7 = null;
    public ?string $strIngredient8 = null;
    public ?string $strIngredient9 = null;
    public ?string $strInstructions = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
    public ?string $strMeasure1 = null;
    public ?string $strMeasure10 = null;
    public ?string $strMeasure11 = null;
    public ?string $strMeasure12 = null;
    public ?string $strMeasure13 = null;
    public ?string $strMeasure14 = null;
    public ?string $strMeasure15 = null;
    public ?string $strMeasure16 = null;
    public ?string $strMeasure17 = null;
    public ?string $strMeasure18 = null;
    public ?string $strMeasure19 = null;
    public ?string $strMeasure2 = null;
    public ?string $strMeasure20 = null;
    public ?string $strMeasure3 = null;
    public ?string $strMeasure4 = null;
    public ?string $strMeasure5 = null;
    public ?string $strMeasure6 = null;
    public ?string $strMeasure7 = null;
    public ?string $strMeasure8 = null;
    public ?string $strMeasure9 = null;
    public ?string $strSource = null;
    public ?string $strTags = null;
    public ?string $strYoutube = null;
}

/** Request payload for Random#list. */
class RandomListMatch
{
    public ?string $dateModified = null;
    public ?string $idMeal = null;
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strCreativeCommonsConfirmed = null;
    public ?string $strDrinkAlternate = null;
    public ?string $strImageSource = null;
    public ?string $strIngredient1 = null;
    public ?string $strIngredient10 = null;
    public ?string $strIngredient11 = null;
    public ?string $strIngredient12 = null;
    public ?string $strIngredient13 = null;
    public ?string $strIngredient14 = null;
    public ?string $strIngredient15 = null;
    public ?string $strIngredient16 = null;
    public ?string $strIngredient17 = null;
    public ?string $strIngredient18 = null;
    public ?string $strIngredient19 = null;
    public ?string $strIngredient2 = null;
    public ?string $strIngredient20 = null;
    public ?string $strIngredient3 = null;
    public ?string $strIngredient4 = null;
    public ?string $strIngredient5 = null;
    public ?string $strIngredient6 = null;
    public ?string $strIngredient7 = null;
    public ?string $strIngredient8 = null;
    public ?string $strIngredient9 = null;
    public ?string $strInstructions = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
    public ?string $strMeasure1 = null;
    public ?string $strMeasure10 = null;
    public ?string $strMeasure11 = null;
    public ?string $strMeasure12 = null;
    public ?string $strMeasure13 = null;
    public ?string $strMeasure14 = null;
    public ?string $strMeasure15 = null;
    public ?string $strMeasure16 = null;
    public ?string $strMeasure17 = null;
    public ?string $strMeasure18 = null;
    public ?string $strMeasure19 = null;
    public ?string $strMeasure2 = null;
    public ?string $strMeasure20 = null;
    public ?string $strMeasure3 = null;
    public ?string $strMeasure4 = null;
    public ?string $strMeasure5 = null;
    public ?string $strMeasure6 = null;
    public ?string $strMeasure7 = null;
    public ?string $strMeasure8 = null;
    public ?string $strMeasure9 = null;
    public ?string $strSource = null;
    public ?string $strTags = null;
    public ?string $strYoutube = null;
}

/** Randomselection entity data model. */
class Randomselection
{
    public ?string $dateModified = null;
    public ?string $idMeal = null;
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strCreativeCommonsConfirmed = null;
    public ?string $strDrinkAlternate = null;
    public ?string $strImageSource = null;
    public ?string $strIngredient1 = null;
    public ?string $strIngredient10 = null;
    public ?string $strIngredient11 = null;
    public ?string $strIngredient12 = null;
    public ?string $strIngredient13 = null;
    public ?string $strIngredient14 = null;
    public ?string $strIngredient15 = null;
    public ?string $strIngredient16 = null;
    public ?string $strIngredient17 = null;
    public ?string $strIngredient18 = null;
    public ?string $strIngredient19 = null;
    public ?string $strIngredient2 = null;
    public ?string $strIngredient20 = null;
    public ?string $strIngredient3 = null;
    public ?string $strIngredient4 = null;
    public ?string $strIngredient5 = null;
    public ?string $strIngredient6 = null;
    public ?string $strIngredient7 = null;
    public ?string $strIngredient8 = null;
    public ?string $strIngredient9 = null;
    public ?string $strInstructions = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
    public ?string $strMeasure1 = null;
    public ?string $strMeasure10 = null;
    public ?string $strMeasure11 = null;
    public ?string $strMeasure12 = null;
    public ?string $strMeasure13 = null;
    public ?string $strMeasure14 = null;
    public ?string $strMeasure15 = null;
    public ?string $strMeasure16 = null;
    public ?string $strMeasure17 = null;
    public ?string $strMeasure18 = null;
    public ?string $strMeasure19 = null;
    public ?string $strMeasure2 = null;
    public ?string $strMeasure20 = null;
    public ?string $strMeasure3 = null;
    public ?string $strMeasure4 = null;
    public ?string $strMeasure5 = null;
    public ?string $strMeasure6 = null;
    public ?string $strMeasure7 = null;
    public ?string $strMeasure8 = null;
    public ?string $strMeasure9 = null;
    public ?string $strSource = null;
    public ?string $strTags = null;
    public ?string $strYoutube = null;
}

/** Request payload for Randomselection#list. */
class RandomselectionListMatch
{
    public ?string $dateModified = null;
    public ?string $idMeal = null;
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strCreativeCommonsConfirmed = null;
    public ?string $strDrinkAlternate = null;
    public ?string $strImageSource = null;
    public ?string $strIngredient1 = null;
    public ?string $strIngredient10 = null;
    public ?string $strIngredient11 = null;
    public ?string $strIngredient12 = null;
    public ?string $strIngredient13 = null;
    public ?string $strIngredient14 = null;
    public ?string $strIngredient15 = null;
    public ?string $strIngredient16 = null;
    public ?string $strIngredient17 = null;
    public ?string $strIngredient18 = null;
    public ?string $strIngredient19 = null;
    public ?string $strIngredient2 = null;
    public ?string $strIngredient20 = null;
    public ?string $strIngredient3 = null;
    public ?string $strIngredient4 = null;
    public ?string $strIngredient5 = null;
    public ?string $strIngredient6 = null;
    public ?string $strIngredient7 = null;
    public ?string $strIngredient8 = null;
    public ?string $strIngredient9 = null;
    public ?string $strInstructions = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
    public ?string $strMeasure1 = null;
    public ?string $strMeasure10 = null;
    public ?string $strMeasure11 = null;
    public ?string $strMeasure12 = null;
    public ?string $strMeasure13 = null;
    public ?string $strMeasure14 = null;
    public ?string $strMeasure15 = null;
    public ?string $strMeasure16 = null;
    public ?string $strMeasure17 = null;
    public ?string $strMeasure18 = null;
    public ?string $strMeasure19 = null;
    public ?string $strMeasure2 = null;
    public ?string $strMeasure20 = null;
    public ?string $strMeasure3 = null;
    public ?string $strMeasure4 = null;
    public ?string $strMeasure5 = null;
    public ?string $strMeasure6 = null;
    public ?string $strMeasure7 = null;
    public ?string $strMeasure8 = null;
    public ?string $strMeasure9 = null;
    public ?string $strSource = null;
    public ?string $strTags = null;
    public ?string $strYoutube = null;
}

/** Search entity data model. */
class Search
{
    public ?string $dateModified = null;
    public ?string $idMeal = null;
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strCreativeCommonsConfirmed = null;
    public ?string $strDrinkAlternate = null;
    public ?string $strImageSource = null;
    public ?string $strIngredient1 = null;
    public ?string $strIngredient10 = null;
    public ?string $strIngredient11 = null;
    public ?string $strIngredient12 = null;
    public ?string $strIngredient13 = null;
    public ?string $strIngredient14 = null;
    public ?string $strIngredient15 = null;
    public ?string $strIngredient16 = null;
    public ?string $strIngredient17 = null;
    public ?string $strIngredient18 = null;
    public ?string $strIngredient19 = null;
    public ?string $strIngredient2 = null;
    public ?string $strIngredient20 = null;
    public ?string $strIngredient3 = null;
    public ?string $strIngredient4 = null;
    public ?string $strIngredient5 = null;
    public ?string $strIngredient6 = null;
    public ?string $strIngredient7 = null;
    public ?string $strIngredient8 = null;
    public ?string $strIngredient9 = null;
    public ?string $strInstructions = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
    public ?string $strMeasure1 = null;
    public ?string $strMeasure10 = null;
    public ?string $strMeasure11 = null;
    public ?string $strMeasure12 = null;
    public ?string $strMeasure13 = null;
    public ?string $strMeasure14 = null;
    public ?string $strMeasure15 = null;
    public ?string $strMeasure16 = null;
    public ?string $strMeasure17 = null;
    public ?string $strMeasure18 = null;
    public ?string $strMeasure19 = null;
    public ?string $strMeasure2 = null;
    public ?string $strMeasure20 = null;
    public ?string $strMeasure3 = null;
    public ?string $strMeasure4 = null;
    public ?string $strMeasure5 = null;
    public ?string $strMeasure6 = null;
    public ?string $strMeasure7 = null;
    public ?string $strMeasure8 = null;
    public ?string $strMeasure9 = null;
    public ?string $strSource = null;
    public ?string $strTags = null;
    public ?string $strYoutube = null;
}

/** Request payload for Search#list. */
class SearchListMatch
{
    public ?string $dateModified = null;
    public ?string $idMeal = null;
    public ?string $strArea = null;
    public ?string $strCategory = null;
    public ?string $strCreativeCommonsConfirmed = null;
    public ?string $strDrinkAlternate = null;
    public ?string $strImageSource = null;
    public ?string $strIngredient1 = null;
    public ?string $strIngredient10 = null;
    public ?string $strIngredient11 = null;
    public ?string $strIngredient12 = null;
    public ?string $strIngredient13 = null;
    public ?string $strIngredient14 = null;
    public ?string $strIngredient15 = null;
    public ?string $strIngredient16 = null;
    public ?string $strIngredient17 = null;
    public ?string $strIngredient18 = null;
    public ?string $strIngredient19 = null;
    public ?string $strIngredient2 = null;
    public ?string $strIngredient20 = null;
    public ?string $strIngredient3 = null;
    public ?string $strIngredient4 = null;
    public ?string $strIngredient5 = null;
    public ?string $strIngredient6 = null;
    public ?string $strIngredient7 = null;
    public ?string $strIngredient8 = null;
    public ?string $strIngredient9 = null;
    public ?string $strInstructions = null;
    public ?string $strMeal = null;
    public ?string $strMealThumb = null;
    public ?string $strMeasure1 = null;
    public ?string $strMeasure10 = null;
    public ?string $strMeasure11 = null;
    public ?string $strMeasure12 = null;
    public ?string $strMeasure13 = null;
    public ?string $strMeasure14 = null;
    public ?string $strMeasure15 = null;
    public ?string $strMeasure16 = null;
    public ?string $strMeasure17 = null;
    public ?string $strMeasure18 = null;
    public ?string $strMeasure19 = null;
    public ?string $strMeasure2 = null;
    public ?string $strMeasure20 = null;
    public ?string $strMeasure3 = null;
    public ?string $strMeasure4 = null;
    public ?string $strMeasure5 = null;
    public ?string $strMeasure6 = null;
    public ?string $strMeasure7 = null;
    public ?string $strMeasure8 = null;
    public ?string $strMeasure9 = null;
    public ?string $strSource = null;
    public ?string $strTags = null;
    public ?string $strYoutube = null;
}


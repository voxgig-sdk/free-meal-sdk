# frozen_string_literal: true

# Typed models for the FreeMeal SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Category entity data model.
#
# @!attribute [rw] idCategory
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strCategoryDescription
#   @return [String, nil]
#
# @!attribute [rw] strCategoryThumb
#   @return [String, nil]
Category = Struct.new(
  :idCategory,
  :strCategory,
  :strCategoryDescription,
  :strCategoryThumb,
  keyword_init: true
)

# Request payload for Category#list.
#
# @!attribute [rw] idCategory
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strCategoryDescription
#   @return [String, nil]
#
# @!attribute [rw] strCategoryThumb
#   @return [String, nil]
CategoryListMatch = Struct.new(
  :idCategory,
  :strCategory,
  :strCategoryDescription,
  :strCategoryThumb,
  keyword_init: true
)

# Filter entity data model.
#
# @!attribute [rw] idMeal
#   @return [String, nil]
#
# @!attribute [rw] strMeal
#   @return [String, nil]
#
# @!attribute [rw] strMealThumb
#   @return [String, nil]
Filter = Struct.new(
  :idMeal,
  :strMeal,
  :strMealThumb,
  keyword_init: true
)

# Request payload for Filter#list.
#
# @!attribute [rw] a
#   @return [String, nil]
#
# @!attribute [rw] c
#   @return [String, nil]
#
# @!attribute [rw] i
#   @return [String, nil]
FilterListMatch = Struct.new(
  :a,
  :c,
  :i,
  keyword_init: true
)

# Latest entity data model.
#
# @!attribute [rw] dateModified
#   @return [String, nil]
#
# @!attribute [rw] idMeal
#   @return [String, nil]
#
# @!attribute [rw] strArea
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strCreativeCommonsConfirmed
#   @return [String, nil]
#
# @!attribute [rw] strDrinkAlternate
#   @return [String, nil]
#
# @!attribute [rw] strImageSource
#   @return [String, nil]
#
# @!attribute [rw] strIngredient1
#   @return [String, nil]
#
# @!attribute [rw] strIngredient10
#   @return [String, nil]
#
# @!attribute [rw] strIngredient11
#   @return [String, nil]
#
# @!attribute [rw] strIngredient12
#   @return [String, nil]
#
# @!attribute [rw] strIngredient13
#   @return [String, nil]
#
# @!attribute [rw] strIngredient14
#   @return [String, nil]
#
# @!attribute [rw] strIngredient15
#   @return [String, nil]
#
# @!attribute [rw] strIngredient16
#   @return [String, nil]
#
# @!attribute [rw] strIngredient17
#   @return [String, nil]
#
# @!attribute [rw] strIngredient18
#   @return [String, nil]
#
# @!attribute [rw] strIngredient19
#   @return [String, nil]
#
# @!attribute [rw] strIngredient2
#   @return [String, nil]
#
# @!attribute [rw] strIngredient20
#   @return [String, nil]
#
# @!attribute [rw] strIngredient3
#   @return [String, nil]
#
# @!attribute [rw] strIngredient4
#   @return [String, nil]
#
# @!attribute [rw] strIngredient5
#   @return [String, nil]
#
# @!attribute [rw] strIngredient6
#   @return [String, nil]
#
# @!attribute [rw] strIngredient7
#   @return [String, nil]
#
# @!attribute [rw] strIngredient8
#   @return [String, nil]
#
# @!attribute [rw] strIngredient9
#   @return [String, nil]
#
# @!attribute [rw] strInstructions
#   @return [String, nil]
#
# @!attribute [rw] strMeal
#   @return [String, nil]
#
# @!attribute [rw] strMealThumb
#   @return [String, nil]
#
# @!attribute [rw] strMeasure1
#   @return [String, nil]
#
# @!attribute [rw] strMeasure10
#   @return [String, nil]
#
# @!attribute [rw] strMeasure11
#   @return [String, nil]
#
# @!attribute [rw] strMeasure12
#   @return [String, nil]
#
# @!attribute [rw] strMeasure13
#   @return [String, nil]
#
# @!attribute [rw] strMeasure14
#   @return [String, nil]
#
# @!attribute [rw] strMeasure15
#   @return [String, nil]
#
# @!attribute [rw] strMeasure16
#   @return [String, nil]
#
# @!attribute [rw] strMeasure17
#   @return [String, nil]
#
# @!attribute [rw] strMeasure18
#   @return [String, nil]
#
# @!attribute [rw] strMeasure19
#   @return [String, nil]
#
# @!attribute [rw] strMeasure2
#   @return [String, nil]
#
# @!attribute [rw] strMeasure20
#   @return [String, nil]
#
# @!attribute [rw] strMeasure3
#   @return [String, nil]
#
# @!attribute [rw] strMeasure4
#   @return [String, nil]
#
# @!attribute [rw] strMeasure5
#   @return [String, nil]
#
# @!attribute [rw] strMeasure6
#   @return [String, nil]
#
# @!attribute [rw] strMeasure7
#   @return [String, nil]
#
# @!attribute [rw] strMeasure8
#   @return [String, nil]
#
# @!attribute [rw] strMeasure9
#   @return [String, nil]
#
# @!attribute [rw] strSource
#   @return [String, nil]
#
# @!attribute [rw] strTags
#   @return [String, nil]
#
# @!attribute [rw] strYoutube
#   @return [String, nil]
Latest = Struct.new(
  :dateModified,
  :idMeal,
  :strArea,
  :strCategory,
  :strCreativeCommonsConfirmed,
  :strDrinkAlternate,
  :strImageSource,
  :strIngredient1,
  :strIngredient10,
  :strIngredient11,
  :strIngredient12,
  :strIngredient13,
  :strIngredient14,
  :strIngredient15,
  :strIngredient16,
  :strIngredient17,
  :strIngredient18,
  :strIngredient19,
  :strIngredient2,
  :strIngredient20,
  :strIngredient3,
  :strIngredient4,
  :strIngredient5,
  :strIngredient6,
  :strIngredient7,
  :strIngredient8,
  :strIngredient9,
  :strInstructions,
  :strMeal,
  :strMealThumb,
  :strMeasure1,
  :strMeasure10,
  :strMeasure11,
  :strMeasure12,
  :strMeasure13,
  :strMeasure14,
  :strMeasure15,
  :strMeasure16,
  :strMeasure17,
  :strMeasure18,
  :strMeasure19,
  :strMeasure2,
  :strMeasure20,
  :strMeasure3,
  :strMeasure4,
  :strMeasure5,
  :strMeasure6,
  :strMeasure7,
  :strMeasure8,
  :strMeasure9,
  :strSource,
  :strTags,
  :strYoutube,
  keyword_init: true
)

# Request payload for Latest#list.
#
# @!attribute [rw] dateModified
#   @return [String, nil]
#
# @!attribute [rw] idMeal
#   @return [String, nil]
#
# @!attribute [rw] strArea
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strCreativeCommonsConfirmed
#   @return [String, nil]
#
# @!attribute [rw] strDrinkAlternate
#   @return [String, nil]
#
# @!attribute [rw] strImageSource
#   @return [String, nil]
#
# @!attribute [rw] strIngredient1
#   @return [String, nil]
#
# @!attribute [rw] strIngredient10
#   @return [String, nil]
#
# @!attribute [rw] strIngredient11
#   @return [String, nil]
#
# @!attribute [rw] strIngredient12
#   @return [String, nil]
#
# @!attribute [rw] strIngredient13
#   @return [String, nil]
#
# @!attribute [rw] strIngredient14
#   @return [String, nil]
#
# @!attribute [rw] strIngredient15
#   @return [String, nil]
#
# @!attribute [rw] strIngredient16
#   @return [String, nil]
#
# @!attribute [rw] strIngredient17
#   @return [String, nil]
#
# @!attribute [rw] strIngredient18
#   @return [String, nil]
#
# @!attribute [rw] strIngredient19
#   @return [String, nil]
#
# @!attribute [rw] strIngredient2
#   @return [String, nil]
#
# @!attribute [rw] strIngredient20
#   @return [String, nil]
#
# @!attribute [rw] strIngredient3
#   @return [String, nil]
#
# @!attribute [rw] strIngredient4
#   @return [String, nil]
#
# @!attribute [rw] strIngredient5
#   @return [String, nil]
#
# @!attribute [rw] strIngredient6
#   @return [String, nil]
#
# @!attribute [rw] strIngredient7
#   @return [String, nil]
#
# @!attribute [rw] strIngredient8
#   @return [String, nil]
#
# @!attribute [rw] strIngredient9
#   @return [String, nil]
#
# @!attribute [rw] strInstructions
#   @return [String, nil]
#
# @!attribute [rw] strMeal
#   @return [String, nil]
#
# @!attribute [rw] strMealThumb
#   @return [String, nil]
#
# @!attribute [rw] strMeasure1
#   @return [String, nil]
#
# @!attribute [rw] strMeasure10
#   @return [String, nil]
#
# @!attribute [rw] strMeasure11
#   @return [String, nil]
#
# @!attribute [rw] strMeasure12
#   @return [String, nil]
#
# @!attribute [rw] strMeasure13
#   @return [String, nil]
#
# @!attribute [rw] strMeasure14
#   @return [String, nil]
#
# @!attribute [rw] strMeasure15
#   @return [String, nil]
#
# @!attribute [rw] strMeasure16
#   @return [String, nil]
#
# @!attribute [rw] strMeasure17
#   @return [String, nil]
#
# @!attribute [rw] strMeasure18
#   @return [String, nil]
#
# @!attribute [rw] strMeasure19
#   @return [String, nil]
#
# @!attribute [rw] strMeasure2
#   @return [String, nil]
#
# @!attribute [rw] strMeasure20
#   @return [String, nil]
#
# @!attribute [rw] strMeasure3
#   @return [String, nil]
#
# @!attribute [rw] strMeasure4
#   @return [String, nil]
#
# @!attribute [rw] strMeasure5
#   @return [String, nil]
#
# @!attribute [rw] strMeasure6
#   @return [String, nil]
#
# @!attribute [rw] strMeasure7
#   @return [String, nil]
#
# @!attribute [rw] strMeasure8
#   @return [String, nil]
#
# @!attribute [rw] strMeasure9
#   @return [String, nil]
#
# @!attribute [rw] strSource
#   @return [String, nil]
#
# @!attribute [rw] strTags
#   @return [String, nil]
#
# @!attribute [rw] strYoutube
#   @return [String, nil]
LatestListMatch = Struct.new(
  :dateModified,
  :idMeal,
  :strArea,
  :strCategory,
  :strCreativeCommonsConfirmed,
  :strDrinkAlternate,
  :strImageSource,
  :strIngredient1,
  :strIngredient10,
  :strIngredient11,
  :strIngredient12,
  :strIngredient13,
  :strIngredient14,
  :strIngredient15,
  :strIngredient16,
  :strIngredient17,
  :strIngredient18,
  :strIngredient19,
  :strIngredient2,
  :strIngredient20,
  :strIngredient3,
  :strIngredient4,
  :strIngredient5,
  :strIngredient6,
  :strIngredient7,
  :strIngredient8,
  :strIngredient9,
  :strInstructions,
  :strMeal,
  :strMealThumb,
  :strMeasure1,
  :strMeasure10,
  :strMeasure11,
  :strMeasure12,
  :strMeasure13,
  :strMeasure14,
  :strMeasure15,
  :strMeasure16,
  :strMeasure17,
  :strMeasure18,
  :strMeasure19,
  :strMeasure2,
  :strMeasure20,
  :strMeasure3,
  :strMeasure4,
  :strMeasure5,
  :strMeasure6,
  :strMeasure7,
  :strMeasure8,
  :strMeasure9,
  :strSource,
  :strTags,
  :strYoutube,
  keyword_init: true
)

# List entity data model.
#
# @!attribute [rw] strArea
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strIngredient
#   @return [String, nil]
List = Struct.new(
  :strArea,
  :strCategory,
  :strIngredient,
  keyword_init: true
)

# Request payload for List#list.
#
# @!attribute [rw] a
#   @return [String, nil]
#
# @!attribute [rw] c
#   @return [String, nil]
#
# @!attribute [rw] i
#   @return [String, nil]
ListListMatch = Struct.new(
  :a,
  :c,
  :i,
  keyword_init: true
)

# Lookup entity data model.
#
# @!attribute [rw] dateModified
#   @return [String, nil]
#
# @!attribute [rw] idMeal
#   @return [String, nil]
#
# @!attribute [rw] strArea
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strCreativeCommonsConfirmed
#   @return [String, nil]
#
# @!attribute [rw] strDrinkAlternate
#   @return [String, nil]
#
# @!attribute [rw] strImageSource
#   @return [String, nil]
#
# @!attribute [rw] strIngredient1
#   @return [String, nil]
#
# @!attribute [rw] strIngredient10
#   @return [String, nil]
#
# @!attribute [rw] strIngredient11
#   @return [String, nil]
#
# @!attribute [rw] strIngredient12
#   @return [String, nil]
#
# @!attribute [rw] strIngredient13
#   @return [String, nil]
#
# @!attribute [rw] strIngredient14
#   @return [String, nil]
#
# @!attribute [rw] strIngredient15
#   @return [String, nil]
#
# @!attribute [rw] strIngredient16
#   @return [String, nil]
#
# @!attribute [rw] strIngredient17
#   @return [String, nil]
#
# @!attribute [rw] strIngredient18
#   @return [String, nil]
#
# @!attribute [rw] strIngredient19
#   @return [String, nil]
#
# @!attribute [rw] strIngredient2
#   @return [String, nil]
#
# @!attribute [rw] strIngredient20
#   @return [String, nil]
#
# @!attribute [rw] strIngredient3
#   @return [String, nil]
#
# @!attribute [rw] strIngredient4
#   @return [String, nil]
#
# @!attribute [rw] strIngredient5
#   @return [String, nil]
#
# @!attribute [rw] strIngredient6
#   @return [String, nil]
#
# @!attribute [rw] strIngredient7
#   @return [String, nil]
#
# @!attribute [rw] strIngredient8
#   @return [String, nil]
#
# @!attribute [rw] strIngredient9
#   @return [String, nil]
#
# @!attribute [rw] strInstructions
#   @return [String, nil]
#
# @!attribute [rw] strMeal
#   @return [String, nil]
#
# @!attribute [rw] strMealThumb
#   @return [String, nil]
#
# @!attribute [rw] strMeasure1
#   @return [String, nil]
#
# @!attribute [rw] strMeasure10
#   @return [String, nil]
#
# @!attribute [rw] strMeasure11
#   @return [String, nil]
#
# @!attribute [rw] strMeasure12
#   @return [String, nil]
#
# @!attribute [rw] strMeasure13
#   @return [String, nil]
#
# @!attribute [rw] strMeasure14
#   @return [String, nil]
#
# @!attribute [rw] strMeasure15
#   @return [String, nil]
#
# @!attribute [rw] strMeasure16
#   @return [String, nil]
#
# @!attribute [rw] strMeasure17
#   @return [String, nil]
#
# @!attribute [rw] strMeasure18
#   @return [String, nil]
#
# @!attribute [rw] strMeasure19
#   @return [String, nil]
#
# @!attribute [rw] strMeasure2
#   @return [String, nil]
#
# @!attribute [rw] strMeasure20
#   @return [String, nil]
#
# @!attribute [rw] strMeasure3
#   @return [String, nil]
#
# @!attribute [rw] strMeasure4
#   @return [String, nil]
#
# @!attribute [rw] strMeasure5
#   @return [String, nil]
#
# @!attribute [rw] strMeasure6
#   @return [String, nil]
#
# @!attribute [rw] strMeasure7
#   @return [String, nil]
#
# @!attribute [rw] strMeasure8
#   @return [String, nil]
#
# @!attribute [rw] strMeasure9
#   @return [String, nil]
#
# @!attribute [rw] strSource
#   @return [String, nil]
#
# @!attribute [rw] strTags
#   @return [String, nil]
#
# @!attribute [rw] strYoutube
#   @return [String, nil]
Lookup = Struct.new(
  :dateModified,
  :idMeal,
  :strArea,
  :strCategory,
  :strCreativeCommonsConfirmed,
  :strDrinkAlternate,
  :strImageSource,
  :strIngredient1,
  :strIngredient10,
  :strIngredient11,
  :strIngredient12,
  :strIngredient13,
  :strIngredient14,
  :strIngredient15,
  :strIngredient16,
  :strIngredient17,
  :strIngredient18,
  :strIngredient19,
  :strIngredient2,
  :strIngredient20,
  :strIngredient3,
  :strIngredient4,
  :strIngredient5,
  :strIngredient6,
  :strIngredient7,
  :strIngredient8,
  :strIngredient9,
  :strInstructions,
  :strMeal,
  :strMealThumb,
  :strMeasure1,
  :strMeasure10,
  :strMeasure11,
  :strMeasure12,
  :strMeasure13,
  :strMeasure14,
  :strMeasure15,
  :strMeasure16,
  :strMeasure17,
  :strMeasure18,
  :strMeasure19,
  :strMeasure2,
  :strMeasure20,
  :strMeasure3,
  :strMeasure4,
  :strMeasure5,
  :strMeasure6,
  :strMeasure7,
  :strMeasure8,
  :strMeasure9,
  :strSource,
  :strTags,
  :strYoutube,
  keyword_init: true
)

# Request payload for Lookup#list.
#
# @!attribute [rw] i
#   @return [String]
LookupListMatch = Struct.new(
  :i,
  keyword_init: true
)

# Random entity data model.
#
# @!attribute [rw] dateModified
#   @return [String, nil]
#
# @!attribute [rw] idMeal
#   @return [String, nil]
#
# @!attribute [rw] strArea
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strCreativeCommonsConfirmed
#   @return [String, nil]
#
# @!attribute [rw] strDrinkAlternate
#   @return [String, nil]
#
# @!attribute [rw] strImageSource
#   @return [String, nil]
#
# @!attribute [rw] strIngredient1
#   @return [String, nil]
#
# @!attribute [rw] strIngredient10
#   @return [String, nil]
#
# @!attribute [rw] strIngredient11
#   @return [String, nil]
#
# @!attribute [rw] strIngredient12
#   @return [String, nil]
#
# @!attribute [rw] strIngredient13
#   @return [String, nil]
#
# @!attribute [rw] strIngredient14
#   @return [String, nil]
#
# @!attribute [rw] strIngredient15
#   @return [String, nil]
#
# @!attribute [rw] strIngredient16
#   @return [String, nil]
#
# @!attribute [rw] strIngredient17
#   @return [String, nil]
#
# @!attribute [rw] strIngredient18
#   @return [String, nil]
#
# @!attribute [rw] strIngredient19
#   @return [String, nil]
#
# @!attribute [rw] strIngredient2
#   @return [String, nil]
#
# @!attribute [rw] strIngredient20
#   @return [String, nil]
#
# @!attribute [rw] strIngredient3
#   @return [String, nil]
#
# @!attribute [rw] strIngredient4
#   @return [String, nil]
#
# @!attribute [rw] strIngredient5
#   @return [String, nil]
#
# @!attribute [rw] strIngredient6
#   @return [String, nil]
#
# @!attribute [rw] strIngredient7
#   @return [String, nil]
#
# @!attribute [rw] strIngredient8
#   @return [String, nil]
#
# @!attribute [rw] strIngredient9
#   @return [String, nil]
#
# @!attribute [rw] strInstructions
#   @return [String, nil]
#
# @!attribute [rw] strMeal
#   @return [String, nil]
#
# @!attribute [rw] strMealThumb
#   @return [String, nil]
#
# @!attribute [rw] strMeasure1
#   @return [String, nil]
#
# @!attribute [rw] strMeasure10
#   @return [String, nil]
#
# @!attribute [rw] strMeasure11
#   @return [String, nil]
#
# @!attribute [rw] strMeasure12
#   @return [String, nil]
#
# @!attribute [rw] strMeasure13
#   @return [String, nil]
#
# @!attribute [rw] strMeasure14
#   @return [String, nil]
#
# @!attribute [rw] strMeasure15
#   @return [String, nil]
#
# @!attribute [rw] strMeasure16
#   @return [String, nil]
#
# @!attribute [rw] strMeasure17
#   @return [String, nil]
#
# @!attribute [rw] strMeasure18
#   @return [String, nil]
#
# @!attribute [rw] strMeasure19
#   @return [String, nil]
#
# @!attribute [rw] strMeasure2
#   @return [String, nil]
#
# @!attribute [rw] strMeasure20
#   @return [String, nil]
#
# @!attribute [rw] strMeasure3
#   @return [String, nil]
#
# @!attribute [rw] strMeasure4
#   @return [String, nil]
#
# @!attribute [rw] strMeasure5
#   @return [String, nil]
#
# @!attribute [rw] strMeasure6
#   @return [String, nil]
#
# @!attribute [rw] strMeasure7
#   @return [String, nil]
#
# @!attribute [rw] strMeasure8
#   @return [String, nil]
#
# @!attribute [rw] strMeasure9
#   @return [String, nil]
#
# @!attribute [rw] strSource
#   @return [String, nil]
#
# @!attribute [rw] strTags
#   @return [String, nil]
#
# @!attribute [rw] strYoutube
#   @return [String, nil]
RandomType = Struct.new(
  :dateModified,
  :idMeal,
  :strArea,
  :strCategory,
  :strCreativeCommonsConfirmed,
  :strDrinkAlternate,
  :strImageSource,
  :strIngredient1,
  :strIngredient10,
  :strIngredient11,
  :strIngredient12,
  :strIngredient13,
  :strIngredient14,
  :strIngredient15,
  :strIngredient16,
  :strIngredient17,
  :strIngredient18,
  :strIngredient19,
  :strIngredient2,
  :strIngredient20,
  :strIngredient3,
  :strIngredient4,
  :strIngredient5,
  :strIngredient6,
  :strIngredient7,
  :strIngredient8,
  :strIngredient9,
  :strInstructions,
  :strMeal,
  :strMealThumb,
  :strMeasure1,
  :strMeasure10,
  :strMeasure11,
  :strMeasure12,
  :strMeasure13,
  :strMeasure14,
  :strMeasure15,
  :strMeasure16,
  :strMeasure17,
  :strMeasure18,
  :strMeasure19,
  :strMeasure2,
  :strMeasure20,
  :strMeasure3,
  :strMeasure4,
  :strMeasure5,
  :strMeasure6,
  :strMeasure7,
  :strMeasure8,
  :strMeasure9,
  :strSource,
  :strTags,
  :strYoutube,
  keyword_init: true
)

# Request payload for Random#list.
#
# @!attribute [rw] dateModified
#   @return [String, nil]
#
# @!attribute [rw] idMeal
#   @return [String, nil]
#
# @!attribute [rw] strArea
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strCreativeCommonsConfirmed
#   @return [String, nil]
#
# @!attribute [rw] strDrinkAlternate
#   @return [String, nil]
#
# @!attribute [rw] strImageSource
#   @return [String, nil]
#
# @!attribute [rw] strIngredient1
#   @return [String, nil]
#
# @!attribute [rw] strIngredient10
#   @return [String, nil]
#
# @!attribute [rw] strIngredient11
#   @return [String, nil]
#
# @!attribute [rw] strIngredient12
#   @return [String, nil]
#
# @!attribute [rw] strIngredient13
#   @return [String, nil]
#
# @!attribute [rw] strIngredient14
#   @return [String, nil]
#
# @!attribute [rw] strIngredient15
#   @return [String, nil]
#
# @!attribute [rw] strIngredient16
#   @return [String, nil]
#
# @!attribute [rw] strIngredient17
#   @return [String, nil]
#
# @!attribute [rw] strIngredient18
#   @return [String, nil]
#
# @!attribute [rw] strIngredient19
#   @return [String, nil]
#
# @!attribute [rw] strIngredient2
#   @return [String, nil]
#
# @!attribute [rw] strIngredient20
#   @return [String, nil]
#
# @!attribute [rw] strIngredient3
#   @return [String, nil]
#
# @!attribute [rw] strIngredient4
#   @return [String, nil]
#
# @!attribute [rw] strIngredient5
#   @return [String, nil]
#
# @!attribute [rw] strIngredient6
#   @return [String, nil]
#
# @!attribute [rw] strIngredient7
#   @return [String, nil]
#
# @!attribute [rw] strIngredient8
#   @return [String, nil]
#
# @!attribute [rw] strIngredient9
#   @return [String, nil]
#
# @!attribute [rw] strInstructions
#   @return [String, nil]
#
# @!attribute [rw] strMeal
#   @return [String, nil]
#
# @!attribute [rw] strMealThumb
#   @return [String, nil]
#
# @!attribute [rw] strMeasure1
#   @return [String, nil]
#
# @!attribute [rw] strMeasure10
#   @return [String, nil]
#
# @!attribute [rw] strMeasure11
#   @return [String, nil]
#
# @!attribute [rw] strMeasure12
#   @return [String, nil]
#
# @!attribute [rw] strMeasure13
#   @return [String, nil]
#
# @!attribute [rw] strMeasure14
#   @return [String, nil]
#
# @!attribute [rw] strMeasure15
#   @return [String, nil]
#
# @!attribute [rw] strMeasure16
#   @return [String, nil]
#
# @!attribute [rw] strMeasure17
#   @return [String, nil]
#
# @!attribute [rw] strMeasure18
#   @return [String, nil]
#
# @!attribute [rw] strMeasure19
#   @return [String, nil]
#
# @!attribute [rw] strMeasure2
#   @return [String, nil]
#
# @!attribute [rw] strMeasure20
#   @return [String, nil]
#
# @!attribute [rw] strMeasure3
#   @return [String, nil]
#
# @!attribute [rw] strMeasure4
#   @return [String, nil]
#
# @!attribute [rw] strMeasure5
#   @return [String, nil]
#
# @!attribute [rw] strMeasure6
#   @return [String, nil]
#
# @!attribute [rw] strMeasure7
#   @return [String, nil]
#
# @!attribute [rw] strMeasure8
#   @return [String, nil]
#
# @!attribute [rw] strMeasure9
#   @return [String, nil]
#
# @!attribute [rw] strSource
#   @return [String, nil]
#
# @!attribute [rw] strTags
#   @return [String, nil]
#
# @!attribute [rw] strYoutube
#   @return [String, nil]
RandomListMatch = Struct.new(
  :dateModified,
  :idMeal,
  :strArea,
  :strCategory,
  :strCreativeCommonsConfirmed,
  :strDrinkAlternate,
  :strImageSource,
  :strIngredient1,
  :strIngredient10,
  :strIngredient11,
  :strIngredient12,
  :strIngredient13,
  :strIngredient14,
  :strIngredient15,
  :strIngredient16,
  :strIngredient17,
  :strIngredient18,
  :strIngredient19,
  :strIngredient2,
  :strIngredient20,
  :strIngredient3,
  :strIngredient4,
  :strIngredient5,
  :strIngredient6,
  :strIngredient7,
  :strIngredient8,
  :strIngredient9,
  :strInstructions,
  :strMeal,
  :strMealThumb,
  :strMeasure1,
  :strMeasure10,
  :strMeasure11,
  :strMeasure12,
  :strMeasure13,
  :strMeasure14,
  :strMeasure15,
  :strMeasure16,
  :strMeasure17,
  :strMeasure18,
  :strMeasure19,
  :strMeasure2,
  :strMeasure20,
  :strMeasure3,
  :strMeasure4,
  :strMeasure5,
  :strMeasure6,
  :strMeasure7,
  :strMeasure8,
  :strMeasure9,
  :strSource,
  :strTags,
  :strYoutube,
  keyword_init: true
)

# Randomselection entity data model.
#
# @!attribute [rw] dateModified
#   @return [String, nil]
#
# @!attribute [rw] idMeal
#   @return [String, nil]
#
# @!attribute [rw] strArea
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strCreativeCommonsConfirmed
#   @return [String, nil]
#
# @!attribute [rw] strDrinkAlternate
#   @return [String, nil]
#
# @!attribute [rw] strImageSource
#   @return [String, nil]
#
# @!attribute [rw] strIngredient1
#   @return [String, nil]
#
# @!attribute [rw] strIngredient10
#   @return [String, nil]
#
# @!attribute [rw] strIngredient11
#   @return [String, nil]
#
# @!attribute [rw] strIngredient12
#   @return [String, nil]
#
# @!attribute [rw] strIngredient13
#   @return [String, nil]
#
# @!attribute [rw] strIngredient14
#   @return [String, nil]
#
# @!attribute [rw] strIngredient15
#   @return [String, nil]
#
# @!attribute [rw] strIngredient16
#   @return [String, nil]
#
# @!attribute [rw] strIngredient17
#   @return [String, nil]
#
# @!attribute [rw] strIngredient18
#   @return [String, nil]
#
# @!attribute [rw] strIngredient19
#   @return [String, nil]
#
# @!attribute [rw] strIngredient2
#   @return [String, nil]
#
# @!attribute [rw] strIngredient20
#   @return [String, nil]
#
# @!attribute [rw] strIngredient3
#   @return [String, nil]
#
# @!attribute [rw] strIngredient4
#   @return [String, nil]
#
# @!attribute [rw] strIngredient5
#   @return [String, nil]
#
# @!attribute [rw] strIngredient6
#   @return [String, nil]
#
# @!attribute [rw] strIngredient7
#   @return [String, nil]
#
# @!attribute [rw] strIngredient8
#   @return [String, nil]
#
# @!attribute [rw] strIngredient9
#   @return [String, nil]
#
# @!attribute [rw] strInstructions
#   @return [String, nil]
#
# @!attribute [rw] strMeal
#   @return [String, nil]
#
# @!attribute [rw] strMealThumb
#   @return [String, nil]
#
# @!attribute [rw] strMeasure1
#   @return [String, nil]
#
# @!attribute [rw] strMeasure10
#   @return [String, nil]
#
# @!attribute [rw] strMeasure11
#   @return [String, nil]
#
# @!attribute [rw] strMeasure12
#   @return [String, nil]
#
# @!attribute [rw] strMeasure13
#   @return [String, nil]
#
# @!attribute [rw] strMeasure14
#   @return [String, nil]
#
# @!attribute [rw] strMeasure15
#   @return [String, nil]
#
# @!attribute [rw] strMeasure16
#   @return [String, nil]
#
# @!attribute [rw] strMeasure17
#   @return [String, nil]
#
# @!attribute [rw] strMeasure18
#   @return [String, nil]
#
# @!attribute [rw] strMeasure19
#   @return [String, nil]
#
# @!attribute [rw] strMeasure2
#   @return [String, nil]
#
# @!attribute [rw] strMeasure20
#   @return [String, nil]
#
# @!attribute [rw] strMeasure3
#   @return [String, nil]
#
# @!attribute [rw] strMeasure4
#   @return [String, nil]
#
# @!attribute [rw] strMeasure5
#   @return [String, nil]
#
# @!attribute [rw] strMeasure6
#   @return [String, nil]
#
# @!attribute [rw] strMeasure7
#   @return [String, nil]
#
# @!attribute [rw] strMeasure8
#   @return [String, nil]
#
# @!attribute [rw] strMeasure9
#   @return [String, nil]
#
# @!attribute [rw] strSource
#   @return [String, nil]
#
# @!attribute [rw] strTags
#   @return [String, nil]
#
# @!attribute [rw] strYoutube
#   @return [String, nil]
Randomselection = Struct.new(
  :dateModified,
  :idMeal,
  :strArea,
  :strCategory,
  :strCreativeCommonsConfirmed,
  :strDrinkAlternate,
  :strImageSource,
  :strIngredient1,
  :strIngredient10,
  :strIngredient11,
  :strIngredient12,
  :strIngredient13,
  :strIngredient14,
  :strIngredient15,
  :strIngredient16,
  :strIngredient17,
  :strIngredient18,
  :strIngredient19,
  :strIngredient2,
  :strIngredient20,
  :strIngredient3,
  :strIngredient4,
  :strIngredient5,
  :strIngredient6,
  :strIngredient7,
  :strIngredient8,
  :strIngredient9,
  :strInstructions,
  :strMeal,
  :strMealThumb,
  :strMeasure1,
  :strMeasure10,
  :strMeasure11,
  :strMeasure12,
  :strMeasure13,
  :strMeasure14,
  :strMeasure15,
  :strMeasure16,
  :strMeasure17,
  :strMeasure18,
  :strMeasure19,
  :strMeasure2,
  :strMeasure20,
  :strMeasure3,
  :strMeasure4,
  :strMeasure5,
  :strMeasure6,
  :strMeasure7,
  :strMeasure8,
  :strMeasure9,
  :strSource,
  :strTags,
  :strYoutube,
  keyword_init: true
)

# Request payload for Randomselection#list.
#
# @!attribute [rw] dateModified
#   @return [String, nil]
#
# @!attribute [rw] idMeal
#   @return [String, nil]
#
# @!attribute [rw] strArea
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strCreativeCommonsConfirmed
#   @return [String, nil]
#
# @!attribute [rw] strDrinkAlternate
#   @return [String, nil]
#
# @!attribute [rw] strImageSource
#   @return [String, nil]
#
# @!attribute [rw] strIngredient1
#   @return [String, nil]
#
# @!attribute [rw] strIngredient10
#   @return [String, nil]
#
# @!attribute [rw] strIngredient11
#   @return [String, nil]
#
# @!attribute [rw] strIngredient12
#   @return [String, nil]
#
# @!attribute [rw] strIngredient13
#   @return [String, nil]
#
# @!attribute [rw] strIngredient14
#   @return [String, nil]
#
# @!attribute [rw] strIngredient15
#   @return [String, nil]
#
# @!attribute [rw] strIngredient16
#   @return [String, nil]
#
# @!attribute [rw] strIngredient17
#   @return [String, nil]
#
# @!attribute [rw] strIngredient18
#   @return [String, nil]
#
# @!attribute [rw] strIngredient19
#   @return [String, nil]
#
# @!attribute [rw] strIngredient2
#   @return [String, nil]
#
# @!attribute [rw] strIngredient20
#   @return [String, nil]
#
# @!attribute [rw] strIngredient3
#   @return [String, nil]
#
# @!attribute [rw] strIngredient4
#   @return [String, nil]
#
# @!attribute [rw] strIngredient5
#   @return [String, nil]
#
# @!attribute [rw] strIngredient6
#   @return [String, nil]
#
# @!attribute [rw] strIngredient7
#   @return [String, nil]
#
# @!attribute [rw] strIngredient8
#   @return [String, nil]
#
# @!attribute [rw] strIngredient9
#   @return [String, nil]
#
# @!attribute [rw] strInstructions
#   @return [String, nil]
#
# @!attribute [rw] strMeal
#   @return [String, nil]
#
# @!attribute [rw] strMealThumb
#   @return [String, nil]
#
# @!attribute [rw] strMeasure1
#   @return [String, nil]
#
# @!attribute [rw] strMeasure10
#   @return [String, nil]
#
# @!attribute [rw] strMeasure11
#   @return [String, nil]
#
# @!attribute [rw] strMeasure12
#   @return [String, nil]
#
# @!attribute [rw] strMeasure13
#   @return [String, nil]
#
# @!attribute [rw] strMeasure14
#   @return [String, nil]
#
# @!attribute [rw] strMeasure15
#   @return [String, nil]
#
# @!attribute [rw] strMeasure16
#   @return [String, nil]
#
# @!attribute [rw] strMeasure17
#   @return [String, nil]
#
# @!attribute [rw] strMeasure18
#   @return [String, nil]
#
# @!attribute [rw] strMeasure19
#   @return [String, nil]
#
# @!attribute [rw] strMeasure2
#   @return [String, nil]
#
# @!attribute [rw] strMeasure20
#   @return [String, nil]
#
# @!attribute [rw] strMeasure3
#   @return [String, nil]
#
# @!attribute [rw] strMeasure4
#   @return [String, nil]
#
# @!attribute [rw] strMeasure5
#   @return [String, nil]
#
# @!attribute [rw] strMeasure6
#   @return [String, nil]
#
# @!attribute [rw] strMeasure7
#   @return [String, nil]
#
# @!attribute [rw] strMeasure8
#   @return [String, nil]
#
# @!attribute [rw] strMeasure9
#   @return [String, nil]
#
# @!attribute [rw] strSource
#   @return [String, nil]
#
# @!attribute [rw] strTags
#   @return [String, nil]
#
# @!attribute [rw] strYoutube
#   @return [String, nil]
RandomselectionListMatch = Struct.new(
  :dateModified,
  :idMeal,
  :strArea,
  :strCategory,
  :strCreativeCommonsConfirmed,
  :strDrinkAlternate,
  :strImageSource,
  :strIngredient1,
  :strIngredient10,
  :strIngredient11,
  :strIngredient12,
  :strIngredient13,
  :strIngredient14,
  :strIngredient15,
  :strIngredient16,
  :strIngredient17,
  :strIngredient18,
  :strIngredient19,
  :strIngredient2,
  :strIngredient20,
  :strIngredient3,
  :strIngredient4,
  :strIngredient5,
  :strIngredient6,
  :strIngredient7,
  :strIngredient8,
  :strIngredient9,
  :strInstructions,
  :strMeal,
  :strMealThumb,
  :strMeasure1,
  :strMeasure10,
  :strMeasure11,
  :strMeasure12,
  :strMeasure13,
  :strMeasure14,
  :strMeasure15,
  :strMeasure16,
  :strMeasure17,
  :strMeasure18,
  :strMeasure19,
  :strMeasure2,
  :strMeasure20,
  :strMeasure3,
  :strMeasure4,
  :strMeasure5,
  :strMeasure6,
  :strMeasure7,
  :strMeasure8,
  :strMeasure9,
  :strSource,
  :strTags,
  :strYoutube,
  keyword_init: true
)

# Search entity data model.
#
# @!attribute [rw] dateModified
#   @return [String, nil]
#
# @!attribute [rw] idMeal
#   @return [String, nil]
#
# @!attribute [rw] strArea
#   @return [String, nil]
#
# @!attribute [rw] strCategory
#   @return [String, nil]
#
# @!attribute [rw] strCreativeCommonsConfirmed
#   @return [String, nil]
#
# @!attribute [rw] strDrinkAlternate
#   @return [String, nil]
#
# @!attribute [rw] strImageSource
#   @return [String, nil]
#
# @!attribute [rw] strIngredient1
#   @return [String, nil]
#
# @!attribute [rw] strIngredient10
#   @return [String, nil]
#
# @!attribute [rw] strIngredient11
#   @return [String, nil]
#
# @!attribute [rw] strIngredient12
#   @return [String, nil]
#
# @!attribute [rw] strIngredient13
#   @return [String, nil]
#
# @!attribute [rw] strIngredient14
#   @return [String, nil]
#
# @!attribute [rw] strIngredient15
#   @return [String, nil]
#
# @!attribute [rw] strIngredient16
#   @return [String, nil]
#
# @!attribute [rw] strIngredient17
#   @return [String, nil]
#
# @!attribute [rw] strIngredient18
#   @return [String, nil]
#
# @!attribute [rw] strIngredient19
#   @return [String, nil]
#
# @!attribute [rw] strIngredient2
#   @return [String, nil]
#
# @!attribute [rw] strIngredient20
#   @return [String, nil]
#
# @!attribute [rw] strIngredient3
#   @return [String, nil]
#
# @!attribute [rw] strIngredient4
#   @return [String, nil]
#
# @!attribute [rw] strIngredient5
#   @return [String, nil]
#
# @!attribute [rw] strIngredient6
#   @return [String, nil]
#
# @!attribute [rw] strIngredient7
#   @return [String, nil]
#
# @!attribute [rw] strIngredient8
#   @return [String, nil]
#
# @!attribute [rw] strIngredient9
#   @return [String, nil]
#
# @!attribute [rw] strInstructions
#   @return [String, nil]
#
# @!attribute [rw] strMeal
#   @return [String, nil]
#
# @!attribute [rw] strMealThumb
#   @return [String, nil]
#
# @!attribute [rw] strMeasure1
#   @return [String, nil]
#
# @!attribute [rw] strMeasure10
#   @return [String, nil]
#
# @!attribute [rw] strMeasure11
#   @return [String, nil]
#
# @!attribute [rw] strMeasure12
#   @return [String, nil]
#
# @!attribute [rw] strMeasure13
#   @return [String, nil]
#
# @!attribute [rw] strMeasure14
#   @return [String, nil]
#
# @!attribute [rw] strMeasure15
#   @return [String, nil]
#
# @!attribute [rw] strMeasure16
#   @return [String, nil]
#
# @!attribute [rw] strMeasure17
#   @return [String, nil]
#
# @!attribute [rw] strMeasure18
#   @return [String, nil]
#
# @!attribute [rw] strMeasure19
#   @return [String, nil]
#
# @!attribute [rw] strMeasure2
#   @return [String, nil]
#
# @!attribute [rw] strMeasure20
#   @return [String, nil]
#
# @!attribute [rw] strMeasure3
#   @return [String, nil]
#
# @!attribute [rw] strMeasure4
#   @return [String, nil]
#
# @!attribute [rw] strMeasure5
#   @return [String, nil]
#
# @!attribute [rw] strMeasure6
#   @return [String, nil]
#
# @!attribute [rw] strMeasure7
#   @return [String, nil]
#
# @!attribute [rw] strMeasure8
#   @return [String, nil]
#
# @!attribute [rw] strMeasure9
#   @return [String, nil]
#
# @!attribute [rw] strSource
#   @return [String, nil]
#
# @!attribute [rw] strTags
#   @return [String, nil]
#
# @!attribute [rw] strYoutube
#   @return [String, nil]
Search = Struct.new(
  :dateModified,
  :idMeal,
  :strArea,
  :strCategory,
  :strCreativeCommonsConfirmed,
  :strDrinkAlternate,
  :strImageSource,
  :strIngredient1,
  :strIngredient10,
  :strIngredient11,
  :strIngredient12,
  :strIngredient13,
  :strIngredient14,
  :strIngredient15,
  :strIngredient16,
  :strIngredient17,
  :strIngredient18,
  :strIngredient19,
  :strIngredient2,
  :strIngredient20,
  :strIngredient3,
  :strIngredient4,
  :strIngredient5,
  :strIngredient6,
  :strIngredient7,
  :strIngredient8,
  :strIngredient9,
  :strInstructions,
  :strMeal,
  :strMealThumb,
  :strMeasure1,
  :strMeasure10,
  :strMeasure11,
  :strMeasure12,
  :strMeasure13,
  :strMeasure14,
  :strMeasure15,
  :strMeasure16,
  :strMeasure17,
  :strMeasure18,
  :strMeasure19,
  :strMeasure2,
  :strMeasure20,
  :strMeasure3,
  :strMeasure4,
  :strMeasure5,
  :strMeasure6,
  :strMeasure7,
  :strMeasure8,
  :strMeasure9,
  :strSource,
  :strTags,
  :strYoutube,
  keyword_init: true
)

# Request payload for Search#list.
#
# @!attribute [rw] f
#   @return [String, nil]
#
# @!attribute [rw] s
#   @return [String, nil]
SearchListMatch = Struct.new(
  :f,
  :s,
  keyword_init: true
)


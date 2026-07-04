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
# @!attribute [rw] id_category
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_category_description
#   @return [String, nil]
#
# @!attribute [rw] str_category_thumb
#   @return [String, nil]
Category = Struct.new(
  :id_category,
  :str_category,
  :str_category_description,
  :str_category_thumb,
  keyword_init: true
)

# Match filter for Category#list (any subset of Category fields).
#
# @!attribute [rw] id_category
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_category_description
#   @return [String, nil]
#
# @!attribute [rw] str_category_thumb
#   @return [String, nil]
CategoryListMatch = Struct.new(
  :id_category,
  :str_category,
  :str_category_description,
  :str_category_thumb,
  keyword_init: true
)

# Filter entity data model.
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
Filter = Struct.new(
  :id_meal,
  :str_meal,
  :str_meal_thumb,
  keyword_init: true
)

# Match filter for Filter#list (any subset of Filter fields).
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
FilterListMatch = Struct.new(
  :id_meal,
  :str_meal,
  :str_meal_thumb,
  keyword_init: true
)

# Latest entity data model.
#
# @!attribute [rw] date_modified
#   @return [String, nil]
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_creative_commons_confirmed
#   @return [String, nil]
#
# @!attribute [rw] str_drink_alternate
#   @return [String, nil]
#
# @!attribute [rw] str_image_source
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient1
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient10
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient11
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient12
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient13
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient14
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient15
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient16
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient17
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient18
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient19
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient2
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient20
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient3
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient4
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient5
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient6
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient7
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient8
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient9
#   @return [String, nil]
#
# @!attribute [rw] str_instruction
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
#
# @!attribute [rw] str_measure1
#   @return [String, nil]
#
# @!attribute [rw] str_measure10
#   @return [String, nil]
#
# @!attribute [rw] str_measure11
#   @return [String, nil]
#
# @!attribute [rw] str_measure12
#   @return [String, nil]
#
# @!attribute [rw] str_measure13
#   @return [String, nil]
#
# @!attribute [rw] str_measure14
#   @return [String, nil]
#
# @!attribute [rw] str_measure15
#   @return [String, nil]
#
# @!attribute [rw] str_measure16
#   @return [String, nil]
#
# @!attribute [rw] str_measure17
#   @return [String, nil]
#
# @!attribute [rw] str_measure18
#   @return [String, nil]
#
# @!attribute [rw] str_measure19
#   @return [String, nil]
#
# @!attribute [rw] str_measure2
#   @return [String, nil]
#
# @!attribute [rw] str_measure20
#   @return [String, nil]
#
# @!attribute [rw] str_measure3
#   @return [String, nil]
#
# @!attribute [rw] str_measure4
#   @return [String, nil]
#
# @!attribute [rw] str_measure5
#   @return [String, nil]
#
# @!attribute [rw] str_measure6
#   @return [String, nil]
#
# @!attribute [rw] str_measure7
#   @return [String, nil]
#
# @!attribute [rw] str_measure8
#   @return [String, nil]
#
# @!attribute [rw] str_measure9
#   @return [String, nil]
#
# @!attribute [rw] str_source
#   @return [String, nil]
#
# @!attribute [rw] str_tag
#   @return [String, nil]
#
# @!attribute [rw] str_youtube
#   @return [String, nil]
Latest = Struct.new(
  :date_modified,
  :id_meal,
  :str_area,
  :str_category,
  :str_creative_commons_confirmed,
  :str_drink_alternate,
  :str_image_source,
  :str_ingredient1,
  :str_ingredient10,
  :str_ingredient11,
  :str_ingredient12,
  :str_ingredient13,
  :str_ingredient14,
  :str_ingredient15,
  :str_ingredient16,
  :str_ingredient17,
  :str_ingredient18,
  :str_ingredient19,
  :str_ingredient2,
  :str_ingredient20,
  :str_ingredient3,
  :str_ingredient4,
  :str_ingredient5,
  :str_ingredient6,
  :str_ingredient7,
  :str_ingredient8,
  :str_ingredient9,
  :str_instruction,
  :str_meal,
  :str_meal_thumb,
  :str_measure1,
  :str_measure10,
  :str_measure11,
  :str_measure12,
  :str_measure13,
  :str_measure14,
  :str_measure15,
  :str_measure16,
  :str_measure17,
  :str_measure18,
  :str_measure19,
  :str_measure2,
  :str_measure20,
  :str_measure3,
  :str_measure4,
  :str_measure5,
  :str_measure6,
  :str_measure7,
  :str_measure8,
  :str_measure9,
  :str_source,
  :str_tag,
  :str_youtube,
  keyword_init: true
)

# Match filter for Latest#list (any subset of Latest fields).
#
# @!attribute [rw] date_modified
#   @return [String, nil]
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_creative_commons_confirmed
#   @return [String, nil]
#
# @!attribute [rw] str_drink_alternate
#   @return [String, nil]
#
# @!attribute [rw] str_image_source
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient1
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient10
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient11
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient12
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient13
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient14
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient15
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient16
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient17
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient18
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient19
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient2
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient20
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient3
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient4
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient5
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient6
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient7
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient8
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient9
#   @return [String, nil]
#
# @!attribute [rw] str_instruction
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
#
# @!attribute [rw] str_measure1
#   @return [String, nil]
#
# @!attribute [rw] str_measure10
#   @return [String, nil]
#
# @!attribute [rw] str_measure11
#   @return [String, nil]
#
# @!attribute [rw] str_measure12
#   @return [String, nil]
#
# @!attribute [rw] str_measure13
#   @return [String, nil]
#
# @!attribute [rw] str_measure14
#   @return [String, nil]
#
# @!attribute [rw] str_measure15
#   @return [String, nil]
#
# @!attribute [rw] str_measure16
#   @return [String, nil]
#
# @!attribute [rw] str_measure17
#   @return [String, nil]
#
# @!attribute [rw] str_measure18
#   @return [String, nil]
#
# @!attribute [rw] str_measure19
#   @return [String, nil]
#
# @!attribute [rw] str_measure2
#   @return [String, nil]
#
# @!attribute [rw] str_measure20
#   @return [String, nil]
#
# @!attribute [rw] str_measure3
#   @return [String, nil]
#
# @!attribute [rw] str_measure4
#   @return [String, nil]
#
# @!attribute [rw] str_measure5
#   @return [String, nil]
#
# @!attribute [rw] str_measure6
#   @return [String, nil]
#
# @!attribute [rw] str_measure7
#   @return [String, nil]
#
# @!attribute [rw] str_measure8
#   @return [String, nil]
#
# @!attribute [rw] str_measure9
#   @return [String, nil]
#
# @!attribute [rw] str_source
#   @return [String, nil]
#
# @!attribute [rw] str_tag
#   @return [String, nil]
#
# @!attribute [rw] str_youtube
#   @return [String, nil]
LatestListMatch = Struct.new(
  :date_modified,
  :id_meal,
  :str_area,
  :str_category,
  :str_creative_commons_confirmed,
  :str_drink_alternate,
  :str_image_source,
  :str_ingredient1,
  :str_ingredient10,
  :str_ingredient11,
  :str_ingredient12,
  :str_ingredient13,
  :str_ingredient14,
  :str_ingredient15,
  :str_ingredient16,
  :str_ingredient17,
  :str_ingredient18,
  :str_ingredient19,
  :str_ingredient2,
  :str_ingredient20,
  :str_ingredient3,
  :str_ingredient4,
  :str_ingredient5,
  :str_ingredient6,
  :str_ingredient7,
  :str_ingredient8,
  :str_ingredient9,
  :str_instruction,
  :str_meal,
  :str_meal_thumb,
  :str_measure1,
  :str_measure10,
  :str_measure11,
  :str_measure12,
  :str_measure13,
  :str_measure14,
  :str_measure15,
  :str_measure16,
  :str_measure17,
  :str_measure18,
  :str_measure19,
  :str_measure2,
  :str_measure20,
  :str_measure3,
  :str_measure4,
  :str_measure5,
  :str_measure6,
  :str_measure7,
  :str_measure8,
  :str_measure9,
  :str_source,
  :str_tag,
  :str_youtube,
  keyword_init: true
)

# List entity data model.
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient
#   @return [String, nil]
List = Struct.new(
  :str_area,
  :str_category,
  :str_ingredient,
  keyword_init: true
)

# Match filter for List#list (any subset of List fields).
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient
#   @return [String, nil]
ListListMatch = Struct.new(
  :str_area,
  :str_category,
  :str_ingredient,
  keyword_init: true
)

# Lookup entity data model.
#
# @!attribute [rw] date_modified
#   @return [String, nil]
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_creative_commons_confirmed
#   @return [String, nil]
#
# @!attribute [rw] str_drink_alternate
#   @return [String, nil]
#
# @!attribute [rw] str_image_source
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient1
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient10
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient11
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient12
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient13
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient14
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient15
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient16
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient17
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient18
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient19
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient2
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient20
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient3
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient4
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient5
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient6
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient7
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient8
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient9
#   @return [String, nil]
#
# @!attribute [rw] str_instruction
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
#
# @!attribute [rw] str_measure1
#   @return [String, nil]
#
# @!attribute [rw] str_measure10
#   @return [String, nil]
#
# @!attribute [rw] str_measure11
#   @return [String, nil]
#
# @!attribute [rw] str_measure12
#   @return [String, nil]
#
# @!attribute [rw] str_measure13
#   @return [String, nil]
#
# @!attribute [rw] str_measure14
#   @return [String, nil]
#
# @!attribute [rw] str_measure15
#   @return [String, nil]
#
# @!attribute [rw] str_measure16
#   @return [String, nil]
#
# @!attribute [rw] str_measure17
#   @return [String, nil]
#
# @!attribute [rw] str_measure18
#   @return [String, nil]
#
# @!attribute [rw] str_measure19
#   @return [String, nil]
#
# @!attribute [rw] str_measure2
#   @return [String, nil]
#
# @!attribute [rw] str_measure20
#   @return [String, nil]
#
# @!attribute [rw] str_measure3
#   @return [String, nil]
#
# @!attribute [rw] str_measure4
#   @return [String, nil]
#
# @!attribute [rw] str_measure5
#   @return [String, nil]
#
# @!attribute [rw] str_measure6
#   @return [String, nil]
#
# @!attribute [rw] str_measure7
#   @return [String, nil]
#
# @!attribute [rw] str_measure8
#   @return [String, nil]
#
# @!attribute [rw] str_measure9
#   @return [String, nil]
#
# @!attribute [rw] str_source
#   @return [String, nil]
#
# @!attribute [rw] str_tag
#   @return [String, nil]
#
# @!attribute [rw] str_youtube
#   @return [String, nil]
Lookup = Struct.new(
  :date_modified,
  :id_meal,
  :str_area,
  :str_category,
  :str_creative_commons_confirmed,
  :str_drink_alternate,
  :str_image_source,
  :str_ingredient1,
  :str_ingredient10,
  :str_ingredient11,
  :str_ingredient12,
  :str_ingredient13,
  :str_ingredient14,
  :str_ingredient15,
  :str_ingredient16,
  :str_ingredient17,
  :str_ingredient18,
  :str_ingredient19,
  :str_ingredient2,
  :str_ingredient20,
  :str_ingredient3,
  :str_ingredient4,
  :str_ingredient5,
  :str_ingredient6,
  :str_ingredient7,
  :str_ingredient8,
  :str_ingredient9,
  :str_instruction,
  :str_meal,
  :str_meal_thumb,
  :str_measure1,
  :str_measure10,
  :str_measure11,
  :str_measure12,
  :str_measure13,
  :str_measure14,
  :str_measure15,
  :str_measure16,
  :str_measure17,
  :str_measure18,
  :str_measure19,
  :str_measure2,
  :str_measure20,
  :str_measure3,
  :str_measure4,
  :str_measure5,
  :str_measure6,
  :str_measure7,
  :str_measure8,
  :str_measure9,
  :str_source,
  :str_tag,
  :str_youtube,
  keyword_init: true
)

# Match filter for Lookup#list (any subset of Lookup fields).
#
# @!attribute [rw] date_modified
#   @return [String, nil]
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_creative_commons_confirmed
#   @return [String, nil]
#
# @!attribute [rw] str_drink_alternate
#   @return [String, nil]
#
# @!attribute [rw] str_image_source
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient1
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient10
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient11
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient12
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient13
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient14
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient15
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient16
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient17
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient18
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient19
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient2
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient20
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient3
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient4
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient5
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient6
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient7
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient8
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient9
#   @return [String, nil]
#
# @!attribute [rw] str_instruction
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
#
# @!attribute [rw] str_measure1
#   @return [String, nil]
#
# @!attribute [rw] str_measure10
#   @return [String, nil]
#
# @!attribute [rw] str_measure11
#   @return [String, nil]
#
# @!attribute [rw] str_measure12
#   @return [String, nil]
#
# @!attribute [rw] str_measure13
#   @return [String, nil]
#
# @!attribute [rw] str_measure14
#   @return [String, nil]
#
# @!attribute [rw] str_measure15
#   @return [String, nil]
#
# @!attribute [rw] str_measure16
#   @return [String, nil]
#
# @!attribute [rw] str_measure17
#   @return [String, nil]
#
# @!attribute [rw] str_measure18
#   @return [String, nil]
#
# @!attribute [rw] str_measure19
#   @return [String, nil]
#
# @!attribute [rw] str_measure2
#   @return [String, nil]
#
# @!attribute [rw] str_measure20
#   @return [String, nil]
#
# @!attribute [rw] str_measure3
#   @return [String, nil]
#
# @!attribute [rw] str_measure4
#   @return [String, nil]
#
# @!attribute [rw] str_measure5
#   @return [String, nil]
#
# @!attribute [rw] str_measure6
#   @return [String, nil]
#
# @!attribute [rw] str_measure7
#   @return [String, nil]
#
# @!attribute [rw] str_measure8
#   @return [String, nil]
#
# @!attribute [rw] str_measure9
#   @return [String, nil]
#
# @!attribute [rw] str_source
#   @return [String, nil]
#
# @!attribute [rw] str_tag
#   @return [String, nil]
#
# @!attribute [rw] str_youtube
#   @return [String, nil]
LookupListMatch = Struct.new(
  :date_modified,
  :id_meal,
  :str_area,
  :str_category,
  :str_creative_commons_confirmed,
  :str_drink_alternate,
  :str_image_source,
  :str_ingredient1,
  :str_ingredient10,
  :str_ingredient11,
  :str_ingredient12,
  :str_ingredient13,
  :str_ingredient14,
  :str_ingredient15,
  :str_ingredient16,
  :str_ingredient17,
  :str_ingredient18,
  :str_ingredient19,
  :str_ingredient2,
  :str_ingredient20,
  :str_ingredient3,
  :str_ingredient4,
  :str_ingredient5,
  :str_ingredient6,
  :str_ingredient7,
  :str_ingredient8,
  :str_ingredient9,
  :str_instruction,
  :str_meal,
  :str_meal_thumb,
  :str_measure1,
  :str_measure10,
  :str_measure11,
  :str_measure12,
  :str_measure13,
  :str_measure14,
  :str_measure15,
  :str_measure16,
  :str_measure17,
  :str_measure18,
  :str_measure19,
  :str_measure2,
  :str_measure20,
  :str_measure3,
  :str_measure4,
  :str_measure5,
  :str_measure6,
  :str_measure7,
  :str_measure8,
  :str_measure9,
  :str_source,
  :str_tag,
  :str_youtube,
  keyword_init: true
)

# Random entity data model.
#
# @!attribute [rw] date_modified
#   @return [String, nil]
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_creative_commons_confirmed
#   @return [String, nil]
#
# @!attribute [rw] str_drink_alternate
#   @return [String, nil]
#
# @!attribute [rw] str_image_source
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient1
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient10
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient11
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient12
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient13
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient14
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient15
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient16
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient17
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient18
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient19
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient2
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient20
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient3
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient4
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient5
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient6
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient7
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient8
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient9
#   @return [String, nil]
#
# @!attribute [rw] str_instruction
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
#
# @!attribute [rw] str_measure1
#   @return [String, nil]
#
# @!attribute [rw] str_measure10
#   @return [String, nil]
#
# @!attribute [rw] str_measure11
#   @return [String, nil]
#
# @!attribute [rw] str_measure12
#   @return [String, nil]
#
# @!attribute [rw] str_measure13
#   @return [String, nil]
#
# @!attribute [rw] str_measure14
#   @return [String, nil]
#
# @!attribute [rw] str_measure15
#   @return [String, nil]
#
# @!attribute [rw] str_measure16
#   @return [String, nil]
#
# @!attribute [rw] str_measure17
#   @return [String, nil]
#
# @!attribute [rw] str_measure18
#   @return [String, nil]
#
# @!attribute [rw] str_measure19
#   @return [String, nil]
#
# @!attribute [rw] str_measure2
#   @return [String, nil]
#
# @!attribute [rw] str_measure20
#   @return [String, nil]
#
# @!attribute [rw] str_measure3
#   @return [String, nil]
#
# @!attribute [rw] str_measure4
#   @return [String, nil]
#
# @!attribute [rw] str_measure5
#   @return [String, nil]
#
# @!attribute [rw] str_measure6
#   @return [String, nil]
#
# @!attribute [rw] str_measure7
#   @return [String, nil]
#
# @!attribute [rw] str_measure8
#   @return [String, nil]
#
# @!attribute [rw] str_measure9
#   @return [String, nil]
#
# @!attribute [rw] str_source
#   @return [String, nil]
#
# @!attribute [rw] str_tag
#   @return [String, nil]
#
# @!attribute [rw] str_youtube
#   @return [String, nil]
Random = Struct.new(
  :date_modified,
  :id_meal,
  :str_area,
  :str_category,
  :str_creative_commons_confirmed,
  :str_drink_alternate,
  :str_image_source,
  :str_ingredient1,
  :str_ingredient10,
  :str_ingredient11,
  :str_ingredient12,
  :str_ingredient13,
  :str_ingredient14,
  :str_ingredient15,
  :str_ingredient16,
  :str_ingredient17,
  :str_ingredient18,
  :str_ingredient19,
  :str_ingredient2,
  :str_ingredient20,
  :str_ingredient3,
  :str_ingredient4,
  :str_ingredient5,
  :str_ingredient6,
  :str_ingredient7,
  :str_ingredient8,
  :str_ingredient9,
  :str_instruction,
  :str_meal,
  :str_meal_thumb,
  :str_measure1,
  :str_measure10,
  :str_measure11,
  :str_measure12,
  :str_measure13,
  :str_measure14,
  :str_measure15,
  :str_measure16,
  :str_measure17,
  :str_measure18,
  :str_measure19,
  :str_measure2,
  :str_measure20,
  :str_measure3,
  :str_measure4,
  :str_measure5,
  :str_measure6,
  :str_measure7,
  :str_measure8,
  :str_measure9,
  :str_source,
  :str_tag,
  :str_youtube,
  keyword_init: true
)

# Match filter for Random#list (any subset of Random fields).
#
# @!attribute [rw] date_modified
#   @return [String, nil]
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_creative_commons_confirmed
#   @return [String, nil]
#
# @!attribute [rw] str_drink_alternate
#   @return [String, nil]
#
# @!attribute [rw] str_image_source
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient1
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient10
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient11
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient12
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient13
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient14
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient15
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient16
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient17
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient18
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient19
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient2
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient20
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient3
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient4
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient5
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient6
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient7
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient8
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient9
#   @return [String, nil]
#
# @!attribute [rw] str_instruction
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
#
# @!attribute [rw] str_measure1
#   @return [String, nil]
#
# @!attribute [rw] str_measure10
#   @return [String, nil]
#
# @!attribute [rw] str_measure11
#   @return [String, nil]
#
# @!attribute [rw] str_measure12
#   @return [String, nil]
#
# @!attribute [rw] str_measure13
#   @return [String, nil]
#
# @!attribute [rw] str_measure14
#   @return [String, nil]
#
# @!attribute [rw] str_measure15
#   @return [String, nil]
#
# @!attribute [rw] str_measure16
#   @return [String, nil]
#
# @!attribute [rw] str_measure17
#   @return [String, nil]
#
# @!attribute [rw] str_measure18
#   @return [String, nil]
#
# @!attribute [rw] str_measure19
#   @return [String, nil]
#
# @!attribute [rw] str_measure2
#   @return [String, nil]
#
# @!attribute [rw] str_measure20
#   @return [String, nil]
#
# @!attribute [rw] str_measure3
#   @return [String, nil]
#
# @!attribute [rw] str_measure4
#   @return [String, nil]
#
# @!attribute [rw] str_measure5
#   @return [String, nil]
#
# @!attribute [rw] str_measure6
#   @return [String, nil]
#
# @!attribute [rw] str_measure7
#   @return [String, nil]
#
# @!attribute [rw] str_measure8
#   @return [String, nil]
#
# @!attribute [rw] str_measure9
#   @return [String, nil]
#
# @!attribute [rw] str_source
#   @return [String, nil]
#
# @!attribute [rw] str_tag
#   @return [String, nil]
#
# @!attribute [rw] str_youtube
#   @return [String, nil]
RandomListMatch = Struct.new(
  :date_modified,
  :id_meal,
  :str_area,
  :str_category,
  :str_creative_commons_confirmed,
  :str_drink_alternate,
  :str_image_source,
  :str_ingredient1,
  :str_ingredient10,
  :str_ingredient11,
  :str_ingredient12,
  :str_ingredient13,
  :str_ingredient14,
  :str_ingredient15,
  :str_ingredient16,
  :str_ingredient17,
  :str_ingredient18,
  :str_ingredient19,
  :str_ingredient2,
  :str_ingredient20,
  :str_ingredient3,
  :str_ingredient4,
  :str_ingredient5,
  :str_ingredient6,
  :str_ingredient7,
  :str_ingredient8,
  :str_ingredient9,
  :str_instruction,
  :str_meal,
  :str_meal_thumb,
  :str_measure1,
  :str_measure10,
  :str_measure11,
  :str_measure12,
  :str_measure13,
  :str_measure14,
  :str_measure15,
  :str_measure16,
  :str_measure17,
  :str_measure18,
  :str_measure19,
  :str_measure2,
  :str_measure20,
  :str_measure3,
  :str_measure4,
  :str_measure5,
  :str_measure6,
  :str_measure7,
  :str_measure8,
  :str_measure9,
  :str_source,
  :str_tag,
  :str_youtube,
  keyword_init: true
)

# Randomselection entity data model.
#
# @!attribute [rw] date_modified
#   @return [String, nil]
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_creative_commons_confirmed
#   @return [String, nil]
#
# @!attribute [rw] str_drink_alternate
#   @return [String, nil]
#
# @!attribute [rw] str_image_source
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient1
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient10
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient11
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient12
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient13
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient14
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient15
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient16
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient17
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient18
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient19
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient2
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient20
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient3
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient4
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient5
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient6
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient7
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient8
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient9
#   @return [String, nil]
#
# @!attribute [rw] str_instruction
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
#
# @!attribute [rw] str_measure1
#   @return [String, nil]
#
# @!attribute [rw] str_measure10
#   @return [String, nil]
#
# @!attribute [rw] str_measure11
#   @return [String, nil]
#
# @!attribute [rw] str_measure12
#   @return [String, nil]
#
# @!attribute [rw] str_measure13
#   @return [String, nil]
#
# @!attribute [rw] str_measure14
#   @return [String, nil]
#
# @!attribute [rw] str_measure15
#   @return [String, nil]
#
# @!attribute [rw] str_measure16
#   @return [String, nil]
#
# @!attribute [rw] str_measure17
#   @return [String, nil]
#
# @!attribute [rw] str_measure18
#   @return [String, nil]
#
# @!attribute [rw] str_measure19
#   @return [String, nil]
#
# @!attribute [rw] str_measure2
#   @return [String, nil]
#
# @!attribute [rw] str_measure20
#   @return [String, nil]
#
# @!attribute [rw] str_measure3
#   @return [String, nil]
#
# @!attribute [rw] str_measure4
#   @return [String, nil]
#
# @!attribute [rw] str_measure5
#   @return [String, nil]
#
# @!attribute [rw] str_measure6
#   @return [String, nil]
#
# @!attribute [rw] str_measure7
#   @return [String, nil]
#
# @!attribute [rw] str_measure8
#   @return [String, nil]
#
# @!attribute [rw] str_measure9
#   @return [String, nil]
#
# @!attribute [rw] str_source
#   @return [String, nil]
#
# @!attribute [rw] str_tag
#   @return [String, nil]
#
# @!attribute [rw] str_youtube
#   @return [String, nil]
Randomselection = Struct.new(
  :date_modified,
  :id_meal,
  :str_area,
  :str_category,
  :str_creative_commons_confirmed,
  :str_drink_alternate,
  :str_image_source,
  :str_ingredient1,
  :str_ingredient10,
  :str_ingredient11,
  :str_ingredient12,
  :str_ingredient13,
  :str_ingredient14,
  :str_ingredient15,
  :str_ingredient16,
  :str_ingredient17,
  :str_ingredient18,
  :str_ingredient19,
  :str_ingredient2,
  :str_ingredient20,
  :str_ingredient3,
  :str_ingredient4,
  :str_ingredient5,
  :str_ingredient6,
  :str_ingredient7,
  :str_ingredient8,
  :str_ingredient9,
  :str_instruction,
  :str_meal,
  :str_meal_thumb,
  :str_measure1,
  :str_measure10,
  :str_measure11,
  :str_measure12,
  :str_measure13,
  :str_measure14,
  :str_measure15,
  :str_measure16,
  :str_measure17,
  :str_measure18,
  :str_measure19,
  :str_measure2,
  :str_measure20,
  :str_measure3,
  :str_measure4,
  :str_measure5,
  :str_measure6,
  :str_measure7,
  :str_measure8,
  :str_measure9,
  :str_source,
  :str_tag,
  :str_youtube,
  keyword_init: true
)

# Match filter for Randomselection#list (any subset of Randomselection fields).
#
# @!attribute [rw] date_modified
#   @return [String, nil]
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_creative_commons_confirmed
#   @return [String, nil]
#
# @!attribute [rw] str_drink_alternate
#   @return [String, nil]
#
# @!attribute [rw] str_image_source
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient1
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient10
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient11
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient12
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient13
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient14
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient15
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient16
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient17
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient18
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient19
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient2
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient20
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient3
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient4
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient5
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient6
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient7
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient8
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient9
#   @return [String, nil]
#
# @!attribute [rw] str_instruction
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
#
# @!attribute [rw] str_measure1
#   @return [String, nil]
#
# @!attribute [rw] str_measure10
#   @return [String, nil]
#
# @!attribute [rw] str_measure11
#   @return [String, nil]
#
# @!attribute [rw] str_measure12
#   @return [String, nil]
#
# @!attribute [rw] str_measure13
#   @return [String, nil]
#
# @!attribute [rw] str_measure14
#   @return [String, nil]
#
# @!attribute [rw] str_measure15
#   @return [String, nil]
#
# @!attribute [rw] str_measure16
#   @return [String, nil]
#
# @!attribute [rw] str_measure17
#   @return [String, nil]
#
# @!attribute [rw] str_measure18
#   @return [String, nil]
#
# @!attribute [rw] str_measure19
#   @return [String, nil]
#
# @!attribute [rw] str_measure2
#   @return [String, nil]
#
# @!attribute [rw] str_measure20
#   @return [String, nil]
#
# @!attribute [rw] str_measure3
#   @return [String, nil]
#
# @!attribute [rw] str_measure4
#   @return [String, nil]
#
# @!attribute [rw] str_measure5
#   @return [String, nil]
#
# @!attribute [rw] str_measure6
#   @return [String, nil]
#
# @!attribute [rw] str_measure7
#   @return [String, nil]
#
# @!attribute [rw] str_measure8
#   @return [String, nil]
#
# @!attribute [rw] str_measure9
#   @return [String, nil]
#
# @!attribute [rw] str_source
#   @return [String, nil]
#
# @!attribute [rw] str_tag
#   @return [String, nil]
#
# @!attribute [rw] str_youtube
#   @return [String, nil]
RandomselectionListMatch = Struct.new(
  :date_modified,
  :id_meal,
  :str_area,
  :str_category,
  :str_creative_commons_confirmed,
  :str_drink_alternate,
  :str_image_source,
  :str_ingredient1,
  :str_ingredient10,
  :str_ingredient11,
  :str_ingredient12,
  :str_ingredient13,
  :str_ingredient14,
  :str_ingredient15,
  :str_ingredient16,
  :str_ingredient17,
  :str_ingredient18,
  :str_ingredient19,
  :str_ingredient2,
  :str_ingredient20,
  :str_ingredient3,
  :str_ingredient4,
  :str_ingredient5,
  :str_ingredient6,
  :str_ingredient7,
  :str_ingredient8,
  :str_ingredient9,
  :str_instruction,
  :str_meal,
  :str_meal_thumb,
  :str_measure1,
  :str_measure10,
  :str_measure11,
  :str_measure12,
  :str_measure13,
  :str_measure14,
  :str_measure15,
  :str_measure16,
  :str_measure17,
  :str_measure18,
  :str_measure19,
  :str_measure2,
  :str_measure20,
  :str_measure3,
  :str_measure4,
  :str_measure5,
  :str_measure6,
  :str_measure7,
  :str_measure8,
  :str_measure9,
  :str_source,
  :str_tag,
  :str_youtube,
  keyword_init: true
)

# Search entity data model.
#
# @!attribute [rw] date_modified
#   @return [String, nil]
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_creative_commons_confirmed
#   @return [String, nil]
#
# @!attribute [rw] str_drink_alternate
#   @return [String, nil]
#
# @!attribute [rw] str_image_source
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient1
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient10
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient11
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient12
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient13
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient14
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient15
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient16
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient17
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient18
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient19
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient2
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient20
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient3
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient4
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient5
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient6
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient7
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient8
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient9
#   @return [String, nil]
#
# @!attribute [rw] str_instruction
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
#
# @!attribute [rw] str_measure1
#   @return [String, nil]
#
# @!attribute [rw] str_measure10
#   @return [String, nil]
#
# @!attribute [rw] str_measure11
#   @return [String, nil]
#
# @!attribute [rw] str_measure12
#   @return [String, nil]
#
# @!attribute [rw] str_measure13
#   @return [String, nil]
#
# @!attribute [rw] str_measure14
#   @return [String, nil]
#
# @!attribute [rw] str_measure15
#   @return [String, nil]
#
# @!attribute [rw] str_measure16
#   @return [String, nil]
#
# @!attribute [rw] str_measure17
#   @return [String, nil]
#
# @!attribute [rw] str_measure18
#   @return [String, nil]
#
# @!attribute [rw] str_measure19
#   @return [String, nil]
#
# @!attribute [rw] str_measure2
#   @return [String, nil]
#
# @!attribute [rw] str_measure20
#   @return [String, nil]
#
# @!attribute [rw] str_measure3
#   @return [String, nil]
#
# @!attribute [rw] str_measure4
#   @return [String, nil]
#
# @!attribute [rw] str_measure5
#   @return [String, nil]
#
# @!attribute [rw] str_measure6
#   @return [String, nil]
#
# @!attribute [rw] str_measure7
#   @return [String, nil]
#
# @!attribute [rw] str_measure8
#   @return [String, nil]
#
# @!attribute [rw] str_measure9
#   @return [String, nil]
#
# @!attribute [rw] str_source
#   @return [String, nil]
#
# @!attribute [rw] str_tag
#   @return [String, nil]
#
# @!attribute [rw] str_youtube
#   @return [String, nil]
Search = Struct.new(
  :date_modified,
  :id_meal,
  :str_area,
  :str_category,
  :str_creative_commons_confirmed,
  :str_drink_alternate,
  :str_image_source,
  :str_ingredient1,
  :str_ingredient10,
  :str_ingredient11,
  :str_ingredient12,
  :str_ingredient13,
  :str_ingredient14,
  :str_ingredient15,
  :str_ingredient16,
  :str_ingredient17,
  :str_ingredient18,
  :str_ingredient19,
  :str_ingredient2,
  :str_ingredient20,
  :str_ingredient3,
  :str_ingredient4,
  :str_ingredient5,
  :str_ingredient6,
  :str_ingredient7,
  :str_ingredient8,
  :str_ingredient9,
  :str_instruction,
  :str_meal,
  :str_meal_thumb,
  :str_measure1,
  :str_measure10,
  :str_measure11,
  :str_measure12,
  :str_measure13,
  :str_measure14,
  :str_measure15,
  :str_measure16,
  :str_measure17,
  :str_measure18,
  :str_measure19,
  :str_measure2,
  :str_measure20,
  :str_measure3,
  :str_measure4,
  :str_measure5,
  :str_measure6,
  :str_measure7,
  :str_measure8,
  :str_measure9,
  :str_source,
  :str_tag,
  :str_youtube,
  keyword_init: true
)

# Match filter for Search#list (any subset of Search fields).
#
# @!attribute [rw] date_modified
#   @return [String, nil]
#
# @!attribute [rw] id_meal
#   @return [String, nil]
#
# @!attribute [rw] str_area
#   @return [String, nil]
#
# @!attribute [rw] str_category
#   @return [String, nil]
#
# @!attribute [rw] str_creative_commons_confirmed
#   @return [String, nil]
#
# @!attribute [rw] str_drink_alternate
#   @return [String, nil]
#
# @!attribute [rw] str_image_source
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient1
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient10
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient11
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient12
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient13
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient14
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient15
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient16
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient17
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient18
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient19
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient2
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient20
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient3
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient4
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient5
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient6
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient7
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient8
#   @return [String, nil]
#
# @!attribute [rw] str_ingredient9
#   @return [String, nil]
#
# @!attribute [rw] str_instruction
#   @return [String, nil]
#
# @!attribute [rw] str_meal
#   @return [String, nil]
#
# @!attribute [rw] str_meal_thumb
#   @return [String, nil]
#
# @!attribute [rw] str_measure1
#   @return [String, nil]
#
# @!attribute [rw] str_measure10
#   @return [String, nil]
#
# @!attribute [rw] str_measure11
#   @return [String, nil]
#
# @!attribute [rw] str_measure12
#   @return [String, nil]
#
# @!attribute [rw] str_measure13
#   @return [String, nil]
#
# @!attribute [rw] str_measure14
#   @return [String, nil]
#
# @!attribute [rw] str_measure15
#   @return [String, nil]
#
# @!attribute [rw] str_measure16
#   @return [String, nil]
#
# @!attribute [rw] str_measure17
#   @return [String, nil]
#
# @!attribute [rw] str_measure18
#   @return [String, nil]
#
# @!attribute [rw] str_measure19
#   @return [String, nil]
#
# @!attribute [rw] str_measure2
#   @return [String, nil]
#
# @!attribute [rw] str_measure20
#   @return [String, nil]
#
# @!attribute [rw] str_measure3
#   @return [String, nil]
#
# @!attribute [rw] str_measure4
#   @return [String, nil]
#
# @!attribute [rw] str_measure5
#   @return [String, nil]
#
# @!attribute [rw] str_measure6
#   @return [String, nil]
#
# @!attribute [rw] str_measure7
#   @return [String, nil]
#
# @!attribute [rw] str_measure8
#   @return [String, nil]
#
# @!attribute [rw] str_measure9
#   @return [String, nil]
#
# @!attribute [rw] str_source
#   @return [String, nil]
#
# @!attribute [rw] str_tag
#   @return [String, nil]
#
# @!attribute [rw] str_youtube
#   @return [String, nil]
SearchListMatch = Struct.new(
  :date_modified,
  :id_meal,
  :str_area,
  :str_category,
  :str_creative_commons_confirmed,
  :str_drink_alternate,
  :str_image_source,
  :str_ingredient1,
  :str_ingredient10,
  :str_ingredient11,
  :str_ingredient12,
  :str_ingredient13,
  :str_ingredient14,
  :str_ingredient15,
  :str_ingredient16,
  :str_ingredient17,
  :str_ingredient18,
  :str_ingredient19,
  :str_ingredient2,
  :str_ingredient20,
  :str_ingredient3,
  :str_ingredient4,
  :str_ingredient5,
  :str_ingredient6,
  :str_ingredient7,
  :str_ingredient8,
  :str_ingredient9,
  :str_instruction,
  :str_meal,
  :str_meal_thumb,
  :str_measure1,
  :str_measure10,
  :str_measure11,
  :str_measure12,
  :str_measure13,
  :str_measure14,
  :str_measure15,
  :str_measure16,
  :str_measure17,
  :str_measure18,
  :str_measure19,
  :str_measure2,
  :str_measure20,
  :str_measure3,
  :str_measure4,
  :str_measure5,
  :str_measure6,
  :str_measure7,
  :str_measure8,
  :str_measure9,
  :str_source,
  :str_tag,
  :str_youtube,
  keyword_init: true
)


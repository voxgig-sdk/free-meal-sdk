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
    public ?string $id_category = null;
    public ?string $str_category = null;
    public ?string $str_category_description = null;
    public ?string $str_category_thumb = null;
}

/** Request payload for Category#list. */
class CategoryListMatch
{
    public ?string $id_category = null;
    public ?string $str_category = null;
    public ?string $str_category_description = null;
    public ?string $str_category_thumb = null;
}

/** Filter entity data model. */
class Filter
{
    public ?string $id_meal = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
}

/** Request payload for Filter#list. */
class FilterListMatch
{
    public ?string $id_meal = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
}

/** Latest entity data model. */
class Latest
{
    public ?string $date_modified = null;
    public ?string $id_meal = null;
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_creative_commons_confirmed = null;
    public ?string $str_drink_alternate = null;
    public ?string $str_image_source = null;
    public ?string $str_ingredient1 = null;
    public ?string $str_ingredient10 = null;
    public ?string $str_ingredient11 = null;
    public ?string $str_ingredient12 = null;
    public ?string $str_ingredient13 = null;
    public ?string $str_ingredient14 = null;
    public ?string $str_ingredient15 = null;
    public ?string $str_ingredient16 = null;
    public ?string $str_ingredient17 = null;
    public ?string $str_ingredient18 = null;
    public ?string $str_ingredient19 = null;
    public ?string $str_ingredient2 = null;
    public ?string $str_ingredient20 = null;
    public ?string $str_ingredient3 = null;
    public ?string $str_ingredient4 = null;
    public ?string $str_ingredient5 = null;
    public ?string $str_ingredient6 = null;
    public ?string $str_ingredient7 = null;
    public ?string $str_ingredient8 = null;
    public ?string $str_ingredient9 = null;
    public ?string $str_instruction = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
    public ?string $str_measure1 = null;
    public ?string $str_measure10 = null;
    public ?string $str_measure11 = null;
    public ?string $str_measure12 = null;
    public ?string $str_measure13 = null;
    public ?string $str_measure14 = null;
    public ?string $str_measure15 = null;
    public ?string $str_measure16 = null;
    public ?string $str_measure17 = null;
    public ?string $str_measure18 = null;
    public ?string $str_measure19 = null;
    public ?string $str_measure2 = null;
    public ?string $str_measure20 = null;
    public ?string $str_measure3 = null;
    public ?string $str_measure4 = null;
    public ?string $str_measure5 = null;
    public ?string $str_measure6 = null;
    public ?string $str_measure7 = null;
    public ?string $str_measure8 = null;
    public ?string $str_measure9 = null;
    public ?string $str_source = null;
    public ?string $str_tag = null;
    public ?string $str_youtube = null;
}

/** Request payload for Latest#list. */
class LatestListMatch
{
    public ?string $date_modified = null;
    public ?string $id_meal = null;
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_creative_commons_confirmed = null;
    public ?string $str_drink_alternate = null;
    public ?string $str_image_source = null;
    public ?string $str_ingredient1 = null;
    public ?string $str_ingredient10 = null;
    public ?string $str_ingredient11 = null;
    public ?string $str_ingredient12 = null;
    public ?string $str_ingredient13 = null;
    public ?string $str_ingredient14 = null;
    public ?string $str_ingredient15 = null;
    public ?string $str_ingredient16 = null;
    public ?string $str_ingredient17 = null;
    public ?string $str_ingredient18 = null;
    public ?string $str_ingredient19 = null;
    public ?string $str_ingredient2 = null;
    public ?string $str_ingredient20 = null;
    public ?string $str_ingredient3 = null;
    public ?string $str_ingredient4 = null;
    public ?string $str_ingredient5 = null;
    public ?string $str_ingredient6 = null;
    public ?string $str_ingredient7 = null;
    public ?string $str_ingredient8 = null;
    public ?string $str_ingredient9 = null;
    public ?string $str_instruction = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
    public ?string $str_measure1 = null;
    public ?string $str_measure10 = null;
    public ?string $str_measure11 = null;
    public ?string $str_measure12 = null;
    public ?string $str_measure13 = null;
    public ?string $str_measure14 = null;
    public ?string $str_measure15 = null;
    public ?string $str_measure16 = null;
    public ?string $str_measure17 = null;
    public ?string $str_measure18 = null;
    public ?string $str_measure19 = null;
    public ?string $str_measure2 = null;
    public ?string $str_measure20 = null;
    public ?string $str_measure3 = null;
    public ?string $str_measure4 = null;
    public ?string $str_measure5 = null;
    public ?string $str_measure6 = null;
    public ?string $str_measure7 = null;
    public ?string $str_measure8 = null;
    public ?string $str_measure9 = null;
    public ?string $str_source = null;
    public ?string $str_tag = null;
    public ?string $str_youtube = null;
}

/** List entity data model. */
class List
{
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_ingredient = null;
}

/** Request payload for List#list. */
class ListListMatch
{
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_ingredient = null;
}

/** Lookup entity data model. */
class Lookup
{
    public ?string $date_modified = null;
    public ?string $id_meal = null;
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_creative_commons_confirmed = null;
    public ?string $str_drink_alternate = null;
    public ?string $str_image_source = null;
    public ?string $str_ingredient1 = null;
    public ?string $str_ingredient10 = null;
    public ?string $str_ingredient11 = null;
    public ?string $str_ingredient12 = null;
    public ?string $str_ingredient13 = null;
    public ?string $str_ingredient14 = null;
    public ?string $str_ingredient15 = null;
    public ?string $str_ingredient16 = null;
    public ?string $str_ingredient17 = null;
    public ?string $str_ingredient18 = null;
    public ?string $str_ingredient19 = null;
    public ?string $str_ingredient2 = null;
    public ?string $str_ingredient20 = null;
    public ?string $str_ingredient3 = null;
    public ?string $str_ingredient4 = null;
    public ?string $str_ingredient5 = null;
    public ?string $str_ingredient6 = null;
    public ?string $str_ingredient7 = null;
    public ?string $str_ingredient8 = null;
    public ?string $str_ingredient9 = null;
    public ?string $str_instruction = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
    public ?string $str_measure1 = null;
    public ?string $str_measure10 = null;
    public ?string $str_measure11 = null;
    public ?string $str_measure12 = null;
    public ?string $str_measure13 = null;
    public ?string $str_measure14 = null;
    public ?string $str_measure15 = null;
    public ?string $str_measure16 = null;
    public ?string $str_measure17 = null;
    public ?string $str_measure18 = null;
    public ?string $str_measure19 = null;
    public ?string $str_measure2 = null;
    public ?string $str_measure20 = null;
    public ?string $str_measure3 = null;
    public ?string $str_measure4 = null;
    public ?string $str_measure5 = null;
    public ?string $str_measure6 = null;
    public ?string $str_measure7 = null;
    public ?string $str_measure8 = null;
    public ?string $str_measure9 = null;
    public ?string $str_source = null;
    public ?string $str_tag = null;
    public ?string $str_youtube = null;
}

/** Request payload for Lookup#list. */
class LookupListMatch
{
    public ?string $date_modified = null;
    public ?string $id_meal = null;
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_creative_commons_confirmed = null;
    public ?string $str_drink_alternate = null;
    public ?string $str_image_source = null;
    public ?string $str_ingredient1 = null;
    public ?string $str_ingredient10 = null;
    public ?string $str_ingredient11 = null;
    public ?string $str_ingredient12 = null;
    public ?string $str_ingredient13 = null;
    public ?string $str_ingredient14 = null;
    public ?string $str_ingredient15 = null;
    public ?string $str_ingredient16 = null;
    public ?string $str_ingredient17 = null;
    public ?string $str_ingredient18 = null;
    public ?string $str_ingredient19 = null;
    public ?string $str_ingredient2 = null;
    public ?string $str_ingredient20 = null;
    public ?string $str_ingredient3 = null;
    public ?string $str_ingredient4 = null;
    public ?string $str_ingredient5 = null;
    public ?string $str_ingredient6 = null;
    public ?string $str_ingredient7 = null;
    public ?string $str_ingredient8 = null;
    public ?string $str_ingredient9 = null;
    public ?string $str_instruction = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
    public ?string $str_measure1 = null;
    public ?string $str_measure10 = null;
    public ?string $str_measure11 = null;
    public ?string $str_measure12 = null;
    public ?string $str_measure13 = null;
    public ?string $str_measure14 = null;
    public ?string $str_measure15 = null;
    public ?string $str_measure16 = null;
    public ?string $str_measure17 = null;
    public ?string $str_measure18 = null;
    public ?string $str_measure19 = null;
    public ?string $str_measure2 = null;
    public ?string $str_measure20 = null;
    public ?string $str_measure3 = null;
    public ?string $str_measure4 = null;
    public ?string $str_measure5 = null;
    public ?string $str_measure6 = null;
    public ?string $str_measure7 = null;
    public ?string $str_measure8 = null;
    public ?string $str_measure9 = null;
    public ?string $str_source = null;
    public ?string $str_tag = null;
    public ?string $str_youtube = null;
}

/** Random entity data model. */
class Random
{
    public ?string $date_modified = null;
    public ?string $id_meal = null;
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_creative_commons_confirmed = null;
    public ?string $str_drink_alternate = null;
    public ?string $str_image_source = null;
    public ?string $str_ingredient1 = null;
    public ?string $str_ingredient10 = null;
    public ?string $str_ingredient11 = null;
    public ?string $str_ingredient12 = null;
    public ?string $str_ingredient13 = null;
    public ?string $str_ingredient14 = null;
    public ?string $str_ingredient15 = null;
    public ?string $str_ingredient16 = null;
    public ?string $str_ingredient17 = null;
    public ?string $str_ingredient18 = null;
    public ?string $str_ingredient19 = null;
    public ?string $str_ingredient2 = null;
    public ?string $str_ingredient20 = null;
    public ?string $str_ingredient3 = null;
    public ?string $str_ingredient4 = null;
    public ?string $str_ingredient5 = null;
    public ?string $str_ingredient6 = null;
    public ?string $str_ingredient7 = null;
    public ?string $str_ingredient8 = null;
    public ?string $str_ingredient9 = null;
    public ?string $str_instruction = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
    public ?string $str_measure1 = null;
    public ?string $str_measure10 = null;
    public ?string $str_measure11 = null;
    public ?string $str_measure12 = null;
    public ?string $str_measure13 = null;
    public ?string $str_measure14 = null;
    public ?string $str_measure15 = null;
    public ?string $str_measure16 = null;
    public ?string $str_measure17 = null;
    public ?string $str_measure18 = null;
    public ?string $str_measure19 = null;
    public ?string $str_measure2 = null;
    public ?string $str_measure20 = null;
    public ?string $str_measure3 = null;
    public ?string $str_measure4 = null;
    public ?string $str_measure5 = null;
    public ?string $str_measure6 = null;
    public ?string $str_measure7 = null;
    public ?string $str_measure8 = null;
    public ?string $str_measure9 = null;
    public ?string $str_source = null;
    public ?string $str_tag = null;
    public ?string $str_youtube = null;
}

/** Request payload for Random#list. */
class RandomListMatch
{
    public ?string $date_modified = null;
    public ?string $id_meal = null;
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_creative_commons_confirmed = null;
    public ?string $str_drink_alternate = null;
    public ?string $str_image_source = null;
    public ?string $str_ingredient1 = null;
    public ?string $str_ingredient10 = null;
    public ?string $str_ingredient11 = null;
    public ?string $str_ingredient12 = null;
    public ?string $str_ingredient13 = null;
    public ?string $str_ingredient14 = null;
    public ?string $str_ingredient15 = null;
    public ?string $str_ingredient16 = null;
    public ?string $str_ingredient17 = null;
    public ?string $str_ingredient18 = null;
    public ?string $str_ingredient19 = null;
    public ?string $str_ingredient2 = null;
    public ?string $str_ingredient20 = null;
    public ?string $str_ingredient3 = null;
    public ?string $str_ingredient4 = null;
    public ?string $str_ingredient5 = null;
    public ?string $str_ingredient6 = null;
    public ?string $str_ingredient7 = null;
    public ?string $str_ingredient8 = null;
    public ?string $str_ingredient9 = null;
    public ?string $str_instruction = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
    public ?string $str_measure1 = null;
    public ?string $str_measure10 = null;
    public ?string $str_measure11 = null;
    public ?string $str_measure12 = null;
    public ?string $str_measure13 = null;
    public ?string $str_measure14 = null;
    public ?string $str_measure15 = null;
    public ?string $str_measure16 = null;
    public ?string $str_measure17 = null;
    public ?string $str_measure18 = null;
    public ?string $str_measure19 = null;
    public ?string $str_measure2 = null;
    public ?string $str_measure20 = null;
    public ?string $str_measure3 = null;
    public ?string $str_measure4 = null;
    public ?string $str_measure5 = null;
    public ?string $str_measure6 = null;
    public ?string $str_measure7 = null;
    public ?string $str_measure8 = null;
    public ?string $str_measure9 = null;
    public ?string $str_source = null;
    public ?string $str_tag = null;
    public ?string $str_youtube = null;
}

/** Randomselection entity data model. */
class Randomselection
{
    public ?string $date_modified = null;
    public ?string $id_meal = null;
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_creative_commons_confirmed = null;
    public ?string $str_drink_alternate = null;
    public ?string $str_image_source = null;
    public ?string $str_ingredient1 = null;
    public ?string $str_ingredient10 = null;
    public ?string $str_ingredient11 = null;
    public ?string $str_ingredient12 = null;
    public ?string $str_ingredient13 = null;
    public ?string $str_ingredient14 = null;
    public ?string $str_ingredient15 = null;
    public ?string $str_ingredient16 = null;
    public ?string $str_ingredient17 = null;
    public ?string $str_ingredient18 = null;
    public ?string $str_ingredient19 = null;
    public ?string $str_ingredient2 = null;
    public ?string $str_ingredient20 = null;
    public ?string $str_ingredient3 = null;
    public ?string $str_ingredient4 = null;
    public ?string $str_ingredient5 = null;
    public ?string $str_ingredient6 = null;
    public ?string $str_ingredient7 = null;
    public ?string $str_ingredient8 = null;
    public ?string $str_ingredient9 = null;
    public ?string $str_instruction = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
    public ?string $str_measure1 = null;
    public ?string $str_measure10 = null;
    public ?string $str_measure11 = null;
    public ?string $str_measure12 = null;
    public ?string $str_measure13 = null;
    public ?string $str_measure14 = null;
    public ?string $str_measure15 = null;
    public ?string $str_measure16 = null;
    public ?string $str_measure17 = null;
    public ?string $str_measure18 = null;
    public ?string $str_measure19 = null;
    public ?string $str_measure2 = null;
    public ?string $str_measure20 = null;
    public ?string $str_measure3 = null;
    public ?string $str_measure4 = null;
    public ?string $str_measure5 = null;
    public ?string $str_measure6 = null;
    public ?string $str_measure7 = null;
    public ?string $str_measure8 = null;
    public ?string $str_measure9 = null;
    public ?string $str_source = null;
    public ?string $str_tag = null;
    public ?string $str_youtube = null;
}

/** Request payload for Randomselection#list. */
class RandomselectionListMatch
{
    public ?string $date_modified = null;
    public ?string $id_meal = null;
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_creative_commons_confirmed = null;
    public ?string $str_drink_alternate = null;
    public ?string $str_image_source = null;
    public ?string $str_ingredient1 = null;
    public ?string $str_ingredient10 = null;
    public ?string $str_ingredient11 = null;
    public ?string $str_ingredient12 = null;
    public ?string $str_ingredient13 = null;
    public ?string $str_ingredient14 = null;
    public ?string $str_ingredient15 = null;
    public ?string $str_ingredient16 = null;
    public ?string $str_ingredient17 = null;
    public ?string $str_ingredient18 = null;
    public ?string $str_ingredient19 = null;
    public ?string $str_ingredient2 = null;
    public ?string $str_ingredient20 = null;
    public ?string $str_ingredient3 = null;
    public ?string $str_ingredient4 = null;
    public ?string $str_ingredient5 = null;
    public ?string $str_ingredient6 = null;
    public ?string $str_ingredient7 = null;
    public ?string $str_ingredient8 = null;
    public ?string $str_ingredient9 = null;
    public ?string $str_instruction = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
    public ?string $str_measure1 = null;
    public ?string $str_measure10 = null;
    public ?string $str_measure11 = null;
    public ?string $str_measure12 = null;
    public ?string $str_measure13 = null;
    public ?string $str_measure14 = null;
    public ?string $str_measure15 = null;
    public ?string $str_measure16 = null;
    public ?string $str_measure17 = null;
    public ?string $str_measure18 = null;
    public ?string $str_measure19 = null;
    public ?string $str_measure2 = null;
    public ?string $str_measure20 = null;
    public ?string $str_measure3 = null;
    public ?string $str_measure4 = null;
    public ?string $str_measure5 = null;
    public ?string $str_measure6 = null;
    public ?string $str_measure7 = null;
    public ?string $str_measure8 = null;
    public ?string $str_measure9 = null;
    public ?string $str_source = null;
    public ?string $str_tag = null;
    public ?string $str_youtube = null;
}

/** Search entity data model. */
class Search
{
    public ?string $date_modified = null;
    public ?string $id_meal = null;
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_creative_commons_confirmed = null;
    public ?string $str_drink_alternate = null;
    public ?string $str_image_source = null;
    public ?string $str_ingredient1 = null;
    public ?string $str_ingredient10 = null;
    public ?string $str_ingredient11 = null;
    public ?string $str_ingredient12 = null;
    public ?string $str_ingredient13 = null;
    public ?string $str_ingredient14 = null;
    public ?string $str_ingredient15 = null;
    public ?string $str_ingredient16 = null;
    public ?string $str_ingredient17 = null;
    public ?string $str_ingredient18 = null;
    public ?string $str_ingredient19 = null;
    public ?string $str_ingredient2 = null;
    public ?string $str_ingredient20 = null;
    public ?string $str_ingredient3 = null;
    public ?string $str_ingredient4 = null;
    public ?string $str_ingredient5 = null;
    public ?string $str_ingredient6 = null;
    public ?string $str_ingredient7 = null;
    public ?string $str_ingredient8 = null;
    public ?string $str_ingredient9 = null;
    public ?string $str_instruction = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
    public ?string $str_measure1 = null;
    public ?string $str_measure10 = null;
    public ?string $str_measure11 = null;
    public ?string $str_measure12 = null;
    public ?string $str_measure13 = null;
    public ?string $str_measure14 = null;
    public ?string $str_measure15 = null;
    public ?string $str_measure16 = null;
    public ?string $str_measure17 = null;
    public ?string $str_measure18 = null;
    public ?string $str_measure19 = null;
    public ?string $str_measure2 = null;
    public ?string $str_measure20 = null;
    public ?string $str_measure3 = null;
    public ?string $str_measure4 = null;
    public ?string $str_measure5 = null;
    public ?string $str_measure6 = null;
    public ?string $str_measure7 = null;
    public ?string $str_measure8 = null;
    public ?string $str_measure9 = null;
    public ?string $str_source = null;
    public ?string $str_tag = null;
    public ?string $str_youtube = null;
}

/** Request payload for Search#list. */
class SearchListMatch
{
    public ?string $date_modified = null;
    public ?string $id_meal = null;
    public ?string $str_area = null;
    public ?string $str_category = null;
    public ?string $str_creative_commons_confirmed = null;
    public ?string $str_drink_alternate = null;
    public ?string $str_image_source = null;
    public ?string $str_ingredient1 = null;
    public ?string $str_ingredient10 = null;
    public ?string $str_ingredient11 = null;
    public ?string $str_ingredient12 = null;
    public ?string $str_ingredient13 = null;
    public ?string $str_ingredient14 = null;
    public ?string $str_ingredient15 = null;
    public ?string $str_ingredient16 = null;
    public ?string $str_ingredient17 = null;
    public ?string $str_ingredient18 = null;
    public ?string $str_ingredient19 = null;
    public ?string $str_ingredient2 = null;
    public ?string $str_ingredient20 = null;
    public ?string $str_ingredient3 = null;
    public ?string $str_ingredient4 = null;
    public ?string $str_ingredient5 = null;
    public ?string $str_ingredient6 = null;
    public ?string $str_ingredient7 = null;
    public ?string $str_ingredient8 = null;
    public ?string $str_ingredient9 = null;
    public ?string $str_instruction = null;
    public ?string $str_meal = null;
    public ?string $str_meal_thumb = null;
    public ?string $str_measure1 = null;
    public ?string $str_measure10 = null;
    public ?string $str_measure11 = null;
    public ?string $str_measure12 = null;
    public ?string $str_measure13 = null;
    public ?string $str_measure14 = null;
    public ?string $str_measure15 = null;
    public ?string $str_measure16 = null;
    public ?string $str_measure17 = null;
    public ?string $str_measure18 = null;
    public ?string $str_measure19 = null;
    public ?string $str_measure2 = null;
    public ?string $str_measure20 = null;
    public ?string $str_measure3 = null;
    public ?string $str_measure4 = null;
    public ?string $str_measure5 = null;
    public ?string $str_measure6 = null;
    public ?string $str_measure7 = null;
    public ?string $str_measure8 = null;
    public ?string $str_measure9 = null;
    public ?string $str_source = null;
    public ?string $str_tag = null;
    public ?string $str_youtube = null;
}


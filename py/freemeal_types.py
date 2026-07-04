# Typed models for the FreeMeal SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Category:
    id_category: Optional[str] = None
    str_category: Optional[str] = None
    str_category_description: Optional[str] = None
    str_category_thumb: Optional[str] = None


@dataclass
class CategoryListMatch:
    id_category: Optional[str] = None
    str_category: Optional[str] = None
    str_category_description: Optional[str] = None
    str_category_thumb: Optional[str] = None


@dataclass
class Filter:
    id_meal: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None


@dataclass
class FilterListMatch:
    id_meal: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None


@dataclass
class Latest:
    date_modified: Optional[str] = None
    id_meal: Optional[str] = None
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_creative_commons_confirmed: Optional[str] = None
    str_drink_alternate: Optional[str] = None
    str_image_source: Optional[str] = None
    str_ingredient1: Optional[str] = None
    str_ingredient10: Optional[str] = None
    str_ingredient11: Optional[str] = None
    str_ingredient12: Optional[str] = None
    str_ingredient13: Optional[str] = None
    str_ingredient14: Optional[str] = None
    str_ingredient15: Optional[str] = None
    str_ingredient16: Optional[str] = None
    str_ingredient17: Optional[str] = None
    str_ingredient18: Optional[str] = None
    str_ingredient19: Optional[str] = None
    str_ingredient2: Optional[str] = None
    str_ingredient20: Optional[str] = None
    str_ingredient3: Optional[str] = None
    str_ingredient4: Optional[str] = None
    str_ingredient5: Optional[str] = None
    str_ingredient6: Optional[str] = None
    str_ingredient7: Optional[str] = None
    str_ingredient8: Optional[str] = None
    str_ingredient9: Optional[str] = None
    str_instruction: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None
    str_measure1: Optional[str] = None
    str_measure10: Optional[str] = None
    str_measure11: Optional[str] = None
    str_measure12: Optional[str] = None
    str_measure13: Optional[str] = None
    str_measure14: Optional[str] = None
    str_measure15: Optional[str] = None
    str_measure16: Optional[str] = None
    str_measure17: Optional[str] = None
    str_measure18: Optional[str] = None
    str_measure19: Optional[str] = None
    str_measure2: Optional[str] = None
    str_measure20: Optional[str] = None
    str_measure3: Optional[str] = None
    str_measure4: Optional[str] = None
    str_measure5: Optional[str] = None
    str_measure6: Optional[str] = None
    str_measure7: Optional[str] = None
    str_measure8: Optional[str] = None
    str_measure9: Optional[str] = None
    str_source: Optional[str] = None
    str_tag: Optional[str] = None
    str_youtube: Optional[str] = None


@dataclass
class LatestListMatch:
    date_modified: Optional[str] = None
    id_meal: Optional[str] = None
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_creative_commons_confirmed: Optional[str] = None
    str_drink_alternate: Optional[str] = None
    str_image_source: Optional[str] = None
    str_ingredient1: Optional[str] = None
    str_ingredient10: Optional[str] = None
    str_ingredient11: Optional[str] = None
    str_ingredient12: Optional[str] = None
    str_ingredient13: Optional[str] = None
    str_ingredient14: Optional[str] = None
    str_ingredient15: Optional[str] = None
    str_ingredient16: Optional[str] = None
    str_ingredient17: Optional[str] = None
    str_ingredient18: Optional[str] = None
    str_ingredient19: Optional[str] = None
    str_ingredient2: Optional[str] = None
    str_ingredient20: Optional[str] = None
    str_ingredient3: Optional[str] = None
    str_ingredient4: Optional[str] = None
    str_ingredient5: Optional[str] = None
    str_ingredient6: Optional[str] = None
    str_ingredient7: Optional[str] = None
    str_ingredient8: Optional[str] = None
    str_ingredient9: Optional[str] = None
    str_instruction: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None
    str_measure1: Optional[str] = None
    str_measure10: Optional[str] = None
    str_measure11: Optional[str] = None
    str_measure12: Optional[str] = None
    str_measure13: Optional[str] = None
    str_measure14: Optional[str] = None
    str_measure15: Optional[str] = None
    str_measure16: Optional[str] = None
    str_measure17: Optional[str] = None
    str_measure18: Optional[str] = None
    str_measure19: Optional[str] = None
    str_measure2: Optional[str] = None
    str_measure20: Optional[str] = None
    str_measure3: Optional[str] = None
    str_measure4: Optional[str] = None
    str_measure5: Optional[str] = None
    str_measure6: Optional[str] = None
    str_measure7: Optional[str] = None
    str_measure8: Optional[str] = None
    str_measure9: Optional[str] = None
    str_source: Optional[str] = None
    str_tag: Optional[str] = None
    str_youtube: Optional[str] = None


@dataclass
class List:
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_ingredient: Optional[str] = None


@dataclass
class ListListMatch:
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_ingredient: Optional[str] = None


@dataclass
class Lookup:
    date_modified: Optional[str] = None
    id_meal: Optional[str] = None
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_creative_commons_confirmed: Optional[str] = None
    str_drink_alternate: Optional[str] = None
    str_image_source: Optional[str] = None
    str_ingredient1: Optional[str] = None
    str_ingredient10: Optional[str] = None
    str_ingredient11: Optional[str] = None
    str_ingredient12: Optional[str] = None
    str_ingredient13: Optional[str] = None
    str_ingredient14: Optional[str] = None
    str_ingredient15: Optional[str] = None
    str_ingredient16: Optional[str] = None
    str_ingredient17: Optional[str] = None
    str_ingredient18: Optional[str] = None
    str_ingredient19: Optional[str] = None
    str_ingredient2: Optional[str] = None
    str_ingredient20: Optional[str] = None
    str_ingredient3: Optional[str] = None
    str_ingredient4: Optional[str] = None
    str_ingredient5: Optional[str] = None
    str_ingredient6: Optional[str] = None
    str_ingredient7: Optional[str] = None
    str_ingredient8: Optional[str] = None
    str_ingredient9: Optional[str] = None
    str_instruction: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None
    str_measure1: Optional[str] = None
    str_measure10: Optional[str] = None
    str_measure11: Optional[str] = None
    str_measure12: Optional[str] = None
    str_measure13: Optional[str] = None
    str_measure14: Optional[str] = None
    str_measure15: Optional[str] = None
    str_measure16: Optional[str] = None
    str_measure17: Optional[str] = None
    str_measure18: Optional[str] = None
    str_measure19: Optional[str] = None
    str_measure2: Optional[str] = None
    str_measure20: Optional[str] = None
    str_measure3: Optional[str] = None
    str_measure4: Optional[str] = None
    str_measure5: Optional[str] = None
    str_measure6: Optional[str] = None
    str_measure7: Optional[str] = None
    str_measure8: Optional[str] = None
    str_measure9: Optional[str] = None
    str_source: Optional[str] = None
    str_tag: Optional[str] = None
    str_youtube: Optional[str] = None


@dataclass
class LookupListMatch:
    date_modified: Optional[str] = None
    id_meal: Optional[str] = None
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_creative_commons_confirmed: Optional[str] = None
    str_drink_alternate: Optional[str] = None
    str_image_source: Optional[str] = None
    str_ingredient1: Optional[str] = None
    str_ingredient10: Optional[str] = None
    str_ingredient11: Optional[str] = None
    str_ingredient12: Optional[str] = None
    str_ingredient13: Optional[str] = None
    str_ingredient14: Optional[str] = None
    str_ingredient15: Optional[str] = None
    str_ingredient16: Optional[str] = None
    str_ingredient17: Optional[str] = None
    str_ingredient18: Optional[str] = None
    str_ingredient19: Optional[str] = None
    str_ingredient2: Optional[str] = None
    str_ingredient20: Optional[str] = None
    str_ingredient3: Optional[str] = None
    str_ingredient4: Optional[str] = None
    str_ingredient5: Optional[str] = None
    str_ingredient6: Optional[str] = None
    str_ingredient7: Optional[str] = None
    str_ingredient8: Optional[str] = None
    str_ingredient9: Optional[str] = None
    str_instruction: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None
    str_measure1: Optional[str] = None
    str_measure10: Optional[str] = None
    str_measure11: Optional[str] = None
    str_measure12: Optional[str] = None
    str_measure13: Optional[str] = None
    str_measure14: Optional[str] = None
    str_measure15: Optional[str] = None
    str_measure16: Optional[str] = None
    str_measure17: Optional[str] = None
    str_measure18: Optional[str] = None
    str_measure19: Optional[str] = None
    str_measure2: Optional[str] = None
    str_measure20: Optional[str] = None
    str_measure3: Optional[str] = None
    str_measure4: Optional[str] = None
    str_measure5: Optional[str] = None
    str_measure6: Optional[str] = None
    str_measure7: Optional[str] = None
    str_measure8: Optional[str] = None
    str_measure9: Optional[str] = None
    str_source: Optional[str] = None
    str_tag: Optional[str] = None
    str_youtube: Optional[str] = None


@dataclass
class Random:
    date_modified: Optional[str] = None
    id_meal: Optional[str] = None
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_creative_commons_confirmed: Optional[str] = None
    str_drink_alternate: Optional[str] = None
    str_image_source: Optional[str] = None
    str_ingredient1: Optional[str] = None
    str_ingredient10: Optional[str] = None
    str_ingredient11: Optional[str] = None
    str_ingredient12: Optional[str] = None
    str_ingredient13: Optional[str] = None
    str_ingredient14: Optional[str] = None
    str_ingredient15: Optional[str] = None
    str_ingredient16: Optional[str] = None
    str_ingredient17: Optional[str] = None
    str_ingredient18: Optional[str] = None
    str_ingredient19: Optional[str] = None
    str_ingredient2: Optional[str] = None
    str_ingredient20: Optional[str] = None
    str_ingredient3: Optional[str] = None
    str_ingredient4: Optional[str] = None
    str_ingredient5: Optional[str] = None
    str_ingredient6: Optional[str] = None
    str_ingredient7: Optional[str] = None
    str_ingredient8: Optional[str] = None
    str_ingredient9: Optional[str] = None
    str_instruction: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None
    str_measure1: Optional[str] = None
    str_measure10: Optional[str] = None
    str_measure11: Optional[str] = None
    str_measure12: Optional[str] = None
    str_measure13: Optional[str] = None
    str_measure14: Optional[str] = None
    str_measure15: Optional[str] = None
    str_measure16: Optional[str] = None
    str_measure17: Optional[str] = None
    str_measure18: Optional[str] = None
    str_measure19: Optional[str] = None
    str_measure2: Optional[str] = None
    str_measure20: Optional[str] = None
    str_measure3: Optional[str] = None
    str_measure4: Optional[str] = None
    str_measure5: Optional[str] = None
    str_measure6: Optional[str] = None
    str_measure7: Optional[str] = None
    str_measure8: Optional[str] = None
    str_measure9: Optional[str] = None
    str_source: Optional[str] = None
    str_tag: Optional[str] = None
    str_youtube: Optional[str] = None


@dataclass
class RandomListMatch:
    date_modified: Optional[str] = None
    id_meal: Optional[str] = None
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_creative_commons_confirmed: Optional[str] = None
    str_drink_alternate: Optional[str] = None
    str_image_source: Optional[str] = None
    str_ingredient1: Optional[str] = None
    str_ingredient10: Optional[str] = None
    str_ingredient11: Optional[str] = None
    str_ingredient12: Optional[str] = None
    str_ingredient13: Optional[str] = None
    str_ingredient14: Optional[str] = None
    str_ingredient15: Optional[str] = None
    str_ingredient16: Optional[str] = None
    str_ingredient17: Optional[str] = None
    str_ingredient18: Optional[str] = None
    str_ingredient19: Optional[str] = None
    str_ingredient2: Optional[str] = None
    str_ingredient20: Optional[str] = None
    str_ingredient3: Optional[str] = None
    str_ingredient4: Optional[str] = None
    str_ingredient5: Optional[str] = None
    str_ingredient6: Optional[str] = None
    str_ingredient7: Optional[str] = None
    str_ingredient8: Optional[str] = None
    str_ingredient9: Optional[str] = None
    str_instruction: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None
    str_measure1: Optional[str] = None
    str_measure10: Optional[str] = None
    str_measure11: Optional[str] = None
    str_measure12: Optional[str] = None
    str_measure13: Optional[str] = None
    str_measure14: Optional[str] = None
    str_measure15: Optional[str] = None
    str_measure16: Optional[str] = None
    str_measure17: Optional[str] = None
    str_measure18: Optional[str] = None
    str_measure19: Optional[str] = None
    str_measure2: Optional[str] = None
    str_measure20: Optional[str] = None
    str_measure3: Optional[str] = None
    str_measure4: Optional[str] = None
    str_measure5: Optional[str] = None
    str_measure6: Optional[str] = None
    str_measure7: Optional[str] = None
    str_measure8: Optional[str] = None
    str_measure9: Optional[str] = None
    str_source: Optional[str] = None
    str_tag: Optional[str] = None
    str_youtube: Optional[str] = None


@dataclass
class Randomselection:
    date_modified: Optional[str] = None
    id_meal: Optional[str] = None
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_creative_commons_confirmed: Optional[str] = None
    str_drink_alternate: Optional[str] = None
    str_image_source: Optional[str] = None
    str_ingredient1: Optional[str] = None
    str_ingredient10: Optional[str] = None
    str_ingredient11: Optional[str] = None
    str_ingredient12: Optional[str] = None
    str_ingredient13: Optional[str] = None
    str_ingredient14: Optional[str] = None
    str_ingredient15: Optional[str] = None
    str_ingredient16: Optional[str] = None
    str_ingredient17: Optional[str] = None
    str_ingredient18: Optional[str] = None
    str_ingredient19: Optional[str] = None
    str_ingredient2: Optional[str] = None
    str_ingredient20: Optional[str] = None
    str_ingredient3: Optional[str] = None
    str_ingredient4: Optional[str] = None
    str_ingredient5: Optional[str] = None
    str_ingredient6: Optional[str] = None
    str_ingredient7: Optional[str] = None
    str_ingredient8: Optional[str] = None
    str_ingredient9: Optional[str] = None
    str_instruction: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None
    str_measure1: Optional[str] = None
    str_measure10: Optional[str] = None
    str_measure11: Optional[str] = None
    str_measure12: Optional[str] = None
    str_measure13: Optional[str] = None
    str_measure14: Optional[str] = None
    str_measure15: Optional[str] = None
    str_measure16: Optional[str] = None
    str_measure17: Optional[str] = None
    str_measure18: Optional[str] = None
    str_measure19: Optional[str] = None
    str_measure2: Optional[str] = None
    str_measure20: Optional[str] = None
    str_measure3: Optional[str] = None
    str_measure4: Optional[str] = None
    str_measure5: Optional[str] = None
    str_measure6: Optional[str] = None
    str_measure7: Optional[str] = None
    str_measure8: Optional[str] = None
    str_measure9: Optional[str] = None
    str_source: Optional[str] = None
    str_tag: Optional[str] = None
    str_youtube: Optional[str] = None


@dataclass
class RandomselectionListMatch:
    date_modified: Optional[str] = None
    id_meal: Optional[str] = None
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_creative_commons_confirmed: Optional[str] = None
    str_drink_alternate: Optional[str] = None
    str_image_source: Optional[str] = None
    str_ingredient1: Optional[str] = None
    str_ingredient10: Optional[str] = None
    str_ingredient11: Optional[str] = None
    str_ingredient12: Optional[str] = None
    str_ingredient13: Optional[str] = None
    str_ingredient14: Optional[str] = None
    str_ingredient15: Optional[str] = None
    str_ingredient16: Optional[str] = None
    str_ingredient17: Optional[str] = None
    str_ingredient18: Optional[str] = None
    str_ingredient19: Optional[str] = None
    str_ingredient2: Optional[str] = None
    str_ingredient20: Optional[str] = None
    str_ingredient3: Optional[str] = None
    str_ingredient4: Optional[str] = None
    str_ingredient5: Optional[str] = None
    str_ingredient6: Optional[str] = None
    str_ingredient7: Optional[str] = None
    str_ingredient8: Optional[str] = None
    str_ingredient9: Optional[str] = None
    str_instruction: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None
    str_measure1: Optional[str] = None
    str_measure10: Optional[str] = None
    str_measure11: Optional[str] = None
    str_measure12: Optional[str] = None
    str_measure13: Optional[str] = None
    str_measure14: Optional[str] = None
    str_measure15: Optional[str] = None
    str_measure16: Optional[str] = None
    str_measure17: Optional[str] = None
    str_measure18: Optional[str] = None
    str_measure19: Optional[str] = None
    str_measure2: Optional[str] = None
    str_measure20: Optional[str] = None
    str_measure3: Optional[str] = None
    str_measure4: Optional[str] = None
    str_measure5: Optional[str] = None
    str_measure6: Optional[str] = None
    str_measure7: Optional[str] = None
    str_measure8: Optional[str] = None
    str_measure9: Optional[str] = None
    str_source: Optional[str] = None
    str_tag: Optional[str] = None
    str_youtube: Optional[str] = None


@dataclass
class Search:
    date_modified: Optional[str] = None
    id_meal: Optional[str] = None
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_creative_commons_confirmed: Optional[str] = None
    str_drink_alternate: Optional[str] = None
    str_image_source: Optional[str] = None
    str_ingredient1: Optional[str] = None
    str_ingredient10: Optional[str] = None
    str_ingredient11: Optional[str] = None
    str_ingredient12: Optional[str] = None
    str_ingredient13: Optional[str] = None
    str_ingredient14: Optional[str] = None
    str_ingredient15: Optional[str] = None
    str_ingredient16: Optional[str] = None
    str_ingredient17: Optional[str] = None
    str_ingredient18: Optional[str] = None
    str_ingredient19: Optional[str] = None
    str_ingredient2: Optional[str] = None
    str_ingredient20: Optional[str] = None
    str_ingredient3: Optional[str] = None
    str_ingredient4: Optional[str] = None
    str_ingredient5: Optional[str] = None
    str_ingredient6: Optional[str] = None
    str_ingredient7: Optional[str] = None
    str_ingredient8: Optional[str] = None
    str_ingredient9: Optional[str] = None
    str_instruction: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None
    str_measure1: Optional[str] = None
    str_measure10: Optional[str] = None
    str_measure11: Optional[str] = None
    str_measure12: Optional[str] = None
    str_measure13: Optional[str] = None
    str_measure14: Optional[str] = None
    str_measure15: Optional[str] = None
    str_measure16: Optional[str] = None
    str_measure17: Optional[str] = None
    str_measure18: Optional[str] = None
    str_measure19: Optional[str] = None
    str_measure2: Optional[str] = None
    str_measure20: Optional[str] = None
    str_measure3: Optional[str] = None
    str_measure4: Optional[str] = None
    str_measure5: Optional[str] = None
    str_measure6: Optional[str] = None
    str_measure7: Optional[str] = None
    str_measure8: Optional[str] = None
    str_measure9: Optional[str] = None
    str_source: Optional[str] = None
    str_tag: Optional[str] = None
    str_youtube: Optional[str] = None


@dataclass
class SearchListMatch:
    date_modified: Optional[str] = None
    id_meal: Optional[str] = None
    str_area: Optional[str] = None
    str_category: Optional[str] = None
    str_creative_commons_confirmed: Optional[str] = None
    str_drink_alternate: Optional[str] = None
    str_image_source: Optional[str] = None
    str_ingredient1: Optional[str] = None
    str_ingredient10: Optional[str] = None
    str_ingredient11: Optional[str] = None
    str_ingredient12: Optional[str] = None
    str_ingredient13: Optional[str] = None
    str_ingredient14: Optional[str] = None
    str_ingredient15: Optional[str] = None
    str_ingredient16: Optional[str] = None
    str_ingredient17: Optional[str] = None
    str_ingredient18: Optional[str] = None
    str_ingredient19: Optional[str] = None
    str_ingredient2: Optional[str] = None
    str_ingredient20: Optional[str] = None
    str_ingredient3: Optional[str] = None
    str_ingredient4: Optional[str] = None
    str_ingredient5: Optional[str] = None
    str_ingredient6: Optional[str] = None
    str_ingredient7: Optional[str] = None
    str_ingredient8: Optional[str] = None
    str_ingredient9: Optional[str] = None
    str_instruction: Optional[str] = None
    str_meal: Optional[str] = None
    str_meal_thumb: Optional[str] = None
    str_measure1: Optional[str] = None
    str_measure10: Optional[str] = None
    str_measure11: Optional[str] = None
    str_measure12: Optional[str] = None
    str_measure13: Optional[str] = None
    str_measure14: Optional[str] = None
    str_measure15: Optional[str] = None
    str_measure16: Optional[str] = None
    str_measure17: Optional[str] = None
    str_measure18: Optional[str] = None
    str_measure19: Optional[str] = None
    str_measure2: Optional[str] = None
    str_measure20: Optional[str] = None
    str_measure3: Optional[str] = None
    str_measure4: Optional[str] = None
    str_measure5: Optional[str] = None
    str_measure6: Optional[str] = None
    str_measure7: Optional[str] = None
    str_measure8: Optional[str] = None
    str_measure9: Optional[str] = None
    str_source: Optional[str] = None
    str_tag: Optional[str] = None
    str_youtube: Optional[str] = None


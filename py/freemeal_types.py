# Typed models for the FreeMeal SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Category(TypedDict, total=False):
    id_category: str
    str_category: str
    str_category_description: str
    str_category_thumb: str


class CategoryListMatch(TypedDict, total=False):
    id_category: str
    str_category: str
    str_category_description: str
    str_category_thumb: str


class Filter(TypedDict, total=False):
    id_meal: str
    str_meal: str
    str_meal_thumb: str


class FilterListMatch(TypedDict, total=False):
    id_meal: str
    str_meal: str
    str_meal_thumb: str


class Latest(TypedDict, total=False):
    date_modified: str
    id_meal: str
    str_area: str
    str_category: str
    str_creative_commons_confirmed: str
    str_drink_alternate: str
    str_image_source: str
    str_ingredient1: str
    str_ingredient10: str
    str_ingredient11: str
    str_ingredient12: str
    str_ingredient13: str
    str_ingredient14: str
    str_ingredient15: str
    str_ingredient16: str
    str_ingredient17: str
    str_ingredient18: str
    str_ingredient19: str
    str_ingredient2: str
    str_ingredient20: str
    str_ingredient3: str
    str_ingredient4: str
    str_ingredient5: str
    str_ingredient6: str
    str_ingredient7: str
    str_ingredient8: str
    str_ingredient9: str
    str_instruction: str
    str_meal: str
    str_meal_thumb: str
    str_measure1: str
    str_measure10: str
    str_measure11: str
    str_measure12: str
    str_measure13: str
    str_measure14: str
    str_measure15: str
    str_measure16: str
    str_measure17: str
    str_measure18: str
    str_measure19: str
    str_measure2: str
    str_measure20: str
    str_measure3: str
    str_measure4: str
    str_measure5: str
    str_measure6: str
    str_measure7: str
    str_measure8: str
    str_measure9: str
    str_source: str
    str_tag: str
    str_youtube: str


class LatestListMatch(TypedDict, total=False):
    date_modified: str
    id_meal: str
    str_area: str
    str_category: str
    str_creative_commons_confirmed: str
    str_drink_alternate: str
    str_image_source: str
    str_ingredient1: str
    str_ingredient10: str
    str_ingredient11: str
    str_ingredient12: str
    str_ingredient13: str
    str_ingredient14: str
    str_ingredient15: str
    str_ingredient16: str
    str_ingredient17: str
    str_ingredient18: str
    str_ingredient19: str
    str_ingredient2: str
    str_ingredient20: str
    str_ingredient3: str
    str_ingredient4: str
    str_ingredient5: str
    str_ingredient6: str
    str_ingredient7: str
    str_ingredient8: str
    str_ingredient9: str
    str_instruction: str
    str_meal: str
    str_meal_thumb: str
    str_measure1: str
    str_measure10: str
    str_measure11: str
    str_measure12: str
    str_measure13: str
    str_measure14: str
    str_measure15: str
    str_measure16: str
    str_measure17: str
    str_measure18: str
    str_measure19: str
    str_measure2: str
    str_measure20: str
    str_measure3: str
    str_measure4: str
    str_measure5: str
    str_measure6: str
    str_measure7: str
    str_measure8: str
    str_measure9: str
    str_source: str
    str_tag: str
    str_youtube: str


class List(TypedDict, total=False):
    str_area: str
    str_category: str
    str_ingredient: str


class ListListMatch(TypedDict, total=False):
    str_area: str
    str_category: str
    str_ingredient: str


class Lookup(TypedDict, total=False):
    date_modified: str
    id_meal: str
    str_area: str
    str_category: str
    str_creative_commons_confirmed: str
    str_drink_alternate: str
    str_image_source: str
    str_ingredient1: str
    str_ingredient10: str
    str_ingredient11: str
    str_ingredient12: str
    str_ingredient13: str
    str_ingredient14: str
    str_ingredient15: str
    str_ingredient16: str
    str_ingredient17: str
    str_ingredient18: str
    str_ingredient19: str
    str_ingredient2: str
    str_ingredient20: str
    str_ingredient3: str
    str_ingredient4: str
    str_ingredient5: str
    str_ingredient6: str
    str_ingredient7: str
    str_ingredient8: str
    str_ingredient9: str
    str_instruction: str
    str_meal: str
    str_meal_thumb: str
    str_measure1: str
    str_measure10: str
    str_measure11: str
    str_measure12: str
    str_measure13: str
    str_measure14: str
    str_measure15: str
    str_measure16: str
    str_measure17: str
    str_measure18: str
    str_measure19: str
    str_measure2: str
    str_measure20: str
    str_measure3: str
    str_measure4: str
    str_measure5: str
    str_measure6: str
    str_measure7: str
    str_measure8: str
    str_measure9: str
    str_source: str
    str_tag: str
    str_youtube: str


class LookupListMatch(TypedDict, total=False):
    date_modified: str
    id_meal: str
    str_area: str
    str_category: str
    str_creative_commons_confirmed: str
    str_drink_alternate: str
    str_image_source: str
    str_ingredient1: str
    str_ingredient10: str
    str_ingredient11: str
    str_ingredient12: str
    str_ingredient13: str
    str_ingredient14: str
    str_ingredient15: str
    str_ingredient16: str
    str_ingredient17: str
    str_ingredient18: str
    str_ingredient19: str
    str_ingredient2: str
    str_ingredient20: str
    str_ingredient3: str
    str_ingredient4: str
    str_ingredient5: str
    str_ingredient6: str
    str_ingredient7: str
    str_ingredient8: str
    str_ingredient9: str
    str_instruction: str
    str_meal: str
    str_meal_thumb: str
    str_measure1: str
    str_measure10: str
    str_measure11: str
    str_measure12: str
    str_measure13: str
    str_measure14: str
    str_measure15: str
    str_measure16: str
    str_measure17: str
    str_measure18: str
    str_measure19: str
    str_measure2: str
    str_measure20: str
    str_measure3: str
    str_measure4: str
    str_measure5: str
    str_measure6: str
    str_measure7: str
    str_measure8: str
    str_measure9: str
    str_source: str
    str_tag: str
    str_youtube: str


class Random(TypedDict, total=False):
    date_modified: str
    id_meal: str
    str_area: str
    str_category: str
    str_creative_commons_confirmed: str
    str_drink_alternate: str
    str_image_source: str
    str_ingredient1: str
    str_ingredient10: str
    str_ingredient11: str
    str_ingredient12: str
    str_ingredient13: str
    str_ingredient14: str
    str_ingredient15: str
    str_ingredient16: str
    str_ingredient17: str
    str_ingredient18: str
    str_ingredient19: str
    str_ingredient2: str
    str_ingredient20: str
    str_ingredient3: str
    str_ingredient4: str
    str_ingredient5: str
    str_ingredient6: str
    str_ingredient7: str
    str_ingredient8: str
    str_ingredient9: str
    str_instruction: str
    str_meal: str
    str_meal_thumb: str
    str_measure1: str
    str_measure10: str
    str_measure11: str
    str_measure12: str
    str_measure13: str
    str_measure14: str
    str_measure15: str
    str_measure16: str
    str_measure17: str
    str_measure18: str
    str_measure19: str
    str_measure2: str
    str_measure20: str
    str_measure3: str
    str_measure4: str
    str_measure5: str
    str_measure6: str
    str_measure7: str
    str_measure8: str
    str_measure9: str
    str_source: str
    str_tag: str
    str_youtube: str


class RandomListMatch(TypedDict, total=False):
    date_modified: str
    id_meal: str
    str_area: str
    str_category: str
    str_creative_commons_confirmed: str
    str_drink_alternate: str
    str_image_source: str
    str_ingredient1: str
    str_ingredient10: str
    str_ingredient11: str
    str_ingredient12: str
    str_ingredient13: str
    str_ingredient14: str
    str_ingredient15: str
    str_ingredient16: str
    str_ingredient17: str
    str_ingredient18: str
    str_ingredient19: str
    str_ingredient2: str
    str_ingredient20: str
    str_ingredient3: str
    str_ingredient4: str
    str_ingredient5: str
    str_ingredient6: str
    str_ingredient7: str
    str_ingredient8: str
    str_ingredient9: str
    str_instruction: str
    str_meal: str
    str_meal_thumb: str
    str_measure1: str
    str_measure10: str
    str_measure11: str
    str_measure12: str
    str_measure13: str
    str_measure14: str
    str_measure15: str
    str_measure16: str
    str_measure17: str
    str_measure18: str
    str_measure19: str
    str_measure2: str
    str_measure20: str
    str_measure3: str
    str_measure4: str
    str_measure5: str
    str_measure6: str
    str_measure7: str
    str_measure8: str
    str_measure9: str
    str_source: str
    str_tag: str
    str_youtube: str


class Randomselection(TypedDict, total=False):
    date_modified: str
    id_meal: str
    str_area: str
    str_category: str
    str_creative_commons_confirmed: str
    str_drink_alternate: str
    str_image_source: str
    str_ingredient1: str
    str_ingredient10: str
    str_ingredient11: str
    str_ingredient12: str
    str_ingredient13: str
    str_ingredient14: str
    str_ingredient15: str
    str_ingredient16: str
    str_ingredient17: str
    str_ingredient18: str
    str_ingredient19: str
    str_ingredient2: str
    str_ingredient20: str
    str_ingredient3: str
    str_ingredient4: str
    str_ingredient5: str
    str_ingredient6: str
    str_ingredient7: str
    str_ingredient8: str
    str_ingredient9: str
    str_instruction: str
    str_meal: str
    str_meal_thumb: str
    str_measure1: str
    str_measure10: str
    str_measure11: str
    str_measure12: str
    str_measure13: str
    str_measure14: str
    str_measure15: str
    str_measure16: str
    str_measure17: str
    str_measure18: str
    str_measure19: str
    str_measure2: str
    str_measure20: str
    str_measure3: str
    str_measure4: str
    str_measure5: str
    str_measure6: str
    str_measure7: str
    str_measure8: str
    str_measure9: str
    str_source: str
    str_tag: str
    str_youtube: str


class RandomselectionListMatch(TypedDict, total=False):
    date_modified: str
    id_meal: str
    str_area: str
    str_category: str
    str_creative_commons_confirmed: str
    str_drink_alternate: str
    str_image_source: str
    str_ingredient1: str
    str_ingredient10: str
    str_ingredient11: str
    str_ingredient12: str
    str_ingredient13: str
    str_ingredient14: str
    str_ingredient15: str
    str_ingredient16: str
    str_ingredient17: str
    str_ingredient18: str
    str_ingredient19: str
    str_ingredient2: str
    str_ingredient20: str
    str_ingredient3: str
    str_ingredient4: str
    str_ingredient5: str
    str_ingredient6: str
    str_ingredient7: str
    str_ingredient8: str
    str_ingredient9: str
    str_instruction: str
    str_meal: str
    str_meal_thumb: str
    str_measure1: str
    str_measure10: str
    str_measure11: str
    str_measure12: str
    str_measure13: str
    str_measure14: str
    str_measure15: str
    str_measure16: str
    str_measure17: str
    str_measure18: str
    str_measure19: str
    str_measure2: str
    str_measure20: str
    str_measure3: str
    str_measure4: str
    str_measure5: str
    str_measure6: str
    str_measure7: str
    str_measure8: str
    str_measure9: str
    str_source: str
    str_tag: str
    str_youtube: str


class Search(TypedDict, total=False):
    date_modified: str
    id_meal: str
    str_area: str
    str_category: str
    str_creative_commons_confirmed: str
    str_drink_alternate: str
    str_image_source: str
    str_ingredient1: str
    str_ingredient10: str
    str_ingredient11: str
    str_ingredient12: str
    str_ingredient13: str
    str_ingredient14: str
    str_ingredient15: str
    str_ingredient16: str
    str_ingredient17: str
    str_ingredient18: str
    str_ingredient19: str
    str_ingredient2: str
    str_ingredient20: str
    str_ingredient3: str
    str_ingredient4: str
    str_ingredient5: str
    str_ingredient6: str
    str_ingredient7: str
    str_ingredient8: str
    str_ingredient9: str
    str_instruction: str
    str_meal: str
    str_meal_thumb: str
    str_measure1: str
    str_measure10: str
    str_measure11: str
    str_measure12: str
    str_measure13: str
    str_measure14: str
    str_measure15: str
    str_measure16: str
    str_measure17: str
    str_measure18: str
    str_measure19: str
    str_measure2: str
    str_measure20: str
    str_measure3: str
    str_measure4: str
    str_measure5: str
    str_measure6: str
    str_measure7: str
    str_measure8: str
    str_measure9: str
    str_source: str
    str_tag: str
    str_youtube: str


class SearchListMatch(TypedDict, total=False):
    date_modified: str
    id_meal: str
    str_area: str
    str_category: str
    str_creative_commons_confirmed: str
    str_drink_alternate: str
    str_image_source: str
    str_ingredient1: str
    str_ingredient10: str
    str_ingredient11: str
    str_ingredient12: str
    str_ingredient13: str
    str_ingredient14: str
    str_ingredient15: str
    str_ingredient16: str
    str_ingredient17: str
    str_ingredient18: str
    str_ingredient19: str
    str_ingredient2: str
    str_ingredient20: str
    str_ingredient3: str
    str_ingredient4: str
    str_ingredient5: str
    str_ingredient6: str
    str_ingredient7: str
    str_ingredient8: str
    str_ingredient9: str
    str_instruction: str
    str_meal: str
    str_meal_thumb: str
    str_measure1: str
    str_measure10: str
    str_measure11: str
    str_measure12: str
    str_measure13: str
    str_measure14: str
    str_measure15: str
    str_measure16: str
    str_measure17: str
    str_measure18: str
    str_measure19: str
    str_measure2: str
    str_measure20: str
    str_measure3: str
    str_measure4: str
    str_measure5: str
    str_measure6: str
    str_measure7: str
    str_measure8: str
    str_measure9: str
    str_source: str
    str_tag: str
    str_youtube: str

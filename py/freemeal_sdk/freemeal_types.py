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
    idCategory: str
    strCategory: str
    strCategoryDescription: str
    strCategoryThumb: str


class CategoryListMatch(TypedDict, total=False):
    idCategory: str
    strCategory: str
    strCategoryDescription: str
    strCategoryThumb: str


class Filter(TypedDict, total=False):
    idMeal: str
    strMeal: str
    strMealThumb: str


class FilterListMatch(TypedDict, total=False):
    idMeal: str
    strMeal: str
    strMealThumb: str


class Latest(TypedDict, total=False):
    dateModified: str
    idMeal: str
    strArea: str
    strCategory: str
    strCreativeCommonsConfirmed: str
    strDrinkAlternate: str
    strImageSource: str
    strIngredient1: str
    strIngredient10: str
    strIngredient11: str
    strIngredient12: str
    strIngredient13: str
    strIngredient14: str
    strIngredient15: str
    strIngredient16: str
    strIngredient17: str
    strIngredient18: str
    strIngredient19: str
    strIngredient2: str
    strIngredient20: str
    strIngredient3: str
    strIngredient4: str
    strIngredient5: str
    strIngredient6: str
    strIngredient7: str
    strIngredient8: str
    strIngredient9: str
    strInstructions: str
    strMeal: str
    strMealThumb: str
    strMeasure1: str
    strMeasure10: str
    strMeasure11: str
    strMeasure12: str
    strMeasure13: str
    strMeasure14: str
    strMeasure15: str
    strMeasure16: str
    strMeasure17: str
    strMeasure18: str
    strMeasure19: str
    strMeasure2: str
    strMeasure20: str
    strMeasure3: str
    strMeasure4: str
    strMeasure5: str
    strMeasure6: str
    strMeasure7: str
    strMeasure8: str
    strMeasure9: str
    strSource: str
    strTags: str
    strYoutube: str


class LatestListMatch(TypedDict, total=False):
    dateModified: str
    idMeal: str
    strArea: str
    strCategory: str
    strCreativeCommonsConfirmed: str
    strDrinkAlternate: str
    strImageSource: str
    strIngredient1: str
    strIngredient10: str
    strIngredient11: str
    strIngredient12: str
    strIngredient13: str
    strIngredient14: str
    strIngredient15: str
    strIngredient16: str
    strIngredient17: str
    strIngredient18: str
    strIngredient19: str
    strIngredient2: str
    strIngredient20: str
    strIngredient3: str
    strIngredient4: str
    strIngredient5: str
    strIngredient6: str
    strIngredient7: str
    strIngredient8: str
    strIngredient9: str
    strInstructions: str
    strMeal: str
    strMealThumb: str
    strMeasure1: str
    strMeasure10: str
    strMeasure11: str
    strMeasure12: str
    strMeasure13: str
    strMeasure14: str
    strMeasure15: str
    strMeasure16: str
    strMeasure17: str
    strMeasure18: str
    strMeasure19: str
    strMeasure2: str
    strMeasure20: str
    strMeasure3: str
    strMeasure4: str
    strMeasure5: str
    strMeasure6: str
    strMeasure7: str
    strMeasure8: str
    strMeasure9: str
    strSource: str
    strTags: str
    strYoutube: str


class List(TypedDict, total=False):
    strArea: str
    strCategory: str
    strIngredient: str


class ListListMatch(TypedDict, total=False):
    strArea: str
    strCategory: str
    strIngredient: str


class Lookup(TypedDict, total=False):
    dateModified: str
    idMeal: str
    strArea: str
    strCategory: str
    strCreativeCommonsConfirmed: str
    strDrinkAlternate: str
    strImageSource: str
    strIngredient1: str
    strIngredient10: str
    strIngredient11: str
    strIngredient12: str
    strIngredient13: str
    strIngredient14: str
    strIngredient15: str
    strIngredient16: str
    strIngredient17: str
    strIngredient18: str
    strIngredient19: str
    strIngredient2: str
    strIngredient20: str
    strIngredient3: str
    strIngredient4: str
    strIngredient5: str
    strIngredient6: str
    strIngredient7: str
    strIngredient8: str
    strIngredient9: str
    strInstructions: str
    strMeal: str
    strMealThumb: str
    strMeasure1: str
    strMeasure10: str
    strMeasure11: str
    strMeasure12: str
    strMeasure13: str
    strMeasure14: str
    strMeasure15: str
    strMeasure16: str
    strMeasure17: str
    strMeasure18: str
    strMeasure19: str
    strMeasure2: str
    strMeasure20: str
    strMeasure3: str
    strMeasure4: str
    strMeasure5: str
    strMeasure6: str
    strMeasure7: str
    strMeasure8: str
    strMeasure9: str
    strSource: str
    strTags: str
    strYoutube: str


class LookupListMatch(TypedDict, total=False):
    dateModified: str
    idMeal: str
    strArea: str
    strCategory: str
    strCreativeCommonsConfirmed: str
    strDrinkAlternate: str
    strImageSource: str
    strIngredient1: str
    strIngredient10: str
    strIngredient11: str
    strIngredient12: str
    strIngredient13: str
    strIngredient14: str
    strIngredient15: str
    strIngredient16: str
    strIngredient17: str
    strIngredient18: str
    strIngredient19: str
    strIngredient2: str
    strIngredient20: str
    strIngredient3: str
    strIngredient4: str
    strIngredient5: str
    strIngredient6: str
    strIngredient7: str
    strIngredient8: str
    strIngredient9: str
    strInstructions: str
    strMeal: str
    strMealThumb: str
    strMeasure1: str
    strMeasure10: str
    strMeasure11: str
    strMeasure12: str
    strMeasure13: str
    strMeasure14: str
    strMeasure15: str
    strMeasure16: str
    strMeasure17: str
    strMeasure18: str
    strMeasure19: str
    strMeasure2: str
    strMeasure20: str
    strMeasure3: str
    strMeasure4: str
    strMeasure5: str
    strMeasure6: str
    strMeasure7: str
    strMeasure8: str
    strMeasure9: str
    strSource: str
    strTags: str
    strYoutube: str


class Random(TypedDict, total=False):
    dateModified: str
    idMeal: str
    strArea: str
    strCategory: str
    strCreativeCommonsConfirmed: str
    strDrinkAlternate: str
    strImageSource: str
    strIngredient1: str
    strIngredient10: str
    strIngredient11: str
    strIngredient12: str
    strIngredient13: str
    strIngredient14: str
    strIngredient15: str
    strIngredient16: str
    strIngredient17: str
    strIngredient18: str
    strIngredient19: str
    strIngredient2: str
    strIngredient20: str
    strIngredient3: str
    strIngredient4: str
    strIngredient5: str
    strIngredient6: str
    strIngredient7: str
    strIngredient8: str
    strIngredient9: str
    strInstructions: str
    strMeal: str
    strMealThumb: str
    strMeasure1: str
    strMeasure10: str
    strMeasure11: str
    strMeasure12: str
    strMeasure13: str
    strMeasure14: str
    strMeasure15: str
    strMeasure16: str
    strMeasure17: str
    strMeasure18: str
    strMeasure19: str
    strMeasure2: str
    strMeasure20: str
    strMeasure3: str
    strMeasure4: str
    strMeasure5: str
    strMeasure6: str
    strMeasure7: str
    strMeasure8: str
    strMeasure9: str
    strSource: str
    strTags: str
    strYoutube: str


class RandomListMatch(TypedDict, total=False):
    dateModified: str
    idMeal: str
    strArea: str
    strCategory: str
    strCreativeCommonsConfirmed: str
    strDrinkAlternate: str
    strImageSource: str
    strIngredient1: str
    strIngredient10: str
    strIngredient11: str
    strIngredient12: str
    strIngredient13: str
    strIngredient14: str
    strIngredient15: str
    strIngredient16: str
    strIngredient17: str
    strIngredient18: str
    strIngredient19: str
    strIngredient2: str
    strIngredient20: str
    strIngredient3: str
    strIngredient4: str
    strIngredient5: str
    strIngredient6: str
    strIngredient7: str
    strIngredient8: str
    strIngredient9: str
    strInstructions: str
    strMeal: str
    strMealThumb: str
    strMeasure1: str
    strMeasure10: str
    strMeasure11: str
    strMeasure12: str
    strMeasure13: str
    strMeasure14: str
    strMeasure15: str
    strMeasure16: str
    strMeasure17: str
    strMeasure18: str
    strMeasure19: str
    strMeasure2: str
    strMeasure20: str
    strMeasure3: str
    strMeasure4: str
    strMeasure5: str
    strMeasure6: str
    strMeasure7: str
    strMeasure8: str
    strMeasure9: str
    strSource: str
    strTags: str
    strYoutube: str


class Randomselection(TypedDict, total=False):
    dateModified: str
    idMeal: str
    strArea: str
    strCategory: str
    strCreativeCommonsConfirmed: str
    strDrinkAlternate: str
    strImageSource: str
    strIngredient1: str
    strIngredient10: str
    strIngredient11: str
    strIngredient12: str
    strIngredient13: str
    strIngredient14: str
    strIngredient15: str
    strIngredient16: str
    strIngredient17: str
    strIngredient18: str
    strIngredient19: str
    strIngredient2: str
    strIngredient20: str
    strIngredient3: str
    strIngredient4: str
    strIngredient5: str
    strIngredient6: str
    strIngredient7: str
    strIngredient8: str
    strIngredient9: str
    strInstructions: str
    strMeal: str
    strMealThumb: str
    strMeasure1: str
    strMeasure10: str
    strMeasure11: str
    strMeasure12: str
    strMeasure13: str
    strMeasure14: str
    strMeasure15: str
    strMeasure16: str
    strMeasure17: str
    strMeasure18: str
    strMeasure19: str
    strMeasure2: str
    strMeasure20: str
    strMeasure3: str
    strMeasure4: str
    strMeasure5: str
    strMeasure6: str
    strMeasure7: str
    strMeasure8: str
    strMeasure9: str
    strSource: str
    strTags: str
    strYoutube: str


class RandomselectionListMatch(TypedDict, total=False):
    dateModified: str
    idMeal: str
    strArea: str
    strCategory: str
    strCreativeCommonsConfirmed: str
    strDrinkAlternate: str
    strImageSource: str
    strIngredient1: str
    strIngredient10: str
    strIngredient11: str
    strIngredient12: str
    strIngredient13: str
    strIngredient14: str
    strIngredient15: str
    strIngredient16: str
    strIngredient17: str
    strIngredient18: str
    strIngredient19: str
    strIngredient2: str
    strIngredient20: str
    strIngredient3: str
    strIngredient4: str
    strIngredient5: str
    strIngredient6: str
    strIngredient7: str
    strIngredient8: str
    strIngredient9: str
    strInstructions: str
    strMeal: str
    strMealThumb: str
    strMeasure1: str
    strMeasure10: str
    strMeasure11: str
    strMeasure12: str
    strMeasure13: str
    strMeasure14: str
    strMeasure15: str
    strMeasure16: str
    strMeasure17: str
    strMeasure18: str
    strMeasure19: str
    strMeasure2: str
    strMeasure20: str
    strMeasure3: str
    strMeasure4: str
    strMeasure5: str
    strMeasure6: str
    strMeasure7: str
    strMeasure8: str
    strMeasure9: str
    strSource: str
    strTags: str
    strYoutube: str


class Search(TypedDict, total=False):
    dateModified: str
    idMeal: str
    strArea: str
    strCategory: str
    strCreativeCommonsConfirmed: str
    strDrinkAlternate: str
    strImageSource: str
    strIngredient1: str
    strIngredient10: str
    strIngredient11: str
    strIngredient12: str
    strIngredient13: str
    strIngredient14: str
    strIngredient15: str
    strIngredient16: str
    strIngredient17: str
    strIngredient18: str
    strIngredient19: str
    strIngredient2: str
    strIngredient20: str
    strIngredient3: str
    strIngredient4: str
    strIngredient5: str
    strIngredient6: str
    strIngredient7: str
    strIngredient8: str
    strIngredient9: str
    strInstructions: str
    strMeal: str
    strMealThumb: str
    strMeasure1: str
    strMeasure10: str
    strMeasure11: str
    strMeasure12: str
    strMeasure13: str
    strMeasure14: str
    strMeasure15: str
    strMeasure16: str
    strMeasure17: str
    strMeasure18: str
    strMeasure19: str
    strMeasure2: str
    strMeasure20: str
    strMeasure3: str
    strMeasure4: str
    strMeasure5: str
    strMeasure6: str
    strMeasure7: str
    strMeasure8: str
    strMeasure9: str
    strSource: str
    strTags: str
    strYoutube: str


class SearchListMatch(TypedDict, total=False):
    dateModified: str
    idMeal: str
    strArea: str
    strCategory: str
    strCreativeCommonsConfirmed: str
    strDrinkAlternate: str
    strImageSource: str
    strIngredient1: str
    strIngredient10: str
    strIngredient11: str
    strIngredient12: str
    strIngredient13: str
    strIngredient14: str
    strIngredient15: str
    strIngredient16: str
    strIngredient17: str
    strIngredient18: str
    strIngredient19: str
    strIngredient2: str
    strIngredient20: str
    strIngredient3: str
    strIngredient4: str
    strIngredient5: str
    strIngredient6: str
    strIngredient7: str
    strIngredient8: str
    strIngredient9: str
    strInstructions: str
    strMeal: str
    strMealThumb: str
    strMeasure1: str
    strMeasure10: str
    strMeasure11: str
    strMeasure12: str
    strMeasure13: str
    strMeasure14: str
    strMeasure15: str
    strMeasure16: str
    strMeasure17: str
    strMeasure18: str
    strMeasure19: str
    strMeasure2: str
    strMeasure20: str
    strMeasure3: str
    strMeasure4: str
    strMeasure5: str
    strMeasure6: str
    strMeasure7: str
    strMeasure8: str
    strMeasure9: str
    strSource: str
    strTags: str
    strYoutube: str

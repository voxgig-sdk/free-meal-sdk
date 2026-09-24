package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "FreeMeal",
			"slug": "free-meal",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"ratelimit": map[string]any{
				"options": map[string]any{
					"active": false,
					"burst": 5,
					"rate": 5,
				},
				"optspec": map[string]any{
					"now": "`$FUNCTION`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"retry": map[string]any{
				"options": map[string]any{
					"active": false,
					"factor": 2,
					"maxDelay": 2000,
					"minDelay": 50,
					"retries": 2,
					"statuses": []any{
						408,
						425,
						429,
						500,
						502,
						503,
						504,
					},
				},
				"optspec": map[string]any{
					"jitter": "`$BOOLEAN`",
					"sleep": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"optspec": map[string]any{
					"entity": "`$MAP`",
					"net": "`$MAP`",
				},
				"strict": false,
				"transport": "base",
			},
			"timeout": map[string]any{
				"options": map[string]any{
					"active": false,
					"ms": 30000,
				},
				"optspec": map[string]any{
					"clearTimer": "`$FUNCTION`",
					"setTimer": "`$FUNCTION`",
				},
				"strict": false,
				"transport": "wrap",
			},
		},
		"options": map[string]any{
			"base": "https://www.themealdb.com/api/json/v1/1",
			"auth": map[string]any{
				"prefix": "",
				"in": "path",
				"name": "api_key",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"category": map[string]any{},
				"filter": map[string]any{},
				"latest": map[string]any{},
				"list": map[string]any{},
				"lookup": map[string]any{},
				"random": map[string]any{},
				"randomselection": map[string]any{},
				"search": map[string]any{},
			},
		},
		"entity": map[string]any{
			"category": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "idCategory",
						"title": "Id Category",
						"type": "`$STRING`",
						"short": "Unique category identifier",
					},
					map[string]any{
						"name": "strCategory",
						"title": "Str Category",
						"type": "`$STRING`",
						"short": "Category name",
					},
					map[string]any{
						"name": "strCategoryDescription",
						"title": "Str Category Description",
						"type": "`$STRING`",
						"short": "Category description",
					},
					map[string]any{
						"name": "strCategoryThumb",
						"title": "Str Category Thumb",
						"type": "`$STRING`",
						"short": "URL to category thumbnail image",
					},
				},
				"name": "category",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/categories.php",
								"segments": []any{
									map[string]any{
										"lit": "categories.php",
									},
								},
								"parts": []any{
									"categories.php",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.categories`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"filter": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "idMeal",
						"title": "Id Meal",
						"type": "`$STRING`",
						"short": "Unique meal identifier",
					},
					map[string]any{
						"name": "strMeal",
						"title": "Str Meal",
						"type": "`$STRING`",
						"short": "Meal name",
					},
					map[string]any{
						"name": "strMealThumb",
						"title": "Str Meal Thumb",
						"type": "`$STRING`",
						"short": "URL to meal thumbnail image",
					},
				},
				"name": "filter",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/filter.php",
								"segments": []any{
									map[string]any{
										"lit": "filter.php",
									},
								},
								"parts": []any{
									"filter.php",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "a",
											"orig": "a",
											"type": "`$STRING`",
											"kind": "query",
											"example": "Canadian",
										},
										map[string]any{
											"name": "c",
											"orig": "c",
											"type": "`$STRING`",
											"kind": "query",
											"example": "Seafood",
										},
										map[string]any{
											"name": "i",
											"orig": "i",
											"type": "`$STRING`",
											"kind": "query",
											"example": "chicken_breast",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"a",
										"c",
										"i",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"latest": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "dateModified",
						"title": "Date Modified",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "idMeal",
						"title": "Id Meal",
						"type": "`$STRING`",
						"short": "Unique meal identifier",
					},
					map[string]any{
						"name": "strArea",
						"title": "Str Area",
						"type": "`$STRING`",
						"short": "Meal area/region",
					},
					map[string]any{
						"name": "strCategory",
						"title": "Str Category",
						"type": "`$STRING`",
						"short": "Meal category",
					},
					map[string]any{
						"name": "strCreativeCommonsConfirmed",
						"title": "Str Creative Commons Confirmed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strDrinkAlternate",
						"title": "Str Drink Alternate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strImageSource",
						"title": "Str Image Source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient1",
						"title": "Str Ingredient1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient10",
						"title": "Str Ingredient10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient11",
						"title": "Str Ingredient11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient12",
						"title": "Str Ingredient12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient13",
						"title": "Str Ingredient13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient14",
						"title": "Str Ingredient14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient15",
						"title": "Str Ingredient15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient16",
						"title": "Str Ingredient16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient17",
						"title": "Str Ingredient17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient18",
						"title": "Str Ingredient18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient19",
						"title": "Str Ingredient19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient2",
						"title": "Str Ingredient2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient20",
						"title": "Str Ingredient20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient3",
						"title": "Str Ingredient3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient4",
						"title": "Str Ingredient4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient5",
						"title": "Str Ingredient5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient6",
						"title": "Str Ingredient6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient7",
						"title": "Str Ingredient7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient8",
						"title": "Str Ingredient8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient9",
						"title": "Str Ingredient9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strInstructions",
						"title": "Str Instructions",
						"type": "`$STRING`",
						"short": "Cooking instructions",
					},
					map[string]any{
						"name": "strMeal",
						"title": "Str Meal",
						"type": "`$STRING`",
						"short": "Meal name",
					},
					map[string]any{
						"name": "strMealThumb",
						"title": "Str Meal Thumb",
						"type": "`$STRING`",
						"short": "URL to meal thumbnail image",
					},
					map[string]any{
						"name": "strMeasure1",
						"title": "Str Measure1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure10",
						"title": "Str Measure10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure11",
						"title": "Str Measure11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure12",
						"title": "Str Measure12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure13",
						"title": "Str Measure13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure14",
						"title": "Str Measure14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure15",
						"title": "Str Measure15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure16",
						"title": "Str Measure16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure17",
						"title": "Str Measure17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure18",
						"title": "Str Measure18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure19",
						"title": "Str Measure19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure2",
						"title": "Str Measure2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure20",
						"title": "Str Measure20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure3",
						"title": "Str Measure3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure4",
						"title": "Str Measure4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure5",
						"title": "Str Measure5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure6",
						"title": "Str Measure6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure7",
						"title": "Str Measure7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure8",
						"title": "Str Measure8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure9",
						"title": "Str Measure9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strSource",
						"title": "Str Source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strTags",
						"title": "Str Tags",
						"type": "`$STRING`",
						"short": "Comma-separated tags",
					},
					map[string]any{
						"name": "strYoutube",
						"title": "Str Youtube",
						"type": "`$STRING`",
						"short": "YouTube video URL",
					},
				},
				"name": "latest",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/latest.php",
								"segments": []any{
									map[string]any{
										"lit": "latest.php",
									},
								},
								"parts": []any{
									"latest.php",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"list": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "strArea",
						"title": "Str Area",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
						"title": "Str Category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient",
						"title": "Str Ingredient",
						"type": "`$STRING`",
					},
				},
				"name": "list",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/list.php",
								"segments": []any{
									map[string]any{
										"lit": "list.php",
									},
								},
								"parts": []any{
									"list.php",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "a",
											"orig": "a",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "c",
											"orig": "c",
											"type": "`$STRING`",
											"kind": "query",
										},
										map[string]any{
											"name": "i",
											"orig": "i",
											"type": "`$STRING`",
											"kind": "query",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"a",
										"c",
										"i",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"lookup": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "dateModified",
						"title": "Date Modified",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "idMeal",
						"title": "Id Meal",
						"type": "`$STRING`",
						"short": "Unique meal identifier",
					},
					map[string]any{
						"name": "strArea",
						"title": "Str Area",
						"type": "`$STRING`",
						"short": "Meal area/region",
					},
					map[string]any{
						"name": "strCategory",
						"title": "Str Category",
						"type": "`$STRING`",
						"short": "Meal category",
					},
					map[string]any{
						"name": "strCreativeCommonsConfirmed",
						"title": "Str Creative Commons Confirmed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strDrinkAlternate",
						"title": "Str Drink Alternate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strImageSource",
						"title": "Str Image Source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient1",
						"title": "Str Ingredient1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient10",
						"title": "Str Ingredient10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient11",
						"title": "Str Ingredient11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient12",
						"title": "Str Ingredient12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient13",
						"title": "Str Ingredient13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient14",
						"title": "Str Ingredient14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient15",
						"title": "Str Ingredient15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient16",
						"title": "Str Ingredient16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient17",
						"title": "Str Ingredient17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient18",
						"title": "Str Ingredient18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient19",
						"title": "Str Ingredient19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient2",
						"title": "Str Ingredient2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient20",
						"title": "Str Ingredient20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient3",
						"title": "Str Ingredient3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient4",
						"title": "Str Ingredient4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient5",
						"title": "Str Ingredient5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient6",
						"title": "Str Ingredient6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient7",
						"title": "Str Ingredient7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient8",
						"title": "Str Ingredient8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient9",
						"title": "Str Ingredient9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strInstructions",
						"title": "Str Instructions",
						"type": "`$STRING`",
						"short": "Cooking instructions",
					},
					map[string]any{
						"name": "strMeal",
						"title": "Str Meal",
						"type": "`$STRING`",
						"short": "Meal name",
					},
					map[string]any{
						"name": "strMealThumb",
						"title": "Str Meal Thumb",
						"type": "`$STRING`",
						"short": "URL to meal thumbnail image",
					},
					map[string]any{
						"name": "strMeasure1",
						"title": "Str Measure1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure10",
						"title": "Str Measure10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure11",
						"title": "Str Measure11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure12",
						"title": "Str Measure12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure13",
						"title": "Str Measure13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure14",
						"title": "Str Measure14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure15",
						"title": "Str Measure15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure16",
						"title": "Str Measure16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure17",
						"title": "Str Measure17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure18",
						"title": "Str Measure18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure19",
						"title": "Str Measure19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure2",
						"title": "Str Measure2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure20",
						"title": "Str Measure20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure3",
						"title": "Str Measure3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure4",
						"title": "Str Measure4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure5",
						"title": "Str Measure5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure6",
						"title": "Str Measure6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure7",
						"title": "Str Measure7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure8",
						"title": "Str Measure8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure9",
						"title": "Str Measure9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strSource",
						"title": "Str Source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strTags",
						"title": "Str Tags",
						"type": "`$STRING`",
						"short": "Comma-separated tags",
					},
					map[string]any{
						"name": "strYoutube",
						"title": "Str Youtube",
						"type": "`$STRING`",
						"short": "YouTube video URL",
					},
				},
				"name": "lookup",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/lookup.php",
								"segments": []any{
									map[string]any{
										"lit": "lookup.php",
									},
								},
								"parts": []any{
									"lookup.php",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "i",
											"orig": "i",
											"type": "`$STRING`",
											"kind": "query",
											"reqd": true,
											"example": "52772",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"i",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"random": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "dateModified",
						"title": "Date Modified",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "idMeal",
						"title": "Id Meal",
						"type": "`$STRING`",
						"short": "Unique meal identifier",
					},
					map[string]any{
						"name": "strArea",
						"title": "Str Area",
						"type": "`$STRING`",
						"short": "Meal area/region",
					},
					map[string]any{
						"name": "strCategory",
						"title": "Str Category",
						"type": "`$STRING`",
						"short": "Meal category",
					},
					map[string]any{
						"name": "strCreativeCommonsConfirmed",
						"title": "Str Creative Commons Confirmed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strDrinkAlternate",
						"title": "Str Drink Alternate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strImageSource",
						"title": "Str Image Source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient1",
						"title": "Str Ingredient1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient10",
						"title": "Str Ingredient10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient11",
						"title": "Str Ingredient11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient12",
						"title": "Str Ingredient12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient13",
						"title": "Str Ingredient13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient14",
						"title": "Str Ingredient14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient15",
						"title": "Str Ingredient15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient16",
						"title": "Str Ingredient16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient17",
						"title": "Str Ingredient17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient18",
						"title": "Str Ingredient18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient19",
						"title": "Str Ingredient19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient2",
						"title": "Str Ingredient2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient20",
						"title": "Str Ingredient20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient3",
						"title": "Str Ingredient3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient4",
						"title": "Str Ingredient4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient5",
						"title": "Str Ingredient5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient6",
						"title": "Str Ingredient6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient7",
						"title": "Str Ingredient7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient8",
						"title": "Str Ingredient8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient9",
						"title": "Str Ingredient9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strInstructions",
						"title": "Str Instructions",
						"type": "`$STRING`",
						"short": "Cooking instructions",
					},
					map[string]any{
						"name": "strMeal",
						"title": "Str Meal",
						"type": "`$STRING`",
						"short": "Meal name",
					},
					map[string]any{
						"name": "strMealThumb",
						"title": "Str Meal Thumb",
						"type": "`$STRING`",
						"short": "URL to meal thumbnail image",
					},
					map[string]any{
						"name": "strMeasure1",
						"title": "Str Measure1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure10",
						"title": "Str Measure10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure11",
						"title": "Str Measure11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure12",
						"title": "Str Measure12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure13",
						"title": "Str Measure13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure14",
						"title": "Str Measure14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure15",
						"title": "Str Measure15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure16",
						"title": "Str Measure16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure17",
						"title": "Str Measure17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure18",
						"title": "Str Measure18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure19",
						"title": "Str Measure19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure2",
						"title": "Str Measure2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure20",
						"title": "Str Measure20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure3",
						"title": "Str Measure3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure4",
						"title": "Str Measure4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure5",
						"title": "Str Measure5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure6",
						"title": "Str Measure6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure7",
						"title": "Str Measure7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure8",
						"title": "Str Measure8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure9",
						"title": "Str Measure9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strSource",
						"title": "Str Source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strTags",
						"title": "Str Tags",
						"type": "`$STRING`",
						"short": "Comma-separated tags",
					},
					map[string]any{
						"name": "strYoutube",
						"title": "Str Youtube",
						"type": "`$STRING`",
						"short": "YouTube video URL",
					},
				},
				"name": "random",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/random.php",
								"segments": []any{
									map[string]any{
										"lit": "random.php",
									},
								},
								"parts": []any{
									"random.php",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"randomselection": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "dateModified",
						"title": "Date Modified",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "idMeal",
						"title": "Id Meal",
						"type": "`$STRING`",
						"short": "Unique meal identifier",
					},
					map[string]any{
						"name": "strArea",
						"title": "Str Area",
						"type": "`$STRING`",
						"short": "Meal area/region",
					},
					map[string]any{
						"name": "strCategory",
						"title": "Str Category",
						"type": "`$STRING`",
						"short": "Meal category",
					},
					map[string]any{
						"name": "strCreativeCommonsConfirmed",
						"title": "Str Creative Commons Confirmed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strDrinkAlternate",
						"title": "Str Drink Alternate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strImageSource",
						"title": "Str Image Source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient1",
						"title": "Str Ingredient1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient10",
						"title": "Str Ingredient10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient11",
						"title": "Str Ingredient11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient12",
						"title": "Str Ingredient12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient13",
						"title": "Str Ingredient13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient14",
						"title": "Str Ingredient14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient15",
						"title": "Str Ingredient15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient16",
						"title": "Str Ingredient16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient17",
						"title": "Str Ingredient17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient18",
						"title": "Str Ingredient18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient19",
						"title": "Str Ingredient19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient2",
						"title": "Str Ingredient2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient20",
						"title": "Str Ingredient20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient3",
						"title": "Str Ingredient3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient4",
						"title": "Str Ingredient4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient5",
						"title": "Str Ingredient5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient6",
						"title": "Str Ingredient6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient7",
						"title": "Str Ingredient7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient8",
						"title": "Str Ingredient8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient9",
						"title": "Str Ingredient9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strInstructions",
						"title": "Str Instructions",
						"type": "`$STRING`",
						"short": "Cooking instructions",
					},
					map[string]any{
						"name": "strMeal",
						"title": "Str Meal",
						"type": "`$STRING`",
						"short": "Meal name",
					},
					map[string]any{
						"name": "strMealThumb",
						"title": "Str Meal Thumb",
						"type": "`$STRING`",
						"short": "URL to meal thumbnail image",
					},
					map[string]any{
						"name": "strMeasure1",
						"title": "Str Measure1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure10",
						"title": "Str Measure10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure11",
						"title": "Str Measure11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure12",
						"title": "Str Measure12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure13",
						"title": "Str Measure13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure14",
						"title": "Str Measure14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure15",
						"title": "Str Measure15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure16",
						"title": "Str Measure16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure17",
						"title": "Str Measure17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure18",
						"title": "Str Measure18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure19",
						"title": "Str Measure19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure2",
						"title": "Str Measure2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure20",
						"title": "Str Measure20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure3",
						"title": "Str Measure3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure4",
						"title": "Str Measure4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure5",
						"title": "Str Measure5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure6",
						"title": "Str Measure6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure7",
						"title": "Str Measure7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure8",
						"title": "Str Measure8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure9",
						"title": "Str Measure9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strSource",
						"title": "Str Source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strTags",
						"title": "Str Tags",
						"type": "`$STRING`",
						"short": "Comma-separated tags",
					},
					map[string]any{
						"name": "strYoutube",
						"title": "Str Youtube",
						"type": "`$STRING`",
						"short": "YouTube video URL",
					},
				},
				"name": "randomselection",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/randomselection.php",
								"segments": []any{
									map[string]any{
										"lit": "randomselection.php",
									},
								},
								"parts": []any{
									"randomselection.php",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"args": map[string]any{},
								"select": map[string]any{},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"search": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "dateModified",
						"title": "Date Modified",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "idMeal",
						"title": "Id Meal",
						"type": "`$STRING`",
						"short": "Unique meal identifier",
					},
					map[string]any{
						"name": "strArea",
						"title": "Str Area",
						"type": "`$STRING`",
						"short": "Meal area/region",
					},
					map[string]any{
						"name": "strCategory",
						"title": "Str Category",
						"type": "`$STRING`",
						"short": "Meal category",
					},
					map[string]any{
						"name": "strCreativeCommonsConfirmed",
						"title": "Str Creative Commons Confirmed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strDrinkAlternate",
						"title": "Str Drink Alternate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strImageSource",
						"title": "Str Image Source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient1",
						"title": "Str Ingredient1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient10",
						"title": "Str Ingredient10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient11",
						"title": "Str Ingredient11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient12",
						"title": "Str Ingredient12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient13",
						"title": "Str Ingredient13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient14",
						"title": "Str Ingredient14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient15",
						"title": "Str Ingredient15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient16",
						"title": "Str Ingredient16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient17",
						"title": "Str Ingredient17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient18",
						"title": "Str Ingredient18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient19",
						"title": "Str Ingredient19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient2",
						"title": "Str Ingredient2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient20",
						"title": "Str Ingredient20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient3",
						"title": "Str Ingredient3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient4",
						"title": "Str Ingredient4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient5",
						"title": "Str Ingredient5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient6",
						"title": "Str Ingredient6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient7",
						"title": "Str Ingredient7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient8",
						"title": "Str Ingredient8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient9",
						"title": "Str Ingredient9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strInstructions",
						"title": "Str Instructions",
						"type": "`$STRING`",
						"short": "Cooking instructions",
					},
					map[string]any{
						"name": "strMeal",
						"title": "Str Meal",
						"type": "`$STRING`",
						"short": "Meal name",
					},
					map[string]any{
						"name": "strMealThumb",
						"title": "Str Meal Thumb",
						"type": "`$STRING`",
						"short": "URL to meal thumbnail image",
					},
					map[string]any{
						"name": "strMeasure1",
						"title": "Str Measure1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure10",
						"title": "Str Measure10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure11",
						"title": "Str Measure11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure12",
						"title": "Str Measure12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure13",
						"title": "Str Measure13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure14",
						"title": "Str Measure14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure15",
						"title": "Str Measure15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure16",
						"title": "Str Measure16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure17",
						"title": "Str Measure17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure18",
						"title": "Str Measure18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure19",
						"title": "Str Measure19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure2",
						"title": "Str Measure2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure20",
						"title": "Str Measure20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure3",
						"title": "Str Measure3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure4",
						"title": "Str Measure4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure5",
						"title": "Str Measure5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure6",
						"title": "Str Measure6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure7",
						"title": "Str Measure7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure8",
						"title": "Str Measure8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure9",
						"title": "Str Measure9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strSource",
						"title": "Str Source",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strTags",
						"title": "Str Tags",
						"type": "`$STRING`",
						"short": "Comma-separated tags",
					},
					map[string]any{
						"name": "strYoutube",
						"title": "Str Youtube",
						"type": "`$STRING`",
						"short": "YouTube video URL",
					},
				},
				"name": "search",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"kind": "http",
								"method": "GET",
								"orig": "/search.php",
								"segments": []any{
									map[string]any{
										"lit": "search.php",
									},
								},
								"parts": []any{
									"search.php",
								},
								"rename": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"name": "f",
											"orig": "f",
											"type": "`$STRING`",
											"kind": "query",
											"example": "a",
										},
										map[string]any{
											"name": "s",
											"orig": "s",
											"type": "`$STRING`",
											"kind": "query",
											"example": "Arrabiata",
										},
									},
								},
								"select": map[string]any{
									"exist": []any{
										"f",
										"s",
									},
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "ratelimit":
		if NewRatelimitFeatureFunc != nil {
			return NewRatelimitFeatureFunc()
		}
	case "retry":
		if NewRetryFeatureFunc != nil {
			return NewRetryFeatureFunc()
		}
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	case "timeout":
		if NewTimeoutFeatureFunc != nil {
			return NewTimeoutFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}

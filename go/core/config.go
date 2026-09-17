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
						"short": "Unique category identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
						"short": "Category name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategoryDescription",
						"short": "Category description",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategoryThumb",
						"short": "URL to category thumbnail image",
						"type": "`$STRING`",
					},
				},
				"name": "category",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/categories.php",
								"segments": []any{
									map[string]any{
										"lit": "categories.php",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.categories`",
								},
								"parts": []any{
									"categories.php",
								},
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
						"short": "Unique meal identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"short": "Meal name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
						"short": "URL to meal thumbnail image",
						"type": "`$STRING`",
					},
				},
				"name": "filter",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "Canadian",
											"kind": "query",
											"name": "a",
											"orig": "a",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "Seafood",
											"kind": "query",
											"name": "c",
											"orig": "c",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "chicken_breast",
											"kind": "query",
											"name": "i",
											"orig": "i",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/filter.php",
								"segments": []any{
									map[string]any{
										"lit": "filter.php",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"a",
										"c",
										"i",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"parts": []any{
									"filter.php",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "idMeal",
						"short": "Unique meal identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strArea",
						"short": "Meal area/region",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
						"short": "Meal category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCreativeCommonsConfirmed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strDrinkAlternate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strImageSource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strInstructions",
						"short": "Cooking instructions",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"short": "Meal name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
						"short": "URL to meal thumbnail image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strSource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strTags",
						"short": "Comma-separated tags",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strYoutube",
						"short": "YouTube video URL",
						"type": "`$STRING`",
					},
				},
				"name": "latest",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/latest.php",
								"segments": []any{
									map[string]any{
										"lit": "latest.php",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"parts": []any{
									"latest.php",
								},
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient",
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
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"kind": "query",
											"name": "a",
											"orig": "a",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "c",
											"orig": "c",
											"type": "`$STRING`",
										},
										map[string]any{
											"kind": "query",
											"name": "i",
											"orig": "i",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/list.php",
								"segments": []any{
									map[string]any{
										"lit": "list.php",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"a",
										"c",
										"i",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"parts": []any{
									"list.php",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "idMeal",
						"short": "Unique meal identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strArea",
						"short": "Meal area/region",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
						"short": "Meal category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCreativeCommonsConfirmed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strDrinkAlternate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strImageSource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strInstructions",
						"short": "Cooking instructions",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"short": "Meal name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
						"short": "URL to meal thumbnail image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strSource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strTags",
						"short": "Comma-separated tags",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strYoutube",
						"short": "YouTube video URL",
						"type": "`$STRING`",
					},
				},
				"name": "lookup",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "52772",
											"kind": "query",
											"name": "i",
											"orig": "i",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/lookup.php",
								"segments": []any{
									map[string]any{
										"lit": "lookup.php",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"i",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"parts": []any{
									"lookup.php",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "idMeal",
						"short": "Unique meal identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strArea",
						"short": "Meal area/region",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
						"short": "Meal category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCreativeCommonsConfirmed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strDrinkAlternate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strImageSource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strInstructions",
						"short": "Cooking instructions",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"short": "Meal name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
						"short": "URL to meal thumbnail image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strSource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strTags",
						"short": "Comma-separated tags",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strYoutube",
						"short": "YouTube video URL",
						"type": "`$STRING`",
					},
				},
				"name": "random",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/random.php",
								"segments": []any{
									map[string]any{
										"lit": "random.php",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"parts": []any{
									"random.php",
								},
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "idMeal",
						"short": "Unique meal identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strArea",
						"short": "Meal area/region",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
						"short": "Meal category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCreativeCommonsConfirmed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strDrinkAlternate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strImageSource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strInstructions",
						"short": "Cooking instructions",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"short": "Meal name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
						"short": "URL to meal thumbnail image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strSource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strTags",
						"short": "Comma-separated tags",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strYoutube",
						"short": "YouTube video URL",
						"type": "`$STRING`",
					},
				},
				"name": "randomselection",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/randomselection.php",
								"segments": []any{
									map[string]any{
										"lit": "randomselection.php",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"parts": []any{
									"randomselection.php",
								},
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "idMeal",
						"short": "Unique meal identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strArea",
						"short": "Meal area/region",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
						"short": "Meal category",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCreativeCommonsConfirmed",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strDrinkAlternate",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strImageSource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strIngredient9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strInstructions",
						"short": "Cooking instructions",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"short": "Meal name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
						"short": "URL to meal thumbnail image",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure1",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure10",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure11",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure12",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure13",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure14",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure15",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure16",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure17",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure18",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure19",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure2",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure20",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure3",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure4",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure5",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure6",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure7",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure8",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeasure9",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strSource",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strTags",
						"short": "Comma-separated tags",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strYoutube",
						"short": "YouTube video URL",
						"type": "`$STRING`",
					},
				},
				"name": "search",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "a",
											"kind": "query",
											"name": "f",
											"orig": "f",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "Arrabiata",
											"kind": "query",
											"name": "s",
											"orig": "s",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/search.php",
								"segments": []any{
									map[string]any{
										"lit": "search.php",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"f",
										"s",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
								},
								"parts": []any{
									"search.php",
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

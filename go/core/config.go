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
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://www.themealdb.com/api/json/v1/1",
			"auth": map[string]any{
				"prefix": "",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategoryDescription",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategoryThumb",
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
								"parts": []any{
									"categories.php",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.categories`",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
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
								"parts": []any{
									"filter.php",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strArea",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strYoutube",
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
								"parts": []any{
									"latest.php",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
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
								"parts": []any{
									"list.php",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strArea",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strYoutube",
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
								"parts": []any{
									"lookup.php",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strArea",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strYoutube",
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
								"parts": []any{
									"random.php",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strArea",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strYoutube",
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
								"parts": []any{
									"randomselection.php",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.meals`",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strArea",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strCategory",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMeal",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strMealThumb",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "strYoutube",
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
								"parts": []any{
									"search.php",
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
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}

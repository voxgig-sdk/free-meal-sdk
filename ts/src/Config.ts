
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'FreeMeal',
        slug: "free-meal",
    version: "0.0.1",
    target: "ts",

  }


  feature = {
     test:     {
      "options": {
        "active": false
      }
    },

  }


  options = {
    base: "https://www.themealdb.com/api/json/v1/1",

    auth: {
      prefix: '',
    },

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      category: {
      },

      filter: {
      },

      latest: {
      },

      list: {
      },

      lookup: {
      },

      random: {
      },

      randomselection: {
      },

      search: {
      },

    }
  }


  entity = {
    "category": {
      "fields": [
        {
          "name": "idCategory",
          "short": "Unique category identifier",
          "type": "`$STRING`"
        },
        {
          "name": "strCategory",
          "short": "Category name",
          "type": "`$STRING`"
        },
        {
          "name": "strCategoryDescription",
          "short": "Category description",
          "type": "`$STRING`"
        },
        {
          "name": "strCategoryThumb",
          "short": "URL to category thumbnail image",
          "type": "`$STRING`"
        }
      ],
      "name": "category",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/categories.php",
              "parts": [
                "categories.php"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.categories`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "filter": {
      "fields": [
        {
          "name": "idMeal",
          "short": "Unique meal identifier",
          "type": "`$STRING`"
        },
        {
          "name": "strMeal",
          "short": "Meal name",
          "type": "`$STRING`"
        },
        {
          "name": "strMealThumb",
          "short": "URL to meal thumbnail image",
          "type": "`$STRING`"
        }
      ],
      "name": "filter",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": "Canadian",
                    "kind": "query",
                    "name": "a",
                    "orig": "a",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "Seafood",
                    "kind": "query",
                    "name": "c",
                    "orig": "c",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "chicken_breast",
                    "kind": "query",
                    "name": "i",
                    "orig": "i",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/filter.php",
              "parts": [
                "filter.php"
              ],
              "select": {
                "exist": [
                  "a",
                  "c",
                  "i"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "latest": {
      "fields": [
        {
          "name": "dateModified",
          "type": "`$STRING`"
        },
        {
          "name": "idMeal",
          "short": "Unique meal identifier",
          "type": "`$STRING`"
        },
        {
          "name": "strArea",
          "short": "Meal area/region",
          "type": "`$STRING`"
        },
        {
          "name": "strCategory",
          "short": "Meal category",
          "type": "`$STRING`"
        },
        {
          "name": "strCreativeCommonsConfirmed",
          "type": "`$STRING`"
        },
        {
          "name": "strDrinkAlternate",
          "type": "`$STRING`"
        },
        {
          "name": "strImageSource",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient1",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient10",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient11",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient12",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient13",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient14",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient15",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient16",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient17",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient18",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient19",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient2",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient20",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient3",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient4",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient5",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient6",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient7",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient8",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient9",
          "type": "`$STRING`"
        },
        {
          "name": "strInstructions",
          "short": "Cooking instructions",
          "type": "`$STRING`"
        },
        {
          "name": "strMeal",
          "short": "Meal name",
          "type": "`$STRING`"
        },
        {
          "name": "strMealThumb",
          "short": "URL to meal thumbnail image",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure1",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure10",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure11",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure12",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure13",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure14",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure15",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure16",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure17",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure18",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure19",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure2",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure20",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure3",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure4",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure5",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure6",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure7",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure8",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure9",
          "type": "`$STRING`"
        },
        {
          "name": "strSource",
          "type": "`$STRING`"
        },
        {
          "name": "strTags",
          "short": "Comma-separated tags",
          "type": "`$STRING`"
        },
        {
          "name": "strYoutube",
          "short": "YouTube video URL",
          "type": "`$STRING`"
        }
      ],
      "name": "latest",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/latest.php",
              "parts": [
                "latest.php"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "list": {
      "fields": [
        {
          "name": "strArea",
          "type": "`$STRING`"
        },
        {
          "name": "strCategory",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient",
          "type": "`$STRING`"
        }
      ],
      "name": "list",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "kind": "query",
                    "name": "a",
                    "orig": "a",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "c",
                    "orig": "c",
                    "type": "`$STRING`"
                  },
                  {
                    "kind": "query",
                    "name": "i",
                    "orig": "i",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/list.php",
              "parts": [
                "list.php"
              ],
              "select": {
                "exist": [
                  "a",
                  "c",
                  "i"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "lookup": {
      "fields": [
        {
          "name": "dateModified",
          "type": "`$STRING`"
        },
        {
          "name": "idMeal",
          "short": "Unique meal identifier",
          "type": "`$STRING`"
        },
        {
          "name": "strArea",
          "short": "Meal area/region",
          "type": "`$STRING`"
        },
        {
          "name": "strCategory",
          "short": "Meal category",
          "type": "`$STRING`"
        },
        {
          "name": "strCreativeCommonsConfirmed",
          "type": "`$STRING`"
        },
        {
          "name": "strDrinkAlternate",
          "type": "`$STRING`"
        },
        {
          "name": "strImageSource",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient1",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient10",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient11",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient12",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient13",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient14",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient15",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient16",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient17",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient18",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient19",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient2",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient20",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient3",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient4",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient5",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient6",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient7",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient8",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient9",
          "type": "`$STRING`"
        },
        {
          "name": "strInstructions",
          "short": "Cooking instructions",
          "type": "`$STRING`"
        },
        {
          "name": "strMeal",
          "short": "Meal name",
          "type": "`$STRING`"
        },
        {
          "name": "strMealThumb",
          "short": "URL to meal thumbnail image",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure1",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure10",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure11",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure12",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure13",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure14",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure15",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure16",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure17",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure18",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure19",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure2",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure20",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure3",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure4",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure5",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure6",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure7",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure8",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure9",
          "type": "`$STRING`"
        },
        {
          "name": "strSource",
          "type": "`$STRING`"
        },
        {
          "name": "strTags",
          "short": "Comma-separated tags",
          "type": "`$STRING`"
        },
        {
          "name": "strYoutube",
          "short": "YouTube video URL",
          "type": "`$STRING`"
        }
      ],
      "name": "lookup",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": "52772",
                    "kind": "query",
                    "name": "i",
                    "orig": "i",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/lookup.php",
              "parts": [
                "lookup.php"
              ],
              "select": {
                "exist": [
                  "i"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "random": {
      "fields": [
        {
          "name": "dateModified",
          "type": "`$STRING`"
        },
        {
          "name": "idMeal",
          "short": "Unique meal identifier",
          "type": "`$STRING`"
        },
        {
          "name": "strArea",
          "short": "Meal area/region",
          "type": "`$STRING`"
        },
        {
          "name": "strCategory",
          "short": "Meal category",
          "type": "`$STRING`"
        },
        {
          "name": "strCreativeCommonsConfirmed",
          "type": "`$STRING`"
        },
        {
          "name": "strDrinkAlternate",
          "type": "`$STRING`"
        },
        {
          "name": "strImageSource",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient1",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient10",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient11",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient12",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient13",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient14",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient15",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient16",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient17",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient18",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient19",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient2",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient20",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient3",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient4",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient5",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient6",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient7",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient8",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient9",
          "type": "`$STRING`"
        },
        {
          "name": "strInstructions",
          "short": "Cooking instructions",
          "type": "`$STRING`"
        },
        {
          "name": "strMeal",
          "short": "Meal name",
          "type": "`$STRING`"
        },
        {
          "name": "strMealThumb",
          "short": "URL to meal thumbnail image",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure1",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure10",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure11",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure12",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure13",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure14",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure15",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure16",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure17",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure18",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure19",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure2",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure20",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure3",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure4",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure5",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure6",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure7",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure8",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure9",
          "type": "`$STRING`"
        },
        {
          "name": "strSource",
          "type": "`$STRING`"
        },
        {
          "name": "strTags",
          "short": "Comma-separated tags",
          "type": "`$STRING`"
        },
        {
          "name": "strYoutube",
          "short": "YouTube video URL",
          "type": "`$STRING`"
        }
      ],
      "name": "random",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/random.php",
              "parts": [
                "random.php"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "randomselection": {
      "fields": [
        {
          "name": "dateModified",
          "type": "`$STRING`"
        },
        {
          "name": "idMeal",
          "short": "Unique meal identifier",
          "type": "`$STRING`"
        },
        {
          "name": "strArea",
          "short": "Meal area/region",
          "type": "`$STRING`"
        },
        {
          "name": "strCategory",
          "short": "Meal category",
          "type": "`$STRING`"
        },
        {
          "name": "strCreativeCommonsConfirmed",
          "type": "`$STRING`"
        },
        {
          "name": "strDrinkAlternate",
          "type": "`$STRING`"
        },
        {
          "name": "strImageSource",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient1",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient10",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient11",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient12",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient13",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient14",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient15",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient16",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient17",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient18",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient19",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient2",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient20",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient3",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient4",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient5",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient6",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient7",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient8",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient9",
          "type": "`$STRING`"
        },
        {
          "name": "strInstructions",
          "short": "Cooking instructions",
          "type": "`$STRING`"
        },
        {
          "name": "strMeal",
          "short": "Meal name",
          "type": "`$STRING`"
        },
        {
          "name": "strMealThumb",
          "short": "URL to meal thumbnail image",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure1",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure10",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure11",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure12",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure13",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure14",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure15",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure16",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure17",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure18",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure19",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure2",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure20",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure3",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure4",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure5",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure6",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure7",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure8",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure9",
          "type": "`$STRING`"
        },
        {
          "name": "strSource",
          "type": "`$STRING`"
        },
        {
          "name": "strTags",
          "short": "Comma-separated tags",
          "type": "`$STRING`"
        },
        {
          "name": "strYoutube",
          "short": "YouTube video URL",
          "type": "`$STRING`"
        }
      ],
      "name": "randomselection",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/randomselection.php",
              "parts": [
                "randomselection.php"
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    },
    "search": {
      "fields": [
        {
          "name": "dateModified",
          "type": "`$STRING`"
        },
        {
          "name": "idMeal",
          "short": "Unique meal identifier",
          "type": "`$STRING`"
        },
        {
          "name": "strArea",
          "short": "Meal area/region",
          "type": "`$STRING`"
        },
        {
          "name": "strCategory",
          "short": "Meal category",
          "type": "`$STRING`"
        },
        {
          "name": "strCreativeCommonsConfirmed",
          "type": "`$STRING`"
        },
        {
          "name": "strDrinkAlternate",
          "type": "`$STRING`"
        },
        {
          "name": "strImageSource",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient1",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient10",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient11",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient12",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient13",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient14",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient15",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient16",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient17",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient18",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient19",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient2",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient20",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient3",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient4",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient5",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient6",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient7",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient8",
          "type": "`$STRING`"
        },
        {
          "name": "strIngredient9",
          "type": "`$STRING`"
        },
        {
          "name": "strInstructions",
          "short": "Cooking instructions",
          "type": "`$STRING`"
        },
        {
          "name": "strMeal",
          "short": "Meal name",
          "type": "`$STRING`"
        },
        {
          "name": "strMealThumb",
          "short": "URL to meal thumbnail image",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure1",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure10",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure11",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure12",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure13",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure14",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure15",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure16",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure17",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure18",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure19",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure2",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure20",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure3",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure4",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure5",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure6",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure7",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure8",
          "type": "`$STRING`"
        },
        {
          "name": "strMeasure9",
          "type": "`$STRING`"
        },
        {
          "name": "strSource",
          "type": "`$STRING`"
        },
        {
          "name": "strTags",
          "short": "Comma-separated tags",
          "type": "`$STRING`"
        },
        {
          "name": "strYoutube",
          "short": "YouTube video URL",
          "type": "`$STRING`"
        }
      ],
      "name": "search",
      "op": {
        "list": {
          "input": "data",
          "name": "list",
          "points": [
            {
              "args": {
                "query": [
                  {
                    "example": "a",
                    "kind": "query",
                    "name": "f",
                    "orig": "f",
                    "type": "`$STRING`"
                  },
                  {
                    "example": "Arrabiata",
                    "kind": "query",
                    "name": "s",
                    "orig": "s",
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/search.php",
              "parts": [
                "search.php"
              ],
              "select": {
                "exist": [
                  "f",
                  "s"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              }
            }
          ]
        }
      },
      "relations": {
        "ancestors": []
      }
    }
  }
}


const config = new Config()

export {
  config
}


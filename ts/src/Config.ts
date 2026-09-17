
import { BaseFeature } from './feature/base/BaseFeature'
import { RatelimitFeature } from './feature/ratelimit/RatelimitFeature'
import { RetryFeature } from './feature/retry/RetryFeature'
import { TestFeature } from './feature/test/TestFeature'
import { TimeoutFeature } from './feature/timeout/TimeoutFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   ratelimit: RatelimitFeature,
 retry: RetryFeature,
 test: TestFeature,
 timeout: TimeoutFeature,

}


// Per-feature plugin DEFINITIONS (voxgig/plugin `Definition` values), from
// the model's active plugin groups. A feature that takes a `plugins` option
// (secrets over sekreto) reads its own entry; a feature with no plugins has
// none. Named imports above make each definition statically reachable, so
// an SDK carries exactly the plugin modules its model selects — the same
// leanness the old side-effect registry imports bought, without a registry.
const FEATURE_PLUGINS: Record<string, any[]> = {
  
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
     ratelimit:     {
      "options": {
        "active": false,
        "burst": 5,
        "rate": 5
      },
      "optspec": {
        "now": "`$FUNCTION`",
        "sleep": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },
 retry:     {
      "options": {
        "active": false,
        "factor": 2,
        "maxDelay": 2000,
        "minDelay": 50,
        "retries": 2,
        "statuses": [
          408,
          425,
          429,
          500,
          502,
          503,
          504
        ]
      },
      "optspec": {
        "jitter": "`$BOOLEAN`",
        "sleep": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },
 test:     {
      "options": {
        "active": false
      },
      "optspec": {
        "entity": "`$MAP`",
        "net": "`$MAP`"
      },
      "strict": false,
      "transport": "base"
    },
 timeout:     {
      "options": {
        "active": false,
        "ms": 30000
      },
      "optspec": {
        "clearTimer": "`$FUNCTION`",
        "setTimer": "`$FUNCTION`"
      },
      "strict": false,
      "transport": "wrap"
    },

  }


  options = {
    base: "https://www.themealdb.com/api/json/v1/1",

    auth: {
      prefix: '',
      in: 'path',
      name: 'api_key',
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
              "segments": [
                {
                  "lit": "categories.php"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.categories`"
              },
              "parts": [
                "categories.php"
              ]
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
              "segments": [
                {
                  "lit": "filter.php"
                }
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
              },
              "parts": [
                "filter.php"
              ]
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
              "segments": [
                {
                  "lit": "latest.php"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              },
              "parts": [
                "latest.php"
              ]
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
              "segments": [
                {
                  "lit": "list.php"
                }
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
              },
              "parts": [
                "list.php"
              ]
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
              "segments": [
                {
                  "lit": "lookup.php"
                }
              ],
              "select": {
                "exist": [
                  "i"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              },
              "parts": [
                "lookup.php"
              ]
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
              "segments": [
                {
                  "lit": "random.php"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              },
              "parts": [
                "random.php"
              ]
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
              "segments": [
                {
                  "lit": "randomselection.php"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body.meals`"
              },
              "parts": [
                "randomselection.php"
              ]
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
              "segments": [
                {
                  "lit": "search.php"
                }
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
              },
              "parts": [
                "search.php"
              ]
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
  config,
  FEATURE_PLUGINS,
}


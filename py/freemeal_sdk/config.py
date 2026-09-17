# FreeMeal SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "FreeMeal",
            "slug": "free-meal",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "ratelimit": {
        "options": {
          "active": False,
          "burst": 5,
          "rate": 5,
        },
        "optspec": {
          "now": "`$FUNCTION`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "retry": {
        "options": {
          "active": False,
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
            504,
          ],
        },
        "optspec": {
          "jitter": "`$BOOLEAN`",
          "sleep": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
            "test": {
        "options": {
          "active": False,
        },
        "optspec": {
          "entity": "`$MAP`",
          "net": "`$MAP`",
        },
        "strict": False,
        "transport": "base",
      },
            "timeout": {
        "options": {
          "active": False,
          "ms": 30000,
        },
        "optspec": {
          "clearTimer": "`$FUNCTION`",
          "setTimer": "`$FUNCTION`",
        },
        "strict": False,
        "transport": "wrap",
      },
        },
        "options": {
            "base": "https://www.themealdb.com/api/json/v1/1",
            "auth": {
                "prefix": "",
                "in": "path",
                "name": "api_key",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "category": {},
                "filter": {},
                "latest": {},
                "list": {},
                "lookup": {},
                "random": {},
                "randomselection": {},
                "search": {},
            },
        },
        "entity": {
      "category": {
        "fields": [
          {
            "name": "idCategory",
            "short": "Unique category identifier",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
            "short": "Category name",
            "type": "`$STRING`",
          },
          {
            "name": "strCategoryDescription",
            "short": "Category description",
            "type": "`$STRING`",
          },
          {
            "name": "strCategoryThumb",
            "short": "URL to category thumbnail image",
            "type": "`$STRING`",
          },
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
                    "lit": "categories.php",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.categories`",
                },
                "parts": [
                  "categories.php",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "filter": {
        "fields": [
          {
            "name": "idMeal",
            "short": "Unique meal identifier",
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "short": "Meal name",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
            "short": "URL to meal thumbnail image",
            "type": "`$STRING`",
          },
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
                      "type": "`$STRING`",
                    },
                    {
                      "example": "Seafood",
                      "kind": "query",
                      "name": "c",
                      "orig": "c",
                      "type": "`$STRING`",
                    },
                    {
                      "example": "chicken_breast",
                      "kind": "query",
                      "name": "i",
                      "orig": "i",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/filter.php",
                "segments": [
                  {
                    "lit": "filter.php",
                  },
                ],
                "select": {
                  "exist": [
                    "a",
                    "c",
                    "i",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meals`",
                },
                "parts": [
                  "filter.php",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "latest": {
        "fields": [
          {
            "name": "dateModified",
            "type": "`$STRING`",
          },
          {
            "name": "idMeal",
            "short": "Unique meal identifier",
            "type": "`$STRING`",
          },
          {
            "name": "strArea",
            "short": "Meal area/region",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
            "short": "Meal category",
            "type": "`$STRING`",
          },
          {
            "name": "strCreativeCommonsConfirmed",
            "type": "`$STRING`",
          },
          {
            "name": "strDrinkAlternate",
            "type": "`$STRING`",
          },
          {
            "name": "strImageSource",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient1",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient10",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient11",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient12",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient13",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient14",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient15",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient16",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient17",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient18",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient19",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient2",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient20",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient3",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient4",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient5",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient6",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient7",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient8",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient9",
            "type": "`$STRING`",
          },
          {
            "name": "strInstructions",
            "short": "Cooking instructions",
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "short": "Meal name",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
            "short": "URL to meal thumbnail image",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure1",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure10",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure11",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure12",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure13",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure14",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure15",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure16",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure17",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure18",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure19",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure2",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure20",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure3",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure4",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure5",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure6",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure7",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure8",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure9",
            "type": "`$STRING`",
          },
          {
            "name": "strSource",
            "type": "`$STRING`",
          },
          {
            "name": "strTags",
            "short": "Comma-separated tags",
            "type": "`$STRING`",
          },
          {
            "name": "strYoutube",
            "short": "YouTube video URL",
            "type": "`$STRING`",
          },
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
                    "lit": "latest.php",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meals`",
                },
                "parts": [
                  "latest.php",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "list": {
        "fields": [
          {
            "name": "strArea",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient",
            "type": "`$STRING`",
          },
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
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "c",
                      "orig": "c",
                      "type": "`$STRING`",
                    },
                    {
                      "kind": "query",
                      "name": "i",
                      "orig": "i",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/list.php",
                "segments": [
                  {
                    "lit": "list.php",
                  },
                ],
                "select": {
                  "exist": [
                    "a",
                    "c",
                    "i",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meals`",
                },
                "parts": [
                  "list.php",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "lookup": {
        "fields": [
          {
            "name": "dateModified",
            "type": "`$STRING`",
          },
          {
            "name": "idMeal",
            "short": "Unique meal identifier",
            "type": "`$STRING`",
          },
          {
            "name": "strArea",
            "short": "Meal area/region",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
            "short": "Meal category",
            "type": "`$STRING`",
          },
          {
            "name": "strCreativeCommonsConfirmed",
            "type": "`$STRING`",
          },
          {
            "name": "strDrinkAlternate",
            "type": "`$STRING`",
          },
          {
            "name": "strImageSource",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient1",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient10",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient11",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient12",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient13",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient14",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient15",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient16",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient17",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient18",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient19",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient2",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient20",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient3",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient4",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient5",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient6",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient7",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient8",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient9",
            "type": "`$STRING`",
          },
          {
            "name": "strInstructions",
            "short": "Cooking instructions",
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "short": "Meal name",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
            "short": "URL to meal thumbnail image",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure1",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure10",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure11",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure12",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure13",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure14",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure15",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure16",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure17",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure18",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure19",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure2",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure20",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure3",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure4",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure5",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure6",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure7",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure8",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure9",
            "type": "`$STRING`",
          },
          {
            "name": "strSource",
            "type": "`$STRING`",
          },
          {
            "name": "strTags",
            "short": "Comma-separated tags",
            "type": "`$STRING`",
          },
          {
            "name": "strYoutube",
            "short": "YouTube video URL",
            "type": "`$STRING`",
          },
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
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/lookup.php",
                "segments": [
                  {
                    "lit": "lookup.php",
                  },
                ],
                "select": {
                  "exist": [
                    "i",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meals`",
                },
                "parts": [
                  "lookup.php",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "random": {
        "fields": [
          {
            "name": "dateModified",
            "type": "`$STRING`",
          },
          {
            "name": "idMeal",
            "short": "Unique meal identifier",
            "type": "`$STRING`",
          },
          {
            "name": "strArea",
            "short": "Meal area/region",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
            "short": "Meal category",
            "type": "`$STRING`",
          },
          {
            "name": "strCreativeCommonsConfirmed",
            "type": "`$STRING`",
          },
          {
            "name": "strDrinkAlternate",
            "type": "`$STRING`",
          },
          {
            "name": "strImageSource",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient1",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient10",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient11",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient12",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient13",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient14",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient15",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient16",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient17",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient18",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient19",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient2",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient20",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient3",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient4",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient5",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient6",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient7",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient8",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient9",
            "type": "`$STRING`",
          },
          {
            "name": "strInstructions",
            "short": "Cooking instructions",
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "short": "Meal name",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
            "short": "URL to meal thumbnail image",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure1",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure10",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure11",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure12",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure13",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure14",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure15",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure16",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure17",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure18",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure19",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure2",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure20",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure3",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure4",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure5",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure6",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure7",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure8",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure9",
            "type": "`$STRING`",
          },
          {
            "name": "strSource",
            "type": "`$STRING`",
          },
          {
            "name": "strTags",
            "short": "Comma-separated tags",
            "type": "`$STRING`",
          },
          {
            "name": "strYoutube",
            "short": "YouTube video URL",
            "type": "`$STRING`",
          },
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
                    "lit": "random.php",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meals`",
                },
                "parts": [
                  "random.php",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "randomselection": {
        "fields": [
          {
            "name": "dateModified",
            "type": "`$STRING`",
          },
          {
            "name": "idMeal",
            "short": "Unique meal identifier",
            "type": "`$STRING`",
          },
          {
            "name": "strArea",
            "short": "Meal area/region",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
            "short": "Meal category",
            "type": "`$STRING`",
          },
          {
            "name": "strCreativeCommonsConfirmed",
            "type": "`$STRING`",
          },
          {
            "name": "strDrinkAlternate",
            "type": "`$STRING`",
          },
          {
            "name": "strImageSource",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient1",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient10",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient11",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient12",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient13",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient14",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient15",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient16",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient17",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient18",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient19",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient2",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient20",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient3",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient4",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient5",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient6",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient7",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient8",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient9",
            "type": "`$STRING`",
          },
          {
            "name": "strInstructions",
            "short": "Cooking instructions",
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "short": "Meal name",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
            "short": "URL to meal thumbnail image",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure1",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure10",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure11",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure12",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure13",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure14",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure15",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure16",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure17",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure18",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure19",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure2",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure20",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure3",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure4",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure5",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure6",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure7",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure8",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure9",
            "type": "`$STRING`",
          },
          {
            "name": "strSource",
            "type": "`$STRING`",
          },
          {
            "name": "strTags",
            "short": "Comma-separated tags",
            "type": "`$STRING`",
          },
          {
            "name": "strYoutube",
            "short": "YouTube video URL",
            "type": "`$STRING`",
          },
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
                    "lit": "randomselection.php",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meals`",
                },
                "parts": [
                  "randomselection.php",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "search": {
        "fields": [
          {
            "name": "dateModified",
            "type": "`$STRING`",
          },
          {
            "name": "idMeal",
            "short": "Unique meal identifier",
            "type": "`$STRING`",
          },
          {
            "name": "strArea",
            "short": "Meal area/region",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
            "short": "Meal category",
            "type": "`$STRING`",
          },
          {
            "name": "strCreativeCommonsConfirmed",
            "type": "`$STRING`",
          },
          {
            "name": "strDrinkAlternate",
            "type": "`$STRING`",
          },
          {
            "name": "strImageSource",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient1",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient10",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient11",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient12",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient13",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient14",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient15",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient16",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient17",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient18",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient19",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient2",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient20",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient3",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient4",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient5",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient6",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient7",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient8",
            "type": "`$STRING`",
          },
          {
            "name": "strIngredient9",
            "type": "`$STRING`",
          },
          {
            "name": "strInstructions",
            "short": "Cooking instructions",
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "short": "Meal name",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
            "short": "URL to meal thumbnail image",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure1",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure10",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure11",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure12",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure13",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure14",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure15",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure16",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure17",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure18",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure19",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure2",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure20",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure3",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure4",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure5",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure6",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure7",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure8",
            "type": "`$STRING`",
          },
          {
            "name": "strMeasure9",
            "type": "`$STRING`",
          },
          {
            "name": "strSource",
            "type": "`$STRING`",
          },
          {
            "name": "strTags",
            "short": "Comma-separated tags",
            "type": "`$STRING`",
          },
          {
            "name": "strYoutube",
            "short": "YouTube video URL",
            "type": "`$STRING`",
          },
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
                      "type": "`$STRING`",
                    },
                    {
                      "example": "Arrabiata",
                      "kind": "query",
                      "name": "s",
                      "orig": "s",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/search.php",
                "segments": [
                  {
                    "lit": "search.php",
                  },
                ],
                "select": {
                  "exist": [
                    "f",
                    "s",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meals`",
                },
                "parts": [
                  "search.php",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
    },
    }

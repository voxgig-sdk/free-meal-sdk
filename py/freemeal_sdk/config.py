# FreeMeal SDK configuration


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
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
      },
        },
        "options": {
            "base": "https://www.themealdb.com/api/json/v1/1",
            "auth": {
                "prefix": "",
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
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
            "type": "`$STRING`",
          },
          {
            "name": "strCategoryDescription",
            "type": "`$STRING`",
          },
          {
            "name": "strCategoryThumb",
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
                "parts": [
                  "categories.php",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.categories`",
                },
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
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
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
                "parts": [
                  "filter.php",
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
            "type": "`$STRING`",
          },
          {
            "name": "strArea",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
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
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
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
            "type": "`$STRING`",
          },
          {
            "name": "strYoutube",
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
                "parts": [
                  "latest.php",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meals`",
                },
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
                "parts": [
                  "list.php",
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
            "type": "`$STRING`",
          },
          {
            "name": "strArea",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
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
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
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
            "type": "`$STRING`",
          },
          {
            "name": "strYoutube",
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
                "parts": [
                  "lookup.php",
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
            "type": "`$STRING`",
          },
          {
            "name": "strArea",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
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
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
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
            "type": "`$STRING`",
          },
          {
            "name": "strYoutube",
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
                "parts": [
                  "random.php",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meals`",
                },
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
            "type": "`$STRING`",
          },
          {
            "name": "strArea",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
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
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
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
            "type": "`$STRING`",
          },
          {
            "name": "strYoutube",
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
                "parts": [
                  "randomselection.php",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.meals`",
                },
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
            "type": "`$STRING`",
          },
          {
            "name": "strArea",
            "type": "`$STRING`",
          },
          {
            "name": "strCategory",
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
            "type": "`$STRING`",
          },
          {
            "name": "strMeal",
            "type": "`$STRING`",
          },
          {
            "name": "strMealThumb",
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
            "type": "`$STRING`",
          },
          {
            "name": "strYoutube",
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
                "parts": [
                  "search.php",
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

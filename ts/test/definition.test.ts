import { describe, test } from 'node:test'
import { SDK } from '..'
import { runDefinitionPoint } from './definition-runner'
import { isControlSkipped } from './utility'


// Generated from the API definition, not from the model this SDK was built
// from: the route, the declared query parameters, the credential the security
// scheme names, and the definition's own response example.
const PLAN: any[] = [
  {
    "entity": "category",
    "accessor": "Category",
    "op": "list",
    "method": "GET",
    "path": "/categories.php",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [],
    "queryArgs": [],
    "auth": [],
    "status": 200,
    "sample": {
      "categories": [
        {
          "idCategory": "x",
          "strCategory": "x",
          "strCategoryThumb": "x",
          "strCategoryDescription": "x"
        }
      ]
    },
    "idField": "id"
  },
  {
    "entity": "filter",
    "accessor": "Filter",
    "op": "list",
    "method": "GET",
    "path": "/filter.php",
    "args": [],
    "select": {
      "a": "Canadian",
      "c": "Seafood",
      "i": "chicken_breast"
    },
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [
      "i",
      "c",
      "a"
    ],
    "queryArgs": [
      {
        "name": "a",
        "wire": "a"
      },
      {
        "name": "c",
        "wire": "c"
      },
      {
        "name": "i",
        "wire": "i"
      }
    ],
    "auth": [],
    "status": 200,
    "sample": {
      "meals": [
        {
          "strMeal": "x",
          "strMealThumb": "x",
          "idMeal": "x"
        }
      ]
    },
    "idField": "id"
  },
  {
    "entity": "latest",
    "accessor": "Latest",
    "op": "list",
    "method": "GET",
    "path": "/latest.php",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [],
    "queryArgs": [],
    "auth": null,
    "status": 200,
    "sample": {
      "meals": [
        {
          "idMeal": "x",
          "strMeal": "x",
          "strDrinkAlternate": "x",
          "strCategory": "x",
          "strArea": "x",
          "strInstructions": "x",
          "strMealThumb": "x",
          "strTags": "x",
          "strYoutube": "x",
          "strIngredient1": "x",
          "strIngredient2": "x",
          "strIngredient3": "x",
          "strIngredient4": "x",
          "strIngredient5": "x",
          "strIngredient6": "x",
          "strIngredient7": "x",
          "strIngredient8": "x",
          "strIngredient9": "x",
          "strIngredient10": "x",
          "strIngredient11": "x",
          "strIngredient12": "x",
          "strIngredient13": "x",
          "strIngredient14": "x",
          "strIngredient15": "x",
          "strIngredient16": "x",
          "strIngredient17": "x",
          "strIngredient18": "x",
          "strIngredient19": "x",
          "strIngredient20": "x",
          "strMeasure1": "x",
          "strMeasure2": "x",
          "strMeasure3": "x",
          "strMeasure4": "x",
          "strMeasure5": "x",
          "strMeasure6": "x",
          "strMeasure7": "x",
          "strMeasure8": "x",
          "strMeasure9": "x",
          "strMeasure10": "x",
          "strMeasure11": "x",
          "strMeasure12": "x",
          "strMeasure13": "x",
          "strMeasure14": "x",
          "strMeasure15": "x",
          "strMeasure16": "x",
          "strMeasure17": "x",
          "strMeasure18": "x",
          "strMeasure19": "x",
          "strMeasure20": "x",
          "strSource": "x",
          "strImageSource": "x",
          "strCreativeCommonsConfirmed": "x",
          "dateModified": "x"
        }
      ]
    },
    "idField": "id"
  },
  {
    "entity": "list",
    "accessor": "List",
    "op": "list",
    "method": "GET",
    "path": "/list.php",
    "args": [],
    "select": {
      "a": "v1",
      "c": "v1",
      "i": "v1"
    },
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [
      "c",
      "a",
      "i"
    ],
    "queryArgs": [
      {
        "name": "a",
        "wire": "a"
      },
      {
        "name": "c",
        "wire": "c"
      },
      {
        "name": "i",
        "wire": "i"
      }
    ],
    "auth": [],
    "status": 200,
    "sample": {
      "meals": [
        {
          "strCategory": "x",
          "strArea": "x",
          "strIngredient": "x"
        }
      ]
    },
    "idField": "id"
  },
  {
    "entity": "lookup",
    "accessor": "Lookup",
    "op": "list",
    "method": "GET",
    "path": "/lookup.php",
    "args": [],
    "select": {
      "i": "52772"
    },
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [
      "i"
    ],
    "queryArgs": [
      {
        "name": "i",
        "wire": "i"
      }
    ],
    "auth": [],
    "status": 200,
    "sample": {
      "meals": [
        {
          "idMeal": "x",
          "strMeal": "x",
          "strDrinkAlternate": "x",
          "strCategory": "x",
          "strArea": "x",
          "strInstructions": "x",
          "strMealThumb": "x",
          "strTags": "x",
          "strYoutube": "x",
          "strIngredient1": "x",
          "strIngredient2": "x",
          "strIngredient3": "x",
          "strIngredient4": "x",
          "strIngredient5": "x",
          "strIngredient6": "x",
          "strIngredient7": "x",
          "strIngredient8": "x",
          "strIngredient9": "x",
          "strIngredient10": "x",
          "strIngredient11": "x",
          "strIngredient12": "x",
          "strIngredient13": "x",
          "strIngredient14": "x",
          "strIngredient15": "x",
          "strIngredient16": "x",
          "strIngredient17": "x",
          "strIngredient18": "x",
          "strIngredient19": "x",
          "strIngredient20": "x",
          "strMeasure1": "x",
          "strMeasure2": "x",
          "strMeasure3": "x",
          "strMeasure4": "x",
          "strMeasure5": "x",
          "strMeasure6": "x",
          "strMeasure7": "x",
          "strMeasure8": "x",
          "strMeasure9": "x",
          "strMeasure10": "x",
          "strMeasure11": "x",
          "strMeasure12": "x",
          "strMeasure13": "x",
          "strMeasure14": "x",
          "strMeasure15": "x",
          "strMeasure16": "x",
          "strMeasure17": "x",
          "strMeasure18": "x",
          "strMeasure19": "x",
          "strMeasure20": "x",
          "strSource": "x",
          "strImageSource": "x",
          "strCreativeCommonsConfirmed": "x",
          "dateModified": "x"
        }
      ]
    },
    "idField": "id"
  },
  {
    "entity": "random",
    "accessor": "Random",
    "op": "list",
    "method": "GET",
    "path": "/random.php",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [],
    "queryArgs": [],
    "auth": [],
    "status": 200,
    "sample": {
      "meals": [
        {
          "idMeal": "x",
          "strMeal": "x",
          "strDrinkAlternate": "x",
          "strCategory": "x",
          "strArea": "x",
          "strInstructions": "x",
          "strMealThumb": "x",
          "strTags": "x",
          "strYoutube": "x",
          "strIngredient1": "x",
          "strIngredient2": "x",
          "strIngredient3": "x",
          "strIngredient4": "x",
          "strIngredient5": "x",
          "strIngredient6": "x",
          "strIngredient7": "x",
          "strIngredient8": "x",
          "strIngredient9": "x",
          "strIngredient10": "x",
          "strIngredient11": "x",
          "strIngredient12": "x",
          "strIngredient13": "x",
          "strIngredient14": "x",
          "strIngredient15": "x",
          "strIngredient16": "x",
          "strIngredient17": "x",
          "strIngredient18": "x",
          "strIngredient19": "x",
          "strIngredient20": "x",
          "strMeasure1": "x",
          "strMeasure2": "x",
          "strMeasure3": "x",
          "strMeasure4": "x",
          "strMeasure5": "x",
          "strMeasure6": "x",
          "strMeasure7": "x",
          "strMeasure8": "x",
          "strMeasure9": "x",
          "strMeasure10": "x",
          "strMeasure11": "x",
          "strMeasure12": "x",
          "strMeasure13": "x",
          "strMeasure14": "x",
          "strMeasure15": "x",
          "strMeasure16": "x",
          "strMeasure17": "x",
          "strMeasure18": "x",
          "strMeasure19": "x",
          "strMeasure20": "x",
          "strSource": "x",
          "strImageSource": "x",
          "strCreativeCommonsConfirmed": "x",
          "dateModified": "x"
        }
      ]
    },
    "idField": "id"
  },
  {
    "entity": "randomselection",
    "accessor": "Randomselection",
    "op": "list",
    "method": "GET",
    "path": "/randomselection.php",
    "args": [],
    "select": {},
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [],
    "queryArgs": [],
    "auth": null,
    "status": 200,
    "sample": {
      "meals": [
        {
          "idMeal": "x",
          "strMeal": "x",
          "strDrinkAlternate": "x",
          "strCategory": "x",
          "strArea": "x",
          "strInstructions": "x",
          "strMealThumb": "x",
          "strTags": "x",
          "strYoutube": "x",
          "strIngredient1": "x",
          "strIngredient2": "x",
          "strIngredient3": "x",
          "strIngredient4": "x",
          "strIngredient5": "x",
          "strIngredient6": "x",
          "strIngredient7": "x",
          "strIngredient8": "x",
          "strIngredient9": "x",
          "strIngredient10": "x",
          "strIngredient11": "x",
          "strIngredient12": "x",
          "strIngredient13": "x",
          "strIngredient14": "x",
          "strIngredient15": "x",
          "strIngredient16": "x",
          "strIngredient17": "x",
          "strIngredient18": "x",
          "strIngredient19": "x",
          "strIngredient20": "x",
          "strMeasure1": "x",
          "strMeasure2": "x",
          "strMeasure3": "x",
          "strMeasure4": "x",
          "strMeasure5": "x",
          "strMeasure6": "x",
          "strMeasure7": "x",
          "strMeasure8": "x",
          "strMeasure9": "x",
          "strMeasure10": "x",
          "strMeasure11": "x",
          "strMeasure12": "x",
          "strMeasure13": "x",
          "strMeasure14": "x",
          "strMeasure15": "x",
          "strMeasure16": "x",
          "strMeasure17": "x",
          "strMeasure18": "x",
          "strMeasure19": "x",
          "strMeasure20": "x",
          "strSource": "x",
          "strImageSource": "x",
          "strCreativeCommonsConfirmed": "x",
          "dateModified": "x"
        }
      ]
    },
    "idField": "id"
  },
  {
    "entity": "search",
    "accessor": "Search",
    "op": "list",
    "method": "GET",
    "path": "/search.php",
    "args": [],
    "select": {
      "f": "a",
      "s": "Arrabiata"
    },
    "headers": [],
    "cookies": [],
    "responseMedia": [
      "application/json"
    ],
    "query": [
      "s",
      "f"
    ],
    "queryArgs": [
      {
        "name": "f",
        "wire": "f"
      },
      {
        "name": "s",
        "wire": "s"
      }
    ],
    "auth": [],
    "status": 200,
    "sample": {
      "meals": [
        {
          "idMeal": "x",
          "strMeal": "x",
          "strDrinkAlternate": "x",
          "strCategory": "x",
          "strArea": "x",
          "strInstructions": "x",
          "strMealThumb": "x",
          "strTags": "x",
          "strYoutube": "x",
          "strIngredient1": "x",
          "strIngredient2": "x",
          "strIngredient3": "x",
          "strIngredient4": "x",
          "strIngredient5": "x",
          "strIngredient6": "x",
          "strIngredient7": "x",
          "strIngredient8": "x",
          "strIngredient9": "x",
          "strIngredient10": "x",
          "strIngredient11": "x",
          "strIngredient12": "x",
          "strIngredient13": "x",
          "strIngredient14": "x",
          "strIngredient15": "x",
          "strIngredient16": "x",
          "strIngredient17": "x",
          "strIngredient18": "x",
          "strIngredient19": "x",
          "strIngredient20": "x",
          "strMeasure1": "x",
          "strMeasure2": "x",
          "strMeasure3": "x",
          "strMeasure4": "x",
          "strMeasure5": "x",
          "strMeasure6": "x",
          "strMeasure7": "x",
          "strMeasure8": "x",
          "strMeasure9": "x",
          "strMeasure10": "x",
          "strMeasure11": "x",
          "strMeasure12": "x",
          "strMeasure13": "x",
          "strMeasure14": "x",
          "strMeasure15": "x",
          "strMeasure16": "x",
          "strMeasure17": "x",
          "strMeasure18": "x",
          "strMeasure19": "x",
          "strMeasure20": "x",
          "strSource": "x",
          "strImageSource": "x",
          "strCreativeCommonsConfirmed": "x",
          "dateModified": "x"
        }
      ]
    },
    "idField": "id"
  }
]


describe('definition', () => {
  for (const point of PLAN) {
    test(point.entity + '.' + point.op + ' ' + point.method + ' ' + point.path, async (t) => {
      const control = isControlSkipped('entityOp', point.entity + '.' + point.op, 'definition')
      if (control.skip) {
        t.skip(control.reason || 'skipped via sdk-test-control.json')
        return
      }
      await runDefinitionPoint(SDK, point)
    })
  }
})

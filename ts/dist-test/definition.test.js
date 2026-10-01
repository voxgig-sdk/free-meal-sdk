"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const node_test_1 = require("node:test");
const __1 = require("..");
const definition_runner_1 = require("./definition-runner");
const utility_1 = require("./utility");
// Generated from the API definition, not from the model this SDK was built
// from: the route, the declared query parameters, the credential the security
// scheme names, and the definition's own response example.
const PLAN = [
    {
        "entity": "category",
        "accessor": "Category",
        "op": "list",
        "method": "GET",
        "path": "/categories.php",
        "args": [],
        "select": {},
        "headers": [],
        "query": [],
        "auth": [],
        "status": 200,
        "sample": {
            "categories": [
                {
                    "idCategory": "x",
                    "strCategory": "x",
                    "strCategoryDescription": "x",
                    "strCategoryThumb": "x"
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
        "query": [
            "i",
            "c",
            "a"
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "meals": [
                {
                    "idMeal": "x",
                    "strMeal": "x",
                    "strMealThumb": "x"
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
        "query": [],
        "auth": null,
        "status": 200,
        "sample": {
            "meals": [
                {
                    "dateModified": "x",
                    "idMeal": "x",
                    "strArea": "x",
                    "strCategory": "x",
                    "strCreativeCommonsConfirmed": "x",
                    "strDrinkAlternate": "x",
                    "strImageSource": "x",
                    "strIngredient1": "x",
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
                    "strIngredient2": "x",
                    "strIngredient20": "x",
                    "strIngredient3": "x",
                    "strIngredient4": "x",
                    "strIngredient5": "x",
                    "strIngredient6": "x",
                    "strIngredient7": "x",
                    "strIngredient8": "x",
                    "strIngredient9": "x",
                    "strInstructions": "x",
                    "strMeal": "x",
                    "strMealThumb": "x",
                    "strMeasure1": "x",
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
                    "strMeasure2": "x",
                    "strMeasure20": "x",
                    "strMeasure3": "x",
                    "strMeasure4": "x",
                    "strMeasure5": "x",
                    "strMeasure6": "x",
                    "strMeasure7": "x",
                    "strMeasure8": "x",
                    "strMeasure9": "x",
                    "strSource": "x",
                    "strTags": "x",
                    "strYoutube": "x"
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
        "query": [
            "c",
            "a",
            "i"
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "meals": [
                {
                    "strArea": "x",
                    "strCategory": "x",
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
        "query": [
            "i"
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "meals": [
                {
                    "dateModified": "x",
                    "idMeal": "x",
                    "strArea": "x",
                    "strCategory": "x",
                    "strCreativeCommonsConfirmed": "x",
                    "strDrinkAlternate": "x",
                    "strImageSource": "x",
                    "strIngredient1": "x",
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
                    "strIngredient2": "x",
                    "strIngredient20": "x",
                    "strIngredient3": "x",
                    "strIngredient4": "x",
                    "strIngredient5": "x",
                    "strIngredient6": "x",
                    "strIngredient7": "x",
                    "strIngredient8": "x",
                    "strIngredient9": "x",
                    "strInstructions": "x",
                    "strMeal": "x",
                    "strMealThumb": "x",
                    "strMeasure1": "x",
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
                    "strMeasure2": "x",
                    "strMeasure20": "x",
                    "strMeasure3": "x",
                    "strMeasure4": "x",
                    "strMeasure5": "x",
                    "strMeasure6": "x",
                    "strMeasure7": "x",
                    "strMeasure8": "x",
                    "strMeasure9": "x",
                    "strSource": "x",
                    "strTags": "x",
                    "strYoutube": "x"
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
        "query": [],
        "auth": [],
        "status": 200,
        "sample": {
            "meals": [
                {
                    "dateModified": "x",
                    "idMeal": "x",
                    "strArea": "x",
                    "strCategory": "x",
                    "strCreativeCommonsConfirmed": "x",
                    "strDrinkAlternate": "x",
                    "strImageSource": "x",
                    "strIngredient1": "x",
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
                    "strIngredient2": "x",
                    "strIngredient20": "x",
                    "strIngredient3": "x",
                    "strIngredient4": "x",
                    "strIngredient5": "x",
                    "strIngredient6": "x",
                    "strIngredient7": "x",
                    "strIngredient8": "x",
                    "strIngredient9": "x",
                    "strInstructions": "x",
                    "strMeal": "x",
                    "strMealThumb": "x",
                    "strMeasure1": "x",
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
                    "strMeasure2": "x",
                    "strMeasure20": "x",
                    "strMeasure3": "x",
                    "strMeasure4": "x",
                    "strMeasure5": "x",
                    "strMeasure6": "x",
                    "strMeasure7": "x",
                    "strMeasure8": "x",
                    "strMeasure9": "x",
                    "strSource": "x",
                    "strTags": "x",
                    "strYoutube": "x"
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
        "query": [],
        "auth": null,
        "status": 200,
        "sample": {
            "meals": [
                {
                    "dateModified": "x",
                    "idMeal": "x",
                    "strArea": "x",
                    "strCategory": "x",
                    "strCreativeCommonsConfirmed": "x",
                    "strDrinkAlternate": "x",
                    "strImageSource": "x",
                    "strIngredient1": "x",
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
                    "strIngredient2": "x",
                    "strIngredient20": "x",
                    "strIngredient3": "x",
                    "strIngredient4": "x",
                    "strIngredient5": "x",
                    "strIngredient6": "x",
                    "strIngredient7": "x",
                    "strIngredient8": "x",
                    "strIngredient9": "x",
                    "strInstructions": "x",
                    "strMeal": "x",
                    "strMealThumb": "x",
                    "strMeasure1": "x",
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
                    "strMeasure2": "x",
                    "strMeasure20": "x",
                    "strMeasure3": "x",
                    "strMeasure4": "x",
                    "strMeasure5": "x",
                    "strMeasure6": "x",
                    "strMeasure7": "x",
                    "strMeasure8": "x",
                    "strMeasure9": "x",
                    "strSource": "x",
                    "strTags": "x",
                    "strYoutube": "x"
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
        "query": [
            "s",
            "f"
        ],
        "auth": [],
        "status": 200,
        "sample": {
            "meals": [
                {
                    "dateModified": "x",
                    "idMeal": "x",
                    "strArea": "x",
                    "strCategory": "x",
                    "strCreativeCommonsConfirmed": "x",
                    "strDrinkAlternate": "x",
                    "strImageSource": "x",
                    "strIngredient1": "x",
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
                    "strIngredient2": "x",
                    "strIngredient20": "x",
                    "strIngredient3": "x",
                    "strIngredient4": "x",
                    "strIngredient5": "x",
                    "strIngredient6": "x",
                    "strIngredient7": "x",
                    "strIngredient8": "x",
                    "strIngredient9": "x",
                    "strInstructions": "x",
                    "strMeal": "x",
                    "strMealThumb": "x",
                    "strMeasure1": "x",
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
                    "strMeasure2": "x",
                    "strMeasure20": "x",
                    "strMeasure3": "x",
                    "strMeasure4": "x",
                    "strMeasure5": "x",
                    "strMeasure6": "x",
                    "strMeasure7": "x",
                    "strMeasure8": "x",
                    "strMeasure9": "x",
                    "strSource": "x",
                    "strTags": "x",
                    "strYoutube": "x"
                }
            ]
        },
        "idField": "id"
    }
];
(0, node_test_1.describe)('definition', () => {
    for (const point of PLAN) {
        (0, node_test_1.test)(point.entity + '.' + point.op + ' ' + point.method + ' ' + point.path, async (t) => {
            const control = (0, utility_1.isControlSkipped)('entityOp', point.entity + '.' + point.op, 'definition');
            if (control.skip) {
                t.skip(control.reason || 'skipped via sdk-test-control.json');
                return;
            }
            await (0, definition_runner_1.runDefinitionPoint)(__1.SDK, point);
        });
    }
});
//# sourceMappingURL=definition.test.js.map
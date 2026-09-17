# FreeMeal Lua SDK



The Lua SDK for the FreeMeal API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:Category()` — each with the same small set of operations (`list`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/free-meal-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("free-meal_sdk")

local client = sdk.new({
  apikey = os.getenv("FREE_MEAL_APIKEY"),
})
```

### 2. List category records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local categorys, err = client:Category():list()
if err then error(err) end

for _, item in ipairs(categorys) do
  print(item["idCategory"])
end
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local latests, err = client:Latest():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Latest():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
  },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
FREE_MEAL_TEST_LIVE=TRUE
FREE_MEAL_APIKEY=<your-key>
```

Then run:

```bash
cd lua && busted test/
```


## Reference

### FreeMealSDK

```lua
local sdk = require("free-meal_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### FreeMealSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
| `Category` | `(data) -> CategoryEntity` | Create a Category entity instance. |
| `Filter` | `(data) -> FilterEntity` | Create a Filter entity instance. |
| `Latest` | `(data) -> LatestEntity` | Create a Latest entity instance. |
| `List` | `(data) -> ListEntity` | Create a List entity instance. |
| `Lookup` | `(data) -> LookupEntity` | Create a Lookup entity instance. |
| `Random` | `(data) -> RandomEntity` | Create a Random entity instance. |
| `Randomselection` | `(data) -> RandomselectionEntity` | Create a Randomselection entity instance. |
| `Search` | `(data) -> SearchEntity` | Create a Search entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local category, err = client:Category():list()
    if err then error(err) end
    -- category is the record list

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

### Entities

#### Category

| Field | Description |
| --- | --- |
| `idCategory` | Unique category identifier |
| `strCategory` | Category name |
| `strCategoryDescription` | Category description |
| `strCategoryThumb` | URL to category thumbnail image |

Operations: List.

API path: `/categories.php`

#### Filter

| Field | Description |
| --- | --- |
| `idMeal` | Unique meal identifier |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |

Operations: List.

API path: `/filter.php`

#### Latest

| Field | Description |
| --- | --- |
| `dateModified` |  |
| `idMeal` | Unique meal identifier |
| `strArea` | Meal area/region |
| `strCategory` | Meal category |
| `strCreativeCommonsConfirmed` |  |
| `strDrinkAlternate` |  |
| `strImageSource` |  |
| `strIngredient1` |  |
| `strIngredient10` |  |
| `strIngredient11` |  |
| `strIngredient12` |  |
| `strIngredient13` |  |
| `strIngredient14` |  |
| `strIngredient15` |  |
| `strIngredient16` |  |
| `strIngredient17` |  |
| `strIngredient18` |  |
| `strIngredient19` |  |
| `strIngredient2` |  |
| `strIngredient20` |  |
| `strIngredient3` |  |
| `strIngredient4` |  |
| `strIngredient5` |  |
| `strIngredient6` |  |
| `strIngredient7` |  |
| `strIngredient8` |  |
| `strIngredient9` |  |
| `strInstructions` | Cooking instructions |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |
| `strMeasure1` |  |
| `strMeasure10` |  |
| `strMeasure11` |  |
| `strMeasure12` |  |
| `strMeasure13` |  |
| `strMeasure14` |  |
| `strMeasure15` |  |
| `strMeasure16` |  |
| `strMeasure17` |  |
| `strMeasure18` |  |
| `strMeasure19` |  |
| `strMeasure2` |  |
| `strMeasure20` |  |
| `strMeasure3` |  |
| `strMeasure4` |  |
| `strMeasure5` |  |
| `strMeasure6` |  |
| `strMeasure7` |  |
| `strMeasure8` |  |
| `strMeasure9` |  |
| `strSource` |  |
| `strTags` | Comma-separated tags |
| `strYoutube` | YouTube video URL |

Operations: List.

API path: `/latest.php`

#### List

| Field | Description |
| --- | --- |
| `strArea` |  |
| `strCategory` |  |
| `strIngredient` |  |

Operations: List.

API path: `/list.php`

#### Lookup

| Field | Description |
| --- | --- |
| `dateModified` |  |
| `idMeal` | Unique meal identifier |
| `strArea` | Meal area/region |
| `strCategory` | Meal category |
| `strCreativeCommonsConfirmed` |  |
| `strDrinkAlternate` |  |
| `strImageSource` |  |
| `strIngredient1` |  |
| `strIngredient10` |  |
| `strIngredient11` |  |
| `strIngredient12` |  |
| `strIngredient13` |  |
| `strIngredient14` |  |
| `strIngredient15` |  |
| `strIngredient16` |  |
| `strIngredient17` |  |
| `strIngredient18` |  |
| `strIngredient19` |  |
| `strIngredient2` |  |
| `strIngredient20` |  |
| `strIngredient3` |  |
| `strIngredient4` |  |
| `strIngredient5` |  |
| `strIngredient6` |  |
| `strIngredient7` |  |
| `strIngredient8` |  |
| `strIngredient9` |  |
| `strInstructions` | Cooking instructions |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |
| `strMeasure1` |  |
| `strMeasure10` |  |
| `strMeasure11` |  |
| `strMeasure12` |  |
| `strMeasure13` |  |
| `strMeasure14` |  |
| `strMeasure15` |  |
| `strMeasure16` |  |
| `strMeasure17` |  |
| `strMeasure18` |  |
| `strMeasure19` |  |
| `strMeasure2` |  |
| `strMeasure20` |  |
| `strMeasure3` |  |
| `strMeasure4` |  |
| `strMeasure5` |  |
| `strMeasure6` |  |
| `strMeasure7` |  |
| `strMeasure8` |  |
| `strMeasure9` |  |
| `strSource` |  |
| `strTags` | Comma-separated tags |
| `strYoutube` | YouTube video URL |

Operations: List.

API path: `/lookup.php`

#### Random

| Field | Description |
| --- | --- |
| `dateModified` |  |
| `idMeal` | Unique meal identifier |
| `strArea` | Meal area/region |
| `strCategory` | Meal category |
| `strCreativeCommonsConfirmed` |  |
| `strDrinkAlternate` |  |
| `strImageSource` |  |
| `strIngredient1` |  |
| `strIngredient10` |  |
| `strIngredient11` |  |
| `strIngredient12` |  |
| `strIngredient13` |  |
| `strIngredient14` |  |
| `strIngredient15` |  |
| `strIngredient16` |  |
| `strIngredient17` |  |
| `strIngredient18` |  |
| `strIngredient19` |  |
| `strIngredient2` |  |
| `strIngredient20` |  |
| `strIngredient3` |  |
| `strIngredient4` |  |
| `strIngredient5` |  |
| `strIngredient6` |  |
| `strIngredient7` |  |
| `strIngredient8` |  |
| `strIngredient9` |  |
| `strInstructions` | Cooking instructions |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |
| `strMeasure1` |  |
| `strMeasure10` |  |
| `strMeasure11` |  |
| `strMeasure12` |  |
| `strMeasure13` |  |
| `strMeasure14` |  |
| `strMeasure15` |  |
| `strMeasure16` |  |
| `strMeasure17` |  |
| `strMeasure18` |  |
| `strMeasure19` |  |
| `strMeasure2` |  |
| `strMeasure20` |  |
| `strMeasure3` |  |
| `strMeasure4` |  |
| `strMeasure5` |  |
| `strMeasure6` |  |
| `strMeasure7` |  |
| `strMeasure8` |  |
| `strMeasure9` |  |
| `strSource` |  |
| `strTags` | Comma-separated tags |
| `strYoutube` | YouTube video URL |

Operations: List.

API path: `/random.php`

#### Randomselection

| Field | Description |
| --- | --- |
| `dateModified` |  |
| `idMeal` | Unique meal identifier |
| `strArea` | Meal area/region |
| `strCategory` | Meal category |
| `strCreativeCommonsConfirmed` |  |
| `strDrinkAlternate` |  |
| `strImageSource` |  |
| `strIngredient1` |  |
| `strIngredient10` |  |
| `strIngredient11` |  |
| `strIngredient12` |  |
| `strIngredient13` |  |
| `strIngredient14` |  |
| `strIngredient15` |  |
| `strIngredient16` |  |
| `strIngredient17` |  |
| `strIngredient18` |  |
| `strIngredient19` |  |
| `strIngredient2` |  |
| `strIngredient20` |  |
| `strIngredient3` |  |
| `strIngredient4` |  |
| `strIngredient5` |  |
| `strIngredient6` |  |
| `strIngredient7` |  |
| `strIngredient8` |  |
| `strIngredient9` |  |
| `strInstructions` | Cooking instructions |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |
| `strMeasure1` |  |
| `strMeasure10` |  |
| `strMeasure11` |  |
| `strMeasure12` |  |
| `strMeasure13` |  |
| `strMeasure14` |  |
| `strMeasure15` |  |
| `strMeasure16` |  |
| `strMeasure17` |  |
| `strMeasure18` |  |
| `strMeasure19` |  |
| `strMeasure2` |  |
| `strMeasure20` |  |
| `strMeasure3` |  |
| `strMeasure4` |  |
| `strMeasure5` |  |
| `strMeasure6` |  |
| `strMeasure7` |  |
| `strMeasure8` |  |
| `strMeasure9` |  |
| `strSource` |  |
| `strTags` | Comma-separated tags |
| `strYoutube` | YouTube video URL |

Operations: List.

API path: `/randomselection.php`

#### Search

| Field | Description |
| --- | --- |
| `dateModified` |  |
| `idMeal` | Unique meal identifier |
| `strArea` | Meal area/region |
| `strCategory` | Meal category |
| `strCreativeCommonsConfirmed` |  |
| `strDrinkAlternate` |  |
| `strImageSource` |  |
| `strIngredient1` |  |
| `strIngredient10` |  |
| `strIngredient11` |  |
| `strIngredient12` |  |
| `strIngredient13` |  |
| `strIngredient14` |  |
| `strIngredient15` |  |
| `strIngredient16` |  |
| `strIngredient17` |  |
| `strIngredient18` |  |
| `strIngredient19` |  |
| `strIngredient2` |  |
| `strIngredient20` |  |
| `strIngredient3` |  |
| `strIngredient4` |  |
| `strIngredient5` |  |
| `strIngredient6` |  |
| `strIngredient7` |  |
| `strIngredient8` |  |
| `strIngredient9` |  |
| `strInstructions` | Cooking instructions |
| `strMeal` | Meal name |
| `strMealThumb` | URL to meal thumbnail image |
| `strMeasure1` |  |
| `strMeasure10` |  |
| `strMeasure11` |  |
| `strMeasure12` |  |
| `strMeasure13` |  |
| `strMeasure14` |  |
| `strMeasure15` |  |
| `strMeasure16` |  |
| `strMeasure17` |  |
| `strMeasure18` |  |
| `strMeasure19` |  |
| `strMeasure2` |  |
| `strMeasure20` |  |
| `strMeasure3` |  |
| `strMeasure4` |  |
| `strMeasure5` |  |
| `strMeasure6` |  |
| `strMeasure7` |  |
| `strMeasure8` |  |
| `strMeasure9` |  |
| `strSource` |  |
| `strTags` | Comma-separated tags |
| `strYoutube` | YouTube video URL |

Operations: List.

API path: `/search.php`



## Entities


### Category

Create an instance: `local category = client:Category(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `idCategory` | `string` | Unique category identifier |
| `strCategory` | `string` | Category name |
| `strCategoryDescription` | `string` | Category description |
| `strCategoryThumb` | `string` | URL to category thumbnail image |

#### Example: List

```lua
local categorys, err = client:Category():list()
```


### Filter

Create an instance: `local filter = client:Filter(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `idMeal` | `string` | Unique meal identifier |
| `strMeal` | `string` | Meal name |
| `strMealThumb` | `string` | URL to meal thumbnail image |

#### Example: List

```lua
local filters, err = client:Filter():list()
```


### Latest

Create an instance: `local latest = client:Latest(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dateModified` | `string` |  |
| `idMeal` | `string` | Unique meal identifier |
| `strArea` | `string` | Meal area/region |
| `strCategory` | `string` | Meal category |
| `strCreativeCommonsConfirmed` | `string` |  |
| `strDrinkAlternate` | `string` |  |
| `strImageSource` | `string` |  |
| `strIngredient1` | `string` |  |
| `strIngredient10` | `string` |  |
| `strIngredient11` | `string` |  |
| `strIngredient12` | `string` |  |
| `strIngredient13` | `string` |  |
| `strIngredient14` | `string` |  |
| `strIngredient15` | `string` |  |
| `strIngredient16` | `string` |  |
| `strIngredient17` | `string` |  |
| `strIngredient18` | `string` |  |
| `strIngredient19` | `string` |  |
| `strIngredient2` | `string` |  |
| `strIngredient20` | `string` |  |
| `strIngredient3` | `string` |  |
| `strIngredient4` | `string` |  |
| `strIngredient5` | `string` |  |
| `strIngredient6` | `string` |  |
| `strIngredient7` | `string` |  |
| `strIngredient8` | `string` |  |
| `strIngredient9` | `string` |  |
| `strInstructions` | `string` | Cooking instructions |
| `strMeal` | `string` | Meal name |
| `strMealThumb` | `string` | URL to meal thumbnail image |
| `strMeasure1` | `string` |  |
| `strMeasure10` | `string` |  |
| `strMeasure11` | `string` |  |
| `strMeasure12` | `string` |  |
| `strMeasure13` | `string` |  |
| `strMeasure14` | `string` |  |
| `strMeasure15` | `string` |  |
| `strMeasure16` | `string` |  |
| `strMeasure17` | `string` |  |
| `strMeasure18` | `string` |  |
| `strMeasure19` | `string` |  |
| `strMeasure2` | `string` |  |
| `strMeasure20` | `string` |  |
| `strMeasure3` | `string` |  |
| `strMeasure4` | `string` |  |
| `strMeasure5` | `string` |  |
| `strMeasure6` | `string` |  |
| `strMeasure7` | `string` |  |
| `strMeasure8` | `string` |  |
| `strMeasure9` | `string` |  |
| `strSource` | `string` |  |
| `strTags` | `string` | Comma-separated tags |
| `strYoutube` | `string` | YouTube video URL |

#### Example: List

```lua
local latests, err = client:Latest():list()
```


### List

Create an instance: `local list = client:List(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `strArea` | `string` |  |
| `strCategory` | `string` |  |
| `strIngredient` | `string` |  |

#### Example: List

```lua
local lists, err = client:List():list()
```


### Lookup

Create an instance: `local lookup = client:Lookup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dateModified` | `string` |  |
| `idMeal` | `string` | Unique meal identifier |
| `strArea` | `string` | Meal area/region |
| `strCategory` | `string` | Meal category |
| `strCreativeCommonsConfirmed` | `string` |  |
| `strDrinkAlternate` | `string` |  |
| `strImageSource` | `string` |  |
| `strIngredient1` | `string` |  |
| `strIngredient10` | `string` |  |
| `strIngredient11` | `string` |  |
| `strIngredient12` | `string` |  |
| `strIngredient13` | `string` |  |
| `strIngredient14` | `string` |  |
| `strIngredient15` | `string` |  |
| `strIngredient16` | `string` |  |
| `strIngredient17` | `string` |  |
| `strIngredient18` | `string` |  |
| `strIngredient19` | `string` |  |
| `strIngredient2` | `string` |  |
| `strIngredient20` | `string` |  |
| `strIngredient3` | `string` |  |
| `strIngredient4` | `string` |  |
| `strIngredient5` | `string` |  |
| `strIngredient6` | `string` |  |
| `strIngredient7` | `string` |  |
| `strIngredient8` | `string` |  |
| `strIngredient9` | `string` |  |
| `strInstructions` | `string` | Cooking instructions |
| `strMeal` | `string` | Meal name |
| `strMealThumb` | `string` | URL to meal thumbnail image |
| `strMeasure1` | `string` |  |
| `strMeasure10` | `string` |  |
| `strMeasure11` | `string` |  |
| `strMeasure12` | `string` |  |
| `strMeasure13` | `string` |  |
| `strMeasure14` | `string` |  |
| `strMeasure15` | `string` |  |
| `strMeasure16` | `string` |  |
| `strMeasure17` | `string` |  |
| `strMeasure18` | `string` |  |
| `strMeasure19` | `string` |  |
| `strMeasure2` | `string` |  |
| `strMeasure20` | `string` |  |
| `strMeasure3` | `string` |  |
| `strMeasure4` | `string` |  |
| `strMeasure5` | `string` |  |
| `strMeasure6` | `string` |  |
| `strMeasure7` | `string` |  |
| `strMeasure8` | `string` |  |
| `strMeasure9` | `string` |  |
| `strSource` | `string` |  |
| `strTags` | `string` | Comma-separated tags |
| `strYoutube` | `string` | YouTube video URL |

#### Example: List

```lua
local lookups, err = client:Lookup():list()
```


### Random

Create an instance: `local random = client:Random(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dateModified` | `string` |  |
| `idMeal` | `string` | Unique meal identifier |
| `strArea` | `string` | Meal area/region |
| `strCategory` | `string` | Meal category |
| `strCreativeCommonsConfirmed` | `string` |  |
| `strDrinkAlternate` | `string` |  |
| `strImageSource` | `string` |  |
| `strIngredient1` | `string` |  |
| `strIngredient10` | `string` |  |
| `strIngredient11` | `string` |  |
| `strIngredient12` | `string` |  |
| `strIngredient13` | `string` |  |
| `strIngredient14` | `string` |  |
| `strIngredient15` | `string` |  |
| `strIngredient16` | `string` |  |
| `strIngredient17` | `string` |  |
| `strIngredient18` | `string` |  |
| `strIngredient19` | `string` |  |
| `strIngredient2` | `string` |  |
| `strIngredient20` | `string` |  |
| `strIngredient3` | `string` |  |
| `strIngredient4` | `string` |  |
| `strIngredient5` | `string` |  |
| `strIngredient6` | `string` |  |
| `strIngredient7` | `string` |  |
| `strIngredient8` | `string` |  |
| `strIngredient9` | `string` |  |
| `strInstructions` | `string` | Cooking instructions |
| `strMeal` | `string` | Meal name |
| `strMealThumb` | `string` | URL to meal thumbnail image |
| `strMeasure1` | `string` |  |
| `strMeasure10` | `string` |  |
| `strMeasure11` | `string` |  |
| `strMeasure12` | `string` |  |
| `strMeasure13` | `string` |  |
| `strMeasure14` | `string` |  |
| `strMeasure15` | `string` |  |
| `strMeasure16` | `string` |  |
| `strMeasure17` | `string` |  |
| `strMeasure18` | `string` |  |
| `strMeasure19` | `string` |  |
| `strMeasure2` | `string` |  |
| `strMeasure20` | `string` |  |
| `strMeasure3` | `string` |  |
| `strMeasure4` | `string` |  |
| `strMeasure5` | `string` |  |
| `strMeasure6` | `string` |  |
| `strMeasure7` | `string` |  |
| `strMeasure8` | `string` |  |
| `strMeasure9` | `string` |  |
| `strSource` | `string` |  |
| `strTags` | `string` | Comma-separated tags |
| `strYoutube` | `string` | YouTube video URL |

#### Example: List

```lua
local randoms, err = client:Random():list()
```


### Randomselection

Create an instance: `local randomselection = client:Randomselection(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dateModified` | `string` |  |
| `idMeal` | `string` | Unique meal identifier |
| `strArea` | `string` | Meal area/region |
| `strCategory` | `string` | Meal category |
| `strCreativeCommonsConfirmed` | `string` |  |
| `strDrinkAlternate` | `string` |  |
| `strImageSource` | `string` |  |
| `strIngredient1` | `string` |  |
| `strIngredient10` | `string` |  |
| `strIngredient11` | `string` |  |
| `strIngredient12` | `string` |  |
| `strIngredient13` | `string` |  |
| `strIngredient14` | `string` |  |
| `strIngredient15` | `string` |  |
| `strIngredient16` | `string` |  |
| `strIngredient17` | `string` |  |
| `strIngredient18` | `string` |  |
| `strIngredient19` | `string` |  |
| `strIngredient2` | `string` |  |
| `strIngredient20` | `string` |  |
| `strIngredient3` | `string` |  |
| `strIngredient4` | `string` |  |
| `strIngredient5` | `string` |  |
| `strIngredient6` | `string` |  |
| `strIngredient7` | `string` |  |
| `strIngredient8` | `string` |  |
| `strIngredient9` | `string` |  |
| `strInstructions` | `string` | Cooking instructions |
| `strMeal` | `string` | Meal name |
| `strMealThumb` | `string` | URL to meal thumbnail image |
| `strMeasure1` | `string` |  |
| `strMeasure10` | `string` |  |
| `strMeasure11` | `string` |  |
| `strMeasure12` | `string` |  |
| `strMeasure13` | `string` |  |
| `strMeasure14` | `string` |  |
| `strMeasure15` | `string` |  |
| `strMeasure16` | `string` |  |
| `strMeasure17` | `string` |  |
| `strMeasure18` | `string` |  |
| `strMeasure19` | `string` |  |
| `strMeasure2` | `string` |  |
| `strMeasure20` | `string` |  |
| `strMeasure3` | `string` |  |
| `strMeasure4` | `string` |  |
| `strMeasure5` | `string` |  |
| `strMeasure6` | `string` |  |
| `strMeasure7` | `string` |  |
| `strMeasure8` | `string` |  |
| `strMeasure9` | `string` |  |
| `strSource` | `string` |  |
| `strTags` | `string` | Comma-separated tags |
| `strYoutube` | `string` | YouTube video URL |

#### Example: List

```lua
local randomselections, err = client:Randomselection():list()
```


### Search

Create an instance: `local search = client:Search(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `dateModified` | `string` |  |
| `idMeal` | `string` | Unique meal identifier |
| `strArea` | `string` | Meal area/region |
| `strCategory` | `string` | Meal category |
| `strCreativeCommonsConfirmed` | `string` |  |
| `strDrinkAlternate` | `string` |  |
| `strImageSource` | `string` |  |
| `strIngredient1` | `string` |  |
| `strIngredient10` | `string` |  |
| `strIngredient11` | `string` |  |
| `strIngredient12` | `string` |  |
| `strIngredient13` | `string` |  |
| `strIngredient14` | `string` |  |
| `strIngredient15` | `string` |  |
| `strIngredient16` | `string` |  |
| `strIngredient17` | `string` |  |
| `strIngredient18` | `string` |  |
| `strIngredient19` | `string` |  |
| `strIngredient2` | `string` |  |
| `strIngredient20` | `string` |  |
| `strIngredient3` | `string` |  |
| `strIngredient4` | `string` |  |
| `strIngredient5` | `string` |  |
| `strIngredient6` | `string` |  |
| `strIngredient7` | `string` |  |
| `strIngredient8` | `string` |  |
| `strIngredient9` | `string` |  |
| `strInstructions` | `string` | Cooking instructions |
| `strMeal` | `string` | Meal name |
| `strMealThumb` | `string` | URL to meal thumbnail image |
| `strMeasure1` | `string` |  |
| `strMeasure10` | `string` |  |
| `strMeasure11` | `string` |  |
| `strMeasure12` | `string` |  |
| `strMeasure13` | `string` |  |
| `strMeasure14` | `string` |  |
| `strMeasure15` | `string` |  |
| `strMeasure16` | `string` |  |
| `strMeasure17` | `string` |  |
| `strMeasure18` | `string` |  |
| `strMeasure19` | `string` |  |
| `strMeasure2` | `string` |  |
| `strMeasure20` | `string` |  |
| `strMeasure3` | `string` |  |
| `strMeasure4` | `string` |  |
| `strMeasure5` | `string` |  |
| `strMeasure6` | `string` |  |
| `strMeasure7` | `string` |  |
| `strMeasure8` | `string` |  |
| `strMeasure9` | `string` |  |
| `strSource` | `string` |  |
| `strTags` | `string` | Comma-separated tags |
| `strYoutube` | `string` | YouTube video URL |

#### Example: List

```lua
local searchs, err = client:Search():list()
```

## Features

This SDK ships 4 optional features. Each is **inactive until you
switch it on**, so an SDK you have not configured behaves exactly as if none of
them existed — no retries, no cache, no logging, no measurable overhead.

Activate a feature by name in the client options, alongside the options shown
above:

| Feature | What it does |
|---|---|
| [`ratelimit`](#ratelimit) | Client-side rate limiting via a token bucket |
| [`retry`](#retry) | Automatic retry of transient failures with exponential backoff |
| [`test`](#test) | In-memory mock transport for testing without a live server |
| [`timeout`](#timeout) | Per-request timeout with transport abort |

> **Order matters for `ratelimit`, `retry`, `timeout`.** These wrap the
> transport, so each one wraps whatever is already installed: the order you
> activate them in IS the nesting order. Activating them as an ordered list
> rather than a map is what fixes that order.

### ratelimit

Client-side rate limiting via a token bucket.

| Option | Default |
|---|---|
| `active` | `false` |
| `burst` | `5` |
| `rate` | `5` |

Set `feature.ratelimit.active` to enable it, then override any of the options above.

`ratelimit` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### retry

Automatic retry of transient failures with exponential backoff.

| Option | Default |
|---|---|
| `active` | `false` |
| `factor` | `2` |
| `maxDelay` | `2000` |
| `minDelay` | `50` |
| `retries` | `2` |
| `statuses` | `[408, 425, 429, 500, 502, 503, 504]` |

Set `feature.retry.active` to enable it, then override any of the options above.

`retry` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.

### test

In-memory mock transport for testing without a live server.

| Option | Default |
|---|---|
| `active` | `false` |

Set `feature.test.active` to enable it, then override any of the options above.

### timeout

Per-request timeout with transport abort.

| Option | Default |
|---|---|
| `active` | `false` |
| `ms` | `30000` |

Set `feature.timeout.active` to enable it, then override any of the options above.

`timeout` wraps the transport, so its position among the other
transport features decides what it sees. A feature activated later wraps one
activated earlier.


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **RatelimitFeature**: Client-side rate limiting via a token bucket
- **RetryFeature**: Automatic retry of transient failures with exponential backoff
- **TestFeature**: In-memory mock transport for testing without a live server
- **TimeoutFeature**: Per-request timeout with transport abort

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── free-meal_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── schema.lua               -- Generated option + entity specs
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`free-meal_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local latest = client:Latest()
latest:list()

-- latest:data_get() now returns the latest data from the last list
-- latest:match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
